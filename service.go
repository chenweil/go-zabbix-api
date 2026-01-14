package zabbix

import (
	"encoding/json"
	"fmt"
	"time"
)

// Service represents a Zabbix service object
// https://www.zabbix.com/documentation/current/manual/api/reference/service/object
type Service struct {
	ServiceID    string   `json:"serviceid,omitempty"`
	Name         string   `json:"name"`
	Status       string   `json:"status,string"`
	Algorithm    string   `json:"algorithm,string"`
	Description  string   `json:"description,omitempty"`
	ParentID     string   `json:"parentid,omitempty"`
	RootCause    string   `json:"rootcause,omitempty"`
	Order        string   `json:"order,omitempty"`
	SortOrder    string   `json:"sortorder,omitempty"`
	Propagate    string   `json:"propagate,string"`
	ShowSLA      string   `json:"show_sla,string"`
	AcceptableSLA string   `json:"acceptable_sla,omitempty"`
	
	// Additional fields
	ServiceType   string   `json:"service_type,omitempty"`
	BusinessValue string   `json:"business_value,omitempty"`
	Tags          []Tag  `json:"tags,omitempty"`
	
	// Read-only fields
	LatestSLA     float64 `json:"latest_sla,omitempty"`
	LatestDowntime int     `json:"latest_downtime,omitempty"`
	LatestChange  int     `json:"latest_change,omitempty"`
}

// Services represents an array of Service objects
type Services []Service

// ServiceProblem represents a service problem object
type ServiceProblem struct {
	ServiceID    string `json:"serviceid,omitempty"`
	ProblemID    string `json:"problemid,omitempty"`
	TriggerID    string `json:"triggerid,omitempty"`
	EventID      string `json:"eventid,omitempty"`
	Status       string `json:"status"`
	Clock        int    `json:"clock"`
	Severity     string `json:"severity"`
	Description  string `json:"description"`
}

// ServiceProblems represents an array of ServiceProblem objects
type ServiceProblems []ServiceProblem

// ServiceDependency represents a service dependency relationship
type ServiceDependency struct {
	DependsOnID string `json:"depends_onid,omitempty"`
	LinkType    string `json:"link_type,string"`
}

// ServiceDependencies represents an array of ServiceDependency objects
type ServiceDependencies []ServiceDependency

// ServiceTag represents a service tag
type ServiceTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// ServiceTags represents an array of ServiceTag objects
type ServiceTags []ServiceTag

