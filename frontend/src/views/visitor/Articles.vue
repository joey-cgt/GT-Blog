<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import ArticleList from '../../components/visitor/ArticleList.vue'
import Pagination from '../../components/visitor/Pagination.vue'
import ArticleFilter from '../../components/visitor/ArticleFilter.vue'
import { useRoute, useRouter } from 'vue-router'
import { getCategoryArticles, getTagArticles, getArticleList } from '../../api/article'

const route = useRoute()
const router = useRouter()
const currentPage = ref(1)
const itemsPerPage = 10
const filterType = ref('all') // all, category, tag
const filterValue = ref('')
const filteredArticles = ref([])

const fetchArticles = async () => {
  try {
    if (filterType.value === 'tag') {
      const response = await getTagArticles(filterValue.value)
      filteredArticles.value = response.data.items || []
    } else if (filterType.value === 'category') {
      const response = await getCategoryArticles(filterValue.value)
      filteredArticles.value = response.data.items || []
    } else {
      const params = {
        status: 1,
        page: 1, // 总是从第一页开始
        pageSize: itemsPerPage,
        sortBy: 'create_time',
        sortOrder: 'desc'
      }
      const response = await getArticleList(params)
      filteredArticles.value = response.data.items || []
    }
  } catch (error) {
    console.error('获取文章失败:', error)
    filteredArticles.value = []
  }
}

// 监听路由参数变化
watch(() => route.query, (newQuery) => {
  if (newQuery.tagId) {
    filterType.value = 'tag'
    filterValue.value = newQuery.tagId
  } else if (newQuery.categoryId) {
    filterType.value = 'category'
    filterValue.value = newQuery.categoryId
  } else {
    filterType.value = 'all'
    filterValue.value = ''
  }
  currentPage.value = 1
  // 调用获取文章数据的函数
  fetchArticles()
}, { immediate: true })

const totalPages = computed(() => {
  return Math.ceil(filteredArticles.value.length / itemsPerPage)
})

function handlePageChange(page) {
  currentPage.value = page
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

function handleFilterChange({ type, value }) {
  filterType.value = type
  filterValue.value = value
  currentPage.value = 1

  // 更新URL参数
  if (type === 'tag') {
    router.push({
      path: '/articles',
      query: { tagId: value }
    })
  } else if (type === 'category') {
    router.push({
      path: '/articles',
      query: { categoryId: value }
    })
  } else {
    router.push({
      path: '/articles',
      query: {}
    })
  }
}

// 监听路由变化，处理URL参数
watch(() => route.query, (query) => {
  // 只有当URL参数与当前筛选状态不同时才更新
  if (query.tagId && (filterType.value !== 'tag' || filterValue.value !== query.tagId)) {
    filterType.value = 'tag'
    filterValue.value = query.tagId
    currentPage.value = 1
  } else if (query.categoryId && (filterType.value !== 'category' || filterValue.value !== query.categoryId)) {
    filterType.value = 'category'
    filterValue.value = query.categoryId
    currentPage.value = 1
  } else if (!query.tagId && !query.categoryId && filterType.value !== 'all') {
    filterType.value = 'all'
    filterValue.value = ''
    currentPage.value = 1
  }
}, { immediate: true })
</script>

<template>
  <div class="articles-page">
    <ArticleFilter 
      :filter-type="filterType"
      :filter-value="filterValue"
      @filter-change="handleFilterChange"
    />

    <ArticleList 
      :customArticles="filteredArticles.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage)"
    />
    
    <Pagination
      :current-page="currentPage"
      :total-pages="totalPages"
      @page-change="handlePageChange"
    />
  </div>
</template>

<style scoped>
.articles-page {
  padding: 0 0px 20px;
}

.page-title {
  font-size: 28px;
  margin-bottom: 10px;
  color: #2c3e50;
}

.page-description {
  color: #7f8c8d;
  margin-bottom: 30px;
}
</style>
