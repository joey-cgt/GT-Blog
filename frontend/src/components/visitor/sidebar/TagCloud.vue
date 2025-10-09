<script setup>
import { tags } from '../../../store/blog.js'
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'
import { getTagList } from '@/api/tag.js'

const router = useRouter()

const tagList = ref([])

onMounted(async () => {
  try {
    const response = await getTagList()
    tagList.value = response.data.items
  } catch (error) {
    console.error('获取标签数据失败:', error)
  }
})

const handleTagClick = (tagName) => {
  router.push({
    path: '/articles',
    query: { tag: tagName }
  })
}
</script>

<template>
  <section class="sidebar-section tags-section">
    <h3 class="sidebar-title">热门标签</h3>
    <div class="tags-cloud">
      <a 
        v-for="tag in tagList" 
        :key="tag.name" 
        href="javascript:void(0)" 
        class="tag-cloud-item"
        :class="tag.size"
        @click="handleTagClick(tag.name)"
      >
        {{ tag.name }}
      </a>
    </div>
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

.tags-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.tag-cloud-item {
  background-color: #f0f0f0;
  color: #555;
  padding: 5px 10px;
  border-radius: 4px;
  font-size: 12px;
  transition: all 0.3s;
}

.tag-cloud-item:hover {
  background-color: #3498db;
  color: #fff;
}

.tag-cloud-item.small {
  font-size: 12px;
}

.tag-cloud-item.medium {
  font-size: 14px;
}

.tag-cloud-item.large {
  font-size: 16px;
}
</style>