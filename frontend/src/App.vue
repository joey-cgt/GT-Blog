<script setup>
import { useRoute } from 'vue-router'
import { computed } from 'vue'
import AdminLayout from './layouts/AdminLayout.vue'
import VisitorLayout from './layouts/VisitorLayout.vue'
import Login from './views/admin/Login.vue'

const route = useRoute()

const layout = computed(()=>{
  // 如果是管理后台路由，则使用 AdminLayout，否则使用 VisitorLayout
  // 如果是login，返回Login
  if(route.path === '/login') return Login
  else if(route.path.startsWith('/admin')) return AdminLayout
  else return VisitorLayout
})

</script>

<template>
  <component :is="layout">
  </component>>
</template>

<style>

</style>