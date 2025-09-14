<script setup>
import { categories } from '../../../store/blog.js'
</script>

<template>
  <section class="sidebar-section categories-section">
    <h3 class="sidebar-title">文章分类</h3>
    <ul class="categories-list">
      <li v-for="category in categories" :key="category.name">
        <a href="#">
          {{ category.name }}
          <span class="count">{{ category.count }}</span>
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

.categories-list li {
  border-bottom: 1px solid #eee;
  padding: 10px 0;
  list-style-type: none;
}

.categories-list li:last-child {
  border-bottom: none;
}

.categories-list a {
  display: flex;
  justify-content: space-between;
  color: #555;
}

.categories-list a:hover {
  color: #3498db;
}

.count {
  background-color: #f0f0f0;
  color: #777;
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 12px;
}
</style>