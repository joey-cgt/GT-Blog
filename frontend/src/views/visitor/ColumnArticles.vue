<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getColumnDetail } from '@/api/column'
import { getColumnArticles } from '@/api/article'
import ColumnCard from '../../components/visitor/ColumnCard.vue'
import ArticleList from '../../components/visitor/ArticleList.vue'
import Pagination from '../../components/visitor/Pagination.vue'

const route = useRoute()
const columnId = computed(() => parseInt(route.params.id))
const loading = ref(true)
const currentPage = ref(1)
const articlesPerPage = 10 // 每页显示的文章数量
const allColumnArticles = ref([])
const currentColumn = ref(null)

// 当前页面显示的文章
const columnArticles = computed(() => {
  const start = (currentPage.value - 1) * articlesPerPage
  const end = start + articlesPerPage
  return allColumnArticles.value.slice(start, end)
})

// 总页数
const totalPages = computed(() => {
  return Math.ceil(allColumnArticles.value.length / articlesPerPage)
})

// 页码变化处理函数
const handlePageChange = (page) => {
  currentPage.value = page
  // 滚动到页面顶部
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}

onMounted(async () => {
  try {
    const response = await getColumnDetail(columnId.value)
    currentColumn.value = response.data
    loading.value = false
  } catch (error) {
    console.error('获取专栏详情失败:', error)
  }

  try {
    const response = await getColumnArticles(columnId.value)
    allColumnArticles.value = response.data.items || []
  } catch (error) {
    console.error('获取专栏文章失败:', error)
  }
})
</script>

<template>
  <div class="column-articles-page">
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="spinner"></div>
      <p>加载中...</p>
    </div>
    
    <!-- 专栏不存在 -->
    <div v-else-if="!currentColumn" class="not-found">
      <h2>未找到该专栏</h2>
      <p>您访问的专栏不存在或已被删除</p>
      <router-link to="/columns" class="back-link" >返回专栏列表</router-link>
    </div>
    
    <!-- 专栏内容 -->
    <template v-else>
      <!-- 专栏介绍 -->
      <ColumnCard :column="currentColumn" />
      
      <!-- 文章列表 -->
      <div class="articles-section">
        <div class="section-header">
          <h2 class="section-title">文章列表</h2>
          <router-link to="/columns" class="back-link">返回专栏列表</router-link>
        </div>
        
        <div v-if="columnArticles.length === 0" class="no-articles">
          <p>该专栏暂无文章</p>
        </div>
        
        <div v-else>
          <ArticleList :custom-articles="columnArticles" :limit="columnArticles.length" />
          
          <!-- 分页组件 -->
          <Pagination 
            :current-page="currentPage" 
            :total-pages="totalPages" 
            @page-change="handlePageChange" 
          />
        </div>
      </div>
    </template>
  </div>
</template>

<style scoped>
.column-articles-page {
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
  padding: 0 0 20px;
}

/* 加载状态 */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 50px 0;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(0, 0, 0, 0.1);
  border-radius: 50%;
  border-top-color: #3498db;
  animation: spin 1s ease-in-out infinite;
  margin-bottom: 20px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* 未找到专栏 */
.not-found {
  text-align: center;
  padding: 50px 0;
}

.back-link {
  display: inline-block;
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #3498db;
  color: white;
  border-radius: 4px;
  text-decoration: none;
  transition: background-color 0.3s;
}

.back-link:hover {
  background-color: #2980b9;
}

.section-header .back-link {
  display: inline-block;
  margin-top: 20px;
  padding: 5px 10px;
  background-color: #3498db;
  color: white;
  border-radius: 4px;
  text-decoration: none;
  transition: background-color 0.3s;
}

.section-header .back-link:hover {
  background-color: #2980b9;
}

/* 专栏头部 */
.column-info {
  display: flex;
  gap: 30px;
}

.column-cover {
  flex: 0 0 250px;
}

.column-cover img {
  width: 100%;
  height: 150px;
  object-fit: cover;
  border-radius: 8px;
}

.column-details {
  flex: 1;
}

.column-title {
  font-size: 28px;
  margin: 0 0 15px 0;
  color: #2c3e50;
}

.column-desc {
  font-size: 16px;
  color: #555;
  margin-bottom: 20px;
  line-height: 1.6;
}

.column-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  color: #7f8c8d;
}

/* 文章列表部分 */
.articles-section {
  margin-top: 10px;
}

.section-title {
  font-size: 22px;
  margin: 20px 0;
  /* padding-bottom: 10px; */
  border-bottom: 4px solid #e6e6e6;
  color: #2c3e50;
}
.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.section-header a {
  font-size: 14px;
  margin: 20px 0;
  /* padding-bottom: 10px; */
  border: 1px solid #3498db;
  color:#3498db;
  border-radius: 4px;
  transition: all 0.3s;
  padding: 5px;
}

.section-header a:hover {
  background-color: #3498db;
  color: white;
  cursor: pointer;
}



.no-articles {
  text-align: center;
  padding: 30px;
  background-color: #f9f9f9;
  border-radius: 8px;
  color: #666;
}

.article-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .column-info {
    flex-direction: column;
  }
  
  .column-cover {
    flex: none;
    margin-bottom: 20px;
  }
  
  .article-item {
    flex-direction: column;
  }
  
  .article-cover {
    flex: none;
    margin-bottom: 15px;
  }
}
</style>