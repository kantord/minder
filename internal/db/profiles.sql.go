// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: profiles.sql

package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const countProfilesByEntityType = `-- name: CountProfilesByEntityType :many
SELECT COUNT(p.id) AS num_profiles, ep.entity AS profile_entity
FROM profiles AS p
         JOIN entity_profiles AS ep ON p.id = ep.profile_id
GROUP BY ep.entity
`

type CountProfilesByEntityTypeRow struct {
	NumProfiles   int64    `json:"num_profiles"`
	ProfileEntity Entities `json:"profile_entity"`
}

func (q *Queries) CountProfilesByEntityType(ctx context.Context) ([]CountProfilesByEntityTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, countProfilesByEntityType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CountProfilesByEntityTypeRow{}
	for rows.Next() {
		var i CountProfilesByEntityTypeRow
		if err := rows.Scan(&i.NumProfiles, &i.ProfileEntity); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const countProfilesByName = `-- name: CountProfilesByName :one
SELECT COUNT(*) AS num_named_profiles FROM profiles WHERE lower(name) = lower($1)
`

func (q *Queries) CountProfilesByName(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countProfilesByName, name)
	var num_named_profiles int64
	err := row.Scan(&num_named_profiles)
	return num_named_profiles, err
}

const createProfile = `-- name: CreateProfile :one
INSERT INTO profiles (  
    provider,
    project_id,
    remediate,
    alert,
    name,
    provider_id,
    subscription_id
) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, provider, project_id, remediate, alert, created_at, updated_at, provider_id, subscription_id
`

type CreateProfileParams struct {
	Provider       string         `json:"provider"`
	ProjectID      uuid.UUID      `json:"project_id"`
	Remediate      NullActionType `json:"remediate"`
	Alert          NullActionType `json:"alert"`
	Name           string         `json:"name"`
	ProviderID     uuid.UUID      `json:"provider_id"`
	SubscriptionID uuid.NullUUID  `json:"subscription_id"`
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, createProfile,
		arg.Provider,
		arg.ProjectID,
		arg.Remediate,
		arg.Alert,
		arg.Name,
		arg.ProviderID,
		arg.SubscriptionID,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.ProjectID,
		&i.Remediate,
		&i.Alert,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProviderID,
		&i.SubscriptionID,
	)
	return i, err
}

const createProfileForEntity = `-- name: CreateProfileForEntity :one
INSERT INTO entity_profiles (
    entity,
    profile_id,
    contextual_rules) VALUES ($1, $2, $3::jsonb) RETURNING id, entity, profile_id, contextual_rules, created_at, updated_at
`

type CreateProfileForEntityParams struct {
	Entity          Entities        `json:"entity"`
	ProfileID       uuid.UUID       `json:"profile_id"`
	ContextualRules json.RawMessage `json:"contextual_rules"`
}

