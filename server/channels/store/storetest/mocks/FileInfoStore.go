// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/v7/model"
	mock "github.com/stretchr/testify/mock"
)

// FileInfoStore is an autogenerated mock type for the FileInfoStore type
type FileInfoStore struct {
	mock.Mock
}

// AttachToPost provides a mock function with given fields: fileID, postID, creatorID
func (_m *FileInfoStore) AttachToPost(fileID string, postID string, creatorID string) error {
	ret := _m.Called(fileID, postID, creatorID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(fileID, postID, creatorID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearCaches provides a mock function with given fields:
func (_m *FileInfoStore) ClearCaches() {
	_m.Called()
}

// CountAll provides a mock function with given fields:
func (_m *FileInfoStore) CountAll() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteForPost provides a mock function with given fields: postID
func (_m *FileInfoStore) DeleteForPost(postID string) (string, error) {
	ret := _m.Called(postID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(postID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *FileInfoStore) Get(id string) (*model.FileInfo, error) {
	ret := _m.Called(id)

	var r0 *model.FileInfo
	if rf, ok := ret.Get(0).(func(string) *model.FileInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfo)
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

// GetByIds provides a mock function with given fields: ids
func (_m *FileInfoStore) GetByIds(ids []string) ([]*model.FileInfo, error) {
	ret := _m.Called(ids)

	var r0 []*model.FileInfo
	if rf, ok := ret.Get(0).(func([]string) []*model.FileInfo); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPath provides a mock function with given fields: path
func (_m *FileInfoStore) GetByPath(path string) (*model.FileInfo, error) {
	ret := _m.Called(path)

	var r0 *model.FileInfo
	if rf, ok := ret.Get(0).(func(string) *model.FileInfo); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFilesBatchForIndexing provides a mock function with given fields: startTime, startFileID, limit
func (_m *FileInfoStore) GetFilesBatchForIndexing(startTime int64, startFileID string, limit int) ([]*model.FileForIndexing, error) {
	ret := _m.Called(startTime, startFileID, limit)

	var r0 []*model.FileForIndexing
	if rf, ok := ret.Get(0).(func(int64, string, int) []*model.FileForIndexing); ok {
		r0 = rf(startTime, startFileID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FileForIndexing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, int) error); ok {
		r1 = rf(startTime, startFileID, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForPost provides a mock function with given fields: postID, readFromMaster, includeDeleted, allowFromCache
func (_m *FileInfoStore) GetForPost(postID string, readFromMaster bool, includeDeleted bool, allowFromCache bool) ([]*model.FileInfo, error) {
	ret := _m.Called(postID, readFromMaster, includeDeleted, allowFromCache)

	var r0 []*model.FileInfo
	if rf, ok := ret.Get(0).(func(string, bool, bool, bool) []*model.FileInfo); ok {
		r0 = rf(postID, readFromMaster, includeDeleted, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, bool, bool) error); ok {
		r1 = rf(postID, readFromMaster, includeDeleted, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForUser provides a mock function with given fields: userID
func (_m *FileInfoStore) GetForUser(userID string) ([]*model.FileInfo, error) {
	ret := _m.Called(userID)

	var r0 []*model.FileInfo
	if rf, ok := ret.Get(0).(func(string) []*model.FileInfo); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFromMaster provides a mock function with given fields: id
func (_m *FileInfoStore) GetFromMaster(id string) (*model.FileInfo, error) {
	ret := _m.Called(id)

	var r0 *model.FileInfo
	if rf, ok := ret.Get(0).(func(string) *model.FileInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfo)
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

// GetStorageUsage provides a mock function with given fields: allowFromCache, includeDeleted
func (_m *FileInfoStore) GetStorageUsage(allowFromCache bool, includeDeleted bool) (int64, error) {
	ret := _m.Called(allowFromCache, includeDeleted)

	var r0 int64
	if rf, ok := ret.Get(0).(func(bool, bool) int64); ok {
		r0 = rf(allowFromCache, includeDeleted)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool, bool) error); ok {
		r1 = rf(allowFromCache, includeDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUptoNSizeFileTime provides a mock function with given fields: n
func (_m *FileInfoStore) GetUptoNSizeFileTime(n int64) (int64, error) {
	ret := _m.Called(n)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithOptions provides a mock function with given fields: page, perPage, opt
func (_m *FileInfoStore) GetWithOptions(page int, perPage int, opt *model.GetFileInfosOptions) ([]*model.FileInfo, error) {
	ret := _m.Called(page, perPage, opt)

	var r0 []*model.FileInfo
	if rf, ok := ret.Get(0).(func(int, int, *model.GetFileInfosOptions) []*model.FileInfo); ok {
		r0 = rf(page, perPage, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, *model.GetFileInfosOptions) error); ok {
		r1 = rf(page, perPage, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvalidateFileInfosForPostCache provides a mock function with given fields: postID, deleted
func (_m *FileInfoStore) InvalidateFileInfosForPostCache(postID string, deleted bool) {
	_m.Called(postID, deleted)
}

// PermanentDelete provides a mock function with given fields: fileID
func (_m *FileInfoStore) PermanentDelete(fileID string) error {
	ret := _m.Called(fileID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *FileInfoStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	ret := _m.Called(endTime, limit)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(endTime, limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteByUser provides a mock function with given fields: userID
func (_m *FileInfoStore) PermanentDeleteByUser(userID string) (int64, error) {
	ret := _m.Called(userID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: info
func (_m *FileInfoStore) Save(info *model.FileInfo) (*model.FileInfo, error) {
	ret := _m.Called(info)

	var r0 *model.FileInfo
	if rf, ok := ret.Get(0).(func(*model.FileInfo) *model.FileInfo); ok {
		r0 = rf(info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.FileInfo) error); ok {
		r1 = rf(info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: paramsList, userID, teamID, page, perPage
func (_m *FileInfoStore) Search(paramsList []*model.SearchParams, userID string, teamID string, page int, perPage int) (*model.FileInfoList, error) {
	ret := _m.Called(paramsList, userID, teamID, page, perPage)

	var r0 *model.FileInfoList
	if rf, ok := ret.Get(0).(func([]*model.SearchParams, string, string, int, int) *model.FileInfoList); ok {
		r0 = rf(paramsList, userID, teamID, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfoList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*model.SearchParams, string, string, int, int) error); ok {
		r1 = rf(paramsList, userID, teamID, page, perPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetContent provides a mock function with given fields: fileID, content
func (_m *FileInfoStore) SetContent(fileID string, content string) error {
	ret := _m.Called(fileID, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(fileID, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: info
func (_m *FileInfoStore) Upsert(info *model.FileInfo) (*model.FileInfo, error) {
	ret := _m.Called(info)

	var r0 *model.FileInfo
	if rf, ok := ret.Get(0).(func(*model.FileInfo) *model.FileInfo); ok {
		r0 = rf(info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.FileInfo) error); ok {
		r1 = rf(info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
