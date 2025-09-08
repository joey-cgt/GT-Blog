<script setup>
import Header from './components/Header.vue'
import Sidebar from './components/Sidebar.vue'
import Footer from './components/Footer.vue'
</script>

<template>
  <!-- 导航栏 -->
  <Header />

  <!-- 主要内容区 -->
  <main class="main">
    <div class="container main-container">
      <!-- 内容区 -->
      <div class="content">
        <router-view />
      </div>
      
      <!-- 侧边栏 -->
      <Sidebar />
    </div>
  </main>

  <!-- 页脚 -->
  <Footer />
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

/* 响应式布局 */
@media (max-width: 992px) {
  .main-container {
    grid-template-columns: 1fr;
  }
}
</style>