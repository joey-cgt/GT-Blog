<script setup>
import { computed } from 'vue'
import { articles } from '../store/blog.js'


const props = defineProps({
  limit: {
    type: Number,
    default: 10
  },
  offset: {
    type: Number,
    default: 0
  },
  customArticles: {
    type: Array,
    default: null
  }
})

const displayedArticles = computed(() => {
  // 如果提供了自定义文章列表，则使用它
  const sourceArticles = props.customArticles || articles
  
  return [...sourceArticles]
    .sort((a, b) => new Date(b.date) - new Date(a.date))
    .slice(props.offset, props.offset + props.limit)
})
</script>

<template>
  <div class="articles-list">
    <div 
      v-for="article in displayedArticles" 
      :key="article.id" 
      class="article-item"
    >
      <div class="article-content">
        <img 
          :src="article.cover" 
          :alt="article.title"
          class="article-image"
        >
        <div class="article-details">
          <h3 class="article-title">{{ article.title }}</h3>
          <p class="article-summary">{{ article.summary }}</p>
        </div>
      </div>
      <div class="article-meta">
        <span class="meta-item">{{ article.date }}</span>
        <span class="meta-item">{{ article.category }}</span>
        <span class="meta-item">{{ article.views }} 阅读</span>
        <div class="tags-container">
          <span 
            v-for="tag in article.tags" 
            :key="tag" 
            class="tag"
          >
            {{ tag }}
          </span>
        </div>
      </div>
    </div>

    <slot name="footer"></slot>
  </div>
</template>

<style scoped>
.articles-list {
  display: flex;
  flex-direction: column;
  gap: 0px;
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.article-item {
  padding: 12px 0;
  border-bottom: 1px solid #dadada;
}

.article-item:last-child {
  border-bottom: none;
}

.article-content {
  display: flex;
  gap: 20px;
}

.article-image {
  width: 120px;
  height: 80px;
  object-fit: cover;
  border-radius: 4px;
}

.article-details {
  flex: 1;
  text-align: left;
}

.article-title {
  font-size: 18px;
  margin: 0 0 8px 0;
  color: #2c3e50;
}

.article-summary {
  color: #666;
  font-size: 14px;
  margin: 0 0 12px 0;
  line-height: 1.5;
}

.article-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  font-size: 12px;
  color: #7f8c8d;
  margin-top: 8px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.tags-container {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  justify-content: flex-end;
}

.tag {
  background-color: #ffffff;
  color: #8d8d8d;
  padding: 2px 4px;
  border-radius: 4px;
  font-size: 12px;
  border: 1px solid #d6d6d6;
}

.view-all {
  color: #3498db;
  font-size: 14px;
  text-decoration: none;
  transition: color 0.3s;
}

.view-all:hover {
  color: #2980b9;
  text-decoration: underline;
}

.view-all-container {
  margin-top: 16px;
  text-align: right;
}
</style>