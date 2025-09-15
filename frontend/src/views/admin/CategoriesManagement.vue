<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Document } from '@element-plus/icons-vue'

// 模拟分类数据
const categories = ref([
  { id: 1, name: '前端开发', description: '前端技术相关文章', articleCount: 25 },
  { id: 2, name: '后端开发', description: '后端技术相关文章', articleCount: 18 },
  { id: 3, name: '数据库', description: '数据库技术相关文章', articleCount: 12 },
  { id: 4, name: 'DevOps', description: '运维部署相关文章', articleCount: 8 }
])

const loading = ref(false)
const router = useRouter()

// 对话框状态
const createDialogVisible = ref(false)
const editDialogVisible = ref(false)

// 表单数据
const formData = reactive({
  id: null,
  name: '',
  description: ''
})

// 当前编辑的分类
const currentEditCategory = ref(null)

// 表单验证规则
const formRules = {
  name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' },
    { min: 2, max: 20, message: '分类名称长度在2到20个字符', trigger: 'blur' }
  ],
  description: [
    { max: 100, message: '分类简介不能超过100个字符', trigger: 'blur' }
  ]
}

// 加载分类数据
const loadCategories = async () => {
  loading.value = true
  try {
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 500))
  } catch (error) {
    ElMessage.error('加载分类失败')
  } finally {
    loading.value = false
  }
}

// 初始化加载
loadCategories()

// 打开创建分类对话框
const openCreateDialog = () => {
  // 重置表单
  formData.id = null
  formData.name = ''
  formData.description = ''
  createDialogVisible.value = true
}

// 打开编辑分类对话框
const openEditDialog = (category) => {
  currentEditCategory.value = category
  formData.id = category.id
  formData.name = category.name
  formData.description = category.description
  editDialogVisible.value = true
}

// 提交表单（创建或编辑）
const submitForm = async (isEdit = false) => {
  try {
    // 表单验证
    if (!formData.name.trim()) {
      ElMessage.warning('请输入分类名称')
      return
    }

    if (isEdit) {
      // 编辑现有分类
      const index = categories.value.findIndex(c => c.id === formData.id)
      if (index !== -1) {
        categories.value[index] = {
          ...categories.value[index],
          name: formData.name.trim(),
          description: formData.description.trim()
        }
        ElMessage.success('分类更新成功')
      }
    } else {
      // 创建新分类
      const newCategory = {
        id: Math.max(...categories.value.map(c => c.id)) + 1,
        name: formData.name.trim(),
        description: formData.description.trim(),
        articleCount: 0
      }
      categories.value.push(newCategory)
      ElMessage.success('分类添加成功')
    }

    // 关闭对话框
    createDialogVisible.value = false
    editDialogVisible.value = false
  } catch (error) {
    ElMessage.error(isEdit ? '更新分类失败' : '添加分类失败')
  }
}

// 取消表单
const cancelForm = (isEdit = false) => {
  if (isEdit) {
    editDialogVisible.value = false
  } else {
    createDialogVisible.value = false
  }
}

// 跳转到分类文章列表
const gotoCategoryArticles = (category) => {
  router.push(`/admin/categories/category-articles?categoryid=${category.id}`)
}

// 编辑分类
const editCategory = (category, event) => {
  event.stopPropagation() // 阻止冒泡到卡片点击事件
  openEditDialog(category)
}

// 删除分类
const deleteCategory = async (category, event) => {
  event.stopPropagation() // 阻止冒泡到卡片点击事件
  try {
    const confirmed = await ElMessageBox.confirm(
      `确定要删除分类"${category.name}"吗？此操作不可恢复。`,
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    if (confirmed) {
      // 调用API删除分类
      categories.value = categories.value.filter(c => c.id !== category.id)
      ElMessage.success('分类删除成功')
    }
  } catch (error) {
    // 用户取消删除
  }
}
</script>

<template>
  <div class="categories-management">
    <div class="page-header">
      <h2 class="page-title">分类管理</h2>
      <div class="action-buttons">
        <el-button type="primary" icon="Plus" @click="openCreateDialog">添加分类</el-button>
      </div>
    </div>

    <!-- 分类卡片列表 -->
    <div class="categories-grid" v-loading="loading">
      <div 
        v-for="category in categories" 
        :key="category.id" 
        class="category-card"
      >
        <div class="card-header" @click="gotoCategoryArticles(category)">
          <h3 class="category-name">{{ category.name }}</h3>
          <div class="article-count">
            <el-icon><document /></el-icon>
            {{ category.articleCount }} 篇文章
          </div>
        </div>
        <p class="category-description">{{ category.description }}</p>
        
        <!-- 操作按钮 -->
        <div class="action-buttons" @click.stop>
          <el-button 
            size="large" 
            type="primary" 
            link 
            @click="editCategory(category, $event)"
          >
            编辑
          </el-button>
          <el-button 
            size="large" 
            type="danger" 
            link 
            @click="deleteCategory(category, $event)"
          >
            删除
          </el-button>
        </div>
      </div>
    </div>

    <!-- 添加分类对话框 -->
    <el-dialog
      v-model="createDialogVisible"
      title="添加分类"
      width="500px"
      :before-close="() => cancelForm(false)"
    >
      <el-form :model="formData" label-width="80px">
        <el-form-item label="分类名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入分类名称"
            maxlength="20"
            show-word-limit
          />
        </el-form-item>
        <el-form-item label="分类简介" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入分类简介"
            maxlength="100"
            show-word-limit
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelForm(false)">取消</el-button>
          <el-button type="primary" @click="submitForm(false)">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 编辑分类对话框 -->
    <el-dialog
      v-model="editDialogVisible"
      title="编辑分类"
      width="500px"
      :before-close="() => cancelForm(true)"
    >
      <el-form :model="formData" label-width="80px">
        <el-form-item label="分类名称" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入分类名称"
            maxlength="20"
            show-word-limit
          />
        </el-form-item>
        <el-form-item label="分类简介" prop="description">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="3"
            placeholder="请输入分类简介"
            maxlength="100"
            show-word-limit
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancelForm(true)">取消</el-button>
          <el-button type="primary" @click="submitForm(true)">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.categories-management {
  padding: 20px;
}

.page-header {
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

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e8e8e8;
}


.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 16px;
  margin-top: 20px;
}

.category-card {
  background: #fff;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  cursor: pointer;
  transition: all 0.3s ease;
  border-left: 4px solid var(--el-color-primary);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.category-name {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2d3d;
}

.category-name:hover {
  color: var(--el-color-primary);
}

.article-count {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #8c8c8c;
  font-size: 14px;
  font-weight: 500;
}

.category-description {
  margin: 0 0 16px 0;
  color: #5e6d82;
  line-height: 1.5;
  font-size: 14px;
  text-align: left;
}

.action-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  padding-top: 12px;
  border-top: 1px solid #f0f0f0;
  margin-top: 12px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .categories-grid {
    grid-template-columns: 1fr;
  }
  
  .card-header {
    flex-direction: column;
    gap: 8px;
  }
}
</style>
