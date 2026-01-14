# Zabbix API 实现计划

## 🎯 目标

按照优先级顺序实现缺失的Zabbix API功能，补充关键监控数据查询、告警管理和可视化界面功能。

## 📋 实施计划

### **阶段1: History API (历史数据查询) - 预计2天**
**优先级**: 🔥 高
**文件**: `history.go`

#### 实现内容
```go
// 历史数据查询API
func (api *API) HistoryGet(params HistoryGetOptions) (HistoryData, error)

// 支持的数据类型
type HistoryData struct {
    ItemID    string `json:"itemid"`
    Clock     int    `json:"clock"`
    Value     string `json:"value"`
    ns        int    `json:"ns"`
}
```

#### 核心方法
- `HistoryGet()` - 查询历史数据
- `HistoryGetByItem()` - 按监控项查询
- `HistoryGetByTimeRange()` - 按时间范围查询
- `HistoryGetFloat()` - 浮点历史数据
- `HistoryGetText()` - 文本历史数据

#### 关键考虑点
- 支持多种数据类型：float, unsigned, text, log, binary
- 时间范围过滤
- 分页支持
- 与现有API模式保持一致

---

### **阶段2: Trends API (趋势数据查询) - 预计2天**
**优先级**: 🔥 高
**文件**: `trends.go`

#### 实现内容
```go
// 趋势数据查询API
func (api *API) TrendsGet(params TrendsGetOptions) (TrendsData, error)

// 趋势数据结构
type TrendsData struct {
    ItemID    string `json:"itemid"`
    Clock     int    `json:"clock"`
    ValueMin  string `json:"value_min"`
    ValueAvg  string `json:"value_avg"`
    ValueMax  string `json:"value_max"`
}
```

#### 核心方法
- `TrendsGet()` - 查询趋势数据
- `TrendsGetByItem()` - 按监控项查询
- `TrendsGetSummary()` - 获取趋势统计信息

#### 关键考虑点
- 支持min/max/avg统计值
- 时间范围聚合
- 数据压缩和优化

---

### **阶段3: Events API (事件查询) - 预计3天**
**优先级**: 🔥 高
**文件**: `events.go`

#### 实现内容
```go
// 事件查询API
func (api *API) EventsGet(params EventsGetOptions) (Events, error)

// 事件数据结构
type Event struct {
    EventID      string `json:"eventid"`
    TriggerID    string `json:"triggerid"`
    Clock        int    `json:"clock"`
    Value        int    `json:"value"`
    Acknowledged int    `json:"acknowledged"`
    Severity     string `json:"severity"`
    HostID       string `json:"hostid"`
}
```

#### 核心方法
- `EventsGet()` - 查询事件
- `EventsGetByTrigger()` - 按触发器查询
- `EventsGetByHost()` - 按主机查询
- `EventsAcknowledge()` - 确认事件
- `EventsGetUnacknowledged()` - 获取未确认事件

#### 关键考虑点
- 事件确认状态管理
- 触发器关联查询
- 严重级别过滤
- 时间范围查询

---

### **阶段4: Action API (告警动作管理) - 预计4天**
**优先级**: 🔥 高
**文件**: `action.go`

#### 实现内容
```go
// 动作管理API
func (api *API) ActionsGet(params ActionGetOptions) (Actions, error)
func (api *API) ActionCreate(actions Actions) error
func (api *API) ActionUpdate(actions Actions) error
func (api *API) ActionDelete(actions Actions) error

// 动作结构
type Action struct {
    ActionID       string   `json:"actionid"`
    Name          string   `json:"name"`
    Eventsource   string   `json:"eventsource"`
    DefLongdata   string   `json:"def_longdata"`
    DefShortdata  string   `json:"def_shortdata"`
    RecoveryLong  string   `json:"recovery_long"`
    RecoveryShort string   `json:"recovery_short"`
    RTriggerID    string   `json:"r_triggerid"`
    Actions       []Operation `json:"operations"`
}
```

#### 核心方法
- `ActionsGet()` - 获取动作列表
- `ActionCreate()` - 创建动作
- `ActionUpdate()` - 更新动作
- `ActionDelete()` - 删除动作
- `ActionExecute()` - 执行动作

#### 关键考虑点
- 条件和操作管理
- 告警模板支持
- 恢复动作处理
- 权限检查

---

### **阶段5: Application API (应用管理) - 预计2天**
**优先级**: 🔥 高
**文件**: `application.go`

#### 实现内容
```go
// 应用管理API
func (api *API) ApplicationsGet(params ApplicationGetOptions) (Applications, error)
func (api *API) ApplicationCreate(applications Applications) error
func (api *API) ApplicationUpdate(applications Applications) error
func (api *API) ApplicationDelete(applications Applications) error

// 应用结构
type Application struct {
    ApplicationID string `json:"applicationid"`
    HostID        string `json:"hostid"`
    Name          string `json:"name"`
    Flags         string `json:"flags"`
}
```

