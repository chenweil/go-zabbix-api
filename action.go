package zabbix

import (
	"encoding/json"
	"fmt"
	"time"
)

// Action represents a Zabbix action object
// https://www.zabbix.com/documentation/current/manual/api/reference/action/object
type Action struct {
	ActionID       string   `json:"actionid,omitempty"`
	Name          string   `json:"name"`
	Eventsource   string   `json:"eventsource"`
	DefLongdata   string   `json:"def_longdata"`
	DefShortdata  string   `json:"def_shortdata"`
	RecoveryLong  string   `json:"recovery_long,omitempty"`
	RecoveryShort string   `json:"recovery_short,omitempty"`
	Status        string   `json:"status,string"`
	Filter        string   `json:"filter,omitempty"`
	EvalType      string   `json:"evaltype,omitempty"`
	RecoveryMode  string   `json:"recovery_mode,string"`
	RecoveryData  string   `json:"recovery_data,string"`
	ProblemDuration       string   `json:"problem_duration,omitempty"`
	ProblemDurationMode   string   `json:"problem_duration_mode,omitempty"`
	ProblemDurationCustom string   `json:"problem_duration_custom,omitempty"`
	RTriggerID    string   `json:"r_triggerid,omitempty"`
	Commands      []Operation `json:"operations,omitempty"`
	RecoveryCommands []Operation `json:"recovery_operations,omitempty"`
	FilterConditions []Condition `json:"conditions,omitempty"`
	
	// Extended fields
	URL           string `json:"url,omitempty"`
	Description   string `json:"description,omitempty"`
	Priority      string `json:"priority,omitempty"`
	NotifyIfAll   string `json:"notify_if_all,omitempty"`
	NotifyIfUnack string `json:"notify_if_unack,omitempty"`
	
	// Read-only fields
	ActionType    string   `json:"action_type,omitempty"`
	ActionTypeName string   `json:"action_type_name,omitempty"`
}

// Actions represents an array of Action objects
type Actions []Action

// Operation represents an operation within an action
// https://www.zabbix.com/documentation/current/manual/api/reference/action/object#operation
type Operation struct {
	OperationID   string `json:"operationid,omitempty"`
	ActionID      string `json:"actionid,omitempty"`
	OperationType string `json:"operationtype,string"`
	EscPeriod     string `json:"esc_period,omitempty"`
	EscStepFrom   string `json:"esc_step_from,omitempty"`
	EscStepTo     string `json:"esc_step_to,omitempty"`
	EvalType      string `json:"evaltype,omitempty"`
	
	// Operation specific fields
	Command       string   `json:"command,omitempty"`
	CommandType   string   `json:"commandtype,string"`
	SSHAgent      string   `json:"sshagent,omitempty"`
	IPMIAgent     string   `json:"ipmiagent,omitempty"`
	ScriptID      string   `json:"scriptid,omitempty"`
	ExecuteOn     string   `json:"execute_on,omitempty"`
	MenuPath      string   `json:"menu_path,omitempty"`
	Delay         string   `json:"delay,omitempty"`
	Port          string   `json:"port,omitempty"`
	InventoryMode string   `json:"inventory_mode,omitempty"`
	
	// User/Group operations
	OpmessageCmd  string   `json:"opmessage_cmd,omitempty"`
	Opmessage     string   `json:"opmessage,omitempty"`
	Subject       string   `json:"subject,omitempty"`
	Message       string   `json:"message,omitempty"`
	MediaTypeID   string   `json:"mediatypeid,omitempty"`
	
	// Target operations
	UserIDs       []string `json:"userids,omitempty"`
	UsrGrpIDs     []string `json:"usrgrps,omitempty"`
	HostIDs       []string `json:"hostids,omitempty"`
	HostGroupIDs  []string `json:"groupids,omitempty"`
}

// Operations represents an array of Operation objects
type Operations []Operation

// Condition represents a condition within an action filter
// https://www.zabbix.com/documentation/current/manual/api/reference/action/object#condition
type Condition struct {
	ConditionID  string `json:"conditionid,omitempty"`
	ActionID     string `json:"actionid,omitempty"`
	ConditionType string `json:"conditiontype,string"`
	Operator     string `json:"operator,string"`
	Value        string `json:"value"`
	Value2       string `json:"value2,omitempty"`
	
	// Extended fields for complex conditions
	FormulaID    string `json:"formulaid,omitempty"`
	Formula      string `json:"formula,omitempty"`
}

