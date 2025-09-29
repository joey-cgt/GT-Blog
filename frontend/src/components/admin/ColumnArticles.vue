<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElButton, ElTable, ElTableColumn, ElPagination } from 'element-plus'
import axios from 'axios'

const route = useRoute()
const router = useRouter()

// 从路由参数获取专栏ID
const columnId = computed(() => route.query.columnid)
const columnName = computed(() => route.query.name)


// 专栏数据
const columnData = ref({})

// 文章数据
const articles = ref([])
const total = ref(0)
const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(10)

// 返回专栏管理页面
const goBack = () => {
  router.push('/admin/columns')
}

// 加载专栏数据
const loadColumnData = async () => {
  if (!columnId.value) {
    ElMessage.error('专栏ID为空')
    return
  }
  
  loading.value = true
  try {
    const response = await axios.get(`/api/v1/columns/${columnId.value}`)
    
    if (response.data && response.data.data) {
      columnData.value = response.data.data
    } else {
      columnData.value = {}
    }
  } catch (error) {
    ElMessage.error('获取专栏详情失败：' + (error.response?.data?.error || error.message))
    console.error('获取专栏详情失败：', error)
  } finally {
    loading.value = false
  }
}

// 加载文章列表
const loadArticles = async () => {
  if (!columnId.value) {
    ElMessage.error('专栏ID为空')
    return
  }
  
  loading.value = true
  try {
    const response = await axios.get('/api/v1/aggregated', {
      params: {
        type: 'column',
        id: columnId.value,
        page: currentPage.value,
        pageSize: pageSize.value
      }
    })
    
    if (response.data && response.data.data) {
      articles.value = response.data.data.list || []
      total.value = response.data.data.total || 0
    } else {
      articles.value = []
      total.value = 0
    }
  } catch (error) {
    ElMessage.error('获取文章列表失败：' + (error.response?.data?.error || error.message))
    console.error('获取文章列表失败：', error)
  } finally {
    loading.value = false
  }
}

// 编辑文章
const editArticle = (article) => {
  router.push(`/editor/article/${article.id}?status=${article.status}`)
}

// 删除文章
const deleteArticle = async (article) => {
  try {
    const confirmed = confirm(`确定要删除文章"${article.title}"吗？`)
    if (confirmed) {
      // 调用API删除文章
      articles.value = articles.value.filter(a => a.id !== article.id)
      ElMessage.success('文章删除成功')
    }
  } catch (error) {
    ElMessage.error('删除文章失败')
  }
}

// 分页变化
const handlePageChange = (page) => {
  currentPage.value = page
  loadArticles()
}

onMounted(() => {
  loadColumnData()
  loadArticles()
})
</script>

<template>
  <div class="column-articles-page">
    <!-- 头部标题和返回按钮 -->
    <div class="page-header">
      <div class="header-left">
        <el-button icon="ArrowLeft" @click="goBack" class="back-button">
          返回
        </el-button>
        <h2 class="column-title">{{ columnName }}</h2>
      </div>
    </div>

    <!-- 文章列表 -->
    <div class="articles-section">
      <h3 class="section-title">文章列表</h3>
      
      <el-table :data="articles" v-loading="loading" class="articles-table">
        <el-table-column prop="title" label="文章标题" min-width="200" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <span :class="['status-badge', row.status]">
              {{ row.status === 'published' ? '已发布' : '草稿' }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="publishTime" label="发布时间" width="120" />
        <el-table-column prop="views" label="浏览数" width="80" />
        <el-table-column prop="likes" label="点赞数" width="80" />
        <el-table-column prop="comments" label="评论数" width="80" />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click.stop="editArticle(row)">
              编辑
            </el-button>
            <el-button size="small" type="danger" @click.stop="deleteArticle(row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <el-pagination
        v-if="total > 0"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        layout="total, prev, pager, next, jumper"
        @current-change="handlePageChange"
        class="pagination"
      />
    </div>
  </div>
</template>

<style scoped>
.column-articles-page {
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

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-button {
  border: none;
  background: #f5f7fa;
}

.back-button:hover {
  background: #e4e7ed;
}

.column-title {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  color: #303133;
}

.articles-section {
  background: #fff;
  border-radius: 12px;
  padding: 20px;
}

.section-title {
  margin: 0 0 20px 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2d3d;
  text-align: left;
}

.articles-table {
  margin-bottom: 20px;
}

.status-badge {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.status-badge.published {
  background: #e1f3d8;
  color: #67c23a;
}

.status-badge.draft {
  background: #f4f4f5;
  color: #909399;
}

.pagination {
  justify-content: center;
  margin-top: 20px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  
  .meta-info {
    flex-direction: column;
    gap: 8px;
  }
  
  .articles-table {
    font-size: 14px;
  }
}
</style>