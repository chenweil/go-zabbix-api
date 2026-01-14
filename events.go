package zabbix

import (
	"encoding/json"
	"fmt"
	"time"
)

// Event represents a Zabbix event object
// https://www.zabbix.com/documentation/current/manual/api/reference/event/object
type Event struct {
	EventID       string `json:"eventid"`
	Source       string `json:"source"`
	Object        string `json:"object"`
	ObjectID      string `json:"objectid"`
	Clock         int    `json:"clock"`
	ns            int    `json:"ns"`
	Value         int    `json:"value"`
	Acknowledged  int    `json:"acknowledged"`
	Flags         string `json:"flags"`
	Severity      string `json:"severity"`
	RecoveryEvent string `json:"recovery_event"`
	CorrelationID string `json:"correlation_id"`
	UserID        string `json:"userid"`
	Tags          []Tag  `json:"tags,omitempty"`
	
	// Additional fields for extended event information
	URL           string `json:"url,omitempty"`
	TriggerID     string `json:"triggerid,omitempty"`
	HostID        string `json:"hostid,omitempty"`
	Host          string `json:"host,omitempty"`
	ItemID        string `json:"itemid,omitempty"`
	Description   string `json:"description,omitempty"`
	Priority      string `json:"priority,omitempty"`
	Status        string `json:"status,omitempty"`
	
	// Related objects (optional based on API response)
	Trigger       *Trigger `json:"trigger,omitempty"`
	Hosts         []Host   `json:"hosts,omitempty"`
	RelatedEvents []Event  `json:"related_events,omitempty"`
}

// Events represents an array of Event objects
type Events []Event

// EventAcknowledgement represents an event acknowledgement
type EventAcknowledgement struct {
	EventID   string `json:"eventid"`
	Message   string `json:"message"`
	Action    int    `json:"action"`
	Severity  string `json:"severity,omitempty"`
	KeepAlert bool   `json:"keep_alert,omitempty"`
}

// EventAcknowledgements represents an array of event acknowledgements
type EventAcknowledgements []EventAcknowledgement

// EventGetOptions represents parameters for event.get API call
// https://www.zabbix.com/documentation/current/manual/api/reference/event/get
type EventGetOptions struct {
	EventIDs         []string             `json:"eventids,omitempty"`
	GroupIDs         []string             `json:"groupids,omitempty"`
	HostIDs          []string             `json:"hostids,omitempty"`
	TriggerIDs       []string             `json:"triggerids,omitempty"`
	EventIDFrom      string               `json:"eventid_from,omitempty"`
	EventIDTill      string               `json:"eventid_till,omitempty"`
	TimeFrom         int                 `json:"time_from,omitempty"`
	TimeTill         int                 `json:"time_till,omitempty"`
	ProblemTimeFrom  int                 `json:"problem_time_from,omitempty"`
	ProblemTimeTill  int                 `json:"problem_time_till,omitempty"`
	Value            []int                `json:"value,omitempty"`
	Severity        []string             `json:"severity,omitempty"`
	Source           []int                `json:"source,omitempty"`
	Object           []int                `json:"object,omitempty"`
	Filter           map[string]interface{} `json:"filter,omitempty"`
	Search           map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string        `json:"searchWildcardsEnabled,omitempty"`
	Output           string               `json:"output,omitempty"`
	ExcludeSearch    bool                 `json:"exclude_search,omitempty"`
	LimitSelectLimit int                  `json:"limit_select_limit,omitempty"`
	SelectHosts      string               `json:"selectHosts,omitempty"`
	SelectRelatedObject string            `json:"selectRelatedObject,omitempty"`
	SelectTags       string               `json:"selectTags,omitempty"`
	SelectAcknowledges string            `json:"selectAcknowledges,omitempty"`
	SortField        string               `json:"sortfield,omitempty"`
	SortOrder        string               `json:"sortorder,omitempty"`
	Limit            int                  `json:"limit,omitempty"`
}

// EventSeverity constants
const (
	EventSeverityNotClassified = "0"
	EventSeverityInformation   = "1"
	EventSeverityWarning       = "2"
	EventSeverityAverage       = "3"
	EventSeverityHigh          = "4"
	EventSeverityDisaster      = "5"
)

