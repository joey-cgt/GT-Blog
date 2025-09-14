<script setup>
import { columns } from '../../store/blog.js'
import ColumnCard from '../../components/visitor/ColumnCard.vue'
</script>

<template>
  <div class="columns-page">
    <div class="columns-list">
      <ColumnCard 
        v-for="column in columns" 
        :key="column.id"
        :column="column"
      />
    </div>
  </div>
</template>

<style scoped>
.columns-page {
  padding: 0 0 20px;
}

.columns-page .column-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 20px rgba(0,0,0,0.12);
  background: linear-gradient(135deg, #f1f3f5 0%, #ffffff 100%);
}

.columns-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 100%;
}

.column-card {
  width: 100%;
  max-width: 100%;
  box-sizing: border-box;
}
</style>