func (q *Queries) CreateProfileForEntity(ctx context.Context, arg CreateProfileForEntityParams) (EntityProfile, error) {
	row := q.db.QueryRowContext(ctx, createProfileForEntity, arg.Entity, arg.ProfileID, arg.ContextualRules)
	var i EntityProfile
	err := row.Scan(
		&i.ID,
		&i.Entity,
		&i.ProfileID,
		&i.ContextualRules,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE id = $1 AND project_id = $2
`

type DeleteProfileParams struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (q *Queries) DeleteProfile(ctx context.Context, arg DeleteProfileParams) error {
	_, err := q.db.ExecContext(ctx, deleteProfile, arg.ID, arg.ProjectID)
	return err
}

const deleteProfileForEntity = `-- name: DeleteProfileForEntity :exec
DELETE FROM entity_profiles WHERE profile_id = $1 AND entity = $2
`

type DeleteProfileForEntityParams struct {
	ProfileID uuid.UUID `json:"profile_id"`
	Entity    Entities  `json:"entity"`
}

func (q *Queries) DeleteProfileForEntity(ctx context.Context, arg DeleteProfileForEntityParams) error {
	_, err := q.db.ExecContext(ctx, deleteProfileForEntity, arg.ProfileID, arg.Entity)
	return err
}

const deleteRuleInstantiation = `-- name: DeleteRuleInstantiation :exec
DELETE FROM entity_profile_rules WHERE entity_profile_id = $1 AND rule_type_id = $2
`

type DeleteRuleInstantiationParams struct {
	EntityProfileID uuid.UUID `json:"entity_profile_id"`
	RuleTypeID      uuid.UUID `json:"rule_type_id"`
}

func (q *Queries) DeleteRuleInstantiation(ctx context.Context, arg DeleteRuleInstantiationParams) error {
	_, err := q.db.ExecContext(ctx, deleteRuleInstantiation, arg.EntityProfileID, arg.RuleTypeID)
	return err
}

const getEntityProfileByProjectAndName = `-- name: GetEntityProfileByProjectAndName :many
SELECT profiles.id, name, provider, project_id, remediate, alert, profiles.created_at, profiles.updated_at, provider_id, subscription_id, entity_profiles.id, entity, profile_id, contextual_rules, entity_profiles.created_at, entity_profiles.updated_at FROM profiles JOIN entity_profiles ON profiles.id = entity_profiles.profile_id
WHERE profiles.project_id = $1 AND lower(profiles.name) = lower($2)
`

type GetEntityProfileByProjectAndNameParams struct {
	ProjectID uuid.UUID `json:"project_id"`
	Name      string    `json:"name"`
}

type GetEntityProfileByProjectAndNameRow struct {
	ID              uuid.UUID       `json:"id"`
	Name            string          `json:"name"`
	Provider        string          `json:"provider"`
	ProjectID       uuid.UUID       `json:"project_id"`
	Remediate       NullActionType  `json:"remediate"`
	Alert           NullActionType  `json:"alert"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	ProviderID      uuid.UUID       `json:"provider_id"`
	SubscriptionID  uuid.NullUUID   `json:"subscription_id"`
	ID_2            uuid.UUID       `json:"id_2"`
	Entity          Entities        `json:"entity"`
	ProfileID       uuid.UUID       `json:"profile_id"`
	ContextualRules json.RawMessage `json:"contextual_rules"`
	CreatedAt_2     time.Time       `json:"created_at_2"`
	UpdatedAt_2     time.Time       `json:"updated_at_2"`
}

func (q *Queries) GetEntityProfileByProjectAndName(ctx context.Context, arg GetEntityProfileByProjectAndNameParams) ([]GetEntityProfileByProjectAndNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getEntityProfileByProjectAndName, arg.ProjectID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEntityProfileByProjectAndNameRow{}
	for rows.Next() {
		var i GetEntityProfileByProjectAndNameRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Provider,
			&i.ProjectID,
			&i.Remediate,
			&i.Alert,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProviderID,
			&i.SubscriptionID,
			&i.ID_2,
			&i.Entity,
			&i.ProfileID,
			&i.ContextualRules,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfileByID = `-- name: GetProfileByID :one
SELECT id, name, provider, project_id, remediate, alert, created_at, updated_at, provider_id, subscription_id FROM profiles WHERE id = $1 AND project_id = $2
`

type GetProfileByIDParams struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (q *Queries) GetProfileByID(ctx context.Context, arg GetProfileByIDParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfileByID, arg.ID, arg.ProjectID)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.ProjectID,
		&i.Remediate,
		&i.Alert,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProviderID,
		&i.SubscriptionID,
	)
	return i, err
}

const getProfileByIDAndLock = `-- name: GetProfileByIDAndLock :one
SELECT id, name, provider, project_id, remediate, alert, created_at, updated_at, provider_id, subscription_id FROM profiles WHERE id = $1 AND project_id = $2 FOR UPDATE
`

type GetProfileByIDAndLockParams struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (q *Queries) GetProfileByIDAndLock(ctx context.Context, arg GetProfileByIDAndLockParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfileByIDAndLock, arg.ID, arg.ProjectID)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.ProjectID,
		&i.Remediate,
		&i.Alert,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProviderID,
		&i.SubscriptionID,
	)
	return i, err
}

const getProfileByNameAndLock = `-- name: GetProfileByNameAndLock :one
SELECT id, name, provider, project_id, remediate, alert, created_at, updated_at, provider_id, subscription_id FROM profiles WHERE lower(name) = lower($2) AND project_id = $1 FOR UPDATE
`

type GetProfileByNameAndLockParams struct {
	ProjectID uuid.UUID `json:"project_id"`
	Name      string    `json:"name"`
}

func (q *Queries) GetProfileByNameAndLock(ctx context.Context, arg GetProfileByNameAndLockParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfileByNameAndLock, arg.ProjectID, arg.Name)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.ProjectID,
		&i.Remediate,
		&i.Alert,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProviderID,
		&i.SubscriptionID,
	)
	return i, err
}

