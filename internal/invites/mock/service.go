// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go
//
// Generated by this command:
//
//	mockgen -package mock_invites -destination=./mock/service.go -source=./service.go
//

// Package mock_invites is a generated GoMock package.
package mock_invites

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	auth "github.com/stacklok/minder/internal/auth"
	authz "github.com/stacklok/minder/internal/authz"
	server "github.com/stacklok/minder/internal/config/server"
	db "github.com/stacklok/minder/internal/db"
	events "github.com/stacklok/minder/internal/events"
	v1 "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockInviteService is a mock of InviteService interface.
type MockInviteService struct {
	ctrl     *gomock.Controller
	recorder *MockInviteServiceMockRecorder
}

// MockInviteServiceMockRecorder is the mock recorder for MockInviteService.
type MockInviteServiceMockRecorder struct {
	mock *MockInviteService
}

// NewMockInviteService creates a new mock instance.
func NewMockInviteService(ctrl *gomock.Controller) *MockInviteService {
	mock := &MockInviteService{ctrl: ctrl}
	mock.recorder = &MockInviteServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInviteService) EXPECT() *MockInviteServiceMockRecorder {
	return m.recorder
}

// RemoveInvite mocks base method.
func (m *MockInviteService) RemoveInvite(ctx context.Context, qtx db.Querier, idClient auth.Resolver, targetProject uuid.UUID, authzRole authz.Role, inviteeEmail string) (*v1.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveInvite", ctx, qtx, idClient, targetProject, authzRole, inviteeEmail)
	ret0, _ := ret[0].(*v1.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveInvite indicates an expected call of RemoveInvite.
func (mr *MockInviteServiceMockRecorder) RemoveInvite(ctx, qtx, idClient, targetProject, authzRole, inviteeEmail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveInvite", reflect.TypeOf((*MockInviteService)(nil).RemoveInvite), ctx, qtx, idClient, targetProject, authzRole, inviteeEmail)
}

// UpdateInvite mocks base method.
func (m *MockInviteService) UpdateInvite(ctx context.Context, qtx db.Querier, idClient auth.Resolver, eventsPub events.Publisher, emailConfig server.EmailConfig, targetProject uuid.UUID, authzRole authz.Role, inviteeEmail string) (*v1.Invitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInvite", ctx, qtx, idClient, eventsPub, emailConfig, targetProject, authzRole, inviteeEmail)
	ret0, _ := ret[0].(*v1.Invitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInvite indicates an expected call of UpdateInvite.
func (mr *MockInviteServiceMockRecorder) UpdateInvite(ctx, qtx, idClient, eventsPub, emailConfig, targetProject, authzRole, inviteeEmail any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInvite", reflect.TypeOf((*MockInviteService)(nil).UpdateInvite), ctx, qtx, idClient, eventsPub, emailConfig, targetProject, authzRole, inviteeEmail)
}
