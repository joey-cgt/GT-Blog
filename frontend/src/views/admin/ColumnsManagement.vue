<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  createColumn,
  getColumnList,
  updateColumn,
  deleteColumn
} from '@/api/column'

const loading = ref(false)
const columns = ref([])

// 初始化加载数据
onMounted(() => {
  loadColumns()
})

const loadColumns = async () => {
  try {
    loading.value = true
    const { data } = await getColumnList({
      page: currentPage.value,
      pageSize: pageSize.value
    })
    columns.value = data.items.map(item => ({
      id: item.id,
      name: item.name,
      description: item.description,
      cover: item.coverUrl || '',
      createdAt: item.createTime,
      articleCount: item.articleCount,
      totalViews: item.viewCount,
      isTop: item.isTop,
      updateTime: item.updateTime
    }))
    total.value = data.total
    currentPage.value = data.page
    pageSize.value = data.pageSize
  } catch (error) {
    ElMessage.error('获取专栏列表失败')
  } finally {
    loading.value = false
  }
}

const searchQuery = ref('')
const activeMenu = ref(null)
const dialogVisible = ref(false)
const editDialogVisible = ref(false)
const router = useRouter()
// 分页状态
const currentPage = ref(1)
const pageSize = ref(5)
const total = ref(0)
const newColumn = reactive({
  name: '',
  description: '',
  cover: ''
})
const editingColumn = reactive({
  id: null,
  name: '',
  description: '',
  cover: ''
})

