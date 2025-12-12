package zabbix

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type (
	// Params Zabbix request param
	Params map[string]interface{}
)

// Version constants
const (
	ZabbixVersion60 = "6.0"
	ZabbixVersion70 = "7.0"
)

// Feature constants for Zabbix 7.0
const (
	FeatureHistoryPush   = "history.push"
	FeatureMFA          = "mfa"
	FeatureProxyGroup   = "proxygroup"
	FeatureBrowserItem  = "browser_item"
	FeatureHeadersV7    = "headers_v7"
	FeatureProxyID      = "proxyid"
	FeatureMonitoredBy  = "monitored_by"
)

// HeaderField represents a header field for Zabbix 7.0 format
type HeaderField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// VersionManager manages Zabbix version detection and feature support
type VersionManager struct {
	serverVersion    string
	majorVersion     int
	minorVersion     int
	is70             bool
	is60             bool
	supportedFeatures map[string]bool
}

// NewVersionManager creates a new version manager
func NewVersionManager() *VersionManager {
	return &VersionManager{
		supportedFeatures: make(map[string]bool),
	}
}

// DetectVersion detects the Zabbix server version
func (vm *VersionManager) DetectVersion(api *API) error {
	version, err := api.Version()
	if err != nil {
		return fmt.Errorf("failed to get version: %w", err)
	}
	
	vm.serverVersion = version
	
	// Parse version string (e.g., "7.0.0", "6.4.5")
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		major, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("invalid major version: %s", parts[0])
		}
		vm.majorVersion = major
		
		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid minor version: %s", parts[1])
		}
		vm.minorVersion = minor
		
		vm.is70 = major == 7
		vm.is60 = major == 6
	}
	
	// Initialize feature support
	vm.initializeFeatureSupport()
	
	return nil
}

// GetVersion returns the detected server version
func (vm *VersionManager) GetVersion() string {
	return vm.serverVersion
}

// Is70 returns true if server is Zabbix 7.0
func (vm *VersionManager) Is70() bool {
	return vm.is70
}

// Is60 returns true if server is Zabbix 6.0
func (vm *VersionManager) Is60() bool {
	return vm.is60
}

// IsFeatureSupported checks if a feature is supported
func (vm *VersionManager) IsFeatureSupported(feature string) bool {
	supported, exists := vm.supportedFeatures[feature]
	return exists && supported
}

// initializeFeatureSupport initializes supported features based on version
func (vm *VersionManager) initializeFeatureSupport() {
	// Clear existing features
	for k := range vm.supportedFeatures {
		delete(vm.supportedFeatures, k)
	}
	
	if vm.is70 {
		// Zabbix 7.0 features
		vm.supportedFeatures[FeatureHistoryPush] = true
		vm.supportedFeatures[FeatureMFA] = true
		vm.supportedFeatures[FeatureProxyGroup] = true
		vm.supportedFeatures[FeatureBrowserItem] = true
		vm.supportedFeatures[FeatureHeadersV7] = true
		vm.supportedFeatures[FeatureProxyID] = true
		vm.supportedFeatures[FeatureMonitoredBy] = true
	} else if vm.is60 {
		// Zabbix 6.0 features (legacy format)
		vm.supportedFeatures[FeatureHeadersV7] = false
		vm.supportedFeatures[FeatureProxyID] = false
		vm.supportedFeatures[FeatureMonitoredBy] = false
	}
}

// Adapter interfaces for multi-version support

// ItemAdapter defines interface for item operations across different Zabbix versions
type ItemAdapter interface {
	CreateItems(items []Item) error
	GetItems(params Params) ([]Item, error)
	UpdateItems(items []Item) error
	DeleteItems(items []Item) error
	PrepareHeaders(item Item) interface{}
}

// HostAdapter defines interface for host operations across different Zabbix versions
type HostAdapter interface {
	CreateHosts(hosts []Host) error
	GetHosts(params Params) ([]Host, error)
	UpdateHosts(hosts []Host) error
	DeleteHosts(hosts []Host) error
	PrepareProxyFields(host Host) map[string]interface{}
}

