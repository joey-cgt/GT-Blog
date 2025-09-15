// 博客数据存储
import { reactive } from 'vue'

// 博客信息
export const blogInfo = reactive({
  title: '欢迎来到 Joey 的个人博客',
  author: 'Joey_Chen',
  avatar: 'https://s21.ax1x.com/2025/09/07/pV2fri6.jpg',
  email:'joycgt@126.com',
  wechat: 'joycgt',
  description: '让学习成为一种习惯',
  socialLinks: [
    { name: 'GitHub', icon: 'github', url: 'https://github.com/joey-cgt' },
    { name: '稀土掘金', icon: 'weibo', url: 'https://juejin.cn/user/2522755752011596/posts' },
  ]
})

// 文章列表
export const articles = reactive([
  {
    id: 1,
    title: 'Vue3 组合式API完全指南',
    cover: 'https://img1.baidu.com/it/u=2565709611,1131746988&fm=253&fmt=auto&app=138&f=JPEG?w=1023&h=500',
    summary: 'Vue3的组合式API（Composition API）是Vue3最重要的特性之一，它提供了一种更灵活的方式来组织组件逻辑...',
    date: '2025-09-05',
    category: '前端开发',
    categoryId: 1,
    tags: ['Vue3', 'JavaScript', '前端框架'],
    views: 1250,
    likes: 100
  },
  {
    id: 2,
    title: 'TypeScript高级类型系统详解',
    cover: 'https://img1.baidu.com/it/u=1262645170,3593634273&fm=253&fmt=auto&app=138&f=JPEG?w=841&h=500',
    summary: 'TypeScript的类型系统非常强大，本文将深入探讨条件类型、映射类型、类型守卫等高级特性的使用方法...',
    date: '2025-09-03',
    category: '编程语言',
    categoryId: 4,
    tags: ['TypeScript', 'JavaScript', '类型系统'],
    views: 980
  },
  {
    id: 3,
    title: '微服务架构设计实践',
    cover: 'https://img2.baidu.com/it/u=1097097777,3390036665&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500',
    summary: '本文分享了在大型项目中实施微服务架构的经验，包括服务拆分、通信方式选择、数据一致性保证等关键问题的解决方案...',
    date: '2025-09-01',
    category: '后端架构',
    categoryId: 2,
    tags: ['微服务', '系统设计', '架构'],
    views: 1560,
    likes: 100
  },
  {
    id: 4,
    title: 'Docker容器化应用部署指南',
    cover: 'https://img1.baidu.com/it/u=2464941554,996813356&fm=253&fmt=auto&app=138&f=PNG?w=750&h=500',
    summary: '容器化技术已成为现代应用部署的标准方式，本文将介绍如何使用Docker高效部署和管理你的应用...',
    date: '2025-08-28',
    category: 'DevOps',
    categoryId: 3,
    tags: ['Docker', '容器化', 'CI/CD'],
    views: 1120,
    likes: 100
  },
  {
    id: 5,
    title: 'React与Vue的技术选型对比',
    cover: 'https://img1.baidu.com/it/u=3877579004,3172355599&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=500',
    summary: '在前端框架选型时，React和Vue常常是主要的候选者。本文从多个维度对比这两个框架，帮助你做出更明智的技术选择...',
    date: '2025-08-25',
    category: '前端开发',
    categoryId: 1,
    tags: ['React', 'Vue', '框架对比'],
    views: 2100,
    likes: 100
  },
  {
    id: 6,
    title: 'Node.js性能优化最佳实践',
    cover: 'https://img2.baidu.com/it/u=1702168081,2133302705&fm=253&fmt=auto&app=138&f=PNG?w=657&h=438',
    summary: 'Node.js作为高性能的JavaScript运行时，性能优化至关重要。本文分享Node.js应用性能优化的实用技巧和最佳实践...',
    date: '2025-08-20',
    category: '后端架构',
    categoryId: 2,
    tags: ['Node.js', '性能优化', 'JavaScript'],
    views: 890,
    likes: 100
  },
  {
    id: 7,
    title: 'Webpack 5配置完全指南',
    cover: 'https://img2.baidu.com/it/u=1329152739,4002170524&fm=253&fmt=auto&app=138&f=JPEG?w=1032&h=500',
    summary: 'Webpack 5带来了许多新特性和改进，本文将详细介绍Webpack 5的配置方法和优化技巧，帮助你构建更高效的前端项目...',
    date: '2025-08-18',
    category: '前端开发',
    categoryId: 1,
    tags: ['Webpack', '构建工具', '前端工程化'],
    views: 1020,
    likes: 100
  },
  {
    id: 8,
    title: 'MySQL索引优化原理与实践',
    cover: 'https://t14.baidu.com/it/u=1901670726,3719096885&fm=224&app=112&f=JPEG?w=500&h=500',
    summary: '数据库索引是提升查询性能的关键技术。本文深入探讨MySQL索引的工作原理，并分享实际项目中的索引优化实践经验...',
    date: '2025-08-15',
    category: '数据库',
    categoryId: 5,
    tags: ['MySQL', '数据库', '索引优化'],
    views: 1350,
    likes: 100
  },
  {
    id: 9,
    title: 'Git高级技巧与团队协作',
    cover: 'https://img2.baidu.com/it/u=47659589,2269443285&fm=253&fmt=auto&app=120&f=JPEG?w=889&h=500',
    summary: 'Git是现代软件开发中不可或缺的版本控制工具。本文分享Git的高级使用技巧和团队协作中的最佳实践...',
    date: '2025-08-12',
    category: 'DevOps',
    categoryId: 3,
    tags: ['Git', '版本控制', '团队协作'],
    views: 780,
    likes: 100
  },
  {
    id: 10,
    title: 'CSS Grid布局实战指南',
    cover: 'https://img2.baidu.com/it/u=1252614614,3252081911&fm=253&fmt=auto&app=138&f=JPEG?w=893&h=500',
    summary: 'CSS Grid是强大的二维布局系统，本文将通过实际案例展示如何使用CSS Grid创建复杂而响应式的网页布局...',
    date: '2025-08-10',
    category: '前端开发',
    categoryId: 1,
    tags: ['CSS', 'Grid布局', '响应式设计'],
    views: 950,
    likes: 100
  },
  {
    id: 11,
    title: 'RESTful API设计规范',
    cover: 'https://img0.baidu.com/it/u=591057997,328567836&fm=253&fmt=auto&app=138&f=JPEG?w=779&h=500',
    summary: '良好的API设计是构建可维护后端服务的基础。本文详细介绍RESTful API的设计原则、最佳实践和常见陷阱...',
    date: '2025-08-08',
    category: '后端架构',
    categoryId: 2,
    tags: ['RESTful', 'API设计', '后端开发'],
    views: 1100,
    likes: 100
  },
  {
    id: 12,
    title: '前端安全最佳实践',
    cover: 'https://img2.baidu.com/it/u=3392737198,17976616&fm=253&fmt=auto&app=120&f=JPEG?w=667&h=500',
    summary: 'Web应用安全是每个开发者都需要重视的问题。本文介绍常见的前端安全漏洞和防护措施，帮助构建更安全的Web应用...',
    date: '2025-08-05',
    category: '前端开发',
    categoryId: 1,
    tags: ['Web安全', 'XSS', 'CSRF'],
    views: 870,
    likes: 100
  },
  {
    id: 13,
    title: 'Docker Compose多容器应用部署',
    cover: 'https://img0.baidu.com/it/u=1876534143,3393961615&fm=253&fmt=auto&app=138&f=JPEG?w=577&h=500',
    summary: 'Docker Compose简化了多容器应用的部署和管理。本文通过实际案例演示如何使用Docker Compose编排复杂的应用环境...',
    date: '2025-08-03',
    category: 'DevOps',
    categoryId: 3,
    tags: ['Docker', 'Compose', '容器编排'],
    views: 980,
    likes: 100
  },
  {
    id: 14,
    title: 'JavaScript异步编程详解',
    cover: 'https://img0.baidu.com/it/u=700915356,2944457677&fm=253&fmt=auto&app=138&f=PNG?w=554&h=308',
    summary: '异步编程是JavaScript的核心特性之一。本文深入探讨Promise、async/await等异步编程模式的使用方法和最佳实践...',
    date: '2025-08-01',
    category: '编程语言',
    categoryId: 4,
    tags: ['JavaScript', '异步编程', 'Promise'],
    views: 1250,
    likes: 100
  },
  {
    id: 15,
    title: 'React Hooks原理与最佳实践',
    cover: 'https://t14.baidu.com/it/u=2084796258,320088213&fm=224&app=112&f=JPEG?w=500&h=500',
    summary: 'React Hooks彻底改变了React组件的编写方式。本文分析Hooks的工作原理，并分享在实际项目中的使用经验和最佳实践...',
    date: '2025-07-28',
    category: '前端开发',
    categoryId: 1,
    tags: ['React', 'Hooks', '函数式组件'],
    views: 1450,
    likes: 100
  }
])