// Conditions represents an array of Condition objects
type Conditions []Condition

// ActionGetOptions represents parameters for action.get API call
// https://www.zabbix.com/documentation/current/manual/api/reference/action/get
type ActionGetOptions struct {
	ActionIDs          []string             `json:"actionids,omitempty"`
	Filter             map[string]interface{} `json:"filter,omitempty"`
	Search             map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string             `json:"searchWildcardsEnabled,omitempty"`
	Output             string               `json:"output,omitempty"`
	SelectOperations   string               `json:"selectOperations,omitempty"`
	SelectConditions   string               `json:"selectConditions,omitempty"`
	SelectFilter      string               `json:"selectFilter,omitempty"`
	SortField         string               `json:"sortfield,omitempty"`
	SortOrder         string               `json:"sortorder,omitempty"`
	Limit             int                  `json:"limit,omitempty"`
}

// ActionCreateOptions represents parameters for action.create API call
// https://www.zabbix.com/documentation/current/manual/api/reference/action/create
type ActionCreateOptions struct {
	Actions           Actions       `json:"actions"`
	SelectOperations  string       `json:"selectOperations,omitempty"`
	SelectConditions  string       `json:"selectConditions,omitempty"`
}

// ActionUpdateOptions represents parameters for action.update API call
// https://www.zabbix.com/documentation/current/manual/api/reference/action/update
type ActionUpdateOptions struct {
	Actions           Actions       `json:"actions"`
	SelectOperations  string       `json:"selectOperations,omitempty"`
	SelectConditions  string       `json:"selectConditions,omitempty"`
}

// ActionDeleteOptions represents parameters for action.delete API call
// https://www.zabbix.com/documentation/current/manual/api/reference/action/delete
type ActionDeleteOptions struct {
	ActionIDs []string `json:"actionids"`
}

// Action constants
const (
	// Action source constants
	ActionSourceTrigger    = "0"
	ActionSourceDiscovery  = "1"
	ActionSourceAutoreg    = "2"
	ActionSourceInternal   = "3"

	// Action status constants
	ActionStatusEnabled  = "0"
	ActionStatusDisabled = "1"

	// Action recovery mode constants
	ActionRecoveryModeNone = "0"
	ActionRecoveryModeOperations = "1"
	ActionRecoveryModeBoth = "2"

	// Operation type constants
	OperationTypeSendMessage = "0"
	OperationTypeRemoteCommand = "1"
	OperationTypeAddHost = "2"
	OperationTypeRemoveHost = "3"
	OperationTypeAddToHostGroup = "4"
	OperationTypeRemoveFromHostGroup = "5"
	OperationTypeLinkTemplate = "6"
	OperationTypeUnlinkTemplate = "7"
	OperationTypeEnableHost = "8"
	OperationTypeDisableHost = "9"
	OperationTypeCustomScript = "10"

	// Condition type constants
	ConditionTypeTrigger = "0"
	ConditionTypeTriggerName = "1"
	ConditionTypeTriggerSeverity = "2"
	ConditionTypeTriggerValue = "3"
	ConditionTypeMaintenanceStatus = "12"
	ConditionTypeEventTag = "25"
	ConditionTypeEventTagValue = "26"
	ConditionTypeHost = "10"
	ConditionTypeHostGroup = "11"
	ConditionTypeTemplate = "13"
	ConditionTypeApplication = "15"
	ConditionTypeTriggerProxy = "16"
	ConditionTypeHostProxy = "17"
	ConditionTypeEventType = "18"
	ConditionTypeHostMetadata = "21"
	ConditionTypeHostInventory = "22"
	ConditionTypeEventTagValueEx = "27"

	// Condition operator constants
	ConditionOperatorEquals = "0"
	ConditionOperatorNotEquals = "1"
	ConditionOperatorContains = "2"
	ConditionOperatorNotContains = "3"
	ConditionOperatorInRange = "4"
	ConditionOperatorNotInRange = "5"
	ConditionOperatorMatchesPattern = "6"
	ConditionOperatorNotMatchesPattern = "7"

	// Command type constants
	CommandTypeIPMI = "0"
	CommandTypeSSH = "1"
	CommandTypeTelnet = "2"
	CommandTypeCustomScript = "3"
	CommandTypeGlobalScript = "4"

	// Message delivery method
	OpmessageCmdUserMedia = "0"
	OpmessageCmdAllMedia = "1"

	// Target operation modes
	OperationModeCurrentHost = "0"
	OperationModeTargetHost = "1"
	OperationModeHostGroup = "2"

	// Execute on host
	ExecuteOnZabbixAgent = "0"
	ExecuteOnZabbixServer = "1"
	ExecuteOnZabbixProxy = "2"

	// Problem duration modes
	ProblemDurationModeNone = "0"
	ProblemDurationModeTime = "1"
	ProblemDurationModeCustom = "2"

	// Inventory mode
	InventoryModeManual = "0"
	InventoryModeAutomatic = "1"
)

