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
</script>

<template>
  <aside class="sidebar">
    <!-- 首页显示分类相关控件 -->
    <template v-if="isHomePage">
      <CategoryList />
      <ColumnsList />
      <TagCloud />
      <EmailSubscription />
    </template>
    
    <!-- 内容页面显示博客信息和热门文章 -->
    <template v-else>
      <BlogInfoCard />
      <div class="sticky-sidebar-content">
        <TableOfContents v-if="route.name === 'ArticleContent'" />
        <RecommendedArticles :article-id="route.params.id" />
       
      </div>
    </template>
    
  </aside>
</template>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.sticky-sidebar-content {
  position: sticky;
  top: 90px; /* Header高度70px + 20px间距 */
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .sticky-sidebar-content {
    position: static; /* 移动端取消粘性定位 */
  }
}
</style>