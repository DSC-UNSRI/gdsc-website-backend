// Code generated by MockGen. DO NOT EDIT.
// Source: ./init.go

// Package mock_division_usecase is a generated GoMock package.
package mock_division_usecase

import (
	reflect "reflect"

	model "github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockDivisionUsecase is a mock of DivisionUsecase interface.
type MockDivisionUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockDivisionUsecaseMockRecorder
}

// MockDivisionUsecaseMockRecorder is the mock recorder for MockDivisionUsecase.
type MockDivisionUsecaseMockRecorder struct {
	mock *MockDivisionUsecase
}

// NewMockDivisionUsecase creates a new mock instance.
func NewMockDivisionUsecase(ctrl *gomock.Controller) *MockDivisionUsecase {
	mock := &MockDivisionUsecase{ctrl: ctrl}
	mock.recorder = &MockDivisionUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDivisionUsecase) EXPECT() *MockDivisionUsecaseMockRecorder {
	return m.recorder
}

// CreateDivision mocks base method.
func (m *MockDivisionUsecase) CreateDivision(arg0 model.CreateDivisionRequest) model.WebServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDivision", arg0)
	ret0, _ := ret[0].(model.WebServiceResponse)
	return ret0
}

// CreateDivision indicates an expected call of CreateDivision.
func (mr *MockDivisionUsecaseMockRecorder) CreateDivision(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDivision", reflect.TypeOf((*MockDivisionUsecase)(nil).CreateDivision), arg0)
}

// DeleteDivision mocks base method.
func (m *MockDivisionUsecase) DeleteDivision(arg0 model.DeleteDivisionRequest) model.WebServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDivision", arg0)
	ret0, _ := ret[0].(model.WebServiceResponse)
	return ret0
}

// DeleteDivision indicates an expected call of DeleteDivision.
func (mr *MockDivisionUsecaseMockRecorder) DeleteDivision(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDivision", reflect.TypeOf((*MockDivisionUsecase)(nil).DeleteDivision), arg0)
}

// GetDivision mocks base method.
func (m *MockDivisionUsecase) GetDivision(arg0 model.GetDivisionRequest) model.WebServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDivision", arg0)
	ret0, _ := ret[0].(model.WebServiceResponse)
	return ret0
}

// GetDivision indicates an expected call of GetDivision.
func (mr *MockDivisionUsecaseMockRecorder) GetDivision(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDivision", reflect.TypeOf((*MockDivisionUsecase)(nil).GetDivision), arg0)
}
