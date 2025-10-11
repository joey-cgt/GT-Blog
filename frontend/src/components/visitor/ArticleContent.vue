<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { getArticleDetail, incrementArticleLike, decrementArticleLike, recordArticleView } from '../../api/article.js'
import Comments from './Comments.vue'
import MarkdownRenderer from '../common/MarkdownRenderer.vue'

const route = useRoute()
const articleId = computed(() => parseInt(route.params.id))
const article = ref(null)
const content = ref('')
const isLiked = ref(false)
const likeCount = ref(0)

const fetchArticleDetail = async () => {
  try {
    const response = await getArticleDetail(articleId.value)
    article.value = response.data
    content.value = response.data.content
    likeCount.value = response.data.likeCount || 0
    
    // 记录文章浏览量，注意这里不等待结果返回，不影响用户体验
    recordArticleView(articleId.value).catch(err => {
      console.warn('记录浏览量失败:', err)
      // 浏览量记录失败不影响用户体验，仅记录日志
    })
  } catch (error) {
    console.error('获取文章详情失败:', error)
  }
}


const handleLike = async() => {
  try {
    if (isLiked.value) {
      // 取消点赞：图标变暗，点赞数-1
      await decrementArticleLike(articleId.value)
      likeCount.value--
      isLiked.value = false
      console.log('取消点赞成功', likeCount.value)
    } else {
      // 点赞：图标亮起，点赞数+1
      await incrementArticleLike(articleId.value)
      likeCount.value++
      isLiked.value = true
      console.log('点赞成功', likeCount.value)
    }
    // 更新文章对象中的点赞数
    if (article.value) {
      article.value.likeCount = likeCount.value
    }
  } catch (error) {
    console.error('点赞操作失败:', error)
    // 如果API调用失败，恢复状态
    if (isLiked.value) {
      likeCount.value++
    } else {
      likeCount.value--
    }
    isLiked.value = !isLiked.value
  }
}

// 滚动到评论区
const scrollToComments = () => {
  const commentsSection = document.getElementById('comments')
  if (commentsSection) {
    commentsSection.scrollIntoView({ behavior: 'smooth' })
  } else {
    console.log('评论区未找到')
  }
}

onMounted(() => {
  fetchArticleDetail()
})


</script>

<template>
  <div v-if="article" class="article-content-page">
    <div class="article-content">
      <!-- 文章头部信息 -->
      <div class="article-header">
        <!-- <div class="article-cover">
          <img :src="article.cover" :alt="article.title" />
        </div> -->
        
        <h1 class="article-title">{{ article.title }}</h1>
        
        <div class="article-meta">
          <span class="meta-item date">{{ article.date }}</span>
          <span class="meta-item category">{{ article.category.categoryName }}</span>
          <span class="meta-item views">{{ article.viewCount }} 阅读</span>
          <span class="meta-item likes">{{ article.likeCount }} 赞</span>
        </div>
        
        <div class="article-tags">
          <span v-for="tag in article.tags" :key="tag" class="tag">{{ tag.tagName }}</span>
        </div>
      </div>
      
      <!-- 文章内容 -->
      <MarkdownRenderer :content="content" />
    </div>
    <!-- 底部粘性操作栏 -->
    <div class="sticky-action-bar">
  <button class="action-btn like-btn" @click="handleLike" :class="{ 'liked': isLiked }">
    <span class="icon">{{ isLiked ? '👍' : '👍' }}</span>
    <span class="text">{{ isLiked ? '已点赞' : '点赞' }}</span>
    <span class="count">{{ likeCount }}</span>
  </button>
  <button class="action-btn comment-btn" @click="scrollToComments">
    <span class="icon">💬</span>
    <span class="text">评论</span>
  </button>
</div>
    
    <!-- 评论区组件 -->
    <Comments :articleId="articleId" />
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
  /* padding: 30px; */
  margin-bottom: 30px;
}

.article-content {
  margin: 30px 30px 0;
}

.article-header {
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

/* GitHub Markdown 样式由 github-markdown-css 提供 */
.article-body {
  /* 保留必要的容器样式 */
  line-height: 1.8;
  font-size: 16px;
  text-align: left;
}

/* KaTeX 样式调整 */
.article-body .katex-display {
  margin: 1.5em 0;
  overflow-x: auto;
  overflow-y: hidden;
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

/* 底部粘性操作栏 */
.sticky-action-bar {
  position: sticky;
  bottom: 0px;
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  margin-top: 40px;
  padding: 8px 12px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  /* border-radius: 50px; */
  /* box-shadow: 0 1px 12px rgba(0, 0, 0, 0.15); */
  border: 1px solid rgba(255, 255, 255, 0.2);
  z-index: 20;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  border: none;
  border-radius: 25px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  background: #f8f9fa;
  color: #555;
}

.action-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.like-btn:hover {
  background: #e8f5e8;
  color: #27ae60;
}

.like-btn.liked {
  background: #e8f5e8;
  color: #27ae60;
}

.comment-btn:hover {
  background: #e8f4fd;
  color: #3498db;
}

.action-btn .icon {
  font-size: 16px;
}

.action-btn .count {
  background: rgba(0, 0, 0, 0.1);
  padding: 2px 6px;
  border-radius: 10px;
  font-size: 12px;
  min-width: 18px;
  text-align: center;
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

  .sticky-action-bar {
    bottom: 10px;
    padding: 10px 16px;
  }

  .action-btn {
    padding: 8px 16px;
    font-size: 13px;
  }

  .action-btn .icon {
    font-size: 14px;
  }
}

</style>