// ActionsGet Wrapper for action.get
// https://www.zabbix.com/documentation/current/manual/api/reference/action/get
func (api *API) ActionsGet(options ActionGetOptions) (Actions, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.ActionIDs) > 0 {
		params["actionids"] = options.ActionIDs
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
	if options.SelectOperations != "" {
		params["selectOperations"] = options.SelectOperations
	}
	if options.SelectConditions != "" {
		params["selectConditions"] = options.SelectConditions
	}
	if options.SelectFilter != "" {
		params["selectFilter"] = options.SelectFilter
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

	var actions Actions
	err := api.CallWithErrorParse("action.get", params, &actions)
	return actions, err
}

// ActionsGetByID Get actions by specific action IDs
func (api *API) ActionsGetByID(actionIDs []string) (Actions, error) {
	options := ActionGetOptions{
		ActionIDs: actionIDs,
		Output:   "extend",
	}
	return api.ActionsGet(options)
}

// ActionsGetByName Get actions by name
func (api *API) ActionsGetByName(name string) (Actions, error) {
	options := ActionGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.ActionsGet(options)
}

// ActionsGetEnabled Get enabled actions
func (api *API) ActionsGetEnabled() (Actions, error) {
	options := ActionGetOptions{
		Filter: map[string]interface{}{
			"status": ActionStatusEnabled,
		},
		Output: "extend",
	}
	return api.ActionsGet(options)
}

// ActionsGetDisabled Get disabled actions
func (api *API) ActionsGetDisabled() (Actions, error) {
	options := ActionGetOptions{
		Filter: map[string]interface{}{
			"status": ActionStatusDisabled,
		},
		Output: "extend",
	}
	return api.ActionsGet(options)
}

// ActionsGetBySource Get actions by source
func (api *API) ActionsGetBySource(source string) (Actions, error) {
	options := ActionGetOptions{
		Filter: map[string]interface{}{
			"eventsource": source,
		},
		Output: "extend",
	}
	return api.ActionsGet(options)
}

// ActionsGetTriggerActions Get trigger-based actions
func (api *API) ActionsGetTriggerActions() (Actions, error) {
	return api.ActionsGetBySource(ActionSourceTrigger)
}

// ActionsGetDiscoveryActions Get discovery-based actions
func (api *API) ActionsGetDiscoveryActions() (Actions, error) {
	return api.ActionsGetBySource(ActionSourceDiscovery)
}

// ActionsGetAutoregistrationActions Get autoregistration actions
func (api *API) ActionsGetAutoregistrationActions() (Actions, error) {
	return api.ActionsGetBySource(ActionSourceAutoreg)
}

// ActionsGetInternalActions Get internal actions
func (api *API) ActionsGetInternalActions() (Actions, error) {
	return api.ActionsGetBySource(ActionSourceInternal)
}

// ActionCreate Wrapper for action.create
// https://www.zabbix.com/documentation/current/manual/api/reference/action/create
func (api *API) ActionCreate(actions Actions) (result []string, err error) {
	options := ActionCreateOptions{
		Actions: actions,
	}

	response, err := api.CallWithError("action.create", options)
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
			if actionMap, ok := item.(map[string]interface{}); ok {
				if actionid, exists := actionMap["actionids"]; exists {
					if idArray, ok := actionid.([]interface{}); ok && len(idArray) > 0 {
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

// ActionCreateSingle Create a single action
func (api *API) ActionCreateSingle(action Action) (actionID string, err error) {
	actions := Actions{action}
	result, err := api.ActionCreate(actions)
	if len(result) > 0 {
		actionID = result[0]
	}
	return
}

// ActionUpdate Wrapper for action.update
// https://www.zabbix.com/documentation/current/manual/api/reference/action/update
func (api *API) ActionUpdate(actions Actions) (result []string, err error) {
	options := ActionUpdateOptions{
		Actions: actions,
	}

	response, err := api.CallWithError("action.update", options)
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
			if actionMap, ok := item.(map[string]interface{}); ok {
				if actionid, exists := actionMap["actionids"]; exists {
					if idArray, ok := actionid.([]interface{}); ok && len(idArray) > 0 {
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

// ActionUpdateSingle Update a single action
func (api *API) ActionUpdateSingle(action Action) (actionID string, err error) {
	actions := Actions{action}
	result, err := api.ActionUpdate(actions)
	if len(result) > 0 {
		actionID = result[0]
	}
	return
}

// ActionDelete Wrapper for action.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/action/delete
func (api *API) ActionDelete(actionIDs []string) (result []string, err error) {
	options := ActionDeleteOptions{
		ActionIDs: actionIDs,
	}

	response, err := api.CallWithError("action.delete", options)
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

// ActionDeleteSingle Delete a single action
func (api *API) ActionDeleteSingle(actionID string) (err error) {
	_, err = api.ActionDelete([]string{actionID})
	return
}

// ActionExecute Wrapper for action.execute (if supported)
// Note: This method may not be available in all Zabbix versions
func (api *API) ActionExecute(actionID string, params map[string]interface{}) (result map[string]interface{}, err error) {
	executeParams := make(map[string]interface{})
	executeParams["actionid"] = actionID
	
	// Merge additional parameters
	for k, v := range params {
		executeParams[k] = v
	}

	var response interface{}
	err = api.CallWithErrorParse("action.execute", executeParams, &response)
	if err != nil {
		return
	}

	if resultMap, ok := response.(map[string]interface{}); ok {
		result = resultMap
	} else {
		result = make(map[string]interface{})
		result["result"] = response
	}

	return
}

// ActionGetStatistics Get statistics about actions
func (api *API) ActionGetStatistics() (map[string]interface{}, error) {
	// Get all actions
	allActions, err := api.ActionsGet(ActionGetOptions{Output: "extend"})
	if err != nil {
		return nil, err
	}

	stats := make(map[string]interface{})
	
	// Basic counts
	totalActions := len(allActions)
	enabledActions := 0
	disabledActions := 0
	
	// Count by source
	sourceCounts := make(map[string]int)
	
	// Count by status
	statusCounts := make(map[string]int)
	
	// Count operations
	totalOperations := 0
	operationsByType := make(map[string]int)
	
	// Count conditions
	totalConditions := 0
	conditionsByType := make(map[string]int)
	
	for _, action := range allActions {
		// Count by status
		if action.Status == ActionStatusEnabled {
			enabledActions++
			statusCounts["enabled"] = enabledActions
		} else {
			disabledActions++
			statusCounts["disabled"] = disabledActions
		}
		
		// Count by source
		sourceCounts[action.Eventsource]++
		
		// Count operations
		totalOperations += len(action.Commands)
		for _, op := range action.Commands {
			operationsByType[op.OperationType]++
		}
		
		// Count conditions
		totalConditions += len(action.FilterConditions)
		for _, cond := range action.FilterConditions {
			conditionsByType[cond.ConditionType]++
		}
	}
	
	// Build statistics result
	stats["total_actions"] = totalActions
	stats["enabled_actions"] = enabledActions
	stats["disabled_actions"] = disabledActions
	stats["total_operations"] = totalOperations
	stats["total_conditions"] = totalConditions
	
	stats["source_distribution"] = sourceCounts
	stats["status_distribution"] = statusCounts
	stats["operations_by_type"] = operationsByType
	stats["conditions_by_type"] = conditionsByType
	
	// Calculate enabled percentage
	if totalActions > 0 {
		stats["enabled_percentage"] = (float64(enabledActions) / float64(totalActions)) * 100
	} else {
		stats["enabled_percentage"] = 0
	}
	
	return stats, nil
}

// ActionValidate Validate an action configuration
func (api *API) ActionValidate(action Action) (validationErrors []string) {
	validationErrors = []string{}
	
	// Check required fields
	if action.Name == "" {
		validationErrors = append(validationErrors, "Action name is required")
	}
	
	if action.Eventsource == "" {
		validationErrors = append(validationErrors, "Event source is required")
	}
	
	// Validate event source
	validSources := []string{
		ActionSourceTrigger,
		ActionSourceDiscovery,
		ActionSourceAutoreg,
		ActionSourceInternal,
	}
	
	sourceValid := false
	for _, source := range validSources {
		if action.Eventsource == source {
			sourceValid = true
			break
		}
	}
	
	if !sourceValid {
		validationErrors = append(validationErrors, fmt.Sprintf("Invalid event source: %s", action.Eventsource))
	}
	
	// Validate status
	if action.Status != "" && action.Status != ActionStatusEnabled && action.Status != ActionStatusDisabled {
		validationErrors = append(validationErrors, fmt.Sprintf("Invalid status: %s", action.Status))
	}
	
	// Validate operations
	for i, operation := range action.Commands {
		if operation.OperationType == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Operation %d: Operation type is required", i))
		}
		
		// Validate remote command operations
		if operation.OperationType == OperationTypeRemoteCommand || operation.OperationType == OperationTypeCustomScript {
			if operation.Command == "" && operation.ScriptID == "" {
				validationErrors = append(validationErrors, fmt.Sprintf("Operation %d: Command or script ID is required", i))
			}
		}
		
		// Validate message operations
		if operation.OperationType == OperationTypeSendMessage {
			if operation.Subject == "" && operation.Message == "" {
				validationErrors = append(validationErrors, fmt.Sprintf("Operation %d: Subject or message is required", i))
			}
		}
		
		// Validate target operations
		if operation.OperationType == OperationTypeAddToHostGroup || 
		   operation.OperationType == OperationTypeRemoveFromHostGroup ||
		   operation.OperationType == OperationTypeLinkTemplate ||
		   operation.OperationType == OperationTypeUnlinkTemplate {
			if len(operation.HostGroupIDs) == 0 {
				validationErrors = append(validationErrors, fmt.Sprintf("Operation %d: Host group IDs are required", i))
			}
		}
	}
	
	// Validate conditions
	for i, condition := range action.FilterConditions {
		if condition.ConditionType == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Condition %d: Condition type is required", i))
		}
		
		if condition.Value == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("Condition %d: Value is required", i))
		}
	}
	
	return validationErrors
}

// ActionIsEnabled Check if action is enabled
func (action *Action) IsEnabled() bool {
	return action.Status == ActionStatusEnabled
}

// ActionIsDisabled Check if action is disabled
func (action *Action) IsDisabled() bool {
	return action.Status == ActionStatusDisabled
}

// ActionIsTriggerBased Check if action is trigger-based
func (action *Action) IsTriggerBased() bool {
	return action.Eventsource == ActionSourceTrigger
}

// ActionIsDiscoveryBased Check if action is discovery-based
func (action *Action) IsDiscoveryBased() bool {
	return action.Eventsource == ActionSourceDiscovery
}

// ActionIsAutoregBased Check if action is autoregistration-based
func (action *Action) IsAutoregBased() bool {
	return action.Eventsource == ActionSourceAutoreg
}

// ActionIsInternalBased Check if action is internal-based
func (action *Action) IsInternalBased() bool {
	return action.Eventsource == ActionSourceInternal
}

// ActionHasOperations Check if action has operations
func (action *Action) HasOperations() bool {
	return len(action.Commands) > 0
}

// ActionHasRecoveryOperations Check if action has recovery operations
func (action *Action) HasRecoveryOperations() bool {
	return len(action.RecoveryCommands) > 0
}

// ActionHasConditions Check if action has conditions
func (action *Action) HasConditions() bool {
	return len(action.FilterConditions) > 0
}

// ActionHasCommands Check if action has remote commands
func (action *Action) HasCommands() bool {
	for _, operation := range action.Commands {
		if operation.OperationType == OperationTypeRemoteCommand || operation.OperationType == OperationTypeCustomScript {
			return true
		}
	}
	return false
}

// ActionHasMessages Check if action has message operations
func (action *Action) HasMessages() bool {
	for _, operation := range action.Commands {
		if operation.OperationType == OperationTypeSendMessage {
			return true
		}
	}
	return false
}

// ActionEnable Enable an action
func (api *API) ActionEnable(actionID string) (err error) {
	actions, err := api.ActionsGetByID([]string{actionID})
	if err != nil {
		return
	}
	
	if len(actions) == 0 {
		return fmt.Errorf("Action not found: %s", actionID)
	}
	
	actions[0].Status = ActionStatusEnabled
	_, err = api.ActionUpdateSingle(actions[0])
	return
}

// ActionDisable Disable an action
func (api *API) ActionDisable(actionID string) (err error) {
	actions, err := api.ActionsGetByID([]string{actionID})
	if err != nil {
		return
	}
	
	if len(actions) == 0 {
		return fmt.Errorf("Action not found: %s", actionID)
	}
	
	actions[0].Status = ActionStatusDisabled
	_, err = api.ActionUpdateSingle(actions[0])
	return
}