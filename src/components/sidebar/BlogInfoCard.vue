<script setup>
// 这里需要从API获取博客统计信息
// 暂时使用模拟数据
import { blogInfo } from '../../store/blog.js'

const blogStats = {
  totalArticles: 156,
  totalViews: 28450,
  totalLikes: 8920,
  author: {
    name: '技术博主',
    avatar: '/src/assets/images/avatar.jpg'
  }
}
</script>

<template>
  <div class="blog-info-card">
    <!-- 博主信息 -->
    <div class="author-section">
      <img :src="blogInfo.avatar" alt="博主头像" class="author-avatar">
      <h3 class="author-name">{{ blogInfo.author }}</h3>
    </div>
    
    <!-- 博客统计 -->
    <div class="stats-section">
      <div class="stat-item">
        <i class="fas fa-file-alt"></i>
        <span class="stat-label">文章总数</span>
        <span class="stat-value">{{ blogStats.totalArticles }}</span>
      </div>
      
      <div class="stat-item">
        <i class="fas fa-eye"></i>
        <span class="stat-label">总浏览量</span>
        <span class="stat-value">{{ blogStats.totalViews.toLocaleString() }}</span>
      </div>
      
      <div class="stat-item">
        <i class="fas fa-heart"></i>
        <span class="stat-label">总点赞数</span>
        <span class="stat-value">{{ blogStats.totalLikes.toLocaleString() }}</span>
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
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #eee;
}

.author-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 12px;
}

.author-name {
  font-size: 18px;
  font-weight: 600;
  color: white;
  margin: 0;
}

.stats-section {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.stat-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 0;
}

.stat-item:not(:last-child) {
  border-bottom: 1px solid #f0f0f0;
}

.stat-label {
  display: flex;
  align-items: center;
  gap: 8px;
  color: white;
  font-size: 14px;
}

.stat-value {
  font-weight: 600;
  color: #3498db;
  font-size: 16px;
}

.fas {
  color: white;
  width: 16px;
}
</style>