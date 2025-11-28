# Zabbix 6.0 API适配需求分析

## 项目概述

本文档分析go-zabbix-api项目对Zabbix 6.0 API的适配需求，为zabbix6分支的开发提供指导。

## 当前项目结构分析

项目包含以下核心模块：

### 基础模块
- **base.go**: 基础API客户端和请求/响应结构
  - API结构体：存储连接信息和认证token
  - Request/Response结构体：处理JSON-RPC通信
  - Config结构体：配置管理

### 核心业务模块
- **host.go**: 主机管理
  - Host, Hosts结构体
  - AvailableType, StatusType常量定义
- **item.go**: 监控项管理
  - Item, Items结构体
  - ItemType, ValueType, DataType, DeltaType常量
- **trigger.go**: 触发器管理
- **template.go**: 模板管理
- **host_group.go**: 主机组管理
- **application.go**: 应用管理
- **macro.go**: 宏管理
- **host_interface.go**: 主机接口管理

### 测试模块
每个核心模块都有对应的测试文件，确保API功能正确性。

## Zabbix 6.0 API关键变更点

### 1. UUID支持（重大变更）
- **变更内容**：所有API对象新增UUID属性
- **影响范围**：host, item, trigger, template, hostgroup, graph等所有对象
- **兼容性**：向后兼容，UUID为可选字段
- **适配要求**：为所有对象结构体添加UUID字段，支持UUID更新操作

### 2. 权限模型变更（重大变更）
- **变更内容**：Admin/User类型用户权限被严格限制
- **具体变更**：
  - `user.get`只返回有限的用户属性
  - `mediatype.get`权限变更，Admin用户只能访问部分属性
  - `alert.get`权限限制
- **影响范围**：用户管理、媒体类型、告警相关API
- **适配要求**：更新相关API调用逻辑，处理权限限制

### 3. 参数要求变更（中等影响）
- **HTTP Agent类型**：
  - item和discoveryrule的interfaceid不再强制要求
  - 影响HTTP agent类型的监控项创建
- **认证方法**：
  - `user.checkAuthentication`新增token参数
- **字段长度**：
  - user.create/update的url字段长度从255增加到2048字符

### 4. 新增功能（增强功能）
- **计算型item**：支持文本、日志、字符类型，不再仅限于数值
- **HTTP方法**：新增PATCH, HEAD, OPTIONS, TRACE, CONNECT方法支持
- **压缩内容**：支持libcurl支持的所有编码格式
- **JavaScript引擎**：增强HTTP请求处理能力

## 需要适配的核心模块

### 高优先级适配（必须完成）

#### 1. 所有对象结构体UUID字段添加
**涉及文件**：
- host.go - Host结构体
- item.go - Item结构体  
- trigger.go - Trigger结构体
- template.go - Template结构体
- host_group.go - HostGroup结构体
- 以及其他所有API对象结构体

**适配内容**：
```go
type Host struct {
    HostID string `json:"hostid,omitempty"`
    UUID   string `json:"uuid,omitempty"`  // 新增字段
    // ... 其他字段
}
```

#### 2. 用户权限API适配
**涉及文件**：
- 可能需要新增user.go文件或修改现有用户相关代码

**适配内容**：
- 更新user.get方法的返回字段限制
- 处理mediatype.get权限变更
- 实现新的权限检查逻辑

#### 3. HTTP Agent参数处理
**涉及文件**：
- item.go - Item结构体和创建逻辑
- 可能涉及discoveryrule相关代码

**适配内容**：
- 修改interfaceid为可选参数
- 更新相关验证逻辑

### 中优先级适配（建议完成）

#### 1. 认证方法增强
- user.checkAuthentication添加token参数支持
- 更新认证相关API调用

#### 2. Item类型扩展
- 支持计算型item的文本、日志、字符类型
- 添加新的item常量定义

#### 3. 字段长度更新
- user.url字段长度从255增加到2048
- 更新相关验证和注释

### 低优先级适配（可选完成）

#### 1. 新HTTP方法支持
- 添加PATCH, HEAD, OPTIONS, TRACE, CONNECT方法
- 更新HTTP客户端配置

#### 2. 测试和验证
- 在Zabbix 6.0环境中测试所有API
- 更新测试用例以适配6.0变更

## 详细开发计划

### 第一阶段：基础结构适配（高优先级）
**目标**：确保基本功能在Zabbix 6.0中正常工作

**任务清单**：
1. 为所有API对象添加UUID字段
2. 用户权限相关API适配
3. HTTP agent类型item的interfaceid参数处理
4. 基础功能测试验证

**预期时间**：2-3天

### 第二阶段：新功能支持（中优先级）
**目标**：支持Zabbix 6.0的新特性

**任务清单**：
1. 认证方法token参数支持
2. Item类型扩展（计算型item）
3. 字段长度限制更新
4. 新功能测试验证

**预期时间**：2-3天

### 第三阶段：完善和测试（低优先级）
**目标**：完善功能和提升兼容性

**任务清单**：
1. 新HTTP方法支持
2. 压缩内容处理
3. 全面测试和性能验证
4. 文档更新

**预期时间**：1-2天

## 风险评估

### 高风险项
1. **权限模型变更**：可能影响现有用户管理功能
2. **UUID字段**：需要确保所有对象正确添加

### 中风险项
1. **参数要求变更**：可能破坏现有创建逻辑
2. **新功能集成**：可能与现有代码产生冲突

### 低风险项
1. **文档更新**：不影响核心功能
2. **测试用例更新**：确保功能正确性

## 兼容性策略

### 向后兼容
- 保持现有API接口不变
- 新增字段使用omitempty标签
- 可选参数不破坏现有调用

### 测试策略
- 保留现有测试用例
- 新增6.0特性测试
- 多版本兼容性测试

## 总结

Zabbix 6.0 API适配工作主要集中在UUID支持、权限模型变更和参数要求调整三个方面。通过分阶段实施，可以确保项目在支持新版本的同时保持向后兼容性。

建议优先完成高优先级适配，确保基本功能正常，然后逐步添加新功能支持。整个适配过程预计需要5-8天时间。