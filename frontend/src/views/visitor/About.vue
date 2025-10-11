<script setup>
import { ref, onMounted } from 'vue'
import { getAuthor } from '../../api/visitor.js'
import MarkdownRenderer from '../../components/common/MarkdownRenderer.vue'

const aboutAuthor = ref('')

const fetchAuthorInfo = async () => {
  try {
    const response = await getAuthor(1)
    // 确保传递给MarkdownRenderer的是字符串
    aboutAuthor.value = response.data.aboutMe || ''
  } catch (error) {
    console.error('获取作者信息失败:', error)
  }
}

onMounted(fetchAuthorInfo)

</script>

<template>
  <div class="about-page">
    <div class="about-header">
      <h1 class="about-title">关于博主</h1>
      <div class="about-divider"></div>
    </div>
    <div class="about-content">
      <MarkdownRenderer :content="aboutAuthor" />
    </div>
  </div>
</template>

<style scoped>
.about-page {
    padding: 30px;
}

.about-header {
    margin-bottom: 30px;
    text-align: center;
}

.about-title {
    font-size: 28px;
    font-weight: 600;
    color: #333;
    margin: 0 0 16px 0;
    position: relative;
    display: inline-block;
}

.about-content {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  padding: 30px
}

.about-title::after {
    content: '';
    position: absolute;
    bottom: -10px;
    left: 50%;
    transform: translateX(-50%);
    width: 60px;
    height: 4px;
    background: linear-gradient(90deg, #4e6ef2, #7c3aed);
    border-radius: 2px;
}
</style>
