# Zabbix 6.0/7.0 多版本支持架构设计

## 🎯 项目目标

创建一个全新的go-zabbix-api分支，专门支持Zabbix 6.0和7.0两个版本的API，实现：
- 完全兼容Zabbix 6.0 API
- 完全支持Zabbix 7.0新特性
- 自动版本检测和自适应
- 向后兼容性保证

## 📊 Zabbix 7.0 vs 6.0 主要变更分析

### 🔴 破坏性变更（必须处理）

#### 1. Item/DiscoveryRule Headers和QueryFields结构变更
**变更内容**：
- 6.0: `headers: {"name": "value"}` (对象格式)
- 7.0: `headers: [{"name": "name", "value": "value"}]` (数组格式)
- 同样适用于`query_fields`字段

**影响范围**：
- `item.go`
- `discoveryrule.go`
- `itemprototype.go`

#### 2. Host Proxy相关字段重命名
**变更内容**：
- `proxy_hostid` → `proxyid`
- 新增`monitored_by`字段（必需）
- 移除`proxy_hosts`参数支持

**影响范围**：
- `host.go`
- `proxy.go`

#### 3. Dashboard Widget重大变更
**变更内容**：
- `plaintext` → `itemhistory`
- 字段命名规则变更：`str.str.index1.index2` → `str.index1.str.index2`
- 坐标范围变更：x(0-23→0-71), y(0-62→0-63), width(1-24→1-72), height(2-32→1-64)

**影响范围**：
- `dashboard.go`（如果存在）

#### 4. Script方法参数变更
**变更内容**：
- `script.getscriptsbyhosts`: 数组参数 → 对象参数
- `script.getscriptsbyevents`: 数组参数 → 对象参数

**影响范围**：
- `script.go`（如果存在）

### 🟢 新增功能（可选实现）

#### 1. History Push API
**新增方法**：`history.push`
**用途**：通过HTTP协议发送数据到Zabbix服务器

#### 2. MFA (多因子认证) 支持
**新增API**：
- `mfa.create`, `mfa.update`, `mfa.get`, `mfa.delete`
- `user.resettotp`
- Authentication对象新增`mfa_status`, `mfaid`字段

#### 3. Proxy Group API
**全新API**：`proxygroup`相关方法

#### 4. 新的Item类型
**新增类型**：`22 - Browser` (浏览器监控)

#### 5. 新的预处理类型
**新增类型**：
- `30 - SNMP get value`
- `14 - Matches regular expression` (LLD规则)

### 🟡 字段增强（渐进实现）

#### 1. Item/DiscoveryRule Timeout支持扩展
**扩展内容**：
- 更多item类型支持timeout字段
- 新增各种类型的timeout配置

#### 2. Proxy配置增强
**新增字段**：
- `address`, `port` (被动代理)
- `custom_timeouts`系列
- `timeout_browser`

## 🏗️ 多版本架构设计

### 核心设计原则

1. **版本检测优先**：连接时自动检测Zabbix版本
2. **渐进式兼容**：低版本功能在高版本中正常工作
3. **特性开关**：高版本特性通过配置启用
4. **结构统一**：使用统一的数据结构，通过标签区分版本

### 架构组件

#### 1. 版本管理器 (Version Manager)
```go
type VersionManager struct {
    serverVersion string
    is70          bool
    is60          bool
}

func (vm *VersionManager) DetectVersion(api *API) error
func (vm *VersionManager) IsFeatureSupported(feature string) bool
```

#### 2. 适配器模式 (Adapter Pattern)
```go
type ItemAdapter interface {
    CreateItems(items []Item) error
    GetItems(params Params) ([]Item, error)
}

type Zabbix6ItemAdapter struct{ ... }
type Zabbix7ItemAdapter struct{ ... }
```

#### 3. 统一数据结构
```go
type Item struct {
    // 通用字段
    ItemID    string `json:"itemid,omitempty"`
    Name      string `json:"name"`
    
    // 版本特定字段
    HeadersV6 map[string]string `json:"headers_v6,omitempty"`
    HeadersV7 []HeaderField     `json:"headers_v7,omitempty"`
    
    // 新特性字段
    Timeout   string `json:"timeout,omitempty"`
}

type HeaderField struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}
```

#### 4. 特性检测器
```go
type FeatureDetector struct {
    supportedFeatures map[string]bool
}

const (
    FeatureHistoryPush    = "history.push"
    FeatureMFA           = "mfa"
    FeatureProxyGroup    = "proxygroup"
    FeatureBrowserItem   = "browser_item"
)
```

### 实现策略

#### 阶段1：基础架构搭建
1. 创建版本管理器
2. 实现基础适配器接口
3. 建立特性检测机制

#### 阶段2：核心功能适配
1. Item/DiscoveryRule Headers结构适配
2. Host Proxy字段适配
3. Dashboard Widget适配（如果需要）

#### 阶段3：新功能实现
1. History Push API
2. MFA支持
3. Proxy Group API
4. Browser Item类型

#### 阶段4：测试和优化
1. 多版本兼容性测试
2. 性能优化
3. 文档完善

## 📋 实现计划

### 第一周：基础架构
- [x] 创建新分支`zabbix6-7-support`
- [ ] 实现版本管理器
- [ ] 设计适配器接口
- [ ] 建立测试框架

### 第二周：核心适配
- [ ] Item/DiscoveryRule Headers适配
- [ ] Host Proxy字段适配
- [ ] 基础功能测试

### 第三周：新功能开发
- [ ] History Push API实现
- [ ] MFA支持实现
- [ ] Proxy Group API实现

### 第四周：测试和文档
- [ ] 全面兼容性测试
- [ ] 性能基准测试
- [ ] 文档编写

## 🔧 技术实现细节

### 版本检测机制
```go
func (api *API) DetectVersion() (string, error) {
    version, err := api.Version()
    if err != nil {
        return "", err
    }
    
    // 解析版本号
    if strings.HasPrefix(version, "7.") {
        api.versionManager.is70 = true
    } else if strings.HasPrefix(version, "6.") {
        api.versionManager.is60 = true
    }
    
    return version, nil
}
```

### 自适应数据转换
```go
func (adapter *Zabbix7ItemAdapter) prepareHeaders(item Item) interface{} {
    if len(item.HeadersV6) > 0 {
        // 转换6.0格式到7.0格式
        headers := make([]HeaderField, 0, len(item.HeadersV6))
        for name, value := range item.HeadersV6 {
            headers = append(headers, HeaderField{
                Name:  name,
                Value: value,
            })
        }
        return headers
    }
    return item.HeadersV7
}
```

### 特性开关机制
```go
func (api *API) HistoryPush(data []HistoryData) error {
    if !api.versionManager.IsFeatureSupported(FeatureHistoryPush) {
        return errors.New("history.push not supported in this Zabbix version")
    }
    
    return api.callHistoryPush(data)
}
```

## 📈 预期收益

1. **兼容性**：同时支持Zabbix 6.0和7.0
2. **前瞻性**：为未来版本升级做好准备
3. **易用性**：开发者无需关心版本差异
4. **可维护性**：清晰的架构设计便于后续扩展

## 🎯 成功标准

1. ✅ 所有Zabbix 6.0功能正常工作
2. ✅ 所有Zabbix 7.0新功能可用
3. ✅ 自动版本检测100%准确
4. ✅ 性能不低于原版本
5. ✅ 向后兼容性100%保证

这个架构设计将使go-zabbix-api成为一个真正现代化、多版本兼容的Zabbix SDK。