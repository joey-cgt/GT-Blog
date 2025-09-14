<script setup>
import { columns } from '../../../store/blog.js'
</script>

<template>
  <section class="sidebar-section columns-section">
    <h3 class="sidebar-title">专栏</h3>
    <ul class="columns-list">
      <li v-for="column in columns" :key="column.title">
        <a href="#">
          {{ column.title }}
          <span class="count">{{ column.articleCount }}</span>
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

.columns-list li {
  border-bottom: 1px solid #eee;
  padding: 10px 0;
  list-style-type: none;
}

.columns-list li:last-child {
  border-bottom: none;
}

.columns-list a {
  display: flex;
  justify-content: space-between;
  color: #555;
}

.columns-list a:hover {
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