<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getRecommendedArticles } from '../../../api/visitor.js'
import { ElMessage } from 'element-plus'

const router = useRouter()

// 定义props，接收当前文章id
const props = defineProps({
  articleId: {
    type: [String, Number],
    default: null
  }
})

// 推荐文章数据
const recommendedArticles = ref([])
const loading = ref(false)

// 获取推荐文章
const fetchRecommendedArticles = async () => {
  try {
    loading.value = true
    // 调用API获取推荐文章，传递当前文章id
    const response = await getRecommendedArticles({
      limit: 5, // 获取5篇推荐文章
      id: props.articleId // 传递当前文章id
    })
    
    // 根据后端返回的数据结构更新
    if (response.code === 200 && response.data && response.data.items) {
      recommendedArticles.value = response.data.items
    } else {
      // 如果后端返回格式不同，使用response.data作为文章列表
      recommendedArticles.value = response.data || []
    }
  } catch (error) {
    console.error('获取推荐文章失败:', error)
    ElMessage.error('获取推荐文章失败')
  } finally {
    loading.value = false
  }
}

const handleArticleClick = (articleId) => {
  router.push({
    path: `/article/${articleId}`
  })
}

// 组件挂载时获取数据
onMounted(() => {
  fetchRecommendedArticles()
})
</script>

<template>
  <section class="sidebar-section popular-section">
    <h3 class="sidebar-title">推荐文章</h3>
    
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="spinner"></div>
      <p>加载中...</p>
    </div>
    
    <!-- 文章列表 -->
    <ul v-else-if="recommendedArticles.length > 0" class="popular-list">
      <li v-for="article in recommendedArticles" :key="article.id">
        <a @click="handleArticleClick(article.id)">
          {{ article.title }}
          <div class="stats">
            <span class="views">{{ article.viewCount }} 阅读</span>
            <span class="likes">{{ article.likeCount || 0 }} 点赞</span>
          </div>
        </a>
      </li>
    </ul>
    
    <!-- 无数据状态 -->
    <div v-else class="no-articles">
      <p>暂无推荐文章</p>
    </div>
  </section>
</template>

<style scoped>
.sidebar-section {
  background-color: #fff;
  border-radius: 10px;
  padding: 25px;
  margin-bottom: 30px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
}

.sidebar-title {
  font-size: 20px;
  margin-bottom: 20px;
  color: #2c3e50;
  position: relative;
  padding-bottom: 10px;
}

.sidebar-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background-color: #d6d6d6;
}

/* 加载状态样式 */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 0;
  color: #777;
}

.spinner {
  width: 20px;
  height: 20px;
  border: 2px solid #f3f3f3;
  border-top: 2px solid #3498db;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 10px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* 无数据状态样式 */
.no-articles {
  padding: 40px 0;
  text-align: center;
  color: #999;
}

/* 文章列表样式 */
.popular-list li {
  border-bottom: 1px solid #eee;
  padding: 12px 0;
  text-align: left;
  list-style-type: none;
  cursor: pointer;
}

.popular-list li:last-child {
  border-bottom: none;
}

.popular-list a {
  display: flex;
  flex-direction: column;
  color: #555;
}

.popular-list a:hover {
  color: #3498db;
}

.popular-list .stats {
  display: flex;
  gap: 12px;
  margin-top: 5px;
}

.popular-list .views,
.popular-list .likes {
  font-size: 12px;
  color: #777;
}

.popular-list .likes {
  color: #777;
}
</style>