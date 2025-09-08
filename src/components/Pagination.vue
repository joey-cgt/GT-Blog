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
  maxVisibleButtons: {
    type: Number,
    default: 5
  }
})

const emit = defineEmits(['page-change'])

const pages = computed(() => {
  const pages = []
  
  // 计算显示的页码范围
  let startPage = Math.max(1, props.currentPage - Math.floor(props.maxVisibleButtons / 2))
  let endPage = Math.min(props.totalPages, startPage + props.maxVisibleButtons - 1)
  
  // 调整起始页码，确保显示足够的页码按钮
  if (endPage - startPage + 1 < props.maxVisibleButtons) {
    startPage = Math.max(1, endPage - props.maxVisibleButtons + 1)
  }
  
  // 生成页码数组
  for (let i = startPage; i <= endPage; i++) {
    pages.push(i)
  }
  
  return pages
})

const isFirstPage = computed(() => props.currentPage === 1)
const isLastPage = computed(() => props.currentPage === props.totalPages)

function changePage(page) {
  emit('page-change', page)
}
</script>

<template>
  <div class="pagination" v-if="totalPages > 1">
    <!-- 首页按钮 -->
    <button 
      class="pagination-button" 
      :class="{ disabled: isFirstPage }"
      @click="changePage(1)" 
      :disabled="isFirstPage"
    >
      首页
    </button>
    
    <!-- 上一页按钮 -->
    <button 
      class="pagination-button" 
      :class="{ disabled: isFirstPage }"
      @click="changePage(currentPage - 1)" 
      :disabled="isFirstPage"
    >
      上一页
    </button>
    
    <!-- 页码按钮 -->
    <button 
      v-for="page in pages" 
      :key="page" 
      class="pagination-button" 
      :class="{ active: page === currentPage }"
      @click="changePage(page)"
    >
      {{ page }}
    </button>
    
    <!-- 下一页按钮 -->
    <button 
      class="pagination-button" 
      :class="{ disabled: isLastPage }"
      @click="changePage(currentPage + 1)" 
      :disabled="isLastPage"
    >
      下一页
    </button>
    
    <!-- 末页按钮 -->
    <button 
      class="pagination-button" 
      :class="{ disabled: isLastPage }"
      @click="changePage(totalPages)" 
      :disabled="isLastPage"
    >
      末页
    </button>
    
    <!-- 页码信息 -->
    <span class="pagination-info">{{ currentPage }} / {{ totalPages }} 页</span>
  </div>
</template>

<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 30px 0;
  gap: 8px;
}

.pagination-button {
  padding: 8px 12px;
  border: 1px solid #e0e0e0;
  background-color: #fff;
  color: #333;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s;
  font-size: 14px;
}

.pagination-button:hover:not(.disabled):not(.active) {
  background-color: #f5f5f5;
  border-color: #d0d0d0;
}

.pagination-button.active {
  background-color: #3498db;
  color: white;
  border-color: #3498db;
}

.pagination-button.disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.pagination-info {
  margin-left: 15px;
  color: #666;
  font-size: 14px;
}

@media (max-width: 768px) {
  .pagination {
    flex-wrap: wrap;
  }
  
  .pagination-button {
    padding: 6px 10px;
    font-size: 13px;
  }
}
</style>