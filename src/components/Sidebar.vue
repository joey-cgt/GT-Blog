<script setup>
import CategoryList from './sidebar/CategoryList.vue'
import TagCloud from './sidebar/TagCloud.vue'
import EmailSubscription from './sidebar/EmailSubscription.vue'
import ColumnsList from './sidebar/ColumnsList.vue';
</script>

<template>
  <aside class="sidebar">
    <CategoryList />
    <ColumnsList />
    <TagCloud />
    <EmailSubscription />
  </aside>
</template>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
}
</style>