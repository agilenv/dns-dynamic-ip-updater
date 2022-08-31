// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package dns is a generated GoMock package.
package dns

import (
	context "context"
	reflect "reflect"

	track "github.com/agilenv/linkip/internal/dns/track"
	gomock "github.com/golang/mock/gomock"
)

// MockDNSProvider is a mock of DNSProvider interface.
type MockDNSProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDNSProviderMockRecorder
}

// MockDNSProviderMockRecorder is the mock recorder for MockDNSProvider.
type MockDNSProviderMockRecorder struct {
	mock *MockDNSProvider
}

// NewMockDNSProvider creates a new mock instance.
func NewMockDNSProvider(ctrl *gomock.Controller) *MockDNSProvider {
	mock := &MockDNSProvider{ctrl: ctrl}
	mock.recorder = &MockDNSProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNSProvider) EXPECT() *MockDNSProviderMockRecorder {
	return m.recorder
}

// GetRecord mocks base method.
func (m *MockDNSProvider) GetRecord(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecord", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecord indicates an expected call of GetRecord.
func (mr *MockDNSProviderMockRecorder) GetRecord(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockDNSProvider)(nil).GetRecord), ctx)
}

// UpdateRecord mocks base method.
func (m *MockDNSProvider) UpdateRecord(ctx context.Context, ip string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecord", ctx, ip)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRecord indicates an expected call of UpdateRecord.
func (mr *MockDNSProviderMockRecorder) UpdateRecord(ctx, ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecord", reflect.TypeOf((*MockDNSProvider)(nil).UpdateRecord), ctx, ip)
}

// MockTrackRepository is a mock of TrackRepository interface.
type MockTrackRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTrackRepositoryMockRecorder
}

// MockTrackRepositoryMockRecorder is the mock recorder for MockTrackRepository.
type MockTrackRepositoryMockRecorder struct {
	mock *MockTrackRepository
}

// NewMockTrackRepository creates a new mock instance.
func NewMockTrackRepository(ctrl *gomock.Controller) *MockTrackRepository {
	mock := &MockTrackRepository{ctrl: ctrl}
	mock.recorder = &MockTrackRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrackRepository) EXPECT() *MockTrackRepositoryMockRecorder {
	return m.recorder
}

// LastEvent mocks base method.
func (m *MockTrackRepository) LastEvent() track.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastEvent")
	ret0, _ := ret[0].(track.Event)
	return ret0
}

// LastEvent indicates an expected call of LastEvent.
func (mr *MockTrackRepositoryMockRecorder) LastEvent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastEvent", reflect.TypeOf((*MockTrackRepository)(nil).LastEvent))
}

// Save mocks base method.
func (m *MockTrackRepository) Save(event track.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockTrackRepositoryMockRecorder) Save(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTrackRepository)(nil).Save), event)
}

// MockPublicIPAPI is a mock of PublicIPAPI interface.
type MockPublicIPAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPublicIPAPIMockRecorder
}

// MockPublicIPAPIMockRecorder is the mock recorder for MockPublicIPAPI.
type MockPublicIPAPIMockRecorder struct {
	mock *MockPublicIPAPI
}

// NewMockPublicIPAPI creates a new mock instance.
func NewMockPublicIPAPI(ctrl *gomock.Controller) *MockPublicIPAPI {
	mock := &MockPublicIPAPI{ctrl: ctrl}
	mock.recorder = &MockPublicIPAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublicIPAPI) EXPECT() *MockPublicIPAPIMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPublicIPAPI) Get(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPublicIPAPIMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPublicIPAPI)(nil).Get), ctx)
}

// Name mocks base method.
func (m *MockPublicIPAPI) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPublicIPAPIMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPublicIPAPI)(nil).Name))
}
