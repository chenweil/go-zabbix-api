package zabbix

import (
	"testing"
	"time"
	"math"
)

func TestTrendValueTypeString(t *testing.T) {
	tests := []struct {
		tt       TrendValueType
		expected string
	}{
		{TrendFloat, "0"},
		{TrendUnsigned, "3"},
	}

	for _, test := range tests {
		if got := string(test.tt); got != test.expected {
			t.Errorf("TrendValueType %v, expected %v, got %v", test.tt, test.expected, got)
		}
	}
}

func TestTrendGetOptions(t *testing.T) {
	// Test default values
	opts := TrendGetOptions{
		ItemIDs: []string{"12345"},
		TrendType: TrendFloat,
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

func TestTrend(t *testing.T) {
	// Test Trend creation and JSON marshaling
	trend := Trend{
		ItemID:   "12345",
		Clock:    1640995200,
		ValueMin: "10.5",
		ValueAvg: "15.3",
		ValueMax: "20.1",
		NS:       123456,
	}

	expectedJSON := `{"itemid":"12345","clock":1640995200,"value_min":"10.5","value_avg":"15.3","value_max":"20.1","ns":123456}`
	
	jsonData, err := json.Marshal(trend)
	if err != nil {
		t.Errorf("Failed to marshal Trend: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestTrendFloat(t *testing.T) {
	// Test TrendFloat creation and JSON marshaling
	tf := TrendFloat{
		ItemID:   "12345",
		Clock:    1640995200,
		ValueMin: 10.5,
		ValueAvg: 15.3,
		ValueMax: 20.1,
		NS:       123456,
	}

	expectedJSON := `{"itemid":"12345","clock":1640995200,"value_min":10.5,"value_avg":15.3,"value_max":20.1,"ns":123456}`
	
	jsonData, err := json.Marshal(tf)
	if err != nil {
		t.Errorf("Failed to marshal TrendFloat: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestTrendUnsigned(t *testing.T) {
	// Test TrendUnsigned creation and JSON marshaling
	tu := TrendUnsigned{
		ItemID:   "12345",
		Clock:    1640995200,
		ValueMin: 100,
		ValueAvg: 150,
		ValueMax: 200,
		NS:       123456,
	}

	expectedJSON := `{"itemid":"12345","clock":1640995200,"value_min":100,"value_avg":150,"value_max":200,"ns":123456}`
	
	jsonData, err := json.Marshal(tu)
	if err != nil {
		t.Errorf("Failed to marshal TrendUnsigned: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestTrendSummary(t *testing.T) {
	// Test TrendSummary creation and JSON marshaling
	summary := TrendSummary{
		ItemID:      "12345",
		Count:       100,
		MinValue:    5.0,
		MaxValue:    95.0,
		AvgValue:    50.0,
		MinClock:    1640995200,
		MaxClock:    1641081600,
		TimeRange:   86400,
		Periodicity: 864.0,
	}

	expectedJSON := `{"itemid":"12345","count":100,"min_value":5,"max_value":95,"avg_value":50,"min_clock":1640995200,"max_clock":1641081600,"time_range":86400,"periodicity":864}`
	
	jsonData, err := json.Marshal(summary)
	if err != nil {
		t.Errorf("Failed to marshal TrendSummary: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestTrendSliceToFloatTrends(t *testing.T) {
	// Test conversion from generic TrendSlice to TrendFloatSlice
	trends := TrendSlice{
		{ItemID: "123", Clock: 1000, ValueMin: "10.5", ValueAvg: "15.3", ValueMax: "20.1", NS: 1},
		{ItemID: "123", Clock: 1001, ValueMin: "11.2", ValueAvg: "16.1", ValueMax: "21.3", NS: 2},
		{ItemID: "123", Clock: 1002, ValueMin: "9.8", ValueAvg: "14.9", ValueMax: "19.7", NS: 3},
		{ItemID: "456", Clock: 1000, ValueMin: "25.1", ValueAvg: "30.2", ValueMax: "35.4", NS: 4},
	}

	floatTrends, err := trends.ToFloatTrends()
	if err != nil {
		t.Errorf("Failed to convert to float trends: %v", err)
	}

	expectedCount := 4
	if len(floatTrends) != expectedCount {
		t.Errorf("Expected %d float trend records, got %d", expectedCount, len(floatTrends))
	}

	// Check specific values
	if floatTrends[0].ValueAvg != 15.3 {
		t.Errorf("Expected first avg value 15.3, got %v", floatTrends[0].ValueAvg)
	}

	if floatTrends[1].ValueMax != 21.3 {
		t.Errorf("Expected second max value 21.3, got %v", floatTrends[1].ValueMax)
	}

	if floatTrends[2].ValueMin != 9.8 {
		t.Errorf("Expected third min value 9.8, got %v", floatTrends[2].ValueMin)
	}
}

func TestParseFloat64(t *testing.T) {
	// Test parsing valid float strings
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"10.5", 10.5, false},
		{"-5.3", -5.3, false},
		{"0", 0, false},
		{"", 0, false},
		{"invalid", 0, true},
		{"10.5.2", 0, true},
	}

	for _, test := range tests {
		result, err := parseFloat64(test.input)
		
		if test.hasError && err == nil {
			t.Errorf("Expected error for input '%s', but got none", test.input)
		} else if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input '%s': %v", test.input, err)
		} else if !test.hasError && result != test.expected {
			t.Errorf("Expected %v for input '%s', got %v", test.expected, result)
		}
	}
}

func TestAbsFunction(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{5.0, 5.0},
		{-5.0, 5.0},
		{0, 0},
		{-123.456, 123.456},
	}

	for _, test := range tests {
		result := abs(test.input)
		if result != test.expected {
			t.Errorf("Expected abs(%v) = %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestSqrtFunction(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
		tolerance float64
	}{
		{4.0, 2.0, 0.0001},
		{9.0, 3.0, 0.0001},
		{0, 0, 0},
		{2.0, 1.41421, 0.0001},
	}

	for _, test := range tests {
		result := sqrt(test.input)
		if math.Abs(result-test.expected) > test.tolerance {
			t.Errorf("Expected sqrt(%v) ≈ %v, got %v", test.input, test.expected, result)
		}
	}
}

func TestCalculateStandardDeviation(t *testing.T) {
	// Test standard deviation calculation
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	mean := 3.0
	
	stdDev := calculateStandardDeviation(values, mean)
	// For values [1,2,3,4,5] with mean 3.0
	// Variance = ((1-3)^2 + (2-3)^2 + (3-3)^2 + (4-3)^2 + (5-3)^2) / 5 = (4+1+0+1+4)/5 = 10/5 = 2
	// Standard deviation = sqrt(2) ≈ 1.41421
	expectedStdDev := math.Sqrt(2.0)
	
	if math.Abs(stdDev-expectedStdDev) > 0.0001 {
		t.Errorf("Expected std dev ≈ %v, got %v", expectedStdDev, stdDev)
	}
}

func TestCalculateStandardDeviationEmpty(t *testing.T) {
	// Test with empty slice
	stdDev := calculateStandardDeviation([]float64{}, 0)
	if stdDev != 0 {
		t.Errorf("Expected std dev 0 for empty slice, got %v", stdDev)
	}
}

func TestMockTrendAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test that methods exist and return appropriate types
	opts := TrendGetOptions{
		ItemIDs: []string{"12345"},
		TrendType: TrendFloat,
		Limit:   10,
	}

	// These calls will fail without a real Zabbix server, but we can verify the method signatures
	_, err := api.TrendGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetFloatByItem("12345", 0, 0, 10)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetUnsignedByItem("12345", 0, 0, 10)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetHourly("12345", TrendFloat, 24)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetDaily("12345", TrendFloat, 30)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetLatest("12345", TrendFloat)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetSummary("12345", TrendFloat, 0, 0)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetMultipleSummaries([]string{"12345"}, TrendFloat, 0, 0)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetComparison("12345", TrendFloat, 0, 0, 0, 0)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.TrendGetAnomalies("12345", TrendFloat, 0, 0, 2.0)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkTrendMarshaling(b *testing.B) {
	trend := Trend{
		ItemID:   "12345",
		Clock:    1640995200,
		ValueMin: "10.5",
		ValueAvg: "15.3",
		ValueMax: "20.1",
		NS:       123456,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(trend)
	}
}

func BenchmarkParseFloat64(b *testing.B) {
	input := "15.375"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parseFloat64(input)
	}
}

func BenchmarkStandardDeviation(b *testing.B) {
	values := make([]float64, 100)
	for i := 0; i < 100; i++ {
		values[i] = float64(i)
	}
	mean := 49.5

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateStandardDeviation(values, mean)
	}
}

// Test integration scenarios
func TestTrendIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test multiple trend types for same item
	opts := TrendGetOptions{
		ItemIDs:   []string{"12345", "67890"},
		TrendType: TrendFloat,
		TimeFrom:  1640995200,
		TimeTill:  1640998800,
		Limit:     100,
		SortField: "clock",
		SortOrder: "DESC",
	}

	_, err := api.TrendGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test time range queries
	_, err = api.TrendGetByTimeRange([]string{"12345"}, TrendFloat, 
		int(time.Now().Unix()-3600), int(time.Now().Unix()), 50)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test summary statistics
	_, err = api.TrendGetSummary("12345", TrendFloat, 
		int(time.Now().Unix()-86400), int(time.Now().Unix()))
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}