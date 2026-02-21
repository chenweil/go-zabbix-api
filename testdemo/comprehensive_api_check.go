package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	zabbix "github.com/tpretz/go-zabbix-api"
)

// TestResult represents a single test result
type TestResult struct {
	Name     string
	Status   string // PASS, FAIL, SKIP
	Message  string
	Duration time.Duration
	Details  map[string]interface{}
}

// TestSummary holds all test results
type TestSummary struct {
	Results   []TestResult
	StartTime time.Time
	EndTime   time.Time
}

func (ts *TestSummary) Add(result TestResult) {
	ts.Results = append(ts.Results, result)
}

func (ts *TestSummary) Print() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("                    ZABBIX API 完整测试报告")
	fmt.Println(strings.Repeat("=", 80))

	passed, failed, skipped := 0, 0, 0

	for _, r := range ts.Results {
		switch r.Status {
		case "PASS":
			passed++
		case "FAIL":
			failed++
		case "SKIP":
			skipped++
		}

		icon := "✅"
		if r.Status == "FAIL" {
			icon = "❌"
		} else if r.Status == "SKIP" {
			icon = "⏭️"
		}

		fmt.Printf("\n%s [%s] %s (%.2fs)\n", icon, r.Status, r.Name, r.Duration.Seconds())
		if r.Message != "" {
			fmt.Printf("   %s\n", r.Message)
		}
		if len(r.Details) > 0 {
			for k, v := range r.Details {
				fmt.Printf("   - %s: %v\n", k, v)
			}
		}
	}

	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Printf("总测试数: %d | 通过: %d | 失败: %d | 跳过: %d\n",
		len(ts.Results), passed, failed, skipped)
	fmt.Printf("测试耗时: %.2f 秒\n", ts.EndTime.Sub(ts.StartTime).Seconds())
	fmt.Println(strings.Repeat("=", 80))
}

