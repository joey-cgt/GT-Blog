# GT-Blog 博客系统

一个基于Vue.js的现代化博客系统，提供文章管理、专栏分类、标签系统等功能。

## 功能特点

- 响应式设计，适配各种设备
- 专栏分类文章展示
- 文章分页浏览
- 简洁美观的UI界面

## 技术栈

- Vue 3 (Composition API)
- Vue Router
- Pinia (状态管理)
- Axios (HTTP请求)
- Tailwind CSS (可选)

## 项目结构

```
gt-blog/
├── public/            # 静态资源
└── src/
    ├── assets/        # 静态资源
    ├── components/    # 公共组件
    ├── router/        # 路由配置
    ├── store/         # 状态管理
    ├── views/         # 页面组件
    ├── App.vue        # 根组件
    └── main.js        # 入口文件
```

## 安装步骤

1. 克隆项目仓库
```bash
git clone https://github.com/your-repo/gt-blog.git
```

2. 安装依赖
```bash
npm install
```

3. 启动开发服务器
```bash
npm run dev
```

4. 构建生产版本
```bash
npm run build
```

## 开发指南

1. 创建新分支
```bash
git checkout -b feature/your-feature
```

2. 提交代码
```bash
git add .
git commit -m "描述你的修改"
git push origin feature/your-feature
```

3. 创建Pull Request

## 贡献指南

欢迎提交Pull Request或报告Issue。