// EventSource constants
const (
	EventSourceTrigger   = "0"
	EventSourceDiscovery = "1"
	EventSourceAutoreg   = "2"
	EventSourceInternal  = "3"
	EventSourceHTTPAgent = "4"
)

// EventObject constants
const (
	EventObjectTrigger   = "0"
	EventObjectDiscovery = "1"
	EventObjectAutoreg   = "2"
	EventObjectITService = "3"
)

// EventValue constants
const (
	EventValueOK       = "0"
	EventValueProblem   = "1"
	EventValueNotClassified = "0"
	EventValueClassified    = "1"
)

// EventAcknowledgementAction constants
const (
	ActionNone       = "0"
	ActionCloseProblem = "1"
	ActionAcknowledge  = "2"
	ActionCloseEvent   = "3"
)

// EventSortField constants
const (
	EventSortClock        = "clock"
	EventSortEventID      = "eventid"
	EventSortSeverity     = "severity"
	EventSortAcknowledged = "acknowledged"
)

// EventsGet Wrapper for event.get
// https://www.zabbix.com/documentation/current/manual/api/reference/event/get
func (api *API) EventsGet(options EventGetOptions) (Events, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.EventIDs) > 0 {
		params["eventids"] = options.EventIDs
	}
	if len(options.GroupIDs) > 0 {
		params["groupids"] = options.GroupIDs
	}
	if len(options.HostIDs) > 0 {
		params["hostids"] = options.HostIDs
	}
	if len(options.TriggerIDs) > 0 {
		params["triggerids"] = options.TriggerIDs
	}
	if options.EventIDFrom != "" {
		params["eventid_from"] = options.EventIDFrom
	}
	if options.EventIDTill != "" {
		params["eventid_till"] = options.EventIDTill
	}
	if options.TimeFrom > 0 {
		params["time_from"] = options.TimeFrom
	}
	if options.TimeTill > 0 {
		params["time_till"] = options.TimeTill
	}
	if options.ProblemTimeFrom > 0 {
		params["problem_time_from"] = options.ProblemTimeFrom
	}
	if options.ProblemTimeTill > 0 {
		params["problem_time_till"] = options.ProblemTimeTill
	}
	if len(options.Value) > 0 {
		params["value"] = options.Value
	}
	if len(options.Severity) > 0 {
		params["severity"] = options.Severity
	}
	if len(options.Source) > 0 {
		params["source"] = options.Source
	}
	if len(options.Object) > 0 {
		params["object"] = options.Object
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
	if options.ExcludeSearch {
		params["exclude_search"] = options.ExcludeSearch
	}
	if options.LimitSelectLimit > 0 {
		params["limit_select_limit"] = options.LimitSelectLimit
	}
	if options.SelectHosts != "" {
		params["selectHosts"] = options.SelectHosts
	}
	if options.SelectRelatedObject != "" {
		params["selectRelatedObject"] = options.SelectRelatedObject
	}
	if options.SelectTags != "" {
		params["selectTags"] = options.SelectTags
	}
	if options.SelectAcknowledges != "" {
		params["selectAcknowledges"] = options.SelectAcknowledges
	}
	if options.SortField != "" {
		params["sortfield"] = options.SortField
	} else {
		// Default sort by clock (most recent first)
		params["sortfield"] = EventSortClock
	}
	if options.SortOrder != "" {
		params["sortorder"] = options.SortOrder
	} else {
		// Default sort order (newest first)
		params["sortorder"] = "DESC"
	}
	if options.Limit > 0 {
		params["limit"] = options.Limit
	} else {
		// Set a reasonable default limit
		params["limit"] = 1000
	}

	var events Events
	err := api.CallWithErrorParse("event.get", params, &events)
	return events, err
}

// EventsGetByID Get events by specific event IDs
func (api *API) EventsGetByID(eventIDs []string) (Events, error) {
	options := EventGetOptions{
		EventIDs: eventIDs,
		Output:   "extend",
	}
	return api.EventsGet(options)
}

// EventsGetByTrigger Get events by trigger IDs
func (api *API) EventsGetByTrigger(triggerIDs []string) (Events, error) {
	options := EventGetOptions{
		TriggerIDs: triggerIDs,
		Output:     "extend",
	}
	return api.EventsGet(options)
}

