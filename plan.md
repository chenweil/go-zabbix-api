# Zabbix 6.0 API适配详细开发计划

## 项目概述

本文档详细规划go-zabbix-api项目对Zabbix 6.0 API的适配工作，确保项目在支持新版本的同时保持向后兼容性。

## 开发时间表

**总预估时间：6-8个工作日**
- 第一阶段：3天（高优先级任务）
- 第二阶段：2天（中优先级任务）
- 第三阶段：1-2天（低优先级任务和测试）

## 第一阶段：基础结构适配（高优先级）

### 任务1.1：为所有API对象添加UUID字段
**优先级：高 | 预估时间：1天 | 依赖：无**

#### 涉及文件和修改点：

1. **host.go** - 修改Host结构体
   ```go
   type Host struct {
       HostID     string        `json:"hostid,omitempty"`
       UUID       string        `json:"uuid,omitempty"`  // 新增字段
       Host       string        `json:"host"`
       // ... 其他字段保持不变
   }
   ```

2. **item.go** - 修改Item结构体
   ```go
   type Item struct {
       ItemID       string    `json:"itemid,omitempty"`
       UUID         string    `json:"uuid,omitempty"`    // 新增字段
       Delay        string    `json:"delay"`
       // ... 其他字段保持不变
   }
   ```

3. **trigger.go** - 修改Trigger结构体
   ```go
   type Trigger struct {
       TriggerID   string `json:"triggerid,omitempty"`
       UUID        string `json:"uuid,omitempty"`       // 新增字段
       Description string `json:"description"`
       // ... 其他字段保持不变
   }
   ```

4. **template.go** - 修改Template结构体
   ```go
   type Template struct {
       TemplateID      string       `json:"templateid,omitempty"`
       UUID            string       `json:"uuid,omitempty"`      // 新增字段
       Host            string       `json:"host"`
       // ... 其他字段保持不变
   }
   ```

5. **host_group.go** - 修改HostGroup结构体
   ```go
   type HostGroup struct {
       GroupID  string       `json:"groupid,omitempty"`
       UUID     string       `json:"uuid,omitempty"`        // 新增字段
       Name     string       `json:"name"`
       // ... 其他字段保持不变
   }
   ```

6. **application.go** - 修改Application结构体（如果存在）
7. **macro.go** - 修改Macro结构体（如果需要）
8. **host_interface.go** - 修改HostInterface结构体（如果需要）

#### 验证方法：
- 运行现有测试确保向后兼容性
- 创建新的测试用例验证UUID字段的处理

### 任务1.2：用户权限API适配
**优先级：高 | 预估时间：0.5天 | 依赖：任务1.1**

#### 涉及文件和修改点：

1. **base.go** - 修改Login方法和认证逻辑
   ```go
   // 修改user.checkAuthentication调用以支持token参数
   func (api *API) CheckAuthentication(token string) (valid bool, err error) {
       params := map[string]string{"token": token}
       response, err := api.CallWithError("user.checkAuthentication", params)
       if err != nil {
           return
       }
       // 处理响应逻辑
       return
   }
   ```

2. **新增或修改user相关代码** - 处理权限限制
   - 更新user.get方法的返回字段处理
   - 添加权限检查逻辑
   - 处理mediatype.get权限变更

#### 验证方法：
- 测试Admin用户权限限制
- 验证mediatype.get的权限控制
- 确保user.get返回正确的字段子集

### 任务1.3：HTTP Agent参数处理
**优先级：高 | 预估时间：0.5天 | 依赖：任务1.1**

#### 涉及文件和修改点：

1. **item.go** - 修改Item结构体和创建逻辑
   ```go
   type Item struct {
       // ... 其他字段
       InterfaceID  string    `json:"interfaceid,omitempty"`  // 确保为可选字段
       // ... HTTP Agent字段
   }
   ```

