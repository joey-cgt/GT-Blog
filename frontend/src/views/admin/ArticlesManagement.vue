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
          :activeTab="activeTab"
          @page-change="handlePublishedPageChange"
          @edit="handleEditArticle"
          @delete="handleDeleteArticle"
          @top="handleTopArticle"
        />
      </el-tab-pane>
      
      <el-tab-pane label="草稿箱" name="drafts">
        <article-list 
          :articles="draftArticles" 
          :loading="loading"
          :pagination="draftPagination"
          :activeTab="activeTab"
          @page-change="handleDraftPageChange"
          @edit="handleEditArticle"
          @delete="handleDeleteArticle"
        />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessageBox, ElMessage } from 'element-plus'
import ArticleList from '../../components/admin/ArticleList.vue'
import { getArticleList, deleteArticle, updateArticleTop } from '../../api/article.js'

const router = useRouter()
const route = useRoute()

// 当前激活的标签页
const activeTab = ref('published')
// 加载状态
const loading = ref(false)

// 文章数据
const publishedArticles = ref([])
const draftArticles = ref([])

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
    const params = {
      page: activeTab.value === 'published' ? publishedPagination.value.currentPage : draftPagination.value.currentPage,
      pageSize: activeTab.value === 'published' ? publishedPagination.value.pageSize : draftPagination.value.pageSize,
      sortBy: 'create_time',
      sortOrder: 'desc'
    }
    
    // 根据标签页设置不同的状态参数
    if (activeTab.value === 'published') {
      params.status = 1 // 已发布文章
    } else if (activeTab.value === 'drafts') {
      params.status = 0 // 草稿文章
    }
    
    const response = await getArticleList(params)
    
    if (activeTab.value === 'published') {
      publishedArticles.value = response.data?.items || []
      publishedPagination.value.total = response.data?.total || 0
    } else {
      console.log('草稿文章:', response.data?.items || [])
      draftArticles.value = response.data?.items || []
      draftPagination.value.total = response.data?.total || 0
    }
  } catch (error) {
    console.error('加载文章失败:', error)
    ElMessage.error('加载文章失败')
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

// 监听标签页变化
watch(activeTab, (newTab) => {
  if (newTab === 'published' || newTab === 'drafts') {
    loadArticles()
  }
})

// 新建文章
const handleCreate = () => {
  router.push('/editor/article/new')
}

// 编辑文章
const handleEditArticle = (article) => {
  router.push(`/editor/article/${article.id}?status=${article.status}`)
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
    await deleteArticle(article.id)
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

// 处理文章置顶
const handleTopArticle = async (article, isTop) => {
  try {
    // 调用置顶API
    await updateArticleTop(article.id, { id: article.id, isTop })
    ElMessage.success(isTop ? '置顶成功' : '取消置顶成功')
    // 重新加载数据
    loadArticles()
  } catch (error) {
    console.error('置顶操作失败:', error)
    ElMessage.error(isTop ? '置顶失败' : '取消置顶失败')
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