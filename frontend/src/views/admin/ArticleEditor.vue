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
        <template v-if="articleStatus === 'new' || articleStatus === 0">
          <el-button type="primary" @click="handleCreateDraft">
            保存草稿
          </el-button>
          <el-button type="success" @click="showPublishDialog">
            发布文章
          </el-button>
        </template>
        <!-- 已发布文章：只显示更新文章 -->
        <template v-else-if="articleStatus === 1">
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
      :title="(articleStatus === 'new' || articleStatus === 0) ? '发布设置' : '更新文章'"
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
            <el-option
              v-for="category in categories"
              :key="category.id"
              :label="category.name"
              :value="category.name"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="加入专栏">
          <el-select
            v-model="publishForm.column"
            placeholder="请选择专栏"
            clearable
            style="width: 100%"
          >
            <el-option
              v-for="column in columns"
              :key="column.id"
              :label="column.name"
              :value="column.name"
            />
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
            <el-option
              v-for="tag in tags"
              :key="tag.id"
              :label="tag.name"
              :value="tag.name"
            />
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
          <el-input
            v-model="publishForm.coverImage"
            placeholder="请输入图片URL"
            clearable
            style="width: 100%; margin-bottom: 10px;"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="publishDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handlePublish">
            {{ (articleStatus === 'new' || articleStatus === 0) ? '确认发布' : '确认更新' }}
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed, nextTick, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import MarkdownRenderer from '../../components/common/MarkdownRenderer.vue'
import { createAndPublishArticle, updatePublishedArticle, createDraft, updateDraft, getArticleDetail, publishDraft } from '../../api/article.js'
import { getCategoryList } from '../../api/category.js'
import { getTags } from '../../api/tag.js'
import { getColumns } from '../../api/column.js'
import 'github-markdown-css/github-markdown.css'

const router = useRouter()
const route = useRoute()

// 从路由参数获取文章ID
const articleId = ref(route.params.id || null)
const publishDialogVisible = ref(false)

// 文章状态：'new'（新建）、0（草稿）、1（已发布）
// 优先使用查询参数中的状态，其次通过路由路径判断
const articleStatus = ref(
  route.query.status || 'new'
)

// 文章表单数据
const articleForm = reactive({
  title: '',
  content: ''
})

// 真实数据列表
const categories = ref([])
const tags = ref([])
const columns = ref([])

// 数据映射关系（从ID到显示名称和从显示名称到ID）
const categoryMap = reactive({})
const reverseCategoryMap = reactive({})
const columnMap = reactive({})
const reverseColumnMap = reactive({})
const tagMap = reactive({})
const reverseTagMap = reactive({})

// 发布表单数据（显示名称）
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

// 获取分类数据
const fetchCategories = async () => {
  try {
    const response = await getCategoryList({ page: 1, pageSize: 100 })
    categories.value = response.data?.items || []
    // 构建映射关系
    categories.value.forEach(category => {
      categoryMap[category.name] = category.id.toString()
      reverseCategoryMap[category.id.toString()] = category.name
    })
  } catch (error) {
    ElMessage.error('获取分类数据失败')
    console.error('获取分类数据失败:', error)
  }
}

// 获取标签数据
const fetchTags = async () => {
  try {
    const response = await getTags({ page: 1, pageSize: 100 })
    tags.value = response.data?.items || []
    // 构建映射关系
    tags.value.forEach(tag => {
      tagMap[tag.name] = tag.id.toString()
      reverseTagMap[tag.id.toString()] = tag.name
    })
  } catch (error) {
    ElMessage.error('获取标签数据失败')
    console.error('获取标签数据失败:', error)
  }
}

// 获取专栏数据
const fetchColumns = async () => {
  try {
    const response = await getColumns({ page: 1, pageSize: 100 })
    columns.value = response.data?.items || []
    // 构建映射关系
    columns.value.forEach(column => {
      columnMap[column.name] = column.id.toString()
      reverseColumnMap[column.id.toString()] = column.name
    })
  } catch (error) {
    ElMessage.error('获取专栏数据失败')
    console.error('获取专栏数据失败:', error)
  }
}

// 初始化数据
const initData = async () => {
  await Promise.all([
    fetchCategories(),
    fetchTags(),
    fetchColumns()
  ])
}