// BaseAdapter provides common functionality for all adapters
type BaseAdapter struct {
	api     *API
	version *VersionManager
}

// NewBaseAdapter creates a new base adapter
func NewBaseAdapter(api *API) *BaseAdapter {
	return &BaseAdapter{
		api:     api,
		version: api.versionManager,
	}
}

// Zabbix6ItemAdapter implements ItemAdapter for Zabbix 6.0
type Zabbix6ItemAdapter struct {
	*BaseAdapter
}

// NewZabbix6ItemAdapter creates a new Zabbix 6.0 item adapter
func NewZabbix6ItemAdapter(api *API) *Zabbix6ItemAdapter {
	return &Zabbix6ItemAdapter{
		BaseAdapter: NewBaseAdapter(api),
	}
}

// CreateItems creates items using Zabbix 6.0 format
func (adapter *Zabbix6ItemAdapter) CreateItems(items []Item) error {
	// Convert headers to 6.0 format
	for i := range items {
		if len(items[i].HeadersV7) > 0 {
			// Convert 7.0 format to 6.0 format
			headers := make(map[string]string)
			for _, header := range items[i].HeadersV7 {
				headers[header.Name] = header.Value
			}
			items[i].HeadersV6 = headers
		}
	}
	return adapter.api.ItemsCreate(items)
}

// GetItems gets items using Zabbix 6.0 format
func (adapter *Zabbix6ItemAdapter) GetItems(params Params) ([]Item, error) {
	return adapter.api.ItemsGet(params)
}

// UpdateItems updates items using Zabbix 6.0 format
func (adapter *Zabbix6ItemAdapter) UpdateItems(items []Item) error {
	// Convert headers to 6.0 format
	for i := range items {
		if len(items[i].HeadersV7) > 0 {
			headers := make(map[string]string)
			for _, header := range items[i].HeadersV7 {
				headers[header.Name] = header.Value
			}
			items[i].HeadersV6 = headers
		}
	}
	return adapter.api.ItemsUpdate(items)
}

// DeleteItems deletes items
func (adapter *Zabbix6ItemAdapter) DeleteItems(items []Item) error {
	return adapter.api.ItemsDelete(items)
}

// PrepareHeaders prepares headers in Zabbix 6.0 format
func (adapter *Zabbix6ItemAdapter) PrepareHeaders(item Item) interface{} {
	if len(item.HeadersV6) > 0 {
		return item.HeadersV6
	}
	if len(item.HeadersV7) > 0 {
		// Convert 7.0 format to 6.0 format
		headers := make(map[string]string)
		for _, header := range item.HeadersV7 {
			headers[header.Name] = header.Value
		}
		return headers
	}
	return nil
}

// Zabbix7ItemAdapter implements ItemAdapter for Zabbix 7.0
type Zabbix7ItemAdapter struct {
	*BaseAdapter
}

// NewZabbix7ItemAdapter creates a new Zabbix 7.0 item adapter
func NewZabbix7ItemAdapter(api *API) *Zabbix7ItemAdapter {
	return &Zabbix7ItemAdapter{
		BaseAdapter: NewBaseAdapter(api),
	}
}

// CreateItems creates items using Zabbix 7.0 format
func (adapter *Zabbix7ItemAdapter) CreateItems(items []Item) error {
	// Convert headers to 7.0 format
	for i := range items {
		if len(items[i].HeadersV6) > 0 {
			// Convert 6.0 format to 7.0 format
			headers := make([]HeaderField, 0, len(items[i].HeadersV6))
			for name, value := range items[i].HeadersV6 {
				headers = append(headers, HeaderField{
					Name:  name,
					Value: value,
				})
			}
			items[i].HeadersV7 = headers
		}
	}
	return adapter.api.ItemsCreate(items)
}

// GetItems gets items using Zabbix 7.0 format
func (adapter *Zabbix7ItemAdapter) GetItems(params Params) ([]Item, error) {
	return adapter.api.ItemsGet(params)
}

