package zabbix

import (
	"testing"
)

func TestDiscoveryRuleGetOptions(t *testing.T) {
	// Test default values
	opts := DiscoveryRuleGetOptions{
		RuleIDs: []string{"12345"},
	}

	if len(opts.RuleIDs) != 1 {
		t.Errorf("Expected 1 rule ID, got %d", len(opts.RuleIDs))
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

func TestDiscoveryRule(t *testing.T) {
	// Test DiscoveryRule creation and JSON marshaling
	rule := DiscoveryRule{
		RuleID:      "12345",
		Name:        "Network Discovery",
		IPRange:     "192.168.1.0/24",
		Delay:       "3600",
		Status:      DiscoveryRuleStatusEnabled,
		ProxyHostID: "67890",
	}

	expectedJSON := `{"druleid":"12345","name":"Network Discovery","iprange":"192.168.1.0/24","delay":"3600","proxy_hostid":"67890","status":"0"}`
	
	jsonData, err := json.Marshal(rule)
	if err != nil {
		t.Errorf("Failed to marshal DiscoveryRule: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestDiscoveryCheck(t *testing.T) {
	// Test DiscoveryCheck creation and JSON marshaling
	check := DiscoveryCheck{
		CheckID:  "67890",
		RuleID:  "12345",
		Type:    DiscoveryCheckTypeSSH,
		Port:    "22",
		HostSource: DiscoveryHostSourceIP,
		NameSource: DiscoveryNameSourceDNS,
	}

	expectedJSON := `{"dcheckid":"67890","druleid":"12345","type":"0","port":"22","host_source":"0","name_source":"0"}`
	
	jsonData, err := json.Marshal(check)
	if err != nil {
		t.Errorf("Failed to marshal DiscoveryCheck: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestDiscoveredHost(t *testing.T) {
	// Test DiscoveredHost creation and JSON marshaling
	host := DiscoveredHost{
		HostID:     "11111",
		Host:       "192.168.1.100",
		LastCheck:  "1640995200",
		HostStatus: "1",
	}

	expectedJSON := `{"dhostid":"11111","host":"192.168.1.100","lastcheck":"1640995200","host":"1"}`
	
	jsonData, err := json.Marshal(host)
	if err != nil {
		t.Errorf("Failed to marshal DiscoveredHost: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestDiscoveredService(t *testing.T) {
	// Test DiscoveredService creation and JSON marshaling
	service := DiscoveredService{
		ServiceID:  "22222",
		HostID:     "11111",
		Service:    "SSH",
		Port:       "22",
		LastCheck:  "1640995200",
		ServiceStatus: "1",
		ServiceType: "9",
		ServiceKey: "ssh[22]",
	}

	expectedJSON := `{"dserviceid":"22222","dhostid":"11111","service":"SSH","port":"22","lastcheck":"1640995200","service":"1","service_type":"9","service_key":"ssh[22]"}`
	
	jsonData, err := json.Marshal(service)
	if err != nil {
		t.Errorf("Failed to marshal DiscoveredService: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestDiscoveryConstants(t *testing.T) {
	// Test discovery constants
	tests := []struct {
		constant string
		expected string
	}{
		// Discovery rule status
		{DiscoveryRuleStatusEnabled, "0"},
		{DiscoveryRuleStatusDisabled, "1"},
		
		// Discovery check types
		{DiscoveryCheckTypeSSH, "0"},
		{DiscoveryCheckTypeLDAP, "1"},
		{DiscoveryCheckTypeSMTP, "2"},
		{DiscoveryCheckTypeFTP, "3"},
		{DiscoveryCheckTypeHTTP, "4"},
		{DiscoveryCheckTypeHTTPS, "5"},
		{DiscoveryCheckTypePOP, "6"},
		{DiscoveryCheckTypeNNTP, "7"},
		{DiscoveryCheckTypeIMAP, "8"},
		{DiscoveryCheckTypeTCP, "9"},
		{DiscoveryCheckTypeICMP, "10"},
		{DiscoveryCheckTypeSNMPv1, "11"},
		{DiscoveryCheckTypeSNMPv2c, "12"},
		{DiscoveryCheckTypeSNMPv3, "13"},
		{DiscoveryCheckTypeTelnet, "14"},
		{DiscoveryCheckTypeCustom, "19"},
		
		// Host source
		{DiscoveryHostSourceIP, "0"},
		{DiscoveryHostSourceDNS, "1"},
		
		// Name source
		{DiscoveryNameSourceDNS, "0"},
		{DiscoveryNameSourceIP, "1"},
		
		// SNMP security levels
		{SNMPv3SecurityLevelNoAuthNoPriv, "0"},
		{SNMPv3SecurityLevelAuthNoPriv, "1"},
		{SNMPv3SecurityLevelAuthPriv, "2"},
		
		// SNMP authentication protocols
		{SNMPv3AuthProtocolMD5, "0"},
		{SNMPv3AuthProtocolSHA, "1"},
		
		// SNMP privacy protocols
		{SNMPv3PrivProtocolDES, "0"},
		{SNMPv3PrivProtocolAES, "1"},
	}

	for _, test := range tests {
		if test.constant != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.constant)
		}
	}
}

func TestDiscoveryRuleIsEnabled(t *testing.T) {
	// Test DiscoveryRule.IsEnabled() method
	enabledRule := DiscoveryRule{Status: DiscoveryRuleStatusEnabled}
	disabledRule := DiscoveryRule{Status: DiscoveryRuleStatusDisabled}
	noStatusRule := DiscoveryRule{Status: ""}

	if !enabledRule.IsEnabled() {
		t.Errorf("Expected enabled rule to be detected as enabled")
	}

	if disabledRule.IsEnabled() {
		t.Errorf("Expected disabled rule to NOT be detected as enabled")
	}

	if noStatusRule.IsEnabled() {
		t.Errorf("Expected rule with no status to NOT be detected as enabled")
	}
}

func TestDiscoveryRuleIsDisabled(t *testing.T) {
	// Test DiscoveryRule.IsDisabled() method
	disabledRule := DiscoveryRule{Status: DiscoveryRuleStatusDisabled}
	enabledRule := DiscoveryRule{Status: DiscoveryRuleStatusEnabled}
	noStatusRule := DiscoveryRule{Status: ""}

	if !disabledRule.IsDisabled() {
		t.Errorf("Expected disabled rule to be detected as disabled")
	}

	if enabledRule.IsDisabled() {
		t.Errorf("Expected enabled rule to NOT be detected as disabled")
	}

	if noStatusRule.IsDisabled() {
		t.Errorf("Expected rule with no status to NOT be detected as disabled")
	}
}

func TestDiscoveryRuleHasChecks(t *testing.T) {
	// Test DiscoveryRule.HasChecks() method
	ruleWithChecks := DiscoveryRule{
		Checks: []DiscoveryCheck{
			{Type: DiscoveryCheckTypeSSH, Port: "22"},
		},
	}
	ruleWithoutChecks := DiscoveryRule{
		Checks: []DiscoveryCheck{},
	}
	ruleNoChecks := DiscoveryRule{}

	if !ruleWithChecks.HasChecks() {
		t.Errorf("Expected rule with checks to have checks")
	}

	if ruleWithoutChecks.HasChecks() {
		t.Errorf("Expected rule without checks to NOT have checks")
	}

	if ruleNoChecks.HasChecks() {
		t.Errorf("Expected rule with nil checks to NOT have checks")
	}
}

func TestDiscoveryRuleGetCheckCount(t *testing.T) {
	// Test DiscoveryRule.GetCheckCount() method
	ruleWithChecks := DiscoveryRule{
		Checks: []DiscoveryCheck{
			{Type: DiscoveryCheckTypeSSH, Port: "22"},
			{Type: DiscoveryCheckTypeHTTP, Port: "80"},
			{Type: DiscoveryCheckTypeHTTPS, Port: "443"},
		},
	}
	ruleWithoutChecks := DiscoveryRule{
		Checks: []DiscoveryCheck{},
	}
	ruleNoChecks := DiscoveryRule{}

	count := ruleWithChecks.GetCheckCount()
	if count != 3 {
		t.Errorf("Expected 3 checks, got %d", count)
	}

	count = ruleWithoutChecks.GetCheckCount()
	if count != 0 {
		t.Errorf("Expected 0 checks, got %d", count)
	}

	count = ruleNoChecks.GetCheckCount()
	if count != 0 {
		t.Errorf("Expected 0 checks, got %d", count)
	}
}

func TestDiscoveredHostIsActive(t *testing.T) {
	// Test DiscoveredHost.IsActive() method
	activeHost := DiscoveredHost{HostStatus: "1"}
	inactiveHost := DiscoveredHost{HostStatus: "0"}
	noStatusHost := DiscoveredHost{HostStatus: ""}

	if !activeHost.IsActive() {
		t.Errorf("Expected active host to be detected as active")
	}

	if inactiveHost.IsActive() {
		t.Errorf("Expected inactive host to NOT be detected as active")
	}

	if noStatusHost.IsActive() {
		t.Errorf("Expected host with no status to NOT be detected as active")
	}
}

func TestDiscoveredServiceIsActive(t *testing.T) {
	// Test DiscoveredService.IsActive() method
	activeService := DiscoveredService{ServiceStatus: "1"}
	inactiveService := DiscoveredService{ServiceStatus: "0"}
	noStatusService := DiscoveredService{ServiceStatus: ""}

	if !activeService.IsActive() {
		t.Errorf("Expected active service to be detected as active")
	}

	if inactiveService.IsActive() {
		t.Errorf("Expected inactive service to NOT be detected as active")
	}

	if noStatusService.IsActive() {
		t.Errorf("Expected service with no status to NOT be detected as active")
	}
}

func TestDiscoveryValidate(t *testing.T) {
	// Test Discovery validation
	api := NewAPI(Config{})
	
	// Valid discovery rule
	validRule := DiscoveryRule{
		Name:        "Valid Discovery Rule",
		IPRange:     "192.168.1.0/24",
		Delay:       "3600",
		Status:      DiscoveryRuleStatusEnabled,
		Checks: []DiscoveryCheck{
			{
				Type: DiscoveryCheckTypeSSH,
				Port: "22",
			},
		},
	}
	
	errors := api.DiscoveryValidate(validRule)
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors for valid rule, got: %v", errors)
	}
	
	// Invalid rule - missing name
	invalidRule1 := DiscoveryRule{
		IPRange: "192.168.1.0/24",
		Delay:   "3600",
		Status:  DiscoveryRuleStatusEnabled,
	}
	
	errors = api.DiscoveryValidate(invalidRule1)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for rule without name")
	}
	
	// Invalid rule - missing IP range
	invalidRule2 := DiscoveryRule{
		Name:   "Test Rule",
		Delay:  "3600",
		Status: DiscoveryRuleStatusEnabled,
	}
	
	errors = api.DiscoveryValidate(invalidRule2)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for rule without IP range")
	}
	
	// Invalid rule - missing delay
	invalidRule3 := DiscoveryRule{
		Name:     "Test Rule",
		IPRange:  "192.168.1.0/24",
		Status:   DiscoveryRuleStatusEnabled,
	}
	
	errors = api.DiscoveryValidate(invalidRule3)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for rule without delay")
	}
	
	// Invalid rule - invalid status
	invalidRule4 := DiscoveryRule{
		Name:     "Test Rule",
		IPRange:  "192.168.1.0/24",
		Delay:    "3600",
		Status:   "invalid_status",
	}
	
	errors = api.DiscoveryValidate(invalidRule4)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for rule with invalid status")
	}
	
	// Invalid check - missing type
	invalidRule5 := DiscoveryRule{
		Name:     "Test Rule",
		IPRange:  "192.168.1.0/24",
		Delay:    "3600",
		Status:   DiscoveryRuleStatusEnabled,
		Checks: []DiscoveryCheck{
			{
				Port: "22",
				// Missing Type
			},
		},
	}
	
	errors = api.DiscoveryValidate(invalidRule5)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for check without type")
	}
	
	// Invalid check - missing port
	invalidRule6 := DiscoveryRule{
		Name:     "Test Rule",
		IPRange:  "192.168.1.0/24",
		Delay:    "3600",
		Status:   DiscoveryRuleStatusEnabled,
		Checks: []DiscoveryCheck{
			{
				Type: DiscoveryCheckTypeSSH,
				// Missing Port
			},
		},
	}
	
	errors = api.DiscoveryValidate(invalidRule6)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for check without port")
	}
}

func TestMockDiscoveryAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test discovery rules
	opts := DiscoveryRuleGetOptions{
		RuleIDs: []string{"12345"},
		Output: "extend",
		Limit:  10,
	}

	_, err := api.DiscoveryRulesGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRulesGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRulesGetByName("Test Discovery")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRulesGetEnabled()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRulesGetDisabled()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRulesGetWithChecks(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRuleGetByID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test discovery checks
	checkOpts := DiscoveryCheckGetOptions{
		RuleIDs: []string{"12345"},
		Output: "extend",
	}

	_, err = api.DiscoveryChecksGet(checkOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryChecksGetByRule("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryChecksGetByID([]string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryChecksGetByType(DiscoveryCheckTypeSSH)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test discovered hosts
	hostOpts := DiscoveredHostGetOptions{
		RuleIDs: []string{"12345"},
		Output: "extend",
	}

	_, err = api.DiscoveredHostsGet(hostOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredHostsGetByRule("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredHostsGetByID([]string{"11111"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredHostsGetActive()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredHostsGetWithServices(hostOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test discovered services
	serviceOpts := DiscoveredServiceGetOptions{
		HostIDs: []string{"11111"},
		Output: "extend",
	}

	_, err = api.DiscoveredServicesGet(serviceOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredServicesGetByHost("11111")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredServicesGetByRule("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredServicesGetByType("9")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveredServicesGetByPort("22")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test CRUD operations
	rule := DiscoveryRule{
		Name:        "Test Discovery Rule",
		IPRange:     "192.168.1.0/24",
		Delay:       "3600",
		Status:      DiscoveryRuleStatusEnabled,
		Checks: []DiscoveryCheck{
			{Type: DiscoveryCheckTypeSSH, Port: "22"},
		},
	}
	
	_, err = api.DiscoveryRuleCreateSingle(rule)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRuleUpdateSingle(rule)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.DiscoveryRuleDeleteSingle("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test enable/disable operations
	_, err = api.DiscoveryRulesEnable([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRulesDisable([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics
	_, err = api.DiscoveryGetStatistics()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkDiscoveryRuleMarshaling(b *testing.B) {
	rule := DiscoveryRule{
		RuleID:      "12345",
		Name:        "Network Discovery Rule",
		IPRange:     "192.168.1.0/24",
		Delay:       "3600",
		Status:      DiscoveryRuleStatusEnabled,
		ProxyHostID: "67890",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(rule)
	}
}

func BenchmarkDiscoveryCheckMarshaling(b *testing.B) {
	check := DiscoveryCheck{
		CheckID:  "67890",
		RuleID:  "12345",
		Type:    DiscoveryCheckTypeSSH,
		Port:    "22",
		HostSource: DiscoveryHostSourceIP,
		NameSource: DiscoveryNameSourceDNS,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(check)
	}
}

// Test integration scenarios
func TestDiscoveryIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test discovery rule creation with complex checks
	rule := DiscoveryRule{
		Name:        "Complex Network Discovery",
		IPRange:     "10.0.0.0/16",
		Delay:       "7200",
		Status:      DiscoveryRuleStatusEnabled,
		ProxyHostID: "12345",
		Checks: []DiscoveryCheck{
			{
				Type:      DiscoveryCheckTypeSSH,
				Port:      "22",
				HostSource: DiscoveryHostSourceIP,
				NameSource: DiscoveryNameSourceDNS,
			},
			{
				Type:      DiscoveryCheckTypeHTTP,
				Port:      "80",
				HostSource: DiscoveryHostSourceIP,
				NameSource: DiscoveryNameSourceDNS,
			},
			{
				Type:      DiscoveryCheckTypeHTTPS,
				Port:      "443",
				HostSource: DiscoveryHostSourceIP,
				NameSource: DiscoveryNameSourceDNS,
			},
			{
				Type:      DiscoveryCheckTypeSNMPv2c,
				Port:      "161",
				HostSource: DiscoveryHostSourceIP,
				NameSource: DiscoveryNameSourceIP,
				SnmpCommunity: "public",
			},
		},
	}
	
	_, err := api.DiscoveryRuleCreateSingle(rule)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test discovery rule update
	rule.RuleID = "11111"
	rule.Name = "Updated Discovery Rule"
	rule.IPRange = "172.16.0.0/16"
	
	_, err = api.DiscoveryRuleUpdateSingle(rule)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test discovery rule filtering
	opts := DiscoveryRuleGetOptions{
		Filter: map[string]interface{}{
			"status": DiscoveryRuleStatusEnabled,
		},
		Output:       "extend",
		SelectChecks: "extend",
		Limit:        100,
	}
	
	_, err = api.DiscoveryRulesGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test discovered hosts with services
	hostOpts := DiscoveredHostGetOptions{
		Output:        "extend",
		SelectServices: "extend",
		SortField:     "host",
		SortOrder:     "ASC",
	}
	
	_, err = api.DiscoveredHostsGetWithServices(hostOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test discovered services filtering
	serviceOpts := DiscoveredServiceGetOptions{
		Filter: map[string]interface{}{
			"service_type": "9", // TCP
		},
		Output:    "extend",
		Limit:     50,
		SortField: "port",
		SortOrder: "ASC",
	}
	
	_, err = api.DiscoveredServicesGet(serviceOpts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test bulk operations
	rules := DiscoveryRules{
		{Name: "Bulk Rule 1", IPRange: "192.168.1.0/24", Delay: "3600", Status: DiscoveryRuleStatusEnabled},
		{Name: "Bulk Rule 2", IPRange: "192.168.2.0/24", Delay: "7200", Status: DiscoveryRuleStatusEnabled},
		{Name: "Bulk Rule 3", IPRange: "192.168.3.0/24", Delay: "10800", Status: DiscoveryRuleStatusDisabled},
	}
	
	_, err = api.DiscoveryRulesCreate(rules)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test enable/disable bulk operations
	_, err = api.DiscoveryRulesEnable([]string{"11111", "22222", "33333"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DiscoveryRulesDisable([]string{"11111", "22222", "33333"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}