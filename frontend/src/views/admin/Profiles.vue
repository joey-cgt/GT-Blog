<template>
  <div class="profiles-page" v-loading="loading">
    <div class="page-header">
      <h2>个人资料</h2>
      <el-button type="primary" @click="handleUpdateInfo">更新信息</el-button>
    </div>
    <div class="content-layout">
      <div class="profile-settings">
        <div class="setting-item">
          <span class="label">头像</span>
          <div class="avatar-section">
            <el-avatar :size="80" :src="profileData.avatarUrl" class="avatar-preview" />
            <el-input v-model="profileData.avatarUrl" placeholder="请输入头像URL" style="width: 300px; margin-left: 16px;" />
          </div>
        </div>
        <div class="setting-item">
          <span class="label">昵称</span>
          <el-input v-model="profileData.nickname" placeholder="请输入昵称" style="width: 200px;" />
        </div>
        <div class="setting-item">
          <span class="label">简介</span>
          <el-input v-model="profileData.bio" placeholder="请输入个人简介" style="width: 300px;" type="textarea" :rows="3" />
        </div>
        <div class="setting-item">
          <span class="label">邮箱</span>
          <el-input v-model="profileData.email" placeholder="请输入邮箱" style="width: 200px;" />
        </div>
        <div class="setting-item">
          <span class="label">微信</span>
          <el-input v-model="profileData.wechat" placeholder="请输入微信号" style="width: 200px;" />
        </div>
        <div class="setting-item">
          <span class="label">链接账号</span>
          <div class="social-media-table">
            <el-table :data="profileData.socialAccounts" style="width: 400px;">
              <el-table-column prop="platform" label="平台" width="120">
                <template #default="{ row, $index }">
                  <el-input v-model="row.platform" :placeholder="'平台' + ($index + 1)" />
                </template>
              </el-table-column>
              <el-table-column prop="url" label="链接">
                <template #default="{ row, $index }">
                  <el-input v-model="row.url" :placeholder="'链接' + ($index + 1)" /> 
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>
        <div class="setting-item">
          <span class="label">关于我</span>
          <el-input v-model="profileData.aboutMe" placeholder="请输入关于我的Markdown内容" type="textarea" :rows="5" style="width: 500px; resize: none;" @input="handleAboutInput" />
        </div>
        <div class="setting-item">
          <span class="label">关于博客</span>
          <el-input v-model="profileData.aboutBlog" placeholder="请输入博客简介" type="textarea" :rows="3" style="width: 500px; resize: none;" @input="handleBlogAboutInput" />
        </div>
        
        
      </div>
      <div class="preview-about">
        <h3>"关于我"预览</h3>
        <MarkdownRenderer :content="profileData.aboutMe" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import MarkdownRenderer from '../../components/common/MarkdownRenderer.vue'
import {
  getProfile,
  updateProfile
 } from '@/api/admin'

const loading = ref(false)

// 初始化数据
const profileData = ref({
  nickname: '',
  avatarUrl: '',
  bio: '',
  email: '',
  wechat: '',
  aboutBlog: '',
  aboutMe: '',
  socialAccounts: [
    { platform: '', url: '' },
    { platform: '', url: '' },
    { platform: '', url: '' },
    { platform: '', url: '' }
  ]
})

// 获取用户资料
const fetchProfile = async () => {
  loading.value = true
  try {
    // 直接调用getProfile，不传递用户名参数，完全依赖JWT token
    const res = await getProfile()
    
    if (res.data) {
      profileData.value = {
        nickname: res.data.nickname || '',
        bio: res.data.bio || '',
        email: res.data.email || '',
        wechat: res.data.wechat || '',
        aboutMe: res.data.aboutMe || '',
        aboutBlog: res.data.aboutBlog || '',
        avatarUrl: res.data.avatarUrl || '',
        socialAccounts: ensureFourSocialAccounts(res.data.socialAccounts)
      }
    }
  } catch (error) {
    console.error('请求错误:', error)
    ElMessage.error('获取资料失败')
  } finally {
    loading.value = false
  }
}

function ensureFourSocialAccounts(socialAccounts) {
  // 如果不存在，初始化为空数组
  const accounts = Array.isArray(socialAccounts) ? [...socialAccounts] : []
  
  // 确保至少有4个元素，不足的用空对象补充
  while (accounts.length < 4) {
    accounts.push({ platform: '', url: '' })
  }
  
  return accounts
}

// 组件挂载时加载数据
onMounted(fetchProfile)

const handleUpdateInfo = async () => {
  try {
    await ElMessageBox.confirm(
      '确定要更新所有个人信息吗？',
      '确认更新',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    // 提交表单
    const submitData = {
      ...profileData.value
    }
    await updateProfile(submitData)
    ElMessage.success('个人信息更新成功！')
    // 刷新个人资料数据
    await fetchProfile()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('更新失败: ' + error.message)
    }
  }
}

const handleAboutInput = () => {
  // 输入时触发更新，确保实时渲染
  // about ref 已经通过 v-model 更新，这里主要是确保响应性
}

const handleBlogAboutInput = () => {
  // 输入时触发更新，确保实时渲染
  // blogAbout ref 已经通过 v-model 更新，这里主要是确保响应性
}
</script>

<style scoped>
.profiles-page {
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e8e8e8;
}

.content-layout {
  display: flex;
  gap: 24px;
}

.profile-settings {
  flex: 1;
  background: #fff;
  border-radius: 8px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  max-height: 750px;
  overflow-y: auto;
}

.preview-about {
  flex: 1;
  background: #fff;
  border-radius: 8px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  max-height: 750px;
  overflow-y: auto;
}

.preview-about h3 {
  margin-bottom: 16px;
  color: #303133;
}

.setting-item {
  display: flex;
  align-items: center;
  margin-bottom: 16px; 
}

.setting-item .label {
  width: 80px;
  font-weight: 500;
  color: #606266;
  align-self: flex-start;
  padding-top: 8px;
}

.preview-about h3 {
  margin-bottom: 16px;
  color: #636363;
  text-align: right;
}

.social-media-table {
  border: 1px solid #ebeef5;
  border-radius: 4px;
  overflow: hidden;
}

.avatar-section {
  display: flex;
  align-items: center;
}

.avatar-preview {
  border: 2px solid #dcdfe6;
}

/* 禁用 textarea 拖拽 */
:deep(.el-textarea__inner) {
  resize: none !important;
}

/* 响应式布局 */
@media (max-width: 1024px) {
  .content-layout {
    flex-direction: column;
  }
}
</style>

