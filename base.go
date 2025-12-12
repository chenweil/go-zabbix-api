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
	config    Config
}

type Config struct {
	Url         string
	TlsNoVerify bool
	Log         *log.Logger
	Serialize   bool
	Version     int
	// Compression support for Zabbix 6.0
	// Enable compression for API requests and responses
	EnableCompression bool
	// Supported compression encodings (gzip, deflate, identity)
	AcceptedEncodings []string
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
		config:         c,
		versionManager: NewVersionManager(),
	}

	// Configure TLS settings first
	if c.TlsNoVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		api.c.Transport = tr
	}

	// Configure compression for Zabbix 6.0+
	if c.EnableCompression {
		api.configureCompression()
	}

	return
}

// configureCompression sets up compression support for Zabbix 6.0+
func (api *API) configureCompression() {
	// Create a custom transport that handles compression
	transport := &http.Transport{}
	
	// Configure TLS if needed
	if api.config.TlsNoVerify {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	
	// Wrap with compression transport
	api.c.Transport = &compressionTransport{
		transport:        transport,
		acceptedEncodings: api.config.AcceptedEncodings,
		logger:          api.Logger,
	}
}

// compressionTransport handles HTTP compression for Zabbix 6.0+
type compressionTransport struct {
	transport        http.RoundTripper
	acceptedEncodings []string
	logger          *log.Logger
}

func (ct *compressionTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add Accept-Encoding header
	if len(ct.acceptedEncodings) > 0 {
		req.Header.Set("Accept-Encoding", strings.Join(ct.acceptedEncodings, ", "))
	}
	
	// Make the request
	resp, err := ct.transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	
	return resp, nil
}

// Login logs in to the Zabbix API and stores the auth token in the API struct.
func (api *API) Login(user, pass string) (auth string, err error) {
	params := map[string]interface{}{
		"user":     user,
		"password": pass,
	}

	response, err := api.CallWithError("user.login", params)
	if err != nil {
		return
	}

	auth = response.Result.(string)
	api.Auth = auth
	
	// Auto-detect Zabbix version after successful login
	if api.versionManager.serverVersion == "" {
		_, err = api.DetectVersion()
		if err != nil && api.Logger != nil {
			api.Logger.Printf("Warning: Failed to detect Zabbix version: %s", err)
		}
	}

	return
}

// LoginWithToken logs in to the Zabbix API using token parameter (Zabbix 6.0+)
func (api *API) LoginWithToken(user, password, token string) (auth string, err error) {
	params := map[string]interface{}{
		"user":     user,
		"password": password,
		"token":    token,
	}

	response, err := api.CallWithError("user.login", params)
	if err != nil {
		return
	}

	auth = response.Result.(string)
	api.Auth = auth

	// Auto-detect Zabbix version after successful login
	if api.versionManager.serverVersion == "" {
		_, err = api.DetectVersion()
		if err != nil && api.Logger != nil {
			api.Logger.Printf("Warning: Failed to detect Zabbix version: %s", err)
		}
	}

	return
}

// CheckAuthentication validates authentication token (Zabbix 6.0+)
func (api *API) CheckAuthentication(token string) (valid bool, err error) {
	params := map[string]string{"token": token}
	response, err := api.CallWithError("user.checkAuthentication", params)
	if err != nil {
		return
	}

	valid = response.Result.(string) == "true"
	return
}

// Logout logs out from the Zabbix API and clears the auth token.
func (api *API) Logout() (err error) {
	_, err = api.CallWithError("user.logout", nil)
	if err == nil {
		api.Auth = ""
	}
	return
}

// Version returns the version of the Zabbix API.
func (api *API) Version() (version string, err error) {
	response, err := api.CallWithError("apiinfo.version", nil)
	if err != nil {
		return
	}

	version = response.Result.(string)
	return
}

// DetectVersion detects the Zabbix server version and initializes appropriate adapters
func (api *API) DetectVersion() (string, error) {
	version, err := api.Version()
	if err != nil {
		return "", err
	}

	// Parse version and set flags
	api.versionManager.SetVersion(version)
	
	// Initialize appropriate adapters based on version
	api.initializeAdapters()

	return version, nil
}

// initializeAdapters sets up the appropriate adapters based on Zabbix version
func (api *API) initializeAdapters() {
	if api.versionManager.IsZabbix7() {
		api.itemAdapter = &Zabbix7ItemAdapter{api: api}
		api.hostAdapter = &Zabbix7HostAdapter{api: api}
	} else {
		api.itemAdapter = &Zabbix6ItemAdapter{api: api}
		api.hostAdapter = &Zabbix6HostAdapter{api: api}
	}
}

// GetServerVersion returns the detected Zabbix server version
func (api *API) GetServerVersion() string {
	return api.versionManager.serverVersion
}

// IsZabbix7 returns true if connected to Zabbix 7.0+
func (api *API) IsZabbix7() bool {
	return api.versionManager.IsZabbix7()
}

// IsZabbix6 returns true if connected to Zabbix 6.0
func (api *API) IsZabbix6() bool {
	return api.versionManager.IsZabbix6()
}

// SupportsFeature checks if a specific feature is supported by the current Zabbix version
func (api *API) SupportsFeature(feature string) bool {
	return api.versionManager.IsFeatureSupported(feature)
}

// SupportsMFA returns true if MFA is supported (Zabbix 7.0+)
func (api *API) SupportsMFA() bool {
	return api.SupportsFeature(FeatureMFA)
}

// SupportsProxyGroup returns true if Proxy Group API is supported (Zabbix 7.0+)
func (api *API) SupportsProxyGroup() bool {
	return api.SupportsFeature(FeatureProxyGroup)
}

// SupportsHistoryPush returns true if History Push API is supported (Zabbix 7.0+)
func (api *API) SupportsHistoryPush() bool {
	return api.SupportsFeature(FeatureHistoryPush)
}

// SupportsBrowserItem returns true if Browser Item type is supported (Zabbix 7.0+)
func (api *API) SupportsBrowserItem() bool {
	return api.SupportsFeature(FeatureBrowserItem)
}

// CallWithError makes a JSON-RPC call to Zabbix API and returns the raw response.
func (api *API) CallWithError(method string, params interface{}) (response *RawResponse, err error) {
	request := request{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Auth:    api.Auth,
		ID:      atomic.AddInt32(&api.id, 1),
	}

	if api.config.Serialize {
		api.ex.Lock()
		defer api.ex.Unlock()
	}

	var requestBytes []byte
	requestBytes, err = json.Marshal(request)
	if err != nil {
		return
	}

	if api.Logger != nil {
		api.Logger.Printf("Request: %s\n", string(requestBytes))
	}

	httpRequest, err := http.NewRequest("POST", api.url, bytes.NewBuffer(requestBytes))
	if err != nil {
		return
	}

	httpRequest.Header.Set("Content-Type", "application/json-rpc")
	if api.UserAgent != "" {
		httpRequest.Header.Set("User-Agent", api.UserAgent)
	}

	httpResponse, err := api.c.Do(httpRequest)
	if err != nil {
		return
	}
	defer httpResponse.Body.Close()

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	if api.Logger != nil {
		api.Logger.Printf("Response: %s\n", string(responseBytes))
	}

	response = &RawResponse{}
	err = json.Unmarshal(responseBytes, response)
	if err != nil {
		return
	}

	if response.Error != nil {
		err = response.Error
	}

	return
}

// CallWithErrorParse makes a JSON-RPC call to Zabbix API and parses the result into the provided interface.
func (api *API) CallWithErrorParse(method string, params interface{}, result interface{}) (err error) {
	response, err := api.CallWithError(method, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(response.Result, result)
	return
}

// Call makes a JSON-RPC call to Zabbix API and returns the result.
func (api *API) Call(method string, params interface{}) (result interface{}, err error) {
	response, err := api.CallWithError(method, params)
	if err != nil {
		return
	}

	result = response.Result
	return
}

// VersionManager handles Zabbix version detection and feature support
type VersionManager struct {
	serverVersion    string
	majorVersion     int
	minorVersion     int
	isZabbix6        bool
	isZabbix7        bool
	supportedFeatures map[string]bool
}

// NewVersionManager creates a new version manager
func NewVersionManager() *VersionManager {
	return &VersionManager{
		supportedFeatures: make(map[string]bool),
	}
}

// SetVersion sets the Zabbix version and initializes feature support
func (vm *VersionManager) SetVersion(version string) {
	vm.serverVersion = version
	vm.parseVersion()
	vm.initializeFeatureSupport()
}

// parseVersion parses the version string and extracts major/minor versions
func (vm *VersionManager) parseVersion() {
	// Parse version like "6.0.0" or "7.0.1"
	parts := strings.Split(vm.serverVersion, ".")
	if len(parts) >= 2 {
		if major, err := strconv.Atoi(parts[0]); err == nil {
			vm.majorVersion = major
		}
		if minor, err := strconv.Atoi(parts[1]); err == nil {
			vm.minorVersion = minor
		}
	}
	
	vm.isZabbix6 = vm.majorVersion == 6
	vm.isZabbix7 = vm.majorVersion == 7
}

// initializeFeatureSupport sets up feature support based on version
func (vm *VersionManager) initializeFeatureSupport() {
	// Zabbix 6.0 features
	vm.supportedFeatures[FeatureUUID] = vm.isZabbix6 || vm.isZabbix7
	vm.supportedFeatures[FeatureTags] = vm.isZabbix6 || vm.isZabbix7
	vm.supportedFeatures[FeatureCompression] = vm.isZabbix6 || vm.isZabbix7
	vm.supportedFeatures[FeatureHTTPMethods] = vm.isZabbix6 || vm.isZabbix7
	vm.supportedFeatures[FeatureCalculatedItemTypes] = vm.isZabbix6 || vm.isZabbix7
	
	// Zabbix 7.0+ features
	vm.supportedFeatures[FeatureMFA] = vm.isZabbix7
	vm.supportedFeatures[FeatureProxyGroup] = vm.isZabbix7
	vm.supportedFeatures[FeatureHistoryPush] = vm.isZabbix7
	vm.supportedFeatures[FeatureBrowserItem] = vm.isZabbix7
	vm.supportedFeatures[FeatureHeadersArrayFormat] = vm.isZabbix7
	vm.supportedFeatures[FeatureProxyFieldsV7] = vm.isZabbix7
}

// IsZabbix6 returns true if the version is Zabbix 6.x
func (vm *VersionManager) IsZabbix6() bool {
	return vm.isZabbix6
}

// IsZabbix7 returns true if the version is Zabbix 7.x
func (vm *VersionManager) IsZabbix7() bool {
	return vm.isZabbix7
}

// IsFeatureSupported checks if a feature is supported by the current version
func (vm *VersionManager) IsFeatureSupported(feature string) bool {
	supported, exists := vm.supportedFeatures[feature]
	return exists && supported
}

// Feature constants for Zabbix versions
const (
	FeatureUUID                = "uuid"
	FeatureTags                = "tags"
	FeatureCompression         = "compression"
	FeatureHTTPMethods         = "http_methods"
	FeatureCalculatedItemTypes = "calculated_item_types"
	FeatureMFA                 = "mfa"
	FeatureProxyGroup          = "proxy_group"
	FeatureHistoryPush         = "history_push"
	FeatureBrowserItem         = "browser_item"
	FeatureHeadersArrayFormat  = "headers_array_format"
	FeatureProxyFieldsV7       = "proxy_fields_v7"
)

// Adapter interfaces for multi-version support
type ItemAdapter interface {
	CreateItems(items Items) error
	GetItems(params Params) (Items, error)
	UpdateItems(items Items) error
	DeleteItems(itemIds []string) error
}

type HostAdapter interface {
	CreateHosts(hosts Hosts) error
	GetHosts(params Params) (Hosts, error)
	UpdateHosts(hosts Hosts) error
	DeleteHosts(hostIds []string) error
}

// Adapter implementations will be added in respective files
var (
	itemAdapter ItemAdapter
	hostAdapter HostAdapter
	versionManager *VersionManager
)

// Initialize adapters and version manager
func init() {
	versionManager = NewVersionManager()
}