// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/auth/types.go

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// MockAuthenticationEndpoint is a mock of AuthenticationEndpoint interface.
type MockAuthenticationEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationEndpointMockRecorder
}

// MockAuthenticationEndpointMockRecorder is the mock recorder for MockAuthenticationEndpoint.
type MockAuthenticationEndpointMockRecorder struct {
	mock *MockAuthenticationEndpoint
}

// NewMockAuthenticationEndpoint creates a new mock instance.
func NewMockAuthenticationEndpoint(ctrl *gomock.Controller) *MockAuthenticationEndpoint {
	mock := &MockAuthenticationEndpoint{ctrl: ctrl}
	mock.recorder = &MockAuthenticationEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticationEndpoint) EXPECT() *MockAuthenticationEndpointMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticationEndpoint) Authenticate(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Authenticate", ctx)
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticationEndpointMockRecorder) Authenticate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationEndpoint)(nil).Authenticate), ctx)
}

// MockAuthenticationService is a mock of AuthenticationService interface.
type MockAuthenticationService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationServiceMockRecorder
}

// MockAuthenticationServiceMockRecorder is the mock recorder for MockAuthenticationService.
type MockAuthenticationServiceMockRecorder struct {
	mock *MockAuthenticationService
}

// NewMockAuthenticationService creates a new mock instance.
func NewMockAuthenticationService(ctrl *gomock.Controller) *MockAuthenticationService {
	mock := &MockAuthenticationService{ctrl: ctrl}
	mock.recorder = &MockAuthenticationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticationService) EXPECT() *MockAuthenticationServiceMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticationService) Authenticate(ctx context.Context, principal *Principal) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, principal)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticationServiceMockRecorder) Authenticate(ctx, principal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationService)(nil).Authenticate), ctx, principal)
}

// MockAuthenticationDelegate is a mock of AuthenticationDelegate interface.
type MockAuthenticationDelegate struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationDelegateMockRecorder
}

// MockAuthenticationDelegateMockRecorder is the mock recorder for MockAuthenticationDelegate.
type MockAuthenticationDelegateMockRecorder struct {
	mock *MockAuthenticationDelegate
}

// NewMockAuthenticationDelegate creates a new mock instance.
func NewMockAuthenticationDelegate(ctrl *gomock.Controller) *MockAuthenticationDelegate {
	mock := &MockAuthenticationDelegate{ctrl: ctrl}
	mock.recorder = &MockAuthenticationDelegateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticationDelegate) EXPECT() *MockAuthenticationDelegateMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticationDelegate) Authenticate(ctx context.Context, principal *Principal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, principal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticationDelegateMockRecorder) Authenticate(ctx, principal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticationDelegate)(nil).Authenticate), ctx, principal)
}

// MockAuthorizationFilter is a mock of AuthorizationFilter interface.
type MockAuthorizationFilter struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationFilterMockRecorder
}

// MockAuthorizationFilterMockRecorder is the mock recorder for MockAuthorizationFilter.
type MockAuthorizationFilterMockRecorder struct {
	mock *MockAuthorizationFilter
}

// NewMockAuthorizationFilter creates a new mock instance.
func NewMockAuthorizationFilter(ctrl *gomock.Controller) *MockAuthorizationFilter {
	mock := &MockAuthorizationFilter{ctrl: ctrl}
	mock.recorder = &MockAuthorizationFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationFilter) EXPECT() *MockAuthorizationFilterMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthorizationFilter) Authorize(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Authorize", ctx)
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizationFilterMockRecorder) Authorize(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizationFilter)(nil).Authorize), ctx)
}

// MockAuthorizationService is a mock of AuthorizationService interface.
type MockAuthorizationService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationServiceMockRecorder
}

// MockAuthorizationServiceMockRecorder is the mock recorder for MockAuthorizationService.
type MockAuthorizationServiceMockRecorder struct {
	mock *MockAuthorizationService
}

// NewMockAuthorizationService creates a new mock instance.
func NewMockAuthorizationService(ctrl *gomock.Controller) *MockAuthorizationService {
	mock := &MockAuthorizationService{ctrl: ctrl}
	mock.recorder = &MockAuthorizationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationService) EXPECT() *MockAuthorizationServiceMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthorizationService) Authorize(ctx context.Context, tokenString string) (*Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, tokenString)
	ret0, _ := ret[0].(*Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizationServiceMockRecorder) Authorize(ctx, tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizationService)(nil).Authorize), ctx, tokenString)
}

