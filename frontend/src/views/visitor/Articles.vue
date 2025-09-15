<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { articles, tags } from '../../store/blog.js'
import ArticleList from '../../components/visitor/ArticleList.vue'
import Pagination from '../../components/visitor/Pagination.vue'
import ArticleFilter from '../../components/visitor/ArticleFilter.vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const currentPage = ref(1)
const itemsPerPage = 10
const filterType = ref('all') // all, category, tag
const filterValue = ref('')

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
})

// 初始化时检查URL参数
onMounted(() => {
  if (route.query.tagId) {
    filterType.value = 'tag'
    filterValue.value = route.query.tagId
  } else if (route.query.categoryId) {
    filterType.value = 'category'
    filterValue.value = route.query.categoryId
  }
})

// 筛选文章
const displayedArticles = computed(() => {
  let filtered = [...articles]
  
  if (filterType.value === 'category' && filterValue.value) {
    filtered = filtered.filter(article => article.categoryId === Number(filterValue.value))
  } else if (filterType.value === 'tag' && filterValue.value) {
    // 根据tagId找到对应的标签名称
    const tag = tags.find(t => t.id === Number(filterValue.value))
    if (tag) {
      filtered = filtered.filter(article => article.tags.includes(tag.name))
    }
  }
  
  // 按时间从近到远排序
  return filtered.sort((a, b) => new Date(b.date) - new Date(a.date))
})

const totalPages = computed(() => {
  return Math.ceil(displayedArticles.value.length / itemsPerPage)
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
      :customArticles="displayedArticles.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage)"
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
