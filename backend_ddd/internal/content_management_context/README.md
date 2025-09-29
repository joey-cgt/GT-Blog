基于内容管理上下文的核心功能（文章、专栏、分类、标签的全生命周期管理），结合 RESTful API 设计规范，以下是前端请求 URL 的系统设计。URL 统一以`/api/admin`为前缀（标识管理后台接口，便于权限校验），按资源类型和操作场景分层，确保语义清晰且与后端接口层（`handler`/`router`）精准匹配：

### 一、文章管理模块（核心资源）

针对文章的草稿 / 发布状态流转、内容编辑、删除等操作：

| 功能场景            | HTTP 方法 | URL 路径                           | 说明（请求参数 / 示例）                                      |
| ------------------- | --------- | ---------------------------------- | ------------------------------------------------------------ |
| 1. 分页查询文章列表 | GET       | `/api/admin/articles`              | 查询参数：- `page=1&size=10`（分页）- `status=draft`（状态筛选：`draft`/`published`）- `category_id=3`（按分类筛选）- `keyword=ddd`（标题模糊搜索） |
| 2. 获取单篇文章详情 | GET       | `/api/admin/articles/{id}`         | `{id}`为文章唯一标识（返回完整内容 + 关联标签 / 分类 / 专栏信息，用于编辑回显） |
| 3. 发布新文章       | POST      | `/api/admin/articles`              | 请求体：`{title: "DDD实践指南", content: "...", category_id: 2, tag_ids: [1,3], column_id: 5, is_top: false}`（`status`自动设为`published`） |
| 4. 保存为草稿       | POST      | `/api/admin/articles/drafts`       | 请求体：同上（`status`自动设为`draft`，与发布接口语义分离，前端可绑定不同按钮） |
| 5. 编辑已发布文章   | PUT       | `/api/admin/articles/{id}`         | 请求体：支持部分字段更新（如仅修改`content`和`title`，保留其他字段） |
| 6. 草稿发布         | PUT       | `/api/admin/articles/{id}/publish` | 无请求体（仅将状态从`draft`转为`published`，触发发布事件）   |
| 7. 删除单篇文章     | DELETE    | `/api/admin/articles/{id}`         | 支持删除草稿 / 已发布文章（后端自动清理关联的标签关系）      |
| 8. 批量删除文章     | DELETE    | `/api/admin/articles/batch`        | 请求体：`{ids: ["a1", "a2", "a3"]}`（优化多文章删除效率）    |
| 9. 置顶 / 取消置顶  | PATCH     | `/api/admin/articles/{id}/top`     | 请求体：`{is_top: true}`（单独处理置顶状态，与内容编辑解耦） |

### 二、专栏管理模块

针对专栏的创建、编辑及文章关联管理：

| 功能场景              | HTTP 方法 | URL 路径                                       | 说明（请求参数 / 示例）                                      |
| --------------------- | --------- | ---------------------------------------------- | ------------------------------------------------------------ |
| 1. 分页查询专栏列表   | GET       | `/api/admin/columns`                           | 查询参数：`page=1&size=10&keyword=架构`（分页 + 标题搜索）   |
| 2. 获取专栏详情       | GET       | `/api/admin/columns/{id}`                      | 返回专栏信息 + 关联文章数（如`{id: 5, title: "微服务", article_count: 8, ...}`） |
| 3. 创建新专栏         | POST      | `/api/admin/columns`                           | 请求体：`{title: "Go语言进阶", description: "从入门到实战", cover_url: "https://xxx.png"}` |
| 4. 编辑专栏信息       | PUT       | `/api/admin/columns/{id}`                      | 请求体：支持部分字段更新（如仅修改`cover_url`或`description`） |
| 5. 删除专栏           | DELETE    | `/api/admin/columns/{id}`                      | 后端自动解除与文章的关联（不删除文章本身）                   |
| 6. 查询专栏关联文章   | GET       | `/api/admin/columns/{id}/articles`             | 查询参数：`page=1&size=10`（用于管理专栏内的文章列表）       |
| 7. 关联文章到专栏     | PUT       | `/api/admin/columns/{id}/articles`             | 请求体：`{article_ids: ["a1", "a2"]}`（批量绑定文章）        |
| 8. 解除文章与专栏关联 | DELETE    | `/api/admin/columns/{id}/articles/{articleId}` | 单独移除某篇文章的关联（精细操作）                           |

