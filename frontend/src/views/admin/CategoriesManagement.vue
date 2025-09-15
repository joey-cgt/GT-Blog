<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

// 模拟分类数据
const categories = ref([
  { id: 1, name: '前端开发', description: '前端技术相关文章', articleCount: 25 },
  { id: 2, name: '后端开发', description: '后端技术相关文章', articleCount: 18 },
  { id: 3, name: '数据库', description: '数据库技术相关文章', articleCount: 12 },
  { id: 4, name: 'DevOps', description: '运维部署相关文章', articleCount: 8 }
])

const loading = ref(false)

// 加载分类数据
const loadCategories = async () => {
  loading.value = true
  try {
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 500))
  } catch (error) {
    ElMessage.error('加载分类失败')
  } finally {
    loading.value = false
  }
}

// 初始化加载
loadCategories()
</script>

<template>
  <div class="categories-management">
    <div class="page-header">
      <h2 class="page-title">分类管理</h2>
    </div>
    
    <div class="content-section">
      <p>分类管理功能开发中...</p>
      <el-table :data="categories" v-loading="loading">
        <el-table-column prop="name" label="分类名称" width="120" />
        <el-table-column prop="description" label="描述" />
        <el-table-column prop="articleCount" label="文章数量" width="100" />
      </el-table>
    </div>
  </div>
</template>

<style scoped>
.categories-management {
  padding: 20px;
}

.page-header {
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e8e8e8;
}

.page-title {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  color: #303133;
}

.content-section {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
}
</style>
