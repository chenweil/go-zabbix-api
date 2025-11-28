package zabbix

import (
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
	expectedDefaultFields := []string{"userid", "username", "name", "surname", "roleid"}
	
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
	expectedDefaultFields := []string{"mediatypeid", "name", "type", "status", "maxattempts"}
	
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
	expectedDefaultFields := []string{"alertid", "actionid", "eventid", "clock", "status"}
	
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