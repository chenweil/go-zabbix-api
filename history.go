package zabbix

import "encoding/json"

// HistoryValueType represents the type of history data
type HistoryValueType int

const (
	// HistoryFloat represents float values
	HistoryFloat HistoryValueType = 0
	// HistoryString represents string values
	HistoryString HistoryValueType = 1
	// HistoryLog represents log values
	HistoryLog HistoryValueType = 2
	// HistoryUnsigned represents unsigned integer values
	HistoryUnsigned HistoryValueType = 3
	// HistoryText represents text values
	HistoryText HistoryValueType = 4
	// HistoryBinary represents binary values (not supported)
	HistoryBinary HistoryValueType = 5
)

// HistoryData represents a single history data point
type HistoryData struct {
	ItemID string `json:"itemid"`
	Clock  int    `json:"clock"`
	Value  string `json:"value"`
	NS     int    `json:"ns"`
}

// HistoryDataSlice represents a slice of history data
type HistoryDataSlice []HistoryData

// HistoryGetOptions represents parameters for history.get API call
type HistoryGetOptions struct {
	ItemIDs      []string           `json:"itemids,omitempty"`
	History      HistoryValueType   `json:"history,string"`
	TimeFrom     int               `json:"time_from,omitempty"`
	TimeTill     int               `json:"time_till,omitempty"`
	Limit        int               `json:"limit,omitempty"`
	SortField    string            `json:"sortfield,omitempty"`
	SortOrder    string            `json:"sortorder,omitempty"`
	Filter       map[string]interface{} `json:"filter,omitempty"`
}

// History represents history data from different value types
type History struct {
	ItemID string `json:"itemid"`
	Clock  int    `json:"clock"`
	NS     int    `json:"ns"`
	
	// Different value types - only one will be populated based on history type
	Value     string `json:"value,omitempty"`      // for text/string
	ValueNum  float64 `json:"value_num,omitempty"` // for float/unsigned
	ValueStr  string  `json:"value_str,omitempty"` // for text/string
	ValueBool int     `json:"value_int,omitempty"` // for unsigned (boolean)
	
	// Log specific fields
	Level    string `json:"level,omitempty"`
	Severity string `json:"severity,omitempty"`
	Source   string `json:"source,omitempty"`
	LogEventID string `json:"logeventid,omitempty"`
}

// HistorySlice represents a slice of history data
type HistorySlice []History

// HistoryFloat represents float history data
type HistoryFloat struct {
	ItemID   string  `json:"itemid"`
	Clock    int     `json:"clock"`
	Value    float64 `json:"value"`
	NS       int     `json:"ns"`
}

// HistoryFloatSlice represents a slice of float history data
type HistoryFloatSlice []HistoryFloat

// HistoryUnsigned represents unsigned integer history data
type HistoryUnsigned struct {
	ItemID   string `json:"itemid"`
	Clock    int    `json:"clock"`
	Value    uint64 `json:"value"`
	NS       int    `json:"ns"`
}

// HistoryUnsignedSlice represents a slice of unsigned history data
type HistoryUnsignedSlice []HistoryUnsigned

// HistoryString represents string/text history data
type HistoryString struct {
	ItemID   string `json:"itemid"`
	Clock    int    `json:"clock"`
	Value    string `json:"value"`
	NS       int    `json:"ns"`
}

// HistoryStringSlice represents a slice of string history data
type HistoryStringSlice []HistoryString

// HistoryLog represents log history data
type HistoryLog struct {
	ItemID      string `json:"itemid"`
	Clock       int    `json:"clock"`
	Value       string `json:"value"`
	NS          int    `json:"ns"`
	Level       string `json:"severity,omitempty"`
	Severity    string `json:"source,omitempty"`
	Source      string `json:"logeventid,omitempty"`
	LogEventID  string `json:"timestamp,omitempty"`
	Timestamp   string `json:"log,omitempty"`
}

// HistoryLogSlice represents a slice of log history data
type HistoryLogSlice []HistoryLog

// HistoryBinary represents binary history data
type HistoryBinary struct {
	ItemID   string `json:"itemid"`
	Clock    int    `json:"clock"`
	Value    string `json:"value"`
	NS       int    `json:"ns"`
}

// HistoryBinarySlice represents a slice of binary history data
type HistoryBinarySlice []HistoryBinary

// HistoryGet Generic method to get history data
// https://www.zabbix.com/documentation/current/manual/api/reference/history/get
func (api *API) HistoryGet(params HistoryGetOptions) (HistorySlice, error) {
	
	// Prepare parameters for API call
	apiParams := make(map[string]interface{})
	
	if len(params.ItemIDs) > 0 {
		apiParams["itemids"] = params.ItemIDs
	}
	
	apiParams["history"] = params.History
	
	if params.TimeFrom > 0 {
		apiParams["time_from"] = params.TimeFrom
	}
	
	if params.TimeTill > 0 {
		apiParams["time_till"] = params.TimeTill
	}
	
	if params.Limit > 0 {
		apiParams["limit"] = params.Limit
	} else {
		// Set a reasonable default limit to prevent overwhelming responses
		apiParams["limit"] = 1000
	}
	
	if params.SortField != "" {
		apiParams["sortfield"] = params.SortField
	} else {
		// Default sort by clock (most recent first)
		apiParams["sortfield"] = "clock"
	}
	
	if params.SortOrder != "" {
		apiParams["sortorder"] = params.SortOrder
	} else {
		// Default sort order (newest first)
		apiParams["sortorder"] = "DESC"
	}
	
	if params.Filter != nil {
		apiParams["filter"] = params.Filter
	}

	var history HistorySlice
	err := api.CallWithErrorParse("history.get", apiParams, &history)
	return history, err
}

