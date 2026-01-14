package zabbix

import (
	"encoding/json"
	"fmt"
)

// Dashboard represents a Zabbix dashboard object
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/object
type Dashboard struct {
	DashboardID string   `json:"dashboardid,omitempty"`
	Name        string   `json:"name"`
	UserID      string   `json:"userid,omitempty"`
	Private     string   `json:"private,string"`
	Description string   `json:"description,omitempty"`
	AutoCommit string   `json:"auto_commit,string,omitempty"`
	ShowProblensInMaintenance string `json:"show_problems_in_maintenance,string,omitempty"`
	Widgets     []Widget `json:"widgets,omitempty"`
	
	// Zabbix 7.0+ additional fields
	Sharing      string   `json:"sharing,omitempty"`
	Tags         []Tag   `json:"tags,omitempty"`
	Users        []User  `json:"users,omitempty"`
	UserGroups   []UserGroup `json:"user_groups,omitempty"`
}

// Dashboards represents an array of Dashboard objects
type Dashboards []Dashboard

// Widget represents a dashboard widget
// Note: Zabbix 7.0 has significant changes to Widget structure
type Widget struct {
	WidgetID     string `json:"widgetid,omitempty"`
	DashboardID  string `json:"dashboardid,omitempty"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	View         *WidgetView `json:"view,omitempty"`
	
	// Zabbix 6.0 Widget fields
	Fields       map[string]interface{} `json:"fields,omitempty"`
	Width        string `json:"width,omitempty"`
	Height       string `json:"height,omitempty"`
	X            string `json:"x,omitempty"`
	Y            string `json:"y,omitempty"`
	
	// Zabbix 7.0+ Widget fields (different structure)
	Header       *WidgetHeader `json:"header,omitempty"`
	Body         *WidgetBody   `json:"body,omitempty"`
}

// WidgetView represents widget view configuration (Zabbix 6.0)
type WidgetView struct {
	RefreshInterval string `json:"refresh_interval,omitempty"`
	ShowHeader     string `json:"show_header,omitempty"`
}

// WidgetHeader represents widget header configuration (Zabbix 7.0+)
type WidgetHeader struct {
	Show        string `json:"show,omitempty"`
	Position    string `json:"position,omitempty"`
	Draggable  string `json:"draggable,omitempty"`
}

// WidgetBody represents widget body configuration (Zabbix 7.0+)
type WidgetBody struct {
	View    *WidgetBodyView `json:"view,omitempty"`
	Fields  []WidgetField   `json:"fields,omitempty"`
}

// WidgetBodyView represents widget body view configuration
type WidgetBodyView struct {
	RefreshInterval string `json:"refresh_interval,omitempty"`
	// Add other view fields as needed
}

// WidgetField represents a widget field (Zabbix 7.0+)
type WidgetField struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// WidgetFieldV6 represents widget field for Zabbix 6.0 compatibility
type WidgetFieldV6 struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// DashboardShare represents dashboard sharing information
type DashboardShare struct {
	Users     []User     `json:"users,omitempty"`
	UserGroups []UserGroup `json:"user_groups,omitempty"`
	Private   string     `json:"private,string"`
}

// DashboardGetOptions represents parameters for dashboard.get API call
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/get
type DashboardGetOptions struct {
	DashboardIDs    []string             `json:"dashboardids,omitempty"`
	Filter          map[string]interface{} `json:"filter,omitempty"`
	Search          map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string        `json:"searchWildcardsEnabled,omitempty"`
	Output          string               `json:"output,omitempty"`
	SelectUsers     string               `json:"selectUsers,omitempty"`
	SelectUserGroups string              `json:"selectUserGroups,omitempty"`
	SelectSharing   string               `json:"selectSharing,omitempty"`
	SelectTags      string               `json:"selectTags,omitempty"`
	SelectWidgets   string               `json:"selectWidgets,omitempty"`
	SortField       string               `json:"sortfield,omitempty"`
	SortOrder       string               `json:"sortorder,omitempty"`
	Limit           int                  `json:"limit,omitempty"`
}

// DashboardCreateOptions represents parameters for dashboard.create API call
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/create
type DashboardCreateOptions struct {
	Dashboards   Dashboards `json:"dashboards"`
	SelectUsers   string     `json:"selectUsers,omitempty"`
	SelectUserGroups string  `json:"selectUserGroups,omitempty"`
	SelectSharing string     `json:"selectSharing,omitempty"`
	SelectTags    string     `json:"selectTags,omitempty"`
	SelectWidgets string     `json:"selectWidgets,omitempty"`
}

// DashboardUpdateOptions represents parameters for dashboard.update API call
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/update
type DashboardUpdateOptions struct {
	Dashboards   Dashboards `json:"dashboards"`
	SelectUsers   string     `json:"selectUsers,omitempty"`
	SelectUserGroups string  `json:"selectUserGroups,omitempty"`
	SelectSharing string     `json:"selectSharing,omitempty"`
	SelectTags    string     `json:"selectTags,omitempty"`
	SelectWidgets string     `json:"selectWidgets,omitempty"`
}

// DashboardDeleteOptions represents parameters for dashboard.delete API call
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/delete
type DashboardDeleteOptions struct {
	DashboardIDs []string `json:"dashboardids"`
}

// Dashboard constants
const (
	// Dashboard sharing modes
	DashboardSharingPrivate  = "0"
	DashboardSharingPublic   = "1"
	
	// Dashboard refresh intervals
	DashboardRefresh30s  = "30"
	DashboardRefresh1m   = "60"
	DashboardRefresh2m   = "120"
	DashboardRefresh5m   = "300"
	DashboardRefresh10m  = "600"
	DashboardRefresh30m  = "1800"
	DashboardRefresh1h   = "3600"
	
	// Widget types (common ones)
	WidgetTypeClock         = "clock"
	WidgetTypeDataOverview  = "data_overview"
	WidgetTypeDiscovery     = "discovery"
	WidgetTypeGraph        = "graph"
	WidgetTypeGraphPrototype = "graph_prototype"
	WidgetTypeHostIssues   = "host_issues"
	WidgetTypeMap          = "map"
	WidgetTypePlainText    = "plain_text" // Deprecated in Zabbix 7.0, use WidgetTypeItemHistory instead
	WidgetTypeSlideShow   = "slideshow"
	WidgetTypeSvgGraph     = "svg_graph"
	WidgetTypeSystemInfo   = "system_info"
	WidgetTypeTopHosts     = "top_hosts"
	WidgetTypeTriggerInfo  = "trigger_info"
	WidgetTypeTriggerOverview = "trigger_overview"
	WidgetTypeUrl         = "url"
	WidgetTypeWebMonitoring = "web"
	
	// Zabbix 7.0+ Widget types
	WidgetTypeItemHistory = "item_history" // Replaces plaintext in Zabbix 7.0
)

// Widget constants
const (
	// Widget position constants
	WidgetPositionLeft   = "left"
	WidgetPositionRight  = "right"
	
	// Widget header options
	WidgetHeaderShow    = "1"
	WidgetHeaderHide    = "0"
	WidgetHeaderDraggable = "1"
	WidgetHeaderFixed   = "0"
)

// DashboardsGet Wrapper for dashboard.get
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/get
func (api *API) DashboardsGet(options DashboardGetOptions) (Dashboards, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.DashboardIDs) > 0 {
		params["dashboardids"] = options.DashboardIDs
	}
	if options.Filter != nil {
		params["filter"] = options.Filter
	}
	if options.Search != nil {
		params["search"] = options.Search
	}
	if options.SearchWildcardsEnabled != "" {
		params["searchWildcardsEnabled"] = options.SearchWildcardsEnabled
	}
	if options.Output != "" {
		params["output"] = options.Output
	} else {
		params["output"] = "extend"
	}
	if options.SelectUsers != "" {
		params["selectUsers"] = options.SelectUsers
	}
	if options.SelectUserGroups != "" {
		params["selectUserGroups"] = options.SelectUserGroups
	}
	if options.SelectSharing != "" {
		params["selectSharing"] = options.SelectSharing
	}
	if options.SelectTags != "" {
		params["selectTags"] = options.SelectTags
	}
	if options.SelectWidgets != "" {
		params["selectWidgets"] = options.SelectWidgets
	}
	if options.SortField != "" {
		params["sortfield"] = options.SortField
	} else {
		params["sortfield"] = "name"
	}
	if options.SortOrder != "" {
		params["sortorder"] = options.SortOrder
	} else {
		params["sortorder"] = "ASC"
	}
	if options.Limit > 0 {
		params["limit"] = options.Limit
	}

	var dashboards Dashboards
	err := api.CallWithErrorParse("dashboard.get", params, &dashboards)
	return dashboards, err
}

// DashboardsGetByID Get dashboards by specific dashboard IDs
func (api *API) DashboardsGetByID(dashboardIDs []string) (Dashboards, error) {
	options := DashboardGetOptions{
		DashboardIDs: dashboardIDs,
		Output:       "extend",
	}
	return api.DashboardsGet(options)
}

// DashboardsGetByName Get dashboards by name
func (api *API) DashboardsGetByName(name string) (Dashboards, error) {
	options := DashboardGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.DashboardsGet(options)
}

// DashboardsGetByUser Get dashboards by user ID
func (api *API) DashboardsGetByUser(userID string) (Dashboards, error) {
	options := DashboardGetOptions{
		Filter: map[string]interface{}{
			"userid": userID,
		},
		Output: "extend",
	}
	return api.DashboardsGet(options)
}

// DashboardsGetPrivate Get private dashboards
func (api *API) DashboardsGetPrivate() (Dashboards, error) {
	options := DashboardGetOptions{
		Filter: map[string]interface{}{
			"private": DashboardSharingPrivate,
		},
		Output: "extend",
	}
	return api.DashboardsGet(options)
}

// DashboardsGetPublic Get public dashboards
func (api *API) DashboardsGetPublic() (Dashboards, error) {
	options := DashboardGetOptions{
		Filter: map[string]interface{}{
			"private": DashboardSharingPublic,
		},
		Output: "extend",
	}
	return api.DashboardsGet(options)
}

// DashboardGetByID Get dashboard by ID (exactly one match required)
func (api *API) DashboardGetByID(dashboardID string) (*Dashboard, error) {
	dashboards, err := api.DashboardsGetByID([]string{dashboardID})
	if err != nil {
		return nil, err
	}

	if len(dashboards) == 1 {
		return &dashboards[0], nil
	} else if len(dashboards) == 0 {
		return nil, fmt.Errorf("Dashboard not found: %s", dashboardID)
	} else {
		return nil, fmt.Errorf("Multiple dashboards found with ID: %s", dashboardID)
	}
}

// DashboardsGetWithWidgets Get dashboards with their widgets
func (api *API) DashboardsGetWithWidgets(options DashboardGetOptions) (Dashboards, error) {
	options.SelectWidgets = "extend"
	return api.DashboardsGet(options)
}

// DashboardsCreate Wrapper for dashboard.create
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/create
func (api *API) DashboardsCreate(dashboards Dashboards) (result []string, err error) {
	options := DashboardCreateOptions{
		Dashboards: dashboards,
	}

	response, err := api.CallWithError("dashboard.create", options)
	if err != nil {
		return
	}

	var rawResult interface{}
	err = json.Unmarshal(response.Result, &rawResult)
	if err != nil {
		return
	}

	if resultArray, ok := rawResult.([]interface{}); ok {
		for _, item := range resultArray {
			if dashboardMap, ok := item.(map[string]interface{}); ok {
				if dashboardid, exists := dashboardMap["dashboardids"]; exists {
					if idArray, ok := dashboardid.([]interface{}); ok && len(idArray) > 0 {
						if id, ok := idArray[0].(string); ok {
							result = append(result, id)
						}
					}
				}
			}
		}
	}
	return
}

// DashboardCreateSingle Create a single dashboard
func (api *API) DashboardCreateSingle(dashboard Dashboard) (dashboardID string, err error) {
	dashboards := Dashboards{dashboard}
	result, err := api.DashboardsCreate(dashboards)
	if len(result) > 0 {
		dashboardID = result[0]
	}
	return
}

// DashboardsUpdate Wrapper for dashboard.update
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/update
func (api *API) DashboardsUpdate(dashboards Dashboards) (result []string, err error) {
	options := DashboardUpdateOptions{
		Dashboards: dashboards,
	}

	response, err := api.CallWithError("dashboard.update", options)
	if err != nil {
		return
	}

	var rawResult interface{}
	err = json.Unmarshal(response.Result, &rawResult)
	if err != nil {
		return
	}

	if resultArray, ok := rawResult.([]interface{}); ok {
		for _, item := range resultArray {
			if dashboardMap, ok := item.(map[string]interface{}); ok {
				if dashboardid, exists := dashboardMap["dashboardids"]; exists {
					if idArray, ok := dashboardid.([]interface{}); ok && len(idArray) > 0 {
						if id, ok := idArray[0].(string); ok {
							result = append(result, id)
						}
					}
				}
			}
		}
	}
	return
}

// DashboardUpdateSingle Update a single dashboard
func (api *API) DashboardUpdateSingle(dashboard Dashboard) (dashboardID string, err error) {
	dashboards := Dashboards{dashboard}
	result, err := api.DashboardsUpdate(dashboards)
	if len(result) > 0 {
		dashboardID = result[0]
	}
	return
}

// DashboardsDelete Wrapper for dashboard.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/dashboard/delete
func (api *API) DashboardsDelete(dashboards Dashboards) (result []string, err error) {
	dashboardIDs := make([]string, len(dashboards))
	for i, dashboard := range dashboards {
		dashboardIDs[i] = dashboard.DashboardID
	}
	
	return api.DashboardsDeleteByIDs(dashboardIDs)
}

// DashboardsDeleteByIDs Wrapper for dashboard.delete with IDs
func (api *API) DashboardsDeleteByIDs(dashboardIDs []string) (result []string, err error) {
	options := DashboardDeleteOptions{
		DashboardIDs: dashboardIDs,
	}

	response, err := api.CallWithError("dashboard.delete", options)
	if err != nil {
		return
	}

	var rawResult interface{}
	err = json.Unmarshal(response.Result, &rawResult)
	if err != nil {
		return
	}

	if resultArray, ok := rawResult.([]interface{}); ok {
		for _, item := range resultArray {
			if id, ok := item.(string); ok {
				result = append(result, id)
			}
		}
	}
	return
}

// DashboardDeleteSingle Delete a single dashboard
func (api *API) DashboardDeleteSingle(dashboardID string) (err error) {
	_, err = api.DashboardsDeleteByIDs([]string{dashboardID})
	return
}

// DashboardShare Set dashboard sharing
func (api *API) DashboardShare(dashboardID string, sharing DashboardShare) error {
	// Get current dashboard
	dashboard, err := api.DashboardGetByID(dashboardID)
	if err != nil {
		return err
	}
	
	// Update sharing settings
	dashboard.Sharing = sharing.Private
	dashboard.Users = sharing.Users
	dashboard.UserGroups = sharing.UserGroups
	
	_, err = api.DashboardUpdateSingle(*dashboard)
	return err
}

// DashboardMakePublic Make a dashboard public
func (api *API) DashboardMakePublic(dashboardID string) error {
	return api.DashboardShare(dashboardID, DashboardShare{
		Private: DashboardSharingPublic,
	})
}

// DashboardMakePrivate Make a dashboard private
func (api *API) DashboardMakePrivate(dashboardID string) error {
	return api.DashboardShare(dashboardID, DashboardShare{
		Private: DashboardSharingPrivate,
	})
}

// DashboardAddUser Share dashboard with specific users
func (api *API) DashboardAddUser(dashboardID string, users []User) error {
	return api.DashboardShare(dashboardID, DashboardShare{
		Private: DashboardSharingPublic,
		Users:   users,
	})
}

// DashboardAddUserGroup Share dashboard with specific user groups
func (api *API) DashboardAddUserGroup(dashboardID string, userGroups []UserGroup) error {
	return api.DashboardShare(dashboardID, DashboardShare{
		Private:    DashboardSharingPublic,
		UserGroups: userGroups,
	})
}

// DashboardsGetStatistics Get statistics about dashboards
func (api *API) DashboardsGetStatistics() (map[string]interface{}, error) {
	// Get all dashboards
	allDashboards, err := api.DashboardsGet(DashboardGetOptions{Output: "extend"})
	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})
	
	// Basic counts
	totalDashboards := len(allDashboards)
	privateDashboards := 0
	publicDashboards := 0
	
	// Group by user
	userCounts := make(map[string]int)
	
	// Count widgets per dashboard
	widgetCounts := make(map[string]int)
	
	for _, dashboard := range allDashboards {
		// Count by sharing type
		if dashboard.Private == DashboardSharingPrivate {
			privateDashboards++
		} else {
			publicDashboards++
		}
		
		// Count by user
		if dashboard.UserID != "" {
			userCounts[dashboard.UserID]++
		}
		
		// Count widgets
		widgetCounts[dashboard.DashboardID] = len(dashboard.Widgets)
	}
	
	// Build statistics result
	stats["total_dashboards"] = totalDashboards
	stats["private_dashboards"] = privateDashboards
	stats["public_dashboards"] = publicDashboards
	stats["sharing_distribution"] = map[string]int{
		"private": privateDashboards,
		"public":  publicDashboards,
	}
	stats["user_distribution"] = userCounts
	stats["widget_counts"] = widgetCounts
	
	// Calculate statistics
	if totalDashboards > 0 {
		stats["private_percentage"] = (float64(privateDashboards) / float64(totalDashboards)) * 100
		stats["public_percentage"] = (float64(publicDashboards) / float64(totalDashboards)) * 100
	} else {
		stats["private_percentage"] = 0
		stats["public_percentage"] = 0
	}
	
	// Calculate average widgets per dashboard
	totalWidgets := 0
	for _, count := range widgetCounts {
		totalWidgets += count
	}
	
	if totalDashboards > 0 {
		stats["average_widgets_per_dashboard"] = float64(totalWidgets) / float64(totalDashboards)
	} else {
		stats["average_widgets_per_dashboard"] = 0
	}
	
	return stats, nil
}

// DashboardValidate Validate a dashboard configuration
func (api *API) DashboardValidate(dashboard Dashboard) (validationErrors []string) {
	validationErrors = []string{}
	
	// Check required fields
	if dashboard.Name == "" {
		validationErrors = append(validationErrors, "Dashboard name is required")
	}
	
	// Validate widgets if present
	for i, widget := range dashboard.Widgets {
		if widget.Type == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Widget %d: Type is required", i))
		}
		
		if widget.Name == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Widget %d: Name is required", i))
		}
		
		// Validate widget positioning
		if widget.X != "" && widget.Y != "" {
			// Basic validation for Zabbix 6.0 coordinates
			if x := parseInt(widget.X); x < 0 || x > 23 {
				validationErrors = append(validationErrors, fmt.Sprintf("Widget %d: X coordinate must be between 0-23", i))
			}
			if y := parseInt(widget.Y); y < 0 || y > 62 {
				validationErrors = append(validationErrors, fmt.Sprintf("Widget %d: Y coordinate must be between 0-62", i))
			}
		}
		
		// Validate widget size
		if widget.Width != "" && widget.Height != "" {
			if width := parseInt(widget.Width); width < 1 || width > 24 {
				validationErrors = append(validationErrors, fmt.Sprintf("Widget %d: Width must be between 1-24", i))
			}
			if height := parseInt(widget.Height); height < 2 || height > 32 {
				validationErrors = append(validationErrors, fmt.Sprintf("Widget %d: Height must be between 2-32", i))
			}
		}
	}
	
	return validationErrors
}

// Helper function to parse integer from string
func parseInt(s string) int {
	if s == "" {
		return 0
	}
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}

// DashboardIsPublic Check if dashboard is public
func (dashboard *Dashboard) IsPublic() bool {
	return dashboard.Private == DashboardSharingPublic
}

// DashboardIsPrivate Check if dashboard is private
func (dashboard *Dashboard) IsPrivate() bool {
	return dashboard.Private == DashboardSharingPrivate
}

// DashboardHasWidgets Check if dashboard has widgets
func (dashboard *Dashboard) HasWidgets() bool {
	return len(dashboard.Widgets) > 0
}

// DashboardGetWidgetCount Get number of widgets
func (dashboard *Dashboard) GetWidgetCount() int {
	return len(dashboard.Widgets)
}

// DashboardAddWidget Add a widget to dashboard
func (dashboard *Dashboard) AddWidget(widget Widget) {
	dashboard.Widgets = append(dashboard.Widgets, widget)
}

// DashboardRemoveWidget Remove a widget from dashboard
func (dashboard *Dashboard) RemoveWidget(widgetID string) {
	for i, widget := range dashboard.Widgets {
		if widget.WidgetID == widgetID {
			dashboard.Widgets = append(dashboard.Widgets[:i], dashboard.Widgets[i+1:]...)
			break
		}
	}
}

// DashboardUpdateWidget Update a widget in dashboard
func (dashboard *Dashboard) UpdateWidget(widget Widget) {
	for i, w := range dashboard.Widgets {
		if w.WidgetID == widget.WidgetID {
			dashboard.Widgets[i] = widget
			break
		}
	}
}

// CreateSimpleDashboard Create a simple dashboard with basic widgets
func (api *API) CreateSimpleDashboard(name, description string, userID string) (string, error) {
	dashboard := Dashboard{
		Name:        name,
		Description: description,
		UserID:      userID,
		Private:     DashboardSharingPrivate,
		AutoCommit:  "1",
	}
	
	return api.DashboardCreateSingle(dashboard)
}

// CreateSystemOverviewDashboard Create a system overview dashboard
func (api *API) CreateSystemOverviewDashboard(userID string) (string, error) {
	dashboard := Dashboard{
		Name:        "System Overview",
		Description: "System overview dashboard with key metrics",
		UserID:      userID,
		Private:     DashboardSharingPrivate,
		AutoCommit:  "1",
		Widgets: []Widget{
			{
				Type:      WidgetTypeHostIssues,
				Name:      "Host Issues",
				X:         "0",
				Y:         "0",
				Width:     "12",
				Height:    "8",
				Fields:    map[string]interface{}{},
			},
			{
				Type:      WidgetTypeTopHosts,
				Name:      "Top Hosts",
				X:         "12",
				Y:         "0",
				Width:     "12",
				Height:    "8",
				Fields:    map[string]interface{}{},
			},
		},
	}
	
	return api.DashboardCreateSingle(dashboard)
}

// GetUserDashboards Get all dashboards owned by a user
func (api *API) GetUserDashboards(userID string) (Dashboards, error) {
	return api.DashboardsGetByUser(userID)
}