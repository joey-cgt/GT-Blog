import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Articles from '../views/Articles.vue'
import Categories from '../views/Categories.vue'
import Tags from '../views/Tags.vue'
import About from '../views/About.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/articles',
    name: 'Articles',
    component: Articles
  },
  {
    path: '/columns',
    name: 'Columns',
    component: () => import('../views/Columns.vue')
  },
  {
    path: '/columns/:id',
    name: 'ColumnArticles',
    component: () => import('../views/ColumnArticles.vue'),
    meta: {
      hideSidebar: true,
      hideFooter: true
    }
  },
  {
    path: '/categories',
    name: 'Categories',
    component: Categories
  },
  {
    path: '/tags',
    name: 'Tags',
    component: Tags
  },
  {
    path: '/about',
    name: 'About',
    component: About
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