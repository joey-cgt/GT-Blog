<script setup>
import { popularArticles } from '../../store/blog.js'
</script>

<template>
  <section class="sidebar-section popular-section">
    <h3 class="sidebar-title">热门文章</h3>
    <ul class="popular-list">
      <li v-for="article in popularArticles" :key="article.id">
        <a href="#">
          {{ article.title }}
          <span class="views">{{ article.views }} 阅读</span>
        </a>
      </li>
    </ul>
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

.popular-list li {
  border-bottom: 1px solid #eee;
  padding: 12px 0;
  text-align: left;
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

.popular-list .views {
  font-size: 12px;
  color: #777;
  margin-top: 5px;
}
</style>