const getProfileByProjectAndID = `-- name: GetProfileByProjectAndID :many
SELECT profiles.id, name, provider, project_id, remediate, alert, profiles.created_at, profiles.updated_at, provider_id, subscription_id, entity_profiles.id, entity, profile_id, contextual_rules, entity_profiles.created_at, entity_profiles.updated_at FROM profiles JOIN entity_profiles ON profiles.id = entity_profiles.profile_id
WHERE profiles.project_id = $1 AND profiles.id = $2
`

type GetProfileByProjectAndIDParams struct {
	ProjectID uuid.UUID `json:"project_id"`
	ID        uuid.UUID `json:"id"`
}

type GetProfileByProjectAndIDRow struct {
	ID              uuid.UUID       `json:"id"`
	Name            string          `json:"name"`
	Provider        string          `json:"provider"`
	ProjectID       uuid.UUID       `json:"project_id"`
	Remediate       NullActionType  `json:"remediate"`
	Alert           NullActionType  `json:"alert"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	ProviderID      uuid.UUID       `json:"provider_id"`
	SubscriptionID  uuid.NullUUID   `json:"subscription_id"`
	ID_2            uuid.UUID       `json:"id_2"`
	Entity          Entities        `json:"entity"`
	ProfileID       uuid.UUID       `json:"profile_id"`
	ContextualRules json.RawMessage `json:"contextual_rules"`
	CreatedAt_2     time.Time       `json:"created_at_2"`
	UpdatedAt_2     time.Time       `json:"updated_at_2"`
}

func (q *Queries) GetProfileByProjectAndID(ctx context.Context, arg GetProfileByProjectAndIDParams) ([]GetProfileByProjectAndIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getProfileByProjectAndID, arg.ProjectID, arg.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProfileByProjectAndIDRow{}
	for rows.Next() {
		var i GetProfileByProjectAndIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Provider,
			&i.ProjectID,
			&i.Remediate,
			&i.Alert,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProviderID,
			&i.SubscriptionID,
			&i.ID_2,
			&i.Entity,
			&i.ProfileID,
			&i.ContextualRules,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfileForEntity = `-- name: GetProfileForEntity :one
SELECT id, entity, profile_id, contextual_rules, created_at, updated_at FROM entity_profiles WHERE profile_id = $1 AND entity = $2
`

type GetProfileForEntityParams struct {
	ProfileID uuid.UUID `json:"profile_id"`
	Entity    Entities  `json:"entity"`
}

func (q *Queries) GetProfileForEntity(ctx context.Context, arg GetProfileForEntityParams) (EntityProfile, error) {
	row := q.db.QueryRowContext(ctx, getProfileForEntity, arg.ProfileID, arg.Entity)
	var i EntityProfile
	err := row.Scan(
		&i.ID,
		&i.Entity,
		&i.ProfileID,
		&i.ContextualRules,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProfilesByProjectID = `-- name: ListProfilesByProjectID :many
SELECT profiles.id, name, provider, project_id, remediate, alert, profiles.created_at, profiles.updated_at, provider_id, subscription_id, entity_profiles.id, entity, profile_id, contextual_rules, entity_profiles.created_at, entity_profiles.updated_at FROM profiles JOIN entity_profiles ON profiles.id = entity_profiles.profile_id
WHERE profiles.project_id = $1
`

type ListProfilesByProjectIDRow struct {
	ID              uuid.UUID       `json:"id"`
	Name            string          `json:"name"`
	Provider        string          `json:"provider"`
	ProjectID       uuid.UUID       `json:"project_id"`
	Remediate       NullActionType  `json:"remediate"`
	Alert           NullActionType  `json:"alert"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	ProviderID      uuid.UUID       `json:"provider_id"`
	SubscriptionID  uuid.NullUUID   `json:"subscription_id"`
	ID_2            uuid.UUID       `json:"id_2"`
	Entity          Entities        `json:"entity"`
	ProfileID       uuid.UUID       `json:"profile_id"`
	ContextualRules json.RawMessage `json:"contextual_rules"`
	CreatedAt_2     time.Time       `json:"created_at_2"`
	UpdatedAt_2     time.Time       `json:"updated_at_2"`
}

