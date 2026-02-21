package main

import (
	"fmt"
	"log"
	"os"

	zabbix "github.com/tpretz/go-zabbix-api"
)

func main() {
	// 从环境变量或直接使用提供的凭据
	url := getEnv("ZABBIX_URL", "https://web.mserver.orb.local/api_jsonrpc.php")
	user := getEnv("ZABBIX_USER", "Admin")
	password := getEnv("ZABBIX_PASSWORD", "privatepassword")

	fmt.Println("=== Zabbix API 实际服务器测试 ===")
	fmt.Printf("服务器地址: %s\n", url)
	fmt.Printf("用户名: %s\n", user)
	fmt.Println()

	// 创建 API 实例
	api := zabbix.NewAPI(url)
	api.SetClient(nil) // 使用默认 HTTP 客户端

	// 测试 1: 获取版本信息
	fmt.Println("【测试 1】获取 Zabbix 版本")
	version, err := api.Version()
	if err != nil {
		log.Fatalf("❌ 获取版本失败: %v", err)
	}
	fmt.Printf("✅ Zabbix 版本: %s\n\n", version)

	// 测试 2: 登录认证
	fmt.Println("【测试 2】用户登录认证")
	auth, err := api.Login(user, password)
	if err != nil {
		log.Fatalf("❌ 登录失败: %v", err)
	}
	fmt.Printf("✅ 登录成功, Auth Token: %s...\n\n", auth[:20])

	// 测试 3: 获取主机组列表
	fmt.Println("【测试 3】获取主机组列表")
	params := make(map[string]interface{})
	params["output"] = "extend"
	params["limit"] = 5

	groups, err := api.HostGroupsGet(params)
	if err != nil {
		log.Printf("❌ 获取主机组失败: %v", err)
	} else {
		fmt.Printf("✅ 找到 %d 个主机组:\n", len(groups))
		for i, group := range groups {
			fmt.Printf("  %d. [%s] %s\n", i+1, group.GroupID, group.Name)
		}
		fmt.Println()
	}

	// 测试 4: 获取主机列表
	fmt.Println("【测试 4】获取主机列表")
	hostParams := make(map[string]interface{})
	hostParams["output"] = []string{"hostid", "host", "name", "status"}
	hostParams["limit"] = 5

	hosts, err := api.HostsGet(hostParams)
	if err != nil {
		log.Printf("❌ 获取主机失败: %v", err)
	} else {
		fmt.Printf("✅ 找到 %d 个主机:\n", len(hosts))
		for i, host := range hosts {
			status := "启用"
			if host.Status != 0 {
				status = "禁用"
			}
			fmt.Printf("  %d. [%s] %s (%s) - %s\n", i+1, host.HostID, host.Host, host.Name, status)
		}
		fmt.Println()
	}

	// 测试 5: 获取模板列表
	fmt.Println("【测试 5】获取模板列表")
	templateParams := make(map[string]interface{})
	templateParams["output"] = []string{"templateid", "host", "name"}
	templateParams["limit"] = 5

	templates, err := api.TemplatesGet(templateParams)
	if err != nil {
		log.Printf("❌ 获取模板失败: %v", err)
	} else {
		fmt.Printf("✅ 找到 %d 个模板:\n", len(templates))
		for i, template := range templates {
			fmt.Printf("  %d. [%s] %s (%s)\n", i+1, template.TemplateID, template.Host, template.Name)
		}
		fmt.Println()
	}

	// 测试 6: 获取触发器列表
	fmt.Println("【测试 6】获取触发器列表")
	triggerParams := make(map[string]interface{})
	triggerParams["output"] = []string{"triggerid", "description", "priority"}
	triggerParams["limit"] = 5
	triggerParams["sortfield"] = "priority"
	triggerParams["sortorder"] = "DESC"

	triggers, err := api.TriggersGet(triggerParams)
	if err != nil {
		log.Printf("❌ 获取触发器失败: %v", err)
	} else {
		fmt.Printf("✅ 找到 %d 个触发器:\n", len(triggers))
		for i, trigger := range triggers {
			priority := getPriorityName(trigger.Priority)
			fmt.Printf("  %d. [%s] %s - %s\n", i+1, trigger.TriggerID, trigger.Description, priority)
		}
		fmt.Println()
	}

	// 测试 7: 获取监控项列表
	fmt.Println("【测试 7】获取监控项列表")
	itemParams := make(map[string]interface{})
	itemParams["output"] = []string{"itemid", "name", "key_", "type"}
	itemParams["limit"] = 5

	items, err := api.ItemsGet(itemParams)
	if err != nil {
		log.Printf("❌ 获取监控项失败: %v", err)
	} else {
		fmt.Printf("✅ 找到 %d 个监控项:\n", len(items))
		for i, item := range items {
			fmt.Printf("  %d. [%s] %s (key: %s)\n", i+1, item.ItemID, item.Name, item.Key)
		}
		fmt.Println()
	}

	// 测试 8: 登出
	fmt.Println("【测试 8】用户登出")
	_, err = api.Logout()
	if err != nil {
		log.Printf("❌ 登出失败: %v", err)
	} else {
		fmt.Println("✅ 登出成功")
	}

	fmt.Println("\n=== 测试完成 ===")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getPriorityName(priority int) string {
	priorities := map[int]string{
		0: "未分类",
		1: "信息",
		2: "警告",
		3: "一般严重",
		4: "严重",
		5: "灾难",
	}
	if name, ok := priorities[priority]; ok {
		return name
	}
	return fmt.Sprintf("未知(%d)", priority)
}
