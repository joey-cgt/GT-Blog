<script setup>
import { computed } from 'vue'
import ArticleCard from './ArticleCard.vue'
import { getPopularArticles, getLatestArticles, getPinnedArticles } from '@/api/article'
import { ref, onMounted } from 'vue'

const props = defineProps({
  type: {
    type: String,
    default: 'latest',
    validator: (value) => ['pinned', 'latest', 'popular'].includes(value)
  },
  title: {
    type: String,
    required: true
  },
  limit: {
    type: Number,
    default: 5
  }
})



// 定义响应式数据
const topArticles = ref([])
const popularArticles = ref([])
const latestArticles = ref([])

// 从API获取文章数据
onMounted(async () => {
  try {
    const pinnedResponse = await getPinnedArticles()
    const latestResponse = await getLatestArticles()
    const popularResponse = await getPopularArticles()
    
    topArticles.value = pinnedResponse.data.items
    popularArticles.value = popularResponse.data.items
    latestArticles.value = latestResponse.data.items
  } catch (error) {
    console.error('获取文章数据失败:', error)
  }
})

const filteredArticles = computed(() => {
  switch (props.type) {
    case 'pinned':
      return topArticles.value.slice(0, 3)
    case 'latest':
      return latestArticles.value.slice(0, props.limit)
    case 'popular':
      return popularArticles.value.slice(0, props.limit)
    default:
      return []
  }
})
</script>

<template>
  <section class="articles-section">
    <div class="section-header">
      <h2 class="section-title">
        <span v-if="type === 'pinned'">📌</span>
        <span v-else-if="type === 'latest'">🕒</span>
        <span v-else-if="type === 'popular'">🔥</span>
        {{ title }}
      </h2>
    </div>
    
    <ArticleCard 
      v-for="article in filteredArticles" 
      :key="article.id" 
      :article="article" 
    />
  </section>
</template>

<style scoped>
.section-header {
  margin-bottom: 16px;
}

.section-title {
  font-size: 20px;
  color: #2c3e50;
}
</style>