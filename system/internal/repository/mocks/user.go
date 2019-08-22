// Code generated by MockGen. DO NOT EDIT.
// Source: system/internal/repository/user.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	repository "github.com/cortezaproject/corteza-server/system/internal/repository"
	types "github.com/cortezaproject/corteza-server/system/types"
	gomock "github.com/golang/mock/gomock"
	factory "github.com/titpetric/factory"
	io "io"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockUserRepository) With(ctx context.Context, db *factory.DB) repository.UserRepository {
	ret := m.ctrl.Call(m, "With", ctx, db)
	ret0, _ := ret[0].(repository.UserRepository)
	return ret0
}

// With indicates an expected call of With
func (mr *MockUserRepositoryMockRecorder) With(ctx, db interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockUserRepository)(nil).With), ctx, db)
}

// FindByEmail mocks base method
func (m *MockUserRepository) FindByEmail(email string) (*types.User, error) {
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockUserRepositoryMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), email)
}

// FindByUsername mocks base method
func (m *MockUserRepository) FindByUsername(username string) (*types.User, error) {
	ret := m.ctrl.Call(m, "FindByUsername", username)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername
func (mr *MockUserRepositoryMockRecorder) FindByUsername(username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockUserRepository)(nil).FindByUsername), username)
}

// FindByID mocks base method
func (m *MockUserRepository) FindByID(id uint64) (*types.User, error) {
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockUserRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), id)
}

// FindByIDs mocks base method
func (m *MockUserRepository) FindByIDs(id ...uint64) (types.UserSet, error) {
	varargs := []interface{}{}
	for _, a := range id {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByIDs", varargs...)
	ret0, _ := ret[0].(types.UserSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDs indicates an expected call of FindByIDs
func (mr *MockUserRepositoryMockRecorder) FindByIDs(id ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDs", reflect.TypeOf((*MockUserRepository)(nil).FindByIDs), id...)
}

// Find mocks base method
func (m *MockUserRepository) Find(filter types.UserFilter) (types.UserSet, types.UserFilter, error) {
	ret := m.ctrl.Call(m, "Find", filter)
	ret0, _ := ret[0].(types.UserSet)
	ret1, _ := ret[1].(types.UserFilter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockUserRepositoryMockRecorder) Find(filter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserRepository)(nil).Find), filter)
}

// Total mocks base method
func (m *MockUserRepository) Total() uint {
	ret := m.ctrl.Call(m, "Total")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Total indicates an expected call of Total
func (mr *MockUserRepositoryMockRecorder) Total() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Total", reflect.TypeOf((*MockUserRepository)(nil).Total))
}

// Create mocks base method
func (m *MockUserRepository) Create(mod *types.User) (*types.User, error) {
	ret := m.ctrl.Call(m, "Create", mod)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserRepositoryMockRecorder) Create(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), mod)
}

// Update mocks base method
func (m *MockUserRepository) Update(mod *types.User) (*types.User, error) {
	ret := m.ctrl.Call(m, "Update", mod)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUserRepositoryMockRecorder) Update(mod interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), mod)
}

// BindAvatar mocks base method
func (m *MockUserRepository) BindAvatar(user *types.User, avatar io.Reader) (*types.User, error) {
	ret := m.ctrl.Call(m, "BindAvatar", user, avatar)
	ret0, _ := ret[0].(*types.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindAvatar indicates an expected call of BindAvatar
func (mr *MockUserRepositoryMockRecorder) BindAvatar(user, avatar interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindAvatar", reflect.TypeOf((*MockUserRepository)(nil).BindAvatar), user, avatar)
}

// SuspendByID mocks base method
func (m *MockUserRepository) SuspendByID(id uint64) error {
	ret := m.ctrl.Call(m, "SuspendByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SuspendByID indicates an expected call of SuspendByID
func (mr *MockUserRepositoryMockRecorder) SuspendByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuspendByID", reflect.TypeOf((*MockUserRepository)(nil).SuspendByID), id)
}

// UnsuspendByID mocks base method
func (m *MockUserRepository) UnsuspendByID(id uint64) error {
	ret := m.ctrl.Call(m, "UnsuspendByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsuspendByID indicates an expected call of UnsuspendByID
func (mr *MockUserRepositoryMockRecorder) UnsuspendByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsuspendByID", reflect.TypeOf((*MockUserRepository)(nil).UnsuspendByID), id)
}

// DeleteByID mocks base method
func (m *MockUserRepository) DeleteByID(id uint64) error {
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockUserRepositoryMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockUserRepository)(nil).DeleteByID), id)
}
