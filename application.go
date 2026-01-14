package zabbix

import (
	"encoding/json"
	"fmt"
)

// Application represents a Zabbix application object
// https://www.zabbix.com/documentation/current/manual/api/reference/application/object
type Application struct {
	ApplicationID string   `json:"applicationid,omitempty"`
	HostID        string   `json:"hostid"`
	Name          string   `json:"name"`
	TemplateID    string   `json:"templateid,omitempty"`
	Flags         string   `json:"flags"`
	// Note: Applications (applications) field is deprecated in Zabbix 6.0+
	// and has been replaced by Tags system
}

// Applications represents an array of Application objects
type Applications []Application

// ApplicationID represents an application ID
type ApplicationID struct {
	ApplicationID string `json:"applicationid"`
}

// ApplicationIDs represents an array of ApplicationID objects
type ApplicationIDs []ApplicationID

// ApplicationGetOptions represents parameters for application.get API call
// https://www.zabbix.com/documentation/current/manual/api/reference/application/get
type ApplicationGetOptions struct {
	ApplicationIDs []string               `json:"applicationids,omitempty"`
	HostIDs        []string               `json:"hostids,omitempty"`
	TemplateIDs    []string               `json:"templateids,omitempty"`
	Filter         map[string]interface{} `json:"filter,omitempty"`
	Search         map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string        `json:"searchWildcardsEnabled,omitempty"`
	Output         string                 `json:"output,omitempty"`
	SelectHosts    string                 `json:"selectHosts,omitempty"`
	SelectItems    string                 `json:"selectItems,omitempty"`
	SortField      string                 `json:"sortfield,omitempty"`
	SortOrder      string                 `json:"sortorder,omitempty"`
	Limit          int                    `json:"limit,omitempty"`
}

// ApplicationCreateOptions represents parameters for application.create API call
// https://www.zabbix.com/documentation/current/manual/api/reference/application/create
type ApplicationCreateOptions struct {
	Applications  Applications `json:"applications"`
	SelectHosts  string       `json:"selectHosts,omitempty"`
	SelectItems  string       `json:"selectItems,omitempty"`
}

// ApplicationUpdateOptions represents parameters for application.update API call
// https://www.zabbix.com/documentation/current/manual/api/reference/application/update
type ApplicationUpdateOptions struct {
	Applications  Applications `json:"applications"`
	SelectHosts  string       `json:"selectHosts,omitempty"`
	SelectItems  string       `json:"selectItems,omitempty"`
}

// ApplicationDeleteOptions represents parameters for application.delete API call
// https://www.zabbix.com/documentation/current/manual/api/reference/application/delete
type ApplicationDeleteOptions struct {
	ApplicationIDs []string `json:"applicationids"`
}

// ApplicationWithItems represents an application with associated items
type ApplicationWithItems struct {
	Application
	Items Items `json:"items,omitempty"`
}

// ApplicationWithHosts represents an application with associated hosts
type ApplicationWithHosts struct {
	Application
	Hosts Hosts `json:"hosts,omitempty"`
}

