<script setup>
import Header from './components/Header.vue'
import Sidebar from './components/Sidebar.vue'
import Footer from './components/Footer.vue'
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <!-- 导航栏 -->
  <Header />

  <!-- 主要内容区 -->
  <main class="main">
    <div class="container main-container">
      <!-- 内容区 -->
      <div class="content" :class="{ 'full-width': route.meta.hideSidebar }">
        <router-view />
      </div>
      
      <!-- 侧边栏 -->
      <Sidebar v-if="!route.meta.hideSidebar" />
    </div>
  </main>

  <!-- 页脚 -->
  <Footer v-if="!route.meta.hideFooter" />
</template>

<style>
@import './assets/styles/global.css';

/* 主要内容区布局 */
.main {
  padding: 40px 0;
}

.main-container {
  display: grid;
  grid-template-columns: 5fr 2fr;
  gap: 30px;
}

.content {
  display: flex;
  flex-direction: column;
}

.content.full-width {
  grid-column: 1 / -1;
}

/* 响应式布局 */
@media (max-width: 992px) {
  .main-container {
    grid-template-columns: 1fr;
  }
}
</style>