// EventsGetByHost Get events by host IDs
func (api *API) EventsGetByHost(hostIDs []string) (Events, error) {
	options := EventGetOptions{
		HostIDs: hostIDs,
		Output:  "extend",
	}
	return api.EventsGet(options)
}

// EventsGetByTimeRange Get events within a time range
func (api *API) EventsGetByTimeRange(timeFrom, timeTill int) (Events, error) {
	options := EventGetOptions{
		TimeFrom: timeFrom,
		TimeTill: timeTill,
		Output:   "extend",
	}
	return api.EventsGet(options)
}

// EventsGetProblems Get problem events (non-OK state events)
func (api *API) EventsGetProblems() (Events, error) {
	options := EventGetOptions{
		Value:   []int{1}, // Problem state
		Output:  "extend",
		Limit:   1000,
	}
	return api.EventsGet(options)
}

// EventsGetOK Get OK state events
func (api *API) EventsGetOK() (Events, error) {
	options := EventGetOptions{
		Value:  []int{0}, // OK state
		Output: "extend",
		Limit:  1000,
	}
	return api.EventsGet(options)
}

// EventsGetBySeverity Get events by severity level
func (api *API) EventsGetBySeverity(severity []string) (Events, error) {
	options := EventGetOptions{
		Severity: severity,
		Output:   "extend",
		Limit:    1000,
	}
	return api.EventsGet(options)
}

// EventsGetHighSeverity Get high severity events (High and Disaster)
func (api *API) EventsGetHighSeverity() (Events, error) {
	severity := []string{EventSeverityHigh, EventSeverityDisaster}
	return api.EventsGetBySeverity(severity)
}

// EventsGetCritical Get critical events (Disaster severity only)
func (api *API) EventsGetCritical() (Events, error) {
	severity := []string{EventSeverityDisaster}
	return api.EventsGetBySeverity(severity)
}

// EventsGetWarning Get warning level events
func (api *API) EventsGetWarning() (Events, error) {
	severity := []string{EventSeverityWarning}
	return api.EventsGetBySeverity(severity)
}

// EventsGetRecent Get recent events (last hour by default)
func (api *API) EventsGetRecent(hours int) (Events, error) {
	currentTime := int(time.Now().Unix())
	timeFrom := currentTime - (hours * 3600)
	
	options := EventGetOptions{
		TimeFrom: timeFrom,
		TimeTill: currentTime,
		Output:   "extend",
		Limit:    500,
	}
	return api.EventsGet(options)
}

// EventsGetUnacknowledged Get unacknowledged events
func (api *API) EventsGetUnacknowledged() (Events, error) {
	options := EventGetOptions{
		Filter: map[string]interface{}{
			"acknowledged": "0",
		},
		Output: "extend",
		Limit:  1000,
	}
	return api.EventsGet(options)
}

// EventsGetAcknowledged Get acknowledged events
func (api *API) EventsGetAcknowledged() (Events, error) {
	options := EventGetOptions{
		Filter: map[string]interface{}{
			"acknowledged": "1",
		},
		Output: "extend",
		Limit:  1000,
	}
	return api.EventsGet(options)
}

// EventsGetBySource Get events by source
func (api *API) EventsGetBySource(source int) (Events, error) {
	options := EventGetOptions{
		Source: []int{source},
		Output: "extend",
		Limit:  1000,
	}
	return api.EventsGet(options)
}

// EventsGetTriggerEvents Get trigger events specifically
func (api *API) EventsGetTriggerEvents() (Events, error) {
	return api.EventsGetBySource(EventSourceTrigger)
}

// EventsGetDiscoveryEvents Get discovery events specifically
func (api *API) EventsGetDiscoveryEvents() (Events, error) {
	return api.EventsGetBySource(EventSourceDiscovery)
}

// EventsGetInternalEvents Get internal events specifically
func (api *API) EventsGetInternalEvents() (Events, error) {
	return api.EventsGetBySource(EventSourceInternal)
}

