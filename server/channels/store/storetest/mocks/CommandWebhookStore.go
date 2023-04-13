// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/v8/public/model"
	mock "github.com/stretchr/testify/mock"
)

// CommandWebhookStore is an autogenerated mock type for the CommandWebhookStore type
type CommandWebhookStore struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields:
func (_m *CommandWebhookStore) Cleanup() {
	_m.Called()
}

// Get provides a mock function with given fields: id
func (_m *CommandWebhookStore) Get(id string) (*model.CommandWebhook, error) {
	ret := _m.Called(id)

	var r0 *model.CommandWebhook
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.CommandWebhook, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.CommandWebhook); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CommandWebhook)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: webhook
func (_m *CommandWebhookStore) Save(webhook *model.CommandWebhook) (*model.CommandWebhook, error) {
	ret := _m.Called(webhook)

	var r0 *model.CommandWebhook
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.CommandWebhook) (*model.CommandWebhook, error)); ok {
		return rf(webhook)
	}
	if rf, ok := ret.Get(0).(func(*model.CommandWebhook) *model.CommandWebhook); ok {
		r0 = rf(webhook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CommandWebhook)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.CommandWebhook) error); ok {
		r1 = rf(webhook)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TryUse provides a mock function with given fields: id, limit
func (_m *CommandWebhookStore) TryUse(id string, limit int) error {
	ret := _m.Called(id, limit)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(id, limit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCommandWebhookStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommandWebhookStore creates a new instance of CommandWebhookStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommandWebhookStore(t mockConstructorTestingTNewCommandWebhookStore) *CommandWebhookStore {
	mock := &CommandWebhookStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
