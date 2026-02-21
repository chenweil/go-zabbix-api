package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tpretz/go-zabbix-api"
)

func main() {
	// Example: Multi-version Zabbix API usage
	
	// Create API configuration
	config := zabbix.Config{
		Url:         "http://your-zabbix-server/api_jsonrpc.php",
		TlsNoVerify: false,
		Log:         log.New(os.Stdout, "[ZABBIX] ", log.LstdFlags),
	}
	
	// Create API instance
	api := zabbix.NewAPI(config)
	
	// Login (this will automatically detect version and initialize appropriate adapters)
	auth, err := api.Login("admin", "zabbix")
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	
	fmt.Printf("Login successful, auth token: %s\n", auth)
	
	// Check detected version
	fmt.Printf("Detected Zabbix version: %s\n", api.GetServerVersion())
	fmt.Printf("Is Zabbix 7.0: %t\n", api.IsZabbix7())
	fmt.Printf("Is Zabbix 6.0: %t\n", api.IsZabbix6())
	
	// Check feature support
	fmt.Printf("History Push supported: %t\n", api.IsFeatureSupported(zabbix.FeatureHistoryPush))
	fmt.Printf("MFA supported: %t\n", api.IsFeatureSupported(zabbix.FeatureMFA))
	fmt.Printf("Proxy Group supported: %t\n", api.IsFeatureSupported(zabbix.FeatureProxyGroup))
	fmt.Printf("Browser Item supported: %t\n", api.IsFeatureSupported(zabbix.FeatureBrowserItem))
	
	// Example: Create an item with headers (works for both 6.0 and 7.0)
	item := zabbix.Item{
		HostID:  "10084", // Example host ID
		Key:     "web.page.get[example.com]",
		Name:    "Example.com page content",
		Type:    zabbix.WebItem,
		Delay:   "1m",
		ValueType: zabbix.Text,
		Url:     "http://example.com",
		Timeout: "10s",
	}
	
	// Set headers using the appropriate format
	if api.IsZabbix7() {
		// Zabbix 7.0 format: array of objects
		item.HeadersV7 = []zabbix.HeaderField{
			{Name: "User-Agent", Value: "Zabbix Monitoring"},
			{Name: "Accept", Value: "text/html"},
		}
	} else {
		// Zabbix 6.0 format: map
		item.HeadersV6 = zabbix.HttpHeaders{
			"User-Agent": "Zabbix Monitoring",
			"Accept":     "text/html",
		}
	}
	
	// Create item using adapter (automatically handles version differences)
	err = api.CreateItems(zabbix.Items{item})
	if err != nil {
		log.Printf("Failed to create item: %v", err)
	} else {
		fmt.Println("Item created successfully")
	}
	
	// Example: Create a host with proxy configuration
	host := zabbix.Host{
		Host:     "example-host",
		Name:     "Example Host",
		Status:   zabbix.Monitored,
		GroupIds: zabbix.HostGroupIDs{{GroupID: "15"}}, // Example group ID
		Interfaces: zabbix.HostInterfaces{
			{
				Type:        zabbix.Agent, // Zabbix agent
				Main:        "1",
				UseIP:       "1",
				DNS:         "",
				IP:          "192.168.1.100",
				Port:        "10050",
			},
		},
	}
	
	// Set proxy configuration based on version
	if api.IsZabbix7() {
		host.ProxyID = "10085"     // New field name in 7.0
		host.MonitoredBy = zabbix.MonitoredByProxy // Required field in 7.0
	} else {
		host.ProxyHostID = "10085" // Old field name in 6.0
	}
	
	// Create host using adapter (automatically handles version differences)
	hostAdapter := api.GetHostAdapter()
	if hostAdapter != nil {
		err = hostAdapter.CreateHosts(zabbix.Hosts{host})
	} else {
		err = api.HostsCreate(zabbix.Hosts{host})
	}
	if err != nil {
		log.Printf("Failed to create host: %v", err)
	} else {
		fmt.Println("Host created successfully")
	}
	
	// Example: Use History Push API (Zabbix 7.0+ only)
	if api.SupportsHistoryPush() {
		historyData := []zabbix.HistoryData{
			{
				Host:  "example-host",
				Key:   "web.page.get[example.com]",
				Value: "Page content here",
				Clock: 1609459200, // Unix timestamp
			},
		}
		fmt.Printf("History Push 支持，准备推送 %d 条数据（当前库未提供封装方法）\n", len(historyData))
		// 可手动调用: api.CallWithError(\"history.push\", historyData)
	} else {
		fmt.Println("History Push API not available in this Zabbix version")
	}
	
	// Example: Manual version forcing (for testing)
	fmt.Println("\n--- Testing manual version forcing ---")
	
	// Force Zabbix 6.0 mode
	err = api.ForceVersion("6.4.0")
	if err != nil {
		log.Printf("Failed to force version: %v", err)
	} else {
		fmt.Printf("Forced to version: %s\n", api.GetServerVersion())
		fmt.Printf("Headers 数组格式支持: %t\n", api.IsFeatureSupported(zabbix.FeatureHeadersArrayFormat))
	}
	
	// Force Zabbix 7.0 mode
	err = api.ForceVersion("7.0.0")
	if err != nil {
		log.Printf("Failed to force version: %v", err)
	} else {
		fmt.Printf("Forced to version: %s\n", api.GetServerVersion())
		fmt.Printf("Headers 数组格式支持: %t\n", api.IsFeatureSupported(zabbix.FeatureHeadersArrayFormat))
	}
	
	fmt.Println("Example completed successfully!")
}
