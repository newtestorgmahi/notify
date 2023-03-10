// Code generated by mockery v2.14.0. DO NOT EDIT.

package wechat

import (
	message "github.com/silenceper/wechat/v2/officialaccount/message"
	mock "github.com/stretchr/testify/mock"
)

// mockWechatMessageManager is an autogenerated mock type for the wechatMessageManager type
type mockWechatMessageManager struct {
	mock.Mock
}

// Send provides a mock function with given fields: msg
func (_m *mockWechatMessageManager) Send(msg *message.CustomerMessage) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.CustomerMessage) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewMockWechatMessageManager interface {
	mock.TestingT
	Cleanup(func())
}

// newMockWechatMessageManager creates a new instance of mockWechatMessageManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockWechatMessageManager(t mockConstructorTestingTnewMockWechatMessageManager) *mockWechatMessageManager {
	mock := &mockWechatMessageManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