2. **修改ItemsCreate方法** - 更新验证逻辑
   - 对于HTTP Agent类型，interfaceid不再是必需参数
   - 更新相关验证和错误处理

#### 验证方法：
- 测试HTTP Agent类型监控项的创建
- 验证interfaceid为可选参数
- 确保其他类型监控项的正常创建

### 任务1.4：基础功能测试验证
**优先级：高 | 预估时间：1天 | 依赖：任务1.1-1.3**

#### 测试任务：
1. 运行所有现有测试用例
2. 创建Zabbix 6.0环境下的基础功能测试
3. 验证CRUD操作的正常工作
4. 确保向后兼容性

#### 测试覆盖：
- Host CRUD操作
- Item CRUD操作
- Trigger CRUD操作
- Template CRUD操作
- HostGroup CRUD操作

## 第二阶段：新功能支持（中优先级）

### 任务2.1：认证方法增强
**优先级：中 | 预估时间：0.5天 | 依赖：任务1.2**

#### 涉及文件和修改点：

1. **base.go** - 增强认证相关方法
   ```go
   // 为Login方法添加token参数支持
   func (api *API) LoginWithToken(user, password, token string) (auth string, err error) {
       params := map[string]interface{}{
           "user": user,
           "password": password,
           "token": token,
       }
       // 实现逻辑
   }
   ```

#### 验证方法：
- 测试token参数的认证
- 验证向后兼容性

### 任务2.2：Item类型扩展
**优先级：中 | 预估时间：0.5天 | 依赖：任务1.3**

#### 涉及文件和修改点：

1. **item.go** - 扩展计算型item支持
   ```go
   // 为计算型item添加新的值类型支持
   const (
       // 现有常量...
       CalculatedText ValueType = 5  // 新增：计算型文本
       CalculatedLog  ValueType = 6  // 新增：计算型日志
       CalculatedChar ValueType = 7  // 新增：计算型字符
   )
   ```

#### 验证方法：
- 创建计算型item的测试用例
- 验证新值类型的支持

### 任务2.3：字段长度更新
**优先级：中 | 预估时间：0.5天 | 依赖：任务1.2**

#### 涉及文件和修改点：

1. **更新User相关结构体** - URL字段长度从255增加到2048
   ```go
   type User struct {
       UserID string `json:"userid,omitempty"`
       Url    string `json:"url,omitempty"`  // 更新注释：最大长度2048字符
       // ... 其他字段
   }
   ```

#### 验证方法：
- 测试长URL的创建和更新
- 验证字段长度限制

### 任务2.4：新功能测试验证
**优先级：中 | 预估时间：0.5天 | 依赖：任务2.1-2.3**

#### 测试任务：
1. 测试认证方法增强
2. 验证Item类型扩展
3. 测试字段长度更新
4. 集成测试

## 第三阶段：完善和测试（低优先级）

### 任务3.1：新HTTP方法支持
**优先级：低 | 预估时间：0.5天 | 依赖：任务1.3**

#### 涉及文件和修改点：

1. **item.go** - 添加新的HTTP方法支持
   ```go
   const (
       // 现有HTTP方法常量...
       HTTPMethodPATCH   string = "PATCH"
       HTTPMethodHEAD    string = "HEAD"
       HTTPMethodOPTIONS string = "OPTIONS"
       HTTPMethodTRACE   string = "TRACE"
       HTTPMethodCONNECT string = "CONNECT"
   )
   ```

#### 验证方法：
- 测试新HTTP方法的监控项创建
- 验证方法参数的正确性

### 任务3.2：压缩内容处理
**优先级：低 | 预估时间：0.5天 | 依赖：任务3.1**

#### 涉及文件和修改点：

1. **base.go** - 增强HTTP客户端配置
   ```go
   // 添加压缩内容支持
   func (api *API) configureCompression() {
       // 配置libcurl支持的所有编码格式
   }
   ```

