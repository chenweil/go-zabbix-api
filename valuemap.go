package zabbix

import (
	"encoding/json"
	"fmt"
)

// ValueMap represents a Zabbix value map object
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/object
type ValueMap struct {
	ValueMapID   string        `json:"valuemapid,omitempty"`
	Name         string        `json:"name"`
	CustomMapping []ValueMapMapping `json:"mappings,omitempty"`
}

// ValueMaps represents an array of ValueMap objects
type ValueMaps []ValueMap

// ValueMapMapping represents a value map mapping
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/object#mapping
type ValueMapMapping struct {
	MappingID string `json:"mappingid,omitempty"`
	ValueMapID string `json:"valuemapid,omitempty"`
	Value     string `json:"value"`
	NewValue string `json:"newvalue"`
}

// ValueMapMappings represents an array of ValueMapMapping objects
type ValueMapMappings []ValueMapMapping

// ValueMapGetOptions represents parameters for valuemap.get API call
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/get
type ValueMapGetOptions struct {
	ValueMapIDs   []string               `json:"valuemapids,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	Search        map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string        `json:"searchWildcardsEnabled,omitempty"`
	Output        string               `json:"output,omitempty"`
	SelectMappings string              `json:"selectMappings,omitempty"`
	SortField     string               `json:"sortfield,omitempty"`
	SortOrder     string               `json:"sortorder,omitempty"`
	Limit         int                  `json:"limit,omitempty"`
}

// ValueMapCreateOptions represents parameters for valuemap.create API call
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/create
type ValueMapCreateOptions struct {
	ValueMaps     ValueMaps         `json:"valuemaps"`
	SelectMappings string           `json:"selectMappings,omitempty"`
}

// ValueMapUpdateOptions represents parameters for valuemap.update API call
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/update
type ValueMapUpdateOptions struct {
	ValueMaps     ValueMaps         `json:"valuemaps"`
	SelectMappings string           `json:"selectMappings,omitempty"`
}

// ValueMapDeleteOptions represents parameters for valuemap.delete API call
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/delete
type ValueMapDeleteOptions struct {
	ValueMapIDs []string `json:"valuemapids"`
}

// ValueMapWithMappings represents a valuemap with its mappings
type ValueMapWithMappings struct {
	ValueMap
	Mappings []ValueMapMapping `json:"mappings,omitempty"`
}

// ValueMapWithMappingsSlice represents a slice of ValueMapWithMappings
type ValueMapWithMappingsSlice []ValueMapWithMappings

// ValueMapImport represents value map for import/export
type ValueMapImport struct {
	ValueMaps []ValueMap `json:"valuemaps,omitempty"`
}

// ValueMapImportResult represents import result
type ValueMapImportResult struct {
	Created    []string `json:"created,omitempty"`
	Updated    []string `json:"updated,omitempty"`
	Failed     []string `json:"failed,omitempty"`
	Count      int      `json:"count,omitempty"`
}

// ValueMapValidation represents value map validation result
type ValueMapValidation struct {
	IsValid    bool                `json:"valid"`
	Errors     []string            `json:"errors,omitempty"`
	Warnings   []string            `json:"warnings,omitempty"`
	Statistics map[string]interface{} `json:"statistics,omitempty"`
}

// ValueMapTemplate represents valuemap-template relationship
type ValueMapTemplate struct {
	ValueMapID string `json:"valuemapid"`
	TemplateID string `json:"templateid"`
}

// ValueMapTemplates represents an array of ValueMapTemplate
type ValueMapTemplates []ValueMapTemplate

// ValueMapPattern represents search pattern for valuemaps
type ValueMapPattern struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Tag         string            `json:"tag"`
	Value       string            `json:"value"`
}

// ValueMapStatistics represents valuemap statistics
type ValueMapStatistics struct {
	TotalValueMaps    int                  `json:"total_valuemaps"`
	ValueMapsInUse    int                  `json:"valuemaps_in_use"`
	UnusedValueMaps   int                  `json:"unused_valuemaps"`
	MappingCounts     map[string]int       `json:"mapping_counts"`
	ValueMapSizes     map[string]int       `json:"valuemap_sizes"`
	TopValueMaps     []ValueMap           `json:"top_valuemaps"`
	LeastUsedMaps     []ValueMap           `json:"least_used_valuemaps"`
}

// ValueMapsGet Wrapper for valuemap.get
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/get
func (api *API) ValueMapsGet(options ValueMapGetOptions) (ValueMaps, error) {
	
	// Prepare parameters for API call
	params := make(map[string]interface{})
	
	if len(options.ValueMapIDs) > 0 {
		params["valuemapids"] = options.ValueMapIDs
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
	if options.SelectMappings != "" {
		params["selectMappings"] = options.SelectMappings
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

	var valueMaps ValueMaps
	err := api.CallWithErrorParse("valuemap.get", params, &valueMaps)
	return valueMaps, err
}

// ValueMapsGetByID Get value maps by specific value map IDs
func (api *API) ValueMapsGetByID(valueMapIDs []string) (ValueMaps, error) {
	options := ValueMapGetOptions{
		ValueMapIDs: valueMapIDs,
		Output:    "extend",
	}
	return api.ValueMapsGet(options)
}

// ValueMapsGetByName Get value maps by name
func (api *API) ValueMapsGetByName(name string) (ValueMaps, error) {
	options := ValueMapGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.ValueMapsGet(options)
}

// ValueMapsGetByNames Get value maps by multiple names
func (api *API) ValueMapsGetByNames(names []string) (ValueMaps, error) {
	allValueMaps := ValueMaps{}
	
	for _, name := range names {
		valueMaps, err := api.ValueMapsGetByName(name)
		if err != nil {
			// Continue with other names even if one fails
			continue
		}
		allValueMaps = append(allValueMaps, valueMaps...)
	}
	
	return allValueMaps, nil
}

// ValueMapsGetByPattern Get value maps matching a pattern
func (api *API) ValueMapsGetByPattern(pattern string) (ValueMaps, error) {
	options := ValueMapGetOptions{
		Search: map[string]interface{}{
			"name": pattern,
		},
		SearchWildcardsEnabled: "true",
		Output:               "extend",
	}
	return api.ValueMapsGet(options)
}

// ValueMapsGetWithMappings Get value maps with their mappings
func (api *API) ValueMapsGetWithMappings(options ValueMapGetOptions) (ValueMapWithMappingsSlice, error) {
	options.SelectMappings = "extend"
	
	valueMaps, err := api.ValueMapsGet(options)
	if err != nil {
		return nil, err
	}
	
	valueMapsWithMappings := make(ValueMapWithMappingsSlice, 0, len(valueMaps))
	for _, valueMap := range valueMaps {
		// Get mappings for this value map
		mappings, err := api.ValueMapMappingsGetByValueMapID(valueMap.ValueMapID)
		if err != nil {
			// If we can't get mappings, just include the value map without mappings
			valueMapsWithMappings = append(valueMapsWithMappings, ValueMapWithMappings{
				ValueMap: valueMap,
				Mappings: nil,
			})
		} else {
			valueMapsWithMappings = append(valueMapsWithMappings, ValueMapWithMappings{
				ValueMap: valueMap,
				Mappings: mappings,
			})
		}
	}
	
	return valueMapsWithMappings, nil
}

// ValueMapGetByID Get value map by ID (exactly one match required)
func (api *API) ValueMapGetByID(valueMapID string) (*ValueMap, error) {
	valueMaps, err := api.ValueMapsGetByID([]string{valueMapID})
	if err != nil {
		return nil, err
	}

	if len(valueMaps) == 1 {
		return &valueMaps[0], nil
	} else if len(valueMaps) == 0 {
		return nil, fmt.Errorf("Value map not found: %s", valueMapID)
	} else {
		return nil, fmt.Errorf("Multiple value maps found with ID: %s", valueMapID)
	}
}

// ValueMapsCreate Wrapper for valuemap.create
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/create
func (api *API) ValueMapsCreate(valueMaps ValueMaps) (result []string, err error) {
	options := ValueMapCreateOptions{
		ValueMaps: valueMaps,
	}

	response, err := api.CallWithError("valuemap.create", options)
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
			if valueMapMap, ok := item.(map[string]interface{}); ok {
				if valuemapid, exists := valueMapMap["valuemapids"]; exists {
					if idArray, ok := valuemapid.([]interface{}); ok && len(idArray) > 0 {
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

// ValueMapCreateSingle Create a single value map
func (api *API) ValueMapCreateSingle(valueMap ValueMap) (valueMapID string, err error) {
	valueMaps := ValueMaps{valueMap}
	result, err := api.ValueMapsCreate(valueMaps)
	if len(result) > 0 {
		valueMapID = result[0]
	}
	return
}

// ValueMapsUpdate Wrapper for valuemap.update
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/update
func (api *API) ValueMapsUpdate(valueMaps ValueMaps) (result []string, err error) {
	options := ValueMapUpdateOptions{
		ValueMaps: valueMaps,
	}

	response, err := api.CallWithError("valuemap.update", options)
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
			if valueMapMap, ok := item.(map[string]interface{}); ok {
				if valuemapid, exists := valueMapMap["valuemapids"]; exists {
					if idArray, ok := valuemapid.([]interface{}); ok && len(idArray) > 0 {
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

// ValueMapUpdateSingle Update a single value map
func (api *API) ValueMapUpdateSingle(valueMap ValueMap) (valueMapID string, err error) {
	valueMaps := ValueMaps{valueMap}
	result, err := api.ValueMapsUpdate(valueMaps)
	if len(result) > 0 {
		valueMapID = result[0]
	}
	return
}

// ValueMapsDelete Wrapper for valuemap.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/valuemap/delete
func (api *API) ValueMapsDelete(valueMaps ValueMaps) (result []string, err error) {
	valueMapIDs := make([]string, len(valueMaps))
	for i, valueMap := range valueMaps {
		valueMapIDs[i] = valueMap.ValueMapID
	}
	
	return api.ValueMapsDeleteByIDs(valueMapIDs)
}

// ValueMapsDeleteByIDs Wrapper for valuemap.delete with IDs
func (api *API) ValueMapsDeleteByIDs(valueMapIDs []string) (result []string, err error) {
	options := ValueMapDeleteOptions{
		ValueMapIDs: valueMapIDs,
	}

	response, err := api.CallWithError("valuemap.delete", options)
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

// ValueMapDeleteSingle Delete a single value map
func (api *API) ValueMapDeleteSingle(valueMapID string) (err error) {
	_, err = api.ValueMapsDeleteByIDs([]string{valueMapID})
	return
}

// ValueMapMappingsGetByValueMapID Get mappings for a value map
func (api *API) ValueMapMappingsGetByValueMapID(valueMapID string) ([]ValueMapMapping, error) {
	// Get the value map with mappings
	valueMapsWithMappings, err := api.ValueMapsGetWithMappings(ValueMapGetOptions{
		ValueMapIDs: []string{valueMapID},
		Output:     "extend",
	})
	if err != nil {
		return nil, err
	}

	if len(valueMapsWithMappings) > 0 {
		return valueMapsWithMappings[0].Mappings, nil
	}

	return nil, fmt.Errorf("Value map not found: %s", valueMapID)
}

// ValueMapsGetCount Get count of value maps matching criteria
func (api *API) ValueMapsGetCount(options ValueMapGetOptions) (int, error) {
	options.Limit = 1
	valueMaps, err := api.ValueMapsGet(options)
	if err != nil {
		return 0, err
	}
	
	// Get the total count from API metadata if available
	// This is a simplified implementation
	return len(valueMaps), nil
}

// ValueMapsGetStatistics Get statistics about value maps
func (api *API) ValueMapsGetStatistics() (ValueMapStatistics, error) {
	// Get all value maps
	allValueMaps, err := api.ValueMapsGet(ValueMapGetOptions{Output: "extend"})
	if err != nil {
		return ValueMapStatistics{}, err
	}

	stats := ValueMapStatistics{
		TotalValueMaps: len(allValueMaps),
		MappingCounts: make(map[string]int),
		ValueMapSizes: make(map[string]int),
	}

	// Analyze value maps
	valueMapsInUse := 0
	unusedValueMaps := 0

	for _, valueMap := range allValueMaps {
		// Count mappings
		mappingCount := len(valueMap.CustomMapping)
		stats.MappingCounts[valueMap.Name] = mappingCount
		stats.ValueMapSizes[valueMap.Name] = mappingCount

		// Determine if in use (simplified logic)
		if mappingCount > 0 {
			valueMapsInUse++
		} else {
			unusedValueMaps++
		}
	}

	stats.ValueMapsInUse = valueMapsInUse
	stats.UnusedValueMaps = unusedValueMaps

	// Find top value maps (simplified)
	if len(allValueMaps) > 0 {
		stats.TopValueMaps = allValueMaps[:min(5, len(allValueMaps))]
		stats.LeastUsedMaps = allValueMaps[max(0, len(allValueMaps)-5):]
	}

	return stats, nil
}

// ValueMapValidate Validate a value map configuration
func (api *API) ValueMapValidate(valueMap ValueMap) (ValueMapValidation) {
	validation := ValueMapValidation{
		IsValid:  true,
		Errors:   []string{},
		Warnings: []string{},
	}
	
	// Check required fields
	if valueMap.Name == "" {
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "Value map name is required")
	}
	
	// Validate mappings if present
	if len(valueMap.CustomMapping) == 0 {
		validation.Warnings = append(validation.Warnings, "Value map has no mappings")
	} else {
		valueMapValidation := api.validateValueMapMappings(valueMap.CustomMapping)
		validation.Errors = append(validation.Errors, valueMapValidation.Errors...)
		validation.Warnings = append(validation.Warnings, valueMapValidation.Warnings...)
		if len(valueMapValidation.Errors) > 0 {
			validation.IsValid = false
		}
	}
	
	// Add statistics
	validation.Statistics = map[string]interface{}{
		"mapping_count": len(valueMap.CustomMapping),
		"name_length":   len(valueMap.Name),
	}
	
	return validation
}

// validateValueMapMappings Validate value map mappings
func (api *API) validateValueMapMappings(mappings []ValueMapMapping) ValueMapValidation {
	validation := ValueMapValidation{
		IsValid:  true,
		Errors:   []string{},
		Warnings: []string{},
	}
	
	// Check for duplicate values
	valueMap := make(map[string]bool)
	newValueMap := make(map[string]bool)
	
	for i, mapping := range mappings {
		if mapping.Value == "" {
			validation.IsValid = false
			validation.Errors = append(validation.Errors, fmt.Sprintf("Mapping %d: value is required", i))
		}
		
		if mapping.NewValue == "" {
			validation.IsValid = false
			validation.Errors = append(validation.Errors, fmt.Sprintf("Mapping %d: new value is required", i))
		}
		
		// Check for duplicate values
		if valueMap[mapping.Value] {
			validation.Warnings = append(validation.Warnings, fmt.Sprintf("Mapping %d: duplicate value '%s'", i, mapping.Value))
		} else {
			valueMap[mapping.Value] = true
		}
		
		// Check for duplicate new values
		if newValueMap[mapping.NewValue] {
			validation.Warnings = append(validation.Warnings, fmt.Sprintf("Mapping %d: duplicate new value '%s'", i, mapping.NewValue))
		} else {
			newValueMap[mapping.NewValue] = true
		}
	}
	
	return validation
}

// ValueMapAddMapping Add a mapping to a value map
func (api *API) ValueMapAddMapping(valueMapID string, mapping ValueMapMapping) (err error) {
	// Get current value map
	valueMap, err := api.ValueMapGetByID(valueMapID)
	if err != nil {
		return err
	}
	
	// Add mapping
	valueMap.CustomMapping = append(valueMap.CustomMapping, mapping)
	
	// Update the value map
	_, err = api.ValueMapUpdateSingle(*valueMap)
	return err
}

// ValueMapRemoveMapping Remove a mapping from a value map
func (api *API) ValueMapRemoveMapping(valueMapID, mappingID string) (err error) {
	// Get current value map
	valueMap, err := api.ValueMapGetByID(valueMapID)
	if err != nil {
		return err
	}
	
	// Remove mapping
	for i, mapping := range valueMap.CustomMapping {
		if mapping.MappingID == mappingID {
			valueMap.CustomMapping = append(valueMap.CustomMapping[:i], valueMap.CustomMapping[i+1:]...)
			break
		}
	}
	
	// Update the value map
	_, err = api.ValueMapUpdateSingle(*valueMap)
	return err
}

// ValueMapUpdateMapping Update a mapping in a value map
func (api *API) ValueMapUpdateMapping(valueMapID string, mapping ValueMapMapping) (err error) {
	// Get current value map
	valueMap, err := api.ValueMapGetByID(valueMapID)
	if err != nil {
		return err
	}
	
	// Update mapping
	for i, m := range valueMap.CustomMapping {
		if m.MappingID == mapping.MappingID {
			valueMap.CustomMapping[i] = mapping
			break
		}
	}
	
	// Update the value map
	_, err = api.ValueMapUpdateSingle(*valueMap)
	return err
}

// CreateSimpleValueMap Create a simple value map with basic mappings
func (api *API) CreateSimpleValueMap(name string, mappings []ValueMapMapping) (string, error) {
	valueMap := ValueMap{
		Name:         name,
		CustomMapping: mappings,
	}
	
	return api.ValueMapCreateSingle(valueMap)
}

// CreateBooleanValueMap Create a boolean value map (0=false, 1=true)
func (api *API) CreateBooleanValueMap(name string) (string, error) {
	mappings := []ValueMapMapping{
		{Value: "0", NewValue: "False"},
		{Value: "1", NewValue: "True"},
	}
	
	return api.CreateSimpleValueMap(name, mappings)
}

// CreateSeverityValueMap Create a severity value map
func (api *API) CreateSeverityValueMap(name string) (string, error) {
	mappings := []ValueMapMapping{
		{Value: "0", NewValue: "Not Classified"},
		{Value: "1", NewValue: "Information"},
		{Value: "2", NewValue: "Warning"},
		{Value: "3", NewValue: "Average"},
		{Value: "4", NewValue: "High"},
		{Value: "5", NewValue: "Disaster"},
	}
	
	return api.CreateSimpleValueMap(name, mappings)
}

// ValueMapTest Test value map mappings against test values
func (api *API) ValueMapTest(valueMapID string, testValues []string) (map[string]string, error) {
	// Get value map with mappings
	valueMapsWithMappings, err := api.ValueMapsGetWithMappings(ValueMapGetOptions{
		ValueMapIDs: []string{valueMapID},
		Output:     "extend",
	})
	if err != nil {
		return nil, err
	}

	if len(valueMapsWithMappings) == 0 {
		return nil, fmt.Errorf("Value map not found: %s", valueMapID)
	}

	// Build mapping lookup
	valueMap := make(map[string]string)
	for _, mapping := range valueMapsWithMappings[0].Mappings {
		valueMap[mapping.Value] = mapping.NewValue
	}

	// Test values
	results := make(map[string]string)
	for _, testValue := range testValues {
		if mappedValue, exists := valueMap[testValue]; exists {
			results[testValue] = mappedValue
		} else {
			results[testValue] = fmt.Sprintf("No mapping for value: %s", testValue)
		}
	}

	return results, nil
}

// ValueMapExport Export value maps for backup/configuration
func (api *API) ValueMapExport(valueMapIDs []string) (ValueMapImport, error) {
	valueMaps, err := api.ValueMapsGetWithMappings(ValueMapGetOptions{
		ValueMapIDs: valueMapIDs,
		Output:     "extend",
	})
	if err != nil {
		return ValueMapImport{}, err
	}

	exportData := ValueMapImport{
		ValueMaps: make([]ValueMap, 0, len(valueMaps)),
	}

	for _, valueMapWithMappings := range valueMaps {
		valueMap := valueMapWithMappings.ValueMap
		valueMap.CustomMapping = valueMapWithMappings.Mappings
		exportData.ValueMaps = append(exportData.ValueMaps, valueMap)
	}

	return exportData, nil
}

// ValueMapImport Import value maps from backup/configuration
func (api *API) ValueMapImport(importData ValueMapImport, overwrite bool) (ValueMapImportResult, error) {
	created := []string{}
	updated := []string{}
	failed := []string{}

	for _, valueMap := range importData.ValueMaps {
		// Check if value map exists
		existingValueMaps, err := api.ValueMapsGetByName(valueMap.Name)
		if err != nil {
			failed = append(failed, fmt.Sprintf("Failed to check existence of %s: %v", valueMap.Name, err))
			continue
		}

		if len(existingValueMaps) == 0 {
			// Create new value map
			valueMapID, err := api.ValueMapCreateSingle(valueMap)
			if err != nil {
				failed = append(failed, fmt.Sprintf("Failed to create %s: %v", valueMap.Name, err))
			} else {
				created = append(created, valueMapID)
			}
		} else if overwrite {
			// Update existing value map
			valueMap := existingValueMaps[0]
			valueMap.CustomMapping = valueMap.CustomMapping
			
			_, err := api.ValueMapUpdateSingle(valueMap)
			if err != nil {
				failed = append(failed, fmt.Sprintf("Failed to update %s: %v", valueMap.Name, err))
			} else {
				updated = append(updated, valueMap.ValueMapID)
			}
		} else {
			// Skip existing value map
			continue
		}
	}

	result := ValueMapImportResult{
		Created: created,
		Updated: updated,
		Failed: failed,
		Count:  len(importData.ValueMaps),
	}

	return result, nil
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ValueMapIsEmpty Check if value map has no mappings
func (valueMap *ValueMap) IsEmpty() bool {
	return len(valueMap.CustomMapping) == 0
}

// ValueMapHasMapping Check if value map has a specific mapping
func (valueMap *ValueMap) HasMapping(value string) bool {
	for _, mapping := range valueMap.CustomMapping {
		if mapping.Value == value {
			return true
		}
	}
	return false
}

// ValueMapGetMappingCount Get number of mappings
func (valueMap *ValueMap) GetMappingCount() int {
	return len(valueMap.CustomMapping)
}

// ValueMapFindMapping Find mapping by value
func (valueMap *ValueMap) FindMapping(value string) (ValueMapMapping, bool) {
	for _, mapping := range valueMap.CustomMapping {
		if mapping.Value == value {
			return mapping, true
		}
	}
	return ValueMapMapping{}, false
}

// ValueMapConvertValue Convert a value using the value map
func (valueMap *ValueMap) ConvertValue(value string) (string, bool) {
	if mapping, found := valueMap.FindMapping(value); found {
		return mapping.NewValue, true
	}
	return value, false
}