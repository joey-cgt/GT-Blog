<script setup>
import { useRouter } from 'vue-router'

const props = defineProps({
  column: {
    type: Object,
    required: true
  }
})

const router = useRouter()

const navigateToColumn = () => {
  router.push(`/columns/${props.column.id}`)
}
</script>

<template>
  <div class="column-card" @click="navigateToColumn">
    <div class="column-cover">
      <img :src="column.cover" :alt="column.title">
    </div>
    <div class="column-content">
      <h3 class="column-title">{{ column.title }}</h3>
      <p class="column-desc">{{ column.description }}</p>
      <div class="column-meta">
        <span class="meta-item">📅 {{ column.createTime }}</span>
        <span class="meta-item">📝 {{ column.articleCount }}篇</span>
        <span class="meta-item">👀 {{ column.totalViews }}次</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.column-card {
  /* border: 1px solid rgba(0,0,0,0.1); */
  border-radius: 12px;
  overflow: hidden;
  transition: all 0.3s ease;
  display: flex;
  padding: 16px;
  gap: 20px;
  background: linear-gradient(135deg, #f8f9fa 0%, #ffffff 100%);
  box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}

.column-cover {
  flex: 0 0 200px;
}

.column-cover img {
  width: 100%;
  height: 120px;
  object-fit: cover;
  border-radius: 4px;
}

.column-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  text-align: left;
}

.column-title {
  font-size: 18px;
  margin: 0 0 10px 0;
  color: #2c3e50;
  transition: color 0.2s ease;
  cursor: pointer;
}

.column-title:hover {
  color: #3498db;
}

.column-desc {
  color: #666;
  font-size: 14px;
  margin: 0 0 auto 0;
  line-height: 1.5;
}

.column-meta {
  display: flex;
  gap: 12px;
  font-size: 13px;
  color: #7f8c8d;
  margin-top: 12px;
}
</style>