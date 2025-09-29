领域驱动设计(DDD)架构的博客系统后端。主要特点如下：

### 架构特点

1. **领域驱动设计(DDD)**:
   - 按业务领域划分为多个上下文(Context)
   - 每个上下文内部采用分层架构
   - 层次结构清晰：接口层(L0) -> 应用层(L1) -> 领域层(L2) -> 基础设施层(L3)
2. **技术栈**:
   - Go 语言(1.25.0)
   - Gin Web 框架
   - GORM ORM 库

### 主要业务上下文

1. **content_management_context**: 内容管理上下文，负责文章、专栏、分类、标签的管理
2. **authentication_context**: 认证上下文(尚未实现)
3. **analytics_context**: 分析上下文
4. **content_presentation_context**: 内容展示上下文
5. **profile_context**: 用户资料上下文

### 分层架构(以 content_management_context 为例)

1. **L0_interface**: 接口层
   - converter: 数据转换
   - dto: 数据传输对象
   - handler: 请求处理器
   - router: API 路由定义
2. **L1_application**: 应用层
   - dto: 应用层数据传输对象
   - service: 应用服务
3. **L2_domain**: 领域层
   - event: 领域事件
   - model: 领域模型
   - repository: 仓储接口
   - service: 领域服务
4. **L3_infrastructure**: 基础设施层
   - eventbus: 事件总线
   - persistence: 持久化实现

### 共享组件(common)

- error: 错误处理
- eventbus: 事件总线
- logging: 日志记录
- middleware: 中间件

这个项目展示了 DDD 在 Go 语言中的实践，通过清晰的分层和上下文边界，使得系统更加模块化、可维护，并且能够更好地应对复杂业务需求的变化。目前看起来项目还在开发中，一些上下文尚未实现。