// UpdateItems updates items using Zabbix 7.0 format
func (adapter *Zabbix7ItemAdapter) UpdateItems(items []Item) error {
	// Convert headers to 7.0 format
	for i := range items {
		if len(items[i].HeadersV6) > 0 {
			headers := make([]HeaderField, 0, len(items[i].HeadersV6))
			for name, value := range items[i].HeadersV6 {
				headers = append(headers, HeaderField{
					Name:  name,
					Value: value,
				})
			}
			items[i].HeadersV7 = headers
		}
	}
	return adapter.api.ItemsUpdate(items)
}

// DeleteItems deletes items
func (adapter *Zabbix7ItemAdapter) DeleteItems(items []Item) error {
	return adapter.api.ItemsDelete(items)
}

// PrepareHeaders prepares headers in Zabbix 7.0 format
func (adapter *Zabbix7ItemAdapter) PrepareHeaders(item Item) interface{} {
	if len(item.HeadersV7) > 0 {
		return item.HeadersV7
	}
	if len(item.HeadersV6) > 0 {
		// Convert 6.0 format to 7.0 format
		headers := make([]HeaderField, 0, len(item.HeadersV6))
		for name, value := range item.HeadersV6 {
			headers = append(headers, HeaderField{
				Name:  name,
				Value: value,
			})
		}
		return headers
	}
	return nil
}

// Zabbix6HostAdapter implements HostAdapter for Zabbix 6.0
type Zabbix6HostAdapter struct {
	*BaseAdapter
}

// NewZabbix6HostAdapter creates a new Zabbix 6.0 host adapter
func NewZabbix6HostAdapter(api *API) *Zabbix6HostAdapter {
	return &Zabbix6HostAdapter{
		BaseAdapter: NewBaseAdapter(api),
	}
}

// CreateHosts creates hosts using Zabbix 6.0 format
func (adapter *Zabbix6HostAdapter) CreateHosts(hosts []Host) error {
	// Convert proxy fields to 6.0 format
	for i := range hosts {
		if hosts[i].ProxyID != "" && hosts[i].ProxyHostID == "" {
			hosts[i].ProxyHostID = hosts[i].ProxyID
		}
	}
	return adapter.api.HostsCreate(hosts)
}

// GetHosts gets hosts using Zabbix 6.0 format
func (adapter *Zabbix6HostAdapter) GetHosts(params Params) ([]Host, error) {
	return adapter.api.HostsGet(params)
}

// UpdateHosts updates hosts using Zabbix 6.0 format
func (adapter *Zabbix6HostAdapter) UpdateHosts(hosts []Host) error {
	// Convert proxy fields to 6.0 format
	for i := range hosts {
		if hosts[i].ProxyID != "" && hosts[i].ProxyHostID == "" {
			hosts[i].ProxyHostID = hosts[i].ProxyID
		}
	}
	return adapter.api.HostsUpdate(hosts)
}

// DeleteHosts deletes hosts
func (adapter *Zabbix6HostAdapter) DeleteHosts(hosts []Host) error {
	return adapter.api.HostsDelete(hosts)
}

// PrepareProxyFields prepares proxy fields in Zabbix 6.0 format
func (adapter *Zabbix6HostAdapter) PrepareProxyFields(host Host) map[string]interface{} {
	result := make(map[string]interface{})
	if host.ProxyID != "" {
		result["proxy_hostid"] = host.ProxyID
	} else if host.ProxyHostID != "" {
		result["proxy_hostid"] = host.ProxyHostID
	}
	return result
}

// Zabbix7HostAdapter implements HostAdapter for Zabbix 7.0
type Zabbix7HostAdapter struct {
	*BaseAdapter
}

// NewZabbix7HostAdapter creates a new Zabbix 7.0 host adapter
func NewZabbix7HostAdapter(api *API) *Zabbix7HostAdapter {
	return &Zabbix7HostAdapter{
		BaseAdapter: NewBaseAdapter(api),
	}
}

