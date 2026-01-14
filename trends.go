package zabbix

import (
	"encoding/json"
	"fmt"
	"math"
)

// TrendValueType represents the type of trend data
type TrendValueType int

const (
	// TrendFloat represents float values
	TrendFloat TrendValueType = 0
	// TrendUnsigned represents unsigned integer values
	TrendUnsigned TrendValueType = 3
)

// Trend represents a single trend data point
type Trend struct {
	ItemID   string `json:"itemid"`
	Clock    int     `json:"clock"`
	ValueMin string  `json:"value_min"`
	ValueAvg string  `json:"value_avg"`
	ValueMax string  `json:"value_max"`
	NS       int     `json:"ns"`
}

// TrendSlice represents a slice of trend data
type TrendSlice []Trend

// TrendGetOptions represents parameters for trends.get API call
type TrendGetOptions struct {
	ItemIDs      []string           `json:"itemids,omitempty"`
	TrendType    TrendValueType     `json:"trend_type,string"`
	TimeFrom     int               `json:"time_from,omitempty"`
	TimeTill     int               `json:"time_till,omitempty"`
	Limit        int               `json:"limit,omitempty"`
	SortField    string            `json:"sortfield,omitempty"`
	SortOrder    string            `json:"sortorder,omitempty"`
	Filter       map[string]interface{} `json:"filter,omitempty"`
}

// TrendFloat represents float trend data
type TrendFloat struct {
	ItemID   string  `json:"itemid"`
	Clock    int     `json:"clock"`
	ValueMin float64 `json:"value_min"`
	ValueAvg float64 `json:"value_avg"`
	ValueMax float64 `json:"value_max"`
	NS       int     `json:"ns"`
}

// TrendFloatSlice represents a slice of float trend data
type TrendFloatSlice []TrendFloat

// TrendUnsigned represents unsigned integer trend data
type TrendUnsigned struct {
	ItemID   string `json:"itemid"`
	Clock    int    `json:"clock"`
	ValueMin uint64 `json:"value_min"`
	ValueAvg uint64 `json:"value_avg"`
	ValueMax uint64 `json:"value_max"`
	NS       int    `json:"ns"`
}

// TrendUnsignedSlice represents a slice of unsigned trend data
type TrendUnsignedSlice []TrendUnsigned

// TrendSummary represents trend statistics summary
type TrendSummary struct {
	ItemID       string  `json:"itemid"`
	Count        int     `json:"count"`
	MinValue     float64 `json:"min_value"`
	MaxValue     float64 `json:"max_value"`
	AvgValue     float64 `json:"avg_value"`
	MinClock     int     `json:"min_clock"`
	MaxClock     int     `json:"max_clock"`
	TimeRange    int     `json:"time_range"` // in seconds
	Periodicity  float64 `json:"periodicity"` // average seconds between data points
}

// TrendSummarySlice represents a slice of trend summaries
type TrendSummarySlice []TrendSummary

