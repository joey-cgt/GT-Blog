<template>
  <div class="tags-management">
    <div class="page-header">
      <h2 class="page-title">标签管理</h2>
      <div class="header-actions">
        <el-button type="primary" icon="Plus" @click="openCreateDialog">
          添加标签
        </el-button>
        <el-button type="info" icon="Setting" @click="openManageDialog">
          管理标签
        </el-button>
      </div>
    </div>
    <div class="tags-cloud">
      <div 
        v-for="tag in tags" 
        :key="tag.id"
        class="tag-item"
        :style="{
          fontSize: getTagSize(tag.articleCount) + 'px',
          opacity: getTagOpacity(tag.articleCount),
          transform: `scale(${1 + (tag.articleCount / 100) * 0.3})`
        }"
        @click="handleTagClick(tag)"
      >
        {{ tag.name }}
        <span class="article-count">({{ tag.articleCount }})</span>
      </div>
    </div>

    <!-- 创建标签对话框 -->
    <el-dialog
      v-model="createDialogVisible"
      title="创建新标签"
      width="500px"
      :before-close="cancelForm"
    >
      <el-form :model="tagForm" label-width="80px">
        <el-form-item label="标签名称" required>
          <el-input
            v-model="tagForm.name"
            placeholder="请输入标签名称"
            maxlength="20"
            show-word-limit
          />
        </el-form-item>
        <el-form-item label="标签描述">
          <el-input
            v-model="tagForm.description"
            type="textarea"
            :rows="3"
            placeholder="请输入标签描述（可选）"
            maxlength="100"
            show-word-limit
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelForm">取消</el-button>
          <el-button 
            type="primary" 
            @click="submitForm"
            :loading="formLoading"
          >
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 管理标签对话框 -->
    <el-dialog
      v-model="manageDialogVisible"
      title="管理标签"
      width="700px"
    >
      <el-table :data="tags" stripe>
        <el-table-column prop="name" label="标签名称" min-width="120" />
        <el-table-column prop="articleCount" label="文章数量" width="100" align="center">
          <template #default="{ row }">
            <el-tag size="small" :type="row.articleCount > 0 ? 'success' : 'info'">
              {{ row.articleCount }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" align="center">
          <template #default="{ row }">
            <el-button size="small" type="primary" @click="editTag(row)">
              编辑
            </el-button>
            <el-button size="small" type="danger" @click="deleteTag(row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="manageDialogVisible = false">关闭</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 编辑标签对话框 -->
    <el-dialog
      v-model="editDialogVisible"
      :title="`编辑标签: ${editingTag?.name || ''}`"
      width="500px"
      :before-close="cancelForm"
    >
      <el-form :model="tagForm" label-width="80px">
        <el-form-item label="标签名称" required>
          <el-input
            v-model="tagForm.name"
            placeholder="请输入标签名称"
            maxlength="20"
            show-word-limit
          />
        </el-form-item>
        <el-form-item label="标签描述">
          <el-input
            v-model="tagForm.description"
            type="textarea"
            :rows="3"
            placeholder="请输入标签描述（可选）"
            maxlength="100"
            show-word-limit
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelForm">取消</el-button>
          <el-button 
            type="primary" 
            @click="submitForm"
            :loading="formLoading"
          >
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'

const route = useRoute()
const router = useRouter()

// 对话框状态
const createDialogVisible = ref(false)
const manageDialogVisible = ref(false)
const editDialogVisible = ref(false)
const formLoading = ref(false)
const editingTag = ref(null)

// 表单数据
const tagForm = ref({
  name: '',
  description: ''
})

// 模拟标签数据
const tags = ref([
  { id: 1, name: 'Vue.js', articleCount: 45 },
  { id: 2, name: 'React', articleCount: 32 },
  { id: 3, name: 'JavaScript', articleCount: 78 },
  { id: 4, name: 'TypeScript', articleCount: 28 },
  { id: 5, name: 'Node.js', articleCount: 36 },
  { id: 6, name: 'CSS', articleCount: 52 },
  { id: 7, name: 'HTML', articleCount: 41 },
  { id: 8, name: '前端框架', articleCount: 23 },
  { id: 9, name: '后端开发', articleCount: 19 },
  { id: 10, name: '数据库', articleCount: 27 },
  { id: 11, name: '算法', articleCount: 15 },
  { id: 12, name: '设计模式', articleCount: 12 }
])

// 计算标签大小（根据文章数量）
const getTagSize = (count) => {
  const minSize = 14
  const maxSize = 32
  const minCount = Math.min(...tags.value.map(t => t.articleCount))
  const maxCount = Math.max(...tags.value.map(t => t.articleCount))
  
  if (maxCount === minCount) return (minSize + maxSize) / 2
  
  return minSize + ((count - minCount) / (maxCount - minCount)) * (maxSize - minSize)
}

// 计算标签透明度（根据文章数量）
const getTagOpacity = (count) => {
  const minOpacity = 0.6
  const maxOpacity = 1
  const minCount = Math.min(...tags.value.map(t => t.articleCount))
  const maxCount = Math.max(...tags.value.map(t => t.articleCount))
  
  if (maxCount === minCount) return maxOpacity
  
  return minOpacity + ((count - minCount) / (maxCount - minCount)) * (maxOpacity - minOpacity)
}

// 打开创建标签对话框
const openCreateDialog = () => {
  createDialogVisible.value = true
  resetForm()
}

// 打开管理标签对话框
const openManageDialog = () => {
  manageDialogVisible.value = true
}

// 编辑标签
const editTag = (tag) => {
  editingTag.value = tag
  tagForm.value = {
    name: tag.name,
    description: tag.description || ''
  }
  editDialogVisible.value = true
}

// 删除标签
const deleteTag = async (tag) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除标签"${tag.name}"吗？此操作不可恢复。`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    // 删除标签
    tags.value = tags.value.filter(t => t.id !== tag.id)
    ElMessage.success('标签删除成功')
  } catch (error) {
    // 用户取消删除
  }
}

// 重置表单
const resetForm = () => {
  tagForm.value = {
    name: '',
    description: ''
  }
}

// 提交表单
const submitForm = async () => {
  if (!tagForm.value.name.trim()) {
    ElMessage.warning('请输入标签名称')
    return
  }

  formLoading.value = true
  try {
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    if (editingTag.value) {
      // 编辑现有标签
      const index = tags.value.findIndex(t => t.id === editingTag.value.id)
      if (index !== -1) {
        tags.value[index] = {
          ...tags.value[index],
          name: tagForm.value.name.trim(),
          description: tagForm.value.description.trim()
        }
      }
      editDialogVisible.value = false
      ElMessage.success('标签更新成功')
    } else {
      // 创建新标签
      const newTag = {
        id: Math.max(...tags.value.map(t => t.id)) + 1,
        name: tagForm.value.name.trim(),
        description: tagForm.value.description.trim(),
        articleCount: 0,
        createdAt: new Date().toISOString().split('T')[0]
      }
      
      tags.value.push(newTag)
      createDialogVisible.value = false
      ElMessage.success('标签创建成功')
    }
    
  } catch (error) {
    ElMessage.error(editingTag.value ? '更新标签失败' : '创建标签失败')
  } finally {
    formLoading.value = false
    editingTag.value = null
  }
}

// 取消表单
const cancelForm = () => {
  createDialogVisible.value = false
  editDialogVisible.value = false
  editingTag.value = null
}

// 处理标签点击
const handleTagClick = (tag) => {
  console.log('点击标签:', tag.name)
  // 这里可以添加跳转到标签文章列表的功能
  router.push(`/admin/tags/tag-articles?tagid=${tag.id}`)
}
</script>

<style scoped>
.tags-management {
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

.page-title {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  color: #303133;
}

.tags-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
  padding: 30px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 20px;
  min-height: 400px;
  align-items: center;
  position: relative;
  overflow: hidden;
}

.tags-cloud::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(255,255,255,0.1) 1px, transparent 1px);
  background-size: 30px 30px;
  animation: float 20s infinite linear;
  pointer-events: none;
}

@keyframes float {
  0% { transform: translate(0, 0) rotate(0deg); }
  100% { transform: translate(-30px, -30px) rotate(360deg); }
}

.tag-item {
  padding: 12px 24px;
  color: white;
  border-radius: 30px;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
  display: inline-block;
  margin: 6px;
  font-weight: 600;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  position: relative;
  overflow: hidden;
  border: 2px solid rgba(255, 255, 255, 0.2);
}

.tag-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.6s ease;
}

.tag-item:hover::before {
  left: 100%;
}

.tag-item:hover {
  transform: translateY(-4px) scale(1.15);
  box-shadow: 0 12px 35px rgba(0, 0, 0, 0.25);
  z-index: 10;
}

/* 为不同标签添加不同的渐变色 */
.tag-item:nth-child(6n+1) {
  background: linear-gradient(45deg, #ff6b6b, #ee5a52);
}
.tag-item:nth-child(6n+2) {
  background: linear-gradient(45deg, #4ecdc4, #44a08d);
}
.tag-item:nth-child(6n+3) {
  background: linear-gradient(45deg, #45b7d1, #96c93d);
}
.tag-item:nth-child(6n+4) {
  background: linear-gradient(45deg, #f093fb, #f5576c);
}
.tag-item:nth-child(6n+5) {
  background: linear-gradient(45deg, #5ee7df, #b490ca);
}
.tag-item:nth-child(6n+6) {
  background: linear-gradient(45deg, #d299c2, #fef9d7);
}

.article-count {
  font-size: 0.75em;
  opacity: 0.95;
  margin-left: 6px;
  background: rgba(255, 255, 255, 0.2);
  padding: 2px 8px;
  border-radius: 12px;
  backdrop-filter: blur(10px);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .tags-cloud {
    gap: 12px;
    padding: 15px;
  }
  
  .tag-item {
    padding: 6px 12px;
    font-size: 12px;
  }
}
</style>