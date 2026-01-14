package zabbix

import (
	"testing"
)

func TestDashboardGetOptions(t *testing.T) {
	// Test default values
	opts := DashboardGetOptions{
		DashboardIDs: []string{"12345"},
	}

	if len(opts.DashboardIDs) != 1 {
		t.Errorf("Expected 1 dashboard ID, got %d", len(opts.DashboardIDs))
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

func TestDashboard(t *testing.T) {
	// Test Dashboard creation and JSON marshaling
	dashboard := Dashboard{
		DashboardID: "12345",
		Name:        "System Monitoring",
		UserID:      "67890",
		Private:     DashboardSharingPrivate,
		Description: "Main system monitoring dashboard",
	}

	expectedJSON := `{"dashboardid":"12345","name":"System Monitoring","userid":"67890","private":"0","description":"Main system monitoring dashboard"}`
	
	jsonData, err := json.Marshal(dashboard)
	if err != nil {
		t.Errorf("Failed to marshal Dashboard: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestWidget(t *testing.T) {
	// Test Widget creation and JSON marshaling
	widget := Widget{
		WidgetID:    "67890",
		DashboardID: "12345",
		Type:        WidgetTypeGraph,
		Name:        "CPU Usage",
		X:           "0",
		Y:           "0",
		Width:       "12",
		Height:      "8",
		Fields:      map[string]interface{}{"name": "value"},
	}

	expectedJSON := `{"widgetid":"67890","dashboardid":"12345","type":"graph","name":"CPU Usage","x":"0","y":"0","width":"12","height":"8","fields":{"name":"value"}}`
	
	jsonData, err := json.Marshal(widget)
	if err != nil {
		t.Errorf("Failed to marshal Widget: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestWidgetView(t *testing.T) {
	// Test WidgetView creation and JSON marshaling
	widgetView := WidgetView{
		RefreshInterval: DashboardRefresh1m,
		ShowHeader:     "1",
	}

	expectedJSON := `{"refresh_interval":"60","show_header":"1"}`
	
	jsonData, err := json.Marshal(widgetView)
	if err != nil {
		t.Errorf("Failed to marshal WidgetView: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestWidgetHeader(t *testing.T) {
	// Test WidgetHeader creation and JSON marshaling
	widgetHeader := WidgetHeader{
		Show:       WidgetHeaderShow,
		Position:   WidgetPositionLeft,
		Draggable:  WidgetHeaderDraggable,
	}

	expectedJSON := `{"show":"1","position":"left","draggable":"1"}`
	
	jsonData, err := json.Marshal(widgetHeader)
	if err != nil {
		t.Errorf("Failed to marshal WidgetHeader: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestWidgetField(t *testing.T) {
	// Test WidgetField creation and JSON marshaling
	widgetField := WidgetField{
		Name:  "hostid",
		Value: "12345",
	}

	expectedJSON := `{"name":"hostid","value":"12345"}`
	
	jsonData, err := json.Marshal(widgetField)
	if err != nil {
		t.Errorf("Failed to marshal WidgetField: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestDashboardShare(t *testing.T) {
	// Test DashboardShare creation and JSON marshaling
	dashboardShare := DashboardShare{
		Private: DashboardSharingPublic,
		Users: []User{
			{UserID: "123", Username: "user1"},
		},
		UserGroups: []UserGroup{
			{UsrGroupID: "456", Name: "admin-group"},
		},
	}

	expectedJSON := `{"users":[{"userid":"123","username":"user1"}],"user_groups":[{"usrgrpid":"456","name":"admin-group"}],"private":"1"}`
	
	jsonData, err := json.Marshal(dashboardShare)
	if err != nil {
		t.Errorf("Failed to marshal DashboardShare: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestDashboardConstants(t *testing.T) {
	// Test dashboard constants
	tests := []struct {
		constant string
		expected string
	}{
		// Dashboard sharing modes
		{DashboardSharingPrivate, "0"},
		{DashboardSharingPublic, "1"},
		
		// Refresh intervals
		{DashboardRefresh30s, "30"},
		{DashboardRefresh1m, "60"},
		{DashboardRefresh2m, "120"},
		{DashboardRefresh5m, "300"},
		{DashboardRefresh10m, "600"},
		{DashboardRefresh30m, "1800"},
		{DashboardRefresh1h, "3600"},
		
		// Widget types
		{WidgetTypeClock, "clock"},
		{WidgetTypeDataOverview, "data_overview"},
		{WidgetTypeDiscovery, "discovery"},
		{WidgetTypeGraph, "graph"},
		{WidgetTypeGraphPrototype, "graph_prototype"},
		{WidgetTypeHostIssues, "host_issues"},
		{WidgetTypeMap, "map"},
		{WidgetTypePlainText, "plain_text"},
		{WidgetTypeSlideShow, "slideshow"},
		{WidgetTypeSvgGraph, "svg_graph"},
		{WidgetTypeSystemInfo, "system_info"},
		{WidgetTypeTopHosts, "top_hosts"},
		{WidgetTypeTriggerInfo, "trigger_info"},
		{WidgetTypeTriggerOverview, "trigger_overview"},
		{WidgetTypeUrl, "url"},
		{WidgetTypeWebMonitoring, "web"},
		
		// Zabbix 7.0+ Widget types
		{WidgetTypeItemHistory, "item_history"},
		
		// Widget positioning
		{WidgetPositionLeft, "left"},
		{WidgetPositionRight, "right"},
		
		// Widget header options
		{WidgetHeaderShow, "1"},
		{WidgetHeaderHide, "0"},
		{WidgetHeaderDraggable, "1"},
		{WidgetHeaderFixed, "0"},
	}

	for _, test := range tests {
		if test.constant != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.constant)
		}
	}
}

func TestDashboardIsPublic(t *testing.T) {
	// Test Dashboard.IsPublic() method
	publicDashboard := Dashboard{Private: DashboardSharingPublic}
	privateDashboard := Dashboard{Private: DashboardSharingPrivate}
	noSharingDashboard := Dashboard{Private: ""}

	if !publicDashboard.IsPublic() {
		t.Errorf("Expected public dashboard to be detected as public")
	}

	if privateDashboard.IsPublic() {
		t.Errorf("Expected private dashboard to NOT be detected as public")
	}

	if noSharingDashboard.IsPublic() {
		t.Errorf("Expected dashboard with no sharing setting to NOT be detected as public")
	}
}

func TestDashboardIsPrivate(t *testing.T) {
	// Test Dashboard.IsPrivate() method
	privateDashboard := Dashboard{Private: DashboardSharingPrivate}
	publicDashboard := Dashboard{Private: DashboardSharingPublic}
	noSharingDashboard := Dashboard{Private: ""}

	if !privateDashboard.IsPrivate() {
		t.Errorf("Expected private dashboard to be detected as private")
	}

	if publicDashboard.IsPrivate() {
		t.Errorf("Expected public dashboard to NOT be detected as private")
	}

	if noSharingDashboard.IsPrivate() {
		t.Errorf("Expected dashboard with no sharing setting to NOT be detected as private")
	}
}

func TestDashboardHasWidgets(t *testing.T) {
	// Test Dashboard.HasWidgets() method
	dashboardWithWidgets := Dashboard{
		Widgets: []Widget{
			{Type: WidgetTypeGraph, Name: "Widget 1"},
		},
	}
	dashboardWithoutWidgets := Dashboard{
		Widgets: []Widget{},
	}
	dashboardNoWidgets := Dashboard{}

	if !dashboardWithWidgets.HasWidgets() {
		t.Errorf("Expected dashboard with widgets to have widgets")
	}

	if dashboardWithoutWidgets.HasWidgets() {
		t.Errorf("Expected dashboard without widgets to NOT have widgets")
	}

	if dashboardNoWidgets.HasWidgets() {
		t.Errorf("Expected dashboard with nil widgets to NOT have widgets")
	}
}

func TestDashboardGetWidgetCount(t *testing.T) {
	// Test Dashboard.GetWidgetCount() method
	dashboardWithWidgets := Dashboard{
		Widgets: []Widget{
			{Type: WidgetTypeGraph, Name: "Widget 1"},
			{Type: WidgetTypeClock, Name: "Widget 2"},
			{Type: WidgetTypeMap, Name: "Widget 3"},
		},
	}
	dashboardWithoutWidgets := Dashboard{
		Widgets: []Widget{},
	}
	dashboardNoWidgets := Dashboard{}

	count := dashboardWithWidgets.GetWidgetCount()
	if count != 3 {
		t.Errorf("Expected 3 widgets, got %d", count)
	}

	count = dashboardWithoutWidgets.GetWidgetCount()
	if count != 0 {
		t.Errorf("Expected 0 widgets, got %d", count)
	}

	count = dashboardNoWidgets.GetWidgetCount()
	if count != 0 {
		t.Errorf("Expected 0 widgets, got %d", count)
	}
}

func TestDashboardAddWidget(t *testing.T) {
	// Test Dashboard.AddWidget() method
	dashboard := Dashboard{
		Name: "Test Dashboard",
		Widgets: []Widget{
			{Type: WidgetTypeGraph, Name: "Widget 1"},
		},
	}

	initialCount := dashboard.GetWidgetCount()
	
	newWidget := Widget{Type: WidgetTypeClock, Name: "Widget 2"}
	dashboard.AddWidget(newWidget)

	newCount := dashboard.GetWidgetCount()
	if newCount != initialCount+1 {
		t.Errorf("Expected widget count to increase by 1, got %d", newCount)
	}

	// Check if the widget was added
	if len(dashboard.Widgets) != 2 {
		t.Errorf("Expected 2 widgets in array, got %d", len(dashboard.Widgets))
	}
}

func TestDashboardRemoveWidget(t *testing.T) {
	// Test Dashboard.RemoveWidget() method
	widget1 := Widget{WidgetID: "1", Type: WidgetTypeGraph, Name: "Widget 1"}
	widget2 := Widget{WidgetID: "2", Type: WidgetTypeClock, Name: "Widget 2"}
	widget3 := Widget{WidgetID: "3", Type: WidgetTypeMap, Name: "Widget 3"}

	dashboard := Dashboard{
		Name: "Test Dashboard",
		Widgets: []Widget{widget1, widget2, widget3},
	}

	initialCount := dashboard.GetWidgetCount()
	
	// Remove widget 2
	dashboard.RemoveWidget("2")

	newCount := dashboard.GetWidgetCount()
	if newCount != initialCount-1 {
		t.Errorf("Expected widget count to decrease by 1, got %d", newCount)
	}

	// Check if widget 2 was removed and others remain
	if len(dashboard.Widgets) != 2 {
		t.Errorf("Expected 2 widgets in array, got %d", len(dashboard.Widgets))
	}

	// Verify the remaining widgets are widget1 and widget3
	if dashboard.Widgets[0].WidgetID != "1" || dashboard.Widgets[1].WidgetID != "3" {
		t.Errorf("Unexpected widgets remaining: %v", dashboard.Widgets)
	}
}

func TestDashboardUpdateWidget(t *testing.T) {
	// Test Dashboard.UpdateWidget() method
	widget1 := Widget{WidgetID: "1", Type: WidgetTypeGraph, Name: "Widget 1"}
	widget2 := Widget{WidgetID: "2", Type: WidgetTypeClock, Name: "Widget 2"}

	dashboard := Dashboard{
		Name: "Test Dashboard",
		Widgets: []Widget{widget1, widget2},
	}

	// Update widget 2
	updatedWidget := Widget{WidgetID: "2", Type: WidgetTypeMap, Name: "Updated Widget 2"}
	dashboard.UpdateWidget(updatedWidget)

	// Check if widget was updated
	if dashboard.Widgets[1].Type != WidgetTypeMap {
		t.Errorf("Expected widget type to be updated to map, got %s", dashboard.Widgets[1].Type)
	}

	if dashboard.Widgets[1].Name != "Updated Widget 2" {
		t.Errorf("Expected widget name to be updated, got %s", dashboard.Widgets[1].Name)
	}

	// Verify widget 1 was not affected
	if dashboard.Widgets[0].Type != WidgetTypeGraph {
		t.Errorf("Expected widget 1 type to remain unchanged")
	}
}

func TestParseInt(t *testing.T) {
	// Test parseInt function
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"0", 0},
		{"-456", -456},
		{"", 0},
		{"invalid", 0},
	}

	for _, test := range tests {
		result := parseInt(test.input)
		if result != test.expected {
			t.Errorf("Expected parseInt('%s') = %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestDashboardValidate(t *testing.T) {
	// Test Dashboard validation
	api := NewAPI(Config{})
	
	// Valid dashboard
	validDashboard := Dashboard{
		Name: "Valid Dashboard",
		Widgets: []Widget{
			{
				Type: WidgetTypeGraph,
				Name: "CPU Graph",
				X:    "0",
				Y:    "0",
				Width: "12",
				Height: "8",
			},
		},
	}
	
	errors := api.DashboardValidate(validDashboard)
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors for valid dashboard, got: %v", errors)
	}
	
	// Invalid dashboard - missing name
	invalidDashboard1 := Dashboard{
		Widgets: []Widget{
			{Type: WidgetTypeGraph, Name: "Widget"},
		},
	}
	
	errors = api.DashboardValidate(invalidDashboard1)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for dashboard without name")
	}
	
	// Invalid widget - missing type
	invalidDashboard2 := Dashboard{
		Name: "Test Dashboard",
		Widgets: []Widget{
			{
				Name: "Widget without type",
				X:    "0",
				Y:    "0",
				Width: "12",
				Height: "8",
			},
		},
	}
	
	errors = api.DashboardValidate(invalidDashboard2)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for widget without type")
	}
	
	// Invalid widget - missing name
	invalidDashboard3 := Dashboard{
		Name: "Test Dashboard",
		Widgets: []Widget{
			{
				Type: WidgetTypeGraph,
				X:    "0",
				Y:    "0",
				Width: "12",
				Height: "8",
			},
		},
	}
	
	errors = api.DashboardValidate(invalidDashboard3)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for widget without name")
	}
	
	// Invalid widget - invalid coordinates
	invalidDashboard4 := Dashboard{
		Name: "Test Dashboard",
		Widgets: []Widget{
			{
				Type: WidgetTypeGraph,
				Name: "Widget",
				X:    "50", // Invalid: too large
				Y:    "0",
				Width: "12",
				Height: "8",
			},
		},
	}
	
	errors = api.DashboardValidate(invalidDashboard4)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for widget with invalid coordinates")
	}
}

func TestMockDashboardsAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test that methods exist and return appropriate types
	opts := DashboardGetOptions{
		DashboardIDs: []string{"12345"},
		Output:      "extend",
		Limit:       10,
	}

	// These calls will fail without a real Zabbix server, but we can verify the method signatures
	_, err := api.DashboardsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardsGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardsGetByName("Test Dashboard")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardsGetByUser("67890")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardsGetPrivate()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardsGetPublic()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardGetByID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardsGetWithWidgets(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test CRUD operations
	dashboard := Dashboard{
		Name:        "Test Dashboard",
		Description: "Test description",
		Private:     DashboardSharingPrivate,
		AutoCommit:  "1",
	}
	
	_, err = api.DashboardCreateSingle(dashboard)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.DashboardUpdateSingle(dashboard)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.DashboardDeleteSingle("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test sharing operations
	err = api.DashboardMakePublic("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.DashboardMakePrivate("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	sharing := DashboardShare{
		Private: DashboardSharingPublic,
		Users:   []User{{UserID: "123", Username: "user1"}},
	}
	
	err = api.DashboardShare("12345", sharing)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics
	_, err = api.DashboardsGetStatistics()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test utility methods
	_, err = api.CreateSimpleDashboard("Simple Dashboard", "Description", "user123")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.CreateSystemOverviewDashboard("user123")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.GetUserDashboards("user123")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkDashboardMarshaling(b *testing.B) {
	dashboard := Dashboard{
		DashboardID: "12345",
		Name:        "System Monitoring Dashboard",
		UserID:      "67890",
		Private:     DashboardSharingPrivate,
		Description: "Main monitoring dashboard",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(dashboard)
	}
}

func BenchmarkWidgetMarshaling(b *testing.B) {
	widget := Widget{
		WidgetID:    "67890",
		DashboardID: "12345",
		Type:        WidgetTypeGraph,
		Name:        "CPU Usage Monitor",
		X:           "0",
		Y:           "0",
		Width:       "12",
		Height:      "8",
		Fields: map[string]interface{}{
			"hostid": "12345",
			"itemid": "67890",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(widget)
	}
}

// Test integration scenarios
func TestDashboardsIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test complex filtering
	opts := DashboardGetOptions{
		Filter: map[string]interface{}{
			"private": DashboardSharingPublic,
		},
		Output:       "extend",
		Limit:        100,
		SortField:    "name",
		SortOrder:    "ASC",
	}

	_, err := api.DashboardsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test dashboard creation with widgets
	dashboard := Dashboard{
		Name:        "Integration Test Dashboard",
		Description: "Dashboard created for integration testing",
		Private:     DashboardSharingPrivate,
		AutoCommit:  "1",
		Widgets: []Widget{
			{
				Type:      WidgetTypeGraph,
				Name:      "System Load",
				X:         "0",
				Y:         "0",
				Width:     "12",
				Height:    "8",
				Fields: map[string]interface{}{
					"hostid": "12345",
					"itemid": "67890",
				},
			},
			{
				Type:      WidgetTypeClock,
				Name:      "Current Time",
				X:         "12",
				Y:         "0",
				Width:     "12",
				Height:    "4",
			},
			{
				Type:      WidgetTypeTopHosts,
				Name:      "Top Processes",
				X:         "0",
				Y:         "8",
				Width:     "24",
				Height:    "8",
			},
		},
	}
	
	_, err = api.DashboardCreateSingle(dashboard)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test dashboard update
	dashboard.DashboardID = "11111"
	dashboard.Description = "Updated dashboard description"
	
	_, err = api.DashboardUpdateSingle(dashboard)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test sharing functionality
	err = api.DashboardMakePublic("11111")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test dashboard deletion
	err = api.DashboardDeleteSingle("11111")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}