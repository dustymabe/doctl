// Generated: please do not edit by hand

package mocks

import do "github.com/digitalocean/doctl/do"
import godo "github.com/digitalocean/godo"
import mock "github.com/stretchr/testify/mock"

// FloatingIPActionsService is an autogenerated mock type for the FloatingIPActionsService type
type FloatingIPActionsService struct {
	mock.Mock
}

// Assign provides a mock function with given fields: ip, dropletID
func (_m *FloatingIPActionsService) Assign(ip string, dropletID int) (*do.Action, error) {
	ret := _m.Called(ip, dropletID)

	var r0 *do.Action
	if rf, ok := ret.Get(0).(func(string, int) *do.Action); ok {
		r0 = rf(ip, dropletID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(ip, dropletID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ip, actionID
func (_m *FloatingIPActionsService) Get(ip string, actionID int) (*do.Action, error) {
	ret := _m.Called(ip, actionID)

	var r0 *do.Action
	if rf, ok := ret.Get(0).(func(string, int) *do.Action); ok {
		r0 = rf(ip, actionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(ip, actionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ip, opt
func (_m *FloatingIPActionsService) List(ip string, opt *godo.ListOptions) ([]do.Action, error) {
	ret := _m.Called(ip, opt)

	var r0 []do.Action
	if rf, ok := ret.Get(0).(func(string, *godo.ListOptions) []do.Action); ok {
		r0 = rf(ip, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.ListOptions) error); ok {
		r1 = rf(ip, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unassign provides a mock function with given fields: ip
func (_m *FloatingIPActionsService) Unassign(ip string) (*do.Action, error) {
	ret := _m.Called(ip)

	var r0 *do.Action
	if rf, ok := ret.Get(0).(func(string) *do.Action); ok {
		r0 = rf(ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Action)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
