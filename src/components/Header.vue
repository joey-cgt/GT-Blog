<script setup>
import { ref } from 'vue'
import { blogInfo } from '../store/blog.js'
import { useRoute } from 'vue-router'

// 移动端菜单状态
const mobileMenuOpen = ref(false)
const route = useRoute()
</script>

<template>
  <!-- 导航栏 -->
  <header class="header">
    <div class="container header-container">
      <div class="logo">{{ blogInfo.title }}</div>
      
      <!-- 桌面端导航 -->
      <nav class="desktop-nav">
        <ul>
          <li><router-link to="/" :class="{ active: route.path === '/' }">首页</router-link></li>
          <li><router-link to="/articles" :class="{ active: route.path === '/articles' }">文章</router-link></li>
          <li><router-link to="/columns" :class="{ active: route.path === '/columns' }">专栏</router-link></li>
          <li><router-link to="/categories" :class="{ active: route.path === '/categories' }">分类</router-link></li>
          <li><router-link to="/tags" :class="{ active: route.path === '/tags' }">标签</router-link></li>
          <li><router-link to="/about" :class="{ active: route.path === '/about' }">关于</router-link></li>
        </ul>
      </nav>
      
      <!-- 移动端菜单按钮 -->
      <div class="mobile-menu-button" @click="mobileMenuOpen = !mobileMenuOpen">
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
      </div>
    </div>
    
    <!-- 移动端导航菜单 -->
    <nav class="mobile-nav" :class="{ 'open': mobileMenuOpen }">
      <ul>
        <li><router-link to="/" :class="{ active: route.path === '/' }" @click="mobileMenuOpen = false">首页</router-link></li>
        <li><router-link to="/articles" :class="{ active: route.path === '/articles' }" @click="mobileMenuOpen = false">文章</router-link></li>
        <li><router-link to="/categories" :class="{ active: route.path === '/categories' }" @click="mobileMenuOpen = false">分类</router-link></li>
        <li><router-link to="/tags" :class="{ active: route.path === '/tags' }" @click="mobileMenuOpen = false">标签</router-link></li>
        <li><router-link to="/about" :class="{ active: route.path === '/about' }" @click="mobileMenuOpen = false">关于</router-link></li>
      </ul>
    </nav>
  </header>
</template>

<style scoped>
/* 头部导航栏 */
.header {
  background-color: #fff;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  color: #2c3e50;
}

.desktop-nav ul {
  display: flex;
}

.desktop-nav li {
  margin-left: 30px;
}

.desktop-nav a {
  color: #555;
  font-weight: 500;
  padding: 10px 0;
  position: relative;
}

.desktop-nav a.active,
.desktop-nav a:hover {
  color: #3498db;
}

.desktop-nav a.active::after,
.desktop-nav a:hover::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #3498db;
}

.mobile-menu-button {
  display: none;
  flex-direction: column;
  justify-content: space-between;
  width: 30px;
  height: 20px;
  cursor: pointer;
}

.mobile-menu-button .bar {
  height: 3px;
  width: 100%;
  background-color: #333;
  border-radius: 3px;
}

.mobile-nav {
  display: none;
  background-color: #fff;
  padding: 20px;
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1);
}

.mobile-nav ul {
  display: flex;
  flex-direction: column;
}

.mobile-nav li {
  margin: 10px 0;
}

.mobile-nav a {
  color: #555;
  font-weight: 500;
  display: block;
  padding: 10px 0;
}

.mobile-nav a.active,
.mobile-nav a:hover {
  color: #3498db;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .desktop-nav {
    display: none;
  }
  
  .mobile-menu-button {
    display: flex;
  }
  
  .mobile-nav.open {
    display: block;
  }
}
</style>