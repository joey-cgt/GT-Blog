<template>
  <div class="admin-dashboard">
    <h1>管理后台首页</h1>
    <p>欢迎来到博客管理系统</p>
    <button @click="handleLogout" class="logout-button">退出登录</button>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const handleLogout = () => {
  localStorage.removeItem('isLoggedIn')
  localStorage.removeItem('rememberedUsername')
  router.push('/admin/login')
}
</script>

<style scoped>
.admin-dashboard {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.logout-button {
  padding: 8px 16px;
  background-color: #ff4d4f;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 20px;
}

.logout-button:hover {
  background-color: #ff7875;
}
</style>