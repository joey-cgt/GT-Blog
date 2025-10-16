# GT-Blog 博客系统

一个现代化、功能完整的博客系统，基于前后端分离架构开发。

## 📚 项目简介

GT-Blog是一个功能丰富的个人博客平台，提供了完整的博客内容管理、前台展示以及统计分析功能。系统采用前后端分离架构，前端负责用户界面展示，后端提供API服务和业务逻辑处理。

## 🛠️ 技术栈

### 前端
- **框架**：Vue 3
- **UI组件库**：Element Plus
- **路由**：Vue Router 4
- **构建工具**：Vite 7
- **HTTP请求**：Axios
- **Markdown处理**：markdown-it、marked
- **数学公式**：KaTeX
- **代码高亮**：highlight.js
- **图表可视化**：ECharts、vue-echarts

### 后端
- **语言**：Go 1.25
- **Web框架**：Gin
- **ORM框架**：GORM
- **数据库**：MySQL
- **认证**：JWT
- **日志**：Zap
- **定时任务**：cron
- **架构模式**：DDD (领域驱动设计)、MVC

## 📁 项目结构

```
gt-blog/
├── backend/         # 后端Go代码
│   ├── internal/    # 内部包（DDD架构实现）
│   ├── pkg/         # 公共包（配置、日志、数据库等）
│   └── main.go      # 后端入口文件
├── frontend/        # 前端Vue代码
│   ├── src/         # 源代码
│   │   ├── api/     # API请求封装
│   │   ├── assets/  # 静态资源
│   │   ├── components/ # Vue组件
│   │   ├── layouts/ # 布局组件
│   │   ├── router/  # 路由配置
│   │   ├── views/   # 页面视图
│   │   └── main.js  # 前端入口文件
│   └── vite.config.js # Vite配置
└── README.md        # 项目说明文档
```

## ✨ 主要功能

### 前台功能
- 博客首页展示
- 文章列表与分类浏览
- 文章详情页（支持Markdown内容）
- 分类和标签筛选
- 专栏功能
- 关于页面

### 后台管理
- 管理员登录认证
- 仪表盘统计展示
- 文章管理（创建、编辑、发布、删除、置顶）
- 分类和标签管理
- 评论管理
- 专栏管理

## 🚀 快速开始

### 前置要求
- Go 1.25+
- Node.js 18+
- MySQL 5.7+

### 安装步骤

#### 1. 克隆项目
```bash
git clone https://github.com/joey-cgt/GT-Blog.git
cd gt-blog
```

#### 2. 配置后端
- 进入后端目录
```bash
cd backend
```
- 配置数据库连接
- 安装依赖
```bash
go mod tidy
```
- 启动后端服务
```bash
go run main.go
```

#### 3. 配置前端
- 进入前端目录
```bash
cd frontend
```
- 安装依赖
```bash
npm install
```
- 启动开发服务器
```bash
npm run dev
```

## 🔧 项目配置

### 后端配置
- 配置文件位置：`backend/pkg/config/config.yaml`
- 主要配置项：数据库连接、服务端口、日志设置等

### 前端配置
- 开发环境：修改相关环境变量
- 构建配置：`vite.config.js`

## 📊 项目特点

- **前后端分离**：清晰的架构分离，便于维护和扩展
- **现代化UI**：基于Element Plus构建的美观界面
- **丰富的Markdown功能**：支持代码高亮、数学公式等
- **完整的博客功能**：文章、分类、标签、专栏等功能齐全
- **统计分析**：包含访问统计和数据可视化功能
- **DDD架构**：后端采用领域驱动设计，代码结构清晰

## 🔒 安全特性
- JWT认证保护后台API
- SQL注入防护（使用GORM参数化查询）
- CORS配置管理
- 输入验证

## 📝 开发指南

### 后端开发
- 遵循DDD架构原则
- 代码结构按照层次分明的方式组织
- 测试驱动开发

### 前端开发
- 组件化开发
- 使用Vue 3 Composition API
- 遵循Element Plus设计规范

## 🤝 贡献指南

欢迎对项目进行贡献！

1. Fork 项目
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 📜 许可证

[MIT](https://choosealicense.com/licenses/mit/)

## 📧 联系方式

如有问题或建议，请联系项目维护者。