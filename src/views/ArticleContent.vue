<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { articles } from '../store/blog.js'

const route = useRoute()
const articleId = computed(() => parseInt(route.params.id))
const article = computed(() => {
  return articles.find(a => a.id === articleId.value) || null
})

// 相关文章推荐（同类别或同标签的其他文章）
const relatedArticles = computed(() => {
  if (!article.value) return []
  
  return articles
    .filter(a => {
      // 排除当前文章
      if (a.id === articleId.value) return false
      
      // 同类别或有相同标签的文章
      const sameCategory = a.category === article.value.category
      const hasCommonTag = a.tags.some(tag => article.value.tags.includes(tag))
      
      return sameCategory || hasCommonTag
    })
    .slice(0, 3) // 最多显示3篇相关文章
})

// 模拟文章内容（实际项目中可能从API获取）
const content = ref('')

onMounted(() => {
  // 模拟文章内容（实际项目中应该从API获取）
  if (article.value) {
    content.value = generateDummyContent(article.value.title)
  }
})

// 生成模拟内容的辅助函数
function generateDummyContent(title) {
  return `
  <h1>${title}</h1>
  
  <p>这是一篇关于${title}的详细文章。在实际应用中，这里应该是从后端API获取的文章详细内容。</p>
  
  <h2>第一部分</h2>
  <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam euismod, nisl eget aliquam ultricies, 
  nunc nisl aliquet nunc, vitae aliquam nisl nunc vitae nisl. Nullam euismod, nisl eget aliquam ultricies, 
  nunc nisl aliquet nunc, vitae aliquam nisl nunc vitae nisl.</p>
  
  <h2>第二部分</h2>
  <p>Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
  quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor 
  in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.</p>
  
  <h2>第三部分</h2>
  <p>Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
  Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, 
  totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.</p>
  
  <h2>总结</h2>
  <p>Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni 
  dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor 
  sit amet, consectetur, adipisci velit.</p>
  `
}
</script>

<template>
  <div v-if="article" class="article-content-page">
    <!-- 文章头部信息 -->
    <div class="article-header">
      <div class="article-cover">
        <img :src="article.cover" :alt="article.title" />
      </div>
      
      <h1 class="article-title">{{ article.title }}</h1>
      
      <div class="article-meta">
        <span class="meta-item date">{{ article.date }}</span>
        <span class="meta-item category">{{ article.category }}</span>
        <span class="meta-item views">{{ article.views }} 阅读</span>
      </div>
      
      <div class="article-tags">
        <span v-for="tag in article.tags" :key="tag" class="tag">{{ tag }}</span>
      </div>
    </div>
    
    <!-- 文章内容 -->
    <div class="article-body" v-html="content"></div>
    
    <!-- 相关文章推荐 -->
    <div v-if="relatedArticles.length > 0" class="related-articles">
      <h3 class="related-title">相关文章</h3>
      <div class="related-list">
        <div v-for="relatedArticle in relatedArticles" :key="relatedArticle.id" class="related-item">
          <router-link :to="`/article/${relatedArticle.id}`" class="related-link">
            <div class="related-cover">
              <img :src="relatedArticle.cover" :alt="relatedArticle.title" />
            </div>
            <div class="related-info">
              <h4 class="related-article-title">{{ relatedArticle.title }}</h4>
              <p class="related-article-date">{{ relatedArticle.date }}</p>
            </div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
  
  <div v-else class="article-not-found">
    <h2>文章不存在</h2>
    <p>抱歉，您请求的文章不存在或已被删除。</p>
    <router-link to="/articles" class="back-link">返回文章列表</router-link>
  </div>
</template>

<style scoped>
.article-content-page {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  padding: 30px;
  margin-bottom: 30px;
}

.article-header {
  margin-bottom: 30px;
  text-align: left;
}

.article-cover {
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 20px;
}

.article-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.article-title {
  font-size: 32px;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 16px;
  line-height: 1.3;
}

.article-meta {
  display: flex;
  gap: 16px;
  margin-bottom: 16px;
  color: #7f8c8d;
  font-size: 14px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.article-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 20px;
}

.tag {
  background-color: #f0f0f0;
  color: #555;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
}

.article-body {
  line-height: 1.8;
  color: #333;
  font-size: 16px;
  text-align: left;
}

.article-body h1,
.article-body h2,
.article-body h3 {
  margin-top: 30px;
  margin-bottom: 15px;
  color: #2c3e50;
}

.article-body p {
  margin-bottom: 20px;
}

.related-articles {
  margin-top: 50px;
  border-top: 1px solid #eee;
  padding-top: 30px;
}

.related-title {
  font-size: 22px;
  margin-bottom: 20px;
  color: #2c3e50;
  text-align: left;
}

.related-list {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.related-item {
  border-radius: 8px;
  overflow: hidden;
  transition: transform 0.3s;
}

.related-item:hover {
  transform: translateY(-5px);
}

.related-link {
  text-decoration: none;
  color: inherit;
  display: block;
}

.related-cover {
  height: 150px;
  overflow: hidden;
}

.related-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s;
}

.related-item:hover .related-cover img {
  transform: scale(1.05);
}

.related-info {
  padding: 12px;
  background-color: #f9f9f9;
  text-align: left;
}

.related-article-title {
  font-size: 16px;
  margin-bottom: 8px;
  color: #2c3e50;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.related-article-date {
  font-size: 12px;
  color: #7f8c8d;
}

.article-not-found {
  text-align: center;
  padding: 50px 20px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.back-link {
  display: inline-block;
  margin-top: 20px;
  color: #3498db;
  text-decoration: none;
}

.back-link:hover {
  text-decoration: underline;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .article-content-page {
    padding: 20px;
  }
  
  .article-cover {
    height: 200px;
  }
  
  .article-title {
    font-size: 24px;
  }
  
  .related-list {
    grid-template-columns: 1fr;
  }
}
</style>