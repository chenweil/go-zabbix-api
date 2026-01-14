package zabbix

import (
	"encoding/json"
	"fmt"
)

// DiscoveryRule represents a discovery rule object
// https://www.zabbix.com/documentation/current/manual/api/reference/drule/object
type DiscoveryRule struct {
	RuleID      string `json:"druleid,omitempty"`
	Name        string `json:"name"`
	IPRange     string `json:"iprange"`
	Delay       string `json:"delay"`
	ProxyHostID string `json:"proxy_hostid,omitempty"`
	Status      string `json:"status,string"`
	Checks      []DiscoveryCheck `json:"checks,omitempty"`
}

// DiscoveryRules represents an array of DiscoveryRule objects
type DiscoveryRules []DiscoveryRule

// DiscoveryCheck represents a discovery check object
// https://www.zabbix.com/documentation/current/manual/api/reference/dcheck/object
type DiscoveryCheck struct {
	CheckID     string `json:"dcheckid,omitempty"`
	RuleID     string `json:"druleid,omitempty"`
	Type       string `json:"type,string"`
	SnmpCommunity string `json:"snmp_community,omitempty"`
	SnmpOID    string `json:"snmp_oid,omitempty"`
	Port       string `json:"port,omitempty"`
	SNMPV3SecurityName string `json:"snmpv3_securityname,omitempty"`
	SNMPV3SecurityPassphrase string `json:"snmpv3_securitypassphrase,omitempty"`
	SNMPV3SecurityLevel string `json:"snmpv3_securitylevel,string"`
	SNMPV3AuthProtocol string `json:"snmpv3_authprotocol,string"`
	SNMPV3PrivPassphrase string `json:"snmpv3_privpassphrase,omitempty"`
	SNMPV3PrivProtocol string `json:"snmpv3_privprotocol,string"`
	SNMPV3ContextName string `json:"snmpv3_contextname,omitempty"`
	HostSource    string `json:"host_source,string"`
	NameSource    string `json:"name_source,string"`
	AllowedHosts string `json:"allowed_hosts,omitempty"`
}

// DiscoveryChecks represents an array of DiscoveryCheck objects
type DiscoveryChecks []DiscoveryCheck

// DiscoveredHost represents a discovered host object
// https://www.zabbix.com/documentation/current/manual/api/reference/dhost/object
type DiscoveredHost struct {
	HostID     string `json:"dhostid,omitempty"`
	Host       string `json:"host"`
	LastCheck  string `json:"lastcheck"`
	HostStatus string `json:"host"`
}

// DiscoveredHosts represents an array of DiscoveredHost objects
type DiscoveredHosts []DiscoveredHost

// DiscoveredService represents a discovered service object
// https://www.zabbix.com/documentation/current/manual/api/reference/dservice/object
type DiscoveredService struct {
	ServiceID  string `json:"dserviceid,omitempty"`
	HostID    string `json:"dhostid,omitempty"`
	Service   string `json:"service"`
	Port      string `json:"port"`
	LastCheck string `json:"lastcheck"`
	ServiceStatus string `json:"service"`
	ServiceType string `json:"service_type"`
	ServiceKey string `json:"service_key"`
}

// DiscoveredServices represents an array of DiscoveredService objects
type DiscoveredServices []DiscoveredService

// DiscoveryRuleGetOptions represents parameters for drule.get API call
type DiscoveryRuleGetOptions struct {
	RuleIDs         []string             `json:"druleids,omitempty"`
	Filter         map[string]interface{} `json:"filter,omitempty"`
	Search         map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string        `json:"searchWildcardsEnabled,omitempty"`
	Output         string               `json:"output,omitempty"`
	SelectChecks   string               `json:"selectChecks,omitempty"`
	SortField      string               `json:"sortfield,omitempty"`
	SortOrder      string               `json:"sortorder,omitempty"`
	Limit          int                  `json:"limit,omitempty"`
}

// DiscoveryCheckGetOptions represents parameters for dcheck.get API call
type DiscoveryCheckGetOptions struct {
	CheckIDs        []string             `json:"dcheckids,omitempty"`
	RuleIDs        []string             `json:"druleids,omitempty"`
	Filter         map[string]interface{} `json:"filter,omitempty"`
	Search         map[string]interface{} `json:"search,omitempty"`
	Output         string               `json:"output,omitempty"`
	SortField      string               `json:"sortfield,omitempty"`
	SortOrder      string               `json:"sortorder,omitempty"`
	Limit          int                  `json:"limit,omitempty"`
}