#### 验证方法：
- 测试压缩内容的处理
- 验证性能改进

### 任务3.3：全面测试和性能验证
**优先级：低 | 预估时间：1天 | 依赖：任务3.1-3.2**

#### 测试任务：
1. 全面功能测试
2. 性能基准测试
3. 多版本兼容性测试
4. 压力测试

### 任务3.4：文档更新
**优先级：低 | 预估时间：0.5天 | 依赖：任务3.3**

#### 文档任务：
1. 更新README.md
2. 添加Zabbix 6.0特性说明
3. 更新API文档
4. 添加迁移指南

## 风险控制和回滚方案

### 高风险项控制

#### 1. UUID字段添加风险
**风险描述**：可能破坏现有API调用
**控制措施**：
- 所有UUID字段使用`omitempty`标签
- 保持现有字段不变
- 逐步添加，每添加一个字段就进行测试

**回滚方案**：
- 移除UUID字段，恢复原始结构体定义
- 重新运行测试确保功能正常

#### 2. 权限模型变更风险
**风险描述**：可能影响现有用户管理功能
**控制措施**：
- 保持现有API接口不变
- 添加新的权限检查方法
- 在非生产环境充分测试

**回滚方案**：
- 恢复原始的权限处理逻辑
- 移除新的权限检查代码

### 中风险项控制

#### 1. 参数要求变更风险
**风险描述**：可能破坏现有创建逻辑
**控制措施**：
- 保持向后兼容性
- 添加参数验证的容错处理
- 详细测试各种参数组合

**回滚方案**：
- 恢复原始的参数验证逻辑
- 更新相关文档说明

#### 2. 新功能集成风险
**风险描述**：可能与现有代码产生冲突
**控制措施**：
- 独立模块实现新功能
- 使用特性开关控制新功能
- 充分的单元测试

**回滚方案**：
- 禁用新功能特性开关
- 移除新功能相关代码

### 通用回滚策略

1. **版本控制**：每个阶段完成后创建git标签
2. **分支管理**：在独立分支进行开发，确保主分支稳定
3. **测试覆盖**：每个修改都有对应的测试用例
4. **渐进式部署**：先在测试环境验证，再部署到生产环境

## 测试策略

### 单元测试
- 每个修改的结构体和方法都有对应的测试用例
- 测试覆盖率达到90%以上
- 包含边界条件和异常情况测试

### 集成测试
- 测试各模块间的协作
- 验证API调用的完整性
- 测试数据流和错误处理

### 兼容性测试
- 在Zabbix 5.x版本测试向后兼容性
- 在Zabbix 6.0版本测试新功能
- 测试不同版本的API响应处理

### 性能测试
- 对比修改前后的性能指标
- 测试大量数据处理的性能
- 验证内存使用情况

## 质量保证

### 代码审查
- 每个代码修改都需要经过代码审查
- 检查代码风格和最佳实践
- 确保符合Go语言规范

### 自动化测试
- 集成CI/CD流水线
- 自动运行测试套件
- 自动生成测试报告

### 文档同步
- 代码修改同步更新文档
- 提供详细的变更日志
- 添加使用示例和最佳实践

## 交付物清单

### 代码交付物
- [ ] 修改后的所有.go源文件
- [ ] 新增的测试文件
- [ ] 更新的示例代码

### 文档交付物
- [ ] 更新的README.md
- [ ] API变更说明文档
- [ ] 迁移指南
- [ ] 测试报告

### 部署交付物
- [ ] 构建脚本
- [ ] 部署说明
- [ ] 配置文件模板

## 总结

本开发计划采用分阶段实施的方式，优先确保基础功能的兼容性，然后逐步添加新功能支持。通过详细的风险控制和回滚方案，确保开发过程的安全性和可控性。

整个适配过程预计需要6-8个工作日，完成后项目将完全支持Zabbix 6.0 API，同时保持对旧版本的向后兼容性。