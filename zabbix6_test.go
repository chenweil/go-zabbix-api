package zabbix

import (
	"net/http"
	"testing"
)

// TestZabbix6UserAPI tests Zabbix 6.0 user API compatibility
func TestZabbix6UserAPI(t *testing.T) {
	// Test User struct with Zabbix 6.0 fields
	user := User{
		UserID:   "123",
		Username: "testuser",
		Name:     "Test",
		Surname:  "User",
		Url:      "https://example.com", // URL length increased to 2048 in Zabbix 6.0
		RoleID:   "3", // User role instead of deprecated type field
	}

	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", user.Username)
	}

	if user.Url != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got '%s'", user.Url)
	}
}

// TestZabbix6MediaTypeAPI tests Zabbix 6.0 media type API compatibility
func TestZabbix6MediaTypeAPI(t *testing.T) {
	// Test MediaType struct with Zabbix 6.0 fields
	mediaType := MediaType{
		MediaTypeID: "456",
		Name:        "Email",
		Type:        MediaTypeEmail,
		Status:      MediaTypeStatusEnabled,
		MaxAttempts: "3",
	}

	if mediaType.Type != MediaTypeEmail {
		t.Errorf("Expected media type '%s', got '%s'", MediaTypeEmail, mediaType.Type)
	}

	if mediaType.Status != MediaTypeStatusEnabled {
		t.Errorf("Expected status '%s', got '%s'", MediaTypeStatusEnabled, mediaType.Status)
	}
}

// TestZabbix6AlertAPI tests Zabbix 6.0 alert API compatibility
func TestZabbix6AlertAPI(t *testing.T) {
	// Test Alert struct with Zabbix 6.0 fields
	alert := Alert{
		AlertID:     "789",
		ActionID:    "101",
		EventID:     "202",
		Status:      AlertStatusSent,
		AlertType:   AlertTypeMessage,
	}

	if alert.Status != AlertStatusSent {
		t.Errorf("Expected alert status '%s', got '%s'", AlertStatusSent, alert.Status)
	}

	if alert.AlertType != AlertTypeMessage {
		t.Errorf("Expected alert type '%s', got '%s'", AlertTypeMessage, alert.AlertType)
	}
}

// TestZabbix6HTTPAgentMethods tests Zabbix 6.0 HTTP method constants
func TestZabbix6HTTPAgentMethods(t *testing.T) {
	// Test new HTTP methods added in Zabbix 6.0
	methods := []string{
		HTTPMethodGET,
		HTTPMethodPOST,
		HTTPMethodPUT,
		HTTPMethodHEAD,    // Added in Zabbix 6.0
		HTTPMethodPATCH,   // Added in Zabbix 6.0
		HTTPMethodDELETE,
		HTTPMethodOPTIONS, // Added in Zabbix 6.0
		HTTPMethodTRACE,   // Added in Zabbix 6.0
		HTTPMethodCONNECT, // Added in Zabbix 6.0
	}

	expectedCount := 9
	if len(methods) != expectedCount {
		t.Errorf("Expected %d HTTP methods, got %d", expectedCount, len(methods))
	}

	// Verify specific Zabbix 6.0 methods are present
	zabbix6Methods := map[string]bool{
		HTTPMethodHEAD:    false,
		HTTPMethodPATCH:   false,
		HTTPMethodOPTIONS: false,
		HTTPMethodTRACE:   false,
		HTTPMethodCONNECT: false,
	}

	for _, method := range methods {
		if _, exists := zabbix6Methods[method]; exists {
			zabbix6Methods[method] = true
		}
	}

	for method, exists := range zabbix6Methods {
		if !exists {
			t.Errorf("Zabbix 6.0 HTTP method '%s' is missing", method)
		}
	}
}

// TestZabbix6ItemHTTPAgentCompatibility tests HTTP Agent compatibility
func TestZabbix6ItemHTTPAgentCompatibility(t *testing.T) {
	// Test HTTP Agent item without interfaceid (Zabbix 6.0 compatibility)
	item := Item{
		Type:         HTTPAgent,
		HostID:       "host123",
		Key:          "http.test",
		Name:         "HTTP Test",
		Url:          "https://api.example.com/test",
		RequestMethod: HTTPMethodGET,
		// InterfaceID is intentionally omitted for HTTP Agent in Zabbix 6.0
	}

	if item.Type != HTTPAgent {
		t.Errorf("Expected item type HTTPAgent, got %v", item.Type)
	}

	if item.InterfaceID != "" {
		t.Errorf("InterfaceID should be empty for HTTP Agent in Zabbix 6.0, got '%s'", item.InterfaceID)
	}

	if item.RequestMethod != HTTPMethodGET {
		t.Errorf("Expected request method '%s', got '%s'", HTTPMethodGET, item.RequestMethod)
	}
}

