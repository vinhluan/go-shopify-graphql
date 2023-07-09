// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vinhluan/go-shopify-graphql (interfaces: OrderService)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/vinhluan/go-shopify-graphql/model"
	shopify "github.com/vinhluan/go-shopify-graphql"
	graphql "github.com/r0busta/graphql"
)

// MockOrderService is a mock of OrderService interface.
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService.
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance.
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockOrderService) Get(arg0 graphql.ID) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrderService)(nil).Get), arg0)
}

// List mocks base method.
func (m *MockOrderService) List(arg0 shopify.ListOptions) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrderServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrderService)(nil).List), arg0)
}

// ListAfterCursor mocks base method.
func (m *MockOrderService) ListAfterCursor(arg0 shopify.ListOptions) ([]model.Order, *string, *string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAfterCursor", arg0)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(*string)
	ret2, _ := ret[2].(*string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListAfterCursor indicates an expected call of ListAfterCursor.
func (mr *MockOrderServiceMockRecorder) ListAfterCursor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAfterCursor", reflect.TypeOf((*MockOrderService)(nil).ListAfterCursor), arg0)
}

// ListAll mocks base method.
func (m *MockOrderService) ListAll() ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll")
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockOrderServiceMockRecorder) ListAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockOrderService)(nil).ListAll))
}

// Update mocks base method.
func (m *MockOrderService) Update(arg0 model.OrderInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrderServiceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderService)(nil).Update), arg0)
}