// 过滤专栏列表
const filteredColumns = computed(() => {
  if (!searchQuery.value) return columns.value
  return columns.value.filter(column => 
    column.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
    column.description.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// 分页后的专栏列表
const paginatedColumns = computed(() => {
  return columns.value
})

// 总页数
const totalPages = computed(() => {
  return Math.ceil(total.value / pageSize.value)
})

// 操作菜单
const handleToggleMenu = (columnId, event) => {
  event.stopPropagation()
  activeMenu.value = activeMenu.value === columnId ? null : columnId
}

// 删除了置顶相关功能

// 修改专栏 - 打开弹窗
const handleEditColumn = (column) => {
  // 设置要编辑的专栏信息
  editingColumn.id = column.id
  editingColumn.name = column.name
  editingColumn.description = column.description
  editingColumn.cover = column.cover
  
  // 打开编辑对话框
  editDialogVisible.value = true
  activeMenu.value = null
}

// 确认修改专栏
const handleConfirmEditColumn = async () => {
  if (!editingColumn.name.trim()) {
    ElMessage.warning('请输入专栏名称')
    return
  }

  if (!editingColumn.description.trim()) {
    ElMessage.warning('请输入专栏简介')
    return
  }

  try {
    await updateColumn(editingColumn.id, {
      id: editingColumn.id,
      name: editingColumn.name.trim(),
      description: editingColumn.description.trim(),
      coverUrl: editingColumn.cover.trim() || ''
    })
    
    // 更新本地数据
    const column = columns.value.find(c => c.id === editingColumn.id)
    if (column) {
      column.name = editingColumn.name.trim()
      column.description = editingColumn.description.trim()
      column.cover = editingColumn.cover.trim() || column.cover
    }
    
    editDialogVisible.value = false
    ElMessage.success('专栏已更新')
  } catch (error) {
    ElMessage.error(error.message || '修改专栏失败')
  }
}

// 删除专栏
const handleDeleteColumn = async (column) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除专栏"${column.name}"吗？此操作不可恢复。`,
      '删除确认',
      { type: 'warning' }
    )
    
    await deleteColumn(column.id)
    columns.value = columns.value.filter(c => c.id !== column.id)
    ElMessage.success('专栏已删除')
    try {
      await loadColumns()
    } catch (error) {
      // 失败时保留临时数据
      console.error('刷新数据失败:', error)
    }
  } catch (error) {
    // 用户取消
  }
  activeMenu.value = null
}

// 新建专栏 - 打开弹窗
const handleOpenCreateDialog = () => {
  // 重置表单
  newColumn.name = ''
  newColumn.description = ''
  newColumn.cover = ''
  dialogVisible.value = true
}

// 确认创建专栏
const handleConfirmCreateColumn = async () => {
  if (!newColumn.name.trim()) {
    ElMessage.warning('请输入专栏名称')
    return
  }

  if (!newColumn.description.trim()) {
    ElMessage.warning('请输入专栏简介')
    return
  }

  try {
    const response = await createColumn({
      name: newColumn.name.trim(),
      description: newColumn.description.trim(),
      coverUrl: newColumn.cover.trim() || ''
    })
    
    if (!response || !response.columnId) {
      throw new Error('创建专栏失败：响应数据为空')
    }
    
    // 1. 先在前端添加临时数据
    const tempColumn = {
      id: response.columnId,
      name: newColumn.name.trim(),
      description: newColumn.description.trim(),
      cover: newColumn.cover.trim() || 'https://th.bing.com/th/id/OIP.dIrXao1MUcqjeluUf9m8XAHaEl?w=229&h=180&c=7&r=0&o=7&pid=1.7&rm=3',
      createdAt: new Date().toISOString().split('T')[0],
      articleCount: 0,
      totalViews: 0,
      isTop: false
    }
    columns.value.unshift(tempColumn)
    
    dialogVisible.value = false
    ElMessage.success('专栏创建成功')

    // 2. 静默刷新数据
    try {
      await loadColumns()
    } catch (error) {
      // 失败时保留临时数据
      console.error('刷新数据失败:', error)
    }
  } catch (error) {
    ElMessage.error(error.message || '创建专栏失败')
  }
}

// 取消创建
const handleCancelCreate = () => {
  dialogVisible.value = false
}

// 取消修改
const handleCancelEdit = () => {
  editDialogVisible.value = false
}

// 跳转到专栏文章列表
const handleGotoColumnArticles = (column) => {
  router.push(`/admin/columns/column-articles?columnid=${column.id}&name=${column.name}`)
}

// 处理页码变化
const handlePageChange = (page) => {
  currentPage.value = page
  loadColumns()
}

// 处理每页数量变化
const handleSizeChange = (size) => {
  pageSize.value = size
  currentPage.value = 1 // 重置到第一页
  loadColumns()
}
</script>

<template>
  <div class="columns-management">
    <div class="page-header">
      <h2 class="page-title">专栏管理</h2>
      
      <div class="action-section">
        <el-input
            v-model="searchQuery"
            placeholder="搜索专栏名称或介绍..."
            clearable
            class="search-input"
        >
            <template #prefix>
            <el-icon><search /></el-icon>
            </template>
        </el-input>
    
        <el-button type="primary" @click="handleOpenCreateDialog">
            <el-icon><plus /></el-icon>
            新建专栏
        </el-button>
      </div>
    </div>
    

    <!-- 专栏列表 -->
    <div class="columns-list">
      <div 
        v-for="column in paginatedColumns" 
        :key="column.id" 
        class="column-item"
      >
        <!-- 左侧封面 -->
        <div class="column-cover" @click="handleGotoColumnArticles(column)">
          <img :src="column.cover" :alt="column.name" class="cover-image">
        </div>

        <!-- 中间内容 -->
        <div class="column-content" @click="handleGotoColumnArticles(column)">
          <h3 class="column-name">{{ column.name }}</h3>
          <p class="column-description">{{ column.description }}</p>
          
          <div class="column-meta">
            <span class="meta-item">
              <el-icon><calendar /></el-icon>
              创建于 {{ column.createdAt }}
            </span>
            <span class="meta-item">
              <el-icon><document /></el-icon>
              {{ column.articleCount }} 篇文章
            </span>
            <span class="meta-item">
              <el-icon><view /></el-icon>
              {{ column.totalViews.toLocaleString() }} 浏览
            </span>
          </div>
        </div>

        <!-- 右侧操作菜单 -->
        <div class="column-actions">
          <el-button 
            circle 
            size="small" 
            @click="handleToggleMenu(column.id, $event)"
            class="menu-button"
          >
            <el-icon><more /></el-icon>
          </el-button>

          <!-- 下拉菜单 -->
          <div v-if="activeMenu === column.id" class="action-menu">
            <div class="menu-item" @click="handleEditColumn(column)">
              <el-icon><edit /></el-icon>
              修改专栏
            </div>
            <div class="menu-item delete" @click="handleDeleteColumn(column)">
              <el-icon><delete /></el-icon>
              删除专栏
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 分页组件 -->
    <div class="pagination-section" v-if="total > 0">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[5, 10, 20, 50]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </div>
  </div>

  <!-- 新建专栏弹窗 -->
  <el-dialog
    v-model="dialogVisible"
    title="新建专栏"
    width="500px"
    :close-on-click-modal="false"
  >
    <el-form :model="newColumn" label-width="80px">
      <el-form-item label="专栏名称" required>
        <el-input
          v-model="newColumn.name"
          placeholder="请输入专栏名称"
          maxlength="50"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="专栏简介" required>
        <el-input
          v-model="newColumn.description"
          type="textarea"
          :rows="3"
          placeholder="请输入专栏简介"
          maxlength="200"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="封面链接">
        <el-input
          v-model="newColumn.cover"
          placeholder="请输入封面图片链接（可选）"
        />
        <div class="cover-tip">提示：可以留空使用默认封面</div>
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleCancelCreate">取消</el-button>
        <el-button type="primary" @click="handleConfirmCreateColumn">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>
  
  <!-- 修改专栏弹窗 -->
  <el-dialog
    v-model="editDialogVisible"
    title="修改专栏"
    width="500px"
    :close-on-click-modal="false"
  >
    <el-form :model="editingColumn" label-width="80px">
      <el-form-item label="专栏名称" required>
        <el-input
          v-model="editingColumn.name"
          placeholder="请输入专栏名称"
          maxlength="50"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="专栏简介" required>
        <el-input
          v-model="editingColumn.description"
          type="textarea"
          :rows="3"
          placeholder="请输入专栏简介"
          maxlength="200"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="封面链接">
        <el-input
          v-model="editingColumn.cover"
          placeholder="请输入封面图片链接（可选）"
        />
        <div class="cover-tip">提示：可以留空使用默认封面</div>
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleCancelEdit">取消</el-button>
        <el-button type="primary" @click="handleConfirmEditColumn">
          保存
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped>
.columns-management {
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

.action-section {
  display: flex;
  gap: 16px;
  align-items: center;
}

.search-input {
  flex: 1;
  max-width: 400px;
}

.columns-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.column-item {
  display: flex;
  align-items: center;
  padding: 20px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  transition: all 0.3s ease;
  position: relative;
  cursor: pointer;
}

/* .column-item:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.12);
  transform: translateY(-2px);
} */

.column-cover {
  position: relative;
  flex-shrink: 0;
  margin-right: 20px;
}

.cover-image {
  width: 80px;
  height: 80px;
  border-radius: 8px;
  object-fit: cover;
  background: #f5f7fa;
}

.column-content {
  flex: 1;
  min-width: 0;
  text-align: left;
}

.column-name {
  margin: 0 0 8px 0;
  font-size: 18px;
  font-weight: 600;
  color: #1f2d3d;
}

.column-description {
  margin: 0 0 12px 0;
  color: #5e6d82;
  line-height: 1.5;
  font-size: 14px;
}

.column-meta {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #8c8c8c;
  font-size: 12px;
}

.column-actions {
  position: absolute;
  top: 20px;
  right: 20px;
  flex-shrink: 0;
}

.menu-button {
  border: none;
  background: transparent;
}

.menu-button:hover {
  background: #f5f7fa;
}

.action-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 9999; /* 确保菜单显示在最上层 */
  min-width: 120px;
  overflow: hidden;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 16px;
  cursor: pointer;
  font-size: 14px;
  color: #1f2d3d;
  transition: background 0.2s;
}

.menu-item:hover {
  background: #f5f7fa;
}

.menu-item.delete {
  color: var(--el-color-danger);
}

.menu-item.delete:hover {
  background: var(--el-color-danger-light-9);
}

.cover-tip {
  font-size: 12px;
  color: #8c8c8c;
  margin-top: 4px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.pagination-section {
  margin-top: 24px;
  display: flex;
  justify-content: center;
  background: #fff;
  padding: 16px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .action-section {
    flex-direction: column;
    align-items: stretch;
  }
  
  .search-input {
    max-width: none;
  }
  
  .column-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
  
  .column-cover {
    margin-right: 0;
  }
  
  .column-meta {
    gap: 12px;
  }
  
  .column-actions {
    position: absolute;
    top: 20px;
    right: 20px;
  }
}
</style>