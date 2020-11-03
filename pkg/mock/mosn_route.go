// Code generated by MockGen. DO NOT EDIT.
// Source: ../types/route.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "mosn.io/api"
	v2 "mosn.io/mosn/pkg/config/v2"
	types "mosn.io/mosn/pkg/types"
	reflect "reflect"
)

// MockRouters is a mock of Routers interface
type MockRouters struct {
	ctrl     *gomock.Controller
	recorder *MockRoutersMockRecorder
}

// MockRoutersMockRecorder is the mock recorder for MockRouters
type MockRoutersMockRecorder struct {
	mock *MockRouters
}

// NewMockRouters creates a new mock instance
func NewMockRouters(ctrl *gomock.Controller) *MockRouters {
	mock := &MockRouters{ctrl: ctrl}
	mock.recorder = &MockRoutersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouters) EXPECT() *MockRoutersMockRecorder {
	return m.recorder
}

// MatchRoute mocks base method
func (m *MockRouters) MatchRoute(ctx context.Context, headers api.HeaderMap, randomValue uint64) api.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchRoute", ctx, headers, randomValue)
	ret0, _ := ret[0].(api.Route)
	return ret0
}

// MatchRoute indicates an expected call of MatchRoute
func (mr *MockRoutersMockRecorder) MatchRoute(ctx, headers, randomValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchRoute", reflect.TypeOf((*MockRouters)(nil).MatchRoute), ctx, headers, randomValue)
}

// MatchAllRoutes mocks base method
func (m *MockRouters) MatchAllRoutes(ctx context.Context, headers api.HeaderMap, randomValue uint64) []api.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchAllRoutes", ctx, headers, randomValue)
	ret0, _ := ret[0].([]api.Route)
	return ret0
}

// MatchAllRoutes indicates an expected call of MatchAllRoutes
func (mr *MockRoutersMockRecorder) MatchAllRoutes(ctx, headers, randomValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchAllRoutes", reflect.TypeOf((*MockRouters)(nil).MatchAllRoutes), ctx, headers, randomValue)
}

// MatchRouteFromHeaderKV mocks base method
func (m *MockRouters) MatchRouteFromHeaderKV(ctx context.Context, headers api.HeaderMap, key, value string) api.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchRouteFromHeaderKV", ctx, headers, key, value)
	ret0, _ := ret[0].(api.Route)
	return ret0
}

// MatchRouteFromHeaderKV indicates an expected call of MatchRouteFromHeaderKV
func (mr *MockRoutersMockRecorder) MatchRouteFromHeaderKV(ctx, headers, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchRouteFromHeaderKV", reflect.TypeOf((*MockRouters)(nil).MatchRouteFromHeaderKV), ctx, headers, key, value)
}

// AddRoute mocks base method
func (m *MockRouters) AddRoute(domain string, route *v2.Router) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", domain, route)
	ret0, _ := ret[0].(int)
	return ret0
}

// AddRoute indicates an expected call of AddRoute
func (mr *MockRoutersMockRecorder) AddRoute(domain, route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockRouters)(nil).AddRoute), domain, route)
}

// RemoveAllRoutes mocks base method
func (m *MockRouters) RemoveAllRoutes(domain string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllRoutes", domain)
	ret0, _ := ret[0].(int)
	return ret0
}

// RemoveAllRoutes indicates an expected call of RemoveAllRoutes
func (mr *MockRoutersMockRecorder) RemoveAllRoutes(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllRoutes", reflect.TypeOf((*MockRouters)(nil).RemoveAllRoutes), domain)
}

// MockRouterManager is a mock of RouterManager interface
type MockRouterManager struct {
	ctrl     *gomock.Controller
	recorder *MockRouterManagerMockRecorder
}

// MockRouterManagerMockRecorder is the mock recorder for MockRouterManager
type MockRouterManagerMockRecorder struct {
	mock *MockRouterManager
}

// NewMockRouterManager creates a new mock instance
func NewMockRouterManager(ctrl *gomock.Controller) *MockRouterManager {
	mock := &MockRouterManager{ctrl: ctrl}
	mock.recorder = &MockRouterManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouterManager) EXPECT() *MockRouterManagerMockRecorder {
	return m.recorder
}

