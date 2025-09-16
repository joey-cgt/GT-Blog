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
          <div class="avatar-upload">
            <el-upload
              action="#"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :http-request="handleAvatarUpload"
            >
              <el-avatar :size="80" :src="profileData.avatar" class="avatar-preview clickable-avatar" />
            </el-upload>
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
            <el-table :data="socialMedia" style="width: 400px;">
              <el-table-column prop="name" label="名称" width="120">
                <template #default="{ row, $index }">
                  <el-input v-model="row.name" :placeholder="'名称' + ($index + 1)" />
                </template>
              </el-table-column>
              <el-table-column prop="link" label="链接">
                <template #default="{ row, $index }">
                  <el-input v-model="row.link" :placeholder="'链接' + ($index + 1)" />
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>
        <div class="setting-item">
          <span class="label">关于我</span>
          <el-input v-model="profileData.about" placeholder="请输入关于我的Markdown内容" type="textarea" :rows="5" style="width: 500px; resize: none;" @input="handleAboutInput" />
        </div>
        <div class="setting-item">
          <span class="label">关于博客</span>
          <el-input v-model="profileData.blogAbout" placeholder="请输入博客简介" type="textarea" :rows="3" style="width: 500px; resize: none;" @input="handleBlogAboutInput" />
        </div>
        
        
      </div>
      <div class="preview-about">
        <h3>"关于我"预览</h3>
        <MarkdownRenderer :content="profileData.about" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import MarkdownRenderer from '../../components/common/MarkdownRenderer.vue'
import { getProfile } from '@/api/profile'

const route = useRoute()
const loading = ref(false)

// 初始化数据
const profileData = ref({
  nickname: '',
  bio: '',
  email: '',
  wechat: '',
  about: '',
  blogAbout: '',
  avatar: '',
  socialMedia: []
})

// 获取用户资料
const fetchProfile = async () => {
  loading.value = true
  try {
    const res = await getProfile(route.params.id || '1')
    
    if (res.data) {
      profileData.value = {
        nickname: res.data.nickname || '',
        bio: res.data.bio || '',
        email: res.data.email || '',
        wechat: res.data.wechat || '',
        about: res.data.about || '',
        blogAbout: res.data.blogAbout || '',
        avatar: res.data.avatar || '',
        socialMedia: res.data.socialMedia ? JSON.parse(res.data.socialMedia) : []
      }
    }
  } catch (error) {
    console.error('请求错误:', error)
    ElMessage.error('获取资料失败')
  } finally {
    loading.value = false
  }
}

// 组件挂载时加载数据
onMounted(fetchProfile)


// 直接使用profileData管理数据
const beforeAvatarUpload = (file) => {
  const isJPGOrPNG = file.type === 'image/jpeg' || file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isJPGOrPNG) {
    ElMessage.error('头像只能是 JPG 或 PNG 格式!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('头像大小不能超过 2MB!')
    return false
  }
  return true
}

const handleAvatarUpload = async ({ file }) => {
  // 这里处理头像上传逻辑
  // 实际项目中应该调用API上传到服务器
  const reader = new FileReader()
  reader.onload = (e) => {
    avatarUrl.value = e.target.result
  }
  reader.readAsDataURL(file)
}

const socialMedia = ref([
  { name: '', link: '' },
  { name: '', link: '' },
  { name: '', link: '' },
  { name: '', link: '' }
])

const handleAboutInput = () => {
  // 输入时触发更新，确保实时渲染
  // about ref 已经通过 v-model 更新，这里主要是确保响应性
}

const handleBlogAboutInput = () => {
  // 输入时触发更新，确保实时渲染
  // blogAbout ref 已经通过 v-model 更新，这里主要是确保响应性
}

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
    
    // 这里执行实际的更新逻辑
    const updateData = {
      nickname: nickname.value,
      bio: bio.value,
      email: email.value,
      wechat: wechat.value,
      about: about.value,
      avatar: avatarUrl.value,
      socialMedia: socialMedia.value
    }
    
    console.log('更新数据:', updateData)
    
    ElMessage.success('个人信息更新成功！')
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.info('已取消更新')
    }
  }
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

.avatar-upload {
  display: flex;
  align-items: center;
  gap: 16px;
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

