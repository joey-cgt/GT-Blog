<script setup>
import { blogInfo } from '../../store/blog.js'
</script>

<template>
  <!-- 博主信息卡片 -->
  <section class="author-card">
    <div class="author-info">
      <div class="author-avatar-container">
        <img :src="blogInfo.avatar" alt="博主头像" class="author-avatar" />
      </div>
      <div class="author-details">
        <div class="author-name-row">
          <h2 class="author-name">{{ blogInfo.author }}</h2>
          <div class="social-links">
            <a v-for="link in blogInfo.socialLinks" :key="link.name" :href="link.url" target="_blank" class="social-link">
              <span class="icon" :class="link.icon"></span>
              {{ link.name }}
            </a>
          </div>
        </div>
        <p class="author-description">{{ blogInfo.description }}</p>
      </div>
    </div>
  </section>
</template>

<style scoped>
/* 博主信息卡片 - 独特样式 */
.author-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 16px;
  padding: 40px;
  margin-bottom: 30px;
  box-shadow: 0 10px 30px rgba(102, 126, 234, 0.3);
  color: white;
  position: relative;
  overflow: hidden;
}

.author-card::before {
  content: '';
  position: absolute;
  top: -50%;
  right: -50%;
  width: 100%;
  height: 200%;
  background: rgba(255, 255, 255, 0.1);
  transform: rotate(45deg);
  pointer-events: none;
}

.author-info {
  display: flex;
  gap: 20px;
}

.author-avatar-container {
  flex-shrink: 0;
}

.author-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.author-details {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.author-name-row {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 8px;
}

.author-name {
  font-size: 24px;
  margin: 0;
  color: white;
  font-weight: 600;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  line-height: 1.2;
  text-align: left;
}

.author-description {
  color: rgba(255, 255, 255, 0.9);
  margin: 0;
  font-size: 14px;
  line-height: 1.6;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
  text-align: left
}

.social-links {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.social-link {
  display: flex;
  align-items: center;
  color: white;
  font-size: 12px;
  font-weight: 500;
  padding: 6px 12px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 16px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  transition: all 0.3s ease;
  text-decoration: none;
}

.social-link:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.social-link .icon {
  margin-right: 4px;
  font-size: 14px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .author-card {
    padding: 25px;
    gap: 15px;
  }
  
  .author-info {
    gap: 15px;
  }
  
  .author-avatar {
    width: 70px;
    height: 70px;
  }
  
  .author-name {
    font-size: 20px;
    margin-bottom: 5px;
  }
  
  .social-links {
    gap: 10px;
  }
  
  .social-link {
    font-size: 11px;
    padding: 5px 10px;
  }
}

@media (max-width: 576px) {
  .author-info {
    flex-direction: column;
    align-items: center;
    text-align: center;
    gap: 12px;
  }
  
  .author-avatar {
    width: 60px;
    height: 60px;
  }
  
  .author-details {
    align-items: center;
  }
  
  .author-name {
    font-size: 18px;
    text-align: center;
  }
  
  .author-description {
    font-size: 13px;
    text-align: center;
  }
  
  .social-links {
    justify-content: center;
  }
}
</style>