// DiscoveredHostGetOptions represents parameters for dhost.get API call
type DiscoveredHostGetOptions struct {
	HostIDs        []string             `json:"dhostids,omitempty"`
	RuleIDs       []string             `json:"druleids,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	Search        map[string]interface{} `json:"search,omitempty"`
	Output        string               `json:"output,omitempty"`
	SelectServices string              `json:"selectServices,omitempty"`
	SortField     string               `json:"sortfield,omitempty"`
	SortOrder     string               `json:"sortorder,omitempty"`
	Limit         int                  `json:"limit,omitempty"`
}

// DiscoveredServiceGetOptions represents parameters for dservice.get API call
type DiscoveredServiceGetOptions struct {
	ServiceIDs    []string             `json:"dserviceids,omitempty"`
	HostIDs      []string             `json:"dhostids,omitempty"`
	RuleIDs     []string             `json:"druleids,omitempty"`
	Filter       map[string]interface{} `json:"filter,omitempty"`
	Search       map[string]interface{} `json:"search,omitempty"`
	Output       string               `json:"output,omitempty"`
	SortField    string               `json:"sortfield,omitempty"`
	SortOrder    string               `json:"sortorder,omitempty"`
	Limit        int                  `json:"limit,omitempty"`
}

// Discovery constants
const (
	// Discovery rule status
	DiscoveryRuleStatusEnabled  = "0"
	DiscoveryRuleStatusDisabled = "1"

	// Discovery check types
	DiscoveryCheckTypeSSH   = "0"
	DiscoveryCheckTypeLDAP  = "1"
	DiscoveryCheckTypeSMTP = "2"
	DiscoveryCheckTypeFTP   = "3"
	DiscoveryCheckTypeHTTP = "4"
	DiscoveryCheckTypeHTTPS = "5"
	DiscoveryCheckTypePOP   = "6"
	DiscoveryCheckTypeNNTP  = "7"
	DiscoveryCheckTypeIMAP  = "8"
	DiscoveryCheckTypeTCP   = "9"
	DiscoveryCheckTypeICMP  = "10"
	DiscoveryCheckTypeSNMPv1 = "11"
	DiscoveryCheckTypeSNMPv2c = "12"
	DiscoveryCheckTypeSNMPv3 = "13"
	DiscoveryCheckTypeTelnet = "14"
	DiscoveryCheckTypeCustom = "19"

	// Host source for discovery
	DiscoveryHostSourceIP    = "0"
	DiscoveryHostSourceDNS   = "1"

	// Name source for discovery
	DiscoveryNameSourceDNS = "0"
	DiscoveryNameSourceIP  = "1"

	// SNMP security levels
	SNMPv3SecurityLevelNoAuthNoPriv = "0"
	SNMPv3SecurityLevelAuthNoPriv   = "1"
	SNMPv3SecurityLevelAuthPriv      = "2"

	// SNMP authentication protocols
	SNMPv3AuthProtocolMD5   = "0"
	SNMPv3AuthProtocolSHA    = "1"

	// SNMP privacy protocols
	SNMPv3PrivProtocolDES   = "0"
	SNMPv3PrivProtocolAES    = "1"
)

// DiscoveryRulesGet Wrapper for drule.get
// https://www.zabbix.com/documentation/current/manual/api/reference/drule/get
func (api *API) DiscoveryRulesGet(options DiscoveryRuleGetOptions) (DiscoveryRules, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.RuleIDs) > 0 {
		params["druleids"] = options.RuleIDs
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
	if options.SelectChecks != "" {
		params["selectChecks"] = options.SelectChecks
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

	var rules DiscoveryRules
	err := api.CallWithErrorParse("drule.get", params, &rules)
	return rules, err
}

// DiscoveryRulesGetByID Get discovery rules by specific rule IDs
func (api *API) DiscoveryRulesGetByID(ruleIDs []string) (DiscoveryRules, error) {
	options := DiscoveryRuleGetOptions{
		RuleIDs: ruleIDs,
		Output: "extend",
	}
	return api.DiscoveryRulesGet(options)
}

// DiscoveryRulesGetByName Get discovery rules by name
func (api *API) DiscoveryRulesGetByName(name string) (DiscoveryRules, error) {
	options := DiscoveryRuleGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.DiscoveryRulesGet(options)
}

// DiscoveryRulesGetEnabled Get enabled discovery rules
func (api *API) DiscoveryRulesGetEnabled() (DiscoveryRules, error) {
	options := DiscoveryRuleGetOptions{
		Filter: map[string]interface{}{
			"status": DiscoveryRuleStatusEnabled,
		},
		Output: "extend",
	}
	return api.DiscoveryRulesGet(options)
}

// DiscoveryRulesGetDisabled Get disabled discovery rules
func (api *API) DiscoveryRulesGetDisabled() (DiscoveryRules, error) {
	options := DiscoveryRuleGetOptions{
		Filter: map[string]interface{}{
			"status": DiscoveryRuleStatusDisabled,
		},
		Output: "extend",
	}
	return api.DiscoveryRulesGet(options)
}

// DiscoveryRulesGetWithChecks Get discovery rules with their checks
func (api *API) DiscoveryRulesGetWithChecks(options DiscoveryRuleGetOptions) (DiscoveryRules, error) {
	options.SelectChecks = "extend"
	return api.DiscoveryRulesGet(options)
}

// DiscoveryRuleGetByID Get discovery rule by ID (exactly one match required)
func (api *API) DiscoveryRuleGetByID(ruleID string) (*DiscoveryRule, error) {
	rules, err := api.DiscoveryRulesGetByID([]string{ruleID})
	if err != nil {
		return nil, err
	}

	if len(rules) == 1 {
		return &rules[0], nil
	} else if len(rules) == 0 {
		return nil, fmt.Errorf("Discovery rule not found: %s", ruleID)
	} else {
		return nil, fmt.Errorf("Multiple discovery rules found with ID: %s", ruleID)
	}
}

// DiscoveryRulesCreate Wrapper for drule.create
// https://www.zabbix.com/documentation/current/manual/api/reference/drule/create
func (api *API) DiscoveryRulesCreate(rules DiscoveryRules) (result []string, err error) {
	response, err := api.CallWithError("drule.create", rules)
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
			if ruleMap, ok := item.(map[string]interface{}); ok {
				if druleid, exists := ruleMap["druleids"]; exists {
					if idArray, ok := druleid.([]interface{}); ok && len(idArray) > 0 {
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

// DiscoveryRuleCreateSingle Create a single discovery rule
func (api *API) DiscoveryRuleCreateSingle(rule DiscoveryRule) (ruleID string, err error) {
	rules := DiscoveryRules{rule}
	result, err := api.DiscoveryRulesCreate(rules)
	if len(result) > 0 {
		ruleID = result[0]
	}
	return
}

// DiscoveryRulesUpdate Wrapper for drule.update
// https://www.zabbix.com/documentation/current/manual/api/reference/drule/update
func (api *API) DiscoveryRulesUpdate(rules DiscoveryRules) (result []string, err error) {
	response, err := api.CallWithError("drule.update", rules)
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
			if ruleMap, ok := item.(map[string]interface{}); ok {
				if druleid, exists := ruleMap["druleids"]; exists {
					if idArray, ok := druleid.([]interface{}); ok && len(idArray) > 0 {
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

// DiscoveryRuleUpdateSingle Update a single discovery rule
func (api *API) DiscoveryRuleUpdateSingle(rule DiscoveryRule) (ruleID string, err error) {
	rules := DiscoveryRules{rule}
	result, err := api.DiscoveryRulesUpdate(rules)
	if len(result) > 0 {
		ruleID = result[0]
	}
	return
}

// DiscoveryRulesDelete Wrapper for drule.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/drule/delete
func (api *API) DiscoveryRulesDelete(rules DiscoveryRules) (result []string, err error) {
	ruleIDs := make([]string, len(rules))
	for i, rule := range rules {
		ruleIDs[i] = rule.RuleID
	}
	
	return api.DiscoveryRulesDeleteByIDs(ruleIDs)
}

// DiscoveryRulesDeleteByIDs Wrapper for drule.delete with IDs
func (api *API) DiscoveryRulesDeleteByIDs(ruleIDs []string) (result []string, err error) {
	response, err := api.CallWithError("drule.delete", ruleIDs)
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

// DiscoveryRuleDeleteSingle Delete a single discovery rule
func (api *API) DiscoveryRuleDeleteSingle(ruleID string) (err error) {
	_, err = api.DiscoveryRulesDeleteByIDs([]string{ruleID})
	return
}

// DiscoveryRulesEnable Enable discovery rules
func (api *API) DiscoveryRulesEnable(ruleIDs []string) (result []string, err error) {
	rules, err := api.DiscoveryRulesGetByID(ruleIDs)
	if err != nil {
		return
	}

	for i := range rules {
		rules[i].Status = DiscoveryRuleStatusEnabled
	}

	return api.DiscoveryRulesUpdate(rules)
}

// DiscoveryRulesDisable Disable discovery rules
func (api *API) DiscoveryRulesDisable(ruleIDs []string) (result []string, err error) {
	rules, err := api.DiscoveryRulesGetByID(ruleIDs)
	if err != nil {
		return
	}

	for i := range rules {
		rules[i].Status = DiscoveryRuleStatusDisabled
	}

	return api.DiscoveryRulesUpdate(rules)
}

// DiscoveryChecksGet Wrapper for dcheck.get
// https://www.zabbix.com/documentation/current/manual/api/reference/dcheck/get
func (api *API) DiscoveryChecksGet(options DiscoveryCheckGetOptions) (DiscoveryChecks, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.CheckIDs) > 0 {
		params["dcheckids"] = options.CheckIDs
	}
	if len(options.RuleIDs) > 0 {
		params["druleids"] = options.RuleIDs
	}
	if options.Filter != nil {
		params["filter"] = options.Filter
	}
	if options.Search != nil {
		params["search"] = options.Search
	}
	if options.Output != "" {
		params["output"] = options.Output
	} else {
		params["output"] = "extend"
	}
	if options.SortField != "" {
		params["sortfield"] = options.SortField
	} else {
		params["sortfield"] = "type"
	}
	if options.SortOrder != "" {
		params["sortorder"] = options.SortOrder
	} else {
		params["sortorder"] = "ASC"
	}
	if options.Limit > 0 {
		params["limit"] = options.Limit
	}

	var checks DiscoveryChecks
	err := api.CallWithErrorParse("dcheck.get", params, &checks)
	return checks, err
}

// DiscoveryChecksGetByRule Get discovery checks by rule ID
func (api *API) DiscoveryChecksGetByRule(ruleID string) (DiscoveryChecks, error) {
	options := DiscoveryCheckGetOptions{
		RuleIDs: []string{ruleID},
		Output: "extend",
	}
	return api.DiscoveryChecksGet(options)
}

// DiscoveryChecksGetByID Get discovery checks by check IDs
func (api *API) DiscoveryChecksGetByID(checkIDs []string) (DiscoveryChecks, error) {
	options := DiscoveryCheckGetOptions{
		CheckIDs: checkIDs,
		Output: "extend",
	}
	return api.DiscoveryChecksGet(options)
}

// DiscoveryChecksGetByType Get discovery checks by type
func (api *API) DiscoveryChecksGetByType(checkType string) (DiscoveryChecks, error) {
	options := DiscoveryCheckGetOptions{
		Filter: map[string]interface{}{
			"type": checkType,
		},
		Output: "extend",
	}
	return api.DiscoveryChecksGet(options)
}

// DiscoveredHostsGet Wrapper for dhost.get
// https://www.zabbix.com/documentation/current/manual/api/reference/dhost/get
func (api *API) DiscoveredHostsGet(options DiscoveredHostGetOptions) (DiscoveredHosts, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.HostIDs) > 0 {
		params["dhostids"] = options.HostIDs
	}
	if len(options.RuleIDs) > 0 {
		params["druleids"] = options.RuleIDs
	}
	if options.Filter != nil {
		params["filter"] = options.Filter
	}
	if options.Search != nil {
		params["search"] = options.Search
	}
	if options.Output != "" {
		params["output"] = options.Output
	} else {
		params["output"] = "extend"
	}
	if options.SelectServices != "" {
		params["selectServices"] = options.SelectServices
	}
	if options.SortField != "" {
		params["sortfield"] = options.SortField
	} else {
		params["sortfield"] = "host"
	}
	if options.SortOrder != "" {
		params["sortorder"] = options.SortOrder
	} else {
		params["sortorder"] = "ASC"
	}
	if options.Limit > 0 {
		params["limit"] = options.Limit
	}

	var hosts DiscoveredHosts
	err := api.CallWithErrorParse("dhost.get", params, &hosts)
	return hosts, err
}

// DiscoveredHostsGetByRule Get discovered hosts by rule ID
func (api *API) DiscoveredHostsGetByRule(ruleID string) (DiscoveredHosts, error) {
	options := DiscoveredHostGetOptions{
		RuleIDs: []string{ruleID},
		Output: "extend",
	}
	return api.DiscoveredHostsGet(options)
}

// DiscoveredHostsGetByID Get discovered hosts by host IDs
func (api *API) DiscoveredHostsGetByID(hostIDs []string) (DiscoveredHosts, error) {
	options := DiscoveredHostGetOptions{
		HostIDs: hostIDs,
		Output: "extend",
	}
	return api.DiscoveredHostsGet(options)
}

// DiscoveredHostsGetActive Get active discovered hosts
func (api *API) DiscoveredHostsGetActive() (DiscoveredHosts, error) {
	options := DiscoveredHostGetOptions{
		Filter: map[string]interface{}{
			"host": "1", // Active status
		},
		Output: "extend",
	}
	return api.DiscoveredHostsGet(options)
}

// DiscoveredHostsGetWithServices Get discovered hosts with their services
func (api *API) DiscoveredHostsGetWithServices(options DiscoveredHostGetOptions) (DiscoveredHosts, error) {
	options.SelectServices = "extend"
	return api.DiscoveredHostsGet(options)
}

// DiscoveredServicesGet Wrapper for dservice.get
// https://www.zabbix.com/documentation/current/manual/api/reference/dservice/get
func (api *API) DiscoveredServicesGet(options DiscoveredServiceGetOptions) (DiscoveredServices, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.ServiceIDs) > 0 {
		params["dserviceids"] = options.ServiceIDs
	}
	if len(options.HostIDs) > 0 {
		params["dhostids"] = options.HostIDs
	}
	if len(options.RuleIDs) > 0 {
		params["druleids"] = options.RuleIDs
	}
	if options.Filter != nil {
		params["filter"] = options.Filter
	}
	if options.Search != nil {
		params["search"] = options.Search
	}
	if options.Output != "" {
		params["output"] = options.Output
	} else {
		params["output"] = "extend"
	}
	if options.SortField != "" {
		params["sortfield"] = options.SortField
	} else {
		params["sortfield"] = "service"
	}
	if options.SortOrder != "" {
		params["sortorder"] = options.SortOrder
	} else {
		params["sortorder"] = "ASC"
	}
	if options.Limit > 0 {
		params["limit"] = options.Limit
	}

	var services DiscoveredServices
	err := api.CallWithErrorParse("dservice.get", params, &services)
	return services, err
}

// DiscoveredServicesGetByHost Get discovered services by host ID
func (api *API) DiscoveredServicesGetByHost(hostID string) (DiscoveredServices, error) {
	options := DiscoveredServiceGetOptions{
		HostIDs: []string{hostID},
		Output: "extend",
	}
	return api.DiscoveredServicesGet(options)
}

// DiscoveredServicesGetByRule Get discovered services by rule ID
func (api *API) DiscoveredServicesGetByRule(ruleID string) (DiscoveredServices, error) {
	options := DiscoveredServiceGetOptions{
		RuleIDs: []string{ruleID},
		Output: "extend",
	}
	return api.DiscoveredServicesGet(options)
}

// DiscoveredServicesGetByType Get discovered services by service type
func (api *API) DiscoveredServicesGetByType(serviceType string) (DiscoveredServices, error) {
	options := DiscoveredServiceGetOptions{
		Filter: map[string]interface{}{
			"service_type": serviceType,
		},
		Output: "extend",
	}
	return api.DiscoveredServicesGet(options)
}

// DiscoveredServicesGetByPort Get discovered services by port
func (api *API) DiscoveredServicesGetByPort(port string) (DiscoveredServices, error) {
	options := DiscoveredServiceGetOptions{
		Filter: map[string]interface{}{
			"port": port,
		},
		Output: "extend",
	}
	return api.DiscoveredServicesGet(options)
}

// DiscoveryGetStatistics Get statistics about discovery system
func (api *API) DiscoveryGetStatistics() (map[string]interface{}, error) {
	stats := make(map[string]interface{})
	
	// Get discovery rules
	rules, err := api.DiscoveryRulesGet(DiscoveryRuleGetOptions{Output: "extend"})
	if err != nil {
		return nil, err
	}
	
	// Get discovered hosts
	hosts, err := api.DiscoveredHostsGet(DiscoveredHostGetOptions{Output: "extend"})
	if err != nil {
		return nil, err
	}
	
	// Get discovered services
	services, err := api.DiscoveredServicesGet(DiscoveredServiceGetOptions{Output: "extend"})
	if err != nil {
		return nil, err
	}
	
	// Basic counts
	totalRules := len(rules)
	enabledRules := 0
	disabledRules := 0
	
	totalHosts := len(hosts)
	activeHosts := 0
	inactiveHosts := 0
	
	totalServices := len(services)
	activeServices := 0
	inactiveServices := 0
	
	// Count rules by status
	for _, rule := range rules {
		if rule.Status == DiscoveryRuleStatusEnabled {
			enabledRules++
		} else {
			disabledRules++
		}
	}
	
	// Count hosts by status
	for _, host := range hosts {
		if host.HostStatus == "1" {
			activeHosts++
		} else {
			inactiveHosts++
		}
	}
	
	// Count services by status
	for _, service := range services {
		if service.ServiceStatus == "1" {
			activeServices++
		} else {
			inactiveServices++
		}
	}
	
	// Build statistics result
	stats["total_rules"] = totalRules
	stats["enabled_rules"] = enabledRules
	stats["disabled_rules"] = disabledRules
	
	stats["total_hosts"] = totalHosts
	stats["active_hosts"] = activeHosts
	stats["inactive_hosts"] = inactiveHosts
	
	stats["total_services"] = totalServices
	stats["active_services"] = activeServices
	stats["inactive_services"] = inactiveServices
	
	// Calculate percentages
	if totalRules > 0 {
		stats["enabled_rules_percentage"] = (float64(enabledRules) / float64(totalRules)) * 100
	} else {
		stats["enabled_rules_percentage"] = 0
	}
	
	if totalHosts > 0 {
		stats["active_hosts_percentage"] = (float64(activeHosts) / float64(totalHosts)) * 100
	} else {
		stats["active_hosts_percentage"] = 0
	}
	
	if totalServices > 0 {
		stats["active_services_percentage"] = (float64(activeServices) / float64(totalServices)) * 100
	} else {
		stats["active_services_percentage"] = 0
	}
	
	return stats, nil
}

// DiscoveryValidate Validate a discovery rule configuration
func (api *API) DiscoveryValidate(rule DiscoveryRule) (validationErrors []string) {
	validationErrors = []string{}
	
	// Check required fields
	if rule.Name == "" {
		validationErrors = append(validationErrors, "Discovery rule name is required")
	}
	
	if rule.IPRange == "" {
		validationErrors = append(validationErrors, "IP range is required")
	}
	
	if rule.Delay == "" {
		validationErrors = append(validationErrors, "Delay is required")
	}
	
	// Validate status
	if rule.Status != "" && rule.Status != DiscoveryRuleStatusEnabled && rule.Status != DiscoveryRuleStatusDisabled {
		validationErrors = append(validationErrors, fmt.Sprintf("Invalid status: %s", rule.Status))
	}
	
	// Validate IP range format (basic validation)
	if rule.IPRange != "" {
		// Basic IP range validation would go here
		// For now, just check it's not empty
	}
	
	// Validate delay format (should be in seconds)
	if rule.Delay != "" {
		// Basic delay validation would go here
		// For now, just check it's not empty
	}
	
	// Validate checks if present
	for i, check := range rule.Checks {
		if check.Type == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Check %d: Type is required", i))
		}
		
		if check.Port == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Check %d: Port is required", i))
		}
	}
	
	return validationErrors
}

// DiscoveryRuleIsEnabled Check if discovery rule is enabled
func (rule *DiscoveryRule) IsEnabled() bool {
	return rule.Status == DiscoveryRuleStatusEnabled
}

// DiscoveryRuleIsDisabled Check if discovery rule is disabled
func (rule *DiscoveryRule) IsDisabled() bool {
	return rule.Status == DiscoveryRuleStatusDisabled
}

// DiscoveryRuleHasChecks Check if discovery rule has checks
func (rule *DiscoveryRule) HasChecks() bool {
	return len(rule.Checks) > 0
}

// DiscoveryRuleGetCheckCount Get number of checks
func (rule *DiscoveryRule) GetCheckCount() int {
	return len(rule.Checks)
}

// DiscoveredHostIsActive Check if discovered host is active
func (host *DiscoveredHost) IsActive() bool {
	return host.HostStatus == "1"
}

// DiscoveredServiceIsActive Check if discovered service is active
func (service *DiscoveredService) IsActive() bool {
	return service.ServiceStatus == "1"
}