// HistoryGetByItem Get history data for a specific item
func (api *API) HistoryGetByItem(itemID string, historyType HistoryValueType, timeFrom, timeTill int, limit int) (HistorySlice, error) {
	params := HistoryGetOptions{
		ItemIDs:    []string{itemID},
		History:    historyType,
		TimeFrom:   timeFrom,
		TimeTill:   timeTill,
		Limit:      limit,
		SortField:  "clock",
		SortOrder:  "DESC",
	}
	
	return api.HistoryGet(params)
}

// HistoryGetFloat Get float history data
// https://www.zabbix.com/documentation/current/manual/api/reference/history/get
func (api *API) HistoryGetFloat(params HistoryGetOptions) (HistoryFloatSlice, error) {
	params.History = HistoryFloat
	
	var history HistoryFloatSlice
	err := api.CallWithErrorParse("history.get", params, &history)
	return history, err
}

// HistoryGetFloatByItem Get float history data for a specific item
func (api *API) HistoryGetFloatByItem(itemID string, timeFrom, timeTill int, limit int) (HistoryFloatSlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   HistoryFloat,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.HistoryGetFloat(params)
}

// HistoryGetUnsigned Get unsigned integer history data
// https://www.zabbix.com/documentation/current/manual/api/reference/history/get
func (api *API) HistoryGetUnsigned(params HistoryGetOptions) (HistoryUnsignedSlice, error) {
	params.History = HistoryUnsigned
	
	var history HistoryUnsignedSlice
	err := api.CallWithErrorParse("history.get", params, &history)
	return history, err
}

// HistoryGetUnsignedByItem Get unsigned history data for a specific item
func (api *API) HistoryGetUnsignedByItem(itemID string, timeFrom, timeTill int, limit int) (HistoryUnsignedSlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   HistoryUnsigned,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.HistoryGetUnsigned(params)
}

// HistoryGetString Get string/text history data
// https://www.zabbix.com/documentation/current/manual/api/reference/history/get
func (api *API) HistoryGetString(params HistoryGetOptions) (HistoryStringSlice, error) {
	params.History = HistoryString
	
	var history HistoryStringSlice
	err := api.CallWithErrorParse("history.get", params, &history)
	return history, err
}

// HistoryGetStringByItem Get string history data for a specific item
func (api *API) HistoryGetStringByItem(itemID string, timeFrom, timeTill int, limit int) (HistoryStringSlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   HistoryString,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.HistoryGetString(params)
}

// HistoryGetText Get text history data
// https://www.zabbix.com/documentation/current/manual/api/reference/history/get
func (api *API) HistoryGetText(params HistoryGetOptions) (HistoryStringSlice, error) {
	params.History = HistoryText
	
	var history HistoryStringSlice
	err := api.CallWithErrorParse("history.get", params, &history)
	return history, err
}

// HistoryGetTextByItem Get text history data for a specific item
func (api *API) HistoryGetTextByItem(itemID string, timeFrom, timeTill int, limit int) (HistoryStringSlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   HistoryText,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.HistoryGetText(params)
}

// HistoryGetLog Get log history data
// https://www.zabbix.com/documentation/current/manual/api/reference/history/get
func (api *API) HistoryGetLog(params HistoryGetOptions) (HistoryLogSlice, error) {
	params.History = HistoryLog
	
	var history HistoryLogSlice
	err := api.CallWithErrorParse("history.get", params, &history)
	return history, err
}

// HistoryGetLogByItem Get log history data for a specific item
func (api *API) HistoryGetLogByItem(itemID string, timeFrom, timeTill int, limit int) (HistoryLogSlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   HistoryLog,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.HistoryGetLog(params)
}

// HistoryGetBinary Get binary history data
// Note: Binary history is not supported in most Zabbix configurations
// https://www.zabbix.com/documentation/current/manual/api/reference/history/get
func (api *API) HistoryGetBinary(params HistoryGetOptions) (HistoryBinarySlice, error) {
	params.History = HistoryBinary
	
	var history HistoryBinarySlice
	err := api.CallWithErrorParse("history.get", params, &history)
	return history, err
}

// HistoryGetBinaryByItem Get binary history data for a specific item
func (api *API) HistoryGetBinaryByItem(itemID string, timeFrom, timeTill int, limit int) (HistoryBinarySlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   HistoryBinary,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.HistoryGetBinary(params)
}

