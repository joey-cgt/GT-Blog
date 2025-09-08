<script setup>
import { computed } from 'vue'

const props = defineProps({
  currentPage: {
    type: Number,
    required: true
  },
  totalPages: {
    type: Number,
    required: true
  },
  maxVisible: {
    type: Number,
    default: 5
  }
})

const emit = defineEmits(['page-change'])

const pages = computed(() => {
  const range = []
  const half = Math.floor(props.maxVisible / 2)
  let start = Math.max(props.currentPage - half, 1)
  let end = Math.min(start + props.maxVisible - 1, props.totalPages)
  
  if (end - start + 1 < props.maxVisible) {
    start = Math.max(end - props.maxVisible + 1, 1)
  }

  for (let i = start; i <= end; i++) {
    range.push(i)
  }
  return range
})

function changePage(page) {
  if (page !== props.currentPage) {
    emit('page-change', page)
  }
}
</script>

<template>
  <div class="pagination">
    <button 
      class="page-nav"
      :disabled="currentPage === 1"
      @click="changePage(currentPage - 1)"
    >
      上一页
    </button>
    
    <button
      v-for="page in pages"
      :key="page"
      class="page-number"
      :class="{ active: page === currentPage }"
      @click="changePage(page)"
    >
      {{ page }}
    </button>
    
    <button 
      class="page-nav"
      :disabled="currentPage === totalPages"
      @click="changePage(currentPage + 1)"
    >
      下一页
    </button>
  </div>
</template>

<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  gap: 8px;
  margin-top: 24px;
}

.page-nav, .page-number {
  padding: 8px 12px;
  border: 1px solid #ddd;
  background: white;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s;
}

.page-nav:hover:not(:disabled),
.page-number:hover:not(.active) {
  background: #f1f1f1;
}

.page-number.active {
  background: #3498db;
  color: white;
  border-color: #3498db;
}

.page-nav:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>