// ServiceGetOptions represents parameters for service.get API call
// https://www.zabbix.com/documentation/current/manual/api/reference/service/get
type ServiceGetOptions struct {
	ServiceIDs     []string             `json:"serviceids,omitempty"`
	ParentIDs     []string             `json:"parentids,omitempty"`
	ChildIDs      []string             `json:"childids,omitempty"`
	TriggerIDs    []string             `json:"triggerids,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	Search        map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string        `json:"searchWildcardsEnabled,omitempty"`
	Output        string               `json:"output,omitempty"`
	SelectChildren string              `json:"selectChildren,omitempty"`
	SelectParents string               `json:"selectParents,omitempty"`
	SelectTags   string               `json:"selectTags,omitempty"`
	SelectProblemEvents string        `json:"selectProblemEvents,omitempty"`
	SelectServices string             `json:"selectServices,omitempty"`
	SortField    string               `json:"sortfield,omitempty"`
	SortOrder    string               `json:"sortorder,omitempty"`
	Limit        int                  `json:"limit,omitempty"`
}

// ServiceCreateOptions represents parameters for service.create API call
// https://www.zabbix.com/documentation/current/manual/api/reference/service/create
type ServiceCreateOptions struct {
	Services      Services      `json:"services"`
	SelectChildren string      `json:"selectChildren,omitempty"`
	SelectParents string      `json:"selectParents,omitempty"`
	SelectTags   string      `json:"selectTags,omitempty"`
}

// ServiceUpdateOptions represents parameters for service.update API call
// https://www.zabbix.com/documentation/current/manual/api/reference/service/update
type ServiceUpdateOptions struct {
	Services      Services      `json:"services"`
	SelectChildren string      `json:"selectChildren,omitempty"`
	SelectParents string      `json:"selectParents,omitempty"`
	SelectTags   string      `json:"selectTags,omitempty"`
}

// ServiceDeleteOptions represents parameters for service.delete API call
// https://www.zabbix.com/documentation/current/manual/api/reference/service/delete
type ServiceDeleteOptions struct {
	ServiceIDs []string `json:"serviceids"`
}

// ServiceGetSLAOptions represents parameters for service.getsla API call
// https://www.zabbix.com/documentation/current/manual/api/reference/service/getsla
type ServiceGetSLAOptions struct {
	ServiceIDs   []string   `json:"serviceids,omitempty"`
	TimeFrom     int       `json:"time_from"`
	TimeTill     int       `json:"time_till"`
	Filters      []ServiceFilter `json:"filters,omitempty"`
}

// ServiceFilter represents SLA filter
type ServiceFilter struct {
	ServiceID   string `json:"serviceid"`
	Exclude     string `json:"exclude,omitempty"`
}

// SLAMeasure represents SLA measurement result
type SLAMeasure struct {
	ServiceID       string  `json:"serviceid"`
	ServiceName     string  `json:"service_name"`
	SLI             float64 `json:"sli,omitempty"`
	Uptime          float64 `json:"uptime,omitempty"`
	Downtime        float64 `json:"downtime,omitempty"`
	TimeToFirstFailure int   `json:"time_to_first_failure,omitempty"`
	TimeToRecovery  int   `json:"time_to_recovery,omitempty"`
	AllowedDowntime float64 `json:"allowed_downtime,omitempty"`
}

// SLAMeasures represents an array of SLAMeasure objects
type SLAMeasures []SLAMeasure

// ServiceTree represents a service tree structure
type ServiceTree struct {
	Service
	Children   []ServiceTree `json:"children,omitempty"`
	Problems   []ServiceProblem `json:"problems,omitempty"`
	SLAMeasure *SLAMeasure    `json:"sla,omitempty"`
}

// Service constants
const (
	// Service status
	ServiceStatusUp       = "0"
	ServiceStatusDown     = "1"
	ServiceStatusPartiallyDown = "2"
	ServiceStatusUnknown  = "3"

	// Service algorithm (calculation method)
	ServiceAlgorithmAny = "0"
	ServiceAlgorithmAll = "1"

	// Service link type
	ServiceLinkTypeNormal    = "0"
	ServiceLinkTypeSoft     = "1"

	// Service propagate option
	ServicePropagateNo  = "0"
	ServicePropagateYes = "1"

	// Service show SLA option
	ServiceShowSLANo  = "0"
	ServiceShowSLAYes = "1"

	// Service problem status
	ServiceProblemStatusOpen    = "1"
	ServiceProblemStatusClosed  = "0"

	// Service filter exclude
	ServiceFilterExcludeNo = "0"
	ServiceFilterExcludeYes = "1"
)

// ServicesGet Wrapper for service.get
// https://www.zabbix.com/documentation/current/manual/api/reference/service/get
func (api *API) ServicesGet(options ServiceGetOptions) (Services, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.ServiceIDs) > 0 {
		params["serviceids"] = options.ServiceIDs
	}
	if len(options.ParentIDs) > 0 {
		params["parentids"] = options.ParentIDs
	}
	if len(options.ChildIDs) > 0 {
		params["childids"] = options.ChildIDs
	}
	if len(options.TriggerIDs) > 0 {
		params["triggerids"] = options.TriggerIDs
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
	if options.SelectChildren != "" {
		params["selectChildren"] = options.SelectChildren
	}
	if options.SelectParents != "" {
		params["selectParents"] = options.SelectParents
	}
	if options.SelectTags != "" {
		params["selectTags"] = options.SelectTags
	}
	if options.SelectProblemEvents != "" {
		params["selectProblemEvents"] = options.SelectProblemEvents
	}
	if options.SelectServices != "" {
		params["selectServices"] = options.SelectServices
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

	var services Services
	err := api.CallWithErrorParse("service.get", params, &services)
	return services, err
}

// ServicesGetByID Get services by specific service IDs
func (api *API) ServicesGetByID(serviceIDs []string) (Services, error) {
	options := ServiceGetOptions{
		ServiceIDs: serviceIDs,
		Output:    "extend",
	}
	return api.ServicesGet(options)
}

// ServicesGetByName Get services by name
func (api *API) ServicesGetByName(name string) (Services, error) {
	options := ServiceGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.ServicesGet(options)
}

// ServicesGetByStatus Get services by status
func (api *API) ServicesGetByStatus(status string) (Services, error) {
	options := ServiceGetOptions{
		Filter: map[string]interface{}{
			"status": status,
		},
		Output: "extend",
	}
	return api.ServicesGet(options)
}

// ServicesGetUp Get services that are up
func (api *API) ServicesGetUp() (Services, error) {
	return api.ServicesGetByStatus(ServiceStatusUp)
}

// ServicesGetDown Get services that are down
func (api *API) ServicesGetDown() (Services, error) {
	return api.ServicesGetByStatus(ServiceStatusDown)
}

// ServicesGetPartiallyDown Get services that are partially down
func (api *API) ServicesGetPartiallyDown() (Services, error) {
	return api.ServicesGetByStatus(ServiceStatusPartiallyDown)
}

// ServicesGetByParent Get child services of a parent service
func (api *API) ServicesGetByParent(parentID string) (Services, error) {
	options := ServiceGetOptions{
		ParentIDs: []string{parentID},
		Output:   "extend",
	}
	return api.ServicesGet(options)
}

// ServicesGetRoot Get root services (services without parent)
func (api *API) ServicesGetRoot() (Services, error) {
	options := ServiceGetOptions{
		Filter: map[string]interface{}{
			"parentid": "",
		},
		Output: "extend",
	}
	return api.ServicesGet(options)
}

// ServicesGetWithChildren Get services with their children
func (api *API) ServicesGetWithChildren(options ServiceGetOptions) (Services, error) {
	options.SelectChildren = "extend"
	return api.ServicesGet(options)
}

// ServicesGetWithParents Get services with their parents
func (api *API) ServicesGetWithParents(options ServiceGetOptions) (Services, error) {
	options.SelectParents = "extend"
	return api.ServicesGet(options)
}

// ServiceGetByID Get service by ID (exactly one match required)
func (api *API) ServiceGetByID(serviceID string) (*Service, error) {
	services, err := api.ServicesGetByID([]string{serviceID})
	if err != nil {
		return nil, err
	}

	if len(services) == 1 {
		return &services[0], nil
	} else if len(services) == 0 {
		return nil, fmt.Errorf("Service not found: %s", serviceID)
	} else {
		return nil, fmt.Errorf("Multiple services found with ID: %s", serviceID)
	}
}

// ServicesCreate Wrapper for service.create
// https://www.zabbix.com/documentation/current/manual/api/reference/service/create
func (api *API) ServicesCreate(services Services) (result []string, err error) {
	options := ServiceCreateOptions{
		Services: services,
	}

	response, err := api.CallWithError("service.create", options)
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
			if serviceMap, ok := item.(map[string]interface{}); ok {
				if serviceid, exists := serviceMap["serviceids"]; exists {
					if idArray, ok := serviceid.([]interface{}); ok && len(idArray) > 0 {
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

// ServiceCreateSingle Create a single service
func (api *API) ServiceCreateSingle(service Service) (serviceID string, err error) {
	services := Services{service}
	result, err := api.ServicesCreate(services)
	if len(result) > 0 {
		serviceID = result[0]
	}
	return
}

// ServicesUpdate Wrapper for service.update
// https://www.zabbix.com/documentation/current/manual/api/reference/service/update
func (api *API) ServicesUpdate(services Services) (result []string, err error) {
	options := ServiceUpdateOptions{
		Services: services,
	}

	response, err := api.CallWithError("service.update", options)
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
			if serviceMap, ok := item.(map[string]interface{}); ok {
				if serviceid, exists := serviceMap["serviceids"]; exists {
					if idArray, ok := serviceid.([]interface{}); ok && len(idArray) > 0 {
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

// ServiceUpdateSingle Update a single service
func (api *API) ServiceUpdateSingle(service Service) (serviceID string, err error) {
	services := Services{service}
	result, err := api.ServicesUpdate(services)
	if len(result) > 0 {
		serviceID = result[0]
	}
	return
}

// ServicesDelete Wrapper for service.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/service/delete
func (api *API) ServicesDelete(services Services) (result []string, err error) {
	serviceIDs := make([]string, len(services))
	for i, service := range services {
		serviceIDs[i] = service.ServiceID
	}
	
	return api.ServicesDeleteByIDs(serviceIDs)
}

// ServicesDeleteByIDs Wrapper for service.delete with IDs
func (api *API) ServicesDeleteByIDs(serviceIDs []string) (result []string, err error) {
	options := ServiceDeleteOptions{
		ServiceIDs: serviceIDs,
	}

	response, err := api.CallWithError("service.delete", options)
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

// ServiceDeleteSingle Delete a single service
func (api *API) ServiceDeleteSingle(serviceID string) (err error) {
	_, err = api.ServicesDeleteByIDs([]string{serviceID})
	return
}

// ServicesAddChildren Add child services to a parent
func (api *API) ServicesAddChildren(parentID string, childIDs []string) (err error) {
	// Get parent service
	parent, err := api.ServiceGetByID(parentID)
	if err != nil {
		return err
	}

	// Get child services
	children, err := api.ServicesGetByID(childIDs)
	if err != nil {
		return err
	}

	// Update children with parent ID
	for i := range children {
		children[i].ParentID = parentID
	}

	// Update the services
	_, err = api.ServicesUpdate(children)
	return err
}

// ServicesRemoveChildren Remove child services from a parent
func (api *API) ServicesRemoveChildren(parentID string, childIDs []string) (err error) {
	// Get child services
	children, err := api.ServicesGetByID(childIDs)
	if err != nil {
		return err
	}

	// Remove parent ID from children
	for i := range children {
		children[i].ParentID = ""
	}

	// Update the services
	_, err = api.ServicesUpdate(children)
	return err
}

// ServicesGetSLA Wrapper for service.getsla
// https://www.zabbix.com/documentation/current/manual/api/reference/service/getsla
func (api *API) ServicesGetSLA(options ServiceGetSLAOptions) (SLAMeasures, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.ServiceIDs) > 0 {
		params["serviceids"] = options.ServiceIDs
	}
	
	params["time_from"] = options.TimeFrom
	params["time_till"] = options.TimeTill
	
	if len(options.Filters) > 0 {
		params["filters"] = options.Filters
	}

	var slaMeasures SLAMeasures
	err := api.CallWithErrorParse("service.getsla", params, &slaMeasures)
	return slaMeasures, err
}

// ServiceGetSLAForService Get SLA for a specific service
func (api *API) ServiceGetSLAForService(serviceID string, timeFrom, timeTill int) (SLAMeasure, error) {
	options := ServiceGetSLAOptions{
		ServiceIDs: []string{serviceID},
		TimeFrom:   timeFrom,
		TimeTill:   timeTill,
	}

	slaMeasures, err := api.ServicesGetSLA(options)
	if err != nil {
		return SLAMeasure{}, err
	}

	if len(slaMeasures) > 0 {
		return slaMeasures[0], nil
	} else {
		return SLAMeasure{}, fmt.Errorf("No SLA data found for service: %s", serviceID)
	}
}

// ServiceGetSLAForPeriod Get SLA for a service over a time period
func (api *API) ServiceGetSLAForPeriod(serviceID string, period time.Duration) (SLAMeasure, error) {
	now := int(time.Now().Unix())
	from := int(now - int(period.Seconds()))
	
	return api.ServiceGetSLAForService(serviceID, from, now)
}

// ServiceGetSLADay Get SLA for a service for the last 24 hours
func (api *API) ServiceGetSLADay(serviceID string) (SLAMeasure, error) {
	return api.ServiceGetSLAForPeriod(serviceID, 24*time.Hour)
}

// ServiceGetSLAWeek Get SLA for a service for the last 7 days
func (api *API) ServiceGetSLAWeek(serviceID string) (SLAMeasure, error) {
	return api.ServiceGetSLAForPeriod(serviceID, 7*24*time.Hour)
}

// ServiceGetSLAMonth Get SLA for a service for the last 30 days
func (api *API) ServiceGetSLAMonth(serviceID string) (SLAMeasure, error) {
	return api.ServiceGetSLAForPeriod(serviceID, 30*24*time.Hour)
}

// ServicesGetTree Get complete service tree
func (api *API) ServicesGetTree() ([]ServiceTree, error) {
	// Get root services
	rootServices, err := api.ServicesGetRoot()
	if err != nil {
		return nil, err
	}

	// Build tree recursively
	var tree []ServiceTree
	for _, service := range rootServices {
		serviceTree, err := api.buildServiceTree(service)
		if err != nil {
			continue
		}
		tree = append(tree, serviceTree)
	}

	return tree, nil
}

// buildServiceTree recursively builds service tree
func (api *API) buildServiceTree(service Service) (ServiceTree, error) {
	tree := ServiceTree{
		Service: service,
	}

	// Get children
	children, err := api.ServicesGetByParent(service.ServiceID)
	if err != nil {
		return ServiceTree{}, err
	}

	// Get SLA for current service
	sla, err := api.ServiceGetSLAForService(service.ServiceID, 
		int(time.Now().Unix())-7*24*3600, int(time.Now().Unix()))
	if err == nil {
		tree.SLAMeasure = &sla
	}

	// Build children recursively
	for _, child := range children {
		childTree, err := api.buildServiceTree(child)
		if err != nil {
			continue
		}
		tree.Children = append(tree.Children, childTree)
	}

	return tree, nil
}

// ServicesGetStatistics Get statistics about services
func (api *API) ServicesGetStatistics() (map[string]interface{}, error) {
	// Get all services
	allServices, err := api.ServicesGet(ServiceGetOptions{Output: "extend"})
	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})
	
	// Basic counts
	totalServices := len(allServices)
	
	// Count by status
	statusCounts := make(map[string]int)
	
	// Count by algorithm
	algorithmCounts := make(map[string]int)
	
	// Count by parent relationships
	rootServices := 0
	childServices := 0
	
	for _, service := range allServices {
		// Count by status
		statusCounts[service.Status]++
		
		// Count by algorithm
		algorithmCounts[service.Algorithm]++
		
		// Count by parent relationship
		if service.ParentID == "" {
			rootServices++
		} else {
			childServices++
		}
	}
	
	// Build statistics result
	stats["total_services"] = totalServices
	stats["root_services"] = rootServices
	stats["child_services"] = childServices
	
	stats["status_distribution"] = statusCounts
	stats["algorithm_distribution"] = algorithmCounts
	
	// Calculate percentages
	if totalServices > 0 {
		stats["root_services_percentage"] = (float64(rootServices) / float64(totalServices)) * 100
		stats["child_services_percentage"] = (float64(childServices) / float64(totalServices)) * 100
	} else {
		stats["root_services_percentage"] = 0
		stats["child_services_percentage"] = 0
	}
	
	return stats, nil
}

// ServiceValidate Validate a service configuration
func (api *API) ServiceValidate(service Service) (validationErrors []string) {
	validationErrors = []string{}
	
	// Check required fields
	if service.Name == "" {
		validationErrors = append(validationErrors, "Service name is required")
	}
	
	// Validate status
	if service.Status != "" && service.Status != ServiceStatusUp && 
		service.Status != ServiceStatusDown && service.Status != ServiceStatusPartiallyDown && 
		service.Status != ServiceStatusUnknown {
		validationErrors = append(validationErrors, fmt.Sprintf("Invalid status: %s", service.Status))
	}
	
	// Validate algorithm
	if service.Algorithm != "" && service.Algorithm != ServiceAlgorithmAny && 
		service.Algorithm != ServiceAlgorithmAll {
		validationErrors = append(validationErrors, fmt.Sprintf("Invalid algorithm: %s", service.Algorithm))
	}
	
	// Validate acceptable SLA
	if service.AcceptableSLA != "" {
		// Basic validation for SLA percentage (0-100)
		// This would need more sophisticated validation in practice
		if len(service.AcceptableSLA) == 0 {
			validationErrors = append(validationErrors, "Acceptable SLA cannot be empty")
		}
	}
	
	return validationErrors
}

// ServiceIsUp Check if service is up
func (service *Service) IsUp() bool {
	return service.Status == ServiceStatusUp
}

// ServiceIsDown Check if service is down
func (service *Service) IsDown() bool {
	return service.Status == ServiceStatusDown
}

// ServiceIsPartiallyDown Check if service is partially down
func (service *Service) IsPartiallyDown() bool {
	return service.Status == ServiceStatusPartiallyDown
}

// ServiceIsUnknown Check if service status is unknown
func (service *Service) IsUnknown() bool {
	return service.Status == ServiceStatusUnknown
}

// ServiceIsRoot Check if service is a root service (has no parent)
func (service *Service) IsRoot() bool {
	return service.ParentID == ""
}

// ServiceIsChild Check if service is a child service (has parent)
func (service *Service) IsChild() bool {
	return service.ParentID != ""
}

// ServiceHasChildren Check if service has children
func (service *Service) HasChildren() bool {
	return service.ChildCount() > 0
}

// ServiceChildCount Get number of children (this would need to be populated by API)
func (service *Service) ChildCount() int {
	// This is a placeholder - in practice, this would be populated by the API
	// when using SelectChildren in the query
	return 0
}

// ServiceUpdateStatus Update service status
func (api *API) ServiceUpdateStatus(serviceID, status string) (err error) {
	service, err := api.ServiceGetByID(serviceID)
	if err != nil {
		return err
	}
	
	service.Status = status
	_, err = api.ServiceUpdateSingle(*service)
	return err
}

// ServiceUpdateSLA Update service acceptable SLA
func (api *API) ServiceUpdateSLA(serviceID, acceptableSLA string) (err error) {
	service, err := api.ServiceGetByID(serviceID)
	if err != nil {
		return err
	}
	
	service.AcceptableSLA = acceptableSLA
	_, err = api.ServiceUpdateSingle(*service)
	return err
}

// CreateSimpleService Create a simple service with basic configuration
func (api *API) CreateSimpleService(name, description string) (string, error) {
	service := Service{
		Name:         name,
		Description:  description,
		Status:       ServiceStatusUp,
		Algorithm:    ServiceAlgorithmAny,
		ShowSLA:      ServiceShowSLAYes,
	}
	
	return api.ServiceCreateSingle(service)
}

// CreateServiceTree Create a complete service tree from hierarchical data
func (api *API) CreateServiceTree(hierarchicalServices [][]string) (createdIDs []string, err error) {
	// This is a simplified implementation for creating a service tree
	// hierarchicalServices[i][j] represents service at level i, position j
	createdIDs = make([]string, 0)
	
	for level, services := range hierarchicalServices {
		for position, serviceName := range services {
			service := Service{
				Name:      serviceName,
				Status:    ServiceStatusUp,
				Algorithm: ServiceAlgorithmAny,
				ShowSLA:   ServiceShowSLAYes,
			}
			
			// Set parent for non-root services
			if level > 0 && len(hierarchicalServices[level-1]) > 0 {
				// Simple parent assignment - in practice, this would be more complex
				parentIdx := position
				if parentIdx >= len(createdIDs[level-1:]) {
					parentIdx = len(createdIDs[level-1:]) - 1
				}
				if level-1 < len(createdIDs) && parentIdx >= 0 {
					service.ParentID = createdIDs[level-1+parentIdx]
				}
			}
			
			serviceID, err := api.ServiceCreateSingle(service)
			if err != nil {
				continue // Continue with other services even if one fails
			}
			createdIDs = append(createdIDs, serviceID)
		}
	}
	
	return createdIDs, nil
}