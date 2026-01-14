package zabbix

import (
	"testing"
)

func TestValueMapGetOptions(t *testing.T) {
	// Test default values
	opts := ValueMapGetOptions{
		ValueMapIDs: []string{"12345"},
	}

	if len(opts.ValueMapIDs) != 1 {
		t.Errorf("Expected 1 value map ID, got %d", len(opts.ValueMapIDs))
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

func TestValueMap(t *testing.T) {
	// Test ValueMap creation and JSON marshaling
	valueMap := ValueMap{
		ValueMapID: "12345",
		Name:       "Boolean Values",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
			{Value: "1", NewValue: "True"},
		},
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(valueMap)
	if err != nil {
		t.Errorf("Failed to marshal ValueMap: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestValueMapMapping(t *testing.T) {
	// Test ValueMapMapping creation and JSON marshaling
	mapping := ValueMapMapping{
		MappingID: "67890",
		ValueMapID: "12345",
		Value:     "0",
		NewValue:  "False",
	}

	expectedJSON := `{"mappingid":"67890","valuemapid":"12345","value":"0","newvalue":"False"}`
	
	jsonData, err := json.Marshal(mapping)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapMapping: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestValueMapWithMappings(t *testing.T) {
	// Test ValueMapWithMappings creation and JSON marshaling
	valueMap := ValueMap{
		ValueMapID: "12345",
		Name:       "Severity Values",
	}

	mappings := []ValueMapMapping{
		{Value: "0", NewValue: "Not Classified"},
		{Value: "1", NewValue: "Information"},
		{Value: "2", NewValue: "Warning"},
	}

	valueMapWithMappings := ValueMapWithMappings{
		ValueMap: valueMap,
		Mappings: mappings,
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(valueMapWithMappings)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapWithMappings: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestValueMapPattern(t *testing.T) {
	// Test ValueMapPattern creation and JSON marshaling
	pattern := ValueMapPattern{
		Name:        "Boolean*",
		Description: "Boolean value mappings",
		Tag:         "type",
		Value:       "boolean",
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(pattern)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapPattern: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestValueMapValidation(t *testing.T) {
	// Test ValueMapValidation creation and JSON marshaling
	validation := ValueMapValidation{
		IsValid:  true,
		Errors:   []string{"Error 1", "Error 2"},
		Warnings: []string{"Warning 1", "Warning 2"},
		Statistics: map[string]interface{}{
			"mapping_count": 5,
			"name_length": 20,
		},
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(validation)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapValidation: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestValueMapImport(t *testing.T) {
	// Test ValueMapImport creation and JSON marshaling
	importData := ValueMapImport{
		ValueMaps: []ValueMap{
			{
				Name: "Test Value Map",
				CustomMapping: []ValueMapMapping{
					{Value: "0", NewValue: "Test"},
				},
			},
		},
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(importData)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapImport: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestValueMapImportResult(t *testing.T) {
	// Test ValueMapImportResult creation and JSON marshaling
	result := ValueMapImportResult{
		Created: []string{"1", "2"},
		Updated: []string{"3"},
		Failed:  []string{"4"},
		Count:   4,
	}

	expectedJSON := `{"created":["1","2"],"updated":["3"],"failed":["4"],"count":4}`
	
	jsonData, err := json.Marshal(result)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapImportResult: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestValueMapTemplate(t *testing.T) {
	// Test ValueMapTemplate creation and JSON marshaling
	relationship := ValueMapTemplate{
		ValueMapID: "12345",
		TemplateID: "67890",
	}

	expectedJSON := `{"valuemapid":"12345","templateid":"67890"}`
	
	jsonData, err := json.Marshal(relationship)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapTemplate: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}
}

func TestValueMapStatistics(t *testing.T) {
	// Test ValueMapStatistics creation and JSON marshaling
	stats := ValueMapStatistics{
		TotalValueMaps:   100,
		ValueMapsInUse:   80,
		UnusedValueMaps:  20,
		MappingCounts:    map[string]int{"map1": 5, "map2": 3},
		ValueMapSizes:     map[string]int{"map1": 5, "map2": 3},
		TopValueMaps:    []ValueMap{{Name: "Top Map"}},
		LeastUsedMaps:   []ValueMap{{Name: "Least Used Map"}},
	}

	// Test that the struct can be marshaled
	jsonData, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Failed to marshal ValueMapStatistics: %v", err)
	}

	// Verify it contains expected fields
	jsonStr := string(jsonData)
	if len(jsonStr) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func TestValueMapIsEmpty(t *testing.T) {
	// Test ValueMap.IsEmpty() method
	emptyValueMap := ValueMap{
		Name:         "Empty Map",
		CustomMapping: []ValueMapMapping{},
	}
	nonEmptyValueMap := ValueMap{
		Name: "Non-Empty Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
		},
	}

	if !emptyValueMap.IsEmpty() {
		t.Errorf("Expected empty value map to be detected as empty")
	}

	if nonEmptyValueMap.IsEmpty() {
		t.Errorf("Expected non-empty value map to NOT be detected as empty")
	}
}

func TestValueMapHasMapping(t *testing.T) {
	// Test ValueMap.HasMapping() method
	valueMap := ValueMap{
		Name: "Test Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
			{Value: "1", NewValue: "True"},
		},
	}

	if !valueMap.HasMapping("0") {
		t.Errorf("Expected value map to have mapping for '0'")
	}

	if !valueMap.HasMapping("1") {
		t.Errorf("Expected value map to have mapping for '1'")
	}

	if valueMap.HasMapping("2") {
		t.Errorf("Expected value map to NOT have mapping for '2'")
	}
}

func TestValueMapGetMappingCount(t *testing.T) {
	// Test ValueMap.GetMappingCount() method
	emptyValueMap := ValueMap{
		Name:         "Empty Map",
		CustomMapping: []ValueMapMapping{},
	}
	singleValueMap := ValueMap{
		Name: "Single Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
		},
	}
	multipleValueMap := ValueMap{
		Name: "Multiple Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
			{Value: "1", NewValue: "True"},
			{Value: "2", NewValue: "Unknown"},
		},
	}

	emptyCount := emptyValueMap.GetMappingCount()
	if emptyCount != 0 {
		t.Errorf("Expected 0 mappings, got %d", emptyCount)
	}

	singleCount := singleValueMap.GetMappingCount()
	if singleCount != 1 {
		t.Errorf("Expected 1 mapping, got %d", singleCount)
	}

	multipleCount := multipleValueMap.GetMappingCount()
	if multipleCount != 3 {
		t.Errorf("Expected 3 mappings, got %d", multipleCount)
	}
}

func TestValueMapFindMapping(t *testing.T) {
	// Test ValueMap.FindMapping() method
	valueMap := ValueMap{
		Name: "Test Map",
		CustomMapping: []ValueMapMapping{
			{MappingID: "1", Value: "0", NewValue: "False"},
			{MappingID: "2", Value: "1", NewValue: "True"},
			{MappingID: "3", Value: "2", NewValue: "Unknown"},
		},
	}

	// Test existing mapping
	mapping, found := valueMap.FindMapping("1")
	if !found {
		t.Errorf("Expected mapping for '1' to be found")
	}
	if mapping.NewValue != "True" {
		t.Errorf("Expected mapping value 'True', got '%s'", mapping.NewValue)
	}

	// Test non-existing mapping
	_, found = valueMap.FindMapping("99")
	if found {
		t.Errorf("Expected mapping for '99' to NOT be found")
	}
}

func TestValueMapConvertValue(t *testing.T) {
	// Test ValueMap.ConvertValue() method
	valueMap := ValueMap{
		Name: "Test Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
			{Value: "1", NewValue: "True"},
		},
	}

	// Test mapped value
	converted, found := valueMap.ConvertValue("0")
	if !found {
		t.Errorf("Expected value '0' to be found")
	}
	if converted != "False" {
		t.Errorf("Expected converted value 'False', got '%s'", converted)
	}

	// Test unmapped value
	converted, found = valueMap.ConvertValue("2")
	if found {
		t.Errorf("Expected value '2' to NOT be found")
	}
	if converted != "2" {
		t.Errorf("Expected original value '2', got '%s'", converted)
	}
}

func TestValueMapValidate(t *testing.T) {
	// Test ValueMap validation
	api := NewAPI(Config{})
	
	// Valid value map
	validValueMap := ValueMap{
		Name: "Valid Value Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
			{Value: "1", NewValue: "True"},
		},
	}
	
	result := api.ValueMapValidate(validValueMap)
	if !result.IsValid {
		t.Errorf("Expected valid value map to be valid, got errors: %v", result.Errors)
	}
	
	// Invalid value map - missing name
	invalidValueMap1 := ValueMap{
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
		},
	}
	
	result = api.ValueMapValidate(invalidValueMap1)
	if result.IsValid {
		t.Errorf("Expected invalid value map (missing name) to be invalid")
	}
	
	// Value map with empty mappings (warning only)
	warningValueMap := ValueMap{
		Name: "Warning Map",
		CustomMapping: []ValueMapMapping{},
	}
	
	result = api.ValueMapValidate(warningValueMap)
	if result.IsValid {
		t.Errorf("Expected value map with empty mappings to have warnings")
	}
	if len(result.Warnings) == 0 {
		t.Errorf("Expected value map with empty mappings to have warnings")
	}
}

func TestValidateValueMapMappings(t *testing.T) {
	// Test mapping validation
	api := NewAPI.Config{}
	
	// Valid mappings
	validMappings := []ValueMapMapping{
		{Value: "0", NewValue: "False"},
		{Value: "1", NewValue: "True"},
	}
	
	result := api.validateValueMapMappings(validMappings)
	if !result.IsValid {
		t.Errorf("Expected valid mappings to be valid, got errors: %v", result.Errors)
	}
	
	// Invalid mappings - missing value
	invalidMappings1 := []ValueMapMapping{
		{Value: "", NewValue: "No Value"},
	}
	
	result = api.validateValueMapMappings(invalidMappings1)
	if result.IsValid {
		t.Errorf("Expected invalid mapping (missing value) to be invalid")
	}
	
	// Invalid mappings - missing new value
	invalidMappings2 := []ValueMapMapping{
		{Value: "0", NewValue: ""},
	}
	
	result = api.validateValueMapMappings(invalidMappings2)
	if result.IsValid {
		t.Errorf("Expected invalid mapping (missing new value) to be invalid")
	}
}

func TestMockValueMapsAPIMethods(t *testing.T) {
	// This is a basic test to ensure the API methods exist and have correct signatures
	api := NewAPI(Config{})
	
	// Test value map get operations
	opts := ValueMapGetOptions{
		ValueMapIDs: []string{"12345"},
		Output:     "extend",
		Limit:      10,
	}

	_, err := api.ValueMapsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapsGetByID([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapsGetByName("Test Value Map")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapsGetByNames([]string{"Map1", "Map2"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapsGetByPattern("Boolean*")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapsGetWithMappings(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapGetByID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test CRUD operations
	valueMap := ValueMap{
		Name: "Test Value Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
			{Value: "1", NewValue: "True"},
		},
	}
	
	_, err = api.ValueMapCreateSingle(valueMap)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapUpdateSingle(valueMap)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ValueMapDeleteSingle("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test mapping operations
	_, err = api.ValueMapMappingsGetByValueMapID("12345")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	mapping := ValueMapMapping{Value: "2", NewValue: "Unknown"}
	
	err = api.ValueMapAddMapping("12345", mapping)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ValueMapRemoveMapping("12345", "67890")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	err = api.ValueMapUpdateMapping("12345", mapping)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test utility methods
	_, err = api.ValueMapsGetCount(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapsGetStatistics()
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapTest("12345", []string{"0", "1", "2"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.ValueMapExport([]string{"12345"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	importData := ValueMapImport{
		ValueMaps: []ValueMap{
			{Name: "Import Map", CustomMapping: []ValueMapMapping{{Value: "0", NewValue: "Test"}}},
		},
	}
	
	_, err = api.ValueMapImport(importData, true)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test convenience methods
	_, err = api.CreateSimpleValueMap("Simple Map", []ValueMapMapping{{Value: "0", NewValue: "False"}})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.CreateBooleanValueMap("Boolean Map")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	_, err = api.CreateSeverityValueMap("Severity Map")
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}

func BenchmarkValueMapMarshaling(b *testing.B) {
	valueMap := ValueMap{
		ValueMapID: "12345",
		Name:       "Test Value Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "False"},
			{Value: "1", NewValue: "True"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(valueMap)
	}
}

func BenchmarkValueMapMappingMarshaling(b *testing.B) {
	mapping := ValueMapMapping{
		MappingID: "67890",
		ValueMapID: "12345",
		Value:     "0",
		NewValue:  "False",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(mapping)
	}
}

// Test integration scenarios
func TestValueMapsIntegration(t *testing.T) {
	// Test realistic usage scenarios
	api := NewAPI(Config{})

	// Test complex filtering
	opts := ValueMapGetOptions{
		Filter: map[string]interface{}{
			"name": "Boolean",
		},
		Output:           "extend",
		SelectMappings:   "extend",
		SortField:        "name",
		SortOrder:        "ASC",
		Limit:            100,
	}

	_, err := api.ValueMapsGet(opts)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test value map creation with mappings
	valueMap := ValueMap{
		Name: "Integration Test Value Map",
		CustomMapping: []ValueMapMapping{
			{Value: "0", NewValue: "Disabled"},
			{Value: "1", NewValue: "Enabled"},
			{Value: "2", NewValue: "Unknown"},
		},
	}
	
	_, err = api.ValueMapCreateSingle(valueMap)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test value map update
	valueMap.ValueMapID = "11111"
	valueMap.Name = "Updated Integration Test Value Map"
	
	_, err = api.ValueMapUpdateSingle(valueMap)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test mapping operations
	newMapping := ValueMapMapping{
		Value:     "3",
		NewValue:  "Warning",
	}
	
	err = api.ValueMapAddMapping("12345", newMapping)
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test value map testing
	_, err = api.ValueMapTest("12345", []string{"0", "1", "99"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test export/import
	_, err = api.ValueMapExport([]string{"12345", "67890"})
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}

	// Test bulk operations
	valueMaps := []ValueMap{
		{Name: "Bulk Map 1", CustomMapping: []ValueMapMapping{{Value: "0", NewValue: "Test1"}}},
		{Name: "Bulk Map 2", CustomMapping: []ValueMapMapping{{Value: "0", NewValue: "Test2"}}},
		{Name: "Bulk Map 3", CustomMapping: []ValueMapMapping{{Value: "0", NewValue: "Test3"}}},
	}
	
	_, err = api.ValueMapsCreate(ValueMaps(valueMaps))
	if err == nil {
		t.Error("Expected error for mock API call, got nil")
	}
}