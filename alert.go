package zabbix

// Alert represent Zabbix Alert object
// https://www.zabbix.com/documentation/current/manual/api/reference/alert/object
type Alert struct {
	AlertID      string `json:"alertid,omitempty"`
	ActionID     string `json:"actionid,omitempty"`
	EventID      string `json:"eventid,omitempty"`
	UserID       string `json:"userid,omitempty"`
	Clock        string `json:"clock,omitempty"`
	MediatypeID  string `json:"mediatypeid,omitempty"`
	SendTo       string `json:"sendto,omitempty"`
	Subject      string `json:"subject,omitempty"`
	Message      string `json:"message,omitempty"`
	Status       string `json:"status,omitempty"`
	Retries      string `json:"retries,omitempty"`
	Error        string `json:"error,omitempty"`
	AlertType    string `json:"alerttype,omitempty"`
	Parameters   string `json:"parameters,omitempty"`
}

// Alerts represents an array of Alert objects
type Alerts []Alert

// AlertGetOptions represents parameters for alert.get API call
type AlertGetOptions struct {
	AlertIDs     []string             `json:"alertids,omitempty"`
	ActionIDs    []string             `json:"actionids,omitempty"`
	EventIDs     []string             `json:"eventids,omitempty"`
	UserIDs      []string             `json:"userids,omitempty"`
	MediaTypeIDs []string             `json:"mediatypeids,omitempty"`
	Filter       map[string]interface{} `json:"filter,omitempty"`
	Search       map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string      `json:"searchWildcardsEnabled,omitempty"`
	TimeFrom     string               `json:"time_from,omitempty"`
	TimeTill     string               `json:"time_till,omitempty"`
	Output       string               `json:"output,omitempty"`
	SelectActions string              `json:"selectActions,omitempty"`
	SelectEvents   string              `json:"selectEvents,omitempty"`
	SelectUsers    string              `json:"selectUsers,omitempty"`
	SelectMediatypes string            `json:"selectMediatypes,omitempty"`
	SortField    string               `json:"sortfield,omitempty"`
	SortOrder    string               `json:"sortorder,omitempty"`
	Limit        int                  `json:"limit,omitempty"`
}

// Alert status constants
const (
	AlertStatusNotSent  = "0"
	AlertStatusSent     = "1"
	AlertStatusFailed   = "2"
	AlertStatusInProgress = "3"
)

// Alert type constants
const (
	AlertTypeMessage = "0"
	AlertTypeCommand = "1"
)

