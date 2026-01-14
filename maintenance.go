package zabbix

import (
	"encoding/json"
	"fmt"
	"time"
)

// Maintenance represents a Zabbix maintenance object
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/object
type Maintenance struct {
	MaintenanceID string   `json:"maintenanceid,omitempty"`
	Name         string   `json:"name"`
	MaintenanceType string `json:"maintenance_type,string"`
	ActiveSince  string   `json:"active_since"`
	ActiveTill  string   `json:"active_till"`
	Description  string   `json:"description,omitempty"`
	Tags        []MaintenanceTag `json:"tags,omitempty"`
	Groups      []MaintenanceGroup `json:"groups,omitempty"`
	Hosts       []MaintenanceHost `json:"hosts,omitempty"`
	TimePeriods []MaintenanceTimePeriod `json:"timeperiods,omitempty"`
}

// MaintenanceTag represents a maintenance tag
type MaintenanceTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// MaintenanceGroup represents a host group for maintenance
type MaintenanceGroup struct {
	GroupID string `json:"groupid"`
	Name   string `json:"name,omitempty"`
}

// MaintenanceHost represents a host for maintenance
type MaintenanceHost struct {
	HostID string `json:"hostid"`
	Host   string `json:"host,omitempty"`
}

// MaintenanceTimePeriod represents a time period for maintenance
type MaintenanceTimePeriod struct {
	TimePeriodID string `json:"timeperiodid,omitempty"`
	MaintenanceID string `json:"maintenanceid,omitempty"`
	TimePeriodType string `json:"timeperiod_type,string"`
	EverySecond  string `json:"every_second,omitempty"`
	EveryMinute  string `json:"every_minute,omitempty"`
	EveryHour   string `json:"every_hour,omitempty"`
	EveryDay    string `json:"every_day,omitempty"`
	EveryWeek   string `json:"every_week,omitempty"`
	EveryMonth  string `json:"every_month,omitempty"`
	EveryYear   string `json:"every_year,omitempty"`
	DayOfWeek   string `json:"dayofweek,omitempty"`
	Day         string `json:"day,omitempty"`
	Hour        string `json:"hour,omitempty"`
	Minute      string `json:"minute,omitempty"`
	Date        string `json:"date,omitempty"`
	Period      string `json:"period,omitempty"`
}

// Maintenances represents an array of Maintenance objects
type Maintenances []Maintenance

// MaintenanceGetOptions represents parameters for maintenance.get API call
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/get
type MaintenanceGetOptions struct {
	MaintenanceIDs   []string               `json:"maintenanceids,omitempty"`
	GroupIDs       []string               `json:"groupids,omitempty"`
	HostIDs        []string               `json:"hostids,omitempty"`
	Filter         map[string]interface{} `json:"filter,omitempty"`
	Search         map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string        `json:"searchWildcardsEnabled,omitempty"`
	Output         string               `json:"output,omitempty"`
	SelectGroups    string               `json:"selectGroups,omitempty"`
	SelectHosts     string               `json:"selectHosts,omitempty"`
	SelectTags     string               `json:"selectTags,omitempty"`
	SelectTimePeriods string            `json:"selectTimePeriods,omitempty"`
	SortField      string               `json:"sortfield,omitempty"`
	SortOrder      string               `json:"sortorder,omitempty"`
	Limit          int                  `json:"limit,omitempty"`
}

// MaintenanceCreateOptions represents parameters for maintenance.create API call
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/create
type MaintenanceCreateOptions struct {
	Maintenances    Maintenances         `json:"maintenances"`
	SelectGroups    string              `json:"selectGroups,omitempty"`
	SelectHosts     string              `json:"selectHosts,omitempty"`
	SelectTags     string              `json:"selectTags,omitempty"`
	SelectTimePeriods string          `json:"selectTimePeriods,omitempty"`
}

