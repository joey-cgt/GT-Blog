<script setup>
import { columns } from '../store/blog.js'
import ColumnCard from '../components/ColumnCard.vue'
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