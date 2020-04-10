// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "c19/security/model"

// SecurityService is an autogenerated mock type for the SecurityService type
type SecurityService struct {
	mock.Mock
}

// CanManage provides a mock function with given fields: ctx
func (_m *SecurityService) CanManage(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CanRead provides a mock function with given fields: ctx
func (_m *SecurityService) CanRead(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CanWrite provides a mock function with given fields: ctx
func (_m *SecurityService) CanWrite(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ChangePassword provides a mock function with given fields: ctx, request
func (_m *SecurityService) ChangePassword(ctx context.Context, request model.LoginRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.LoginRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, request
func (_m *SecurityService) CreateUser(ctx context.Context, request model.UserCreateRequest) (string, error) {
	ret := _m.Called(ctx, request)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, model.UserCreateRequest) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.UserCreateRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: request
func (_m *SecurityService) Login(request model.LoginRequest) (model.LoginResponse, error) {
	ret := _m.Called(request)

	var r0 model.LoginResponse
	if rf, ok := ret.Get(0).(func(model.LoginRequest) model.LoginResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(model.LoginResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.LoginRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx
func (_m *SecurityService) Logout(ctx context.Context) {
	_m.Called(ctx)
}