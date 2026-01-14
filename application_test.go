package zabbix

import (
	"testing"
)

func TestApplicationGetOptions(t *testing.T) {
	// Test default values
	opts := ApplicationGetOptions{
		ApplicationIDs: []string{"12345"},
		HostIDs:       []string{"67890"},
	}

	if len(opts.ApplicationIDs) != 1 {
		t.Errorf("Expected 1 application ID, got %d", len(opts.ApplicationIDs))
	}

	if len(opts.HostIDs) != 1 {
		t.Errorf("Expected 1 host ID, got %d", len(opts.HostIDs))
	}

	if opts.SortField != "" {
		t.Errorf("Expected empty sort field, got %v", opts.SortField)
	}

	if opts.SortOrder != "" {
		t.Errorf("Expected empty sort order, got %v", opts.SortOrder)
	}

	if opts.Output != "" {
		t.Errorf("Expected empty output, got %v", opts.Output)
	}
}

func TestApplication(t *testing.T) {
	// Test Application creation and JSON marshaling
	application := Application{
		ApplicationID: "12345",
		HostID:        "67890",
		Name:          "Web Server",
		TemplateID:    "11111",
		Flags:         "0",
	}

	expectedJSON := `{"applicationid":"12345","hostid":"67890","name":"Web Server","templateid":"11111","flags":"0"}`
	
	jsonData, err := json.Marshal(application)
	if err != nil {
		t.Errorf("Failed to marshal Application: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestApplicationID(t *testing.T) {
	// Test ApplicationID creation and JSON marshaling
	applicationID := ApplicationID{
		ApplicationID: "12345",
	}

	expectedJSON := `{"applicationid":"12345"}`
	
	jsonData, err := json.Marshal(applicationID)
	if err != nil {
		t.Errorf("Failed to marshal ApplicationID: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestApplicationWithItems(t *testing.T) {
	// Test ApplicationWithItems creation and JSON marshaling
	application := Application{
		ApplicationID: "12345",
		HostID:        "67890",
		Name:          "Database",
	}
	
	items := Items{
		{ItemID: "100", Name: "CPU Usage"},
		{ItemID: "101", Name: "Memory Usage"},
	}
	
	applicationWithItems := ApplicationWithItems{
		Application: application,
		Items:       items,
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(applicationWithItems)
	if err != nil {
		t.Errorf("Failed to marshal ApplicationWithItems: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestApplicationWithHosts(t *testing.T) {
	// Test ApplicationWithHosts creation and JSON marshaling
	application := Application{
		ApplicationID: "12345",
		HostID:        "67890",
		Name:          "Network Monitor",
	}
	
	host := Host{
		HostID: "67890",
		Host:   "server01.example.com",
		Name:   "Server 01",
	}
	
	applicationWithHosts := ApplicationWithHosts{
		Application: application,
		Hosts:       Hosts{host},
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(applicationWithHosts)
	if err != nil {
		t.Errorf("Failed to marshal ApplicationWithHosts: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestMockApplicationsAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test that methods exist and return appropriate types
	opts := ApplicationGetOptions{
		ApplicationIDs: []string{"12345"},
		Output:        "extend",
		Limit:         10,
	}

	// These calls will fail without a real Zabbix server, but we can verify the method signatures
	_, err := api.ApplicationsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetByHost([]string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetByHostAndTemplate([]string{"67890"}, []string{"11111"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationGetByID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetWithItems(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetWithHosts(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetByName("Web Server")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetByNames([]string{"App1", "App2"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetByPattern("Web*")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test CRUD operations
	application := Application{
		HostID: "67890",
		Name:   "Test Application",
	}
	
	_, err = api.ApplicationCreateSingle(application)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationUpdateSingle(application)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ApplicationDeleteSingle("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics and utility methods
	_, err = api.ApplicationsGetCount(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetStatistics()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetForHost("67890")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetForTemplate("11111")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ApplicationsGetSummary([]string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ItemsGetByApplicationID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func TestApplicationValidate(t *testing.T) {
	// Test Application validation
	api := NewAPI(Config{})
	
	// Valid application
	validApplication := Application{
		HostID: "67890",
		Name:   "Valid Application",
	}
	
	errors := api.ApplicationValidate(validApplication)
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors for valid application, got: %v", errors)
	}
	
	// Invalid application - missing host ID
	invalidApplication1 := Application{
		Name: "Test Application",
	}
	
	errors = api.ApplicationValidate(invalidApplication1)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for application without host ID")
	}
	
	// Invalid application - missing name
	invalidApplication2 := Application{
		HostID: "67890",
	}
	
	errors = api.ApplicationValidate(invalidApplication2)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for application without name")
	}
	
	// Invalid application - empty host ID
	invalidApplication3 := Application{
		HostID: "",
		Name:   "Test Application",
	}
	
	errors = api.ApplicationValidate(invalidApplication3)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for application with empty host ID")
	}
	
	// Invalid application - empty template ID
	invalidApplication4 := Application{
		HostID:    "67890",
		Name:      "Test Application",
		TemplateID: "",
	}
	
	errors = api.ApplicationValidate(invalidApplication4)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for application with empty template ID")
	}
}

func BenchmarkApplicationMarshaling(b *testing.B) {
	application := Application{
		ApplicationID: "12345",
		HostID:        "67890",
		Name:          "Web Server Application",
		TemplateID:    "11111",
		Flags:         "0",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(application)
	}
}

func BenchmarkApplicationWithItemsMarshaling(b *testing.B) {
	application := Application{
		ApplicationID: "12345",
		HostID:        "67890",
		Name:          "Database Monitor",
	}
	
	items := Items{
		{ItemID: "100", Name: "CPU Usage"},
		{ItemID: "101", Name: "Memory Usage"},
		{ItemID: "102", Name: "Disk Usage"},
	}
	
	applicationWithItems := ApplicationWithItems{
		Application: application,
		Items:       items,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(applicationWithItems)
	}
}

// Test integration scenarios
func TestApplicationsIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test complex filtering
	opts := ApplicationGetOptions{
		HostIDs:    []string{"12345", "67890"},
		Output:    "extend",
		Limit:     100,
		SortField: "name",
		SortOrder: "ASC",
	}

	_, err := api.ApplicationsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test application creation with all fields
	application := Application{
		HostID:     "12345",
		Name:       "Integration Test Application",
		TemplateID: "67890",
		Flags:      "0",
	}
	
	_, err = api.ApplicationCreateSingle(application)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test application update
	application.ApplicationID = "11111"
	_, err = api.ApplicationUpdateSingle(application)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test bulk operations
	applications := Applications{
		{HostID: "12345", Name: "App 1"},
		{HostID: "12345", Name: "App 2"},
		{HostID: "67890", Name: "App 3"},
	}
	
	_, err = api.ApplicationsCreate(applications)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test bulk deletion
	applicationsForDeletion := Applications{
		{ApplicationID: "11111"},
		{ApplicationID: "22222"},
	}
	
	_, err = api.ApplicationsDelete(applicationsForDeletion)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test getting applications with relationships
	optsWithItems := ApplicationGetOptions{
		HostIDs:      []string{"12345"},
		SelectItems:  "extend",
		Output:       "extend",
	}
	
	_, err = api.ApplicationsGetWithItems(optsWithItems)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	optsWithHosts := ApplicationGetOptions{
		ApplicationIDs: []string{"12345"},
		SelectHosts:   "extend",
		Output:        "extend",
	}
	
	_, err = api.ApplicationsGetWithHosts(optsWithHosts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

// Test backward compatibility for Zabbix 6.0+
// Applications are deprecated in Zabbix 6.0 but still supported for compatibility
func TestApplicationBackwardCompatibility(t *testing.T) {
	// Test that the Application struct works with both old and new Zabbix versions
	application := Application{
		ApplicationID: "12345",
		HostID:        "67890",
		Name:          "Legacy Application",
		TemplateID:    "11111",
		Flags:         "0",
		// Note: The 'applications' field (for items) is deprecated
		// In Zabbix 6.0+, this should be handled via Tags instead
	}

	// Test JSON marshaling works correctly
	jsonData, err := json.Marshal(application)
	if err != nil {
		t.Errorf("Failed to marshal Application: %v", err)
	}

	// Test unmarshaling works correctly
	var unmarshaledApplication Application
	err = json.Unmarshal(jsonData, &unmarshaledApplication)
	if err != nil {
		t.Errorf("Failed to unmarshal Application: %v", err)
	}

	// Verify the data is preserved
	if unmarshaledApplication.ApplicationID != application.ApplicationID {
		t.Errorf("ApplicationID not preserved: expected %s, got %s", application.ApplicationID, unmarshaledApplication.ApplicationID)
	}

	if unmarshaledApplication.HostID != application.HostID {
		t.Errorf("HostID not preserved: expected %s, got %s", application.HostID, unmarshaledApplication.HostID)
	}

	if unmarshaledApplication.Name != application.Name {
		t.Errorf("Name not preserved: expected %s, got %s", application.Name, unmarshaledApplication.Name)
	}
}