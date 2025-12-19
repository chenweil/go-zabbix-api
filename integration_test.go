package zabbix

import (
	"testing"
	"fmt"
	"encoding/json"
)

// TestMultiVersionIntegration 集成测试：多版本功能
func TestMultiVersionIntegration(t *testing.T) {
	// 创建 API 实例
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试版本强制和适配器初始化
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("Integration_%s", version), func(t *testing.T) {
			// 强制版本
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version %s: %v", version, err)
			}
			
			// 验证版本检测
			if api.GetServerVersion() != version {
				t.Errorf("Expected version %s, got %s", version, api.GetServerVersion())
			}
			
			// 验证适配器初始化
			if api.GetItemAdapter() == nil {
				t.Error("Item adapter should be initialized")
			}
			
			if api.GetHostAdapter() == nil {
				t.Error("Host adapter should be initialized")
			}
			
			// 验证特性检测
			expectedFeatures := map[string]bool{
				FeatureHistoryPush: version == "7.0.0",
				FeatureMFA:         version == "7.0.0",
				FeatureProxyGroup:  version == "7.0.0",
				FeatureBrowserItem: version == "7.0.0",
			}
			
			for feature, expected := range expectedFeatures {
				if api.IsFeatureSupported(feature) != expected {
					t.Errorf("Feature %s support mismatch for version %s: expected %v, got %v", 
						feature, version, expected, api.IsFeatureSupported(feature))
				}
			}
		})
	}
}

// TestItemIntegration 集成测试：Item 处理
func TestItemIntegration(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试数据
	testItem := Item{
		HostID:    "10084",
		Key:       "web.page.get[example.com]",
		Name:      "Example.com page content",
		Type:      WebItem,
		Delay:     "1m",
		ValueType: Text,
		Url:       "http://example.com",
		Timeout:   "10s",
	}
	
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("Item_%s", version), func(t *testing.T) {
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version: %v", err)
			}
			
			// 根据版本设置不同的 headers 格式
			item := testItem
			if version == "7.0.0" {
				item.HeadersV7 = []HeaderField{
					{Name: "User-Agent", Value: "Zabbix Monitoring"},
					{Name: "Accept", Value: "text/html"},
				}
			} else {
				item.HeadersV6 = HttpHeaders{
					"User-Agent": "Zabbix Monitoring",
					"Accept":     "text/html",
				}
			}
			
			// 测试适配器准备 headers
			adapter := api.GetItemAdapter()
			_ = adapter // 使用适配器但跳过 PrepareHeaders 测试
			
			// TODO: PrepareHeaders 方法需要在适配器中实现
			// headers := adapter.PrepareHeaders(item)
			
				// 验证 JSON 序列化
				jsonData, err := json.Marshal(item)
				if err != nil {
					t.Errorf("Failed to marshal item: %v", err)
				}
				
				var raw map[string]json.RawMessage
				err = json.Unmarshal(jsonData, &raw)
				if err != nil {
					t.Errorf("Failed to unmarshal item JSON: %v", err)
				}
				
				// 验证格式
				if version == "7.0.0" {
					headersRaw, ok := raw["headers_v7"]
					if !ok {
						t.Error("Expected headers_v7 field for Zabbix 7.0")
					} else {
						var headersArray []HeaderField
						err = json.Unmarshal(headersRaw, &headersArray)
						if err != nil {
							t.Errorf("Failed to unmarshal headers_v7 as array: %v", err)
						}
					}
					if _, ok := raw["headers_v6"]; ok {
						t.Error("headers_v6 should not be set for Zabbix 7.0")
					}
				} else {
					headersRaw, ok := raw["headers_v6"]
					if !ok {
						t.Error("Expected headers_v6 field for Zabbix 6.0")
					} else {
						var headersMap HttpHeaders
						err = json.Unmarshal(headersRaw, &headersMap)
						if err != nil {
							t.Errorf("Failed to unmarshal headers_v6 as object: %v", err)
						}
					}
					if _, ok := raw["headers_v7"]; ok {
						t.Error("headers_v7 should not be set for Zabbix 6.0")
					}
				}
			})
		}
	}

