<script setup>
import {ref, onMounted} from 'vue'
import { getAuthor } from '../../../api/visitor';
import { getBlogStatistics } from '../../../api/blogstats';

const authorInfo = ref({
  name: '',
  avatarUrl: '',
  email: '',
  wechat: ''
})

const blogStats = ref({
  totalArticles: 0,
  totalViews: 0,
  totalLikes: 0
})

onMounted(async () => {
  try {
    const res = await getAuthor(1)
    if (res.data) {
      authorInfo.value.name = res.data.nickname
      authorInfo.value.avatarUrl = res.data.avatarUrl
      authorInfo.value.email = res.data.email
      authorInfo.value.wechat = res.data.wechat
    }
  } catch (error) {
    console.error('获取作者信息失败:', error)
  }

  try {
    const response = await getBlogStatistics()
    if (response.data) {
      blogStats.value.totalArticles = response.data.totalArticles || 0
      blogStats.value.totalViews = response.data.totalViews || 0
      blogStats.value.totalLikes = response.data.totalLikes || 0
    }
  } catch (error) {
    console.error('获取博客统计数据失败:', error)
  }
})

</script>

<template>
  <div class="blog-info-card">
    <!-- 博主信息 -->
    <div class="author-section">
      <img :src="authorInfo.avatarUrl" alt="博主头像" class="author-avatar">
      <div class="author-info">
        <h3 class="author-name">{{ authorInfo.name }}</h3>
        <div class="email-tooltip">
          <i class="fas fa-envelope email-icon"></i>
          <span class="tooltip-text">{{ authorInfo.email }}</span>
        </div>
        <div class="wechat-tooltip">
          <i class="fab fa-weixin wechat-icon"></i>
          <span class="tooltip-text">{{ authorInfo.wechat }}</span>
        </div>
      </div>
    </div>
    
    <!-- 博客统计 -->    
    <div class="stats-section">
    <div class="stat-item">
        <i class="fas fa-file-alt"></i>
        <span class="stat-label">总文章数</span>
        <span class="stat-value">{{ blogStats.totalArticles }}</span>
    </div>
    
    <div class="stat-item">
        <i class="fas fa-eye"></i>
        <span class="stat-label">总浏览量</span>
        <span class="stat-value">{{ blogStats.totalViews }}</span>
    </div>
    
    <div class="stat-item">
        <i class="fas fa-heart"></i>
        <span class="stat-label">总点赞数</span>
        <span class="stat-value">{{ blogStats.totalLikes }}</span>
    </div>
    </div>
    </div>
</template>

<style scoped>
.blog-info-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
}

.author-section {
  text-align: center;
  margin-bottom: 0px;
  padding-bottom: 20px;
  /* border-bottom: 1px solid #eee; */
}

.author-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 12px;
}

.author-info {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.author-name {
  font-size: 18px;
  font-weight: 600;
  color: white;
  margin: 0;
}

.email-tooltip {
  position: relative;
  display: inline-block;
}

.email-icon {
  color: white;
  cursor: pointer;
  transition: opacity 0.3s;
}

.email-icon:hover {
  opacity: 0.8;
}

.tooltip-text {
  visibility: hidden;
  background-color: #333;
  color: white;
  text-align: center;
  border-radius: 4px;
  padding: 5px 8px;
  position: absolute;
  z-index: 1;
  bottom: 125%;
  left: 50%;
  transform: translateX(-50%);
  white-space: nowrap;
  opacity: 0;
  transition: opacity 0.3s;
}

.email-tooltip:hover .tooltip-text,
.wechat-tooltip:hover .tooltip-text {
  visibility: visible;
  opacity: 1;
}

.wechat-tooltip {
  position: relative;
  display: inline-block;
}

.wechat-icon {
  color: white;
  cursor: pointer;
  transition: opacity 0.3s;
}

.wechat-icon:hover {
  opacity: 0.8;
}
.stats-section {
  display: flex;
  flex-direction: column;
  gap: 0px;
}

.stat-item {
  display: flex;
  align-items: center;
  /* justify-content: space-between; */
  gap: 20px;
  padding: 2px 0;
}


.stat-label {
  display: flex;
  align-items: center;
  gap: 8px;
  color: white;
  font-size: 14px;
}

.stat-value {
  /* font-weight: 600; */
  color: #ffffff;
  font-size: 16px;
}

.fas {
  color: white;
  width: 16px;
}
</style>