func main() {
	summary := &TestSummary{
		Results:   []TestResult{},
		StartTime: time.Now(),
	}

	// Get environment variables
	zabbixURL := os.Getenv("TEST_ZABBIX_URL")
	zabbixUser := os.Getenv("TEST_ZABBIX_USER")
	zabbixPass := os.Getenv("TEST_ZABBIX_PASSWORD")

	if zabbixURL == "" || zabbixUser == "" || zabbixPass == "" {
		fmt.Println("错误: 请设置以下环境变量:")
		fmt.Println("  - TEST_ZABBIX_URL")
		fmt.Println("  - TEST_ZABBIX_USER")
		fmt.Println("  - TEST_ZABBIX_PASSWORD")
		os.Exit(1)
	}

	// Create API client
	api := zabbix.NewAPI(zabbix.Config{
		Url: zabbixURL,
	})

	// ============ 连接与认证测试 ============
	fmt.Println("\n📡 连接与认证测试")
	fmt.Println(strings.Repeat("-", 60))

	// Test 1: Get Zabbix Version
	testAPIVersion(api, summary)

	// Test 2: Login
	testLogin(api, zabbixUser, zabbixPass, summary)

	// Test 3: Get Server Version
	testServerVersion(api, summary)

	// ============ 主机组 API 测试 ============
	fmt.Println("\n📁 主机组 (HostGroup) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testHostGroupAPI(api, summary)

	// ============ 主机 API 测试 ============
	fmt.Println("\n🖥️ 主机 (Host) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testHostAPI(api, summary)

	// ============ 模板 API 测试 ============
	fmt.Println("\n📋 模板 (Template) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testTemplateAPI(api, summary)

	// ============ 监控项 API 测试 ============
	fmt.Println("\n📊 监控项 (Item) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testItemAPI(api, summary)

	// ============ 触发器 API 测试 ============
	fmt.Println("\n🚨 触发器 (Trigger) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testTriggerAPI(api, summary)

	// ============ 图表 API 测试 ============
	fmt.Println("\n📈 图表 (Graph) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testGraphAPI(api, summary)

	// ============ 用户宏 API 测试 ============
	fmt.Println("\n🔧 用户宏 (Macro) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testMacroAPI(api, summary)

	// ============ 代理 API 测试 ============
	fmt.Println("\n🔌 代理 (Proxy) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testProxyAPI(api, summary)

	// ============ 用户 API 测试 ============
	fmt.Println("\n👤 用户 (User) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testUserAPI(api, summary)

	// ============ 媒介类型 API 测试 ============
	fmt.Println("\n📧 媒介类型 (MediaType) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testMediaTypeAPI(api, summary)

	// ============ 告警 API 测试 ============
	fmt.Println("\n🔔 告警 (Alert) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testAlertAPI(api, summary)

	// ============ LLD API 测试 ============
	fmt.Println("\n🔍 低级别发现规则 (LLD) API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testLLDAPI(api, summary)

	// ============ Item Prototype API 测试 ============
	fmt.Println("\n📦 Item Prototype API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testItemPrototypeAPI(api, summary)

	// ============ Host Prototype API 测试 ============
	fmt.Println("\n🖥️ Host Prototype API 测试")
	fmt.Println(strings.Repeat("-", 60))
	testHostPrototypeAPI(api, summary)

	// ============ 多版本支持测试 ============
	fmt.Println("\n🔄 多版本支持测试")
	fmt.Println(strings.Repeat("-", 60))
	testMultiVersionSupport(api, summary)

	// ============ Zabbix 7.0+ 特性测试 ============
	if api.IsZabbix7() {
		fmt.Println("\n✨ Zabbix 7.0+ 特性测试")
		fmt.Println(strings.Repeat("-", 60))
		testZabbix7Features(api, summary)
	}

	// Test: Logout
	testLogout(api, summary)

	summary.EndTime = time.Now()
	summary.Print()
}

// ============ 具体测试函数 ============

func testAPIVersion(api *zabbix.API, summary *TestSummary) {
	start := time.Now()
	version, err := api.Version()
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "获取 Zabbix API 版本",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "获取 Zabbix API 版本",
			Status:   "PASS",
			Message:  fmt.Sprintf("版本: %s", version),
			Duration: duration,
		})
	}
}

func testLogin(api *zabbix.API, user, pass string, summary *TestSummary) {
	start := time.Now()
	_, err := api.Login(user, pass)
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "用户登录",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "用户登录",
			Status:   "PASS",
			Message:  "登录成功",
			Duration: duration,
		})
	}
}

func testServerVersion(api *zabbix.API, summary *TestSummary) {
	start := time.Now()
	version, err := api.DetectVersion()
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "检测服务器版本",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		details := map[string]interface{}{
			"版本":        version,
			"是 Zabbix 7": api.IsZabbix7(),
			"是 Zabbix 6": api.IsZabbix6(),
		}

		// Get supported features
		features := api.GetSupportedFeatures()
		enabledFeatures := []string{}
		for feature, supported := range features {
			if supported {
				enabledFeatures = append(enabledFeatures, feature)
			}
		}
		details["支持特性"] = enabledFeatures

		summary.Add(TestResult{
			Name:     "检测服务器版本",
			Status:   "PASS",
			Message:  fmt.Sprintf("版本: %s", version),
			Duration: duration,
			Details:  details,
		})
	}
}

func testHostGroupAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all host groups
	start := time.Now()
	groups, err := api.HostGroupsGet(zabbix.Params{})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "HostGroup.Get - 获取所有主机组",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "HostGroup.Get - 获取所有主机组",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个主机组", len(groups)),
			Duration: duration,
			Details: map[string]interface{}{
				"主机组数量": len(groups),
			},
		})
	}

	// Test 2: Get host group by ID
	if len(groups) > 0 {
		start = time.Now()
		group, err := api.HostGroupGetByID(groups[0].GroupID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "HostGroup.GetByID - 通过 ID 获取主机组",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "HostGroup.GetByID - 通过 ID 获取主机组",
				Status:   "PASS",
				Message:  fmt.Sprintf("主机组: %s (ID: %s)", group.Name, group.GroupID),
				Duration: duration,
			})
		}
	}

	// Test 3: Create a test host group
	start = time.Now()
	testGroupName := fmt.Sprintf("TestGroup-%d", time.Now().Unix())
	newGroups := zabbix.HostGroups{
		{Name: testGroupName},
	}
	err = api.HostGroupsCreate(newGroups)
	duration = time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "HostGroup.Create - 创建主机组",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "HostGroup.Create - 创建主机组",
			Status:   "PASS",
			Message:  fmt.Sprintf("创建主机组: %s (ID: %s)", testGroupName, newGroups[0].GroupID),
			Duration: duration,
		})

		// Test 4: Delete the test host group
		start = time.Now()
		err = api.HostGroupsDelete(newGroups)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "HostGroup.Delete - 删除主机组",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "HostGroup.Delete - 删除主机组",
				Status:   "PASS",
				Message:  "主机组删除成功",
				Duration: duration,
			})
		}
	}
}

func testHostAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all hosts
	start := time.Now()
	hosts, err := api.HostsGet(zabbix.Params{})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Host.Get - 获取所有主机",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "Host.Get - 获取所有主机",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个主机", len(hosts)),
			Duration: duration,
			Details: map[string]interface{}{
				"主机数量": len(hosts),
			},
		})
	}

	// Test 2: Get host by ID
	if len(hosts) > 0 {
		start = time.Now()
		host, err := api.HostGetByID(hosts[0].HostID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Host.GetByID - 通过 ID 获取主机",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			details := map[string]interface{}{
				"主机名": host.Host,
				"显示名": host.Name,
				"状态":   host.Status,
				"可用性": host.Available,
			}
			if host.UUID != "" {
				details["UUID"] = host.UUID
			}

			summary.Add(TestResult{
				Name:     "Host.GetByID - 通过 ID 获取主机",
				Status:   "PASS",
				Message:  fmt.Sprintf("主机: %s (ID: %s)", host.Name, host.HostID),
				Duration: duration,
				Details:  details,
			})
		}
	}

	// Test 3: Get hosts by host group
	groups, _ := api.HostGroupsGet(zabbix.Params{"limit": 1})
	if len(groups) > 0 {
		start = time.Now()
		hostsByGroup, err := api.HostsGetByHostGroupIds([]string{groups[0].GroupID})
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Host.GetByHostGroupIds - 通过主机组获取主机",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "Host.GetByHostGroupIds - 通过主机组获取主机",
				Status:   "PASS",
				Message:  fmt.Sprintf("主机组 %s 中有 %d 个主机", groups[0].Name, len(hostsByGroup)),
				Duration: duration,
			})
		}
	}
}

func testTemplateAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all templates
	start := time.Now()
	templates, err := api.TemplatesGet(zabbix.Params{})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Template.Get - 获取所有模板",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "Template.Get - 获取所有模板",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个模板", len(templates)),
			Duration: duration,
			Details: map[string]interface{}{
				"模板数量": len(templates),
			},
		})
	}

	// Test 2: Get template by ID
	if len(templates) > 0 {
		start = time.Now()
		template, err := api.TemplateGetByID(templates[0].TemplateID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Template.GetByID - 通过 ID 获取模板",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			details := map[string]interface{}{
				"模板名": template.Host,
				"描述":  template.Description,
			}
			if template.UUID != "" {
				details["UUID"] = template.UUID
			}

			summary.Add(TestResult{
				Name:     "Template.GetByID - 通过 ID 获取模板",
				Status:   "PASS",
				Message:  fmt.Sprintf("模板: %s (ID: %s)", template.Host, template.TemplateID),
				Duration: duration,
				Details:  details,
			})
		}
	}
}

func testItemAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all items
	start := time.Now()
	items, err := api.ItemsGet(zabbix.Params{"limit": 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Item.Get - 获取所有监控项",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "Item.Get - 获取所有监控项",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个监控项", len(items)),
			Duration: duration,
			Details: map[string]interface{}{
				"监控项数量": len(items),
			},
		})
	}

	// Test 2: Get item by ID
	if len(items) > 0 {
		start = time.Now()
		item, err := api.ItemGetByID(items[0].ItemID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Item.GetByID - 通过 ID 获取监控项",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			itemTypes := map[zabbix.ItemType]string{
				zabbix.ZabbixAgent:       "Zabbix Agent",
				zabbix.SNMPv1Agent:       "SNMPv1 Agent",
				zabbix.ZabbixTrapper:     "Zabbix Trapper",
				zabbix.SimpleCheck:       "Simple Check",
				zabbix.ZabbixInternal:    "Zabbix Internal",
				zabbix.ZabbixAgentActive: "Zabbix Agent (Active)",
				zabbix.ExternalCheck:     "External Check",
				zabbix.DatabaseMonitor:   "Database Monitor",
				zabbix.HTTPAgent:         "HTTP Agent",
				zabbix.Browser:           "Browser (Zabbix 7.0+)",
			}

			itemType := itemTypes[item.Type]
			if itemType == "" {
				itemType = fmt.Sprintf("Type %d", item.Type)
			}

			summary.Add(TestResult{
				Name:     "Item.GetByID - 通过 ID 获取监控项",
				Status:   "PASS",
				Message:  fmt.Sprintf("监控项: %s (类型: %s)", item.Name, itemType),
				Duration: duration,
				Details: map[string]interface{}{
					"键值":     item.Key,
					"类型":     itemType,
					"更新间隔": item.Delay,
				},
			})
		}
	}

	// Test 3: Get items by host
	hosts, _ := api.HostsGet(zabbix.Params{"limit": 1})
	if len(hosts) > 0 {
		start = time.Now()
		itemsByHost, err := api.ItemsGetByHostIds([]string{hosts[0].HostID})
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Item.GetByHostIds - 通过主机获取监控项",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "Item.GetByHostIds - 通过主机获取监控项",
				Status:   "PASS",
				Message:  fmt.Sprintf("主机 %s 有 %d 个监控项", hosts[0].Name, len(itemsByHost)),
				Duration: duration,
			})
		}
	}
}

func testTriggerAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all triggers
	start := time.Now()
	triggers, err := api.TriggersGet(zabbix.Params{"limit": 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Trigger.Get - 获取所有触发器",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "Trigger.Get - 获取所有触发器",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个触发器", len(triggers)),
			Duration: duration,
			Details: map[string]interface{}{
				"触发器数量": len(triggers),
			},
		})
	}

	// Test 2: Get trigger by ID
	if len(triggers) > 0 {
		start = time.Now()
		trigger, err := api.TriggerGetByID(triggers[0].TriggerID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Trigger.GetByID - 通过 ID 获取触发器",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			severities := map[zabbix.SeverityType]string{
				zabbix.NotClassified: "未分类",
				zabbix.Information:   "信息",
				zabbix.Warning:       "警告",
				zabbix.Average:       "一般",
				zabbix.High:          "严重",
				zabbix.Critical:      "灾难",
			}

			summary.Add(TestResult{
				Name:     "Trigger.GetByID - 通过 ID 获取触发器",
				Status:   "PASS",
				Message:  fmt.Sprintf("触发器: %s (严重性: %s)", trigger.Description, severities[trigger.Priority]),
				Duration: duration,
				Details: map[string]interface{}{
					"表达式": trigger.Expression,
					"严重性": severities[trigger.Priority],
					"状态":   trigger.Status,
				},
			})
		}
	}
}

func testGraphAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all graphs
	start := time.Now()
	graphs, err := api.GraphsGet(zabbix.Params{"limit": 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Graph.Get - 获取所有图表",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "Graph.Get - 获取所有图表",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个图表", len(graphs)),
			Duration: duration,
			Details: map[string]interface{}{
				"图表数量": len(graphs),
			},
		})
	}

	// Test 2: Get graph by ID
	if len(graphs) > 0 {
		start = time.Now()
		graph, err := api.GraphGetByID(graphs[0].GraphID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Graph.GetByID - 通过 ID 获取图表",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			graphTypes := map[zabbix.GraphType]string{
				zabbix.GraphNormal:   "普通",
				zabbix.GraphStacked:  "堆叠",
				zabbix.GraphPie:      "饼图",
				zabbix.GraphExploded: "爆炸饼图",
			}

			summary.Add(TestResult{
				Name:     "Graph.GetByID - 通过 ID 获取图表",
				Status:   "PASS",
				Message:  fmt.Sprintf("图表: %s (类型: %s)", graph.Name, graphTypes[graph.Type]),
				Duration: duration,
				Details: map[string]interface{}{
					"宽度":     graph.Width,
					"高度":     graph.Height,
					"类型":     graphTypes[graph.Type],
					"图表项数": len(graph.GraphItems),
				},
			})
		}
	}
}

func testMacroAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all macros
	start := time.Now()
	macros, err := api.MacrosGet(zabbix.Params{"limit": 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Macro.Get - 获取所有用户宏",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "Macro.Get - 获取所有用户宏",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个用户宏", len(macros)),
			Duration: duration,
			Details: map[string]interface{}{
				"宏数量": len(macros),
			},
		})
	}

	// Test 2: Create a test macro (need a host first)
	hosts, _ := api.HostsGet(zabbix.Params{"limit": 1})
	if len(hosts) > 0 {
		hostID := hosts[0].HostID
		macroName := fmt.Sprintf("{$TEST_MACRO_%d}", time.Now().Unix())

		// Create macro
		start = time.Now()
		newMacros := zabbix.Macros{
			{
				HostID:    hostID,
				MacroName: macroName,
				Value:     "test_value_123",
			},
		}
		err = api.MacrosCreate(newMacros)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "Macro.Create - 创建用户宏",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "Macro.Create - 创建用户宏",
				Status:   "PASS",
				Message:  fmt.Sprintf("创建宏: %s = test_value_123", macroName),
				Duration: duration,
			})

			// Test 3: Get macro by ID (using the created macro's ID)
			if newMacros[0].MacroID != "" {
				start = time.Now()
				macro, err := api.MacroGetByID(newMacros[0].MacroID)
				duration = time.Since(start)

				if err != nil {
					summary.Add(TestResult{
						Name:     "Macro.GetByID - 通过 ID 获取用户宏",
						Status:   "FAIL",
						Message:  err.Error(),
						Duration: duration,
					})
				} else {
					summary.Add(TestResult{
						Name:     "Macro.GetByID - 通过 ID 获取用户宏",
						Status:   "PASS",
						Message:  fmt.Sprintf("宏: %s = %s", macro.MacroName, macro.Value),
						Duration: duration,
					})
				}

				// Test 4: Delete the macro
				start = time.Now()
				err = api.MacrosDelete(newMacros)
				duration = time.Since(start)

				if err != nil {
					summary.Add(TestResult{
						Name:     "Macro.Delete - 删除用户宏",
						Status:   "FAIL",
						Message:  err.Error(),
						Duration: duration,
					})
				} else {
					summary.Add(TestResult{
						Name:     "Macro.Delete - 删除用户宏",
						Status:   "PASS",
						Message:  "宏删除成功",
						Duration: duration,
					})
				}
			}
		}
	} else {
		summary.Add(TestResult{
			Name:     "Macro.Create - 创建用户宏",
			Status:   "SKIP",
			Message:  "没有可用的主机来创建宏",
			Duration: 0,
		})
	}
}

func testProxyAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all proxies
	start := time.Now()
	proxies, err := api.ProxiesGet(zabbix.Params{})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Proxy.Get - 获取所有代理",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "Proxy.Get - 获取所有代理",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个代理", len(proxies)),
			Duration: duration,
			Details: map[string]interface{}{
				"代理数量": len(proxies),
			},
		})
	}
}

func testUserAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all users
	start := time.Now()
	users, err := api.UsersGet(zabbix.UserGetOptions{Limit: 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "User.Get - 获取所有用户",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "User.Get - 获取所有用户",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个用户", len(users)),
			Duration: duration,
			Details: map[string]interface{}{
				"用户数量": len(users),
			},
		})
	}

	// Test 2: Get user by ID
	if len(users) > 0 {
		start = time.Now()
		usersById, err := api.UsersGetById([]string{users[0].UserID})
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "User.GetById - 通过 ID 获取用户",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else if len(usersById) > 0 {
			summary.Add(TestResult{
				Name:     "User.GetById - 通过 ID 获取用户",
				Status:   "PASS",
				Message:  fmt.Sprintf("用户: %s (%s %s)", usersById[0].Username, usersById[0].Name, usersById[0].Surname),
				Duration: duration,
			})
		}
	}
}

func testMediaTypeAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all media types
	start := time.Now()
	mediaTypes, err := api.MediaTypesGet(zabbix.MediaTypeGetOptions{})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "MediaType.Get - 获取所有媒介类型",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		typeNames := map[string]string{
			"0": "Email",
			"1": "Script",
			"2": "SMS",
			"3": "Jabber",
			"4": "Ez Texting",
			"5": "Webhook",
		}

		typeCount := make(map[string]int)
		for _, mt := range mediaTypes {
			typeName := typeNames[mt.Type]
			if typeName == "" {
				typeName = fmt.Sprintf("Type %s", mt.Type)
			}
			typeCount[typeName]++
		}

		// Convert typeCount to map[string]interface{}
		typeCountInterface := make(map[string]interface{})
		for k, v := range typeCount {
			typeCountInterface[k] = v
		}

		summary.Add(TestResult{
			Name:     "MediaType.Get - 获取所有媒介类型",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个媒介类型", len(mediaTypes)),
			Duration: duration,
			Details:  typeCountInterface,
		})
	}
}

func testAlertAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get recent alerts
	start := time.Now()
	alerts, err := api.AlertsGetRecent()
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "Alert.GetRecent - 获取最近告警",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		statusCount := make(map[string]int)
		for _, alert := range alerts {
			statusCount[alert.Status]++
		}

		summary.Add(TestResult{
			Name:     "Alert.GetRecent - 获取最近告警",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 条告警记录", len(alerts)),
			Duration: duration,
			Details: map[string]interface{}{
				"告警总数": len(alerts),
				"状态分布": statusCount,
			},
		})
	}
}

func testLLDAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all LLD rules
	start := time.Now()
	llds, err := api.LLDsGet(zabbix.Params{"limit": 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "LLD.Get - 获取所有发现规则",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "LLD.Get - 获取所有发现规则",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个发现规则", len(llds)),
			Duration: duration,
			Details: map[string]interface{}{
				"发现规则数量": len(llds),
			},
		})
	}

	// Test 2: Get LLD by ID
	if len(llds) > 0 {
		start = time.Now()
		lld, err := api.LLDGetByID(llds[0].ItemID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "LLD.GetByID - 通过 ID 获取发现规则",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "LLD.GetByID - 通过 ID 获取发现规则",
				Status:   "PASS",
				Message:  fmt.Sprintf("发现规则: %s", lld.Name),
				Duration: duration,
				Details: map[string]interface{}{
					"键值": lld.Key,
					"类型": lld.Type,
				},
			})
		}
	}
}

func testItemPrototypeAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all item prototypes
	start := time.Now()
	items, err := api.ItemPrototypesGet(zabbix.Params{"limit": 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "ItemPrototype.Get - 获取所有监控项原型",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "ItemPrototype.Get - 获取所有监控项原型",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个监控项原型", len(items)),
			Duration: duration,
			Details: map[string]interface{}{
				"监控项原型数量": len(items),
			},
		})
	}

	// Test 2: Get item prototype by ID
	if len(items) > 0 {
		start = time.Now()
		item, err := api.ItemPrototypeGetByID(items[0].ItemID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "ItemPrototype.GetByID - 通过 ID 获取监控项原型",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "ItemPrototype.GetByID - 通过 ID 获取监控项原型",
				Status:   "PASS",
				Message:  fmt.Sprintf("监控项原型: %s", item.Name),
				Duration: duration,
				Details: map[string]interface{}{
					"键值": item.Key,
					"类型": item.Type,
				},
			})
		}
	}
}

