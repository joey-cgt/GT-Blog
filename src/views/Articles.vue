<script setup>
import { ref, computed, onMounted } from 'vue'
import { articles } from '../store/blog.js'
import ArticleList from '../components/ArticleList.vue'
import Pagination from '../components/Pagination.vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const currentPage = ref(1)
const itemsPerPage = 10

const totalPages = computed(() => {
  return Math.ceil(articles.length / itemsPerPage)
})

function handlePageChange(page) {
  currentPage.value = page
  window.scrollTo({ top: 0, behavior: 'smooth' })
}
</script>

<template>
  <div class="articles-page">
    <ArticleList 
      type="latest" 
      :limit="itemsPerPage"
      :offset="(currentPage - 1) * itemsPerPage"
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
  max-width: 800px;
  margin: 0 auto;
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
