import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Articles from '../views/Articles.vue'
import Categories from '../views/Categories.vue'
import Tags from '../views/Tags.vue'
import About from '../views/About.vue'
import ArticleContent from '../views/ArticleContent.vue'

const routes = [
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
    component: () => import('../views/Columns.vue'),
    meta: {
      showSidebar: true
    }
  },
  {
    path: '/columns/:id',
    name: 'ColumnArticles',
    component: () => import('../views/ColumnArticles.vue'),
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
  next()
})

export default router