// MaintenanceUpdateOptions represents parameters for maintenance.update API call
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/update
type MaintenanceUpdateOptions struct {
	Maintenances    Maintenances         `json:"maintenances"`
	SelectGroups    string              `json:"selectGroups,omitempty"`
	SelectHosts     string              `json:"selectHosts,omitempty"`
	SelectTags     string              `json:"selectTags,omitempty"`
	SelectTimePeriods string          `json:"selectTimePeriods,omitempty"`
}

// MaintenanceDeleteOptions represents parameters for maintenance.delete API call
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/delete
type MaintenanceDeleteOptions struct {
	MaintenanceIDs []string `json:"maintenanceids"`
}

// MaintenanceWithDetails represents a maintenance with full details
type MaintenanceWithDetails struct {
	Maintenance
	TimePeriodDetails []MaintenanceTimePeriodDetail `json:"timeperiods,omitempty"`
}

// MaintenanceTimePeriodDetail represents detailed time period information
type MaintenanceTimePeriodDetail struct {
	MaintenanceTimePeriod
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Duration  string `json:"duration,omitempty"`
	Recurrence string `json:"recurrence,omitempty"`
}

// MaintenanceStatistics represents maintenance statistics
type MaintenanceStatistics struct {
	TotalMaintenances    int                  `json:"total_maintenances"`
	ActiveMaintenances   int                  `json:"active_maintenances"`
	ScheduledMaintenances int                 `json:"scheduled_maintenances"`
	CompletedMaintenances int                 `json:"completed_maintenances"`
	ByType              map[string]int        `json:"by_type"`
	ByStatus            map[string]int        `json:"by_status"`
	UpcomingMaintenances []MaintenanceWithDetails `json:"upcoming_maintenances"`
	ActiveHostGroups     []MaintenanceGroup     `json:"active_host_groups"`
	ActiveHosts         []MaintenanceHost      `json:"active_hosts"`
}

// Maintenance constants
const (
	// Maintenance types
	MaintenanceTypeMaintenance = "0"
	MaintenanceTypeNoDataCollection = "1"

	// Time period types
	MaintenanceTimePeriodTypeOnetime = "0"
	MaintenanceTimePeriodTypeDaily = "1"
	MaintenanceTimePeriodTypeWeekly = "2"
	MaintenanceTimePeriodTypeMonthly = "3"
	MaintenanceTimePeriodTypeYearly = "4"

	// Day of week constants
	MaintenanceDayMonday    = "1"
	MaintenanceDayTuesday   = "2"
	MaintenanceDayWednesday = "3"
	MaintenanceDayThursday  = "4"
	MaintenanceDayFriday   = "5"
	MaintenanceDaySaturday = "6"
	MaintenanceDaySunday   = "7"

	// Month constants
	MaintenanceMonthJanuary   = "1"
	MaintenanceMonthFebruary = "2"
	MaintenanceMonthMarch    = "3"
	MaintenanceMonthApril   = "4"
	MaintenanceMonthMay      = "5"
	MaintenanceMonthJune    = "6"
	MaintenanceMonthJuly     = "7"
	MaintenanceMonthAugust   = "8"
	MaintenanceMonthSeptember = "9"
	MaintenanceMonthOctober  = "10"
	MaintenanceMonthNovember = "11"
	MaintenanceMonthDecember = "12"

	// Maintenance status
	MaintenanceStatusActive = "0"
	MaintenanceStatusScheduled = "1"
	MaintenanceStatusExpired = "2"
)