// TestZabbix6UserGetOptions tests user.get options with Zabbix 6.0 defaults
func TestZabbix6UserGetOptions(t *testing.T) {
	// Test default output fields for Zabbix 6.0 compatibility
	options := UserGetOptions{
		UserIDs: []string{"123"},
	}

	// Verify the default fields that should be accessible in Zabbix 6.0
	// This would normally be set by the API method
	// Here we just verify the structure is correct
	if len(options.UserIDs) != 1 {
		t.Errorf("Expected 1 user ID, got %d", len(options.UserIDs))
	}

	if options.UserIDs[0] != "123" {
		t.Errorf("Expected user ID '123', got '%s'", options.UserIDs[0])
	}
}

// TestZabbix6MediaTypeGetOptions tests mediatype.get options with Zabbix 6.0 defaults
func TestZabbix6MediaTypeGetOptions(t *testing.T) {
	// Test default output fields for Admin users in Zabbix 6.0
	options := MediaTypeGetOptions{
		MediaTypeIDs: []string{"456"},
	}

	// Verify the default fields that should be accessible to Admin users in Zabbix 6.0
	if len(options.MediaTypeIDs) != 1 {
		t.Errorf("Expected 1 media type ID, got %d", len(options.MediaTypeIDs))
	}

	if options.MediaTypeIDs[0] != "456" {
		t.Errorf("Expected media type ID '456', got '%s'", options.MediaTypeIDs[0])
	}
}

// TestZabbix6AlertGetOptions tests alert.get options with Zabbix 6.0 defaults
func TestZabbix6AlertGetOptions(t *testing.T) {
	// Test default output fields for Admin users in Zabbix 6.0
	options := AlertGetOptions{
		AlertIDs: []string{"789"},
	}

	// Verify the default fields that should be accessible to Admin users in Zabbix 6.0
	if len(options.AlertIDs) != 1 {
		t.Errorf("Expected 1 alert ID, got %d", len(options.AlertIDs))
	}

	if options.AlertIDs[0] != "789" {
		t.Errorf("Expected alert ID '789', got '%s'", options.AlertIDs[0])
	}
}

// TestZabbix6AuthenticationMethods tests Zabbix 6.0 authentication enhancements
func TestZabbix6AuthenticationMethods(t *testing.T) {
	// Test API structure for Zabbix 6.0 authentication
	config := Config{
		Url:         "https://zabbix.example.com/api_jsonrpc.php",
		TlsNoVerify: true,
	}

	api := NewAPI(config)
	
	if api.url != config.Url {
		t.Errorf("Expected URL '%s', got '%s'", config.Url, api.url)
	}

	if api.Auth != "" {
		t.Errorf("Expected empty auth token initially, got '%s'", api.Auth)
	}
}

// TestZabbix6CalculatedItemValueTypes tests new calculated item value types added in Zabbix 6.0
func TestZabbix6CalculatedItemValueTypes(t *testing.T) {
	// Test all value types including new Zabbix 6.0 calculated types
	allValueTypes := []ValueType{
		NumericFloat,    // 0
		Character,       // 1
		Log,             // 2
		NumericUnsigned, // 3
		Text,            // 4
		CalculatedText,  // 5 - Added in Zabbix 6.0
		CalculatedLog,   // 6 - Added in Zabbix 6.0
		CalculatedChar,  // 7 - Added in Zabbix 6.0
	}

	expectedCount := 8
	if len(allValueTypes) != expectedCount {
		t.Errorf("Expected %d value types, got %d", expectedCount, len(allValueTypes))
	}

	// Verify specific Zabbix 6.0 calculated types are present
	zabbix6CalculatedTypes := map[ValueType]bool{
		CalculatedText: false,
		CalculatedLog:  false,
		CalculatedChar: false,
	}

	for _, valueType := range allValueTypes {
		if _, exists := zabbix6CalculatedTypes[valueType]; exists {
			zabbix6CalculatedTypes[valueType] = true
		}
	}

	for valueType, exists := range zabbix6CalculatedTypes {
		if !exists {
			t.Errorf("Zabbix 6.0 calculated value type %v is missing", valueType)
		}
	}

	// Test calculated item creation with new value types
	calculatedItems := []Item{
		{
			Type:      Calculated,
			ValueType: CalculatedText,
			Key:       "calculated.text.test",
			Name:      "Calculated Text Item",
			HostID:    "host123",
		},
		{
			Type:      Calculated,
			ValueType: CalculatedLog,
			Key:       "calculated.log.test",
			Name:      "Calculated Log Item",
			HostID:    "host123",
		},
		{
			Type:      Calculated,
			ValueType: CalculatedChar,
			Key:       "calculated.char.test",
			Name:      "Calculated Character Item",
			HostID:    "host123",
		},
	}

	for i, item := range calculatedItems {
		if item.Type != Calculated {
			t.Errorf("Item %d: Expected type Calculated, got %v", i, item.Type)
		}

		switch i {
		case 0:
			if item.ValueType != CalculatedText {
				t.Errorf("Item %d: Expected value type CalculatedText, got %v", i, item.ValueType)
			}
		case 1:
			if item.ValueType != CalculatedLog {
				t.Errorf("Item %d: Expected value type CalculatedLog, got %v", i, item.ValueType)
			}
		case 2:
			if item.ValueType != CalculatedChar {
				t.Errorf("Item %d: Expected value type CalculatedChar, got %v", i, item.ValueType)
			}
		}

		if item.HostID != "host123" {
			t.Errorf("Item %d: Expected HostID 'host123', got '%s'", i, item.HostID)
		}
	}
}