// AlertsGet Wrapper for alert.get
// https://www.zabbix.com/documentation/current/manual/api/reference/alert/get
//
// Zabbix 6.0 Permission Notes:
// - Admin users have limited access to alert data
// - Only Super Admin can access all alert properties
// - Admin users may be restricted from accessing certain alerts based on user permissions
func (api *API) AlertsGet(options AlertGetOptions) (alerts Alerts, err error) {
	params := make(map[string]interface{})
	
	// Convert options to params
	if options.AlertIDs != nil {
		params["alertids"] = options.AlertIDs
	}
	if options.ActionIDs != nil {
		params["actionids"] = options.ActionIDs
	}
	if options.EventIDs != nil {
		params["eventids"] = options.EventIDs
	}
	if options.UserIDs != nil {
		params["userids"] = options.UserIDs
	}
	if options.MediaTypeIDs != nil {
		params["mediatypeids"] = options.MediaTypeIDs
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
	if options.TimeFrom != "" {
		params["time_from"] = options.TimeFrom
	}
	if options.TimeTill != "" {
		params["time_till"] = options.TimeTill
	}
	if options.Output != "" {
		params["output"] = options.Output
	} else {
		// Default to basic fields for Zabbix 6.0 Admin user compatibility
		params["output"] = []string{"alertid", "actionid", "eventid", "clock", "status"}
	}
	if options.SelectActions != "" {
		params["selectActions"] = options.SelectActions
	}
	if options.SelectEvents != "" {
		params["selectEvents"] = options.SelectEvents
	}
	if options.SelectUsers != "" {
		params["selectUsers"] = options.SelectUsers
	}
	if options.SelectMediatypes != "" {
		params["selectMediatypes"] = options.SelectMediatypes
	}
	if options.SortField != "" {
		params["sortfield"] = options.SortField
	}
	if options.SortOrder != "" {
		params["sortorder"] = options.SortOrder
	}
	if options.Limit > 0 {
		params["limit"] = options.Limit
	}

	err = api.CallWithErrorParse("alert.get", params, &alerts)
	return
}

// AlertsGetById Wrapper for alert.get with specific alert IDs
func (api *API) AlertsGetById(alertIds []string) (alerts Alerts, err error) {
	options := AlertGetOptions{
		AlertIDs: alertIds,
		Output: "extend", // Try to get all fields, will be limited by Zabbix 6.0 permissions
	}
	return api.AlertsGet(options)
}

// AlertsGetByAction Wrapper for alert.get with action ID filter
func (api *API) AlertsGetByAction(actionId string) (alerts Alerts, err error) {
	options := AlertGetOptions{
		ActionIDs: []string{actionId},
		Output: "extend",
	}
	return api.AlertsGet(options)
}

// AlertsGetByEvent Wrapper for alert.get with event ID filter
func (api *API) AlertsGetByEvent(eventId string) (alerts Alerts, err error) {
	options := AlertGetOptions{
		EventIDs: []string{eventId},
		Output: "extend",
	}
	return api.AlertsGet(options)
}

// AlertsGetByUser Wrapper for alert.get with user ID filter
func (api *API) AlertsGetByUser(userId string) (alerts Alerts, err error) {
	options := AlertGetOptions{
		UserIDs: []string{userId},
		Output: "extend",
	}
	return api.AlertsGet(options)
}

// AlertsGetByStatus Wrapper for alert.get with status filter
func (api *API) AlertsGetByStatus(status string) (alerts Alerts, err error) {
	options := AlertGetOptions{
		Filter: map[string]interface{}{
			"status": status,
		},
		Output: "extend",
	}
	return api.AlertsGet(options)
}

// AlertsGetFailed Wrapper for alert.get with failed status filter
func (api *API) AlertsGetFailed() (alerts Alerts, err error) {
	return api.AlertsGetByStatus(AlertStatusFailed)
}

// AlertsGetSent Wrapper for alert.get with sent status filter
func (api *API) AlertsGetSent() (alerts Alerts, err error) {
	return api.AlertsGetByStatus(AlertStatusSent)
}

// AlertsGetByTimeRange Wrapper for alert.get with time range filter
func (api *API) AlertsGetByTimeRange(timeFrom, timeTill string) (alerts Alerts, err error) {
	options := AlertGetOptions{
		TimeFrom: timeFrom,
		TimeTill: timeTill,
		Output: "extend",
	}
	return api.AlertsGet(options)
}

// AlertsGetRecent Wrapper for alert.get for recent alerts (last 24 hours)
func (api *API) AlertsGetRecent() (alerts Alerts, err error) {
	// Get current time and subtract 24 hours
	// Note: This is a simplified implementation. In production, you might want to
	// calculate the exact timestamp based on current time
	options := AlertGetOptions{
		TimeFrom: "now-24h",
		Output: "extend",
		SortField: "clock",
		SortOrder: "DESC",
		Limit: 1000, // Limit to prevent too much data
	}
	return api.AlertsGet(options)
}

// AlertsGetByMediaType Wrapper for alert.get with media type ID filter
func (api *API) AlertsGetByMediaType(mediaTypeId string) (alerts Alerts, err error) {
	options := AlertGetOptions{
		MediaTypeIDs: []string{mediaTypeId},
		Output: "extend",
	}
	return api.AlertsGet(options)
}

// AlertsCountByStatus Wrapper for alert.get to count alerts by status
func (api *API) AlertsCountByStatus() (counts map[string]int, err error) {
	counts = make(map[string]int)
	
	// Get alerts for each status
	statuses := []string{AlertStatusNotSent, AlertStatusSent, AlertStatusFailed, AlertStatusInProgress}
	
	for _, status := range statuses {
		options := AlertGetOptions{
			Filter: map[string]interface{}{
				"status": status,
			},
			Output: "count", // Use count output for efficiency
		}
		
		var alerts Alerts
		alerts, err = api.AlertsGet(options)
		if err != nil {
			return
		}
		
		counts[status] = len(alerts)
	}
	
	return
}