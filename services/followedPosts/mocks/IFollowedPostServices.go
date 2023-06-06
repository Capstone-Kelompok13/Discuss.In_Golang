// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "discusiin/dto"

	mock "github.com/stretchr/testify/mock"
)

// IFollowedPostServices is an autogenerated mock type for the IFollowedPostServices type
type IFollowedPostServices struct {
	mock.Mock
}

// AddFollowedPost provides a mock function with given fields: token, postID
func (_m *IFollowedPostServices) AddFollowedPost(token dto.Token, postID int) error {
	ret := _m.Called(token, postID)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.Token, int) error); ok {
		r0 = rf(token, postID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFollowedPost provides a mock function with given fields: token, postID
func (_m *IFollowedPostServices) DeleteFollowedPost(token dto.Token, postID int) error {
	ret := _m.Called(token, postID)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.Token, int) error); ok {
		r0 = rf(token, postID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllFollowedPost provides a mock function with given fields: token
func (_m *IFollowedPostServices) GetAllFollowedPost(token dto.Token) ([]dto.PublicFollowedPost, error) {
	ret := _m.Called(token)

	var r0 []dto.PublicFollowedPost
	if rf, ok := ret.Get(0).(func(dto.Token) []dto.PublicFollowedPost); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.PublicFollowedPost)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.Token) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIFollowedPostServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewIFollowedPostServices creates a new instance of IFollowedPostServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIFollowedPostServices(t mockConstructorTestingTNewIFollowedPostServices) *IFollowedPostServices {
	mock := &IFollowedPostServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
