<template>
  <div class="article-editor">
    <!-- 顶部标题和操作区 -->
    <div class="editor-header">
      <div class="title-section">
        <el-input
          v-model="articleForm.title"
          placeholder="请输入文章标题"
          maxlength="100"
          show-word-limit
          class="title-input"
          size="large"
        />
      </div>
      <div class="action-section">
        <el-button @click="handleCancel">取消</el-button>
        <!-- 新建文章或草稿文章：显示保存草稿和发布文章 -->
        <template v-if="articleStatus === 'new' || articleStatus === 'draft'">
          <el-button type="primary" @click="handleSaveDraft">
            保存草稿
          </el-button>
          <el-button type="success" @click="showPublishDialog">
            发布文章
          </el-button>
        </template>
        <!-- 已发布文章：只显示更新文章 -->
        <template v-else-if="articleStatus === 'published'">
          <el-button type="success" @click="showPublishDialog">
            更新文章
          </el-button>
        </template>
      </div>
    </div>

    <!-- 主编辑区域 -->
    <div class="editor-main">
      <!-- 左侧Markdown编辑区 -->
      <div class="editor-panel">
        <textarea
          v-model="articleForm.content"
          placeholder="请输入文章内容（支持Markdown语法）"
          class="markdown-editor"
          @input="handleContentInput"
          @keydown.tab.prevent="handleTabKey"
        ></textarea>
      </div>

      <!-- 右侧预览区 -->
      <div class="preview-panel">
        <div class="markdown-preview">
          <MarkdownRenderer :content="compiledMarkdown" />
        </div>
      </div>
    </div>

    <!-- 发布设置对话框 -->
    <el-dialog
      v-model="publishDialogVisible"
      :title="(articleStatus === 'new' || articleStatus === 'draft') ? '发布设置' : '更新文章'"
      width="600px"
      :before-close="handlePublishDialogClose"
    >
      <el-form :model="publishForm" label-width="80px">
        <el-form-item label="文章分类">
          <el-select
            v-model="publishForm.category"
            placeholder="请选择分类"
            clearable
            style="width: 100%"
          >
            <el-option label="前端开发" value="frontend" />
            <el-option label="后端开发" value="backend" />
            <el-option label="UI框架" value="ui" />
            <el-option label="性能优化" value="performance" />
            <el-option label="CSS" value="css" />
            <el-option label="JavaScript" value="javascript" />
          </el-select>
        </el-form-item>

        <el-form-item label="加入专栏">
          <el-select
            v-model="publishForm.column"
            placeholder="请选择专栏"
            clearable
            style="width: 100%"
          >
            <el-option label="Vue.js 精讲" value="vuejs" />
            <el-option label="React 实战" value="react" />
            <el-option label="TypeScript 进阶" value="typescript" />
            <el-option label="前端性能优化" value="frontend-performance" />
            <el-option label="CSS 魔法" value="css-magic" />
            <el-option label="Node.js 开发" value="nodejs" />
          </el-select>
        </el-form-item>

        <el-form-item label="文章标签">
          <el-select
            v-model="publishForm.tags"
            multiple
            placeholder="请选择标签"
            clearable
            style="width: 100%"
          >
            <el-option label="Vue" value="vue" />
            <el-option label="React" value="react" />
            <el-option label="TypeScript" value="typescript" />
            <el-option label="JavaScript" value="javascript" />
            <el-option label="CSS" value="css" />
            <el-option label="HTML" value="html" />
            <el-option label="Node.js" value="nodejs" />
            <el-option label="Webpack" value="webpack" />
          </el-select>
        </el-form-item>

        <el-form-item label="文章摘要">
          <el-input
            v-model="publishForm.summary"
            type="textarea"
            :rows="3"
            placeholder="请输入文章摘要"
            maxlength="200"
            show-word-limit
          />
        </el-form-item>

        <el-form-item label="封面图片">
          <el-upload
            class="cover-uploader"
            action="#"
            :show-file-list="false"
            :before-upload="beforeCoverUpload"
          >
            <img v-if="publishForm.coverImage" :src="publishForm.coverImage" class="cover-image" />
            <el-icon v-else class="cover-uploader-icon"><Plus /></el-icon>
          </el-upload>
        </el-form-item>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="publishDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handlePublish">
            {{ (articleStatus === 'new' || articleStatus === 'draft') ? '确认发布' : '确认更新' }}
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import MarkdownRenderer from '../../components/common/MarkdownRenderer.vue'
import 'github-markdown-css/github-markdown.css'

