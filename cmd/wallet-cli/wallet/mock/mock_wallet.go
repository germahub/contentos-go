// Code generated by MockGen. DO NOT EDIT.
// Source: wallet.go

// Package mock_wallet is a generated GoMock package.
package mock_wallet

import (
	wallet "github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWallet is a mock of Wallet interface
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockWallet) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockWalletMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockWallet)(nil).Name))
}

// Path mocks base method
func (m *MockWallet) Path() string {
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockWalletMockRecorder) Path() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockWallet)(nil).Path))
}

// GenerateNewKey mocks base method
func (m *MockWallet) GenerateNewKey() (string, string, error) {
	ret := m.ctrl.Call(m, "GenerateNewKey")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateNewKey indicates an expected call of GenerateNewKey
func (mr *MockWalletMockRecorder) GenerateNewKey() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateNewKey", reflect.TypeOf((*MockWallet)(nil).GenerateNewKey))
}

// Create mocks base method
func (m *MockWallet) Create(name, passphrase, pubKeyStr, privKeyStr string) error {
	ret := m.ctrl.Call(m, "Create", name, passphrase, pubKeyStr, privKeyStr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockWalletMockRecorder) Create(name, passphrase, pubKeyStr, privKeyStr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWallet)(nil).Create), name, passphrase, pubKeyStr, privKeyStr)
}

// GetUnlockedAccount mocks base method
func (m *MockWallet) GetUnlockedAccount(name string) (*wallet.PrivAccount, bool) {
	ret := m.ctrl.Call(m, "GetUnlockedAccount", name)
	ret0, _ := ret[0].(*wallet.PrivAccount)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetUnlockedAccount indicates an expected call of GetUnlockedAccount
func (mr *MockWalletMockRecorder) GetUnlockedAccount(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnlockedAccount", reflect.TypeOf((*MockWallet)(nil).GetUnlockedAccount), name)
}

// Load mocks base method
func (m *MockWallet) Load(name string) error {
	ret := m.ctrl.Call(m, "Load", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockWalletMockRecorder) Load(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockWallet)(nil).Load), name)
}

// Lock mocks base method
func (m *MockWallet) Lock(name string) error {
	ret := m.ctrl.Call(m, "Lock", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Lock indicates an expected call of Lock
func (mr *MockWalletMockRecorder) Lock(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockWallet)(nil).Lock), name)
}

// Unlock mocks base method
func (m *MockWallet) Unlock(name, passphrase string) error {
	ret := m.ctrl.Call(m, "Unlock", name, passphrase)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock
func (mr *MockWalletMockRecorder) Unlock(name, passphrase interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockWallet)(nil).Unlock), name, passphrase)
}

// List mocks base method
func (m *MockWallet) List() []string {
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	return ret0
}

// List indicates an expected call of List
func (mr *MockWalletMockRecorder) List() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWallet)(nil).List))
}

// Info mocks base method
func (m *MockWallet) Info(name string) string {
	ret := m.ctrl.Call(m, "Info", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Info indicates an expected call of Info
func (mr *MockWalletMockRecorder) Info(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockWallet)(nil).Info), name)
}

// Close mocks base method
func (m *MockWallet) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockWalletMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWallet)(nil).Close))
}

// IsLocked mocks base method
func (m *MockWallet) IsLocked(name string) bool {
	ret := m.ctrl.Call(m, "IsLocked", name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLocked indicates an expected call of IsLocked
func (mr *MockWalletMockRecorder) IsLocked(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLocked", reflect.TypeOf((*MockWallet)(nil).IsLocked), name)
}

// IsExist mocks base method
func (m *MockWallet) IsExist(name string) bool {
	ret := m.ctrl.Call(m, "IsExist", name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExist indicates an expected call of IsExist
func (mr *MockWalletMockRecorder) IsExist(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockWallet)(nil).IsExist), name)
}