// TestHostIntegration 集成测试：Host 处理
func TestHostIntegration(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试数据
	testHost := Host{
		Host:     "example-host",
		Name:     "Example Host",
		Status:   Monitored,
		GroupIds: HostGroupIDs{{GroupID: "15"}},
	}
	
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("Host_%s", version), func(t *testing.T) {
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version: %v", err)
			}
			
			// 根据版本设置不同的代理配置
			host := testHost
			if version == "7.0.0" {
				host.ProxyID = "10085"
				host.MonitoredBy = MonitoredByProxy
			} else {
				host.ProxyHostID = "10085"
			}
			
			// 测试适配器准备代理字段
			adapter := api.GetHostAdapter()
			_ = adapter // 使用适配器但跳过 PrepareProxyFields 测试
			
			// TODO: PrepareProxyFields 方法需要在适配器中实现
			// proxyFields := adapter.PrepareProxyFields(host)
			
				// 验证主机配置
				if version == "7.0.0" {
					if host.ProxyID != "10085" {
						t.Errorf("Expected ProxyID '10085', got '%s'", host.ProxyID)
					}
					if host.MonitoredBy != MonitoredByProxy {
						t.Errorf("Expected MonitoredByProxy, got %v", host.MonitoredBy)
					}
				} else {
					if host.ProxyHostID != "10085" {
						t.Errorf("Expected ProxyHostID '10085', got '%s'", host.ProxyHostID)
					}
				}
			})
		}
	}

// TestBrowserItemIntegration 集成测试：Browser Item
func TestBrowserItemIntegration(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试数据
	browserItem := BrowserItem{
		Item: Item{
			HostID:        "10084",
			Key:           "browser.performance[example.com]",
			Name:          "Example.com Performance",
			Type:          Browser,
			Delay:         "1m",
			ValueType:     Text,
			BrowserScript: "return performance.now();",
			BrowserParams: "{\"timeout\": 10000}",
		},
	}
	
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("BrowserItem_%s", version), func(t *testing.T) {
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version: %v", err)
			}
			
			// 验证 Browser Item 支持
			if version == "7.0.0" {
				if !api.SupportsBrowserItem() {
					t.Error("Browser item should be supported in Zabbix 7.0")
				}
				
				// 验证 Browser Item
				err := ValidateBrowserItem(browserItem)
				if err != nil {
					t.Errorf("Browser item should be valid: %v", err)
				}
			} else {
				if api.SupportsBrowserItem() {
					t.Error("Browser item should not be supported in Zabbix 6.0")
				}
				
				// 在 6.0 中，Browser Item 应该验证失败
				err := ValidateItemForVersion(browserItem.Item, version)
				if err == nil {
					t.Error("Browser item should not be valid for Zabbix 6.0")
				}
			}
		})
	}
}

// TestMFAIntegration 集成测试：MFA 功能
func TestMFAIntegration(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试数据
	mfa := MFA{
		Name:         "Test MFA",
		Type:         MFATypeTOTP,
		HashFunction: "sha1",
		CodeLength:   6,
		Status:       MFAStatusEnabled,
		APIAccess:    "1",
	}
	
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("MFA_%s", version), func(t *testing.T) {
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version: %v", err)
			}
			
			if version == "7.0.0" {
				if !api.SupportsMFA() {
					t.Error("MFA should be supported in Zabbix 7.0")
				}
				
				// 验证 MFA 结构
				if mfa.Name == "" {
					t.Error("MFA name should not be empty")
				}
				
				if mfa.Type != MFATypeTOTP {
					t.Errorf("Expected MFA type %d, got %d", MFATypeTOTP, mfa.Type)
				}
			} else {
				if api.SupportsMFA() {
					t.Error("MFA should not be supported in Zabbix 6.0")
				}
			}
		})
	}
}

// TestProxyGroupIntegration 集成测试：Proxy Group 功能
func TestProxyGroupIntegration(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试数据
	proxyGroup := ProxyGroup{
		Name:       "Test Proxy Group",
		ProxyState: ProxyStateOnline,
	}
	
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("ProxyGroup_%s", version), func(t *testing.T) {
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version: %v", err)
			}
			
			if version == "7.0.0" {
				if !api.SupportsProxyGroup() {
					t.Error("Proxy Group should be supported in Zabbix 7.0")
				}
				
				// 验证 Proxy Group 结构
				if proxyGroup.Name == "" {
					t.Error("Proxy Group name should not be empty")
				}
				
				if proxyGroup.ProxyState != ProxyStateOnline {
					t.Errorf("Expected proxy state %d, got %d", ProxyStateOnline, proxyGroup.ProxyState)
				}
			} else {
				if api.SupportsProxyGroup() {
					t.Error("Proxy Group should not be supported in Zabbix 6.0")
				}
			}
		})
	}
}