const router = useRouter()
const route = useRoute()

// 从路由参数获取文章ID
const articleId = ref(route.params.id || null)
const publishDialogVisible = ref(false)

// 文章状态：new（新建）、draft（草稿）、published（已发布）
// 优先使用查询参数中的状态，其次通过路由路径判断
const articleStatus = ref(
  route.query.status || 'new'
)

// 文章表单数据
const articleForm = reactive({
  title: '',
  content: ''
})

// 发布表单数据
const publishForm = reactive({
  category: '',
  column: '',
  tags: [],
  summary: '',
  coverImage: ''
})

// 初始化markdown-it实例
// 实时编译Markdown - 使用MarkdownRenderer组件处理
const compiledMarkdown = computed(() => {
  return articleForm.content || '请输入Markdown内容...'
})

// 加载文章数据（如果是编辑现有文章）
onMounted(async () => {  
  // 如果有文章ID，加载文章数据
  if (articleId.value) {
    try {
      // 这里应该调用API加载文章数据
      // 模拟API调用，根据查询参数设置正确的文章状态
      const mockArticleData = {
        title: '示例文章标题',
        content: '# 这是示例内容',
        status: route.query.status || 'published' // 优先使用查询参数中的状态
      }
      
      // 更新表单数据
      Object.assign(articleForm, mockArticleData)
      
      // 只有当没有查询参数时才使用加载的文章数据中的状态
      if (!route.query.status) {
        articleStatus.value = mockArticleData.status
      }
    } catch (error) {
      ElMessage.error('加载文章失败')
      console.error('加载文章失败:', error)
    }
  }
})

// 内容输入处理
const handleContentInput = (event) => {
  // 可以在这里添加实时保存草稿等功能
}

// 显示发布对话框
const showPublishDialog = () => {
  if (!articleForm.title.trim()) {
    ElMessage.warning('请输入文章标题')
    return
  }
  if (!articleForm.content.trim()) {
    ElMessage.warning('请输入文章内容')
    return
  }
  publishDialogVisible.value = true
}

// 发布对话框关闭前处理
const handlePublishDialogClose = (done) => {
  ElMessageBox.confirm('确定要取消发布吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    done()
  }).catch(() => {
    // 取消关闭
  })
}

// 确认发布
const handlePublish = async () => {
  try {
    // 合并文章数据和发布数据
    const publishData = {
      ...articleForm,
      ...publishForm,
      status: 'published'
    }

    console.log('发布文章:', publishData)
    
    if (articleStatus.value === 'new') {
      console.log('创建并发布新文章')
      ElMessage.success('文章发布成功')
    } else {
      console.log('更新并发布文章')
      ElMessage.success('文章更新成功')
    }
    
    publishDialogVisible.value = false
    
    // 返回来源页面
    if (document.referrer.includes('/admin/')) {
      router.go(-1)
    } else {
      router.push('/admin/articles')
    }
  } catch (error) {
    ElMessage.error('发布失败')
    console.error('发布失败:', error)
  }
}