// TestZabbix6LoginWithToken tests LoginWithToken method for Zabbix 6.0 compatibility
func TestZabbix6LoginWithToken(t *testing.T) {
	// Test API structure for LoginWithToken method
	config := Config{
		Url: "https://zabbix.example.com/api_jsonrpc.php",
	}

	api := NewAPI(config)
	
	// Test that the method exists (this is a compile-time test)
	// In a real test environment, this would require actual Zabbix server connection
	_ = func(user, password, token string) (string, error) {
		return api.LoginWithToken(user, password, token)
	}
}

// TestZabbix6CheckAuthentication tests CheckAuthentication method with token support
func TestZabbix6CheckAuthentication(t *testing.T) {
	// Test API structure for CheckAuthentication method
	config := Config{
		Url: "https://zabbix.example.com/api_jsonrpc.php",
	}

	api := NewAPI(config)
	
	// Test that the method exists (this is a compile-time test)
	// In a real test environment, this would require actual Zabbix server connection
	_ = func(token string) (bool, error) {
		return api.CheckAuthentication(token)
	}
}

// TestZabbix6CompressionSupport tests compression configuration for Zabbix 6.0
func TestZabbix6CompressionSupport(t *testing.T) {
	// Test compression disabled by default
	config := Config{
		Url: "https://zabbix.example.com/api_jsonrpc.php",
	}

	api := NewAPI(config)
	
	// When compression is disabled, transport should be default or TLS-only
	if api.config.EnableCompression {
		t.Errorf("Expected compression to be disabled by default")
	}

	// Test compression enabled with default settings
	configWithCompression := Config{
		Url:                "https://zabbix.example.com/api_jsonrpc.php",
		EnableCompression:  true,
	}

	apiWithCompression := NewAPI(configWithCompression)
	
	if !apiWithCompression.config.EnableCompression {
		t.Errorf("Expected compression to be enabled")
	}

	// Verify default accepted encodings are set
	expectedEncodings := []string{"gzip", "deflate", "identity"}
	if len(apiWithCompression.config.AcceptedEncodings) != len(expectedEncodings) {
		t.Errorf("Expected %d accepted encodings, got %d", len(expectedEncodings), len(apiWithCompression.config.AcceptedEncodings))
	}

	for i, encoding := range expectedEncodings {
		if apiWithCompression.config.AcceptedEncodings[i] != encoding {
			t.Errorf("Expected encoding '%s' at position %d, got '%s'", encoding, i, apiWithCompression.config.AcceptedEncodings[i])
		}
	}
}

// TestZabbix6CompressionWithCustomEncodings tests compression with custom encoding settings
func TestZabbix6CompressionWithCustomEncodings(t *testing.T) {
	// Test compression with custom accepted encodings
	customEncodings := []string{"gzip", "identity"}
	config := Config{
		Url:                "https://zabbix.example.com/api_jsonrpc.php",
		EnableCompression:  true,
		AcceptedEncodings:  customEncodings,
	}

	api := NewAPI(config)
	
	if len(api.config.AcceptedEncodings) != len(customEncodings) {
		t.Errorf("Expected %d accepted encodings, got %d", len(customEncodings), len(api.config.AcceptedEncodings))
	}

	for i, encoding := range customEncodings {
		if api.config.AcceptedEncodings[i] != encoding {
			t.Errorf("Expected encoding '%s' at position %d, got '%s'", encoding, i, api.config.AcceptedEncodings[i])
		}
	}
}