// CreateHosts creates hosts using Zabbix 7.0 format
func (adapter *Zabbix7HostAdapter) CreateHosts(hosts []Host) error {
	// Convert proxy fields to 7.0 format
	for i := range hosts {
		if hosts[i].ProxyHostID != "" && hosts[i].ProxyID == "" {
			hosts[i].ProxyID = hosts[i].ProxyHostID
		}
		// Set monitored_by if not specified
		if hosts[i].MonitoredBy == 0 {
			hosts[i].MonitoredBy = MonitoredByProxy // Default to proxy if proxy is specified
		}
	}
	return adapter.api.HostsCreate(hosts)
}

// GetHosts gets hosts using Zabbix 7.0 format
func (adapter *Zabbix7HostAdapter) GetHosts(params Params) ([]Host, error) {
	return adapter.api.HostsGet(params)
}

// UpdateHosts updates hosts using Zabbix 7.0 format
func (adapter *Zabbix7HostAdapter) UpdateHosts(hosts []Host) error {
	// Convert proxy fields to 7.0 format
	for i := range hosts {
		if hosts[i].ProxyHostID != "" && hosts[i].ProxyID == "" {
			hosts[i].ProxyID = hosts[i].ProxyHostID
		}
	}
	return adapter.api.HostsUpdate(hosts)
}

// DeleteHosts deletes hosts
func (adapter *Zabbix7HostAdapter) DeleteHosts(hosts []Host) error {
	return adapter.api.HostsDelete(hosts)
}

// PrepareProxyFields prepares proxy fields in Zabbix 7.0 format
func (adapter *Zabbix7HostAdapter) PrepareProxyFields(host Host) map[string]interface{} {
	result := make(map[string]interface{})
	if host.ProxyID != "" {
		result["proxyid"] = host.ProxyID
	}
	if host.MonitoredBy != 0 {
		result["monitored_by"] = host.MonitoredBy
	}
	return result
}

type request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth,omitempty"`
	ID      int32       `json:"id"`
}

// Response format of zabbix api
type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Error   *Error      `json:"error"`
	Result  interface{} `json:"result"`
	ID      int32       `json:"id"`
}

// RawResponse format of zabbix api
type RawResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Error   *Error          `json:"error"`
	Result  json.RawMessage `json:"result"`
	ID      int32           `json:"id"`
}

// Error contains error data and code
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d (%s): %s", e.Code, e.Message, e.Data)
}

// ExpectedOneResult use to generate error when you expect one result
type ExpectedOneResult int

func (e *ExpectedOneResult) Error() string {
	return fmt.Sprintf("Expected exactly one result, got %d.", *e)
}

// ExpectedMore use to generate error when you expect more element
type ExpectedMore struct {
	Expected int
	Got      int
}

func (e *ExpectedMore) Error() string {
	return fmt.Sprintf("Expected %d, got %d.", e.Expected, e.Got)
}

// API use to store connection information
type API struct {
	Auth      string      // auth token, filled by Login()
	Logger    *log.Logger // request/response logger, nil by default
	UserAgent string
	url       string
	c         http.Client
	id        int32
	ex        sync.Mutex
	Config    Config
	
	// Multi-version support
	versionManager *VersionManager
	itemAdapter    ItemAdapter
	hostAdapter    HostAdapter
}

type Config struct {
	Url         string
	TlsNoVerify bool
	Log         *log.Logger
	Serialize   bool
	Version     int
}

// NewAPI Creates new API access object.
// Typical URL is http://host/api_jsonrpc.php or http://host/zabbix/api_jsonrpc.php.
// It also may contain HTTP basic auth username and password like
// http://username:password@host/api_jsonrpc.php.
func NewAPI(c Config) (api *API) {
	api = &API{
		url:            c.Url,
		c:              http.Client{},
		UserAgent:      "github.com/tpretz/go-zabbix-api",
		Logger:         c.Log,
		Config:         c,
		versionManager: NewVersionManager(),
	}

	if c.TlsNoVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		api.c = http.Client{
			Transport: tr,
		}
		api.printf("TLS running in insecure mode, do not use this configuration in production")
	}

	return
}