// 分类列表
export const categories = reactive([
  { id: 1, name: '前端开发', count: 28 },
  { id: 2, name: '后端架构', count: 16 },
  { id: 3, name: 'DevOps', count: 12 },
  { id: 4, name: '编程语言', count: 22 },
  { id: 5, name: '数据库', count: 9 },
  { id: 6, name: '云计算', count: 14 },
  { id: 7, name: '人工智能', count: 7 }
])

// 专栏列表
export const columns = reactive([
  {
    id: 1,
    title: 'Vue3深度解析',
    cover: 'https://img1.baidu.com/it/u=1514230067,94251875&fm=253&fmt=auto&app=138&f=JPEG?w=888&h=500',
    description: '深入探讨Vue3核心原理和高级用法',
    createTime: '2025-01-15',
    articleCount: 12,
    totalViews: 15600
  },
  {
    id: 2,
    title: '前端工程化实践',
    cover: 'https://img2.baidu.com/it/u=3036548280,1019325136&fm=253&fmt=auto&app=120&f=JPEG?w=640&h=325',
    description: '现代前端工程化建设最佳实践',
    createTime: '2025-03-22',
    articleCount: 8,
    totalViews: 9800
  },
  {
    id: 3,
    title: 'TypeScript进阶',
    cover: 'https://img2.baidu.com/it/u=4196342142,1557743768&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=500',
    description: 'TypeScript高级类型和设计模式',
    createTime: '2025-05-10',
    articleCount: 6,
    totalViews: 7500
  }
])