func (q *Queries) ListProfilesByProjectID(ctx context.Context, projectID uuid.UUID) ([]ListProfilesByProjectIDRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfilesByProjectID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfilesByProjectIDRow{}
	for rows.Next() {
		var i ListProfilesByProjectIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Provider,
			&i.ProjectID,
			&i.Remediate,
			&i.Alert,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProviderID,
			&i.SubscriptionID,
			&i.ID_2,
			&i.Entity,
			&i.ProfileID,
			&i.ContextualRules,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProfilesInstantiatingRuleType = `-- name: ListProfilesInstantiatingRuleType :many
SELECT profiles.id, profiles.name, profiles.created_at FROM profiles
JOIN entity_profiles ON profiles.id = entity_profiles.profile_id 
JOIN entity_profile_rules ON entity_profiles.id = entity_profile_rules.entity_profile_id
WHERE entity_profile_rules.rule_type_id = $1
GROUP BY profiles.id
`

type ListProfilesInstantiatingRuleTypeRow struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// get profile information that instantiate a rule. This is done by joining the profiles with entity_profiles, then correlating those
// with entity_profile_rules. The rule_type_id is used to filter the results. Note that we only really care about the overal profile,
// so we only return the profile information. We also should group the profiles so that we don't get duplicates.
func (q *Queries) ListProfilesInstantiatingRuleType(ctx context.Context, ruleTypeID uuid.UUID) ([]ListProfilesInstantiatingRuleTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfilesInstantiatingRuleType, ruleTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfilesInstantiatingRuleTypeRow{}
	for rows.Next() {
		var i ListProfilesInstantiatingRuleTypeRow
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProfile = `-- name: UpdateProfile :one
UPDATE profiles SET
    remediate = $3,
    alert = $4,
    updated_at = NOW()
WHERE id = $1 AND project_id = $2 RETURNING id, name, provider, project_id, remediate, alert, created_at, updated_at, provider_id, subscription_id
`

type UpdateProfileParams struct {
	ID        uuid.UUID      `json:"id"`
	ProjectID uuid.UUID      `json:"project_id"`
	Remediate NullActionType `json:"remediate"`
	Alert     NullActionType `json:"alert"`
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, updateProfile,
		arg.ID,
		arg.ProjectID,
		arg.Remediate,
		arg.Alert,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.ProjectID,
		&i.Remediate,
		&i.Alert,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProviderID,
		&i.SubscriptionID,
	)
	return i, err
}

const upsertProfileForEntity = `-- name: UpsertProfileForEntity :one
INSERT INTO entity_profiles (
    entity,
    profile_id,
    contextual_rules) VALUES ($1, $2, $3::jsonb)
ON CONFLICT (entity, profile_id) DO UPDATE SET
    contextual_rules = $3::jsonb
RETURNING id, entity, profile_id, contextual_rules, created_at, updated_at
`

type UpsertProfileForEntityParams struct {
	Entity          Entities        `json:"entity"`
	ProfileID       uuid.UUID       `json:"profile_id"`
	ContextualRules json.RawMessage `json:"contextual_rules"`
}

func (q *Queries) UpsertProfileForEntity(ctx context.Context, arg UpsertProfileForEntityParams) (EntityProfile, error) {
	row := q.db.QueryRowContext(ctx, upsertProfileForEntity, arg.Entity, arg.ProfileID, arg.ContextualRules)
	var i EntityProfile
	err := row.Scan(
		&i.ID,
		&i.Entity,
		&i.ProfileID,
		&i.ContextualRules,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertRuleInstantiation = `-- name: UpsertRuleInstantiation :one
INSERT INTO entity_profile_rules (entity_profile_id, rule_type_id)
VALUES ($1, $2)
ON CONFLICT (entity_profile_id, rule_type_id) DO NOTHING RETURNING id, entity_profile_id, rule_type_id, created_at
`

type UpsertRuleInstantiationParams struct {
	EntityProfileID uuid.UUID `json:"entity_profile_id"`
	RuleTypeID      uuid.UUID `json:"rule_type_id"`
}

func (q *Queries) UpsertRuleInstantiation(ctx context.Context, arg UpsertRuleInstantiationParams) (EntityProfileRule, error) {
	row := q.db.QueryRowContext(ctx, upsertRuleInstantiation, arg.EntityProfileID, arg.RuleTypeID)
	var i EntityProfileRule
	err := row.Scan(
		&i.ID,
		&i.EntityProfileID,
		&i.RuleTypeID,
		&i.CreatedAt,
	)
	return i, err
}
