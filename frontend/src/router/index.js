import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/visitor/Home.vue'
import Articles from '../views/visitor/Articles.vue'
import Categories from '../views/visitor/Categories.vue'
import Tags from '../views/visitor/Tags.vue'
import About from '../views/visitor/About.vue'
import ArticleContent from '../components/visitor/ArticleContent.vue'
import Columns from '../views/visitor/Columns.vue'
import ColumnArticles from '../views/visitor/ColumnArticles.vue'
import AdminLogin from '../views/admin/Login.vue'
import AdminDashboard from '../views/admin/Dashboard.vue'
import AdminLayout from '../layouts/AdminLayout.vue'
import VisitorLayout from '../layouts/VisitorLayout.vue'
import ArticleEditor from '../views/admin/ArticleEditor.vue'

const routes = [
  // 前台访客路由
  {
    path: '/',
    component: VisitorLayout,
    children: [
      {
        path: '', // 首页
        name: 'Home',
        component: Home,
        meta: {
          showSidebar: true,
          showFooter: true
        }
      },
      {
        path: 'articles', // 文章列表页
        name: 'Articles', 
        component: Articles,
        meta: {
          showSidebar: true
        }
      },
      {
        path: 'columns', // 专栏页面
        name: 'Columns',
        component: Columns,
        meta: {
          showSidebar: true
        }
      },
      {
        path: 'columns/:id',
        name: 'ColumnArticles',
        component: ColumnArticles,
        meta: {
          showSidebar: true
        }
      },
      {
        path: 'about',
        name: 'About',
        component: About,
        meta: {
          showSidebar: true
        }
      },
      {
        path: 'article/:id',
        name: 'ArticleContent',
        component: ArticleContent,
        meta: {
          showSidebar: true
        }
      },
    ]
  },

  
  // 后台管理路由
  {
    path: '/login',
    name: 'AdminLogin',
    component: AdminLogin,
    meta: {
      requiresAuth: false
    }
  },
  {
    path: '/admin',
    component: AdminLayout,
    redirect: '/admin/dashboard',
    children: [
      {
        path: 'dashboard',
        component: AdminDashboard,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'articles',
        component: () => import('../views/admin/ArticlesManagement.vue'),
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'columns',
        component: () => import('../views/admin/ColumnsManagement.vue'),
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'columns/column-articles',
        component: () => import('../components/admin/ColumnArticles.vue'),
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'categories',
        component: () => import('../views/admin/CategoriesManagement.vue'),
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'categories/category-articles',
        component: () => import('../components/admin/CategoryArticles.vue'),
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'tags',
        component: () => import('../views/admin/TagsManagement.vue'),
        meta: {
          requiresAuth: true
        }
      },
    ]
  },

  // 独立文章编辑器路由（不使用AdminLayout）
  {
    path: '/editor/drafts/new',
    name: 'NewArticleEditor',
    component: ArticleEditor,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/editor/drafts/:id',
    name: 'DraftArticleEditor',
    component: ArticleEditor,
    meta: {
      requiresAuth: true
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 全局路由守卫
router.beforeEach((to, from, next) => {
  // 确保每次路由切换时滚动到顶部
  window.scrollTo(0, 0)
  
  // 检查是否需要登录验证
  if (to.meta.requiresAuth) {
    const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true'
    if (!isLoggedIn) {
      // 未登录，跳转到登录页面
      next('/login')
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router