package zabbix

import (
	"testing"
	"fmt"
)

// TestVersionManager 测试版本管理器
func TestVersionManager(t *testing.T) {
	vm := NewVersionManager()
	
	// 测试版本解析
	testCases := []struct {
		version    string
		expected60 bool
		expected70 bool
	}{
		{"6.0.0", true, false},
		{"6.4.5", true, false},
		{"7.0.0", false, true},
		{"7.2.1", false, true},
		{"5.0.0", false, false},
	}
	
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Version_%s", tc.version), func(t *testing.T) {
			vm.ForceVersion(tc.version)
			
			if vm.Is60() != tc.expected60 {
				t.Errorf("Expected Is60() to be %v for version %s, got %v", tc.expected60, tc.version, vm.Is60())
			}
			
			if vm.Is70() != tc.expected70 {
				t.Errorf("Expected Is70() to be %v for version %s, got %v", tc.expected70, tc.version, vm.Is70())
			}
		})
	}
}

// TestFeatureDetection 测试特性检测
func TestFeatureDetection(t *testing.T) {
	vm := NewVersionManager()
	
	// 测试 Zabbix 6.0 特性
	vm.ForceVersion("6.4.0")
	
	if vm.IsFeatureSupported(FeatureHistoryPush) {
		t.Error("History Push should not be supported in Zabbix 6.0")
	}
	
	if vm.IsFeatureSupported(FeatureMFA) {
		t.Error("MFA should not be supported in Zabbix 6.0")
	}
	
	if vm.IsFeatureSupported(FeatureHeadersArrayFormat) {
		t.Error("Headers V7 should not be supported in Zabbix 6.0")
	}
	
	// 测试 Zabbix 7.0 特性
	vm.ForceVersion("7.0.0")
	
	if !vm.IsFeatureSupported(FeatureHistoryPush) {
		t.Error("History Push should be supported in Zabbix 7.0")
	}
	
	if !vm.IsFeatureSupported(FeatureMFA) {
		t.Error("MFA should be supported in Zabbix 7.0")
	}
	
	if !vm.IsFeatureSupported(FeatureHeadersArrayFormat) {
		t.Error("Headers V7 should be supported in Zabbix 7.0")
	}
	
	if !vm.IsFeatureSupported(FeatureBrowserItem) {
		t.Error("Browser Item should be supported in Zabbix 7.0")
	}
}

// TestHeaderConversion 测试 Header 格式转换
func TestHeaderConversion(t *testing.T) {
	// 测试 6.0 到 7.0 转换
	headersV6 := HttpHeaders{
		"User-Agent": "Zabbix Monitoring",
		"Accept":     "text/html",
		"X-Custom":   "custom-value",
	}
	
	headersV7 := ConvertHeadersToV7(headersV6)
	
	if len(headersV7) != len(headersV6) {
		t.Errorf("Expected %d headers in V7 format, got %d", len(headersV6), len(headersV7))
	}
	
	// 验证转换结果
	expectedHeaders := map[string]string{
		"User-Agent": "Zabbix Monitoring",
		"Accept":     "text/html",
		"X-Custom":   "custom-value",
	}
	
	for _, header := range headersV7 {
		expectedValue, exists := expectedHeaders[header.Name]
		if !exists {
			t.Errorf("Unexpected header name: %s", header.Name)
		}
		if header.Value != expectedValue {
			t.Errorf("Expected header value %s for %s, got %s", expectedValue, header.Name, header.Value)
		}
	}
	
	// 测试 7.0 到 6.0 转换
	convertedV6 := ConvertHeadersToV6(headersV7)
	
	if len(convertedV6) != len(headersV6) {
		t.Errorf("Expected %d headers in converted V6 format, got %d", len(headersV6), len(convertedV6))
	}
	
	// 验证往返转换
	for name, value := range headersV6 {
		convertedValue, exists := convertedV6[name]
		if !exists {
			t.Errorf("Header %s lost during conversion", name)
		}
		if convertedValue != value {
			t.Errorf("Header %s value changed during conversion: expected %s, got %s", name, value, convertedValue)
		}
	}
}

// TestItemValidation 测试 Item 验证
func TestItemValidation(t *testing.T) {
	// 测试 Zabbix 6.0 验证
	item := Item{
		Type:    WebItem,
		HeadersV6: HttpHeaders{"User-Agent": "Test"},
	}
	
	err := ValidateItemForVersion(item, "6.4.0")
	if err != nil {
		t.Errorf("Item should be valid for Zabbix 6.0: %v", err)
	}
	
	// 测试 Zabbix 7.0 验证
	item7 := Item{
		Type:      Browser,
		HeadersV7: []HeaderField{{Name: "User-Agent", Value: "Test"}},
	}
	
	err = ValidateItemForVersion(item7, "7.0.0")
	if err != nil {
		t.Errorf("Browser item should be valid for Zabbix 7.0: %v", err)
	}
	
	// 测试无效的 Browser Item 在 6.0 中
	err = ValidateItemForVersion(item7, "6.4.0")
	if err == nil {
		t.Error("Browser item should not be valid for Zabbix 6.0")
	}
}

