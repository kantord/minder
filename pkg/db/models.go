// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"
)

type Group struct {
	ID             int32          `json:"id"`
	OrganizationID int32          `json:"organization_id"`
	Name           string         `json:"name"`
	Description    sql.NullString `json:"description"`
	IsProtected    bool           `json:"is_protected"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type Organization struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Company   string    `json:"company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID          int32     `json:"id"`
	GroupID     int32     `json:"group_id"`
	Name        string    `json:"name"`
	IsAdmin     bool      `json:"is_admin"`
	IsProtected bool      `json:"is_protected"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID                 int32          `json:"id"`
	RoleID             int32          `json:"role_id"`
	Email              sql.NullString `json:"email"`
	Username           string         `json:"username"`
	Password           string         `json:"password"`
	FirstName          sql.NullString `json:"first_name"`
	LastName           sql.NullString `json:"last_name"`
	IsProtected        bool           `json:"is_protected"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	MinTokenIssuedTime sql.NullTime   `json:"min_token_issued_time"`
}