// InitializeVersionSupport initializes version-specific adapters
func (api *API) InitializeVersionSupport() error {
	if err := api.versionManager.DetectVersion(api); err != nil {
		return fmt.Errorf("failed to detect version: %w", err)
	}
	
	api.printf("Detected Zabbix version: %s", api.versionManager.GetVersion())
	
	// Validate version compatibility
	if !IsCompatibleVersion(api.versionManager.GetVersion()) {
		return fmt.Errorf("unsupported Zabbix version: %s", api.versionManager.GetVersion())
	}
	
	// Initialize adapters based on version
	if api.versionManager.Is70() {
		api.itemAdapter = NewZabbix7ItemAdapter(api)
		api.hostAdapter = NewZabbix7HostAdapter(api)
		api.printf("Initialized Zabbix 7.0 adapters")
	} else if api.versionManager.Is60() {
		api.itemAdapter = NewZabbix6ItemAdapter(api)
		api.hostAdapter = NewZabbix6HostAdapter(api)
		api.printf("Initialized Zabbix 6.0 adapters")
	} else {
		// Default to 6.0 adapters for other 6.x versions
		api.itemAdapter = NewZabbix6ItemAdapter(api)
		api.hostAdapter = NewZabbix6HostAdapter(api)
		api.printf("Initialized Zabbix 6.x compatible adapters for version %s", api.versionManager.GetVersion())
	}
	
	return nil
}

// ForceVersion forces a specific version and initializes adapters accordingly
// Use this for testing or when you know the version in advance
func (api *API) ForceVersion(version string) error {
	api.versionManager.serverVersion = version
	
	// Parse version string
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		major, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("invalid major version: %s", parts[0])
		}
		api.versionManager.majorVersion = major
		
		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid minor version: %s", parts[1])
		}
		api.versionManager.minorVersion = minor
		
		api.versionManager.is70 = major == 7
		api.versionManager.is60 = major == 6
	}
	
	api.versionManager.initializeFeatureSupport()
	
	// Initialize adapters
	if api.versionManager.Is70() {
		api.itemAdapter = NewZabbix7ItemAdapter(api)
		api.hostAdapter = NewZabbix7HostAdapter(api)
	} else {
		api.itemAdapter = NewZabbix6ItemAdapter(api)
		api.hostAdapter = NewZabbix6HostAdapter(api)
	}
	
	api.printf("Forced Zabbix version to: %s", version)
	return nil
}

// GetVersionManager returns the version manager
func (api *API) GetVersionManager() *VersionManager {
	return api.versionManager
}

// GetItemAdapter returns the item adapter
func (api *API) GetItemAdapter() ItemAdapter {
	return api.itemAdapter
}

// GetHostAdapter returns the host adapter
func (api *API) GetHostAdapter() HostAdapter {
	return api.hostAdapter
}

// HistoryPush pushes historical data (Zabbix 7.0+)
func (api *API) HistoryPush(data []HistoryData) error {
	if !api.versionManager.IsFeatureSupported(FeatureHistoryPush) {
		return fmt.Errorf("history.push not supported in Zabbix version %s", api.versionManager.GetVersion())
	}
	
	return api.CallWithErrorParse("history.push", data, nil)
}

// HistoryData represents historical data for history.push API
type HistoryData struct {
	Host   string      `json:"host"`
	Key    string      `json:"key"`
	Value  interface{} `json:"value"`
	Clock  int64       `json:"clock,omitempty"`
	NS     int64       `json:"ns,omitempty"`
}

// ConvertHeadersToV6 converts headers from 7.0 format to 6.0 format
func ConvertHeadersToV6(headersV7 []HeaderField) HttpHeaders {
	headers := make(HttpHeaders)
	for _, header := range headersV7 {
		headers[header.Name] = header.Value
	}
	return headers
}