// MockAuthorizationDelegate is a mock of AuthorizationDelegate interface.
type MockAuthorizationDelegate struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationDelegateMockRecorder
}

// MockAuthorizationDelegateMockRecorder is the mock recorder for MockAuthorizationDelegate.
type MockAuthorizationDelegateMockRecorder struct {
	mock *MockAuthorizationDelegate
}

// NewMockAuthorizationDelegate creates a new mock instance.
func NewMockAuthorizationDelegate(ctrl *gomock.Controller) *MockAuthorizationDelegate {
	mock := &MockAuthorizationDelegate{ctrl: ctrl}
	mock.recorder = &MockAuthorizationDelegateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationDelegate) EXPECT() *MockAuthorizationDelegateMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthorizationDelegate) Authorize(ctx context.Context, principal *Principal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, principal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizationDelegateMockRecorder) Authorize(ctx, principal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizationDelegate)(nil).Authorize), ctx, principal)
}

// MockTokenManager is a mock of TokenManager interface.
type MockTokenManager struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerMockRecorder
}

// MockTokenManagerMockRecorder is the mock recorder for MockTokenManager.
type MockTokenManagerMockRecorder struct {
	mock *MockTokenManager
}

// NewMockTokenManager creates a new mock instance.
func NewMockTokenManager(ctrl *gomock.Controller) *MockTokenManager {
	mock := &MockTokenManager{ctrl: ctrl}
	mock.recorder = &MockTokenManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManager) EXPECT() *MockTokenManagerMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockTokenManager) Generate(principal *Principal) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", principal)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockTokenManagerMockRecorder) Generate(principal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTokenManager)(nil).Generate), principal)
}

// Validate mocks base method.
func (m *MockTokenManager) Validate(tokenString string) (*Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", tokenString)
	ret0, _ := ret[0].(*Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenManagerMockRecorder) Validate(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenManager)(nil).Validate), tokenString)
}

// MockPrincipalManager is a mock of PrincipalManager interface.
type MockPrincipalManager struct {
	ctrl     *gomock.Controller
	recorder *MockPrincipalManagerMockRecorder
}

// MockPrincipalManagerMockRecorder is the mock recorder for MockPrincipalManager.
type MockPrincipalManagerMockRecorder struct {
	mock *MockPrincipalManager
}

// NewMockPrincipalManager creates a new mock instance.
func NewMockPrincipalManager(ctrl *gomock.Controller) *MockPrincipalManager {
	mock := &MockPrincipalManager{ctrl: ctrl}
	mock.recorder = &MockPrincipalManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrincipalManager) EXPECT() *MockPrincipalManagerMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockPrincipalManager) ChangePassword(ctx context.Context, principal *Principal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, principal)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockPrincipalManagerMockRecorder) ChangePassword(ctx, principal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockPrincipalManager)(nil).ChangePassword), ctx, principal)
}

// Create mocks base method.
func (m *MockPrincipalManager) Create(ctx context.Context, principal *Principal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, principal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPrincipalManagerMockRecorder) Create(ctx, principal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPrincipalManager)(nil).Create), ctx, principal)
}

// Delete mocks base method.
func (m *MockPrincipalManager) Delete(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPrincipalManagerMockRecorder) Delete(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPrincipalManager)(nil).Delete), ctx, username)
}

// Exists mocks base method.
func (m *MockPrincipalManager) Exists(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockPrincipalManagerMockRecorder) Exists(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockPrincipalManager)(nil).Exists), ctx, username)
}

// Find mocks base method.
func (m *MockPrincipalManager) Find(ctx context.Context, username string) (*Principal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, username)
	ret0, _ := ret[0].(*Principal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockPrincipalManagerMockRecorder) Find(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPrincipalManager)(nil).Find), ctx, username)
}

// Update mocks base method.
func (m *MockPrincipalManager) Update(ctx context.Context, principal *Principal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, principal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPrincipalManagerMockRecorder) Update(ctx, principal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPrincipalManager)(nil).Update), ctx, principal)
}
