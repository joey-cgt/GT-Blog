<template>
  <div class="password-settings-container">
    <h2 class="page-title">修改密码</h2>
    
    <el-card class="password-form-card">
      <el-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-width="100px"
        class="password-form"
      >
        <el-form-item label="当前密码" prop="currentPassword">
          <el-input
            v-model="passwordForm.currentPassword"
            type="password"
            placeholder="请输入当前密码"
            show-password
            class="form-input"
          />
        </el-form-item>
        
        <el-form-item label="新密码" prop="newPassword">
          <el-input
            v-model="passwordForm.newPassword"
            type="password"
            placeholder="请输入新密码"
            show-password
            class="form-input"
          />
          <div class="password-hint">
            密码长度8-20位，包含字母和数字
          </div>
        </el-form-item>
        
        <el-form-item label="确认新密码" prop="confirmPassword">
          <el-input
            v-model="passwordForm.confirmPassword"
            type="password"
            placeholder="请再次输入新密码"
            show-password
            class="form-input"
          />
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="handleSubmit" :loading="loading">
            确认修改
          </el-button>
          <el-button @click="handleCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { updatePassword } from '@/api/admin'

const router = useRouter()
const passwordFormRef = ref(null)
const loading = ref(false)

// 表单数据
const passwordForm = reactive({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

// 表单验证规则
const passwordRules = {
  currentPassword: [
    { required: true, message: '请输入当前密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度应在6-20位之间', trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 8, max: 20, message: '密码长度应在8-20位之间', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (!/^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,20}$/.test(value)) {
          callback(new Error('密码必须包含字母和数字'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  confirmPassword: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.newPassword) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

// 提交表单
const handleSubmit = async () => {
  try {
    await passwordFormRef.value.validate()
    loading.value = true
    
    const response = await updatePassword({
      currentPassword: passwordForm.currentPassword,
      newPassword: passwordForm.newPassword
    })
    
    if (response.code === 200) {
      ElMessage.success('密码修改成功，请重新登录')
      // 清除本地存储并跳转到登录页
      localStorage.removeItem('isLoggedIn')
      localStorage.removeItem('rememberedUsername')
      setTimeout(() => {
        router.push('/login')
      }, 1500)
    } else {
      ElMessage.error(response.message || '密码修改失败')
    }
  } catch (error) {
    // 增强错误处理，提供更具体的错误提示
    if (error.name === 'Error') {
      ElMessage.error(error.message)
    } else if (error.response) {
      // 服务器返回了错误响应
      // 首先检查响应体中的code字段（业务错误码）
      if (error.response.data && error.response.data.code === 400) {
        // 业务逻辑错误：密码错误等
        if (error.response.data.message) {
          // 根据后端返回的具体错误信息显示对应提示
          if (error.response.data.message === 'password is incorrect') {
            ElMessage.error('当前密码错误，请重新输入')
          } else {
            ElMessage.error(error.response.data.message)
          }
        } else {
          ElMessage.error('操作失败，请稍后重试')
        }
      } else {
        // 其他HTTP错误
        ElMessage.error(`请求失败: ${error.response.status} - ${error.response.statusText}`)
      }
    } else if (error.request) {
      // 请求已发送但没有收到响应
      ElMessage.error('网络错误，请检查您的网络连接')
    } else {
      // 其他类型的错误
      ElMessage.error('密码修改失败，请稍后重试')
    }
    console.error('密码修改错误:', error)
  } finally {
    loading.value = false
  }
}

// 取消操作
const handleCancel = () => {
  passwordFormRef.value.resetFields()
  router.push('/admin/dashboard')
}
</script>

<style scoped>
.password-settings-container {
  padding: 20px;
  max-width: 600px;
  margin: 0 auto;
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 24px;
  color: #303133;
}

.password-form-card {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.password-form {
  padding: 20px;
}

.form-input {
  width: 100%;
  max-width: 400px;
}

.password-hint {
  color: #909399;
  font-size: 12px;
  margin-top: 4px;
}

.el-button {
  margin-right: 12px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .password-settings-container {
    padding: 10px;
  }
  
  .password-form {
    padding: 15px;
  }
  
  .form-input {
    max-width: 100%;
  }
}
</style>