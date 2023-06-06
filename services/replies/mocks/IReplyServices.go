// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "discusiin/dto"
	models "discusiin/models"

	mock "github.com/stretchr/testify/mock"
)

// IReplyServices is an autogenerated mock type for the IReplyServices type
type IReplyServices struct {
	mock.Mock
}

// CreateReply provides a mock function with given fields: reply, co, token
func (_m *IReplyServices) CreateReply(reply models.Reply, co int, token dto.Token) error {
	ret := _m.Called(reply, co, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Reply, int, dto.Token) error); ok {
		r0 = rf(reply, co, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReply provides a mock function with given fields: replyId, token
func (_m *IReplyServices) DeleteReply(replyId int, token dto.Token) error {
	ret := _m.Called(replyId, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, dto.Token) error); ok {
		r0 = rf(replyId, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllReply provides a mock function with given fields: commentId
func (_m *IReplyServices) GetAllReply(commentId int) ([]dto.PublicReply, error) {
	ret := _m.Called(commentId)

	var r0 []dto.PublicReply
	if rf, ok := ret.Get(0).(func(int) []dto.PublicReply); ok {
		r0 = rf(commentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.PublicReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(commentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReply provides a mock function with given fields: newReply, replyId, token
func (_m *IReplyServices) UpdateReply(newReply models.Reply, replyId int, token dto.Token) error {
	ret := _m.Called(newReply, replyId, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Reply, int, dto.Token) error); ok {
		r0 = rf(newReply, replyId, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIReplyServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewIReplyServices creates a new instance of IReplyServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIReplyServices(t mockConstructorTestingTNewIReplyServices) *IReplyServices {
	mock := &IReplyServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
