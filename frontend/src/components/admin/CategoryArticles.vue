<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElButton, ElTable, ElTableColumn, ElPagination } from 'element-plus'

const route = useRoute()
const router = useRouter()

// 从路由参数获取分类ID
const categoryId = computed(() => route.query.categoryid)

// 模拟分类数据
const categoryData = ref({
  id: 1,
  name: '前端开发',
  description: '前端技术相关文章',
  articleCount: 25
})

// 模拟文章数据
const articles = ref([
  {
    id: 1,
    title: 'Vue 3 组合式API入门',
    status: 'published',
    publishTime: '2024-01-20',
    views: 1200,
    likes: 45,
    comments: 12
  },
  {
    id: 2,
    title: 'React Hooks 深度解析',
    status: 'published',
    publishTime: '2024-02-15',
    views: 890,
    likes: 32,
    comments: 8
  },
  {
    id: 3,
    title: 'TypeScript 类型系统进阶',
    status: 'draft',
    publishTime: '2024-03-10',
    views: 0,
    likes: 0,
    comments: 0
  }
])

const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(10)

// 返回分类管理页面
const goBack = () => {
  router.push('/admin/categories')
}

// 加载分类数据
const loadCategoryData = () => {
  // 这里应该根据categoryId从API获取分类数据
  // 模拟数据加载
  loading.value = true
  setTimeout(() => {
    // 根据categoryId设置不同的分类数据
    if (categoryId.value === '1') {
      categoryData.value = {
        id: 1,
        name: '前端开发',
        description: '前端技术相关文章',
        articleCount: 25
      }
    } else if (categoryId.value === '2') {
      categoryData.value = {
        id: 2,
        name: '后端开发',
        description: '后端技术相关文章',
        articleCount: 18
      }
    } else if (categoryId.value === '3') {
      categoryData.value = {
        id: 3,
        name: '数据库',
        description: '数据库技术相关文章',
        articleCount: 12
      }
    } else if (categoryId.value === '4') {
      categoryData.value = {
        id: 4,
        name: 'DevOps',
        description: '运维部署相关文章',
        articleCount: 8
      }
    }
    loading.value = false
  }, 500)
}

// 加载文章列表
const loadArticles = () => {
  // 这里应该根据categoryId从API获取文章列表
  // 模拟数据加载
  loading.value = true
  setTimeout(() => {
    loading.value = false
  }, 500)
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
  loadCategoryData()
  loadArticles()
})
</script>

<template>
  <div class="category-articles-page">
    <!-- 头部标题和返回按钮 -->
    <div class="page-header">
      <div class="header-left">
        <el-button icon="ArrowLeft" @click="goBack" class="back-button">
          返回
        </el-button>
        <h2 class="category-title">{{ categoryData.name }}</h2>
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
        v-if="articles.length > 0"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="articles.length"
        layout="total, prev, pager, next, jumper"
        @current-change="handlePageChange"
        class="pagination"
      />
    </div>
  </div>
</template>

<style scoped>
.category-articles-page {
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

.category-title {
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