<script setup>
import { ref } from 'vue'

const emailSubscription = ref('')

const subscribeEmail = () => {
  if (emailSubscription.value && /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(emailSubscription.value)) {
    alert(`感谢订阅！邮件 ${emailSubscription.value} 已添加到订阅列表`)
    emailSubscription.value = ''
  } else {
    alert('请输入有效的邮箱地址')
  }
}
</script>

<template>
  <section class="sidebar-section subscribe-section">
    <h3 class="sidebar-title">邮件订阅</h3>
    <p>订阅博客更新，不错过任何文章</p>
    <div class="subscribe-form">
      <input 
        type="email" 
        v-model="emailSubscription" 
        placeholder="请输入您的邮箱" 
        class="email-input"
      />
      <button @click="subscribeEmail" class="subscribe-button">订阅</button>
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

.subscribe-section p {
  margin-bottom: 15px;
  color: #666;
}

.subscribe-form {
  display: flex;
}

.email-input {
  flex: 1;
  padding: 10px 15px;
  border: 1px solid #ddd;
  border-radius: 4px 0 0 4px;
  font-size: 14px;
}

.subscribe-button {
  background-color: #3498db;
  color: #fff;
  border: none;
  padding: 0 20px;
  border-radius: 0 4px 4px 0;
  cursor: pointer;
  transition: background-color 0.3s;
  font-size: 14px;
}

.subscribe-button:hover {
  background-color: #2980b9;
}
</style>