// MaintenancesGet Wrapper for maintenance.get
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/get
func (api *API) MaintenancesGet(options MaintenanceGetOptions) (Maintenances, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.MaintenanceIDs) > 0 {
		params["maintenanceids"] = options.MaintenanceIDs
	}
	if len(options.GroupIDs) > 0 {
		params["groupids"] = options.GroupIDs
	}
	if len(options.HostIDs) > 0 {
		params["hostids"] = options.HostIDs
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
	if options.SelectGroups != "" {
		params["selectGroups"] = options.SelectGroups
	}
	if options.SelectHosts != "" {
		params["selectHosts"] = options.SelectHosts
	}
	if options.SelectTags != "" {
		params["selectTags"] = options.SelectTags
	}
	if options.SelectTimePeriods != "" {
		params["selectTimePeriods"] = options.SelectTimePeriods
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

	var maintenances Maintenances
	err := api.CallWithErrorParse("maintenance.get", params, &maintenances)
	return maintenances, err
}

// MaintenancesGetByID Get maintenances by specific maintenance IDs
func (api *API) MaintenancesGetByID(maintenanceIDs []string) (Maintenances, error) {
	options := MaintenanceGetOptions{
		MaintenanceIDs: maintenanceIDs,
		Output: "extend",
	}
	return api.MaintenancesGet(options)
}

// MaintenancesGetByName Get maintenances by name
func (api *API) MaintenancesGetByName(name string) (Maintenances, error) {
	options := MaintenanceGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.MaintenancesGet(options)
}

// MaintenancesGetActive Get active maintenances
func (api *API) MaintenancesGetActive() (Maintenances, error) {
	options := MaintenanceGetOptions{
		Filter: map[string]interface{}{
			"active_since": fmt.Sprintf("%d", int(time.Now().Unix())),
		},
		Output: "extend",
	}
	return api.MaintenancesGet(options)
}

// MaintenancesGetScheduled Get scheduled maintenances
func (api *API) MaintenancesGetScheduled() (Maintenances, error) {
	now := int(time.Now().Unix())
	options := MaintenanceGetOptions{
		Filter: map[string]interface{}{
			"active_since": fmt.Sprintf("%d", now),
			"active_till": fmt.Sprintf("%d", now+1), // Future start
		},
		Output: "extend",
	}
	return api.MaintenancesGet(options)
}

// MaintenancesGetByGroup Get maintenances by host group
func (api *API) MaintenancesGetByGroup(groupID string) (Maintenances, error) {
	options := MaintenanceGetOptions{
		GroupIDs: []string{groupID},
		Output: "extend",
	}
	return api.MaintenancesGet(options)
}

// MaintenancesGetByHost Get maintenances by host
func (api *API) MaintenancesGetByHost(hostID string) (Maintenances, error) {
	options := MaintenanceGetOptions{
		HostIDs: []string{hostID},
		Output: "extend",
	}
	return api.MaintenancesGet(options)
}

// MaintenancesGetWithTimePeriods Get maintenances with their time periods
func (api *API) MaintenancesGetWithTimePeriods(options MaintenanceGetOptions) (Maintenances, error) {
	options.SelectTimePeriods = "extend"
	return api.MaintenancesGet(options)
}

// MaintenancesGetWithGroups Get maintenances with their host groups
func (api *API) MaintenancesGetWithGroups(options MaintenanceGetOptions) (Maintenances, error) {
	options.SelectGroups = "extend"
	return api.MaintenancesGet(options)
}

// MaintenancesGetWithHosts Get maintenances with their hosts
func (api *API) MaintenancesGetWithHosts(options MaintenanceGetOptions) (Maintenances, error) {
	options.SelectHosts = "extend"
	return api.MaintenancesGet(options)
}

// MaintenancesGetWithTags Get maintenances with their tags
func (api *API) MaintenancesGetWithTags(options MaintenanceGetOptions) (Maintenances, error) {
	options.SelectTags = "extend"
	return api.MaintenancesGet(options)
}

// MaintenanceGetByID Get maintenance by ID (exactly one match required)
func (api *API) MaintenanceGetByID(maintenanceID string) (*Maintenance, error) {
	maintenances, err := api.MaintenancesGetByID([]string{maintenanceID})
	if err != nil {
		return nil, err
	}

	if len(maintenances) == 1 {
		return &maintenances[0], nil
	} else if len(maintenances) == 0 {
		return nil, fmt.Errorf("Maintenance not found: %s", maintenanceID)
	} else {
		return nil, fmt.Errorf("Multiple maintenances found with ID: %s", maintenanceID)
	}
}

// MaintenancesCreate Wrapper for maintenance.create
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/create
func (api *API) MaintenancesCreate(maintenances Maintenances) (result []string, err error) {
	options := MaintenanceCreateOptions{
		Maintenances: maintenances,
	}

	response, err := api.CallWithError("maintenance.create", options)
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
			if maintenanceMap, ok := item.(map[string]interface{}); ok {
				if maintenanceid, exists := maintenanceMap["maintenanceids"]; exists {
					if idArray, ok := maintenanceid.([]interface{}); ok && len(idArray) > 0 {
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

// MaintenanceCreateSingle Create a single maintenance
func (api *API) MaintenanceCreateSingle(maintenance Maintenance) (maintenanceID string, err error) {
	maintenances := Maintenances{maintenance}
	result, err := api.MaintenancesCreate(maintenances)
	if len(result) > 0 {
		maintenanceID = result[0]
	}
	return
}

// MaintenancesUpdate Wrapper for maintenance.update
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/update
func (api *API) MaintenancesUpdate(maintenances Maintenances) (result []string, err error) {
	options := MaintenanceUpdateOptions{
		Maintenances: maintenances,
	}

	response, err := api.CallWithError("maintenance.update", options)
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
			if maintenanceMap, ok := item.(map[string]interface{}); ok {
				if maintenanceid, exists := maintenanceMap["maintenanceids"]; exists {
					if idArray, ok := maintenanceid.([]interface{}); ok && len(idArray) > 0 {
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

// MaintenanceUpdateSingle Update a single maintenance
func (api *API) MaintenanceUpdateSingle(maintenance Maintenance) (maintenanceID string, err error) {
	maintenances := Maintenances{maintenance}
	result, err := api.MaintenancesUpdate(maintenances)
	if len(result) > 0 {
		maintenanceID = result[0]
	}
	return
}

// MaintenancesDelete Wrapper for maintenance.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/maintenance/delete
func (api *API) MaintenancesDelete(maintenances Maintenances) (result []string, err error) {
	maintenanceIDs := make([]string, len(maintenances))
	for i, maintenance := range maintenances {
		maintenanceIDs[i] = maintenance.MaintenanceID
	}
	
	return api.MaintenancesDeleteByIDs(maintenanceIDs)
}

// MaintenancesDeleteByIDs Wrapper for maintenance.delete with IDs
func (api *API) MaintenancesDeleteByIDs(maintenanceIDs []string) (result []string, err error) {
	options := MaintenanceDeleteOptions{
		MaintenanceIDs: maintenanceIDs,
	}

	response, err := api.CallWithError("maintenance.delete", options)
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

// MaintenanceDeleteSingle Delete a single maintenance
func (api *API) MaintenanceDeleteSingle(maintenanceID string) (err error) {
	_, err = api.MaintenancesDeleteByIDs([]string{maintenanceID})
	return
}

// MaintenancesGetStatistics Get statistics about maintenances
func (api *API) MaintenancesGetStatistics() (MaintenanceStatistics, error) {
	// Get all maintenances
	allMaintenances, err := api.MaintenancesGet(MaintenanceGetOptions{Output: "extend"})
	if err != nil {
		return MaintenanceStatistics{}, err
	}

	stats := MaintenanceStatistics{
		TotalMaintenances: len(allMaintenances),
		ByType: make(map[string]int),
		ByStatus: make(map[string]int),
	}

	now := int(time.Now().Unix())
	activeMaintenances := 0
	scheduledMaintenances := 0
	completedMaintenances := 0

	for _, maintenance := range allMaintenances {
		// Count by type
		stats.ByType[maintenance.MaintenanceType]++
		
		// Determine status and count
		if maintenance.IsActive(now) {
			stats.ByStatus["active"]++
			activeMaintenances++
		} else if maintenance.IsScheduled(now) {
			stats.ByStatus["scheduled"]++
			scheduledMaintenances++
		} else if maintenance.IsExpired(now) {
			stats.ByStatus["expired"]++
			completedMaintenances++
		}
	}

	stats.ActiveMaintenances = activeMaintenances
	stats.ScheduledMaintenances = scheduledMaintenances
	stats.CompletedMaintenances = completedMaintenances

	// Get upcoming maintenances (next 30 days)
	upcomingStart := now
	upcomingEnd := now + 30*24*3600 // 30 days
	upcomingMaintenances, err := api.MaintenancesGet(MaintenanceGetOptions{
		Filter: map[string]interface{}{
			"active_since": fmt.Sprintf("%d", upcomingStart),
			"active_till": fmt.Sprintf("%d", upcomingEnd),
		},
		Output: "extend",
		SelectGroups: "extend",
		SelectHosts: "extend",
		SelectTimePeriods: "extend",
	})
	
	if err == nil {
		// Convert to detailed view (simplified)
		for _, maintenance := range upcomingMaintenances {
			stats.UpcomingMaintenances = append(stats.UpcomingMaintenances, MaintenanceWithDetails{
				Maintenance: maintenance,
			})
		}
	}

	// Get active host groups
	activeGroups, err := api.MaintenancesGetActive()
	if err == nil {
		for _, maintenance := range activeGroups {
			if maintenance.Groups != nil {
				stats.ActiveHostGroups = append(stats.ActiveHostGroups, maintenance.Groups...)
			}
		}
	}

	// Get active hosts
	activeHosts, err := api.MaintenancesGetActive()
	if err == nil {
		for _, maintenance := range activeHosts {
			if maintenance.Hosts != nil {
				stats.ActiveHosts = append(stats.ActiveHosts, maintenance.Hosts...)
			}
		}
	}

	return stats, nil
}

// MaintenanceValidate Validate a maintenance configuration
func (api *API) MaintenanceValidate(maintenance Maintenance) (validationErrors []string) {
	validationErrors = []string{}
	
	// Check required fields
	if maintenance.Name == "" {
		validationErrors = append(validationErrors, "Maintenance name is required")
	}
	
	if maintenance.ActiveSince == "" {
		validationErrors = append(validationErrors, "Active since time is required")
	}
	
	if maintenance.ActiveTill == "" {
		validationErrors = append(validationErrors, "Active till time is required")
	}
	
	// Validate maintenance type
	if maintenance.MaintenanceType != "" && 
		maintenance.MaintenanceType != MaintenanceTypeMaintenance && 
		maintenance.MaintenanceType != MaintenanceTypeNoDataCollection {
		validationErrors = append(validationErrors, fmt.Sprintf("Invalid maintenance type: %s", maintenance.MaintenanceType))
	}
	
	// Validate time periods if present
	for i, timePeriod := range maintenance.TimePeriods {
		if timePeriod.TimePeriodType == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Time period %d: Type is required", i))
		}
		
		if timePeriod.Period == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Time period %d: Period is required", i))
		}
	}
	
	// Validate that either groups or hosts are specified
	if len(maintenance.Groups) == 0 && len(maintenance.Hosts) == 0 {
		validationErrors = append(validationErrors, "Either host groups or hosts must be specified")
	}
	
	return validationErrors
}

// MaintenanceIsActive Check if maintenance is currently active
func (maintenance *Maintenance) IsActive(now int) bool {
	// Parse active_since and active_till times
	since, _ := time.Parse("2006-01-02 15:04:05", maintenance.ActiveSince)
	till, _ := time.Parse("2006-01-02 15:04:05", maintenance.ActiveTill)
	
	return now >= int(since.Unix()) && now <= int(till.Unix())
}

// MaintenanceIsScheduled Check if maintenance is scheduled for future
func (maintenance *Maintenance) IsScheduled(now int) bool {
	// Parse active_since time
	since, _ := time.Parse("2006-01-02 15:04:05", maintenance.ActiveSince)
	
	return now < int(since.Unix())
}

// MaintenanceIsExpired Check if maintenance has expired
func (maintenance *Maintenance) IsExpired(now int) bool {
	// Parse active_till time
	till, _ := time.Parse("2006-01-02 15:04:05", maintenance.ActiveTill)
	
	return now > int(till.Unix())
}

// MaintenanceIsNoDataCollection Check if maintenance is of no data collection type
func (maintenance *Maintenance) IsNoDataCollection() bool {
	return maintenance.MaintenanceType == MaintenanceTypeNoDataCollection
}

// MaintenanceIsRegularMaintenance Check if maintenance is of regular maintenance type
func (maintenance *Maintenance) IsRegularMaintenance() bool {
	return maintenance.MaintenanceType == MaintenanceTypeMaintenance
}

// MaintenanceHasTimePeriods Check if maintenance has time periods
func (maintenance *Maintenance) HasTimePeriods() bool {
	return len(maintenance.TimePeriods) > 0
}

// MaintenanceGetTimePeriodCount Get number of time periods
func (maintenance *Maintenance) GetTimePeriodCount() int {
	return len(maintenance.TimePeriods)
}

// MaintenanceHasGroups Check if maintenance has host groups
func (maintenance *Maintenance) HasGroups() bool {
	return len(maintenance.Groups) > 0
}

// MaintenanceHasHosts Check if maintenance has hosts
func (maintenance *Maintenance) HasHosts() bool {
	return len(maintenance.Hosts) > 0
}

// MaintenanceGetGroupCount Get number of host groups
func (maintenance *Maintenance) GetGroupCount() int {
	return len(maintenance.Groups)
}

// MaintenanceGetHostCount Get number of hosts
func (maintenance *Maintenance) GetHostCount() int {
	return len(maintenance.Hosts)
}

// MaintenanceGetTags Get maintenance tags
func (maintenance *Maintenance) GetTags() []MaintenanceTag {
	return maintenance.Tags
}

// MaintenanceAddTag Add a tag to maintenance
func (maintenance *Maintenance) AddTag(tag MaintenanceTag) {
	maintenance.Tags = append(maintenance.Tags, tag)
}

// MaintenanceAddGroup Add a host group to maintenance
func (maintenance *Maintenance) AddGroup(group MaintenanceGroup) {
	maintenance.Groups = append(maintenance.Groups, group)
}

// MaintenanceAddHost Add a host to maintenance
func (maintenance *Maintenance) AddHost(host MaintenanceHost) {
	maintenance.Hosts = append(maintenance.Hosts, host)
}

// MaintenanceAddTimePeriod Add a time period to maintenance
func (maintenance *Maintenance) AddTimePeriod(timePeriod MaintenanceTimePeriod) {
	maintenance.TimePeriods = append(maintenance.TimePeriods, timePeriod)
}

// CreateSimpleMaintenance Create a simple maintenance with basic configuration
func (api *API) CreateSimpleMaintenance(name, description string, activeSince, activeTill string, groupIDs, hostIDs []string) (string, error) {
	maintenance := Maintenance{
		Name:            name,
		Description:     description,
		MaintenanceType: MaintenanceTypeMaintenance,
		ActiveSince:    activeSince,
		ActiveTill:     activeTill,
	}
	
	// Add groups
	for _, groupID := range groupIDs {
		maintenance.AddGroup(MaintenanceGroup{GroupID: groupID})
	}
	
	// Add hosts
	for _, hostID := range hostIDs {
		maintenance.AddHost(MaintenanceHost{HostID: hostID})
	}
	
	return api.MaintenanceCreateSingle(maintenance)
}

// CreateOnetimeMaintenance Create a one-time maintenance
func (api *API) CreateOnetimeMaintenance(name, description string, startTime, duration int64, groupIDs, hostIDs []string) (string, error) {
	start := time.Unix(startTime, 0).Format("2006-01-02 15:04:05")
	end := time.Unix(startTime+duration, 0).Format("2006-01-02 15:04:05")
	
	maintenance := Maintenance{
		Name:            name,
		Description:     description,
		MaintenanceType: MaintenanceTypeMaintenance,
		ActiveSince:    start,
		ActiveTill:     end,
		TimePeriods: []MaintenanceTimePeriod{
			{
				TimePeriodType: MaintenanceTimePeriodTypeOnetime,
				Period:        fmt.Sprintf("%d", duration), // Duration in seconds
			},
		},
	}
	
	// Add groups
	for _, groupID := range groupIDs {
		maintenance.AddGroup(MaintenanceGroup{GroupID: groupID})
	}
	
	// Add hosts
	for _, hostID := range hostIDs {
		maintenance.AddHost(MaintenanceHost{HostID: hostID})
	}
	
	return api.MaintenanceCreateSingle(maintenance)
}

// CreateRecurringMaintenance Create a recurring maintenance
func (api *API) CreateRecurringMaintenance(name, description string, startTime, duration int64, groupIDs, hostIDs []string, timePeriodType string) (string, error) {
	start := time.Unix(startTime, 0).Format("2006-01-02 15:04:05")
	end := time.Unix(startTime+duration, 0).Format("2006-01-02 15:04:05")
	
	maintenance := Maintenance{
		Name:            name,
		Description:     description,
		MaintenanceType: MaintenanceTypeMaintenance,
		ActiveSince:    start,
		ActiveTill:     end,
		TimePeriods: []MaintenanceTimePeriod{
			{
				TimePeriodType: timePeriodType,
				Period:        fmt.Sprintf("%d", duration), // Duration in seconds
			},
		},
	}
	
	// Add groups
	for _, groupID := range groupIDs {
		maintenance.AddGroup(MaintenanceGroup{GroupID: groupID})
	}
	
	// Add hosts
	for _, hostID := range hostIDs {
		maintenance.AddHost(MaintenanceHost{HostID: hostID})
	}
	
	return api.MaintenanceCreateSingle(maintenance)
}

// MaintenanceGetUpcoming Get upcoming maintenances
func (api *API) MaintenanceGetUpcoming(days int) (Maintenances, error) {
	now := int(time.Now().Unix())
	future := now + (days * 24 * 3600) // days in seconds
	
	options := MaintenanceGetOptions{
		Filter: map[string]interface{}{
			"active_since": fmt.Sprintf("%d", now),
			"active_till": fmt.Sprintf("%d", future),
		},
		Output: "extend",
		SelectGroups: "extend",
		SelectHosts: "extend",
		SelectTimePeriods: "extend",
	}
	
	return api.MaintenancesGet(options)
}

// MaintenanceGetCurrent Get currently active maintenances
func (api *API) MaintenanceGetCurrent() (Maintenances, error) {
	return api.MaintenancesGetActive()
}

// MaintenanceCheckMaintenance Check if hosts or groups are under maintenance
func (api *API) MaintenanceCheckMaintenance(hostIDs, groupIDs []string) (map[string]bool, error) {
	result := make(map[string]bool)
	
	// Check host maintenances
	for _, hostID := range hostIDs {
		maintenances, err := api.MaintenancesGetByHost(hostID)
		if err == nil {
			now := int(time.Now().Unix())
			for _, maintenance := range maintenances {
				if maintenance.IsActive(now) {
					result[fmt.Sprintf("host:%s", hostID)] = true
					break
				}
			}
		}
	}
	
	// Check group maintenances
	for _, groupID := range groupIDs {
		maintenances, err := api.MaintenancesGetByGroup(groupID)
		if err == nil {
			now := int(time.Now().Unix())
			for _, maintenance := range maintenances {
				if maintenance.IsActive(now) {
					result[fmt.Sprintf("group:%s", groupID)] = true
					break
				}
			}
		}
	}
	
	return result, nil
}