// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/zi-bot/simple-gin-rest/models"
)

// AssetService is an autogenerated mock type for the AssetService type
type AssetService struct {
	mock.Mock
}

// CreateAsset provides a mock function with given fields: ctx, asset
func (_m *AssetService) CreateAsset(ctx context.Context, asset *models.Asset) error {
	ret := _m.Called(ctx, asset)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Asset) error); ok {
		r0 = rf(ctx, asset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAsset provides a mock function with given fields: ctx, id
func (_m *AssetService) DeleteAsset(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllAssets provides a mock function with given fields: ctx, pagination
func (_m *AssetService) GetAllAssets(ctx context.Context, pagination *models.Pagination) ([]*models.Asset, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []*models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Pagination) ([]*models.Asset, error)); ok {
		return rf(ctx, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Pagination) []*models.Asset); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Asset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Pagination) error); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAssetById provides a mock function with given fields: ctx, id
func (_m *AssetService) GetAssetById(ctx context.Context, id uint64) (*models.Asset, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*models.Asset, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Asset); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Asset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAsset provides a mock function with given fields: ctx, asset
func (_m *AssetService) UpdateAsset(ctx context.Context, asset *models.Asset) error {
	ret := _m.Called(ctx, asset)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Asset) error); ok {
		r0 = rf(ctx, asset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAssetService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAssetService creates a new instance of AssetService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAssetService(t mockConstructorTestingTNewAssetService) *AssetService {
	mock := &AssetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