// 标签云
export const tags = reactive([
  { name: 'JavaScript', size: 'large', count: 42 },
  { name: 'Vue', size: 'large', count: 38 },
  { name: 'React', size: 'medium', count: 25 },
  { name: 'TypeScript', size: 'large', count: 30 },
  { name: 'Node.js', size: 'medium', count: 22 },
  { name: 'Docker', size: 'medium', count: 18 },
  { name: 'Kubernetes', size: 'small', count: 12 },
  { name: 'MongoDB', size: 'small', count: 15 },
  { name: 'MySQL', size: 'medium', count: 20 },
  { name: 'Python', size: 'small', count: 16 },
  { name: 'Go', size: 'small', count: 10 },
  { name: 'Redis', size: 'small', count: 8 }
])

// 热门文章
export const popularArticles = reactive([
  {
    id: 1,
    title: 'Vue3 组合式API完全指南',
    cover: 'https://placekitten.com/800/400',
    summary: 'Vue3的组合式API（Composition API）是Vue3最重要的特性之一，它提供了一种更灵活的方式来组织组件逻辑...',
    date: '2025-09-05',
    category: '前端开发',
    categoryId: 1,
    tags: ['Vue3', 'JavaScript', '前端框架'],
    views: 1250,
    likes: 89
  },
  {
    id: 2,
    title: 'TypeScript高级类型系统详解',
    cover: 'https://placekitten.com/801/400',
    summary: 'TypeScript的类型系统非常强大，本文将深入探讨条件类型、映射类型、类型守卫等高级特性的使用方法...',
    date: '2025-09-03',
    category: '编程语言',
    categoryId: 4,
    tags: ['TypeScript', 'JavaScript', '类型系统'],
    views: 980,
    likes: 76
  },
  {
    id: 3,
    title: '微服务架构设计实践',
    cover: 'https://placekitten.com/802/400',
    summary: '本文分享了在大型项目中实施微服务架构的经验，包括服务拆分、通信方式选择、数据一致性保证等关键问题的解决方案...',
    date: '2025-09-01',
    category: '后端架构',
    categoryId: 2,
    tags: ['微服务', '系统设计', '架构'],
    views: 1560,
    likes: 112
  }
])