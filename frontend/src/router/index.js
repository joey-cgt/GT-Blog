import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/visitor/Home.vue'
import Articles from '../views/visitor/Articles.vue'
import Categories from '../views/visitor/Categories.vue'
import Tags from '../views/visitor/Tags.vue'
import About from '../views/visitor/About.vue'
import ArticleContent from '../components/ArticleContent.vue'
import Columns from '../views/visitor/Columns.vue'
import ColumnArticles from '../views/visitor/ColumnArticles.vue'

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