// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "golang-clean-arch-hateoas-example/domain"

	mock "github.com/stretchr/testify/mock"
)

// ItemRepository is an autogenerated mock type for the ItemRepository type
type ItemRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: item
func (_m *ItemRepository) Create(item domain.Item) (*domain.Item, error) {
	ret := _m.Called(item)

	var r0 *domain.Item
	if rf, ok := ret.Get(0).(func(domain.Item) *domain.Item); ok {
		r0 = rf(item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Item) error); ok {
		r1 = rf(item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ItemRepository) Delete(id int64) (*domain.Item, error) {
	ret := _m.Called(id)

	var r0 *domain.Item
	if rf, ok := ret.Get(0).(func(int64) *domain.Item); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: sigla
func (_m *ItemRepository) Fetch(sigla string) (*[]domain.Item, error) {
	ret := _m.Called(sigla)

	var r0 *[]domain.Item
	if rf, ok := ret.Get(0).(func(string) *[]domain.Item); ok {
		r0 = rf(sigla)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sigla)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *ItemRepository) GetByID(id int64) (*domain.Item, error) {
	ret := _m.Called(id)

	var r0 *domain.Item
	if rf, ok := ret.Get(0).(func(int64) *domain.Item); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: item, id
func (_m *ItemRepository) Update(item domain.Item, id int64) (*domain.Item, error) {
	ret := _m.Called(item, id)

	var r0 *domain.Item
	if rf, ok := ret.Get(0).(func(domain.Item, int64) *domain.Item); ok {
		r0 = rf(item, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Item, int64) error); ok {
		r1 = rf(item, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