// EventsAcknowledge Wrapper for event.acknowledge
// https://www.zabbix.com/documentation/current/manual/api/reference/event/acknowledge
func (api *API) EventsAcknowledge(eventIDs []string, message string) (success bool, err error) {
	params := map[string]interface{}{
		"eventids": eventIDs,
		"message":  message,
	}

	var result map[string]interface{}
	err = api.CallWithErrorParse("event.acknowledge", params, &result)
	if err != nil {
		return false, err
	}

	// Check if the operation was successful
	if result["success"] != nil {
		return true, nil
	}

	return false, fmt.Errorf("failed to acknowledge events")
}

// EventsAcknowledgeWithAction Acknowledge events with specific action
func (api *API) EventsAcknowledgeWithAction(eventIDs []string, message string, action int, severity string) (success bool, err error) {
	params := map[string]interface{}{
		"eventids": eventIDs,
		"message":  message,
		"action":   action,
	}
	
	if severity != "" {
		params["severity"] = severity
	}

	var result map[string]interface{}
	err = api.CallWithErrorParse("event.acknowledge", params, &result)
	if err != nil {
		return false, err
	}

	// Check if the operation was successful
	if result["success"] != nil {
		return true, nil
	}

	return false, fmt.Errorf("failed to acknowledge events")
}

// EventsCloseProblem Close problem events
func (api *API) EventsCloseProblem(eventIDs []string, message string) (success bool, err error) {
	return api.EventsAcknowledgeWithAction(eventIDs, message, ActionCloseProblem, "")
}

// EventsCloseEvent Close events
func (api *API) EventsCloseEvent(eventIDs []string, message string) (success bool, err error) {
	return api.EventsAcknowledgeWithAction(eventIDs, message, ActionCloseEvent, "")
}

// EventsAcknowledgeSingle Acknowledge a single event
func (api *API) EventsAcknowledgeSingle(eventID string, message string) (success bool, err error) {
	return api.EventsAcknowledge([]string{eventID}, message)
}

// EventsAcknowledgeProblem Acknowledge problem events
func (api *API) EventsAcknowledgeProblem(eventIDs []string, message string) (success bool, err error) {
	return api.EventsAcknowledgeWithAction(eventIDs, message, ActionAcknowledge, "")
}

// EventsFilterByTimeRange Filter events by time range with additional criteria
func (api *API) EventsFilterByTimeRange(timeFrom, timeTill int, criteria map[string]interface{}) (Events, error) {
	options := EventGetOptions{
		TimeFrom: timeFrom,
		TimeTill: timeTill,
		Output:   "extend",
		Limit:    1000,
	}
	
	if criteria != nil {
		options.Filter = criteria
	}
	
	return api.EventsGet(options)
}

// EventsFilterByHostAndTrigger Filter events by host and trigger
func (api *API) EventsFilterByHostAndTrigger(hostIDs, triggerIDs []string) (Events, error) {
	options := EventGetOptions{
		HostIDs:     hostIDs,
		TriggerIDs:  triggerIDs,
		Output:     "extend",
		Limit:      1000,
	}
	
	return api.EventsGet(options)
}

// EventsFilterBySeverityAndTime Filter events by severity and time range
func (api *API) EventsFilterBySeverityAndTime(severity []string, timeFrom, timeTill int) (Events, error) {
	options := EventGetOptions{
		Severity: severity,
		TimeFrom: timeFrom,
		TimeTill: timeTill,
		Output:   "extend",
		Limit:    1000,
	}
	
	return api.EventsGet(options)
}