// ConvertHeadersToV7 converts headers from 6.0 format to 7.0 format
func ConvertHeadersToV7(headersV6 HttpHeaders) []HeaderField {
	headers := make([]HeaderField, 0, len(headersV6))
	for name, value := range headersV6 {
		headers = append(headers, HeaderField{
			Name:  name,
			Value: value,
		})
	}
	return headers
}

// IsCompatibleVersion checks if the given version is compatible with current implementation
func IsCompatibleVersion(version string) bool {
	return strings.HasPrefix(version, "6.") || strings.HasPrefix(version, "7.")
}

// ValidateItemForVersion validates item configuration for specific Zabbix version
func ValidateItemForVersion(item Item, version string) error {
	if strings.HasPrefix(version, "7.") {
		// Zabbix 7.0 validation
		if item.Type == Browser && !strings.HasPrefix(version, "7.") {
			return fmt.Errorf("browser item type not supported in Zabbix version %s", version)
		}
	}
	return nil
}

// ValidateHostForVersion validates host configuration for specific Zabbix version
func ValidateHostForVersion(host Host, version string) error {
	if strings.HasPrefix(version, "7.") {
		// Zabbix 7.0 validation
		if host.ProxyID != "" && host.MonitoredBy == 0 {
			return fmt.Errorf("monitored_by field is required when proxy is specified in Zabbix 7.0")
		}
	}
	return nil
}

// Convenience methods for using adapters

