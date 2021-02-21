// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "atwell/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: email
func (_m *UserUsecase) Login(email string) (string, error) {
	ret := _m.Called(email)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignIn provides a mock function with given fields: email
func (_m *UserUsecase) SignIn(email string) (domain.User, error) {
	ret := _m.Called(email)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}