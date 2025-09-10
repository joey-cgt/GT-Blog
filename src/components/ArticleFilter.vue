<script setup>
import { computed } from 'vue'
import { articles } from '../store/blog.js'

const props = defineProps({
  filterType: {
    type: String,
    default: 'all'
  },
  filterValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['filter-change'])

// 获取所有分类和标签
const allCategories = computed(() => {
  return [...new Set(articles.map(article => article.category))]
})

const allTags = computed(() => {
  const tagSet = new Set()
  articles.forEach(article => {
    article.tags.forEach(tag => tagSet.add(tag))
  })
  return Array.from(tagSet)
})

function handleFilter(type, value = '') {
  emit('filter-change', { type, value })
}
</script>

<template>
  <div class="filter-container">
    <!-- 全部 -->
    <div class="filter-row">
      <span class="filter-label">全部：</span>
      <button 
        :class="['filter-button', { active: filterType === 'all' }]"
        @click="handleFilter('all')"
      >
        全部
      </button>
    </div>
    
    <!-- 分类 -->
    <div class="filter-row">
      <span class="filter-label">分类：</span>
      <div class="filter-items">
        <button 
          v-for="category in allCategories" 
          :key="category"
          :class="['filter-button', { active: filterType === 'category' && filterValue === category }]"
          @click="handleFilter('category', category)"
        >
          {{ category }}
        </button>
      </div>
    </div>
    
    <!-- 标签 -->
    <div class="filter-row">
      <span class="filter-label">标签：</span>
      <div class="filter-items">
        <button 
          v-for="tag in allTags" 
          :key="tag"
          :class="['filter-button', { active: filterType === 'tag' && filterValue === tag }]"
          @click="handleFilter('tag', tag)"
        >
          {{ tag }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.filter-container {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  margin-bottom: 20px;
}

.filter-row {
  display: flex;
  align-items: flex-start;
  margin-bottom: 15px;
}

.filter-row:last-child {
  margin-bottom: 0;
}

.filter-label {
  font-weight: 500;
  color: #2c3e50;
  min-width: 60px;
  padding-top: 6px;
}

.filter-items {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  flex: 1;
}

.filter-button {
  padding: 6px 12px;
  border: 1px solid #e0e0e0;
  border-radius: 6px;
  background: white;
  color: #666;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 14px;
}

.filter-button:hover {
  border-color: #3498db;
  color: #3498db;
}

.filter-button.active {
  background: #7048ff;
  color: white;
  border-color: #6338ff;
}
</style>