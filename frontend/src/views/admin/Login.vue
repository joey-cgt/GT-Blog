<template>
  <div class="admin-login-container">
    <div class="login-box">
      <div class="login-header">
        <h1>GT-Blog 博客管理系统</h1>
        <p>请登录以继续</p>
      </div>
      <div class="login-form">
        <div class="form-group">
          <label for="username">用户名：</label>
          <div class="input-container">
            <input 
              type="text" 
              id="username" 
              v-model="loginForm.username" 
              placeholder="请输入用户名"
              @keyup.enter="handleLogin"
            >
          </div>
        </div>
        <div class="form-group">
          <label for="password">密　码：</label>
          <div class="input-container">
            <input 
              type="password" 
              id="password" 
              v-model="loginForm.password" 
              placeholder="请输入密码"
              @keyup.enter="handleLogin"
            >
          </div>
        </div>
        <div class="form-group">
          <button 
            class="login-button" 
            @click="handleLogin" 
            :disabled="isLoading"
          >
            {{ isLoading ? '登录中...' : '登录' }}
          </button>
        </div>
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isLoading = ref(false)
const errorMessage = ref('')

const loginForm = reactive({
  username: '',
  password: '',
  remember: false
})

const handleLogin = async () => {
  // 表单验证
  if (!loginForm.username || !loginForm.password) {
    errorMessage.value = '用户名和密码不能为空'
    return
  }

  try {
    isLoading.value = true
    errorMessage.value = ''
    
    // 模拟登录过程
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // 验证用户名和密码
    if (loginForm.username === 'admin' && loginForm.password === '123456') {
      // 存储登录状态
      localStorage.setItem('isLoggedIn', 'true')
      
      // 如果选择了"记住我"，可以存储用户名
      if (loginForm.remember) {
        localStorage.setItem('rememberedUsername', loginForm.username)
      } else {
        localStorage.removeItem('rememberedUsername')
      }
      
      // 登录成功后跳转到管理后台首页
      router.push('/admin')
    } else {
      errorMessage.value = '用户名或密码错误'
    }
  } catch (error) {
    console.error('登录失败:', error)
    errorMessage.value = '登录失败，请检查用户名和密码'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.admin-login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f0f2f5;
  padding: 20px;
}

.login-box {
  width: 100%;
  max-width: 400px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding: 30px;
}

.login-header {
  text-align: center;
  margin-bottom: 15px;
}

.login-header h1 {
  font-size: 24px;
  color: #333;
  margin-bottom: 8px;
}

.login-header p {
  color: #666;
  font-size: 14px;
}

.login-form {
  width: 100%;
  max-width: 320px;
  margin: 0 auto;
}

.form-group {
  margin-bottom: 20px;
  display: flex;
  align-items: center;
}

.form-group label {
  width: 70px;
  font-size: 14px;
  color: #333;
  text-align: justify;
  text-align-last: justify;
  flex-shrink: 0;
}

.input-container {
  flex: 1;
}

.form-group input[type="text"],
.form-group input[type="password"] {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  transition: border-color 0.3s;
}

.form-group input:focus {
  border-color: #1890ff;
  outline: none;
  box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
}

.remember-group input {
  margin-right: 8px;
}

.remember-group label {
  font-size: 14px;
  color: #333;
  cursor: pointer;
}

.login-button {
  width: 100%;
  background-color: #1890ff;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: #40a9ff;
}

.login-button:disabled {
  background-color: #bae7ff;
  cursor: not-allowed;
}

.error-message {
  color: #f5222d;
  font-size: 14px;
  text-align: center;
  margin-top: 16px;
}

@media (max-width: 576px) {
  .login-box {
    padding: 20px;
  }
}
</style>