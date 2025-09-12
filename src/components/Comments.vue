<script setup>
import { ref } from 'vue'

const newComment = ref('')
const comments = ref([])

const emit = defineEmits(['commentSubmitted'])

const submitComment = () => {
  if (newComment.value.trim()) {
    const comment = {
      id: Date.now(),
      author: '匿名用户',
      content: newComment.value.trim(),
      time: new Date().toLocaleString('zh-CN')
    }
    comments.value.unshift(comment)
    newComment.value = ''
    emit('commentSubmitted', comment)
  }
}
</script>

<template>
  <div id="comments" class="comments-section">
    <h3 class="comments-title">评论</h3>
    <div class="comment-form">
      <textarea v-model="newComment" placeholder="写下你的评论..." class="comment-input"></textarea>
      <button @click="submitComment" class="submit-comment-btn">发表评论</button>
    </div>
    <div class="comments-list">
      <div v-for="comment in comments" :key="comment.id" class="comment-item">
        <div class="comment-header">
          <span class="comment-author">{{ comment.author }}</span>
          <span class="comment-time">{{ comment.time }}</span>
        </div>
        <div class="comment-content">{{ comment.content }}</div>
      </div>
      <div v-if="comments.length === 0" class="no-comments">
        暂无评论，快来发表第一条评论吧！
      </div>
    </div>
  </div>
</template>

<style scoped>
.comments-section {
  padding: 20px;
  margin-top: 40px;
  padding-top: 30px;
  border-top: 1px solid #eaeaea;
}

.comments-title {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 20px;
  text-align: left;
}

.comment-form {
  margin-bottom: 30px;
}

.comment-input {
  width: 100%;
  min-height: 100px;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 14px;
  line-height: 1.5;
  resize: vertical;
  margin-bottom: 12px;
}

.comment-input:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.2);
}

.submit-comment-btn {
  background-color: #3498db;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
  margin-left: auto;
  display: block;
}

.submit-comment-btn:hover {
  background-color: #2980b9;
}

.comments-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.comment-item {
  padding: 16px;
  background-color: #f8f9fa;
  border-radius: 8px;
  /* border-left: 3px solid #3498db; */
}

.comment-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.comment-author {
  font-weight: 500;
  color: #2c3e50;
}

.comment-time {
  font-size: 12px;
  color: #7f8c8d;
}

.comment-content {
  color: #333;
  line-height: 1.5;
  font-size: 14px;
  text-align: left;
}

.no-comments {
  text-align: center;
  color: #7f8c8d;
  font-style: italic;
  padding: 40px 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .comments-section {
    margin-top: 30px;
    padding-top: 20px;
  }

  .comments-title {
    font-size: 18px;
  }

  .comment-item {
    padding: 12px;
  }
}
</style>