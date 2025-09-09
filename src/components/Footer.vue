<script setup>
import { blogInfo } from '../store/blog.js'
</script>

<template>
  <footer class="footer">
    <div class="footer-container">
      <div class="footer-content">
        <div class="footer-section about">
          <h3>关于博客</h3>
          <p>专注于技术分享的个人博客，涵盖前端、后端、DevOps等多个技术领域。</p>
        </div>
        
        <div class="footer-section links">
          <h3>快速链接</h3>
          <ul>
            <li><a href="#">首页</a></li>
            <li><a href="#">文章归档</a></li>
            <li><a href="#">关于我</a></li>
            <li><a href="#">联系方式</a></li>
          </ul>
        </div>
        
        <div class="footer-section contact">
          <h3>联系方式</h3>
          <p>邮箱：joycgt@126.com</p>
          <p>微信：joycgt</p>
          <div class="social-links">
            <a v-for="link in blogInfo.socialLinks" :key="link.name" :href="link.url" target="_blank" class="social-link">
              <span class="icon" :class="link.icon"></span>
            </a>
          </div>
        </div>
      </div>
      
      <div class="footer-bottom">
        <p>&copy; 2025 {{ blogInfo.title }} | 由 Vue3 强力驱动</p>
      </div>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  max-width: 80%;
  background-color: #2c3e50;
  color: #ecf0f1;
  padding: 60px 0 20px;
}

.footer-container {
  max-width: 90%;
  margin: 0 auto;
  padding: 0 15px;
}

.footer-content {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 30px;
  margin-bottom: 40px;
}

.footer-section h3 {
  font-size: 18px;
  margin-bottom: 20px;
}

.footer-section p {
  margin-bottom: 10px;
  color: #bdc3c7;
}

.footer-section.links ul li {
  margin-bottom: 10px;
}

.footer-section.links a {
  color: #bdc3c7;
}

.footer-section.links a:hover {
  color: #3498db;
}

.footer-section.contact .social-links {
  margin-top: 15px;
}

.footer-section.contact .social-link {
  color: #bdc3c7;
  font-size: 18px;
  margin-right: 15px;
}

.footer-section.contact .social-link:hover {
  color: #3498db;
}

.footer-bottom {
  text-align: center;
  padding-top: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.footer-bottom p {
  font-size: 14px;
  color: #95a5a6;
}

/* 响应式设计 */
@media (max-width: 992px) {
  .footer-content {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 576px) {
  .footer-content {
    grid-template-columns: 1fr;
  }
}
</style>