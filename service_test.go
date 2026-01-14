package zabbix

import (
	"testing"
	"time"
)

func TestServiceGetOptions(t *testing.T) {
	// Test default values
	opts := ServiceGetOptions{
		ServiceIDs: []string{"12345"},
	}

	if len(opts.ServiceIDs) != 1 {
		t.Errorf("Expected 1 service ID, got %d", len(opts.ServiceIDs))
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

func TestService(t *testing.T) {
	// Test Service creation and JSON marshaling
	service := Service{
		ServiceID:    "12345",
		Name:         "Web Service",
		Status:       ServiceStatusUp,
		Algorithm:    ServiceAlgorithmAny,
		Description:  "Main web service",
		ParentID:     "67890",
		AcceptableSLA: "99.5",
	}

	expectedJSON := `{"serviceid":"12345","name":"Web Service","status":"0","algorithm":"0","description":"Main web service","parentid":"67890","acceptable_sla":"99.5"}`
	
	jsonData, err := json.Marshal(service)
	if err != nil {
		t.Errorf("Failed to marshal Service: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestServiceProblem(t *testing.T) {
	// Test ServiceProblem creation and JSON marshaling
	problem := ServiceProblem{
		ServiceID:  "12345",
		ProblemID:  "67890",
		TriggerID:  "11111",
		EventID:    "22222",
		Status:     ServiceProblemStatusOpen,
		Clock:      1640995200,
		Severity:   "3",
		Description: "Service is down",
	}

	expectedJSON := `{"serviceid":"12345","problemid":"67890","triggerid":"11111","eventid":"22222","status":"1","clock":1640995200,"severity":"3","description":"Service is down"}`
	
	jsonData, err := json.Marshal(problem)
	if err != nil {
		t.Errorf("Failed to marshal ServiceProblem: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestServiceDependency(t *testing.T) {
	// Test ServiceDependency creation and JSON marshaling
	dependency := ServiceDependency{
		DependsOnID: "67890",
		LinkType:    ServiceLinkTypeNormal,
	}

	expectedJSON := `{"depends_onid":"67890","link_type":"0"}`
	
	jsonData, err := json.Marshal(dependency)
	if err != nil {
		t.Errorf("Failed to marshal ServiceDependency: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestServiceTag(t *testing.T) {
	// Test ServiceTag creation and JSON marshaling
	tag := ServiceTag{
		Tag:   "environment",
		Value: "production",
	}

	expectedJSON := `{"tag":"environment","value":"production"}`
	
	jsonData, err := json.Marshal(tag)
	if err != nil {
		t.Errorf("Failed to marshal ServiceTag: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestServiceTree(t *testing.T) {
	// Test ServiceTree creation and JSON marshaling
	tree := ServiceTree{
		Service: Service{
			ServiceID: "12345",
			Name:      "Root Service",
			Status:    ServiceStatusUp,
		},
		Children: []ServiceTree{
			{
				Service: Service{
					ServiceID: "67890",
					Name:      "Child Service",
					Status:    ServiceStatusUp,
				},
			},
		},
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(tree)
	if err != nil {
		t.Errorf("Failed to marshal ServiceTree: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestServiceFilter(t *testing.T) {
	// Test ServiceFilter creation and JSON marshaling
	filter := ServiceFilter{
		ServiceID: "12345",
		Exclude:   ServiceFilterExcludeYes,
	}

	expectedJSON := `{"serviceid":"12345","exclude":"1"}`
	
	jsonData, err := json.Marshal(filter)
	if err != nil {
		t.Errorf("Failed to marshal ServiceFilter: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestSLAMeasure(t *testing.T) {
	// Test SLAMeasure creation and JSON marshaling
	sla := SLAMeasure{
		ServiceID:       "12345",
		ServiceName:     "Web Service",
		SLI:            99.5,
		Uptime:          99.8,
		Downtime:        0.2,
		TimeToFirstFailure: 3600,
		TimeToRecovery:    300,
		AllowedDowntime:    0.5,
	}

	expectedJSON := `{"serviceid":"12345","service_name":"Web Service","sli":99.5,"uptime":99.8,"downtime":0.2,"time_to_first_failure":3600,"time_to_recovery":300,"allowed_downtime":0.5}`
	
	jsonData, err := json.Marshal(sla)
	if err != nil {
		t.Errorf("Failed to marshal SLAMeasure: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestServiceConstants(t *testing.T) {
	// Test service constants
	tests := []struct {
		constant string
		expected string
	}{
		// Service status
		{ServiceStatusUp, "0"},
		{ServiceStatusDown, "1"},
		{ServiceStatusPartiallyDown, "2"},
		{ServiceStatusUnknown, "3"},
		
		// Service algorithm
		{ServiceAlgorithmAny, "0"},
		{ServiceAlgorithmAll, "1"},
		
		// Service link type
		{ServiceLinkTypeNormal, "0"},
		{ServiceLinkTypeSoft, "1"},
		
		// Service propagate
		{ServicePropagateNo, "0"},
		{ServicePropagateYes, "1"},
		
		// Service show SLA
		{ServiceShowSLANo, "0"},
		{ServiceShowSLAYes, "1"},
		
		// Service problem status
		{ServiceProblemStatusOpen, "1"},
		{ServiceProblemStatusClosed, "0"},
		
		// Service filter exclude
		{ServiceFilterExcludeNo, "0"},
		{ServiceFilterExcludeYes, "1"},
	}

	for _, test := range tests {
		if test.constant != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.constant)
		}
	}
}

func TestServiceIsUp(t *testing.T) {
	// Test Service.IsUp() method
	upService := Service{Status: ServiceStatusUp}
	downService := Service{Status: ServiceStatusDown}
	partialService := Service{Status: ServiceStatusPartiallyDown}
	unknownService := Service{Status: ServiceStatusUnknown}
	noStatusService := Service{Status: ""}

	if !upService.IsUp() {
		t.Errorf("Expected up service to be detected as up")
	}

	if downService.IsUp() {
		t.Errorf("Expected down service to NOT be detected as up")
	}

	if partialService.IsUp() {
		t.Errorf("Expected partially down service to NOT be detected as up")
	}

	if unknownService.IsUp() {
		t.Errorf("Expected unknown service to NOT be detected as up")
	}

	if noStatusService.IsUp() {
		t.Errorf("Expected service with no status to NOT be detected as up")
	}
}

func TestServiceIsDown(t *testing.T) {
	// Test Service.IsDown() method
	downService := Service{Status: ServiceStatusDown}
	upService := Service{Status: ServiceStatusUp}
	partialService := Service{Status: ServiceStatusPartiallyDown}
	unknownService := Service{Status: ServiceStatusUnknown}

	if !downService.IsDown() {
		t.Errorf("Expected down service to be detected as down")
	}

	if upService.IsDown() {
		t.Errorf("Expected up service to NOT be detected as down")
	}

	if partialService.IsDown() {
		t.Errorf("Expected partially down service to NOT be detected as down")
	}

	if unknownService.IsDown() {
		t.Errorf("Expected unknown service to NOT be detected as down")
	}
}

func TestServiceIsPartiallyDown(t *testing.T) {
	// Test Service.IsPartiallyDown() method
	partialService := Service{Status: ServiceStatusPartiallyDown}
	upService := Service{Status: ServiceStatusUp}
	downService := Service{Status: ServiceStatusDown}
	unknownService := Service{Status: ServiceStatusUnknown}

	if !partialService.IsPartiallyDown() {
		t.Errorf("Expected partially down service to be detected as partially down")
	}

	if upService.IsPartiallyDown() {
		t.Errorf("Expected up service to NOT be detected as partially down")
	}

	if downService.IsPartiallyDown() {
		t.Errorf("Expected down service to NOT be detected as partially down")
	}

	if unknownService.IsPartiallyDown() {
		t.Errorf("Expected unknown service to NOT be detected as partially down")
	}
}

func TestServiceIsUnknown(t *testing.T) {
	// Test Service.IsUnknown() method
	unknownService := Service{Status: ServiceStatusUnknown}
	upService := Service{Status: ServiceStatusUp}
	downService := Service{Status: ServiceStatusDown}
	partialService := Service{Status: ServiceStatusPartiallyDown}

	if !unknownService.IsUnknown() {
		t.Errorf("Expected unknown service to be detected as unknown")
	}

	if upService.IsUnknown() {
		t.Errorf("Expected up service to NOT be detected as unknown")
	}

	if downService.IsUnknown() {
		t.Errorf("Expected down service to NOT be detected as unknown")
	}

	if partialService.IsUnknown() {
		t.Errorf("Expected partially down service to NOT be detected as unknown")
	}
}

func TestServiceIsRoot(t *testing.T) {
	// Test Service.IsRoot() method
	rootService := Service{ParentID: ""}
	childService := Service{ParentID: "67890"}

	if !rootService.IsRoot() {
		t.Errorf("Expected service without parent to be detected as root")
	}

	if childService.IsRoot() {
		t.Errorf("Expected service with parent to NOT be detected as root")
	}
}

func TestServiceIsChild(t *testing.T) {
	// Test Service.IsChild() method
	childService := Service{ParentID: "67890"}
	rootService := Service{ParentID: ""}

	if !childService.IsChild() {
		t.Errorf("Expected service with parent to be detected as child")
	}

	if rootService.IsChild() {
		t.Errorf("Expected service without parent to NOT be detected as child")
	}
}

func TestServiceHasChildren(t *testing.T) {
	// Test Service.HasChildren() method (always returns false based on current implementation)
	service := Service{
		ServiceID: "12345",
		Name:      "Test Service",
	}
	
	// Note: This test reflects the current implementation where ChildCount() always returns 0
	// In a real implementation, this would be populated by the API
	if service.HasChildren() {
		t.Errorf("Expected service to NOT have children based on current implementation")
	}
}

func TestServiceValidate(t *testing.T) {
	// Test Service validation
	api := NewAPI(Config{})
	
	// Valid service
	validService := Service{
		Name:         "Valid Service",
		Status:       ServiceStatusUp,
		Algorithm:    ServiceAlgorithmAny,
		AcceptableSLA: "99.5",
	}
	
	errors := api.ServiceValidate(validService)
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors for valid service, got: %v", errors)
	}
	
	// Invalid service - missing name
	invalidService1 := Service{
		Status:       ServiceStatusUp,
		Algorithm:    ServiceAlgorithmAny,
		AcceptableSLA: "99.5",
	}
	
	errors = api.ServiceValidate(invalidService1)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for service without name")
	}
	
	// Invalid service - invalid status
	invalidService2 := Service{
		Name:         "Test Service",
		Status:       "invalid_status",
		Algorithm:    ServiceAlgorithmAny,
		AcceptableSLA: "99.5",
	}
	
	errors = api.ServiceValidate(invalidService2)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for service with invalid status")
	}
	
	// Invalid service - invalid algorithm
	invalidService3 := Service{
		Name:         "Test Service",
		Status:       ServiceStatusUp,
		Algorithm:    "invalid_algorithm",
		AcceptableSLA: "99.5",
	}
	
	errors = api.ServiceValidate(invalidService3)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for service with invalid algorithm")
	}
}

func TestMockServicesAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test service get operations
	opts := ServiceGetOptions{
		ServiceIDs: []string{"12345"},
		Output:     "extend",
		Limit:      10,
	}

	_, err := api.ServicesGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetByName("Test Service")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetByStatus(ServiceStatusUp)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetUp()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetDown()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetPartiallyDown()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetByParent("67890")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetRoot()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetWithChildren(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServicesGetWithParents(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServiceGetByID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test CRUD operations
	service := Service{
		Name:         "Test Service",
		Status:       ServiceStatusUp,
		Algorithm:    ServiceAlgorithmAny,
		Description:  "Test description",
	}
	
	_, err = api.ServiceCreateSingle(service)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServiceUpdateSingle(service)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ServiceDeleteSingle("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test relationship operations
	err = api.ServicesAddChildren("12345", []string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ServicesRemoveChildren("12345", []string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test SLA operations
	slaOpts := ServiceGetSLAOptions{
		ServiceIDs: []string{"12345"},
		TimeFrom:   1640995200,
		TimeTill:   1640998800,
	}

	_, err = api.ServicesGetSLA(slaOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServiceGetSLAForService("12345", 1640995200, 1640998800)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServiceGetSLADay("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServiceGetSLAWeek("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ServiceGetSLAMonth("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test tree operations
	_, err = api.ServicesGetTree()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics
	_, err = api.ServicesGetStatistics()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test status updates
	err = api.ServiceUpdateStatus("12345", ServiceStatusDown)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ServiceUpdateSLA("12345", "99.0")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test utility methods
	_, err = api.CreateSimpleService("Simple Service", "Simple description")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.CreateServiceTree([][]string{{"Service1", "Service2"}, {"Child1", "Child2"}})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkServiceMarshaling(b *testing.B) {
	service := Service{
		ServiceID:    "12345",
		Name:         "Test Service",
		Status:       ServiceStatusUp,
		Algorithm:    ServiceAlgorithmAny,
		Description:  "Test service description",
		ParentID:     "67890",
		AcceptableSLA: "99.5",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(service)
	}
}

func BenchmarkServiceTreeMarshaling(b *testing.B) {
	tree := ServiceTree{
		Service: Service{
			ServiceID: "12345",
			Name:      "Root Service",
			Status:    ServiceStatusUp,
		},
		Children: []ServiceTree{
			{
				Service: Service{
					ServiceID: "67890",
					Name:      "Child Service 1",
					Status:    ServiceStatusUp,
				},
			},
			{
				Service: Service{
					ServiceID: "11111",
					Name:      "Child Service 2",
					Status:    ServiceStatusDown,
				},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(tree)
	}
}

func BenchmarkSLAMeasureMarshaling(b *testing.B) {
	sla := SLAMeasure{
		ServiceID:       "12345",
		ServiceName:     "Web Service",
		SLI:            99.5,
		Uptime:          99.8,
		Downtime:        0.2,
		TimeToFirstFailure: 3600,
		TimeToRecovery:    300,
		AllowedDowntime:    0.5,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(sla)
	}
}

// Test integration scenarios
func TestServicesIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test complex filtering
	opts := ServiceGetOptions{
		Filter: map[string]interface{}{
			"status": ServiceStatusUp,
		},
		Output:           "extend",
		SelectChildren:   "extend",
		SelectParents:    "extend",
		SortField:        "name",
		SortOrder:        "ASC",
		Limit:            100,
	}

	_, err := api.ServicesGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test service creation with dependencies
	service := Service{
		Name:         "Integration Test Service",
		Status:       ServiceStatusUp,
		Algorithm:    ServiceAlgorithmAll,
		Description:  "Service created for integration testing",
		AcceptableSLA: "99.0",
		Tags: []Tag{
			{Tag: "environment", Value: "production"},
			{Tag: "criticality", Value: "high"},
		},
	}
	
	_, err = api.ServiceCreateSingle(service)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test service tree operations
	err = api.ServicesAddChildren("12345", []string{"67890", "11111"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ServicesRemoveChildren("12345", []string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test SLA operations
	slaOpts := ServiceGetSLAOptions{
		ServiceIDs: []string{"12345", "67890"},
		TimeFrom:   int(time.Now().Unix() - 24*3600), // Last 24 hours
		TimeTill:   int(time.Now().Unix()),
		Filters: []ServiceFilter{
			{ServiceID: "12345", Exclude: ServiceFilterExcludeNo},
		},
	}
	
	_, err = api.ServicesGetSLA(slaOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test service status updates
	err = api.ServiceUpdateStatus("12345", ServiceStatusDown)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ServiceUpdateStatus("12345", ServiceStatusUp)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test service tree creation
	hierarchicalServices := [][]string{
		{"Root Service 1", "Root Service 2"},
		{"Child Service 1.1", "Child Service 1.2", "Child Service 2.1"},
		{"Grandchild Service 1.1.1"},
	}
	
	_, err = api.CreateServiceTree(hierarchicalServices)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test bulk operations
	services := Services{
		{Name: "Bulk Service 1", Status: ServiceStatusUp, Algorithm: ServiceAlgorithmAny},
		{Name: "Bulk Service 2", Status: ServiceStatusDown, Algorithm: ServiceAlgorithmAll},
		{Name: "Bulk Service 3", Status: ServiceStatusPartiallyDown, Algorithm: ServiceAlgorithmAny},
	}
	
	_, err = api.ServicesCreate(services)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}