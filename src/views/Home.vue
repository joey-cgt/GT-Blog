<script setup>
import AuthorCard from '../components/AuthorCard.vue'
import SimpleArticleList from '../components/SimpleArticleList.vue'
</script>

<template>
  <div class="home-container">
    <AuthorCard />
    
    <div class="articles-grid">
      <div class="featured-section">
        <SimpleArticleList type="pinned" title="置顶文章" />
      </div>
      
      <div class="main-section">
        <SimpleArticleList type="latest" title="最新文章" :limit="5" />
      </div>
      
      <div class="sidebar-section">
        <SimpleArticleList type="popular" title="热门文章" :limit="5" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.home-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px 20px;
}

.articles-grid {
  display: flex;
  flex-direction: column;
  gap: 30px;
  margin-top: 30px;
}

.featured-section,
.main-section,
.sidebar-section {
  width: 100%;
}
</style>