// AddOrUpdateRouters mocks base method
func (m *MockRouterManager) AddOrUpdateRouters(routerConfig *v2.RouterConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateRouters", routerConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateRouters indicates an expected call of AddOrUpdateRouters
func (mr *MockRouterManagerMockRecorder) AddOrUpdateRouters(routerConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateRouters", reflect.TypeOf((*MockRouterManager)(nil).AddOrUpdateRouters), routerConfig)
}

// GetRouterWrapperByName mocks base method
func (m *MockRouterManager) GetRouterWrapperByName(routerConfigName string) types.RouterWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouterWrapperByName", routerConfigName)
	ret0, _ := ret[0].(types.RouterWrapper)
	return ret0
}

// GetRouterWrapperByName indicates an expected call of GetRouterWrapperByName
func (mr *MockRouterManagerMockRecorder) GetRouterWrapperByName(routerConfigName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouterWrapperByName", reflect.TypeOf((*MockRouterManager)(nil).GetRouterWrapperByName), routerConfigName)
}

// AddRoute mocks base method
func (m *MockRouterManager) AddRoute(routerConfigName, domain string, route *v2.Router) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", routerConfigName, domain, route)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoute indicates an expected call of AddRoute
func (mr *MockRouterManagerMockRecorder) AddRoute(routerConfigName, domain, route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockRouterManager)(nil).AddRoute), routerConfigName, domain, route)
}

// RemoveAllRoutes mocks base method
func (m *MockRouterManager) RemoveAllRoutes(routerConfigName, domain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllRoutes", routerConfigName, domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllRoutes indicates an expected call of RemoveAllRoutes
func (mr *MockRouterManagerMockRecorder) RemoveAllRoutes(routerConfigName, domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllRoutes", reflect.TypeOf((*MockRouterManager)(nil).RemoveAllRoutes), routerConfigName, domain)
}

// MockRouteHandler is a mock of RouteHandler interface
type MockRouteHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRouteHandlerMockRecorder
}

// MockRouteHandlerMockRecorder is the mock recorder for MockRouteHandler
type MockRouteHandlerMockRecorder struct {
	mock *MockRouteHandler
}

// NewMockRouteHandler creates a new mock instance
func NewMockRouteHandler(ctrl *gomock.Controller) *MockRouteHandler {
	mock := &MockRouteHandler{ctrl: ctrl}
	mock.recorder = &MockRouteHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteHandler) EXPECT() *MockRouteHandlerMockRecorder {
	return m.recorder
}

// IsAvailable mocks base method
func (m *MockRouteHandler) IsAvailable(arg0 context.Context, arg1 types.ClusterManager) (types.ClusterSnapshot, types.HandlerStatus) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAvailable", arg0, arg1)
	ret0, _ := ret[0].(types.ClusterSnapshot)
	ret1, _ := ret[1].(types.HandlerStatus)
	return ret0, ret1
}

// IsAvailable indicates an expected call of IsAvailable
func (mr *MockRouteHandlerMockRecorder) IsAvailable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAvailable", reflect.TypeOf((*MockRouteHandler)(nil).IsAvailable), arg0, arg1)
}

// Route mocks base method
func (m *MockRouteHandler) Route() api.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route")
	ret0, _ := ret[0].(api.Route)
	return ret0
}

// Route indicates an expected call of Route
func (mr *MockRouteHandlerMockRecorder) Route() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockRouteHandler)(nil).Route))
}

// MockRouterWrapper is a mock of RouterWrapper interface
type MockRouterWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockRouterWrapperMockRecorder
}

// MockRouterWrapperMockRecorder is the mock recorder for MockRouterWrapper
type MockRouterWrapperMockRecorder struct {
	mock *MockRouterWrapper
}

// NewMockRouterWrapper creates a new mock instance
func NewMockRouterWrapper(ctrl *gomock.Controller) *MockRouterWrapper {
	mock := &MockRouterWrapper{ctrl: ctrl}
	mock.recorder = &MockRouterWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouterWrapper) EXPECT() *MockRouterWrapperMockRecorder {
	return m.recorder
}

// GetRouters mocks base method
func (m *MockRouterWrapper) GetRouters() types.Routers {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouters")
	ret0, _ := ret[0].(types.Routers)
	return ret0
}

// GetRouters indicates an expected call of GetRouters
func (mr *MockRouterWrapperMockRecorder) GetRouters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouters", reflect.TypeOf((*MockRouterWrapper)(nil).GetRouters))
}

// GetRoutersConfig mocks base method
func (m *MockRouterWrapper) GetRoutersConfig() v2.RouterConfiguration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoutersConfig")
	ret0, _ := ret[0].(v2.RouterConfiguration)
	return ret0
}

