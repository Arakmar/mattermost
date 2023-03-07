// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	store "github.com/mattermost/mattermost-server/server/v7/channels/store"
	model "github.com/mattermost/mattermost-server/server/v7/model"
	mock "github.com/stretchr/testify/mock"
)

// ThreadStore is an autogenerated mock type for the ThreadStore type
type ThreadStore struct {
	mock.Mock
}

// DeleteMembershipForUser provides a mock function with given fields: userId, postID
func (_m *ThreadStore) DeleteMembershipForUser(userId string, postID string) error {
	ret := _m.Called(userId, postID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, postID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedRows provides a mock function with given fields: limit
func (_m *ThreadStore) DeleteOrphanedRows(limit int) (int64, error) {
	ret := _m.Called(limit)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *ThreadStore) Get(id string) (*model.Thread, error) {
	ret := _m.Called(id)

	var r0 *model.Thread
	if rf, ok := ret.Get(0).(func(string) *model.Thread); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembershipForUser provides a mock function with given fields: userId, postID
func (_m *ThreadStore) GetMembershipForUser(userId string, postID string) (*model.ThreadMembership, error) {
	ret := _m.Called(userId, postID)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(string, string) *model.ThreadMembership); ok {
		r0 = rf(userId, postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembershipsForUser provides a mock function with given fields: userId, teamID
func (_m *ThreadStore) GetMembershipsForUser(userId string, teamID string) ([]*model.ThreadMembership, error) {
	ret := _m.Called(userId, teamID)

	var r0 []*model.ThreadMembership
	if rf, ok := ret.Get(0).(func(string, string) []*model.ThreadMembership); ok {
		r0 = rf(userId, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamsUnreadForUser provides a mock function with given fields: userID, teamIDs, includeUrgentMentionCount
func (_m *ThreadStore) GetTeamsUnreadForUser(userID string, teamIDs []string, includeUrgentMentionCount bool) (map[string]*model.TeamUnread, error) {
	ret := _m.Called(userID, teamIDs, includeUrgentMentionCount)

	var r0 map[string]*model.TeamUnread
	if rf, ok := ret.Get(0).(func(string, []string, bool) map[string]*model.TeamUnread); ok {
		r0 = rf(userID, teamIDs, includeUrgentMentionCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.TeamUnread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, bool) error); ok {
		r1 = rf(userID, teamIDs, includeUrgentMentionCount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadFollowers provides a mock function with given fields: threadID, fetchOnlyActive
func (_m *ThreadStore) GetThreadFollowers(threadID string, fetchOnlyActive bool) ([]string, error) {
	ret := _m.Called(threadID, fetchOnlyActive)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, bool) []string); ok {
		r0 = rf(threadID, fetchOnlyActive)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(threadID, fetchOnlyActive)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadForUser provides a mock function with given fields: threadMembership, extended, postPriorityIsEnabled
func (_m *ThreadStore) GetThreadForUser(threadMembership *model.ThreadMembership, extended bool, postPriorityIsEnabled bool) (*model.ThreadResponse, error) {
	ret := _m.Called(threadMembership, extended, postPriorityIsEnabled)

	var r0 *model.ThreadResponse
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership, bool, bool) *model.ThreadResponse); ok {
		r0 = rf(threadMembership, extended, postPriorityIsEnabled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ThreadMembership, bool, bool) error); ok {
		r1 = rf(threadMembership, extended, postPriorityIsEnabled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadUnreadReplyCount provides a mock function with given fields: threadMembership
func (_m *ThreadStore) GetThreadUnreadReplyCount(threadMembership *model.ThreadMembership) (int64, error) {
	ret := _m.Called(threadMembership)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) int64); ok {
		r0 = rf(threadMembership)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(threadMembership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadsForUser provides a mock function with given fields: userId, teamID, opts
func (_m *ThreadStore) GetThreadsForUser(userId string, teamID string, opts model.GetUserThreadsOpts) ([]*model.ThreadResponse, error) {
	ret := _m.Called(userId, teamID, opts)

	var r0 []*model.ThreadResponse
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) []*model.ThreadResponse); ok {
		r0 = rf(userId, teamID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userId, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopThreadsForTeamSince provides a mock function with given fields: teamID, userID, since, offset, limit
func (_m *ThreadStore) GetTopThreadsForTeamSince(teamID string, userID string, since int64, offset int, limit int) (*model.TopThreadList, error) {
	ret := _m.Called(teamID, userID, since, offset, limit)

	var r0 *model.TopThreadList
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) *model.TopThreadList); ok {
		r0 = rf(teamID, userID, since, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TopThreadList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int64, int, int) error); ok {
		r1 = rf(teamID, userID, since, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopThreadsForUserSince provides a mock function with given fields: teamID, userID, since, offset, limit
func (_m *ThreadStore) GetTopThreadsForUserSince(teamID string, userID string, since int64, offset int, limit int) (*model.TopThreadList, error) {
	ret := _m.Called(teamID, userID, since, offset, limit)

	var r0 *model.TopThreadList
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) *model.TopThreadList); ok {
		r0 = rf(teamID, userID, since, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TopThreadList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int64, int, int) error); ok {
		r1 = rf(teamID, userID, since, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalThreads provides a mock function with given fields: userId, teamID, opts
func (_m *ThreadStore) GetTotalThreads(userId string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userId, teamID, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userId, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userId, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUnreadMentions provides a mock function with given fields: userId, teamID, opts
func (_m *ThreadStore) GetTotalUnreadMentions(userId string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userId, teamID, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userId, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userId, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUnreadThreads provides a mock function with given fields: userId, teamID, opts
func (_m *ThreadStore) GetTotalUnreadThreads(userId string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userId, teamID, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userId, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userId, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUnreadUrgentMentions provides a mock function with given fields: userId, teamID, opts
func (_m *ThreadStore) GetTotalUnreadUrgentMentions(userId string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userId, teamID, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userId, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userId, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaintainMembership provides a mock function with given fields: userID, postID, opts
func (_m *ThreadStore) MaintainMembership(userID string, postID string, opts store.ThreadMembershipOpts) (*model.ThreadMembership, error) {
	ret := _m.Called(userID, postID, opts)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(string, string, store.ThreadMembershipOpts) *model.ThreadMembership); ok {
		r0 = rf(userID, postID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, store.ThreadMembershipOpts) error); ok {
		r1 = rf(userID, postID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkAllAsRead provides a mock function with given fields: userID, threadIds
func (_m *ThreadStore) MarkAllAsRead(userID string, threadIds []string) error {
	ret := _m.Called(userID, threadIds)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(userID, threadIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllAsReadByChannels provides a mock function with given fields: userID, channelIDs
func (_m *ThreadStore) MarkAllAsReadByChannels(userID string, channelIDs []string) error {
	ret := _m.Called(userID, channelIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(userID, channelIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllAsReadByTeam provides a mock function with given fields: userID, teamID
func (_m *ThreadStore) MarkAllAsReadByTeam(userID string, teamID string) error {
	ret := _m.Called(userID, teamID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, teamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAsRead provides a mock function with given fields: userID, threadID, timestamp
func (_m *ThreadStore) MarkAsRead(userID string, threadID string, timestamp int64) error {
	ret := _m.Called(userID, threadID, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(userID, threadID, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PermanentDeleteBatchForRetentionPolicies provides a mock function with given fields: now, globalPolicyEndTime, limit, cursor
func (_m *ThreadStore) PermanentDeleteBatchForRetentionPolicies(now int64, globalPolicyEndTime int64, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	ret := _m.Called(now, globalPolicyEndTime, limit, cursor)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) int64); ok {
		r0 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 model.RetentionPolicyCursor
	if rf, ok := ret.Get(1).(func(int64, int64, int64, model.RetentionPolicyCursor) model.RetentionPolicyCursor); ok {
		r1 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r1 = ret.Get(1).(model.RetentionPolicyCursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int64, int64, int64, model.RetentionPolicyCursor) error); ok {
		r2 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PermanentDeleteBatchThreadMembershipsForRetentionPolicies provides a mock function with given fields: now, globalPolicyEndTime, limit, cursor
func (_m *ThreadStore) PermanentDeleteBatchThreadMembershipsForRetentionPolicies(now int64, globalPolicyEndTime int64, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	ret := _m.Called(now, globalPolicyEndTime, limit, cursor)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) int64); ok {
		r0 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 model.RetentionPolicyCursor
	if rf, ok := ret.Get(1).(func(int64, int64, int64, model.RetentionPolicyCursor) model.RetentionPolicyCursor); ok {
		r1 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r1 = ret.Get(1).(model.RetentionPolicyCursor)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int64, int64, int64, model.RetentionPolicyCursor) error); ok {
		r2 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateMembership provides a mock function with given fields: membership
func (_m *ThreadStore) UpdateMembership(membership *model.ThreadMembership) (*model.ThreadMembership, error) {
	ret := _m.Called(membership)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) *model.ThreadMembership); ok {
		r0 = rf(membership)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(membership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