#### 核心方法
- `ApplicationsGet()` - 获取应用列表
- `ApplicationCreate()` - 创建应用
- `ApplicationUpdate()` - 更新应用
- `ApplicationDelete()` - 删除应用
- `ApplicationsGetByHost()` - 按主机获取应用

#### 关键考虑点
- 与Item的关联管理
- 主机范围控制
- 应用权限检查

---

### **阶段6: Dashboard API (仪表板管理) - 预计5天**
**优先级**: 🔥 高
**文件**: `dashboard.go`

#### 实现内容
```go
// 仪表板管理API
func (api *API) DashboardsGet(params DashboardGetOptions) (Dashboards, error)
func (api *API) DashboardCreate(dashboards Dashboards) error
func (api *API) DashboardUpdate(dashboards Dashboards) error
func (api *API) DashboardDelete(dashboards Dashboards) error

// 仪表板结构
type Dashboard struct {
    DashboardID string   `json:"dashboardid"`
    Name        string   `json:"name"`
    UserID      string   `json:"userid"`
    Private     string   `json:"private"`
    Widgets     []Widget `json:"widgets"`
}
```

#### 核心方法
- `DashboardsGet()` - 获取仪表板列表
- `DashboardCreate()` - 创建仪表板
- `DashboardUpdate()` - 更新仪表板
- `DashboardDelete()` - 删除仪表板
- `DashboardShare()` - 共享仪表板

#### 关键考虑点
- Widget配置管理
- 权限和共享机制
- Zabbix 7.0变更支持
- 用户界面兼容性

---

### **阶段7: Discovery相关API (网络发现) - 预计3天**
**优先级**: 🟡 中
**文件**: `discovery.go`, `dhost.go`, `dservice.go`

#### 实现内容
```go
// 发现规则管理
func (api *API) DiscoveryRulesGet(params DiscoveryRuleGetOptions) (DiscoveryRules, error)
func (api *API) DiscoveryRuleCreate(rules DiscoveryRules) error

// 发现主机管理
func (api *API) DiscoveredHostsGet(params DHostGetOptions) (DiscoveredHosts, error)

// 发现服务管理
func (api *API) DiscoveredServicesGet(params DServiceGetOptions) (DiscoveredServices, error)
```

#### 核心功能
- 网络发现规则CRUD
- 发现主机管理
- 发现服务管理
- 自动注册处理

---

### **阶段8: Service API (服务监控) - 预计4天**
**优先级**: 🟡 中
**文件**: `service.go`

#### 实现内容
```go
// 服务监控API
func (api *API) ServicesGet(params ServiceGetOptions) (Services, error)
func (api *API) ServiceCreate(services Services) error
func (api *API) ServiceUpdate(services Services) error
func (api *API) ServiceDelete(services Services) error
func (api *API) ServiceGetSLA(params ServiceSLAOptions) (SLAs, error)
```

#### 核心功能
- 服务树形结构管理
- SLA计算和查询
- 服务依赖关系
- 监控状态管理

---

### **阶段9: ValueMap API (值映射管理) - 预计2天**
**优先级**: 🟡 中
**文件**: `valuemap.go`

#### 实现内容
```go
// 值映射API
func (api *API) ValueMapsGet(params ValueMapGetOptions) (ValueMaps, error)
func (api *API) ValueMapCreate(valueMaps ValueMaps) error
func (api *API) ValueMapUpdate(valueMaps ValueMaps) error
func (api *API) ValueMapDelete(valueMaps ValueMaps) error
```

#### 核心功能
- 映射规则定义
- 数值转换
- 模板引用管理

---

### **阶段10: Maintenance API (维护期管理) - 预计3天**
**优先级**: 🟡 中
**文件**: `maintenance.go`

#### 实现内容
```go
// 维护期API
func (api *API) MaintenancesGet(params MaintenanceGetOptions) (Maintenances, error)
func (api *API) MaintenanceCreate(maintenances Maintenances) error
func (api *API) MaintenanceUpdate(maintenances Maintenances) error
func (api *API) MaintenanceDelete(maintenances Maintenances) error
```

#### 核心功能
- 时间段定义
- 主机范围控制
- 维护期类型支持

---

## 🛠️ 实施策略

### 1. 代码结构
- 遵循现有项目结构和命名约定
- 保持API接口一致性
- 添加完整注释和文档

### 2. 测试策略
- 每个API模块添加单元测试
- 实现集成测试
- 验证多版本兼容性

### 3. 质量保证
- 代码审查和重构
- 性能测试和优化
- 错误处理完善

### 4. 文档更新
- 更新README.md
- 添加API使用示例
- 维护变更日志

## 📈 预期成果

完成所有阶段后，API覆盖率将从41.7%提升到约75%，主要收益：

- ✅ **监控数据查询**: 100%完整
- ✅ **告警动作管理**: 100%完整
- ✅ **可视化界面**: 80%完整
- ✅ **发现和自动化**: 70%完整
- ✅ **服务监控**: 100%完整

## 🎯 开始实施

现在开始阶段1：History API实现。