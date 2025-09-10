<script setup>
import { ref, computed } from 'vue'
import { articles } from '../store/blog.js'
import ArticleList from '../components/ArticleList.vue'
import Pagination from '../components/Pagination.vue'
import ArticleFilter from '../components/ArticleFilter.vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const currentPage = ref(1)
const itemsPerPage = 10
const filterType = ref('all') // all, category, tag
const filterValue = ref('')

// 筛选文章
const displayedArticles = computed(() => {
  let filtered = [...articles]
  
  if (filterType.value === 'category' && filterValue.value) {
    filtered = filtered.filter(article => article.category === filterValue.value)
  } else if (filterType.value === 'tag' && filterValue.value) {
    filtered = filtered.filter(article => article.tags.includes(filterValue.value))
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
}
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
  padding: 0 20px 20px;
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
