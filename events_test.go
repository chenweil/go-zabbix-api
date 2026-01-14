package zabbix

import (
	"testing"
	"time"
)

func TestEventGetOptions(t *testing.T) {
	// Test default values
	opts := EventGetOptions{
		EventIDs: []string{"12345"},
		TimeFrom: 1640995200,
		TimeTill: 1640998800,
	}

	if len(opts.EventIDs) != 1 {
		t.Errorf("Expected 1 event ID, got %d", len(opts.EventIDs))
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

func TestEvent(t *testing.T) {
	// Test Event creation and JSON marshaling
	event := Event{
		EventID:     "12345",
		Source:      "0",
		Object:      "0",
		ObjectID:    "67890",
		Clock:       1640995200,
		ns:          123456,
		Value:       1,
		Acknowledged: 0,
		Flags:       "0",
		Severity:    "3",
		Tags: []Tag{
			{Tag: "service", Value: "web"},
		},
	}

	expectedJSON := `{"eventid":"12345","source":"0","object":"0","objectid":"67890","clock":1640995200,"ns":123456,"value":1,"acknowledged":0,"flags":"0","severity":"3","tags":[{"tag":"service","value":"web"}]}`
	
	jsonData, err := json.Marshal(event)
	if err != nil {
		t.Errorf("Failed to marshal Event: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestEventAcknowledgement(t *testing.T) {
	// Test EventAcknowledgement creation and JSON marshaling
	ack := EventAcknowledgement{
		EventID:  "12345",
		Message:  "Acknowledged by admin",
		Action:   2,
		Severity: "3",
		KeepAlert: false,
	}

	expectedJSON := `{"eventid":"12345","message":"Acknowledged by admin","action":2,"severity":"3","keep_alert":false}`
	
	jsonData, err := json.Marshal(ack)
	if err != nil {
		t.Errorf("Failed to marshal EventAcknowledgement: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestEventIsProblem(t *testing.T) {
	// Test Event.IsProblem() method
	problemEvent := Event{Value: 1}
	okEvent := Event{Value: 0}

	if !problemEvent.IsProblem() {
		t.Errorf("Expected problem event to be detected as problem")
	}

	if okEvent.IsProblem() {
		t.Errorf("Expected OK event to NOT be detected as problem")
	}
}

func TestEventIsOK(t *testing.T) {
	// Test Event.IsOK() method
	problemEvent := Event{Value: 1}
	okEvent := Event{Value: 0}

	if okEvent.IsOK() {
		t.Errorf("Expected OK event to be detected as OK")
	}

	if problemEvent.IsOK() {
		t.Errorf("Expected problem event to NOT be detected as OK")
	}
}

func TestEventIsAcknowledged(t *testing.T) {
	// Test Event.IsAcknowledged() method
	ackEvent := Event{Acknowledged: 1}
	unackEvent := Event{Acknowledged: 0}

	if !ackEvent.IsAcknowledged() {
		t.Errorf("Expected acknowledged event to be detected as acknowledged")
	}

	if unackEvent.IsAcknowledged() {
		t.Errorf("Expected unacknowledged event to NOT be detected as acknowledged")
	}
}

func TestEventIsCritical(t *testing.T) {
	// Test Event.IsCritical() method
	criticalEvent := Event{Severity: EventSeverityDisaster}
	highEvent := Event{Severity: EventSeverityHigh}
	warningEvent := Event{Severity: EventSeverityWarning}

	if !criticalEvent.IsCritical() {
		t.Errorf("Expected disaster event to be detected as critical")
	}

	if highEvent.IsCritical() {
		t.Errorf("Expected high event to NOT be detected as critical")
	}

	if warningEvent.IsCritical() {
		t.Errorf("Expected warning event to NOT be detected as critical")
	}
}

func TestEventIsHighSeverity(t *testing.T) {
	// Test Event.IsHighSeverity() method
	disasterEvent := Event{Severity: EventSeverityDisaster}
	highEvent := Event{Severity: EventSeverityHigh}
	averageEvent := Event{Severity: EventSeverityAverage}
	warningEvent := Event{Severity: EventSeverityWarning}

	if !disasterEvent.IsHighSeverity() {
		t.Errorf("Expected disaster event to be detected as high severity")
	}

	if !highEvent.IsHighSeverity() {
		t.Errorf("Expected high event to be detected as high severity")
	}

	if averageEvent.IsHighSeverity() {
		t.Errorf("Expected average event to NOT be detected as high severity")
	}

	if warningEvent.IsHighSeverity() {
		t.Errorf("Expected warning event to NOT be detected as high severity")
	}
}

func TestEventIsWarning(t *testing.T) {
	// Test Event.IsWarning() method
	warningEvent := Event{Severity: EventSeverityWarning}
	disasterEvent := Event{Severity: EventSeverityDisaster}

	if !warningEvent.IsWarning() {
		t.Errorf("Expected warning event to be detected as warning")
	}

	if disasterEvent.IsWarning() {
		t.Errorf("Expected disaster event to NOT be detected as warning")
	}
}

func TestEventAge(t *testing.T) {
	// Test Event.Age() method
	// Create event with timestamp from 1 hour ago
	hourAgo := int(time.Now().Unix()) - 3600
	event := Event{Clock: hourAgo}

	age := event.Age()
	
	// Age should be approximately 3600 seconds (1 hour)
	if age < 3500 || age > 3700 {
		t.Errorf("Expected age around 3600 seconds, got %d", age)
	}
}

func TestEventAgeFormatted(t *testing.T) {
	// Test Event.AgeFormatted() method
	
	// Test minutes
	minuteAgo := int(time.Now().Unix()) - 90
	event := Event{Clock: minuteAgo}
	formatted := event.AgeFormatted()
	
	if formatted != "1m" && formatted != "2m" {
		t.Errorf("Expected age to be formatted as minutes, got %s", formatted)
	}
	
	// Test hours
	hourAgo := int(time.Now().Unix()) - 7200 // 2 hours
	event = Event{Clock: hourAgo}
	formatted = event.AgeFormatted()
	
	if formatted != "2h 0m" && formatted != "2h" {
		t.Errorf("Expected age to be formatted as hours, got %s", formatted)
	}
}

func TestEventConstants(t *testing.T) {
	// Test event constants
	tests := []struct {
		constant string
		expected string
	}{
		{EventSeverityNotClassified, "0"},
		{EventSeverityInformation, "1"},
		{EventSeverityWarning, "2"},
		{EventSeverityAverage, "3"},
		{EventSeverityHigh, "4"},
		{EventSeverityDisaster, "5"},
		{EventSourceTrigger, "0"},
		{EventSourceDiscovery, "1"},
		{EventSourceAutoreg, "2"},
		{EventSourceInternal, "3"},
		{EventSourceHTTPAgent, "4"},
		{EventValueOK, "0"},
		{EventValueProblem, "1"},
		{ActionNone, "0"},
		{ActionCloseProblem, "1"},
		{ActionAcknowledge, "2"},
		{ActionCloseEvent, "3"},
		{EventSortClock, "clock"},
		{EventSortEventID, "eventid"},
		{EventSortSeverity, "severity"},
		{EventSortAcknowledged, "acknowledged"},
	}

	for _, test := range tests {
		if test.constant != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, test.constant)
		}
	}
}

func TestMockEventsAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test that methods exist and return appropriate types
	opts := EventGetOptions{
		EventIDs: []string{"12345"},
		Output:   "extend",
		Limit:    10,
	}

	// These calls will fail without a real Zabbix server, but we can verify the method signatures
	_, err := api.EventsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetByTrigger([]string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetByHost([]string{"11111"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetProblems()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetOK()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetHighSeverity()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetCritical()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetWarning()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetRecent(24)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetUnacknowledged()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetAcknowledged()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetTriggerEvents()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test acknowledgement methods
	_, err = api.EventsAcknowledge([]string{"12345"}, "Test acknowledgement")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsCloseProblem([]string{"12345"}, "Closing problem")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics methods
	_, err = api.EventsGetStatistics(1640995200, 1640998800)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsGetSummary(24)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.EventsCheckHealth()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkEventMarshaling(b *testing.B) {
	event := Event{
		EventID:     "12345",
		Source:      "0",
		Object:      "0",
		ObjectID:    "67890",
		Clock:       1640995200,
		ns:          123456,
		Value:       1,
		Acknowledged: 0,
		Flags:       "0",
		Severity:    "3",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(event)
	}
}

func BenchmarkEventAgeCalculation(b *testing.B) {
	event := Event{Clock: 1640995200}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		event.Age()
	}
}

func BenchmarkEventAgeFormatted(b *testing.B) {
	event := Event{Clock: 1640995200}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		event.AgeFormatted()
	}
}

// Test integration scenarios
func TestEventsIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test complex filtering
	opts := EventGetOptions{
		HostIDs:    []string{"12345", "67890"},
		TriggerIDs: []string{"11111"},
		TimeFrom:   1640995200,
		TimeTill:   1640998800,
		Source:     []int{0}, // Trigger events
		Severity:   []string{"4", "5"}, // High and Disaster
		Output:     "extend",
		Limit:      100,
		SortField:  "clock",
		SortOrder:  "DESC",
	}

	_, err := api.EventsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test filtering by host and trigger
	_, err = api.EventsFilterByHostAndTrigger([]string{"12345"}, []string{"67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test filtering by severity and time
	_, err = api.EventsFilterBySeverityAndTime([]string{"5"}, 1640995200, 1640998800)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test filtering by time range with criteria
	criteria := map[string]interface{}{
		"acknowledged": "0",
		"severity":     "4",
	}
	_, err = api.EventsFilterByTimeRange(1640995200, 1640998800, criteria)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}