// GetRoutersConfig indicates an expected call of GetRoutersConfig
func (mr *MockRouterWrapperMockRecorder) GetRoutersConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutersConfig", reflect.TypeOf((*MockRouterWrapper)(nil).GetRoutersConfig))
}

// MockVirtualHost is a mock of VirtualHost interface
type MockVirtualHost struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualHostMockRecorder
}

// MockVirtualHostMockRecorder is the mock recorder for MockVirtualHost
type MockVirtualHostMockRecorder struct {
	mock *MockVirtualHost
}

// NewMockVirtualHost creates a new mock instance
func NewMockVirtualHost(ctrl *gomock.Controller) *MockVirtualHost {
	mock := &MockVirtualHost{ctrl: ctrl}
	mock.recorder = &MockVirtualHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualHost) EXPECT() *MockVirtualHostMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockVirtualHost) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockVirtualHostMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockVirtualHost)(nil).Name))
}

// GetRouteFromEntries mocks base method
func (m *MockVirtualHost) GetRouteFromEntries(ctx context.Context, headers api.HeaderMap, randomValue uint64) api.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteFromEntries", ctx, headers, randomValue)
	ret0, _ := ret[0].(api.Route)
	return ret0
}

// GetRouteFromEntries indicates an expected call of GetRouteFromEntries
func (mr *MockVirtualHostMockRecorder) GetRouteFromEntries(ctx, headers, randomValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteFromEntries", reflect.TypeOf((*MockVirtualHost)(nil).GetRouteFromEntries), ctx, headers, randomValue)
}

// GetAllRoutesFromEntries mocks base method
func (m *MockVirtualHost) GetAllRoutesFromEntries(ctx context.Context, headers api.HeaderMap, randomValue uint64) []api.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRoutesFromEntries", ctx, headers, randomValue)
	ret0, _ := ret[0].([]api.Route)
	return ret0
}

// GetAllRoutesFromEntries indicates an expected call of GetAllRoutesFromEntries
func (mr *MockVirtualHostMockRecorder) GetAllRoutesFromEntries(ctx, headers, randomValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRoutesFromEntries", reflect.TypeOf((*MockVirtualHost)(nil).GetAllRoutesFromEntries), ctx, headers, randomValue)
}

// GetRouteFromHeaderKV mocks base method
func (m *MockVirtualHost) GetRouteFromHeaderKV(key, value string) api.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteFromHeaderKV", key, value)
	ret0, _ := ret[0].(api.Route)
	return ret0
}

// GetRouteFromHeaderKV indicates an expected call of GetRouteFromHeaderKV
func (mr *MockVirtualHostMockRecorder) GetRouteFromHeaderKV(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteFromHeaderKV", reflect.TypeOf((*MockVirtualHost)(nil).GetRouteFromHeaderKV), key, value)
}

// AddRoute mocks base method
func (m *MockVirtualHost) AddRoute(route *v2.Router) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", route)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoute indicates an expected call of AddRoute
func (mr *MockVirtualHostMockRecorder) AddRoute(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockVirtualHost)(nil).AddRoute), route)
}

// RemoveAllRoutes mocks base method
func (m *MockVirtualHost) RemoveAllRoutes() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAllRoutes")
}

// RemoveAllRoutes indicates an expected call of RemoveAllRoutes
func (mr *MockVirtualHostMockRecorder) RemoveAllRoutes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllRoutes", reflect.TypeOf((*MockVirtualHost)(nil).RemoveAllRoutes))
}

// MockHeaderFormat is a mock of HeaderFormat interface
type MockHeaderFormat struct {
	ctrl     *gomock.Controller
	recorder *MockHeaderFormatMockRecorder
}

// MockHeaderFormatMockRecorder is the mock recorder for MockHeaderFormat
type MockHeaderFormatMockRecorder struct {
	mock *MockHeaderFormat
}

// NewMockHeaderFormat creates a new mock instance
func NewMockHeaderFormat(ctrl *gomock.Controller) *MockHeaderFormat {
	mock := &MockHeaderFormat{ctrl: ctrl}
	mock.recorder = &MockHeaderFormatMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHeaderFormat) EXPECT() *MockHeaderFormatMockRecorder {
	return m.recorder
}

// Format mocks base method
func (m *MockHeaderFormat) Format(info api.RequestInfo) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Format", info)
	ret0, _ := ret[0].(string)
	return ret0
}

// Format indicates an expected call of Format
func (mr *MockHeaderFormatMockRecorder) Format(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Format", reflect.TypeOf((*MockHeaderFormat)(nil).Format), info)
}