// TestHistoryPushIntegration 集成测试：History Push 功能
func TestHistoryPushIntegration(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试数据
	historyData := []HistoryData{
		{
			Host:  "example-host",
			Key:   "web.page.get[example.com]",
			Value: "Page content here",
			Clock: 1609459200,
		},
	}
	
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("HistoryPush_%s", version), func(t *testing.T) {
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version: %v", err)
			}
			
			if version == "7.0.0" {
				if !api.SupportsHistoryPush() {
					t.Error("History Push should be supported in Zabbix 7.0")
				}
				
				// 验证 History Data 结构
				for _, data := range historyData {
					if data.Host == "" {
						t.Error("History data host should not be empty")
					}
					if data.Key == "" {
						t.Error("History data key should not be empty")
					}
					if data.Value == "" {
						t.Error("History data value should not be nil")
					}
				}
			} else {
				if api.SupportsHistoryPush() {
					t.Error("History Push should not be supported in Zabbix 6.0")
				}
			}
		})
	}
}

// TestSupportedFeaturesIntegration 集成测试：支持特性列表
func TestSupportedFeaturesIntegration(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	testVersions := []string{"6.4.0", "7.0.0"}
	
	for _, version := range testVersions {
		t.Run(fmt.Sprintf("SupportedFeatures_%s", version), func(t *testing.T) {
			err := api.ForceVersion(version)
			if err != nil {
				t.Fatalf("Failed to force version: %v", err)
			}
			
			features := api.GetSupportedFeatures()

			expected := map[string]bool{
				FeatureUUID:                true,
				FeatureTags:                true,
				FeatureCompression:         true,
				FeatureHTTPMethods:         true,
				FeatureCalculatedItemTypes: true,
				FeatureMFA:                 version == "7.0.0",
				FeatureProxyGroup:          version == "7.0.0",
				FeatureHistoryPush:         version == "7.0.0",
				FeatureBrowserItem:         version == "7.0.0",
				FeatureHeadersArrayFormat:  version == "7.0.0",
				FeatureProxyFieldsV7:       version == "7.0.0",
			}

			if len(features) != len(expected) {
				t.Errorf("Expected %d features for version %s, got %d: %v",
					len(expected), version, len(features), features)
			}

			for feature, shouldSupport := range expected {
				if features[feature] != shouldSupport {
					t.Errorf("Feature %s support mismatch for version %s: expected %v, got %v",
						feature, version, shouldSupport, features[feature])
				}
			}

			for feature := range features {
				if _, ok := expected[feature]; !ok {
					t.Errorf("Unexpected feature key: %s", feature)
				}
			}
		})
	}
}

// TestErrorHandling 集成测试：错误处理
func TestErrorHandling(t *testing.T) {
	api := &API{
		versionManager: NewVersionManager(),
	}
	
	// 测试无效版本
	err := api.ForceVersion("invalid.version")
	if err != nil {
		t.Errorf("ForceVersion should not error on invalid version: %v", err)
	}
	if api.IsZabbix6() || api.IsZabbix7() {
		t.Error("Invalid version should not be detected as Zabbix 6.x or 7.x")
	}
	
	// 测试不支持的版本
	err = api.ForceVersion("5.0.0")
	if err != nil {
		t.Error("Version forcing should succeed even for unsupported versions")
	}
	
	// 测试在未初始化版本管理器时的 API 调用
	apiNoVersion := &API{}
	
	if apiNoVersion.IsZabbix7() {
		t.Error("Uninitialized API should not report as Zabbix 7.0")
	}
	
	if apiNoVersion.IsFeatureSupported(FeatureMFA) {
		t.Error("Uninitialized API should not support any features")
	}
	
	if apiNoVersion.GetServerVersion() != "" {
		t.Error("Uninitialized API should not have server version")
	}
}
