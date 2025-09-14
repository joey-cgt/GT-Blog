<script setup>
import { useRouter } from 'vue-router'

const props = defineProps({
  article: {
    type: Object,
    required: true
  }
})

const router = useRouter()

function navigateToArticle(articleId) {
  router.push(`/article/${articleId}`)
}
</script>

<template>
  <article class="article-card">
    <div class="article-content">
      <div class="article-header">
        <div class="article-cover">
          <img :src="article.cover" :alt="article.title" />
        </div>
        <div class="article-info">
          <a @click.prevent="navigateToArticle(article.id)" href="#" class="article-title-link">
            <h3 class="article-title">{{ article.title }}</h3>
          </a>
          <p class="article-summary">{{ article.summary }}</p>
        </div>
      </div>
      <div class="article-footer">
        <div class="article-meta">
          <span class="date">{{ article.date }}</span>
          <span class="category">{{ article.category }}</span>
          <span class="views">{{ article.views }} 阅读</span>
        </div>
        <div class="article-tags">
          <span v-for="tag in article.tags" :key="tag" class="tag">{{ tag }}</span>
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped>
.article-card {
  background-color: #fff;
  overflow: hidden;
  margin-bottom: 10px;
  border-radius: 8px;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.05);
}

.article-content {
  padding: 16px;
}

.article-header {
  display: flex;
  gap: 16px;
  margin-bottom: 12px;
}

.article-cover {
  flex-shrink: 0;
  width: 120px;
  height: 80px;
  border-radius: 6px;
  overflow: hidden;
}

.article-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s;
}

.article-card:hover .article-cover img {
  transform: scale(1.05);
}

.article-info {
  flex: 1;
  min-width: 0;
  text-align: left;
}

.article-title-link {
  text-decoration: none;
  color: inherit;
}

.article-title {
  font-size: 18px;
  margin-bottom: 8px;
  color: #2c3e50;
  line-height: 1.4;
  display: -webkit-box;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  transition: color 0.3s;
}

.article-title:hover {
  color: #3498db;
}

.article-summary {
  color: #666;
  font-size: 14px;
  line-height: 1.5;
  display: -webkit-box;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 12px;
  border-top: 1px solid #f0f0f0;
}

.article-meta {
  display: flex;
  gap: 12px;
  font-size: 12px;
  color: #888;
}

.article-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.tag {
  background-color: #f0f0f0;
  color: #555;
  padding: 3px 8px;
  border-radius: 3px;
  font-size: 11px;
}



/* 响应式设计 */
@media (max-width: 768px) {
  .article-header {
    flex-direction: column;
    gap: 12px;
  }
  
  .article-cover {
    width: 100%;
    height: 160px;
  }
  
  .article-footer {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }
  
  .article-meta {
    flex-wrap: wrap;
  }
}
</style>