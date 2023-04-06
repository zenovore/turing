// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	models "github.com/caraml-dev/turing/api/turing/models"
	service "github.com/caraml-dev/turing/api/turing/service"
)

// EnsemblersService is an autogenerated mock type for the EnsemblersService type
type EnsemblersService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ensembler
func (_m *EnsemblersService) Delete(ensembler models.EnsemblerLike) error {
	ret := _m.Called(ensembler)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.EnsemblerLike) error); ok {
		r0 = rf(ensembler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: id, options
func (_m *EnsemblersService) FindByID(id models.ID, options service.EnsemblersFindByIDOptions) (models.EnsemblerLike, error) {
	ret := _m.Called(id, options)

	var r0 models.EnsemblerLike
	var r1 error
	if rf, ok := ret.Get(0).(func(models.ID, service.EnsemblersFindByIDOptions) (models.EnsemblerLike, error)); ok {
		return rf(id, options)
	}
	if rf, ok := ret.Get(0).(func(models.ID, service.EnsemblersFindByIDOptions) models.EnsemblerLike); ok {
		r0 = rf(id, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.EnsemblerLike)
		}
	}

	if rf, ok := ret.Get(1).(func(models.ID, service.EnsemblersFindByIDOptions) error); ok {
		r1 = rf(id, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: options
func (_m *EnsemblersService) List(options service.EnsemblersListOptions) (*service.PaginatedResults, error) {
	ret := _m.Called(options)

	var r0 *service.PaginatedResults
	var r1 error
	if rf, ok := ret.Get(0).(func(service.EnsemblersListOptions) (*service.PaginatedResults, error)); ok {
		return rf(options)
	}
	if rf, ok := ret.Get(0).(func(service.EnsemblersListOptions) *service.PaginatedResults); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.PaginatedResults)
		}
	}

	if rf, ok := ret.Get(1).(func(service.EnsemblersListOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ensembler
func (_m *EnsemblersService) Save(ensembler models.EnsemblerLike) (models.EnsemblerLike, error) {
	ret := _m.Called(ensembler)

	var r0 models.EnsemblerLike
	var r1 error
	if rf, ok := ret.Get(0).(func(models.EnsemblerLike) (models.EnsemblerLike, error)); ok {
		return rf(ensembler)
	}
	if rf, ok := ret.Get(0).(func(models.EnsemblerLike) models.EnsemblerLike); ok {
		r0 = rf(ensembler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.EnsemblerLike)
		}
	}

	if rf, ok := ret.Get(1).(func(models.EnsemblerLike) error); ok {
		r1 = rf(ensembler)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEnsemblersService interface {
	mock.TestingT
	Cleanup(func())
}

// NewEnsemblersService creates a new instance of EnsemblersService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEnsemblersService(t mockConstructorTestingTNewEnsemblersService) *EnsemblersService {
	mock := &EnsemblersService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
