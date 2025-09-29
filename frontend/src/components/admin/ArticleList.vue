<template>
  <div class="article-list">
    <!-- 搜索和筛选 -->
    <div class="filter-section">
      <el-input
        v-model="searchKeyword"
        placeholder="搜索文章标题..."
        style="width: 300px; margin-right: 16px;"
        clearable
        @input="handleSearch"
      >
        <template #prefix>
          <el-icon><Search /></el-icon>
        </template>
      </el-input>
      
      <el-select
        v-model="filterCategory"
        placeholder="选择分类"
        style="width: 200px; margin-right: 16px;"
        clearable
        @change="handleFilterChange"
      >
        <el-option
          v-for="category in categories"
          :key="category.id"
          :label="category.name"
          :value="category.name"
        />
      </el-select>
      
      <el-select
        v-model="filterTag"
        placeholder="选择标签"
        style="width: 200px; margin-right: 16px;"
        clearable
        @change="handleFilterChange"
      >
        <el-option
          v-for="tag in tags"
          :key="tag.id"
          :label="tag.name"
          :value="tag.name"
        />
      </el-select>

      <el-select
        v-model="filterColumn"
        placeholder="选择专栏"
        style="width: 200px;"
        clearable
        @change="handleFilterChange"
      >
        <el-option
          v-for="column in columns"
          :key="column.id"
          :label="column.name"
          :value="column.name"
        />
      </el-select>
    </div>

    <!-- 文章列表 -->
    <el-table
      :data="filteredArticles"
      v-loading="loading"
      style="width: 100%; margin-top: 20px;"
      stripe
    >
      <el-table-column label="文章标题" min-width="300">
        <template #default="{ row }">
          <div class="article-title">
            <span class="title-text">{{ row.title }}</span>
            <el-tag
              v-if="row.status === 'draft'"
              size="small"
              type="warning"
              style="margin-left: 8px;"
            >
              草稿
            </el-tag>
          </div>
        </template>
      </el-table-column>
      
      <el-table-column label="分类" width="120">
        <template #default="{ row }">
          <span class="category">{{ row.category.categoryName }}</span>
        </template>
      </el-table-column>
      
      <el-table-column label="标签" width="150">
        <template #default="{ row }">
          <div class="tags">
            <el-tag
              v-for="tag in row.tags.slice(0, 2)"
              :key="tag"
              size="small"
              style="margin-right: 4px; margin-bottom: 4px;"
            >
              {{ tag.tagName }}
            </el-tag>
            <el-tag v-if="row.tags.length > 2" size="small">+{{ row.tags.length - 2 }}</el-tag>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="专栏" width="180">
        <template #default="{ row }">
          <span class="column">{{ row.column.columnName }}</span>
        </template>
      </el-table-column>
      
      <el-table-column label="发布时间" width="150" v-if="props.articles[0]?.publishTime">
        <template #default="{ row }">
          <span class="time">{{ row.publishTime || row.createTime }}</span>
        </template>
      </el-table-column>
      
      <el-table-column label="创建时间" width="150" v-else>
        <template #default="{ row }">
          <span class="time">{{ row.createTime || row.publishTime }}</span>
        </template>
      </el-table-column>
      
      <el-table-column label="浏览量" width="80" align="center" v-if="props.articles[0]?.views">
        <template #default="{ row }">
          <span class="views">{{ row.views || 0 }}</span>
        </template>
      </el-table-column>
      
      <el-table-column label="点赞数" width="80" align="center" v-if="props.articles[0]?.likes">
        <template #default="{ row }">
          <span class="likes">{{ row.likes || 0 }}</span>
        </template>
      </el-table-column>
      
      <el-table-column label="评论数" width="80" align="center" v-if="props.articles[0]?.comments">
        <template #default="{ row }">
          <span class="comments">{{ row.comments || 0 }}</span>
        </template>
      </el-table-column>
      
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <div class="action-buttons">
            <el-button
              size="small"
              type="primary"
              link
              @click="handleEdit(row)"
            >
              编辑
            </el-button>
            <el-button
              size="small"
              type="danger"
              link
              @click="handleDelete(row)"
            >
              删除
            </el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination-section" v-if="props.pagination.total > 0">
      <el-pagination
        v-model:current-page="props.pagination.currentPage"
        v-model:page-size="props.pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="props.pagination.total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { getCategoryList } from '@/api/category'
