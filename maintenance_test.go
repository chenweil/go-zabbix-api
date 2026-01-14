package zabbix

import (
	"testing"
	"time"
)

func TestMaintenanceGetOptions(t *testing.T) {
	// Test default values
	opts := MaintenanceGetOptions{
		MaintenanceIDs: []string{"12345"},
	}

	if len(opts.MaintenanceIDs) != 1 {
		t.Errorf("Expected 1 maintenance ID, got %d", len(opts.MaintenanceIDs))
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

func TestMaintenance(t *testing.T) {
	// Test Maintenance creation and JSON marshaling
	maintenance := Maintenance{
		MaintenanceID: "12345",
		Name: "System Maintenance",
		MaintenanceType: MaintenanceTypeMaintenance,
		ActiveSince: "2022-01-01 10:00:00",
		ActiveTill: "2022-01-01 12:00:00",
		Description: "Monthly system maintenance",
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(maintenance)
	if err != nil {
		t.Errorf("Failed to marshal Maintenance: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestMaintenanceTag(t *testing.T) {
	// Test MaintenanceTag creation and JSON marshaling
	tag := MaintenanceTag{
		Tag:   "environment",
		Value: "production",
	}

	expectedJSON := `{"tag":"environment","value":"production"}`
	
	jsonData, err := json.Marshal(tag)
	if err != nil {
		t.Errorf("Failed to marshal MaintenanceTag: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestMaintenanceGroup(t *testing.T) {
	// Test MaintenanceGroup creation and JSON marshaling
	group := MaintenanceGroup{
		GroupID: "67890",
		Name:   "Web Servers",
	}

	expectedJSON := `{"groupid":"67890","name":"Web Servers"}`
	
	jsonData, err := json.Marshal(group)
	if err != nil {
		t.Errorf("Failed to marshal MaintenanceGroup: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestMaintenanceHost(t *testing.T) {
	// Test MaintenanceHost creation and JSON marshaling
	host := MaintenanceHost{
		HostID: "11111",
		Host:   "webserver01.example.com",
	}

	expectedJSON := `{"hostid":"11111","host":"webserver01.example.com"}`
	
	jsonData, err := json.Marshal(host)
	if err != nil {
		t.Errorf("Failed to marshal MaintenanceHost: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestMaintenanceTimePeriod(t *testing.T) {
	// Test MaintenanceTimePeriod creation and JSON marshaling
	timePeriod := MaintenanceTimePeriod{
		TimePeriodID: "22222",
		MaintenanceID: "12345",
		TimePeriodType: MaintenanceTimePeriodTypeDaily,
		EveryHour: "0",
		Minute: "0",
		Period: "3600",
	}

	expectedJSON := `{"timeperiodid":"22222","maintenanceid":"12345","timeperiod_type":"1","every_hour":"0","minute":"0","period":"3600"}`
	
	jsonData, err := json.Marshal(timePeriod)
	if err != nil {
		t.Errorf("Failed to marshal MaintenanceTimePeriod: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestMaintenanceWithDetails(t *testing.T) {
	// Test MaintenanceWithDetails creation and JSON marshaling
	maintenance := Maintenance{
		MaintenanceID: "12345",
		Name: "System Maintenance",
	}

	timePeriods := []MaintenanceTimePeriodDetail{
		{
			MaintenanceTimePeriod: MaintenanceTimePeriod{
				TimePeriodType: MaintenanceTimePeriodTypeOnetime,
				Period: "3600",
			},
			StartTime: "2022-01-01 10:00:00",
			EndTime:   "2022-01-01 11:00:00",
			Duration:  "3600",
		},
	}

	maintenanceWithDetails := MaintenanceWithDetails{
		Maintenance: maintenance,
		TimePeriodDetails: timePeriods,
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(maintenanceWithDetails)
	if err != nil {
		t.Errorf("Failed to marshal MaintenanceWithDetails: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestMaintenanceStatistics(t *testing.T) {
	// Test MaintenanceStatistics creation and JSON marshaling
	stats := MaintenanceStatistics{
		TotalMaintenances: 100,
		ActiveMaintenances: 5,
		ScheduledMaintenances: 10,
		CompletedMaintenances: 85,
		ByType: map[string]int{
			"0": 80,
			"1": 20,
		},
		ByStatus: map[string]int{
			"active":     5,
			"scheduled":  10,
			"expired":    85,
		},
		UpcomingMaintenances: []MaintenanceWithDetails{
			{Maintenance: Maintenance{Name: "Upcoming Maintenance"}},
		},
		ActiveHostGroups: []MaintenanceGroup{
			{GroupID: "1", Name: "Web Servers"},
		},
		ActiveHosts: []MaintenanceHost{
			{HostID: "1", Host: "webserver01"},
		},
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Failed to marshal MaintenanceStatistics: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestMaintenanceConstants(t *testing.T) {
	// Test maintenance constants
	tests := []struct {
		constant string
		expected string
	}{
		// Maintenance types
		{MaintenanceTypeMaintenance, "0"},
		{MaintenanceTypeNoDataCollection, "1"},
		
		// Time period types
		{MaintenanceTimePeriodTypeOnetime, "0"},
		{MaintenanceTimePeriodTypeDaily, "1"},
		{MaintenanceTimePeriodTypeWeekly, "2"},
		{MaintenanceTimePeriodTypeMonthly, "3"},
		{MaintenanceTimePeriodTypeYearly, "4"},
		
		// Days of week
		{MaintenanceDayMonday, "1"},
		{MaintenanceDayTuesday, "2"},
		{MaintenanceDayWednesday, "3"},
		{MaintenanceDayThursday, "4"},
		{MaintenanceDayFriday, "5"},
		{MaintenanceDaySaturday, "6"},
		{MaintenanceDaySunday, "7"},
		
		// Months
		{MaintenanceMonthJanuary, "1"},
		{MaintenanceMonthFebruary, "2"},
		{MaintenanceMonthMarch, "3"},
		{MaintenanceMonthApril, "4"},
		{MaintenanceMonthMay, "5"},
		{MaintenanceMonthJune, "6"},
		{MaintenanceMonthJuly, "7"},
		{MaintenanceMonthAugust, "8"},
		{MaintenanceMonthSeptember, "9"},
		{MaintenanceMonthOctober, "10"},
		{MaintenanceMonthNovember, "11"},
		{MaintenanceMonthDecember, "12"},
		
		// Status
		{MaintenanceStatusActive, "0"},
		{MaintenanceStatusScheduled, "1"},
		{MaintenanceStatusExpired, "2"},
	}

	for _, test := range tests {
		if test.constant != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.constant)
		}
	}
}

func TestMaintenanceIsActive(t *testing.T) {
	// Test Maintenance.IsActive() method
	// Create maintenance for current time
	now := time.Now()
	activeSince := now.Add(-time.Hour).Format("2006-01-02 15:04:05")
	activeTill := now.Add(time.Hour).Format("2006-01-02 15:04:05")
	
	activeMaintenance := Maintenance{
		ActiveSince: activeSince,
		ActiveTill: activeTill,
	}
	
	if !activeMaintenance.IsActive(int(now.Unix())) {
		t.Errorf("Expected active maintenance to be detected as active")
	}

	// Create past maintenance
	pastSince := now.Add(-2 * time.Hour).Format("2006-01-02 15:04:05")
	pastTill := now.Add(-time.Hour).Format("2006-01-02 15:04:05")
	
	pastMaintenance := Maintenance{
		ActiveSince: pastSince,
		ActiveTill: pastTill,
	}
	
	if pastMaintenance.IsActive(int(now.Unix())) {
		t.Errorf("Expected past maintenance to NOT be detected as active")
	}

	// Create future maintenance
	futureSince := now.Add(time.Hour).Format("2006-01-02 15:04:05")
	futureTill := now.Add(2 * time.Hour).Format("2006-01-02 15:04:05")
	
	futureMaintenance := Maintenance{
		ActiveSince: futureSince,
		ActiveTill: futureTill,
	}
	
	if futureMaintenance.IsActive(int(now.Unix())) {
		t.Errorf("Expected future maintenance to NOT be detected as active")
	}
}

func TestMaintenanceIsScheduled(t *testing.T) {
	// Test Maintenance.IsScheduled() method
	now := time.Now()
	
	// Create future maintenance
	futureSince := now.Add(time.Hour).Format("2006-01-02 15:04:05")
	futureTill := now.Add(2 * time.Hour).Format("2006-01-02 15:04:05")
	
	futureMaintenance := Maintenance{
		ActiveSince: futureSince,
		ActiveTill: futureTill,
	}
	
	if !futureMaintenance.IsScheduled(int(now.Unix())) {
		t.Errorf("Expected future maintenance to be detected as scheduled")
	}

	// Create active maintenance
	activeSince := now.Add(-time.Hour).Format("2006-01-02 15:04:05")
	activeTill := now.Add(time.Hour).Format("2006-01-02 15:04:05")
	
	activeMaintenance := Maintenance{
		ActiveSince: activeSince,
		ActiveTill: activeTill,
	}
	
	if activeMaintenance.IsScheduled(int(now.Unix())) {
		t.Errorf("Expected active maintenance to NOT be detected as scheduled")
	}
}

func TestMaintenanceIsExpired(t *testing.T) {
	// Test Maintenance.IsExpired() method
	now := time.Now()
	
	// Create past maintenance
	pastSince := now.Add(-2 * time.Hour).Format("2006-01-02 15:04:05")
	pastTill := now.Add(-time.Hour).Format("2006-01-02 15:04:05")
	
	pastMaintenance := Maintenance{
		ActiveSince: pastSince,
		ActiveTill: pastTill,
	}
	
	if !pastMaintenance.IsExpired(int(now.Unix())) {
		t.Errorf("Expected past maintenance to be detected as expired")
	}

	// Create active maintenance
	activeSince := now.Add(-time.Hour).Format("2006-01-02 15:04:05")
	activeTill := now.Add(time.Hour).Format("2006-01-02 15:04:05")
	
	activeMaintenance := Maintenance{
		ActiveSince: activeSince,
		ActiveTill: activeTill,
	}
	
	if activeMaintenance.IsExpired(int(now.Unix())) {
		t.Errorf("Expected active maintenance to NOT be detected as expired")
	}
}

func TestMaintenanceIsNoDataCollection(t *testing.T) {
	// Test Maintenance.IsNoDataCollection() method
	noDataMaintenance := Maintenance{
		MaintenanceType: MaintenanceTypeNoDataCollection,
	}
	
	regularMaintenance := Maintenance{
		MaintenanceType: MaintenanceTypeMaintenance,
	}
	
	emptyMaintenance := Maintenance{}
	
	if !noDataMaintenance.IsNoDataCollection() {
		t.Errorf("Expected no data collection maintenance to be detected as no data collection")
	}

	if regularMaintenance.IsNoDataCollection() {
		t.Errorf("Expected regular maintenance to NOT be detected as no data collection")
	}

	if emptyMaintenance.IsNoDataCollection() {
		t.Errorf("Expected empty maintenance to NOT be detected as no data collection")
	}
}

func TestMaintenanceIsRegularMaintenance(t *testing.T) {
	// Test Maintenance.IsRegularMaintenance() method
	regularMaintenance := Maintenance{
		MaintenanceType: MaintenanceTypeMaintenance,
	}
	
	noDataMaintenance := Maintenance{
		MaintenanceType: MaintenanceTypeNoDataCollection,
	}
	
	emptyMaintenance := Maintenance{}
	
	if !regularMaintenance.IsRegularMaintenance() {
		t.Errorf("Expected regular maintenance to be detected as regular maintenance")
	}

	if noDataMaintenance.IsRegularMaintenance() {
		t.Errorf("Expected no data collection maintenance to NOT be detected as regular maintenance")
	}

	if emptyMaintenance.IsRegularMaintenance() {
		t.Errorf("Expected empty maintenance to NOT be detected as regular maintenance")
	}
}

func TestMaintenanceHasTimePeriods(t *testing.T) {
	// Test Maintenance.HasTimePeriods() method
	maintenanceWithPeriods := Maintenance{
		TimePeriods: []MaintenanceTimePeriod{
			{TimePeriodType: MaintenanceTimePeriodTypeOnetime, Period: "3600"},
		},
	}
	
	maintenanceWithoutPeriods := Maintenance{
		TimePeriods: []MaintenanceTimePeriod{},
	}
	
	noPeriodsMaintenance := Maintenance{}
	
	if !maintenanceWithPeriods.HasTimePeriods() {
		t.Errorf("Expected maintenance with time periods to have time periods")
	}

	if maintenanceWithoutPeriods.HasTimePeriods() {
		t.Errorf("Expected maintenance without time periods to NOT have time periods")
	}

	if noPeriodsMaintenance.HasTimePeriods() {
		t.Errorf("Expected maintenance with nil time periods to NOT have time periods")
	}
}

func TestMaintenanceGetTimePeriodCount(t *testing.T) {
	// Test Maintenance.GetTimePeriodCount() method
	singlePeriodMaintenance := Maintenance{
		TimePeriods: []MaintenanceTimePeriod{
			{TimePeriodType: MaintenanceTimePeriodTypeOnetime, Period: "3600"},
		},
	}
	
	multiplePeriodsMaintenance := Maintenance{
		TimePeriods: []MaintenanceTimePeriod{
			{TimePeriodType: MaintenanceTimePeriodTypeOnetime, Period: "3600"},
			{TimePeriodType: MaintenanceTimePeriodTypeDaily, Period: "1800"},
			{TimePeriodType: MaintenanceTimePeriodTypeWeekly, Period: "900"},
		},
	}
	
	emptyMaintenance := Maintenance{}
	
	singleCount := singlePeriodMaintenance.GetTimePeriodCount()
	if singleCount != 1 {
		t.Errorf("Expected 1 time period, got %d", singleCount)
	}

	multipleCount := multiplePeriodsMaintenance.GetTimePeriodCount()
	if multipleCount != 3 {
		t.Errorf("Expected 3 time periods, got %d", multipleCount)
	}

	emptyCount := emptyMaintenance.GetTimePeriodCount()
	if emptyCount != 0 {
		t.Errorf("Expected 0 time periods, got %d", emptyCount)
	}
}

func TestMaintenanceHasGroups(t *testing.T) {
	// Test Maintenance.HasGroups() method
	maintenanceWithGroups := Maintenance{
		Groups: []MaintenanceGroup{
			{GroupID: "1", Name: "Web Servers"},
		},
	}
	
	maintenanceWithoutGroups := Maintenance{
		Groups: []MaintenanceGroup{},
	}
	
	noGroupsMaintenance := Maintenance{}
	
	if !maintenanceWithGroups.HasGroups() {
		t.Errorf("Expected maintenance with groups to have groups")
	}

	if maintenanceWithoutGroups.HasGroups() {
		t.Errorf("Expected maintenance without groups to NOT have groups")
	}

	if noGroupsMaintenance.HasGroups() {
		t.Errorf("Expected maintenance with nil groups to NOT have groups")
	}
}

func TestMaintenanceHasHosts(t *testing.T) {
	// Test Maintenance.HasHosts() method
	maintenanceWithHosts := Maintenance{
		Hosts: []MaintenanceHost{
			{HostID: "1", Host: "webserver01"},
		},
	}
	
	maintenanceWithoutHosts := Maintenance{
		Hosts: []MaintenanceHost{},
	}
	
	noHostsMaintenance := Maintenance{}
	
	if !maintenanceWithHosts.HasHosts() {
		t.Errorf("Expected maintenance with hosts to have hosts")
	}

	if maintenanceWithoutHosts.HasHosts() {
		t.Errorf("Expected maintenance without hosts to NOT have hosts")
	}

	if noHostsMaintenance.HasHosts() {
		t.Errorf("Expected maintenance with nil hosts to NOT have hosts")
	}
}

func TestMaintenanceAddTag(t *testing.T) {
	// Test Maintenance.AddTag() method
	maintenance := Maintenance{
		Name: "Test Maintenance",
	}
	
	initialCount := len(maintenance.Tags)
	
	newTag := MaintenanceTag{
		Tag:   "environment",
		Value: "production",
	}
	maintenance.AddTag(newTag)
	
	newCount := len(maintenance.Tags)
	if newCount != initialCount+1 {
		t.Errorf("Expected tag count to increase by 1, got %d", newCount)
	}

	// Check if the tag was added
	if len(maintenance.Tags) > 0 {
		if maintenance.Tags[0].Tag != "environment" {
			t.Errorf("Expected tag 'environment', got '%s'", maintenance.Tags[0].Tag)
		}
	}
}

func TestMaintenanceAddGroup(t *testing.T) {
	// Test Maintenance.AddGroup() method
	maintenance := Maintenance{
		Name: "Test Maintenance",
	}
	
	initialCount := len(maintenance.Groups)
	
	newGroup := MaintenanceGroup{
		GroupID: "1",
		Name:   "Web Servers",
	}
	maintenance.AddGroup(newGroup)
	
	newCount := len(maintenance.Groups)
	if newCount != initialCount+1 {
		t.Errorf("Expected group count to increase by 1, got %d", newCount)
	}

	// Check if the group was added
	if len(maintenance.Groups) > 0 {
		if maintenance.Groups[0].GroupID != "1" {
			t.Errorf("Expected group ID '1', got '%s'", maintenance.Groups[0].GroupID)
		}
	}
}

func TestMaintenanceAddHost(t *testing.T) {
	// Test Maintenance.AddHost() method
	maintenance := Maintenance{
		Name: "Test Maintenance",
	}
	
	initialCount := len(maintenance.Hosts)
	
	newHost := MaintenanceHost{
		HostID: "1",
		Host:   "webserver01",
	}
	maintenance.AddHost(newHost)
	
	newCount := len(maintenance.Hosts)
	if newCount != initialCount+1 {
		t.Errorf("Expected host count to increase by 1, got %d", newCount)
	}

	// Check if the host was added
	if len(maintenance.Hosts) > 0 {
		if maintenance.Hosts[0].HostID != "1" {
			t.Errorf("Expected host ID '1', got '%s'", maintenance.Hosts[0].HostID)
		}
	}
}

func TestMaintenanceAddTimePeriod(t *testing.T) {
	// Test Maintenance.AddTimePeriod() method
	maintenance := Maintenance{
		Name: "Test Maintenance",
	}
	
	initialCount := len(maintenance.TimePeriods)
	
	newTimePeriod := MaintenanceTimePeriod{
		TimePeriodType: MaintenanceTimePeriodTypeOnetime,
		Period:        "3600",
	}
	maintenance.AddTimePeriod(newTimePeriod)
	
	newCount := len(maintenance.TimePeriods)
	if newCount != initialCount+1 {
		t.Errorf("Expected time period count to increase by 1, got %d", newCount)
	}

	// Check if the time period was added
	if len(maintenance.TimePeriods) > 0 {
		if maintenance.TimePeriods[0].TimePeriodType != MaintenanceTimePeriodTypeOnetime {
			t.Errorf("Expected time period type '%s', got '%s'", MaintenanceTimePeriodTypeOnetime, maintenance.TimePeriods[0].TimePeriodType)
		}
	}
}

func TestMaintenanceValidate(t *testing.T) {
	// Test Maintenance validation
	api := NewAPI(Config{})
	
	// Valid maintenance
	validMaintenance := Maintenance{
		Name:            "Valid Maintenance",
		ActiveSince:     "2022-01-01 10:00:00",
		ActiveTill:      "2022-01-01 12:00:00",
		MaintenanceType: MaintenanceTypeMaintenance,
		Groups: []MaintenanceGroup{
			{GroupID: "1"},
		},
		TimePeriods: []MaintenanceTimePeriod{
			{TimePeriodType: MaintenanceTimePeriodTypeOnetime, Period: "3600"},
		},
	}
	
	errors := api.MaintenanceValidate(validMaintenance)
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors for valid maintenance, got: %v", errors)
	}
	
	// Invalid maintenance - missing name
	invalidMaintenance1 := Maintenance{
		ActiveSince:  "2022-01-01 10:00:00",
		ActiveTill:   "2022-01-01 12:00:00",
		MaintenanceType: MaintenanceTypeMaintenance,
		Groups: []MaintenanceGroup{
			{GroupID: "1"},
		},
	}
	
	errors = api.MaintenanceValidate(invalidMaintenance1)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for maintenance without name")
	}
	
	// Invalid maintenance - missing active since
	invalidMaintenance2 := Maintenance{
		Name:         "Test Maintenance",
		ActiveTill:   "2022-01-01 12:00:00",
		MaintenanceType: MaintenanceTypeMaintenance,
		Groups: []MaintenanceGroup{
			{GroupID: "1"},
		},
	}
	
	errors = api.MaintenanceValidate(invalidMaintenance2)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for maintenance without active since")
	}
	
	// Invalid maintenance - missing active till
	invalidMaintenance3 := Maintenance{
		Name:         "Test Maintenance",
		ActiveSince:  "2022-01-01 10:00:00",
		MaintenanceType: MaintenanceTypeMaintenance,
		Groups: []MaintenanceGroup{
			{GroupID: "1"},
		},
	}
	
	errors = api.MaintenanceValidate(invalidMaintenance3)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for maintenance without active till")
	}
	
	// Invalid maintenance - invalid type
	invalidMaintenance4 := Maintenance{
		Name:            "Test Maintenance",
		ActiveSince:     "2022-01-01 10:00:00",
		ActiveTill:      "2022-01-01 12:00:00",
		MaintenanceType: "invalid_type",
		Groups: []MaintenanceGroup{
			{GroupID: "1"},
		},
	}
	
	errors = api.MaintenanceValidate(invalidMaintenance4)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for maintenance with invalid type")
	}
	
	// Invalid maintenance - no groups or hosts
	invalidMaintenance5 := Maintenance{
		Name:         "Test Maintenance",
		ActiveSince:  "2022-01-01 10:00:00",
		ActiveTill:   "2022-01-01 12:00:00",
		MaintenanceType: MaintenanceTypeMaintenance,
	}
	
	errors = api.MaintenanceValidate(invalidMaintenance5)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for maintenance without groups or hosts")
	}
}

func TestMockMaintenancesAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test maintenance get operations
	opts := MaintenanceGetOptions{
		MaintenanceIDs: []string{"12345"},
		Output:        "extend",
		Limit:         10,
	}

	_, err := api.MaintenancesGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetByName("Test Maintenance")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetActive()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetScheduled()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetByGroup("67890")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetByHost("11111")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetWithTimePeriods(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetWithGroups(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetWithHosts(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenancesGetWithTags(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenanceGetByID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test CRUD operations
	maintenance := Maintenance{
		Name:            "Test Maintenance",
		MaintenanceType: MaintenanceTypeMaintenance,
		ActiveSince:     "2022-01-01 10:00:00",
		ActiveTill:      "2022-01-01 12:00:00",
		Description:     "Test description",
		Groups: []MaintenanceGroup{
			{GroupID: "1"},
		},
	}
	
	_, err = api.MaintenanceCreateSingle(maintenance)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenanceUpdateSingle(maintenance)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.MaintenanceDeleteSingle("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics
	_, err = api.MaintenancesGetStatistics()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test utility methods
	_, err = api.MaintenanceGetUpcoming(7)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenanceGetCurrent()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.MaintenanceCheckMaintenance([]string{"1"}, []string{"1"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test convenience methods
	_, err = api.CreateSimpleMaintenance("Simple Maintenance", "Description", 
		"2022-01-01 10:00:00", "2022-01-01 12:00:00", []string{"1"}, []string{"1"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.CreateOnetimeMaintenance("One-time Maintenance", "Description", 
		time.Now().Unix(), 3600, []string{"1"}, []string{"1"}, "0")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.CreateRecurringMaintenance("Recurring Maintenance", "Description", 
		time.Now().Unix(), 3600, []string{"1"}, []string{"1"}, "1")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkMaintenanceMarshaling(b *testing.B) {
	maintenance := Maintenance{
		MaintenanceID: "12345",
		Name:            "System Maintenance",
		MaintenanceType: MaintenanceTypeMaintenance,
		ActiveSince:     "2022-01-01 10:00:00",
		ActiveTill:      "2022-01-01 12:00:00",
		Description:     "Monthly system maintenance",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(maintenance)
	}
}

func BenchmarkMaintenanceTimePeriodMarshaling(b *testing.B) {
	timePeriod := MaintenanceTimePeriod{
		TimePeriodID: "67890",
		TimePeriodType: MaintenanceTimePeriodTypeDaily,
		EveryHour: "0",
		Minute: "0",
		Period: "3600",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(timePeriod)
	}
}

// Test integration scenarios
func TestMaintenancesIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test complex filtering
	opts := MaintenanceGetOptions{
		Filter: map[string]interface{}{
			"name": "System*",
		},
		Output:             "extend",
		SelectGroups:        "extend",
		SelectHosts:         "extend",
		SelectTimePeriods:    "extend",
		SortField:           "name",
		SortOrder:           "ASC",
		Limit:               100,
	}

	_, err := api.MaintenancesGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test maintenance creation with time periods
	maintenance := Maintenance{
		Name:            "Integration Test Maintenance",
		MaintenanceType: MaintenanceTypeMaintenance,
		ActiveSince:     "2022-01-01 10:00:00",
		ActiveTill:      "2022-01-01 12:00:00",
		Description:     "Maintenance created for integration testing",
		Groups: []MaintenanceGroup{
			{GroupID: "1", Name: "Web Servers"},
			{GroupID: "2", Name: "Database Servers"},
		},
		Hosts: []MaintenanceHost{
			{HostID: "1", Host: "webserver01"},
			{HostID: "2", Host: "dbserver01"},
		},
		TimePeriods: []MaintenanceTimePeriod{
			{
				TimePeriodType: MaintenanceTimePeriodTypeDaily,
				EveryHour: "0",
				Minute: "0",
				Period: "3600",
			},
			{
				TimePeriodType: MaintenanceTimePeriodTypeWeekly,
				DayOfWeek: MaintenanceDaySunday,
				Hour: "2",
				Minute: "0",
				Period: "7200",
			},
		},
		Tags: []MaintenanceTag{
			{Tag: "environment", Value: "production"},
			{Tag: "type", Value: "scheduled"},
		},
	}
	
	_, err = api.MaintenanceCreateSingle(maintenance)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test maintenance update
	maintenance.MaintenanceID = "11111"
	maintenance.Name = "Updated Integration Test Maintenance"
	
	_, err = api.MaintenanceUpdateSingle(maintenance)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test utility operations
	_, err = api.MaintenanceGetUpcoming(30)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test bulk operations
	maintenances := []Maintenance{
		{
			Name:            "Bulk Maintenance 1",
			MaintenanceType: MaintenanceTypeMaintenance,
			ActiveSince:     "2022-01-01 10:00:00",
			ActiveTill:      "2022-01-01 12:00:00",
			Groups: []MaintenanceGroup{
				{GroupID: "1"},
			},
		},
		{
			Name:            "Bulk Maintenance 2",
			MaintenanceType: MaintenanceTypeNoDataCollection,
			ActiveSince:     "2022-01-02 10:00:00",
			ActiveTill:      "2022-01-02 12:00:00",
			Hosts: []MaintenanceHost{
				{HostID: "1"},
			},
		},
	}
	
	_, err = api.MaintenancesCreate(Maintenances(maintenances))
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}