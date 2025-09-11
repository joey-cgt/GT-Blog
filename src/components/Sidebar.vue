<script setup>
import { useRoute } from 'vue-router'
import { computed } from 'vue'
import CategoryList from './sidebar/CategoryList.vue'
import TagCloud from './sidebar/TagCloud.vue'
import EmailSubscription from './sidebar/EmailSubscription.vue'
import ColumnsList from './sidebar/ColumnsList.vue'
import RecommendedArticles from './sidebar/RecommendedArticles.vue'
import BlogInfoCard from './sidebar/BlogInfoCard.vue'
import TableOfContents from './sidebar/TableOfContents.vue'

const route = useRoute()

// 使用computed确保响应式更新
const isHomePage = computed(() => route.name === 'Home')
const isContentPage = computed(() => ['Articles', 'Columns', 'ColumnArticles', 'Categories', 'Tags', 'About', 'ArticleContent'].includes(route.name))
</script>

<template>
  <aside class="sidebar">
    <!-- 首页显示分类相关控件 -->
    <template v-if="isHomePage">
      <CategoryList />
      <ColumnsList />
      <TagCloud />
    </template>
    
    <!-- 内容页面显示博客信息和热门文章 -->
    <template v-else>
      <BlogInfoCard />
      <TableOfContents v-if="route.name === 'ArticleContent'" />
      <RecommendedArticles />
    </template>
    
    <!-- 所有页面都显示邮件订阅 -->
    <EmailSubscription />
  </aside>
</template>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
}
</style>