// Append mocks base method
func (m *MockHeaderFormat) Append() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Append indicates an expected call of Append
func (mr *MockHeaderFormatMockRecorder) Append() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockHeaderFormat)(nil).Append))
}

// MockQueryParameterMatcher is a mock of QueryParameterMatcher interface
type MockQueryParameterMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockQueryParameterMatcherMockRecorder
}

// MockQueryParameterMatcherMockRecorder is the mock recorder for MockQueryParameterMatcher
type MockQueryParameterMatcherMockRecorder struct {
	mock *MockQueryParameterMatcher
}

// NewMockQueryParameterMatcher creates a new mock instance
func NewMockQueryParameterMatcher(ctrl *gomock.Controller) *MockQueryParameterMatcher {
	mock := &MockQueryParameterMatcher{ctrl: ctrl}
	mock.recorder = &MockQueryParameterMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryParameterMatcher) EXPECT() *MockQueryParameterMatcherMockRecorder {
	return m.recorder
}

// Matches mocks base method
func (m *MockQueryParameterMatcher) Matches(requestQueryParams types.QueryParams) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Matches", requestQueryParams)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Matches indicates an expected call of Matches
func (mr *MockQueryParameterMatcherMockRecorder) Matches(requestQueryParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Matches", reflect.TypeOf((*MockQueryParameterMatcher)(nil).Matches), requestQueryParams)
}

// MockConfigUtility is a mock of ConfigUtility interface
type MockConfigUtility struct {
	ctrl     *gomock.Controller
	recorder *MockConfigUtilityMockRecorder
}

// MockConfigUtilityMockRecorder is the mock recorder for MockConfigUtility
type MockConfigUtilityMockRecorder struct {
	mock *MockConfigUtility
}

// NewMockConfigUtility creates a new mock instance
func NewMockConfigUtility(ctrl *gomock.Controller) *MockConfigUtility {
	mock := &MockConfigUtility{ctrl: ctrl}
	mock.recorder = &MockConfigUtilityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigUtility) EXPECT() *MockConfigUtilityMockRecorder {
	return m.recorder
}

// MatchHeaders mocks base method
func (m *MockConfigUtility) MatchHeaders(requestHeaders map[string]string, configHeaders []*types.HeaderData) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchHeaders", requestHeaders, configHeaders)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MatchHeaders indicates an expected call of MatchHeaders
func (mr *MockConfigUtilityMockRecorder) MatchHeaders(requestHeaders, configHeaders interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchHeaders", reflect.TypeOf((*MockConfigUtility)(nil).MatchHeaders), requestHeaders, configHeaders)
}

// MatchQueryParams mocks base method
func (m *MockConfigUtility) MatchQueryParams(queryParams types.QueryParams, configQueryParams []types.QueryParameterMatcher) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchQueryParams", queryParams, configQueryParams)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MatchQueryParams indicates an expected call of MatchQueryParams
func (mr *MockConfigUtilityMockRecorder) MatchQueryParams(queryParams, configQueryParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchQueryParams", reflect.TypeOf((*MockConfigUtility)(nil).MatchQueryParams), queryParams, configQueryParams)
}

// MockLowerCaseString is a mock of LowerCaseString interface
type MockLowerCaseString struct {
	ctrl     *gomock.Controller
	recorder *MockLowerCaseStringMockRecorder
}

// MockLowerCaseStringMockRecorder is the mock recorder for MockLowerCaseString
type MockLowerCaseStringMockRecorder struct {
	mock *MockLowerCaseString
}

// NewMockLowerCaseString creates a new mock instance
func NewMockLowerCaseString(ctrl *gomock.Controller) *MockLowerCaseString {
	mock := &MockLowerCaseString{ctrl: ctrl}
	mock.recorder = &MockLowerCaseStringMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLowerCaseString) EXPECT() *MockLowerCaseStringMockRecorder {
	return m.recorder
}

// Lower mocks base method
func (m *MockLowerCaseString) Lower() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lower")
}

// Lower indicates an expected call of Lower
func (mr *MockLowerCaseStringMockRecorder) Lower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lower", reflect.TypeOf((*MockLowerCaseString)(nil).Lower))
}

// Equal mocks base method
func (m *MockLowerCaseString) Equal(rhs types.LowerCaseString) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", rhs)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockLowerCaseStringMockRecorder) Equal(rhs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockLowerCaseString)(nil).Equal), rhs)
}

// Get mocks base method
func (m *MockLowerCaseString) Get() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockLowerCaseStringMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLowerCaseString)(nil).Get))
}