// EventsGetStatistics Get statistics about events
func (api *API) EventsGetStatistics(timeFrom, timeTill int) (map[string]interface{}, error) {
	events, err := api.EventsGetByTimeRange(timeFrom, timeTill)
	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})
	
	// Basic counts
	totalEvents := len(events)
	problemEvents := 0
	okEvents := 0
	acknowledgedEvents := 0
	unacknowledgedEvents := 0
	
	// Count by severity
	severityCounts := make(map[string]int)
	
	// Count by source
	sourceCounts := make(map[string]int)
	
	// Count by hour
	hourlyCounts := make(map[string]int)
	
	for _, event := range events {
		// Count states
		if event.Value == 1 {
			problemEvents++
		} else {
			okEvents++
		}
		
		// Count acknowledgements
		if event.Acknowledged == 1 {
			acknowledgedEvents++
		} else {
			unacknowledgedEvents++
		}
		
		// Count by severity
		if event.Severity != "" {
			severityCounts[event.Severity]++
		}
		
		// Count by source
		sourceCounts[event.Source]++
		
		// Count by hour
		hour := time.Unix(int64(event.Clock), 0).Format("2006-01-02 15:00")
		hourlyCounts[hour]++
	}
	
	// Build statistics result
	stats["total_events"] = totalEvents
	stats["problem_events"] = problemEvents
	stats["ok_events"] = okEvents
	stats["acknowledged_events"] = acknowledgedEvents
	stats["unacknowledged_events"] = unacknowledgedEvents
	stats["acknowledgment_rate"] = float64(acknowledgedEvents) / float64(totalEvents) * 100
	
	stats["severity_distribution"] = severityCounts
	stats["source_distribution"] = sourceCounts
	stats["hourly_distribution"] = hourlyCounts
	
	// Time range info
	stats["time_range"] = map[string]interface{}{
		"from": timeFrom,
		"till": timeTill,
		"duration_seconds": timeTill - timeFrom,
	}
	
	return stats, nil
}

// EventsGetSummary Get summary of events for a specific time period
func (api *API) EventsGetSummary(hours int) (map[string]interface{}, error) {
	currentTime := int(time.Now().Unix())
	timeFrom := currentTime - (hours * 3600)
	
	return api.EventsGetStatistics(timeFrom, currentTime)
}

// EventsCheckHealth Check the health of events system
func (api *API) EventsCheckHealth() (map[string]interface{}, error) {
	health := make(map[string]interface{})
	
	// Check recent events (last 24 hours)
	recentEvents, err := api.EventsGetRecent(24)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent events: %v", err)
	}
	
	// Check unacknowledged events
	unackEvents, err := api.EventsGetUnacknowledged()
	if err != nil {
		return nil, fmt.Errorf("failed to get unacknowledged events: %v", err)
	}
	
	// Calculate health metrics
	totalRecent := len(recentEvents)
	totalUnack := len(unackEvents)
	
	// Identify potential issues
	issues := []string{}
	
	if totalUnack > 100 {
		issues = append(issues, fmt.Sprintf("High number of unacknowledged events: %d", totalUnack))
	}
	
	// Check for recent critical events
	recentCritical, _ := api.EventsGetCritical()
	if len(recentCritical) > 0 {
		issues = append(issues, fmt.Sprintf("Recent critical events detected: %d", len(recentCritical)))
	}
	
	// Build health report
	health["status"] = "healthy"
	if len(issues) > 0 {
		health["status"] = "warning"
		health["issues"] = issues
	}
	
	health["metrics"] = map[string]interface{}{
		"recent_events_24h": totalRecent,
		"unacknowledged_events": totalUnack,
		"unacknowledged_percentage": float64(totalUnack) / float64(totalRecent) * 100,
		"recent_critical_events": len(recentCritical),
	}
	
	health["timestamp"] = time.Now().Unix()
	
	return health, nil
}

// EventIsProblem Check if event represents a problem state
func (event *Event) IsProblem() bool {
	return event.Value == 1
}

// EventIsOK Check if event represents an OK state
func (event *Event) IsOK() bool {
	return event.Value == 0
}

// EventIsAcknowledged Check if event is acknowledged
func (event *Event) IsAcknowledged() bool {
	return event.Acknowledged == 1
}

// EventIsCritical Check if event is critical (disaster severity)
func (event *Event) IsCritical() bool {
	return event.Severity == EventSeverityDisaster
}

// EventIsHighSeverity Check if event has high severity (high or disaster)
func (event *Event) IsHighSeverity() bool {
	return event.Severity == EventSeverityHigh || event.Severity == EventSeverityDisaster
}

// EventIsWarning Check if event is warning severity
func (event *Event) IsWarning() bool {
	return event.Severity == EventSeverityWarning
}

// EventAge Calculate event age in seconds
func (event *Event) Age() int {
	return int(time.Now().Unix()) - event.Clock
}

// EventAgeFormatted Get formatted event age (e.g., "2h 30m")
func (event *Event) AgeFormatted() string {
	age := event.Age()
	
	days := age / 86400
	hours := (age % 86400) / 3600
	minutes := (age % 3600) / 60
	
	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	} else {
		return fmt.Sprintf("%dm", minutes)
	}
}