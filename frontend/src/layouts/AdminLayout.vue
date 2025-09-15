<template>
  <div class="admin-layout">
    <!-- 顶部工具栏 -->
    <header class="admin-header">
      <div class="header-left">
        <h1 class="logo">GT-Blog 博客管理系统</h1>
      </div>
      <div class="header-right">
        <el-dropdown>
          <span class="user-info">
            <el-avatar :size="32" :src="userAvatar" />
            <span class="username">管理员</span>
            <el-icon><arrow-down /></el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>个人设置</el-dropdown-item>
              <el-dropdown-item @click="handleLogout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </header>

    <div class="admin-main">
      <!-- 侧边菜单 -->
      <aside class="admin-sidebar">
        <el-menu
          :default-active="activeMenu"
          router
          class="sidebar-menu"
          background-color="#304156"
          text-color="#bfcbd9"
          active-text-color="#409eff"
        >
          <el-menu-item index="/admin">
            <el-icon><home-filled /></el-icon>
            <span>控制面板</span>
          </el-menu-item>
          <el-sub-menu index="content">
            <template #title>
              <el-icon><document /></el-icon>
              <span>内容管理</span>
            </template>
            <el-menu-item index="/admin/articles">文章管理</el-menu-item>
            <el-menu-item index="/admin/columns">专栏管理</el-menu-item>
            <el-menu-item index="/admin/categories">分类管理</el-menu-item>
            <el-menu-item index="/admin/tags">标签管理</el-menu-item>
          </el-sub-menu>
          <el-menu-item index="/admin/comments">
            <el-icon><chat-dot-round /></el-icon>
            <span>评论管理</span>
          </el-menu-item>
          <el-menu-item index="/admin/users">
            <el-icon><user /></el-icon>
            <span>用户管理</span>
          </el-menu-item>
          <el-sub-menu index="/admin/settings">
            <template #title>
              <el-icon><setting /></el-icon>
              <span>博客设置</span>
            </template>
            <el-menu-item index="/admin/settings/profile">个人设置</el-menu-item>
            <el-menu-item index="/admin/settings/system">系统管理</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </aside>

      <!-- 主要内容区域 -->
      <main class="admin-content">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  HomeFilled,
  Document,
  ChatDotRound,
  User,
  Setting,
  ArrowDown
} from '@element-plus/icons-vue'

const router = useRouter()
const route = useRoute()

const userAvatar = ref('')
const activeMenu = computed(() => {
  const path = route.path
  
  // 处理 /admin/xxx/xxx 格式的路由
  if (path.startsWith('/admin/')) {
    const parts = path.split('/').filter(part => part)
    // 如果路径格式为 /admin/something/other，则返回 /admin/something
    if (parts.length >= 3) {
      return `/${parts.slice(0, 2).join('/')}`
    }
  }
  
  return path
})

const handleLogout = () => {
  localStorage.removeItem('isLoggedIn')
  localStorage.removeItem('rememberedUsername')
  router.push('/login')
}
</script>

<style scoped>


.admin-layout {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.admin-header {
  height: 60px;
  background: #fff;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
}

.header-left .logo {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: #303133;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 8px 12px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.user-info:hover {
  background: #f5f7fa;
}

.username {
  font-size: 14px;
  color: #606266;
}

.admin-main {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.admin-sidebar {
  width: 220px;
  background: #304156;
  overflow-y: auto;
}

.sidebar-menu {
  border-right: none;
}

.admin-content {
  flex: 1;
  padding: 24px;
  background: #f5f7fa;
  overflow-y: auto;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .admin-sidebar {
    width: 64px;
  }
  
  .header-left .logo {
    font-size: 16px;
  }
  
  .username {
    display: none;
  }
}
</style>