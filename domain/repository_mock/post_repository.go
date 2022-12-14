// Code generated by mockery v2.14.0. DO NOT EDIT.

package repository_mock

import (
	context "context"

	entity "github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// PostRepository is an autogenerated mock type for the PostRepository type
type PostRepository struct {
	mock.Mock
}

// CreatePost provides a mock function with given fields: ctx, post
func (_m *PostRepository) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	ret := _m.Called(ctx, post)

	var r0 *entity.Post
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Post) *entity.Post); ok {
		r0 = rf(ctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Post) error); ok {
		r1 = rf(ctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePost provides a mock function with given fields: ctx, id
func (_m *PostRepository) DeletePost(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPost provides a mock function with given fields: ctx, id
func (_m *PostRepository) GetPost(ctx context.Context, id int) (*entity.Post, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Post
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Post); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPosts provides a mock function with given fields: ctx
func (_m *PostRepository) ListPosts(ctx context.Context) ([]*entity.Post, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.Post
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Post); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePost provides a mock function with given fields: ctx, post
func (_m *PostRepository) UpdatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	ret := _m.Called(ctx, post)

	var r0 *entity.Post
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Post) *entity.Post); ok {
		r0 = rf(ctx, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Post) error); ok {
		r1 = rf(ctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostRepository creates a new instance of PostRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostRepository(t mockConstructorTestingTNewPostRepository) *PostRepository {
	mock := &PostRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