// TrendGet Generic method to get trend data
// https://www.zabbix.com/documentation/current/manual/api/reference/trends/get
func (api *API) TrendGet(params TrendGetOptions) (TrendSlice, error) {
	
	// Prepare parameters for API call
	apiParams := make(map[string]interface{})
	
	if len(params.ItemIDs) > 0 {
		apiParams["itemids"] = params.ItemIDs
	}
	
	if params.TrendType > 0 {
		apiParams["trend_type"] = params.TrendType
	}
	
	if params.TimeFrom > 0 {
		apiParams["time_from"] = params.TimeFrom
	}
	
	if params.TimeTill > 0 {
		apiParams["time_till"] = params.TimeTill
	}
	
	if params.Limit > 0 {
		apiParams["limit"] = params.Limit
	} else {
		// Set a reasonable default limit
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

	var trends TrendSlice
	err := api.CallWithErrorParse("trends.get", apiParams, &trends)
	return trends, err
}

// TrendGetByItem Get trend data for a specific item
func (api *API) TrendGetByItem(itemID string, trendType TrendValueType, timeFrom, timeTill int, limit int) (TrendSlice, error) {
	params := TrendGetOptions{
		ItemIDs:    []string{itemID},
		TrendType:  trendType,
		TimeFrom:   timeFrom,
		TimeTill:   timeTill,
		Limit:      limit,
		SortField:  "clock",
		SortOrder:  "DESC",
	}
	
	return api.TrendGet(params)
}

// TrendGetFloat Get float trend data
// https://www.zabbix.com/documentation/current/manual/api/reference/trends/get
func (api *API) TrendGetFloat(params TrendGetOptions) (TrendFloatSlice, error) {
	params.TrendType = TrendFloat
	
	var trends TrendFloatSlice
	err := api.CallWithErrorParse("trends.get", params, &trends)
	return trends, err
}

// TrendGetFloatByItem Get float trend data for a specific item
func (api *API) TrendGetFloatByItem(itemID string, timeFrom, timeTill int, limit int) (TrendFloatSlice, error) {
	params := TrendGetOptions{
		ItemIDs:   []string{itemID},
		TrendType: TrendFloat,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.TrendGetFloat(params)
}

// TrendGetUnsigned Get unsigned integer trend data
// https://www.zabbix.com/documentation/current/manual/api/reference/trends/get
func (api *API) TrendGetUnsigned(params TrendGetOptions) (TrendUnsignedSlice, error) {
	params.TrendType = TrendUnsigned
	
	var trends TrendUnsignedSlice
	err := api.CallWithErrorParse("trends.get", params, &trends)
	return trends, err
}

// TrendGetUnsignedByItem Get unsigned trend data for a specific item
func (api *API) TrendGetUnsignedByItem(itemID string, timeFrom, timeTill int, limit int) (TrendUnsignedSlice, error) {
	params := TrendGetOptions{
		ItemIDs:   []string{itemID},
		TrendType: TrendUnsigned,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
	}
	
	return api.TrendGetUnsigned(params)
}

// TrendGetByTimeRange Get trend data for multiple items within a time range
func (api *API) TrendGetByTimeRange(itemIDs []string, trendType TrendValueType, timeFrom, timeTill int, limit int) (TrendSlice, error) {
	params := TrendGetOptions{
		ItemIDs:   itemIDs,
		TrendType: trendType,
		TimeFrom:  timeFrom,
		TimeTill:  timeTill,
		Limit:     limit,
		SortField: "clock",
		SortOrder: "DESC",
	}
	
	return api.TrendGet(params)
}

// TrendGetHourly Get hourly trend data (last 24 hours by default)
func (api *API) TrendGetHourly(itemID string, trendType TrendValueType, hours int) (TrendSlice, error) {
	currentTime := int(getCurrentTimestamp())
	timeFrom := currentTime - (hours * 3600)
	
	params := TrendGetOptions{
		ItemIDs:   []string{itemID},
		TrendType: trendType,
		TimeFrom:  timeFrom,
		TimeTill:  currentTime,
		Limit:     hours, // One data point per hour
		SortField: "clock",
		SortOrder: "DESC",
	}
	
	return api.TrendGet(params)
}

// TrendGetDaily Get daily trend data (last 30 days by default)
func (api *API) TrendGetDaily(itemID string, trendType TrendValueType, days int) (TrendSlice, error) {
	currentTime := int(getCurrentTimestamp())
	timeFrom := currentTime - (days * 86400)
	
	params := TrendGetOptions{
		ItemIDs:   []string{itemID},
		TrendType: trendType,
		TimeFrom:  timeFrom,
		TimeTill:  currentTime,
		Limit:     days, // One data point per day
		SortField: "clock",
		SortOrder: "DESC",
	}
	
	return api.TrendGet(params)
}

// TrendGetLatest Get the latest trend data point for an item
func (api *API) TrendGetLatest(itemID string, trendType TrendValueType) (TrendSlice, error) {
	params := TrendGetOptions{
		ItemIDs:   []string{itemID},
		TrendType: trendType,
		Limit:     1,
		SortField: "clock",
		SortOrder: "DESC",
	}
	
	return api.TrendGet(params)
}

// TrendGetSummary Get trend summary statistics for an item
func (api *API) TrendGetSummary(itemID string, trendType TrendValueType, timeFrom, timeTill int) (TrendSummary, error) {
	trends, err := api.TrendGetByTimeRange([]string{itemID}, trendType, timeFrom, timeTill, 0)
	if err != nil {
		return TrendSummary{}, err
	}
	
	if len(trends) == 0 {
		return TrendSummary{
			ItemID: itemID,
			Count:  0,
		}, nil
	}
	
	// Calculate statistics
	summary := TrendSummary{
		ItemID:     itemID,
		Count:      len(trends),
		MinClock:   trends[0].Clock,
		MaxClock:   trends[0].Clock,
		TimeRange:  timeTill - timeFrom,
	}
	
	// Initialize with first values
	var minVal, maxVal, sumVal float64
	var hasValues bool
	
	for i, trend := range trends {
		if i == 0 {
			summary.MinClock = trend.Clock
			summary.MaxClock = trend.Clock
		}
		
		if trend.Clock < summary.MinClock {
			summary.MinClock = trend.Clock
		}
		if trend.Clock > summary.MaxClock {
			summary.MaxClock = trend.Clock
		}
		
		// Parse numeric values
		if trend.ValueMin != "" {
			if val, err := parseFloat64(trend.ValueMin); err == nil {
				if !hasValues {
					minVal = val
					maxVal = val
					sumVal = val
					hasValues = true
				} else {
					if val < minVal {
						minVal = val
					}
					if val > maxVal {
						maxVal = val
					}
					sumVal += val
				}
			}
		}
		
		if trend.ValueAvg != "" {
			if val, err := parseFloat64(trend.ValueAvg); err == nil {
				if !hasValues {
					minVal = val
					maxVal = val
					sumVal = val
					hasValues = true
				} else {
					if val < minVal {
						minVal = val
					}
					if val > maxVal {
						maxVal = val
					}
					sumVal += val
				}
			}
		}
		
		if trend.ValueMax != "" {
			if val, err := parseFloat64(trend.ValueMax); err == nil {
				if !hasValues {
					minVal = val
					maxVal = val
					sumVal = val
					hasValues = true
				} else {
					if val < minVal {
						minVal = val
					}
					if val > maxVal {
						maxVal = val
					}
					sumVal += val
				}
			}
		}
	}
	
	if hasValues {
		summary.MinValue = minVal
		summary.MaxValue = maxVal
		summary.AvgValue = sumVal / float64(len(trends)*3) // Divide by number of value types
	}
	
	// Calculate periodicity (average seconds between data points)
	if summary.Count > 1 && summary.TimeRange > 0 {
		summary.Periodicity = float64(summary.TimeRange) / float64(summary.Count-1)
	}
	
	return summary, nil
}

// TrendGetMultipleSummaries Get trend summaries for multiple items
func (api *API) TrendGetMultipleSummaries(itemIDs []string, trendType TrendValueType, timeFrom, timeTill int) (TrendSummarySlice, error) {
	summaries := make(TrendSummarySlice, 0, len(itemIDs))
	
	for _, itemID := range itemIDs {
		summary, err := api.TrendGetSummary(itemID, trendType, timeFrom, timeTill)
		if err != nil {
			continue // Skip items with errors
		}
		summaries = append(summaries, summary)
	}
	
	return summaries, nil
}

// TrendGetComparison Compare trends between two time periods
func (api *API) TrendGetComparison(itemID string, trendType TrendValueType, 
	period1From, period1Till, period2From, period2Till int) (map[string]interface{}, error) {
	
	// Get trends for both periods
	trends1, err := api.TrendGetByTimeRange([]string{itemID}, trendType, period1From, period1Till, 0)
	if err != nil {
		return nil, err
	}
	
	trends2, err := api.TrendGetByTimeRange([]string{itemID}, trendType, period2From, period2Till, 0)
	if err != nil {
		return nil, err
	}
	
	// Calculate summaries for both periods
	summary1, err := api.TrendGetSummary(itemID, trendType, period1From, period1Till)
	if err != nil {
		return nil, err
	}
	
	summary2, err := api.TrendGetSummary(itemID, trendType, period2From, period2Till)
	if err != nil {
		return nil, err
	}
	
	// Build comparison result
	comparison := map[string]interface{}{
		"itemid": itemID,
		"period1": map[string]interface{}{
			"time_from": period1From,
			"time_till": period1Till,
			"summary":   summary1,
			"data_points": len(trends1),
		},
		"period2": map[string]interface{}{
			"time_from": period2From,
			"time_till": period2Till,
			"summary":   summary2,
			"data_points": len(trends2),
		},
		"changes": map[string]interface{}{},
	}
	
	// Calculate changes between periods
	if summary1.Count > 0 && summary2.Count > 0 {
		changes := comparison["changes"].(map[string]interface{})
		
		// Calculate percentage changes
		if summary1.AvgValue != 0 {
			changes["avg_change_percent"] = ((summary2.AvgValue - summary1.AvgValue) / summary1.AvgValue) * 100
		}
		
		if summary1.MaxValue != 0 {
			changes["max_change_percent"] = ((summary2.MaxValue - summary1.MaxValue) / summary1.MaxValue) * 100
		}
		
		if summary1.MinValue != 0 {
			changes["min_change_percent"] = ((summary2.MinValue - summary1.MinValue) / summary1.MinValue) * 100
		}
		
		changes["data_points_change"] = summary2.Count - summary1.Count
	}
	
	return comparison, nil
}

// TrendGetAnomalies Detect anomalies in trend data
func (api *API) TrendGetAnomalies(itemID string, trendType TrendValueType, 
	timeFrom, timeTill int, stdDevThreshold float64) ([]Trend, error) {
	
	trends, err := api.TrendGetByTimeRange([]string{itemID}, trendType, timeFrom, timeTill, 0)
	if err != nil {
		return nil, err
	}
	
	if len(trends) < 3 {
		return nil, nil // Need at least 3 points for anomaly detection
	}
	
	// Extract values for statistical analysis
	values := make([]float64, 0, len(trends))
	for _, trend := range trends {
		if trend.ValueAvg != "" {
			if val, err := parseFloat64(trend.ValueAvg); err == nil {
				values = append(values, val)
			}
		}
	}
	
	if len(values) < 3 {
		return nil, nil
	}
	
	// Calculate mean and standard deviation
	mean := calculateAvg(values)
	stdDev := calculateStandardDeviation(values, mean)
	
	threshold := stdDev * stdDevThreshold
	
	// Find anomalies
	var anomalies []Trend
	for _, trend := range trends {
		if trend.ValueAvg != "" {
			if val, err := parseFloat64(trend.ValueAvg); err == nil {
				if abs(val-mean) > threshold {
					anomalies = append(anomalies, trend)
				}
			}
		}
	}
	
	return anomalies, nil
}

// Helper function to parse float64 from string
func parseFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err
}

// Helper function to calculate absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Helper function to calculate standard deviation
func calculateStandardDeviation(values []float64, mean float64) float64 {
	if len(values) == 0 {
		return 0
	}
	
	variance := 0.0
	for _, value := range values {
		variance += (value - mean) * (value - mean)
	}
	variance /= float64(len(values))
	
	return sqrt(variance)
}

// Helper function to calculate square root
func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Convert TrendSlice to typed trend data
func (trends TrendSlice) ToFloatTrends() (TrendFloatSlice, error) {
	floatTrends := make(TrendFloatSlice, 0, len(trends))
	
	for _, trend := range trends {
		minVal, _ := parseFloat64(trend.ValueMin)
		avgVal, _ := parseFloat64(trend.ValueAvg)
		maxVal, _ := parseFloat64(trend.ValueMax)
		
		floatTrends = append(floatTrends, TrendFloat{
			ItemID:   trend.ItemID,
			Clock:    trend.Clock,
			ValueMin: minVal,
			ValueAvg: avgVal,
			ValueMax: maxVal,
			NS:       trend.NS,
		})
	}
	
	return floatTrends, nil
}