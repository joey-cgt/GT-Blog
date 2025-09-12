<script setup>
import { computed } from 'vue'
import { articles, popularArticles } from '../store/blog.js'
import ArticleCard from './ArticleCard.vue'

const props = defineProps({
  type: {
    type: String,
    default: 'latest',
    validator: (value) => ['pinned', 'latest', 'popular'].includes(value)
  },
  title: {
    type: String,
    required: true
  },
  limit: {
    type: Number,
    default: 5
  }
})

const filteredArticles = computed(() => {
  switch (props.type) {
    case 'pinned':
      return articles.slice(0, 3)
    case 'latest':
      return [...articles].sort((a, b) => new Date(b.date) - new Date(a.date)).slice(0, props.limit)
    case 'popular':
      return popularArticles.slice(0, props.limit)
    default:
      return articles.slice(0, props.limit)
  }
})
</script>

<template>
  <section class="articles-section">
    <div class="section-header">
      <h2 class="section-title">
        <span v-if="type === 'pinned'">📌</span>
        <span v-else-if="type === 'latest'">🕒</span>
        <span v-else-if="type === 'popular'">🔥</span>
        {{ title }}
      </h2>
    </div>
    
    <ArticleCard 
      v-for="article in filteredArticles" 
      :key="article.id" 
      :article="article" 
    />
  </section>
</template>

<style scoped>
.section-header {
  margin-bottom: 16px;
}

.section-title {
  font-size: 20px;
  color: #2c3e50;
}
</style>