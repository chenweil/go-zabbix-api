package zabbix

import (
	"testing"
	"time"
)

func TestHistoryValueTypeString(t *testing.T) {
	tests := []struct {
		ht       HistoryValueType
		expected string
	}{
		{HistoryFloat, "0"},
		{HistoryString, "1"},
		{HistoryLog, "2"},
		{HistoryUnsigned, "3"},
		{HistoryText, "4"},
		{HistoryBinary, "5"},
	}

	for _, test := range tests {
		if got := string(test.ht); got != test.expected {
			t.Errorf("HistoryValueType %v, expected %v, got %v", test.ht, test.expected, got)
		}
	}
}

func TestHistoryGetOptions(t *testing.T) {
	// Test default values
	opts := HistoryGetOptions{
		ItemIDs: []string{"12345"},
		History: HistoryFloat,
	}

	if len(opts.ItemIDs) != 1 {
		t.Errorf("Expected 1 item ID, got %d", len(opts.ItemIDs))
	}

	if opts.SortField != "" {
		t.Errorf("Expected empty sort field, got %v", opts.SortField)
	}

	if opts.SortOrder != "" {
		t.Errorf("Expected empty sort order, got %v", opts.SortOrder)
	}
}

func TestHistoryData(t *testing.T) {
	// Test HistoryData creation and JSON marshaling
	hd := HistoryData{
		ItemID: "12345",
		Clock:  1640995200,
		Value:  "test_value",
		NS:     123456,
	}

	expectedJSON := `{"itemid":"12345","clock":1640995200,"value":"test_value","ns":123456}`
	
	jsonData, err := json.Marshal(hd)
	if err != nil {
		t.Errorf("Failed to marshal HistoryData: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestHistoryFloat(t *testing.T) {
	// Test HistoryFloat creation and JSON marshaling
	hf := HistoryFloat{
		ItemID: "12345",
		Clock:  1640995200,
		Value:  123.45,
		NS:     123456,
	}

	expectedJSON := `{"itemid":"12345","clock":1640995200,"value":123.45,"ns":123456}`
	
	jsonData, err := json.Marshal(hf)
	if err != nil {
		t.Errorf("Failed to marshal HistoryFloat: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestHistoryString(t *testing.T) {
	// Test HistoryString creation and JSON marshaling
	hs := HistoryString{
		ItemID: "12345",
		Clock:  1640995200,
		Value:  "test_string",
		NS:     123456,
	}

	expectedJSON := `{"itemid":"12345","clock":1640995200,"value":"test_string","ns":123456}`
	
	jsonData, err := json.Marshal(hs)
	if err != nil {
		t.Errorf("Failed to marshal HistoryString: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestHistoryLog(t *testing.T) {
	// Test HistoryLog creation and JSON marshaling
	hl := HistoryLog{
		ItemID:     "12345",
		Clock:      1640995200,
		Value:      "log entry",
		NS:         123456,
		Level:      "INFO",
		Severity:   "3",
		Source:     "application",
		LogEventID: "456",
		Timestamp:  "2021-12-31 12:00:00",
	}

	expectedJSON := `{"itemid":"12345","clock":1640995200,"value":"log entry","ns":123456,"severity":"3","source":"application","logeventid":"456","timestamp":"2021-12-31 12:00:00"}`
	
	jsonData, err := json.Marshal(hl)
	if err != nil {
		t.Errorf("Failed to marshal HistoryLog: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestHistorySliceToFloatHistory(t *testing.T) {
	// Test conversion from generic HistorySlice to HistoryFloatSlice
	history := HistorySlice{
		{ItemID: "123", Clock: 1000, ValueNum: 10.5, ValueStr: "", NS: 1},
		{ItemID: "123", Clock: 1001, ValueNum: 20.3, ValueStr: "", NS: 2},
		{ItemID: "123", Clock: 1002, ValueNum: 0, ValueStr: "not_numeric", NS: 3},
		{ItemID: "456", Clock: 1000, ValueNum: 30.7, ValueStr: "", NS: 4},
	}

	floatHistory, err := history.ToFloatHistory()
	if err != nil {
		t.Errorf("Failed to convert to float history: %v", err)
	}

	// Should only include records with numeric values
	expectedCount := 3
	if len(floatHistory) != expectedCount {
		t.Errorf("Expected %d float history records, got %d", expectedCount, len(floatHistory))
	}

	// Check specific values
	if floatHistory[0].Value != 10.5 {
		t.Errorf("Expected first value 10.5, got %v", floatHistory[0].Value)
	}

	if floatHistory[1].Value != 20.3 {
		t.Errorf("Expected second value 20.3, got %v", floatHistory[1].Value)
	}

	if floatHistory[2].Value != 30.7 {
		t.Errorf("Expected third value 30.7, got %v", floatHistory[2].Value)
	}
}

func TestHistorySliceToStringHistory(t *testing.T) {
	// Test conversion from generic HistorySlice to HistoryStringSlice
	history := HistorySlice{
		{ItemID: "123", Clock: 1000, ValueNum: 0, ValueStr: "text1", NS: 1},
		{ItemID: "123", Clock: 1001, ValueNum: 0, ValueStr: "text2", NS: 2},
		{ItemID: "123", Clock: 1002, ValueNum: 0, ValueStr: "", NS: 3},
		{ItemID: "456", Clock: 1000, ValueNum: 0, ValueStr: "text3", NS: 4},
	}

	stringHistory, err := history.ToStringHistory()
	if err != nil {
		t.Errorf("Failed to convert to string history: %v", err)
	}

	// Should only include records with string values
	expectedCount := 3
	if len(stringHistory) != expectedCount {
		t.Errorf("Expected %d string history records, got %d", expectedCount, len(stringHistory))
	}

	// Check specific values
	if stringHistory[0].Value != "text1" {
		t.Errorf("Expected first value 'text1', got %v", stringHistory[0].Value)
	}

	if stringHistory[1].Value != "text2" {
		t.Errorf("Expected second value 'text2', got %v", stringHistory[1].Value)
	}

	if stringHistory[2].Value != "text3" {
		t.Errorf("Expected third value 'text3', got %v", stringHistory[2].Value)
	}
}

func TestHistorySliceToLogHistory(t *testing.T) {
	// Test conversion from generic HistorySlice to HistoryLogSlice
	history := HistorySlice{
		{ItemID: "123", Clock: 1000, Value: "log1", NS: 1, Level: "INFO", Source: "app1"},
		{ItemID: "123", Clock: 1001, Value: "log2", NS: 2, Level: "ERROR", Source: "app2"},
		{ItemID: "123", Clock: 1002, Value: "log3", NS: 3, Level: "", Source: ""},
		{ItemID: "456", Clock: 1000, Value: "log4", NS: 4, Level: "DEBUG", Source: "app3"},
	}

	logHistory, err := history.ToLogHistory()
	if err != nil {
		t.Errorf("Failed to convert to log history: %v", err)
	}

	// Should only include records with log-related data
	expectedCount := 3
	if len(logHistory) != expectedCount {
		t.Errorf("Expected %d log history records, got %d", expectedCount, len(logHistory))
	}

	// Check specific values
	if logHistory[0].Level != "INFO" {
		t.Errorf("Expected first level 'INFO', got %v", logHistory[0].Level)
	}

	if logHistory[1].Source != "app2" {
		t.Errorf("Expected second source 'app2', got %v", logHistory[1].Source)
	}

	if logHistory[2].Level != "DEBUG" {
		t.Errorf("Expected third level 'DEBUG', got %v", logHistory[2].Level)
	}
}

func TestStatisticsCalculation(t *testing.T) {
	// Test statistics calculation helper functions
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	
	min := calculateMin(values)
	if min != 1.0 {
		t.Errorf("Expected min 1.0, got %v", min)
	}

	max := calculateMax(values)
	if max != 5.0 {
		t.Errorf("Expected max 5.0, got %v", max)
	}

	avg := calculateAvg(values)
	if avg != 3.0 {
		t.Errorf("Expected avg 3.0, got %v", avg)
	}
}

func TestEmptyStatisticsCalculation(t *testing.T) {
	// Test statistics calculation with empty values
	values := []float64{}
	
	min := calculateMin(values)
	if min != 0 {
		t.Errorf("Expected min 0 for empty slice, got %v", min)
	}

	max := calculateMax(values)
	if max != 0 {
		t.Errorf("Expected max 0 for empty slice, got %v", max)
	}

	avg := calculateAvg(values)
	if avg != 0 {
		t.Errorf("Expected avg 0 for empty slice, got %v", avg)
	}
}

func TestGetCurrentTimestamp(t *testing.T) {
	// Test that getCurrentTimestamp returns a reasonable timestamp
	timestamp := getCurrentTimestamp()
	
	// The example timestamp should be greater than 0
	if timestamp <= 0 {
		t.Errorf("Expected positive timestamp, got %v", timestamp)
	}

	// The example timestamp should be reasonable (not too far in future/past)
	// 1640995200 is roughly 2022-01-01, which is reasonable
	if timestamp < 1600000000 || timestamp > 1800000000 {
		t.Logf("Timestamp %v seems unusual, but this is just an example", timestamp)
	}
}

// Mock API test for History methods
func TestHistoryAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test that methods exist and return appropriate types
	opts := HistoryGetOptions{
		ItemIDs: []string{"12345"},
		History: HistoryFloat,
		Limit:   10,
	}

	// These calls will fail without a real Zabbix server, but we can verify the method signatures
	_, err := api.HistoryGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.HistoryGetFloatByItem("12345", 0, 0, 10)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.HistoryGetStringByItem("12345", 0, 0, 10)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.HistoryGetRecent("12345", HistoryFloat, 10)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.HistoryGetLatest("12345", HistoryFloat)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

// Benchmark tests
func BenchmarkHistoryDataMarshaling(b *testing.B) {
	hd := HistoryData{
		ItemID: "12345",
		Clock:  1640995200,
		Value:  "test_value",
		NS:     123456,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(hd)
	}
}

func BenchmarkStatisticsCalculation(b *testing.B) {
	values := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = float64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateMin(values)
		calculateMax(values)
		calculateAvg(values)
	}
}

// Test integration scenarios
func TestHistoryIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test multiple history types for same item
	opts := HistoryGetOptions{
		ItemIDs:   []string{"12345", "67890"},
		History:   HistoryFloat,
		TimeFrom:  1640995200,
		TimeTill:  1640998800,
		Limit:     100,
		SortField: "clock",
		SortOrder: "DESC",
	}

	_, err := api.HistoryGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test time range queries
	_, err = api.HistoryGetByTimeRange([]string{"12345"}, HistoryFloat, 
		int(time.Now().Unix()-3600), int(time.Now().Unix()), 50)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test statistics
	_, err = api.HistoryGetStats("12345", HistoryFloat, 
		int(time.Now().Unix()-3600), int(time.Now().Unix()))
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}