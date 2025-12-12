# Zabbix 6.0/7.0 多版本支持

这个版本的 go-zabbix-api 现在支持 Zabbix 6.0 和 7.0 两个版本，提供自动版本检测和自适应功能。

## 🚀 新特性

### 自动版本检测
- 连接时自动检测 Zabbix 服务器版本
- 根据版本自动选择合适的适配器
- 无需手动配置版本信息

### 适配器模式
- 为不同版本提供专门的适配器
- 自动处理版本间的数据格式差异
- 保持 API 接口的一致性

### 向后兼容
- 完全兼容 Zabbix 6.0 API
- 现有代码无需修改即可使用
- 渐进式升级支持

## 📋 主要变更

### 1. Item Headers 格式变更
```go
// Zabbix 6.0 格式 (对象)
item.HeadersV6 = zabbix.HttpHeaders{
    "User-Agent": "Zabbix Monitoring",
    "Accept": "text/html",
}

// Zabbix 7.0 格式 (数组)
item.HeadersV7 = []zabbix.HeaderField{
    {Name: "User-Agent", Value: "Zabbix Monitoring"},
    {Name: "Accept", Value: "text/html"},
}
```

### 2. Host Proxy 字段变更
```go
// Zabbix 6.0 格式
host.ProxyHostID = "10085"

// Zabbix 7.0 格式
host.ProxyID = "10085"
host.MonitoredBy = zabbix.MonitoredByProxy // 必需字段
```

### 3. 新增 Item 类型
```go
// Zabbix 7.0+ 新增
item.Type = zabbix.Browser // 浏览器监控
```

## 🔧 使用方法

### 基本用法
```go
package main

import (
    "fmt"
    "log"
    "github.com/tpretz/go-zabbix-api"
)

func main() {
    // 创建 API 配置
    config := zabbix.Config{
        Url:         "http://your-zabbix-server/api_jsonrpc.php",
        TlsNoVerify: false,
    }
    
    // 创建 API 实例
    api := zabbix.NewAPI(config)
    
    // 登录 (自动检测版本)
    auth, err := api.Login("admin", "zabbix")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("登录成功，检测到版本: %s\n", api.GetServerVersion())
    
    // 使用统一的 API 接口
    items, err := api.GetItems(zabbix.Params{"output": "extend"})
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("获取到 %d 个监控项\n", len(items))
}
```

### 版本检测
```go
// 检查版本
fmt.Printf("是 Zabbix 7.0: %t\n", api.IsZabbix7())
fmt.Printf("是 Zabbix 6.0: %t\n", api.IsZabbix6())

// 检查特性支持
fmt.Printf("支持 History Push: %t\n", api.IsFeatureSupported(zabbix.FeatureHistoryPush))
fmt.Printf("支持 MFA: %t\n", api.IsFeatureSupported(zabbix.FeatureMFA))
fmt.Printf("支持代理组: %t\n", api.IsFeatureSupported(zabbix.FeatureProxyGroup))
```

### 创建监控项 (多版本兼容)
```go
item := zabbix.Item{
    HostID:    "10084",
    Key:       "web.page.get[example.com]",
    Name:      "Example.com 页面内容",
    Type:      zabbix.WebItem,
    Delay:     "1m",
    ValueType: zabbix.Text,
    Url:       "http://example.com",
    Timeout:   "10s",
}

// 自动适配版本格式
if api.IsZabbix7() {
    item.HeadersV7 = []zabbix.HeaderField{
        {Name: "User-Agent", Value: "Zabbix Monitoring"},
    }
} else {
    item.HeadersV6 = zabbix.HttpHeaders{
        "User-Agent": "Zabbix Monitoring",
    }
}

// 创建监控项 (自动处理版本差异)
err := api.CreateItems(zabbix.Items{item})
```

### 创建主机 (多版本兼容)
```go
host := zabbix.Host{
    Host:     "example-host",
    Name:     "Example Host",
    Status:   zabbix.Monitored,
    GroupIds: zabbix.HostGroupIDs{{GroupID: "15"}},
}

// 自动适配代理配置
if api.IsZabbix7() {
    host.ProxyID = "10085"
    host.MonitoredBy = zabbix.MonitoredByProxy
} else {
    host.ProxyHostID = "10085"
}

// 创建主机
err := api.CreateHosts(zabbix.Hosts{host})
```