// 加载文章数据（如果是编辑现有文章）
onMounted(async () => {  
  // 先初始化分类、标签和专栏数据
  await initData()
  
  // 如果有文章ID，加载文章数据
  if (articleId.value) {
    await loadArticleData()
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
    // 构建发布数据，只包含有值的字段
    const publishData = {
      title: articleForm.title,
      content: articleForm.content
    }

    // 如果是编辑现有文章（非新建），添加文章ID
    if (articleId.value) {
      publishData.id = articleId.value
    }

    // 只添加有值的可选字段，并将显示名称转换为ID
    if (publishForm.summary.trim()) {
      publishData.abstract = publishForm.summary
    }
    if (publishForm.coverImage.trim()) {
      publishData.coverUrl = publishForm.coverImage
    }
    if (publishForm.category.trim()) {
      publishData.categoryID = convertDisplayNameToId(publishForm.category, categoryMap)
    }
    if (publishForm.column.trim()) {
      publishData.columnID = convertDisplayNameToId(publishForm.column, columnMap)
    }
    if (publishForm.tags.length > 0) {
      publishData.tagIDs = publishForm.tags.map(tag => convertDisplayNameToId(tag, tagMap))
    }

    console.log('发布文章数据（转换后）:', publishData)
    console.log('原始表单数据:', publishForm)
    
    let response
    if (articleStatus.value === 'new') {
      // 创建并发布新文章
      response = await createAndPublishArticle(publishData)
      console.log('创建并发布新文章成功:', response)
      ElMessage.success('文章发布成功')
    } else if (articleStatus.value === 0) {
      // 发布草稿
      response = await publishDraft(articleId.value, publishData)
      console.log('发布草稿成功:', response)
      ElMessage.success('草稿发布成功')
    } else {
      // 更新已发布文章
      response = await updatePublishedArticle(articleId.value, publishData)
      console.log('更新文章成功:', response)
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
const handleCreateDraft = async () => {
  try {
    if (!articleForm.title.trim()) {
      ElMessage.warning('请输入文章标题')
      return
    }
    if (!articleForm.content.trim()) {
      ElMessage.warning('请输入文章内容')
      return
    }

    // 构建草稿数据，只包含有值的字段
    const draftData = {
      title: articleForm.title,
      content: articleForm.content
    }

    // 只添加有值的可选字段，并将显示名称转换为ID
    if (publishForm.summary.trim()) {
      draftData.abstract = publishForm.summary
    }
    if (publishForm.coverImage.trim()) {
      draftData.coverUrl = publishForm.coverImage
    }
    if (publishForm.category.trim()) {
      draftData.categoryID = convertDisplayNameToId(publishForm.category, categoryMap)
    }
    if (publishForm.column.trim()) {
      draftData.columnID = convertDisplayNameToId(publishForm.column, columnMap)
    }
    if (publishForm.tags.length > 0) {
      draftData.tagIDs = publishForm.tags.map(tag => convertDisplayNameToId(tag, tagMap))
    }

    console.log('保存草稿数据（转换后）:', draftData)
    console.log('原始表单数据:', publishForm)
    
    let response
    if (articleStatus.value === 'new') {
      // 创建新草稿
      response = await createDraft(draftData)
      console.log('创建新草稿成功:', response)
      ElMessage.success('草稿保存成功')
    } else {
      // 更新草稿
      response = await updateDraft(articleId.value, draftData)
      console.log('更新草稿成功:', response)
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

// 监听封面图片URL变化，可选：添加简单的URL验证
watch(
  () => publishForm.coverImage,
  (newVal) => {
    if (newVal && !isValidImageUrl(newVal)) {
      ElMessage.warning('图片URL格式可能不正确，请确保是有效的图片链接')
    }
  }
)

// 简单的图片URL验证函数
const isValidImageUrl = (url) => {
  // 基本的URL格式验证
  try {
    new URL(url)
    // 检查是否以常见图片扩展名结尾或包含图片MIME类型
    const imageExtensions = /\.(jpeg|jpg|gif|png|bmp|webp)$/i
    const dataUrlRegex = /^data:image\/(jpeg|jpg|gif|png|bmp|webp);base64,/i
    return imageExtensions.test(url) || dataUrlRegex.test(url)
  } catch (e) {
    return false
  }
}

// 将显示名称转换为ID的辅助函数（保留以兼容现有代码）
const convertDisplayNameToId = (displayName, map) => {
  return map[displayName] || displayName
}

// 加载文章数据
const loadArticleData = async () => {
  try {
    console.log('加载文章数据:', articleId.value)
    
    // 调用后端API获取文章详情
    const response = await getArticleDetail(articleId.value)
    const articleData = response.data
    
    // 更新文章表单数据
    Object.assign(articleForm, {
      title: articleData.title || '',
      content: articleData.content || ''
    })
    
    // 更新发布表单数据（将ID转换为显示名称）
    Object.assign(publishForm, {
      category: reverseCategoryMap[articleData.category?.categoryId?.toString()] || '',
      column: reverseColumnMap[articleData.column?.columnId?.toString()] || '',
      tags: articleData.tags ? articleData.tags.map(tag => tag.tagName || '') : [],
      summary: articleData.abstract || '',
      coverImage: articleData.coverUrl || ''
    })
    
    // 过滤掉可能的空标签
    publishForm.tags = publishForm.tags.filter(tag => tag)
    
    // 设置文章状态（从后端获取）
    if (articleData.status !== undefined) {
      articleStatus.value = articleData.status
    }
    
    console.log('文章数据加载成功:', articleData)
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

.cover-preview {
  border: 1px solid #e8e8e8;
  border-radius: 6px;
  padding: 10px;
  width: 150px;
  background: #fafafa;
}

.cover-placeholder {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  padding: 30px 10px;
  width: 150px;
  text-align: center;
  color: #909399;
  font-size: 12px;
}

.cover-image {
  width: 100%;
  height: 100px;
  object-fit: cover;
  border-radius: 4px;
  margin-bottom: 8px;
}
</style>