// TestHostValidation 测试 Host 验证
func TestHostValidation(t *testing.T) {
	// 测试 Zabbix 6.0 验证
	host := Host{
		Host:       "test-host",
		ProxyHostID: "10085",
	}
	
	err := ValidateHostForVersion(host, "6.4.0")
	if err != nil {
		t.Errorf("Host should be valid for Zabbix 6.0: %v", err)
	}
	
	// 测试 Zabbix 7.0 验证 (正确配置)
	host7 := Host{
		Host:        "test-host",
		ProxyID:     "10085",
		MonitoredBy: MonitoredByProxy,
	}
	
	err = ValidateHostForVersion(host7, "7.0.0")
	if err != nil {
		t.Errorf("Host should be valid for Zabbix 7.0: %v", err)
	}
	
	// 测试 Zabbix 7.0 验证 (缺少 MonitoredBy)
	host7Invalid := Host{
		Host:    "test-host",
		ProxyID: "10085",
		// MonitoredBy 缺失
	}
	
	err = ValidateHostForVersion(host7Invalid, "7.0.0")
	if err == nil {
		t.Error("Host should be invalid for Zabbix 7.0 without MonitoredBy")
	}
}

// TestBrowserItemValidation 测试 Browser Item 验证
func TestBrowserItemValidation(t *testing.T) {
	// 测试有效的 Browser Item
	validItem := BrowserItem{
		Item: Item{
			Type:          Browser,
			Name:          "Test Browser Item",
			Key:           "browser.test[example.com]",
			BrowserScript: "return document.title;",
		},
	}
	
	err := ValidateBrowserItem(validItem)
	if err != nil {
		t.Errorf("Browser item should be valid: %v", err)
	}
	
	// 测试缺少脚本的 Browser Item
	invalidItem := BrowserItem{
		Item: Item{
			Type: Browser,
			Name: "Invalid Browser Item",
			Key:  "browser.test[example.com]",
		},
		// BrowserScript 缺失
	}
	
	err = ValidateBrowserItem(invalidItem)
	if err == nil {
		t.Error("Browser item without script should be invalid")
	}
	
	// 测试类型错误的 Browser Item
	invalidTypeItem := BrowserItem{
		Item: Item{
			Type: WebItem, // 错误的类型
			Name: "Invalid Type Browser Item",
			Key:  "browser.test[example.com]",
			BrowserScript: "return document.title;",
		},
	}
	
	err = ValidateBrowserItem(invalidTypeItem)
	if err == nil {
		t.Error("Browser item with wrong type should be invalid")
	}
}

// TestAdapterInterface 测试适配器接口
func TestAdapterInterface(t *testing.T) {
	// 创建模拟 API (在实际测试中需要真实的 Zabbix 服务器)
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试 Zabbix 6.0 适配器
	err := api.ForceVersion("6.4.0")
	if err != nil {
		t.Fatalf("Failed to force version: %v", err)
	}
	
	if _, ok := api.GetItemAdapter().(*Zabbix6ItemAdapter); !ok {
		t.Error("Expected Zabbix 6.0 item adapter")
	}
	if _, ok := api.GetHostAdapter().(*Zabbix6HostAdapter); !ok {
		t.Error("Expected Zabbix 6.0 host adapter")
	}
	
	// 测试 Zabbix 7.0 适配器
	err = api.ForceVersion("7.0.0")
	if err != nil {
		t.Fatalf("Failed to force version: %v", err)
	}
	
	if _, ok := api.GetItemAdapter().(*Zabbix7ItemAdapter); !ok {
		t.Error("Expected Zabbix 7.0 item adapter")
	}
	if _, ok := api.GetHostAdapter().(*Zabbix7HostAdapter); !ok {
		t.Error("Expected Zabbix 7.0 host adapter")
	}
}

// TestFeatureConstants 测试特性常量
func TestFeatureConstants(t *testing.T) {
	expectedFeatures := []string{
		FeatureUUID,
		FeatureTags,
		FeatureCompression,
		FeatureHTTPMethods,
		FeatureCalculatedItemTypes,
		FeatureHistoryPush,
		FeatureMFA,
		FeatureProxyGroup,
		FeatureBrowserItem,
		FeatureHeadersArrayFormat,
		FeatureProxyFieldsV7,
	}
	
	// 确保所有特性常量都有值
	for _, feature := range expectedFeatures {
		if feature == "" {
			t.Errorf("Feature constant should not be empty")
		}
	}
	
	// 确保特性常量是唯一的
	featureMap := make(map[string]bool)
	for _, feature := range expectedFeatures {
		if featureMap[feature] {
			t.Errorf("Duplicate feature constant: %s", feature)
		}
		featureMap[feature] = true
	}
}

// BenchmarkHeaderConversion 性能测试：Header 转换
func BenchmarkHeaderConversion(b *testing.B) {
	// 创建大量的 headers
	headersV6 := make(HttpHeaders)
	for i := 0; i < 1000; i++ {
		headersV6[fmt.Sprintf("Header-%d", i)] = fmt.Sprintf("Value-%d", i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		headersV7 := ConvertHeadersToV7(headersV6)
		_ = ConvertHeadersToV6(headersV7)
	}
}

// BenchmarkVersionDetection 性能测试：版本检测
func BenchmarkVersionDetection(b *testing.B) {
	vm := NewVersionManager()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vm.ForceVersion("7.0.0")
		_ = vm.IsFeatureSupported(FeatureHistoryPush)
		vm.ForceVersion("6.4.0")
		_ = vm.IsFeatureSupported(FeatureHistoryPush)
	}
}
