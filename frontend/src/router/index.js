import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/visitor/Home.vue'
import Articles from '../views/visitor/Articles.vue'
import Categories from '../views/visitor/Categories.vue'
import Tags from '../views/visitor/Tags.vue'
import About from '../views/visitor/About.vue'
import ArticleContent from '../components/ArticleContent.vue'
import Columns from '../views/visitor/Columns.vue'
import ColumnArticles from '../views/visitor/ColumnArticles.vue'
import AdminLogin from '../views/admin/Login.vue'
import AdminDashboard from '../views/admin/Dashboard.vue'

const routes = [
  // 前台访客路由
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      showFooter: true,
      showSidebar: true
    }
  },
  {
    path: '/articles',
    name: 'Articles',
    component: Articles,
    meta: {
      showSidebar: true
    }
  },
  {
    path: '/columns',
    name: 'Columns',
    component: Columns,
    meta: {
      showSidebar: true
    }
  },
  {
    path: '/columns/:id',
    name: 'ColumnArticles',
    component: ColumnArticles,
    meta: {
      showSidebar: true
    }
  },
  {
    path: '/about',
    name: 'About',
    component: About,
    meta: {
      showSidebar: true
    }
  },
  {
    path: '/article/:id',
    name: 'ArticleContent',
    component: ArticleContent,
    meta: {
      showSidebar: true
    }
  },
  
  // 后台管理路由
  {
    path: '/admin/login',
    name: 'AdminLogin',
    component: AdminLogin,
    meta: {
      showFooter: false,
      showSidebar: false,
      requiresAuth: false
    }
  },
  {
    path: '/admin',
    name: 'AdminDashboard',
    component: AdminDashboard,
    meta: {
      showFooter: false,
      showSidebar: false,
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
      next('/admin/login')
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router