// CreateItems creates items using the appropriate adapter
func (api *API) CreateItems(items Items) error {
	if api.itemAdapter == nil {
		return fmt.Errorf("item adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.itemAdapter.CreateItems(items)
}

// GetItems gets items using the appropriate adapter
func (api *API) GetItems(params Params) (Items, error) {
	if api.itemAdapter == nil {
		return nil, fmt.Errorf("item adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.itemAdapter.GetItems(params)
}

// UpdateItems updates items using the appropriate adapter
func (api *API) UpdateItems(items Items) error {
	if api.itemAdapter == nil {
		return fmt.Errorf("item adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.itemAdapter.UpdateItems(items)
}

// DeleteItems deletes items using the appropriate adapter
func (api *API) DeleteItems(items Items) error {
	if api.itemAdapter == nil {
		return fmt.Errorf("item adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.itemAdapter.DeleteItems(items)
}

// CreateHosts creates hosts using the appropriate adapter
func (api *API) CreateHosts(hosts Hosts) error {
	if api.hostAdapter == nil {
		return fmt.Errorf("host adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.hostAdapter.CreateHosts(hosts)
}

// GetHosts gets hosts using the appropriate adapter
func (api *API) GetHosts(params Params) (Hosts, error) {
	if api.hostAdapter == nil {
		return nil, fmt.Errorf("host adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.hostAdapter.GetHosts(params)
}

// UpdateHosts updates hosts using the appropriate adapter
func (api *API) UpdateHosts(hosts Hosts) error {
	if api.hostAdapter == nil {
		return fmt.Errorf("host adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.hostAdapter.UpdateHosts(hosts)
}

// DeleteHosts deletes hosts using the appropriate adapter
func (api *API) DeleteHosts(hosts Hosts) error {
	if api.hostAdapter == nil {
		return fmt.Errorf("host adapter not initialized. Call Login() or InitializeVersionSupport() first")
	}
	return api.hostAdapter.DeleteHosts(hosts)
}

// IsFeatureSupported checks if a feature is supported by the current Zabbix version
func (api *API) IsFeatureSupported(feature string) bool {
	if api.versionManager == nil {
		return false
	}
	return api.versionManager.IsFeatureSupported(feature)
}

// GetServerVersion returns the detected server version
func (api *API) GetServerVersion() string {
	if api.versionManager == nil {
		return ""
	}
	return api.versionManager.GetVersion()
}

// IsZabbix7 returns true if the server is Zabbix 7.0+
func (api *API) IsZabbix7() bool {
	if api.versionManager == nil {
		return false
	}
	return api.versionManager.Is70()
}

// IsZabbix6 returns true if the server is Zabbix 6.0+
func (api *API) IsZabbix6() bool {
	if api.versionManager == nil {
		return false
	}
	return api.versionManager.Is60()
}

// Enhanced methods for Zabbix 7.0+ features

// CreateMFA creates MFA configuration (Zabbix 7.0+)
func (api *API) CreateMFA(mfas MFAs) error {
	return api.MFACreate(mfas)
}

// GetMFA gets MFA configurations (Zabbix 7.0+)
func (api *API) GetMFA(params Params) (MFAs, error) {
	return api.MFAGet(params)
}

// UpdateMFA updates MFA configuration (Zabbix 7.0+)
func (api *API) UpdateMFA(mfas MFAs) error {
	return api.MFAUpdate(mfas)
}

// DeleteMFA deletes MFA configuration (Zabbix 7.0+)
func (api *API) DeleteMFA(mfas MFAs) error {
	return api.MFADelete(mfas)
}

// ResetUserTOTP resets TOTP for users (Zabbix 7.0+)
func (api *API) ResetUserTOTP(userIDs []string) error {
	return api.UserResetTOTP(userIDs)
}

// CreateProxyGroup creates proxy groups (Zabbix 7.0+)
func (api *API) CreateProxyGroup(proxyGroups ProxyGroups) error {
	return api.ProxyGroupCreate(proxyGroups)
}

// GetProxyGroup gets proxy groups (Zabbix 7.0+)
func (api *API) GetProxyGroup(params Params) (ProxyGroups, error) {
	return api.ProxyGroupGet(params)
}

// UpdateProxyGroup updates proxy groups (Zabbix 7.0+)
func (api *API) UpdateProxyGroup(proxyGroups ProxyGroups) error {
	return api.ProxyGroupUpdate(proxyGroups)
}

// DeleteProxyGroup deletes proxy groups (Zabbix 7.0+)
func (api *API) DeleteProxyGroup(proxyGroups ProxyGroups) error {
	return api.ProxyGroupDelete(proxyGroups)
}

// CreateBrowserItems creates browser items (Zabbix 7.0+)
func (api *API) CreateBrowserItems(items BrowserItems) error {
	return api.CreateBrowserItems(items)
}

// GetBrowserItems gets browser items (Zabbix 7.0+)
func (api *API) GetBrowserItems(params Params) (BrowserItems, error) {
	return api.GetBrowserItems(params)
}

// Utility methods for feature detection

// SupportsHistoryPush returns true if History Push API is supported
func (api *API) SupportsHistoryPush() bool {
	return api.IsFeatureSupported(FeatureHistoryPush)
}

// SupportsMFA returns true if MFA is supported
func (api *API) SupportsMFA() bool {
	return api.IsFeatureSupported(FeatureMFA)
}

// SupportsProxyGroup returns true if Proxy Group is supported
func (api *API) SupportsProxyGroup() bool {
	return api.IsFeatureSupported(FeatureProxyGroup)
}

// SupportsBrowserItem returns true if Browser Item is supported
func (api *API) SupportsBrowserItem() bool {
	return api.IsFeatureSupported(FeatureBrowserItem)
}

// GetSupportedFeatures returns a list of supported features
func (api *API) GetSupportedFeatures() []string {
	if api.versionManager == nil {
		return []string{}
	}
	
	features := []string{
		FeatureHistoryPush,
		FeatureMFA,
		FeatureProxyGroup,
		FeatureBrowserItem,
		FeatureHeadersV7,
		FeatureProxyID,
		FeatureMonitoredBy,
	}
	
	var supported []string
	for _, feature := range features {
		if api.versionManager.IsFeatureSupported(feature) {
			supported = append(supported, feature)
		}
	}
	
	return supported
}

// SetClient Allows one to use specific http.Client, for example with InsecureSkipVerify transport.
func (api *API) SetClient(c *http.Client) {
	api.c = *c
}

func (api *API) printf(format string, v ...interface{}) {
	if api.Logger != nil {
		api.Logger.Printf(format, v...)
	}
}

func (api *API) callBytes(method string, params interface{}) (b []byte, err error) {
	id := atomic.AddInt32(&api.id, 1)
	jsonobj := request{"2.0", method, params, api.Auth, id}
	b, err = json.Marshal(jsonobj)
	if err != nil {
		return
	}
	api.printf("Request (POST): %s", b)

	req, err := http.NewRequest("POST", api.url, bytes.NewReader(b))
	if err != nil {
		return
	}
	req.ContentLength = int64(len(b))
	req.Header.Add("Content-Type", "application/json-rpc")
	req.Header.Add("User-Agent", api.UserAgent)

	if api.Config.Serialize {
		api.ex.Lock()
		defer api.ex.Unlock()
	}

	res, err := api.c.Do(req)
	if err != nil {
		api.printf("Error   : %s", err)
		return
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	api.printf("Response (%d): %s", res.StatusCode, b)
	return
}

// Call Calls specified API method. Uses api.Auth if not empty.
// err is something network or marshaling related. Caller should inspect response.Error to get API error.
func (api *API) Call(method string, params interface{}) (response Response, err error) {
	b, err := api.callBytes(method, params)
	if err == nil {
		err = json.Unmarshal(b, &response)
	}
	return
}

// CallWithError Uses Call() and then sets err to response.Error if former is nil and latter is not.
func (api *API) CallWithError(method string, params interface{}) (response Response, err error) {
	response, err = api.Call(method, params)
	if err == nil && response.Error != nil {
		err = response.Error
	}
	return
}

// CallWithErrorParse Calls specified API method.
// Parse the response of the api in the result variable.
func (api *API) CallWithErrorParse(method string, params interface{}, result interface{}) (err error) {
	var rawResult RawResponse

	response, err := api.callBytes(method, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &rawResult)
	if err != nil {
		return
	}
	if rawResult.Error != nil {
		return rawResult.Error
	}
	err = json.Unmarshal(rawResult.Result, &result)
	return
}

// Login Calls "user.login" API method and fills api.Auth field.
// This method modifies API structure and should not be called concurrently with other methods.
func (api *API) Login(user, password string) (auth string, err error) {
	params := map[string]string{"user": user, "password": password}
	response, err := api.CallWithError("user.login", params)
	if err != nil {
		return
	}

	auth = response.Result.(string)
	api.Auth = auth
	
	// Auto-initialize version support after successful login
	if err := api.InitializeVersionSupport(); err != nil {
		api.printf("Warning: Failed to initialize version support: %v", err)
		// Don't fail login if version detection fails, just log warning
	}
	
	return
}

// LoginWithoutVersionInit Calls "user.login" without automatic version detection.
// Use this when you want to manually control version detection.
func (api *API) LoginWithoutVersionInit(user, password string) (auth string, err error) {
	params := map[string]string{"user": user, "password": password}
	response, err := api.CallWithError("user.login", params)
	if err != nil {
		return
	}

	auth = response.Result.(string)
	api.Auth = auth
	return
}

// Version Calls "APIInfo.version" API method.
// This method temporary modifies API structure and should not be called concurrently with other methods.
func (api *API) Version() (v string, err error) {
	// temporary remove auth for this method to succeed
	// https://www.zabbix.com/documentation/2.2/manual/appendix/api/apiinfo/version
	auth := api.Auth
	api.Auth = ""
	response, err := api.CallWithError("APIInfo.version", Params{})
	api.Auth = auth

	// despite what documentation says, Zabbix 2.2 requires auth, so we try again
	if e, ok := err.(*Error); ok && e.Code == -32602 {
		response, err = api.CallWithError("APIInfo.version", Params{})
	}
	if err != nil {
		return
	}

	v = response.Result.(string)
	return
}
