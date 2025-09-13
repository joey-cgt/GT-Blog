<script setup>
import Header from '../components/Header.vue'
import Sidebar from '../components/Sidebar.vue'
import Footer from '../components/Footer.vue'
import { useRoute } from 'vue-router'

const route = useRoute()

</script>

<template>
  <div>
    <!-- 导航栏 -->
    <Header />

    <!-- 主要内容区 -->
    <main class="main">
      <div class="main-container">
        <!-- 内容区 -->
        <div class="content">
          <router-view />
        </div>
        
        <!-- 侧边栏 -->
        <Sidebar v-if="route.meta.showSidebar" />
      </div>
    </main>

    <!-- 页脚 -->
    <Footer v-if="route.meta.showFooter" />
  </div>
</template>

<style>
@import '../assets/styles/global.css';

/* 管理后台布局 */
.admin-layout {
  min-height: 100vh;
  width: 100%;
}

/* 主要内容区布局 */
.main {
  padding: 40px 0;
  margin: 0 auto;
}

.main-container {
  max-width: 1400px;
  margin: 0 auto;
  display: grid;
  grid-template-columns:3fr 1fr;
  gap: 30px;
}

.content {
  display: flex;
  flex-direction: column;
  max-width: 100%;
}



/* 响应式布局 */
@media (max-width: 992px) {
  .main-container {
    grid-template-columns: 1fr;
  }
}
</style>