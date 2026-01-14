package zabbix

import (
	"testing"
)

func TestActionGetOptions(t *testing.T) {
	// Test default values
	opts := ActionGetOptions{
		ActionIDs: []string{"12345"},
	}

	if len(opts.ActionIDs) != 1 {
		t.Errorf("Expected 1 action ID, got %d", len(opts.ActionIDs))
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

func TestAction(t *testing.T) {
	// Test Action creation and JSON marshaling
	action := Action{
		ActionID:      "12345",
		Name:          "Test Action",
		Eventsource:   ActionSourceTrigger,
		Status:        ActionStatusEnabled,
		DefLongdata:   "Problem: {EVENT.NAME}\\nHost: {HOST.NAME}\\nTime: {EVENT.DATE} {EVENT.TIME}",
		DefShortdata:  "Problem on {HOST.NAME}",
	}

	expectedJSON := `{"actionid":"12345","name":"Test Action","eventsource":"0","def_longdata":"Problem: {EVENT.NAME}\\nHost: {HOST.NAME}\\nTime: {EVENT.DATE} {EVENT.TIME}","def_shortdata":"Problem on {HOST.NAME}","status":"0"}`
	
	jsonData, err := json.Marshal(action)
	if err != nil {
		t.Errorf("Failed to marshal Action: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestOperation(t *testing.T) {
	// Test Operation creation and JSON marshaling
	operation := Operation{
		OperationID:  "67890",
		ActionID:    "12345",
		OperationType: OperationTypeSendMessage,
		EscPeriod:   "60s",
		EscStepFrom: "0",
		EscStepTo:   "10",
		Subject:     "Problem detected",
		Message:     "Please check the system",
		MediaTypeID: "1",
	}

	expectedJSON := `{"operationid":"67890","actionid":"12345","operationtype":"0","esc_period":"60s","esc_step_from":"0","esc_step_to":"10","subject":"Problem detected","message":"Please check the system","mediatypeid":"1"}`
	
	jsonData, err := json.Marshal(operation)
	if err != nil {
		t.Errorf("Failed to marshal Operation: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestCondition(t *testing.T) {
	// Test Condition creation and JSON marshaling
	condition := Condition{
		ConditionID:  "11111",
		ActionID:     "12345",
		ConditionType: ConditionTypeTriggerSeverity,
		Operator:     ConditionOperatorEquals,
		Value:        "3", // Average severity
		FormulaID:    "A",
	}

	expectedJSON := `{"conditionid":"11111","actionid":"12345","conditiontype":"2","operator":"0","value":"3","formulaid":"A"}`
	
	jsonData, err := json.Marshal(condition)
	if err != nil {
		t.Errorf("Failed to marshal Condition: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestActionConstants(t *testing.T) {
	// Test action constants
	tests := []struct {
		constant string
		expected string
	}{
		// Action sources
		{ActionSourceTrigger, "0"},
		{ActionSourceDiscovery, "1"},
		{ActionSourceAutoreg, "2"},
		{ActionSourceInternal, "3"},
		
		// Action status
		{ActionStatusEnabled, "0"},
		{ActionStatusDisabled, "1"},
		
		// Recovery modes
		{ActionRecoveryModeNone, "0"},
		{ActionRecoveryModeOperations, "1"},
		{ActionRecoveryModeBoth, "2"},
		
		// Operation types
		{OperationTypeSendMessage, "0"},
		{OperationTypeRemoteCommand, "1"},
		{OperationTypeAddHost, "2"},
		{OperationTypeRemoveHost, "3"},
		{OperationTypeAddToHostGroup, "4"},
		{OperationTypeRemoveFromHostGroup, "5"},
		{OperationTypeLinkTemplate, "6"},
		{OperationTypeUnlinkTemplate, "7"},
		{OperationTypeEnableHost, "8"},
		{OperationTypeDisableHost, "9"},
		{OperationTypeCustomScript, "10"},
		
		// Condition types
		{ConditionTypeTrigger, "0"},
		{ConditionTypeTriggerName, "1"},
		{ConditionTypeTriggerSeverity, "2"},
		{ConditionTypeTriggerValue, "3"},
		{ConditionTypeMaintenanceStatus, "12"},
		{ConditionTypeEventTag, "25"},
		{ConditionTypeEventTagValue, "26"},
		{ConditionTypeHost, "10"},
		{ConditionTypeHostGroup, "11"},
		{ConditionTypeTemplate, "13"},
		{ConditionTypeApplication, "15"},
		{ConditionTypeTriggerProxy, "16"},
		{ConditionTypeHostProxy, "17"},
		{ConditionTypeEventType, "18"},
		{ConditionTypeHostMetadata, "21"},
		{ConditionTypeHostInventory, "22"},
		{ConditionTypeEventTagValueEx, "27"},
		
		// Condition operators
		{ConditionOperatorEquals, "0"},
		{ConditionOperatorNotEquals, "1"},
		{ConditionOperatorContains, "2"},
		{ConditionOperatorNotContains, "3"},
		{ConditionOperatorInRange, "4"},
		{ConditionOperatorNotInRange, "5"},
		{ConditionOperatorMatchesPattern, "6"},
		{ConditionOperatorNotMatchesPattern, "7"},
		
		// Command types
		{CommandTypeIPMI, "0"},
		{CommandTypeSSH, "1"},
		{CommandTypeTelnet, "2"},
		{CommandTypeCustomScript, "3"},
		{CommandTypeGlobalScript, "4"},
		
		// Message delivery methods
		{OpmessageCmdUserMedia, "0"},
		{OpmessageCmdAllMedia, "1"},
		
		// Target operation modes
		{OperationModeCurrentHost, "0"},
		{OperationModeTargetHost, "1"},
		{OperationModeHostGroup, "2"},
		
		// Execute on host
		{ExecuteOnZabbixAgent, "0"},
		{ExecuteOnZabbixServer, "1"},
		{ExecuteOnZabbixProxy, "2"},
		
		// Problem duration modes
		{ProblemDurationModeNone, "0"},
		{ProblemDurationModeTime, "1"},
		{ProblemDurationModeCustom, "2"},
		
		// Inventory mode
		{InventoryModeManual, "0"},
		{InventoryModeAutomatic, "1"},
	}

	for _, test := range tests {
		if test.constant != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.constant)
		}
	}
}

func TestActionIsEnabled(t *testing.T) {
	// Test Action.IsEnabled() method
	enabledAction := Action{Status: ActionStatusEnabled}
	disabledAction := Action{Status: ActionStatusDisabled}
	noStatusAction := Action{Status: ""}

	if !enabledAction.IsEnabled() {
		t.Errorf("Expected enabled action to be detected as enabled")
	}

	if disabledAction.IsEnabled() {
		t.Errorf("Expected disabled action to NOT be detected as enabled")
	}

	if noStatusAction.IsEnabled() {
		t.Errorf("Expected action with no status to NOT be detected as enabled")
	}
}

func TestActionIsDisabled(t *testing.T) {
	// Test Action.IsDisabled() method
	enabledAction := Action{Status: ActionStatusEnabled}
	disabledAction := Action{Status: ActionStatusDisabled}
	noStatusAction := Action{Status: ""}

	if disabledAction.IsDisabled() {
		t.Errorf("Expected disabled action to be detected as disabled")
	}

	if enabledAction.IsDisabled() {
		t.Errorf("Expected enabled action to NOT be detected as disabled")
	}

	if noStatusAction.IsDisabled() {
		t.Errorf("Expected action with no status to NOT be detected as disabled")
	}
}

func TestActionIsTriggerBased(t *testing.T) {
	// Test Action.IsTriggerBased() method
	triggerAction := Action{Eventsource: ActionSourceTrigger}
	discoveryAction := Action{Eventsource: ActionSourceDiscovery}
	autoregAction := Action{Eventsource: ActionSourceAutoreg}
	internalAction := Action{Eventsource: ActionSourceInternal}
	noSourceAction := Action{Eventsource: ""}

	if !triggerAction.IsTriggerBased() {
		t.Errorf("Expected trigger action to be detected as trigger-based")
	}

	if discoveryAction.IsTriggerBased() {
		t.Errorf("Expected discovery action to NOT be detected as trigger-based")
	}

	if autoregAction.IsTriggerBased() {
		t.Errorf("Expected autoreg action to NOT be detected as trigger-based")
	}

	if internalAction.IsTriggerBased() {
		t.Errorf("Expected internal action to NOT be detected as trigger-based")
	}

	if noSourceAction.IsTriggerBased() {
		t.Errorf("Expected action with no source to NOT be detected as trigger-based")
	}
}

func TestActionIsDiscoveryBased(t *testing.T) {
	// Test Action.IsDiscoveryBased() method
	discoveryAction := Action{Eventsource: ActionSourceDiscovery}
	triggerAction := Action{Eventsource: ActionSourceTrigger}

	if !discoveryAction.IsDiscoveryBased() {
		t.Errorf("Expected discovery action to be detected as discovery-based")
	}

	if triggerAction.IsDiscoveryBased() {
		t.Errorf("Expected trigger action to NOT be detected as discovery-based")
	}
}

func TestActionIsAutoregBased(t *testing.T) {
	// Test Action.IsAutoregBased() method
	autoregAction := Action{Eventsource: ActionSourceAutoreg}
	triggerAction := Action{Eventsource: ActionSourceTrigger}

	if !autoregAction.IsAutoregBased() {
		t.Errorf("Expected autoreg action to be detected as autoregistration-based")
	}

	if triggerAction.IsAutoregBased() {
		t.Errorf("Expected trigger action to NOT be detected as autoregistration-based")
	}
}

func TestActionIsInternalBased(t *testing.T) {
	// Test Action.IsInternalBased() method
	internalAction := Action{Eventsource: ActionSourceInternal}
	triggerAction := Action{Eventsource: ActionSourceTrigger}

	if !internalAction.IsInternalBased() {
		t.Errorf("Expected internal action to be detected as internal-based")
	}

	if triggerAction.IsInternalBased() {
		t.Errorf("Expected trigger action to NOT be detected as internal-based")
	}
}

func TestActionHasOperations(t *testing.T) {
	// Test Action.HasOperations() method
	actionWithOps := Action{
		Commands: []Operation{
			{OperationType: OperationTypeSendMessage},
		},
	}
	actionWithoutOps := Action{
		Commands: []Operation{},
	}
	actionNoOps := Action{}

	if !actionWithOps.HasOperations() {
		t.Errorf("Expected action with operations to have operations")
	}

	if actionWithoutOps.HasOperations() {
		t.Errorf("Expected action without operations to NOT have operations")
	}

	if actionNoOps.HasOperations() {
		t.Errorf("Expected action with nil operations to NOT have operations")
	}
}

func TestActionHasRecoveryOperations(t *testing.T) {
	// Test Action.HasRecoveryOperations() method
	actionWithRecovery := Action{
		RecoveryCommands: []Operation{
			{OperationType: OperationTypeSendMessage},
		},
	}
	actionWithoutRecovery := Action{
		RecoveryCommands: []Operation{},
	}
	actionNoRecovery := Action{}

	if !actionWithRecovery.HasRecoveryOperations() {
		t.Errorf("Expected action with recovery operations to have recovery operations")
	}

	if actionWithoutRecovery.HasRecoveryOperations() {
		t.Errorf("Expected action without recovery operations to NOT have recovery operations")
	}

	if actionNoRecovery.HasRecoveryOperations() {
		t.Errorf("Expected action with nil recovery operations to NOT have recovery operations")
	}
}

func TestActionHasConditions(t *testing.T) {
	// Test Action.HasConditions() method
	actionWithConds := Action{
		FilterConditions: []Condition{
			{ConditionType: ConditionTypeTriggerSeverity},
		},
	}
	actionWithoutConds := Action{
		FilterConditions: []Condition{},
	}
	actionNoConds := Action{}

	if !actionWithConds.HasConditions() {
		t.Errorf("Expected action with conditions to have conditions")
	}

	if actionWithoutConds.HasConditions() {
		t.Errorf("Expected action without conditions to NOT have conditions")
	}

	if actionNoConds.HasConditions() {
		t.Errorf("Expected action with nil conditions to NOT have conditions")
	}
}

func TestActionHasCommands(t *testing.T) {
	// Test Action.HasCommands() method
	actionWithCommands := Action{
		Commands: []Operation{
			{OperationType: OperationTypeRemoteCommand},
			{OperationType: OperationTypeSendMessage},
		},
	}
	actionWithCustomScript := Action{
		Commands: []Operation{
			{OperationType: OperationTypeCustomScript},
		},
	}
	actionWithoutCommands := Action{
		Commands: []Operation{
			{OperationType: OperationTypeSendMessage},
		},
	}

	if !actionWithCommands.HasCommands() {
		t.Errorf("Expected action with remote commands to have commands")
	}

	if !actionWithCustomScript.HasCommands() {
		t.Errorf("Expected action with custom script to have commands")
	}

	if actionWithoutCommands.HasCommands() {
		t.Errorf("Expected action with only messages to NOT have commands")
	}
}

func TestActionHasMessages(t *testing.T) {
	// Test Action.HasMessages() method
	actionWithMessages := Action{
		Commands: []Operation{
			{OperationType: OperationTypeSendMessage},
			{OperationType: OperationTypeRemoteCommand},
		},
	}
	actionWithoutMessages := Action{
		Commands: []Operation{
			{OperationType: OperationTypeRemoteCommand},
		},
	}

	if !actionWithMessages.HasMessages() {
		t.Errorf("Expected action with message operations to have messages")
	}

	if actionWithoutMessages.HasMessages() {
		t.Errorf("Expected action without message operations to NOT have messages")
	}
}

func TestActionValidate(t *testing.T) {
	// Test Action validation
	api := NewAPI(Config{})
	
	// Valid action
	validAction := Action{
		Name:         "Valid Action",
		Eventsource:  ActionSourceTrigger,
		Status:       ActionStatusEnabled,
		Commands: []Operation{
			{
				OperationType: OperationTypeSendMessage,
				Subject:      "Test Subject",
				Message:      "Test Message",
			},
		},
	}
	
	errors := api.ActionValidate(validAction)
	if len(errors) > 0 {
		t.Errorf("Expected no validation errors for valid action, got: %v", errors)
	}
	
	// Invalid action - missing name
	invalidAction1 := Action{
		Eventsource: ActionSourceTrigger,
		Status:      ActionStatusEnabled,
	}
	
	errors = api.ActionValidate(invalidAction1)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for action without name")
	}
	
	// Invalid action - missing event source
	invalidAction2 := Action{
		Name:    "Test Action",
		Status:  ActionStatusEnabled,
	}
	
	errors = api.ActionValidate(invalidAction2)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for action without event source")
	}
	
	// Invalid action - invalid status
	invalidAction3 := Action{
		Name:        "Test Action",
		Eventsource: ActionSourceTrigger,
		Status:      "invalid_status",
	}
	
	errors = api.ActionValidate(invalidAction3)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for action with invalid status")
	}
	
	// Invalid operation - missing command content
	invalidAction4 := Action{
		Name:        "Test Action",
		Eventsource: ActionSourceTrigger,
		Status:      ActionStatusEnabled,
		Commands: []Operation{
			{
				OperationType: OperationTypeRemoteCommand,
				// Missing Command or ScriptID
			},
		},
	}
	
	errors = api.ActionValidate(invalidAction4)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for action with incomplete remote command")
	}
	
	// Invalid condition - missing value
	invalidAction5 := Action{
		Name:        "Test Action",
		Eventsource: ActionSourceTrigger,
		Status:      ActionStatusEnabled,
		FilterConditions: []Condition{
			{
				ConditionType: ConditionTypeTriggerSeverity,
				// Missing Value
			},
		},
	}
	
	errors = api.ActionValidate(invalidAction5)
	if len(errors) == 0 {
		t.Errorf("Expected validation errors for action with condition missing value")
	}
}

func TestMockActionsAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test that methods exist and return appropriate types
	opts := ActionGetOptions{
		ActionIDs: []string{"12345"},
		Output:   "extend",
		Limit:    10,
	}

	// These calls will fail without a real Zabbix server, but we can verify the method signatures
	_, err := api.ActionsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetByName("Test Action")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetEnabled()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetDisabled()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetTriggerActions()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetDiscoveryActions()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetAutoregistrationActions()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionsGetInternalActions()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test CRUD operations
	action := Action{
		Name:         "Test Action",
		Eventsource:  ActionSourceTrigger,
		Status:       ActionStatusEnabled,
		DefLongdata:  "Test long data",
		DefShortdata: "Test short data",
	}
	
	_, err = api.ActionCreateSingle(action)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ActionUpdateSingle(action)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ActionDeleteSingle("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ActionEnable("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ActionDisable("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics
	_, err = api.ActionGetStatistics()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test execute
	params := map[string]interface{}{
		"test": "value",
	}
	_, err = api.ActionExecute("12345", params)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkActionMarshaling(b *testing.B) {
	action := Action{
		ActionID:     "12345",
		Name:         "Test Action",
		Eventsource:  ActionSourceTrigger,
		Status:       ActionStatusEnabled,
		DefLongdata:  "Problem: {EVENT.NAME}",
		DefShortdata: "Problem detected",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(action)
	}
}

func BenchmarkOperationMarshaling(b *testing.B) {
	operation := Operation{
		OperationID:  "67890",
		OperationType: OperationTypeSendMessage,
		EscPeriod:   "60s",
		Subject:     "Problem detected",
		Message:     "Please check",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(operation)
	}
}

func BenchmarkConditionMarshaling(b *testing.B) {
	condition := Condition{
		ConditionID:  "11111",
		ConditionType: ConditionTypeTriggerSeverity,
		Operator:     ConditionOperatorEquals,
		Value:        "3",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(condition)
	}
}

// Test integration scenarios
func TestActionsIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test complex filtering
	opts := ActionGetOptions{
		Filter: map[string]interface{}{
			"eventsource": ActionSourceTrigger,
			"status":     ActionStatusEnabled,
		},
		Output: "extend",
		Limit:  100,
	}

	_, err := api.ActionsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test action creation with operations and conditions
	action := Action{
		Name:         "Integration Test Action",
		Eventsource:  ActionSourceTrigger,
		Status:       ActionStatusEnabled,
		DefLongdata:  "Problem: {EVENT.NAME}\\nHost: {HOST.NAME}",
		DefShortdata: "Problem on {HOST.NAME}",
		Commands: []Operation{
			{
				OperationType: OperationTypeSendMessage,
				EscPeriod:   "0s",
				EscStepFrom: "0",
				EscStepTo:   "0",
				Subject:     "Alert: {EVENT.NAME}",
				Message:     "Problem on {HOST.NAME}\\nSeverity: {TRIGGER.SEVERITY}\\nTime: {EVENT.DATE} {EVENT.TIME}",
			},
			{
				OperationType: OperationTypeRemoteCommand,
				EscPeriod:   "60s",
				EscStepFrom: "5",
				EscStepTo:   "10",
				Command:     "/usr/local/bin/restart-service.sh {HOST.NAME}",
				CommandType: CommandTypeCustomScript,
			},
		},
		FilterConditions: []Condition{
			{
				ConditionType: ConditionTypeTriggerSeverity,
				Operator:     ConditionOperatorEquals,
				Value:        "3", // Average or higher
			},
			{
				ConditionType: ConditionTypeHostGroup,
				Operator:     ConditionOperatorInRange,
				Value:        "5", // Group ID
			},
		},
	}
	
	_, err = api.ActionCreateSingle(action)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}