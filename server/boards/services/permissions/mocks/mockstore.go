// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-server/server/v7/boards/services/permissions (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/server/v7/boards/model"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetBoard mocks base method.
func (m *MockStore) GetBoard(arg0 string) (*model.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoard", arg0)
	ret0, _ := ret[0].(*model.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoard indicates an expected call of GetBoard.
func (mr *MockStoreMockRecorder) GetBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoard", reflect.TypeOf((*MockStore)(nil).GetBoard), arg0)
}

// GetBoardHistory mocks base method.
func (m *MockStore) GetBoardHistory(arg0 string, arg1 model.QueryBoardHistoryOptions) ([]*model.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardHistory", arg0, arg1)
	ret0, _ := ret[0].([]*model.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoardHistory indicates an expected call of GetBoardHistory.
func (mr *MockStoreMockRecorder) GetBoardHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardHistory", reflect.TypeOf((*MockStore)(nil).GetBoardHistory), arg0, arg1)
}

// GetMemberForBoard mocks base method.
func (m *MockStore) GetMemberForBoard(arg0, arg1 string) (*model.BoardMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMemberForBoard", arg0, arg1)
	ret0, _ := ret[0].(*model.BoardMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberForBoard indicates an expected call of GetMemberForBoard.
func (mr *MockStoreMockRecorder) GetMemberForBoard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberForBoard", reflect.TypeOf((*MockStore)(nil).GetMemberForBoard), arg0, arg1)
}