import { getTags } from '@/api/tag'
import { getColumns } from '@/api/column'

const props = defineProps({
  articles: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  },
  pagination: {
    type: Object,
    default: () => ({
      currentPage: 1,
      pageSize: 10,
      total: 0
    })
  }
})

const emit = defineEmits(['page-change', 'edit', 'delete', 'change-status'])

// 搜索关键词
const searchKeyword = ref('')
// 筛选分类
const filterCategory = ref('')
// 筛选标签
const filterTag = ref('')
// 筛选专栏
const filterColumn = ref('')

// 分类列表
const categories = ref([])
// 标签列表
const tags = ref([])
// 专栏列表
const columns = ref([])

// 获取分类数据
const fetchCategories = async () => {
  try {
    const response = await getCategoryList({})
    categories.value = response.data?.items || []
  } catch (error) {
    console.error('获取分类数据失败:', error)
  }
}

// 获取标签数据
const fetchTags = async () => {
  try {
    const response = await getTags({ page: 1, pageSize: 100 })
    tags.value = response.data?.items || []
  } catch (error) {
    console.error('获取标签数据失败:', error)
  }
}

// 获取专栏数据
const fetchColumns = async () => {
  try {
    const response = await getColumns({ page: 1, pageSize: 100 })
    columns.value = response.data?.items || []
  } catch (error) {
    console.error('获取专栏数据失败:', error)
  }
}

// 在组件挂载时获取数据
onMounted(async () => {
  await Promise.all([
    fetchCategories(),
    fetchTags(),
    fetchColumns()
  ])
})

// 过滤后的文章列表
const filteredArticles = computed(() => {
  return props.articles.filter(article => {
    // 标题搜索
    const matchesSearch = !searchKeyword.value || 
      article.title.toLowerCase().includes(searchKeyword.value.toLowerCase())
    
    // 分类筛选
    const matchesCategory = !filterCategory.value || 
      article.category === filterCategory.value
    
    // 标签筛选
    const matchesTag = !filterTag.value || 
      article.tags.includes(filterTag.value)
    
    // 专栏筛选
    const matchesColumn = !filterColumn.value || 
      article.column.includes(filterColumn.value)
    
    return matchesSearch && matchesCategory && matchesTag && matchesColumn
  })
})

// 处理搜索
const handleSearch = () => {
  // 可以添加防抖处理
  console.log('搜索关键词:', searchKeyword.value)
}

// 处理筛选变化
const handleFilterChange = () => {
  console.log('筛选条件变化:', {
    category: filterCategory.value,
    tag: filterTag.value,
    column: filterColumn.value
  })
}

// 处理分页大小变化
const handleSizeChange = (size) => {
  props.pagination.pageSize = size
  emit('page-change', 1) // 重置到第一页
}

// 处理当前页变化
const handleCurrentChange = (page) => {
  emit('page-change', page)
}

// 编辑文章
const handleEdit = (article) => {
  emit('edit', article)
}

// 删除文章
const handleDelete = (article) => {
  emit('delete', article)
}

// 更改文章状态
const handleChangeStatus = (article, newStatus) => {
  emit('change-status', article, newStatus)
}
</script>

<style scoped>
.article-list {
  padding: 20px;
}

.filter-section {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.article-title {
  display: flex;
  align-items: center;
}

.title-text {
  font-weight: 500;
  color: #303133;
}

.category {
  color: #606266;
}
.column {
  color: #606266;
}

.tags {
  display: flex;
  flex-wrap: wrap;
}

.time {
  color: #909399;
  font-size: 12px;
}

.views,
.likes,
.comments {
  font-weight: 600;
}

.views {
  color: #409EFF;
}

.likes {
  color: #E6A23C;
}

.comments {
  color: #67C23A;
}

.action-buttons {
  display: flex;
  gap: 8px;
}

.pagination-section {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>