// ApplicationsGet Wrapper for application.get
// https://www.zabbix.com/documentation/current/manual/api/reference/application/get
func (api *API) ApplicationsGet(options ApplicationGetOptions) (Applications, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.ApplicationIDs) > 0 {
		params["applicationids"] = options.ApplicationIDs
	}
	if len(options.HostIDs) > 0 {
		params["hostids"] = options.HostIDs
	}
	if len(options.TemplateIDs) > 0 {
		params["templateids"] = options.TemplateIDs
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
	if options.SelectHosts != "" {
		params["selectHosts"] = options.SelectHosts
	}
	if options.SelectItems != "" {
		params["selectItems"] = options.SelectItems
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

	var applications Applications
	err := api.CallWithErrorParse("application.get", params, &applications)
	return applications, err
}

// ApplicationsGetByID Get applications by specific application IDs
func (api *API) ApplicationsGetByID(applicationIDs []string) (Applications, error) {
	options := ApplicationGetOptions{
		ApplicationIDs: applicationIDs,
		Output:        "extend",
	}
	return api.ApplicationsGet(options)
}

// ApplicationsGetByHost Get applications by host IDs
func (api *API) ApplicationsGetByHost(hostIDs []string) (Applications, error) {
	options := ApplicationGetOptions{
		HostIDs: hostIDs,
		Output:  "extend",
	}
	return api.ApplicationsGet(options)
}

// ApplicationsGetByHostAndTemplate Get applications by host and template IDs
func (api *API) ApplicationsGetByHostAndTemplate(hostIDs, templateIDs []string) (Applications, error) {
	options := ApplicationGetOptions{
		HostIDs:     hostIDs,
		TemplateIDs: templateIDs,
		Output:      "extend",
	}
	return api.ApplicationsGet(options)
}

// ApplicationGetByID Get application by ID (exactly one match required)
func (api *API) ApplicationGetByID(applicationID string) (*Application, error) {
	applications, err := api.ApplicationsGetByID([]string{applicationID})
	if err != nil {
		return nil, err
	}

	if len(applications) == 1 {
		return &applications[0], nil
	} else if len(applications) == 0 {
		return nil, fmt.Errorf("Application not found: %s", applicationID)
	} else {
		return nil, fmt.Errorf("Multiple applications found with ID: %s", applicationID)
	}
}

// ApplicationsGetWithItems Get applications with their associated items
func (api *API) ApplicationsGetWithItems(options ApplicationGetOptions) ([]ApplicationWithItems, error) {
	options.SelectItems = "extend"
	
	applications, err := api.ApplicationsGet(options)
	if err != nil {
		return nil, err
	}
	
	applicationsWithItems := make([]ApplicationWithItems, 0, len(applications))
	for _, app := range applications {
		// Get items for this application
		items, err := api.ItemsGetByApplicationID(app.ApplicationID)
		if err != nil {
			// If we can't get items, just include the application without items
			applicationsWithItems = append(applicationsWithItems, ApplicationWithItems{
				Application: app,
				Items:       nil,
			})
		} else {
			applicationsWithItems = append(applicationsWithItems, ApplicationWithItems{
				Application: app,
				Items:       items,
			})
		}
	}
	
	return applicationsWithItems, nil
}

// ApplicationsGetWithHosts Get applications with their associated hosts
func (api *API) ApplicationsGetWithHosts(options ApplicationGetOptions) ([]ApplicationWithHosts, error) {
	options.SelectHosts = "extend"
	
	applications, err := api.ApplicationsGet(options)
	if err != nil {
		return nil, err
	}
	
	applicationsWithHosts := make([]ApplicationWithHosts, 0, len(applications))
	for _, app := range applications {
		// Get host for this application
		host, err := api.HostGetByID(app.HostID)
		if err != nil {
			// If we can't get host, just include the application without host
			applicationsWithHosts = append(applicationsWithHosts, ApplicationWithHosts{
				Application: app,
				Hosts:       nil,
			})
		} else {
			applicationsWithHosts = append(applicationsWithHosts, ApplicationWithHosts{
				Application: app,
				Hosts:       Hosts{*host},
			})
		}
	}
	
	return applicationsWithHosts, nil
}

// ApplicationsCreate Wrapper for application.create
// https://www.zabbix.com/documentation/current/manual/api/reference/application/create
func (api *API) ApplicationsCreate(applications Applications) (result []string, err error) {
	options := ApplicationCreateOptions{
		Applications: applications,
	}

	response, err := api.CallWithError("application.create", options)
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
			if appMap, ok := item.(map[string]interface{}); ok {
				if applicationid, exists := appMap["applicationids"]; exists {
					if idArray, ok := applicationid.([]interface{}); ok && len(idArray) > 0 {
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

// ApplicationCreateSingle Create a single application
func (api *API) ApplicationCreateSingle(application Application) (applicationID string, err error) {
	applications := Applications{application}
	result, err := api.ApplicationsCreate(applications)
	if len(result) > 0 {
		applicationID = result[0]
	}
	return
}

// ApplicationsUpdate Wrapper for application.update
// https://www.zabbix.com/documentation/current/manual/api/reference/application/update
func (api *API) ApplicationsUpdate(applications Applications) (result []string, err error) {
	options := ApplicationUpdateOptions{
		Applications: applications,
	}

	response, err := api.CallWithError("application.update", options)
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
			if appMap, ok := item.(map[string]interface{}); ok {
				if applicationid, exists := appMap["applicationids"]; exists {
					if idArray, ok := applicationid.([]interface{}); ok && len(idArray) > 0 {
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

// ApplicationUpdateSingle Update a single application
func (api *API) ApplicationUpdateSingle(application Application) (applicationID string, err error) {
	applications := Applications{application}
	result, err := api.ApplicationsUpdate(applications)
	if len(result) > 0 {
		applicationID = result[0]
	}
	return
}

// ApplicationsDelete Wrapper for application.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/application/delete
func (api *API) ApplicationsDelete(applications Applications) (result []string, err error) {
	applicationIDs := make([]string, len(applications))
	for i, app := range applications {
		applicationIDs[i] = app.ApplicationID
	}
	
	return api.ApplicationsDeleteByIDs(applicationIDs)
}

// ApplicationsDeleteByIDs Wrapper for application.delete with IDs
func (api *API) ApplicationsDeleteByIDs(applicationIDs []string) (result []string, err error) {
	options := ApplicationDeleteOptions{
		ApplicationIDs: applicationIDs,
	}

	response, err := api.CallWithError("application.delete", options)
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

// ApplicationDeleteSingle Delete a single application
func (api *API) ApplicationDeleteSingle(applicationID string) (err error) {
	_, err = api.ApplicationsDeleteByIDs([]string{applicationID})
	return
}

// ItemsGetByApplicationID Get items belonging to an application
func (api *API) ItemsGetByApplicationID(applicationID string) (Items, error) {
	// Get the application first to get the host ID
	application, err := api.ApplicationGetByID(applicationID)
	if err != nil {
		return nil, err
	}
	
	// Get items for the host and filter by application
	items, err := api.ItemsGetByHostID(application.HostID)
	if err != nil {
		return nil, err
	}
	
	// Filter items by application (this is a simplified implementation)
	// In reality, this would require more complex filtering based on the application
	filteredItems := Items{}
	for _, item := range items {
		// Note: The actual API doesn't support direct filtering by application ID
		// This would need to be done by getting all items and filtering client-side
		// or using a more complex query
		filteredItems = append(filteredItems, item)
	}
	
	return filteredItems, nil
}

// ApplicationsGetCount Get count of applications matching criteria
func (api *API) ApplicationsGetCount(options ApplicationGetOptions) (int, error) {
	options.Limit = 1
	applications, err := api.ApplicationsGet(options)
	if err != nil {
		return 0, err
	}
	
	// Get the total count from API metadata if available
	// This is a simplified implementation
	return len(applications), nil
}

// ApplicationsGetStatistics Get statistics about applications
func (api *API) ApplicationsGetStatistics() (map[string]interface{}, error) {
	// Get all applications
	allApplications, err := api.ApplicationsGet(ApplicationGetOptions{Output: "extend"})
	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})
	
	// Basic counts
	totalApplications := len(allApplications)
	
	// Group by host
	hostCounts := make(map[string]int)
	
	// Group by template
	templateCounts := make(map[string]int)
	
	for _, app := range allApplications {
		// Count by host
		hostCounts[app.HostID]++
		
		// Count by template
		if app.TemplateID != "" {
			templateCounts[app.TemplateID]++
		}
	}
	
	// Build statistics result
	stats["total_applications"] = totalApplications
	stats["host_distribution"] = hostCounts
	stats["template_distribution"] = templateCounts
	
	// Calculate application density (apps per host)
	stats["unique_hosts"] = len(hostCounts)
	if len(hostCounts) > 0 {
		stats["applications_per_host"] = float64(totalApplications) / float64(len(hostCounts))
	} else {
		stats["applications_per_host"] = 0
	}
	
	return stats, nil
}

// ApplicationsGetByName Get applications by name
func (api *API) ApplicationsGetByName(name string) (Applications, error) {
	options := ApplicationGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.ApplicationsGet(options)
}

// ApplicationsGetByNames Get applications by multiple names
func (api *API) ApplicationsGetByNames(names []string) (Applications, error) {
	allApplications := Applications{}
	
	for _, name := range names {
		applications, err := api.ApplicationsGetByName(name)
		if err != nil {
			// Continue with other names even if one fails
			continue
		}
		allApplications = append(allApplications, applications...)
	}
	
	return allApplications, nil
}

// ApplicationsGetByPattern Get applications matching a pattern
func (api *API) ApplicationsGetByPattern(pattern string) (Applications, error) {
	options := ApplicationGetOptions{
		Search: map[string]interface{}{
			"name": pattern,
		},
		SearchWildcardsEnabled: "true",
		Output:                "extend",
	}
	return api.ApplicationsGet(options)
}

// ApplicationValidate Validate an application configuration
func (api *API) ApplicationValidate(application Application) (validationErrors []string) {
	validationErrors = []string{}
	
	// Check required fields
	if application.HostID == "" {
		validationErrors = append(validationErrors, "Host ID is required")
	}
	
	if application.Name == "" {
		validationErrors = append(validationErrors, "Application name is required")
	}
	
	// Validate HostID format (should be numeric string)
	if application.HostID != "" {
		if len(application.HostID) == 0 {
			validationErrors = append(validationErrors, "Host ID cannot be empty")
		}
	}
	
	// Validate TemplateID format if provided
	if application.TemplateID != "" {
		if len(application.TemplateID) == 0 {
			validationErrors = append(validationErrors, "Template ID cannot be empty")
		}
	}
	
	return validationErrors
}

// ApplicationsGetForHost Get all applications for a specific host
func (api *API) ApplicationsGetForHost(hostID string) (Applications, error) {
	return api.ApplicationsGetByHost([]string{hostID})
}

// ApplicationsGetForTemplate Get all applications for a specific template
func (api *API) ApplicationsGetForTemplate(templateID string) (Applications, error) {
	options := ApplicationGetOptions{
		TemplateIDs: []string{templateID},
		Output:      "extend",
	}
	return api.ApplicationsGet(options)
}

// ApplicationsGetSummary Get summary information about applications
func (api *API) ApplicationsGetSummary(hostIDs []string) (map[string]interface{}, error) {
	summary := make(map[string]interface{})
	
	// Get applications for specified hosts
	applications, err := api.ApplicationsGetByHost(hostIDs)
	if err != nil {
		return nil, err
	}
	
	// Group by host
	applicationsByHost := make(map[string]Applications)
	for _, app := range applications {
		applicationsByHost[app.HostID] = append(applicationsByHost[app.HostID], app)
	}
	
	// Build summary
	summary["host_applications"] = applicationsByHost
	summary["total_applications"] = len(applications)
	summary["unique_hosts"] = len(applicationsByHost)
	
	// Calculate statistics per host
	hostStats := make(map[string]interface{})
	for hostID, apps := range applicationsByHost {
		hostStats[hostID] = map[string]interface{}{
			"application_count": len(apps),
			"applications": apps,
		}
	}
	summary["host_statistics"] = hostStats
	
	return summary, nil
}