### History Push API (Zabbix 7.0+)
```go
if api.IsFeatureSupported(zabbix.FeatureHistoryPush) {
    data := []zabbix.HistoryData{
        {
            Host:  "example-host",
            Key:   "web.page.get[example.com]",
            Value: "页面内容",
            Clock: 1609459200,
        },
    }
    
    err := api.HistoryPush(data)
    if err != nil {
        log.Printf("推送历史数据失败: %v", err)
    }
}
```

### 手动版本控制
```go
// 强制指定版本 (用于测试)
err := api.ForceVersion("7.0.0")
if err != nil {
    log.Fatal(err)
}

// 不自动检测版本的登录
auth, err := api.LoginWithoutVersionInit("admin", "zabbix")
if err != nil {
    log.Fatal(err)
}

// 手动初始化版本支持
err = api.InitializeVersionSupport()
if err != nil {
    log.Fatal(err)
}
```

## 🔍 特性常量

```go
const (
    FeatureHistoryPush   = "history.push"   // History Push API
    FeatureMFA          = "mfa"            // 多因子认证
    FeatureProxyGroup   = "proxygroup"     // 代理组
    FeatureBrowserItem  = "browser_item"   // 浏览器监控项
    FeatureHeadersV7    = "headers_v7"     // 7.0 格式 headers
    FeatureProxyID      = "proxyid"        // 代理 ID 字段
    FeatureMonitoredBy  = "monitored_by"   // 监控方式字段
)
```

## 🏗️ 架构组件

### VersionManager
负责版本检测和特性支持管理：
- `DetectVersion()` - 检测服务器版本
- `IsFeatureSupported()` - 检查特性支持
- `GetVersion()` - 获取版本信息

### 适配器接口
提供统一的操作接口：
- `ItemAdapter` - 监控项操作接口
- `HostAdapter` - 主机操作接口
- 具体实现：`Zabbix6ItemAdapter`, `Zabbix7ItemAdapter`

### 数据结构
统一的数据结构，支持多版本格式：
- `HeadersV6` / `HeadersV7` - 不同版本的 headers 格式
- `ProxyHostID` / `ProxyID` - 不同版本的代理字段
- `MonitoredBy` - Zabbix 7.0+ 监控方式

## 🧪 测试和验证

### 版本兼容性测试
```go
// 测试不同版本的数据转换
item := zabbix.Item{
    HeadersV6: zabbix.HttpHeaders{"User-Agent": "Test"},
}

// 验证转换
headersV7 := zabbix.ConvertHeadersToV7(item.HeadersV6)
headersV6 := zabbix.ConvertHeadersToV6(headersV7)

// 验证数据一致性
if len(headersV7) == len(headersV6) {
    fmt.Println("版本转换正确")
}
```

### 特性检测测试
```go
// 测试特性检测
features := []string{
    zabbix.FeatureHistoryPush,
    zabbix.FeatureMFA,
    zabbix.FeatureProxyGroup,
}

for _, feature := range features {
    supported := api.IsFeatureSupported(feature)
    fmt.Printf("%s: %t\n", feature, supported)
}
```

## 📈 性能优化

### 版本缓存
- 版本检测结果会被缓存
- 避免重复的版本检测请求
- 提高后续操作的性能

### 适配器复用
- 适配器对象会被复用
- 避免重复创建适配器实例
- 减少内存分配开销

### 数据预处理
- 数据在发送前进行预处理
- 减少网络传输数据量
- 提高序列化性能

## ⚠️ 注意事项

1. **版本检测失败处理**
   - 如果版本检测失败，会默认使用 6.0 兼容模式
   - 可以通过日志查看版本检测的详细信息

2. **特性使用检查**
   - 使用新特性前请检查特性支持
   - 避免在不支持的版本上调用新功能

3. **数据格式兼容**
   - 优先使用新版本格式 (HeadersV7, ProxyID)
   - 库会自动处理格式转换

4. **错误处理**
   - 新增的错误类型包含版本相关信息
   - 便于调试版本兼容性问题

## 🔄 升级指南

### 从旧版本升级
1. 现有代码无需修改，保持完全兼容
2. 登录时会自动检测版本并初始化适配器
3. 可以逐步迁移到新的多版本 API

### 代码迁移建议
1. 使用 `api.IsZabbix7()` 检查版本
2. 使用 `api.IsFeatureSupported()` 检查特性
3. 使用新的数据结构字段 (HeadersV7, ProxyID)
4. 利用适配器简化版本处理逻辑

这个多版本支持的实现确保了 go-zabbix-api 能够同时支持 Zabbix 6.0 和 7.0，为用户提供了无缝的升级体验。