// 保存草稿
const handleSaveDraft = async () => {
  try {
    if (!articleForm.title.trim()) {
      ElMessage.warning('请输入文章标题')
      return
    }
    if (!articleForm.content.trim()) {
      ElMessage.warning('请输入文章内容')
      return
    }

    const draftData = {
      ...articleForm,
      status: 'draft'
    }

    console.log('保存草稿:', draftData)
    
    if (articleStatus.value === 'new') {
      console.log('创建新草稿')
      ElMessage.success('草稿保存成功')
    } else {
      console.log('更新草稿')
      ElMessage.success('草稿更新成功')
    }
    
    // 返回来源页面
    if (document.referrer.includes('/admin/')) {
      router.go(-1)
    } else {
      router.push('/admin/articles')
    }
  } catch (error) {
    ElMessage.error('保存失败')
    console.error('保存失败:', error)
  }
}

// 取消编辑
const handleCancel = () => {
  if (document.referrer.includes('/admin/')) {
    router.go(-1)
  } else {
    router.push('/admin/articles')
  }
}

// 封面图片上传前处理
const beforeCoverUpload = (file) => {
  const isJPGOrPNG = file.type === 'image/jpeg' || file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isJPGOrPNG) {
    ElMessage.error('封面图片只能是 JPG 或 PNG 格式!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('封面图片大小不能超过 2MB!')
    return false
  }

  const reader = new FileReader()
  reader.onload = (e) => {
    publishForm.coverImage = e.target.result
  }
  reader.readAsDataURL(file)
  
  return false
}

// 组件挂载时加载文章数据（如果是编辑模式）
onMounted(() => {
  if (articleId.value) {
    loadArticleData()
  }
})

// 加载文章数据
const loadArticleData = async () => {
  try {
    console.log('加载文章数据:', articleId.value)
    // 模拟数据
    Object.assign(articleForm, {
      title: '示例文章标题',
      content: '# 这是文章内容\n\n这里可以使用Markdown格式编写文章内容。'
    })
    Object.assign(publishForm, {
      category: 'frontend',
      tags: ['vue', 'javascript'],
      summary: '这是文章的摘要内容',
      coverImage: ''
    })
  } catch (error) {
    ElMessage.error('加载文章失败')
    console.error('加载文章失败:', error)
  }
}

const handleTabKey = (event) => {
  const textarea = event.target
  const start = textarea.selectionStart
  const end = textarea.selectionEnd
  
  // 在光标位置插入4个空格
  articleForm.content = 
    articleForm.content.substring(0, start) +
    '    ' + // 4个空格
    articleForm.content.substring(end)
  
  // 设置光标位置
  nextTick(() => {
    textarea.selectionStart = textarea.selectionEnd = start + 4
  })
}
</script>

<style scoped>
.article-editor {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: #fff;
}

.editor-header {
  display: flex;
  align-items: center;
  padding: 16px 24px;
  border-bottom: 1px solid #e8e8e8;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
}

.title-section {
  flex: 1;
  margin-right: 24px;
}

.title-input :deep(.el-input__inner) {
  font-size: 20px;
  font-weight: 600;
  border: none;
  padding: 0;
  height: 40px;
  line-height: 40px;
}

.title-input :deep(.el-input__inner):focus {
  box-shadow: none;
}

.action-section {
  display: flex;
  gap: 12px;
}

.editor-main {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.editor-panel,
.preview-panel {
  flex: 1;
  height: 100%;
  overflow: hidden;
}

.editor-panel {
  border-right: 1px solid #e8e8e8;
}

.markdown-editor {
  width: 100%;
  height: 100%;
  border: none;
  padding: 24px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 14px;
  line-height: 1.6;
  resize: none;
  outline: none;
  background: #fafafa;
}

.markdown-preview {
    padding: 24px;
}

.markdown-editor:focus {
  background: #fff;
}

.cover-uploader {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  width: 150px;
  height: 100px;
}

.cover-uploader:hover {
  border-color: #409EFF;
}

.cover-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 150px;
  height: 100px;
  text-align: center;
  line-height: 100px;
}

.cover-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>