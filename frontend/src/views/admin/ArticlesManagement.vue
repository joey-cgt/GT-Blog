<template>
  <div class="articles-management">
    <!-- 页面标题和操作按钮 -->
    <div class="page-header">
      <h2 class="page-title">文章管理</h2>
      <div class="action-buttons">
        <el-button type="primary" icon="Plus" @click="handleCreate">新建文章</el-button>
      </div>
    </div>

    <!-- 标签页切换 -->
    <el-tabs v-model="activeTab" class="articles-tabs">
      <el-tab-pane label="已发布文章" name="published">
        <article-list 
          :articles="publishedArticles" 
          :loading="loading"
          :pagination="publishedPagination"
          @page-change="handlePublishedPageChange"
          @edit="handleEditArticle"
          @delete="handleDeleteArticle"
          @change-status="handleChangeStatus"
        />
      </el-tab-pane>
      
      <el-tab-pane label="草稿箱" name="drafts">
        <article-list 
          :articles="draftArticles" 
          :loading="loading"
          :pagination="draftPagination"
          @page-change="handleDraftPageChange"
          @edit="handleEditArticle"
          @delete="handleDeleteArticle"
          @change-status="handleChangeStatus"
        />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import ArticleList from '../../components/admin/ArticleList.vue'

const router = useRouter()
const route = useRoute()

// 当前激活的标签页
const activeTab = ref('published')
// 加载状态
const loading = ref(false)

// 模拟数据 - 已发布文章
const publishedArticles = ref([
  {
    id: 1,
    title: 'Vue 3 组合式API最佳实践',
    category: '前端开发',
    tags: ['Vue', 'JavaScript'],
    publishTime: '2024-01-15 10:30',
    views: 2456,
    likes: 324,
    comments: 89,
    status: 'published'
  },
  {
    id: 2,
    title: 'TypeScript 高级类型技巧',
    category: '前端开发',
    tags: ['TypeScript', 'JavaScript'],
    publishTime: '2024-01-12 14:20',
    views: 1987,
    likes: 287,
    comments: 67,
    status: 'published'
  },
  {
    id: 3,
    title: 'Element Plus 组件库深度解析',
    category: 'UI框架',
    tags: ['Element Plus', 'Vue'],
    publishTime: '2024-01-10 09:15',
    views: 1765,
    likes: 234,
    comments: 45,
    status: 'published'
  }
])

// 模拟数据 - 草稿文章
const draftArticles = ref([
  {
    id: 4,
    title: '前端性能优化完全指南',
    category: '性能优化',
    tags: ['性能', '优化'],
    createTime: '2024-01-08 16:45',
    status: 'draft'
  },
  {
    id: 5,
    title: 'CSS Grid 布局实战教程',
    category: 'CSS',
    tags: ['CSS', '布局'],
    createTime: '2024-01-05 11:20',
    status: 'draft'
  },
  {
    id: 6,
    title: 'JavaScript 异步编程详解',
    category: 'JavaScript',
    tags: ['JavaScript', '异步'],
    createTime: '2024-01-18 08:30',
    status: 'draft'
  }
])

// 分页配置
const publishedPagination = ref({
  currentPage: 1,
  pageSize: 10,
  total: 35
})

const draftPagination = ref({
  currentPage: 1,
  pageSize: 10,
  total: 12
})

// 页面加载时获取数据
onMounted(() => {
  // 根据查询参数设置激活的标签页
  if (route.query.tab === 'drafts') {
    activeTab.value = 'drafts'
  }
  loadArticles()
})

// 加载文章数据
const loadArticles = async () => {
  loading.value = true
  try {
    // 这里应该是API调用，暂时用模拟数据
    await new Promise(resolve => setTimeout(resolve, 500))
  } catch (error) {
    console.error('加载文章失败:', error)
  } finally {
    loading.value = false
  }
}

// 处理分页变化
const handlePublishedPageChange = (page) => {
  publishedPagination.value.currentPage = page
  loadArticles()
}

const handleDraftPageChange = (page) => {
  draftPagination.value.currentPage = page
  loadArticles()
}

// 新建文章
const handleCreate = () => {
  router.push('/editor/drafts/new')
}

// 编辑文章
const handleEditArticle = (article) => {
  router.push(`/editor/drafts/${article.id}?status=${article.status}`)
}

// 删除文章
const handleDeleteArticle = async (article) => {
  try {
    // 确认删除
    await ElMessageBox.confirm(
      `确定要删除文章"${article.title}"吗？此操作不可恢复。`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    // 调用删除API
    console.log('删除文章:', article.id)
    ElMessage.success('删除成功')
    
    // 重新加载数据
    loadArticles()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

// 更改文章状态
const handleChangeStatus = async (article, newStatus) => {
  try {
    console.log(`更改文章状态: ${article.id} -> ${newStatus}`)
    // 调用API更改状态
    ElMessage.success('状态更新成功')
    loadArticles()
  } catch (error) {
    ElMessage.error('状态更新失败')
  }
}
</script>

<style scoped>
.articles-management {
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
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

.articles-tabs {
  background: #fff;
  border-radius: 8px;
  padding: 20px;
}

:deep(.el-tabs__header) {
  margin: 0 0 20px 0;
}

:deep(.el-tabs__nav-wrap::after) {
  background-color: #f0f0f0;
}
</style>