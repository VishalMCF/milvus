// Code generated by mockery v2.14.0. DO NOT EDIT.

package querynode

import mock "github.com/stretchr/testify/mock"

// MockTSafeReplicaInterface is an autogenerated mock type for the TSafeReplicaInterface type
type MockTSafeReplicaInterface struct {
	mock.Mock
}

type MockTSafeReplicaInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTSafeReplicaInterface) EXPECT() *MockTSafeReplicaInterface_Expecter {
	return &MockTSafeReplicaInterface_Expecter{mock: &_m.Mock}
}

// Watch provides a mock function with given fields:
func (_m *MockTSafeReplicaInterface) Watch() Listener {
	ret := _m.Called()

	var r0 Listener
	if rf, ok := ret.Get(0).(func() Listener); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Listener)
		}
	}

	return r0
}

// MockTSafeReplicaInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockTSafeReplicaInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
func (_e *MockTSafeReplicaInterface_Expecter) Watch() *MockTSafeReplicaInterface_Watch_Call {
	return &MockTSafeReplicaInterface_Watch_Call{Call: _e.mock.On("Watch")}
}

func (_c *MockTSafeReplicaInterface_Watch_Call) Run(run func()) *MockTSafeReplicaInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTSafeReplicaInterface_Watch_Call) Return(_a0 Listener) *MockTSafeReplicaInterface_Watch_Call {
	_c.Call.Return(_a0)
	return _c
}

// WatchChannel provides a mock function with given fields: channel
func (_m *MockTSafeReplicaInterface) WatchChannel(channel string) Listener {
	ret := _m.Called(channel)

	var r0 Listener
	if rf, ok := ret.Get(0).(func(string) Listener); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Listener)
		}
	}

	return r0
}

// MockTSafeReplicaInterface_WatchChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchChannel'
type MockTSafeReplicaInterface_WatchChannel_Call struct {
	*mock.Call
}

// WatchChannel is a helper method to define mock.On call
//  - channel string
func (_e *MockTSafeReplicaInterface_Expecter) WatchChannel(channel interface{}) *MockTSafeReplicaInterface_WatchChannel_Call {
	return &MockTSafeReplicaInterface_WatchChannel_Call{Call: _e.mock.On("WatchChannel", channel)}
}

func (_c *MockTSafeReplicaInterface_WatchChannel_Call) Run(run func(channel string)) *MockTSafeReplicaInterface_WatchChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTSafeReplicaInterface_WatchChannel_Call) Return(_a0 Listener) *MockTSafeReplicaInterface_WatchChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

// addTSafe provides a mock function with given fields: vChannel
func (_m *MockTSafeReplicaInterface) addTSafe(vChannel string) {
	_m.Called(vChannel)
}

// MockTSafeReplicaInterface_addTSafe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'addTSafe'
type MockTSafeReplicaInterface_addTSafe_Call struct {
	*mock.Call
}

// addTSafe is a helper method to define mock.On call
//  - vChannel string
func (_e *MockTSafeReplicaInterface_Expecter) addTSafe(vChannel interface{}) *MockTSafeReplicaInterface_addTSafe_Call {
	return &MockTSafeReplicaInterface_addTSafe_Call{Call: _e.mock.On("addTSafe", vChannel)}
}

func (_c *MockTSafeReplicaInterface_addTSafe_Call) Run(run func(vChannel string)) *MockTSafeReplicaInterface_addTSafe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTSafeReplicaInterface_addTSafe_Call) Return() *MockTSafeReplicaInterface_addTSafe_Call {
	_c.Call.Return()
	return _c
}

// getTSafe provides a mock function with given fields: vChannel
func (_m *MockTSafeReplicaInterface) getTSafe(vChannel string) (uint64, error) {
	ret := _m.Called(vChannel)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(string) uint64); ok {
		r0 = rf(vChannel)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(vChannel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTSafeReplicaInterface_getTSafe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getTSafe'
type MockTSafeReplicaInterface_getTSafe_Call struct {
	*mock.Call
}

// getTSafe is a helper method to define mock.On call
//  - vChannel string
func (_e *MockTSafeReplicaInterface_Expecter) getTSafe(vChannel interface{}) *MockTSafeReplicaInterface_getTSafe_Call {
	return &MockTSafeReplicaInterface_getTSafe_Call{Call: _e.mock.On("getTSafe", vChannel)}
}

func (_c *MockTSafeReplicaInterface_getTSafe_Call) Run(run func(vChannel string)) *MockTSafeReplicaInterface_getTSafe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTSafeReplicaInterface_getTSafe_Call) Return(_a0 uint64, _a1 error) *MockTSafeReplicaInterface_getTSafe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// removeTSafe provides a mock function with given fields: vChannel
func (_m *MockTSafeReplicaInterface) removeTSafe(vChannel string) {
	_m.Called(vChannel)
}

// MockTSafeReplicaInterface_removeTSafe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'removeTSafe'
type MockTSafeReplicaInterface_removeTSafe_Call struct {
	*mock.Call
}

// removeTSafe is a helper method to define mock.On call
//  - vChannel string
func (_e *MockTSafeReplicaInterface_Expecter) removeTSafe(vChannel interface{}) *MockTSafeReplicaInterface_removeTSafe_Call {
	return &MockTSafeReplicaInterface_removeTSafe_Call{Call: _e.mock.On("removeTSafe", vChannel)}
}

func (_c *MockTSafeReplicaInterface_removeTSafe_Call) Run(run func(vChannel string)) *MockTSafeReplicaInterface_removeTSafe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTSafeReplicaInterface_removeTSafe_Call) Return() *MockTSafeReplicaInterface_removeTSafe_Call {
	_c.Call.Return()
	return _c
}

// setTSafe provides a mock function with given fields: vChannel, timestamp
func (_m *MockTSafeReplicaInterface) setTSafe(vChannel string, timestamp uint64) error {
	ret := _m.Called(vChannel, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint64) error); ok {
		r0 = rf(vChannel, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTSafeReplicaInterface_setTSafe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setTSafe'
type MockTSafeReplicaInterface_setTSafe_Call struct {
	*mock.Call
}

// setTSafe is a helper method to define mock.On call
//  - vChannel string
//  - timestamp uint64
func (_e *MockTSafeReplicaInterface_Expecter) setTSafe(vChannel interface{}, timestamp interface{}) *MockTSafeReplicaInterface_setTSafe_Call {
	return &MockTSafeReplicaInterface_setTSafe_Call{Call: _e.mock.On("setTSafe", vChannel, timestamp)}
}

func (_c *MockTSafeReplicaInterface_setTSafe_Call) Run(run func(vChannel string, timestamp uint64)) *MockTSafeReplicaInterface_setTSafe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64))
	})
	return _c
}

func (_c *MockTSafeReplicaInterface_setTSafe_Call) Return(_a0 error) *MockTSafeReplicaInterface_setTSafe_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockTSafeReplicaInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTSafeReplicaInterface creates a new instance of MockTSafeReplicaInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTSafeReplicaInterface(t mockConstructorTestingTNewMockTSafeReplicaInterface) *MockTSafeReplicaInterface {
	mock := &MockTSafeReplicaInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}