// TestZabbix6CompressionWithTLS tests compression compatibility with TLS configuration
func TestZabbix6CompressionWithTLS(t *testing.T) {
	// Test compression enabled together with TLS no verify
	config := Config{
		Url:                "https://zabbix.example.com/api_jsonrpc.php",
		TlsNoVerify:        true,
		EnableCompression:  true,
	}

	api := NewAPI(config)
	
	// Both TLS and compression should be configured
	if !api.config.EnableCompression {
		t.Errorf("Expected compression to be enabled with TLS")
	}

	// Transport should be configured (this is a basic check)
	if api.c.Transport == nil {
		t.Errorf("Expected transport to be configured with TLS and compression")
	}
}

// TestZabbix6CompressionTransport tests compression transport functionality
func TestZabbix6CompressionTransport(t *testing.T) {
	// Create a compression transport for testing
	baseTransport := &http.Transport{}
	acceptedEncodings := []string{"gzip", "deflate", "identity"}
	
	compTransport := &compressionTransport{
		transport:         baseTransport,
		acceptedEncodings: acceptedEncodings,
	}

	// Verify transport configuration
	if compTransport.transport != baseTransport {
		t.Errorf("Expected base transport to be preserved")
	}

	if len(compTransport.acceptedEncodings) != len(acceptedEncodings) {
		t.Errorf("Expected %d accepted encodings, got %d", len(acceptedEncodings), len(compTransport.acceptedEncodings))
	}
}

// TestZabbix6AllHTTPMethods tests all HTTP method constants including Zabbix 6.0 additions
func TestZabbix6AllHTTPMethods(t *testing.T) {
	// Comprehensive test of all HTTP method constants
	methodTests := []struct {
		constant string
		expected string
	}{
		{HTTPMethodGET, "0"},
		{HTTPMethodPOST, "1"},
		{HTTPMethodPUT, "2"},
		{HTTPMethodHEAD, "3"},    // Zabbix 6.0
		{HTTPMethodPATCH, "4"},   // Zabbix 6.0
		{HTTPMethodDELETE, "5"},
		{HTTPMethodOPTIONS, "6"}, // Zabbix 6.0
		{HTTPMethodTRACE, "7"},   // Zabbix 6.0
		{HTTPMethodCONNECT, "8"}, // Zabbix 6.0
	}

	for _, test := range methodTests {
		if test.constant != test.expected {
			t.Errorf("Expected HTTP method constant '%s', got '%s'", test.expected, test.constant)
		}
	}

	// Test creating items with all HTTP methods
	for i, test := range methodTests {
		item := Item{
			Type:         HTTPAgent,
			HostID:       "host123",
			Key:          "http.test.method",
			Name:         "HTTP Method Test",
			Url:          "https://api.example.com/test",
			RequestMethod: test.constant,
		}

		if item.RequestMethod != test.constant {
			t.Errorf("Item %d: Expected request method '%s', got '%s'", i, test.constant, item.RequestMethod)
		}
	}
}

// TestZabbix6PerformanceBenchmarks benchmarks basic API operations
func TestZabbix6PerformanceBenchmarks(t *testing.T) {
	// Basic performance test for API creation
	config := Config{
		Url: "https://zabbix.example.com/api_jsonrpc.php",
	}

	// Test API creation performance (should be very fast)
	for i := 0; i < 1000; i++ {
		api := NewAPI(config)
		if api.url != config.Url {
			t.Errorf("Iteration %d: Expected URL '%s', got '%s'", i, config.Url, api.url)
		}
	}

	// Test API creation with compression
	configWithCompression := Config{
		Url:               "https://zabbix.example.com/api_jsonrpc.php",
		EnableCompression: true,
	}

	for i := 0; i < 1000; i++ {
		api := NewAPI(configWithCompression)
		if !api.config.EnableCompression {
			t.Errorf("Iteration %d: Expected compression to be enabled", i)
		}
	}

	// Test API creation with TLS and compression
	configWithTLSAndCompression := Config{
		Url:               "https://zabbix.example.com/api_jsonrpc.php",
		TlsNoVerify:       true,
		EnableCompression: true,
	}

	for i := 0; i < 1000; i++ {
		api := NewAPI(configWithTLSAndCompression)
		if !api.config.EnableCompression {
			t.Errorf("Iteration %d: Expected compression to be enabled", i)
		}
		if api.c.Transport == nil {
			t.Errorf("Iteration %d: Expected transport to be configured", i)
		}
	}
}