func testHostPrototypeAPI(api *zabbix.API, summary *TestSummary) {
	// Test 1: Get all host prototypes
	start := time.Now()
	hosts, err := api.HostPrototypesGet(zabbix.Params{"limit": 50})
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "HostPrototype.Get - 获取所有主机原型",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "HostPrototype.Get - 获取所有主机原型",
			Status:   "PASS",
			Message:  fmt.Sprintf("找到 %d 个主机原型", len(hosts)),
			Duration: duration,
			Details: map[string]interface{}{
				"主机原型数量": len(hosts),
			},
		})
	}

	// Test 2: Get host prototype by ID
	if len(hosts) > 0 {
		start = time.Now()
		host, err := api.HostPrototypeGetByID(hosts[0].HostID)
		duration = time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "HostPrototype.GetByID - 通过 ID 获取主机原型",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "HostPrototype.GetByID - 通过 ID 获取主机原型",
				Status:   "PASS",
				Message:  fmt.Sprintf("主机原型: %s", host.Name),
				Duration: duration,
				Details: map[string]interface{}{
					"主机名": host.Host,
				},
			})
		}
	}
}

func testMultiVersionSupport(api *zabbix.API, summary *TestSummary) {
	// Test 1: Check version adapters
	start := time.Now()
	itemAdapter := api.GetItemAdapter()
	hostAdapter := api.GetHostAdapter()
	duration := time.Since(start)

	if itemAdapter == nil || hostAdapter == nil {
		summary.Add(TestResult{
			Name:     "多版本适配器",
			Status:   "FAIL",
			Message:  "适配器未初始化",
			Duration: duration,
		})
	} else {
		version := "6.x"
		if api.IsZabbix7() {
			version = "7.x"
		}

		summary.Add(TestResult{
			Name:     "多版本适配器",
			Status:   "PASS",
			Message:  fmt.Sprintf("已初始化 %s 适配器", version),
			Duration: duration,
			Details: map[string]interface{}{
				"ItemAdapter": itemAdapter != nil,
				"HostAdapter": hostAdapter != nil,
			},
		})
	}

	// Test 2: Check supported features
	start = time.Now()
	features := api.GetSupportedFeatures()
	duration = time.Since(start)

	enabledFeatures := []string{}
	for feature, supported := range features {
		if supported {
			enabledFeatures = append(enabledFeatures, feature)
		}
	}

	summary.Add(TestResult{
		Name:     "支持的特性列表",
		Status:   "PASS",
		Message:  fmt.Sprintf("检测到 %d 个支持特性", len(enabledFeatures)),
		Duration: duration,
		Details: map[string]interface{}{
			"支持特性": enabledFeatures,
		},
	})
}

func testZabbix7Features(api *zabbix.API, summary *TestSummary) {
	// Test 1: Check MFA support
	if api.SupportsMFA() {
		start := time.Now()
		mfas, err := api.MFAGet(zabbix.Params{"limit": 10})
		duration := time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "MFA.Get - 获取 MFA 配置",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "MFA.Get - 获取 MFA 配置",
				Status:   "PASS",
				Message:  fmt.Sprintf("找到 %d 个 MFA 配置", len(mfas)),
				Duration: duration,
			})
		}
	}

	// Test 2: Check Proxy Group support
	if api.SupportsProxyGroup() {
		start := time.Now()
		proxyGroups, err := api.ProxyGroupGet(zabbix.Params{})
		duration := time.Since(start)

		if err != nil {
			summary.Add(TestResult{
				Name:     "ProxyGroup.Get - 获取代理组",
				Status:   "FAIL",
				Message:  err.Error(),
				Duration: duration,
			})
		} else {
			summary.Add(TestResult{
				Name:     "ProxyGroup.Get - 获取代理组",
				Status:   "PASS",
				Message:  fmt.Sprintf("找到 %d 个代理组", len(proxyGroups)),
				Duration: duration,
			})
		}
	}
}

func testLogout(api *zabbix.API, summary *TestSummary) {
	start := time.Now()
	err := api.Logout()
	duration := time.Since(start)

	if err != nil {
		summary.Add(TestResult{
			Name:     "用户登出",
			Status:   "FAIL",
			Message:  err.Error(),
			Duration: duration,
		})
	} else {
		summary.Add(TestResult{
			Name:     "用户登出",
			Status:   "PASS",
			Message:  "登出成功",
			Duration: duration,
		})
	}
}