// HistoryGetByTimeRange Get history data for multiple items within a time range
func (api *API) HistoryGetByTimeRange(itemIDs []string, historyType HistoryValueType, timeFrom, timeTill int, limit int) (HistorySlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   itemIDs,
		History:   historyType,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
		SortField: "clock",
		SortOrder: "DESC",
	}
	
	return api.HistoryGet(params)
}

// HistoryGetRecent Get recent history data (last hour by default)
func (api *API) HistoryGetRecent(itemID string, historyType HistoryValueType, limit int) (HistorySlice, error) {
	// Get current timestamp
	currentTime := int(getCurrentTimestamp())
	oneHourAgo := currentTime - 3600
	
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   historyType,
		TimeFrom:  oneHourAgo,
		TimeTill:  currentTime,
		Limit:     limit,
		SortField: "clock",
		SortOrder: "DESC",
	}
	
	return api.HistoryGet(params)
}

// HistoryGetLatest Get the latest history data point for an item
func (api *API) HistoryGetLatest(itemID string, historyType HistoryValueType) (HistorySlice, error) {
	params := HistoryGetOptions{
		ItemIDs:   []string{itemID},
		History:   historyType,
		Limit:     1,
		SortField: "clock",
		SortOrder: "DESC",
	}
	
	return api.HistoryGet(params)
}

// HistoryGetStats Get statistics for history data
func (api *API) HistoryGetStats(itemID string, historyType HistoryValueType, timeFrom, timeTill int) (map[string]interface{}, error) {
	historyData, err := api.HistoryGetByTimeRange([]string{itemID}, historyType, timeFrom, timeTill, 0)
	if err != nil {
		return nil, err
	}
	
	stats := make(map[string]interface{})
	
	switch historyType {
	case HistoryFloat, HistoryUnsigned:
		if len(historyData) > 0 {
			stats["count"] = len(historyData)
			
			// Calculate basic statistics for numeric data
			values := make([]float64, 0, len(historyData))
			for _, h := range historyData {
				if h.ValueNum != 0 {
					values = append(values, h.ValueNum)
				}
			}
			
			if len(values) > 0 {
				stats["min"] = calculateMin(values)
				stats["max"] = calculateMax(values)
				stats["avg"] = calculateAvg(values)
			}
		}
	case HistoryString, HistoryText:
		if len(historyData) > 0 {
			stats["count"] = len(historyData)
			// For string data, we could add unique value count, most common value, etc.
			uniqueValues := make(map[string]int)
			for _, h := range historyData {
				uniqueValues[h.ValueStr]++
			}
			stats["unique_values"] = len(uniqueValues)
		}
	case HistoryLog:
		if len(historyData) > 0 {
			stats["count"] = len(historyData)
			// For log data, we could analyze severity levels, sources, etc.
			severityCount := make(map[string]int)
			sourceCount := make(map[string]int)
			for _, h := range historyData {
				if h.Severity != "" {
					severityCount[h.Severity]++
				}
				if h.Source != "" {
					sourceCount[h.Source]++
				}
			}
			stats["severity_distribution"] = severityCount
			stats["source_distribution"] = sourceCount
		}
	}
	
	return stats, nil
}

// Helper function to get current timestamp
func getCurrentTimestamp() int64 {
	// This is a simplified implementation
	// In production, you might want to use time.Now().Unix()
	return 1640995200 // Example timestamp
}

// Helper functions for statistics calculation
func calculateMin(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func calculateMax(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func calculateAvg(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Convert HistorySlice to specific type based on value type
func (history HistorySlice) ToFloatHistory() (HistoryFloatSlice, error) {
	floatHistory := make(HistoryFloatSlice, 0, len(history))
	
	for _, h := range history {
		if h.ValueNum != 0 {
			floatHistory = append(floatHistory, HistoryFloat{
				ItemID: h.ItemID,
				Clock:  h.Clock,
				Value:  h.ValueNum,
				NS:     h.NS,
			})
		}
	}
	
	return floatHistory, nil
}

// Convert HistorySlice to string history
func (history HistorySlice) ToStringHistory() (HistoryStringSlice, error) {
	stringHistory := make(HistoryStringSlice, 0, len(history))
	
	for _, h := range history {
		if h.ValueStr != "" {
			stringHistory = append(stringHistory, HistoryString{
				ItemID: h.ItemID,
				Clock:  h.Clock,
				Value:  h.ValueStr,
				NS:     h.NS,
			})
		}
	}
	
	return stringHistory, nil
}

// Convert HistorySlice to log history
func (history HistorySlice) ToLogHistory() (HistoryLogSlice, error) {
	logHistory := make(HistoryLogSlice, 0, len(history))
	
	for _, h := range history {
		if h.Level != "" || h.Source != "" {
			logHistory = append(logHistory, HistoryLog{
				ItemID:      h.ItemID,
				Clock:       h.Clock,
				Value:       h.Value,
				NS:          h.NS,
				Level:       h.Level,
				Severity:    h.Severity,
				Source:      h.Source,
				LogEventID:  h.LogEventID,
				Timestamp:   h.Value,
			})
		}
	}
	
	return logHistory, nil
}