### 三、分类管理模块

针对文章分类的基础管理（分类为文章必填属性）：

| 功能场景                  | HTTP 方法 | URL 路径                     | 说明（请求参数 / 示例）                                      |
| ------------------------- | --------- | ---------------------------- | ------------------------------------------------------------ |
| 1. 获取所有分类（带统计） | GET       | `/api/admin/categories`      | 返回列表：`[{id: 2, name: "技术", article_count: 30}, ...]`（含关联文章数，无需分页） |
| 2. 创建新分类             | POST      | `/api/admin/categories`      | 请求体：`{name: "后端开发", description: "Java/Go相关文章"}`（名称唯一校验） |
| 3. 编辑分类信息           | PUT       | `/api/admin/categories/{id}` | 请求体：`{name: "服务端开发", description: "更新描述"}`（支持重命名和描述修改） |
| 4. 删除分类               | DELETE    | `/api/admin/categories/{id}` | 后端校验：若分类下有关联文章，返回错误提示（需先迁移文章至其他分类） |

### 四、标签管理模块

针对标签的基础管理（文章可关联多个标签）：

| 功能场景                  | HTTP 方法 | URL 路径               | 说明（请求参数 / 示例）                                      |
| ------------------------- | --------- | ---------------------- | ------------------------------------------------------------ |
| 1. 获取所有标签（带统计） | GET       | `/api/admin/tags`      | 返回列表：`[{id: 3, name: "领域驱动设计", article_count: 12}, ...]`（含关联文章数） |
| 2. 创建新标签             | POST      | `/api/admin/tags`      | 请求体：`{name: "DDD"}`（名称唯一校验，避免重复）            |
| 3. 编辑标签名称           | PUT       | `/api/admin/tags/{id}` | 请求体：`{name: "领域驱动设计"}`（支持标签重命名）           |
| 4. 删除标签               | DELETE    | `/api/admin/tags/{id}` | 后端自动解除与所有文章的关联（通过`article_tags`中间表）     |

### 设计规范说明

1. **权限隔离**：所有 URL 以`/api/admin`为前缀，后端可通过该前缀统一拦截并验证管理员 Token（如 JWT），与面向访客的公开 API（如`/api/articles`）严格区分，避免权限越界。
2. RESTful 语义
   - 资源以复数名词表示（`articles`/`columns`），符合 RESTful 规范；
   - HTTP 方法表达操作类型（GET 查询、POST 创建、PUT 全量更新、PATCH 部分更新、DELETE 删除）；
   - 特殊操作（如草稿发布）通过子路径（`/publish`）明确语义，避免接口职责模糊。
3. 前端适配性
   - 分页、筛选参数通过 URL 查询参数传递（如`?page=1&status=draft`），无需修改请求体结构；
   - 批量操作通过`/batch`路径统一处理，便于前端封装批量删除组件；
   - 关联操作（如专栏 - 文章绑定）通过嵌套路径（`/columns/{id}/articles`）清晰表达资源关系，前端可按层级组织 API 调用逻辑。
4. **与后端结构匹配**：URL 路径与后端`router`配置一一对应（如`/api/admin/articles`对应`articles_handler`），请求 / 响应 DTO 与前端参数结构完全对齐，减少前后端协作成本。

该设计覆盖了内容管理上下文的所有核心操作，前端可按资源类型封装 API 模块（如`ArticleApi`、`ColumnApi`），每个模块对应一组 URL，实现代码的模块化和可维护性。