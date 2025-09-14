 <template>
  <div class="admin-dashboard">
    <!-- 操作和统计卡片区域 -->
    <div class="dashboard-main">
      <!-- 左侧操作卡片 -->
      <div class="action-cards">
        <!-- 发布文章卡片 -->
        <el-card class="action-card publish-card" @click="handlePublishArticle">
          <div class="card-content">
            <div class="action-icon">
              <el-icon size="32px"><Notebook /></el-icon>
            </div>
            <div class="action-info">
              <div class="action-title">发布文章</div>
              <div class="action-desc">创建新的博客内容</div>
            </div>
          </div>
        </el-card>

        <!-- 草稿箱卡片 -->
        <el-card class="action-card draft-card">
          <div class="card-content">
            <div class="action-icon">
              <el-icon size="32px"><Files /></el-icon>
            </div>
            <div class="action-info">
              <div class="action-title">草稿箱</div>
              <div class="action-desc">3篇草稿待完成</div>
            </div>
          </div>
        </el-card>
      </div>

      <!-- 垂直分隔线 -->
      <div class="divider"></div>

      <!-- 右侧统计卡片 -->
      <div class="stats-cards">
      <!-- 文章数量卡片 -->
        <el-card class="stat-card">
            <div class="card-content">
            <div class="stat-icon">
                <el-icon size="24px"><Document /></el-icon>
            </div>
            <div class="stat-info">
                <div class="stat-label">文章数量</div>
                <div class="stat-value">1,248</div>
            </div>
            </div>
        </el-card>
   
        <!-- 总浏览量卡片 -->
        <el-card class="stat-card">
            <div class="card-content">
            <div class="stat-icon">
                <el-icon size="24px"><View /></el-icon>
            </div>
            <div class="stat-info">
                <div class="stat-label">总浏览量</div>
                <div class="stat-value">24,891</div>
            </div>
            </div>
        </el-card>

        <!-- 总点赞数卡片 -->
        <el-card class="stat-card">
            <div class="card-content">
            <div class="stat-icon">
                <el-icon size="24px"><Star /></el-icon>
            </div>
            <div class="stat-info">
                <div class="stat-label">总点赞数</div>
                <div class="stat-value">5,672</div>
            </div>
            </div>
        </el-card>

        <!-- 专栏数量卡片 -->
        <el-card class="stat-card">
            <div class="card-content">
            <div class="stat-icon">
                <el-icon size="24px"><Collection /></el-icon>
            </div>
            <div class="stat-info">
                <div class="stat-label">专栏数量</div>
                <div class="stat-value">24</div>
            </div>
            </div>
        </el-card>

        <!-- 分类数量卡片 -->
        <el-card class="stat-card">
            <div class="card-content">
            <div class="stat-icon">
                <el-icon size="24px"><Folder /></el-icon>
            </div>
            <div class="stat-info">
                <div class="stat-label">分类数量</div>
                <div class="stat-value">12</div>
            </div>
            </div>
        </el-card>

        <!-- 标签数量卡片 -->
        <el-card class="stat-card">
            <div class="card-content">
            <div class="stat-icon">
                <el-icon size="24px"><PriceTag /></el-icon>
            </div>
            <div class="stat-info">
                <div class="stat-label">标签数量</div>
                <div class="stat-value">156</div>
            </div>
            </div>
        </el-card>
      </div>
  </div>

    <!-- 趋势图表区域 -->
    <div class="chart-section">
      <el-card class="chart-card">
        <template #header>
          <div class="chart-header">
            <h3>浏览量趋势</h3>
            <div class="time-selector">
              <el-radio-group v-model="timeRange" size="small">
                <el-radio-button label="7d">近7天</el-radio-button>
                <el-radio-button label="30d">近30天</el-radio-button>
                <el-radio-button label="90d">近90天</el-radio-button>
              </el-radio-group>
            </div>
          </div>
        </template>
        
        <div class="chart-container">
          <div ref="chartRef" style="width: 100%; height: 300px;"></div>
        </div>
      </el-card>
    </div>

    <!-- 热门文章区域 -->
    <div class="popular-articles">
      <div class="article-column">
        <el-card class="article-card">
          <template #header>
            <div class="card-header">
              <h3>📈 浏览量最高文章</h3>
            </div>
          </template>
          <el-table :data="mostViewedArticles" size="small" stripe>
            <el-table-column prop="rank" label="排名" width="60">
              <template #default="{ $index }">
                <div class="rank-badge">{{ $index + 1 }}</div>
              </template>
            </el-table-column>
            <el-table-column prop="title" label="文章标题" min-width="200">
              <template #default="{ row }">
                <div class="article-title">{{ row.title }}</div>
              </template>
            </el-table-column>
            <el-table-column prop="views" label="浏览量" width="80" align="center">
              <template #default="{ row }">
                <span class="views-count">{{ row.views }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="date" label="日期" width="100" align="center">
              <template #default="{ row }">
                <span class="article-date">{{ row.date }}</span>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </div>

      <div class="article-column">
        <el-card class="article-card">
          <template #header>
            <div class="card-header">
              <h3>⭐ 点赞最多文章</h3>
            </div>
          </template>
          <el-table :data="mostLikedArticles" size="small" stripe>
            <el-table-column prop="rank" label="排名" width="60">
              <template #default="{ $index }">
                <div class="rank-badge">{{ $index + 1 }}</div>
              </template>
            </el-table-column>
            <el-table-column prop="title" label="文章标题" min-width="200">
              <template #default="{ row }">
                <div class="article-title">{{ row.title }}</div>
              </template>
            </el-table-column>
            <el-table-column prop="likes" label="点赞数" width="80" align="center">
              <template #default="{ row }">
                <span class="likes-count">{{ row.likes }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="date" label="日期" width="100" align="center">
              <template #default="{ row }">
                <span class="article-date">{{ row.date }}</span>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, nextTick, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { Document, Star, View, Top, Collection, Folder, PriceTag, Right, Notebook, Files } from '@element-plus/icons-vue'
import * as echarts from 'echarts'

const router = useRouter()
const timeRange = ref('7d')
const chartRef = ref(null)
let chartInstance = null

// 发布新文章
const handlePublishArticle = () => {
  router.push('/editor/drafts/new')
}

// 模拟热门文章数据
const mostViewedArticles = ref([
  { id: 1, title: 'Vue 3 组合式API最佳实践', views: 2456, date: '2024-01-15' },
  { id: 2, title: 'TypeScript 高级类型技巧', views: 1987, date: '2024-01-12' },
  { id: 3, title: 'Element Plus 组件库深度解析', views: 1765, date: '2024-01-10' },
  { id: 4, title: '前端性能优化完全指南', views: 1543, date: '2024-01-08' },
  { id: 5, title: 'CSS Grid 布局实战教程', views: 1321, date: '2024-01-05' }
])

const mostLikedArticles = ref([
  { id: 1, title: 'Vue 3 组合式API最佳实践', likes: 324, date: '2024-01-15' },
  { id: 2, title: 'TypeScript 高级类型技巧', likes: 287, date: '2024-01-12' },
  { id: 6, title: 'JavaScript 异步编程详解', likes: 256, date: '2024-01-18' },
  { id: 3, title: 'Element Plus 组件库深度解析', likes: 234, date: '2024-01-10' },
  { id: 7, title: 'Webpack 5 配置优化指南', likes: 198, date: '2024-01-20' }
])

// 模拟数据
const generateChartData = (days) => {
  const data = []
  const now = new Date()
  
  for (let i = days - 1; i >= 0; i--) {
    const date = new Date(now)
    date.setDate(date.getDate() - i)
    const dateStr = date.toLocaleDateString('zh-CN', { month: '2-digit', day: '2-digit' })
    
    // 生成随机浏览量数据
    const views = Math.floor(Math.random() * 100) + 5
    data.push({ date: dateStr, views })
  }
  
  return data
}

const initChart = () => {
  if (!chartRef.value) return
  
  chartInstance = echarts.init(chartRef.value)
  
  const updateChart = () => {
    const days = timeRange.value === '7d' ? 7 : timeRange.value === '30d' ? 30 : 90
    const chartData = generateChartData(days)
    
    const option = {
      tooltip: {
        trigger: 'axis',
        backgroundColor: 'rgba(255, 255, 255, 0.95)',
        borderColor: '#e4e7ed',
        textStyle: {
          color: '#606266'
        }
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
      },
      xAxis: {
        type: 'category',
        data: chartData.map(item => item.date),
        axisLine: {
          lineStyle: {
            color: '#dcdfe6'
          }
        },
        axisLabel: {
          color: '#909399'
        }
      },
      yAxis: {
        type: 'value',
        axisLine: {
          show: false
        },
        axisTick: {
          show: false
        },
        axisLabel: {
          color: '#909399'
        },
        splitLine: {
          lineStyle: {
            color: '#f0f0f0'
          }
        }
      },
      series: [
        {
          name: '浏览量',
          type: 'line',
          data: chartData.map(item => item.views),
          smooth: true,
          symbol: 'circle',
          symbolSize: 6,
          itemStyle: {
            color: '#409EFF'
          },
          lineStyle: {
            color: '#409EFF',
            width: 3
          },
          areaStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              { offset: 0, color: 'rgba(64, 158, 255, 0.3)' },
              { offset: 1, color: 'rgba(64, 158, 255, 0.1)' }
            ])
          }
        }
      ]
    }
    
    chartInstance.setOption(option)
  }
  
  updateChart()
  
  // 监听时间范围变化
  watch(timeRange, () => {
    updateChart()
  })
}

onMounted(() => {
  nextTick(() => {
    initChart()
    // 添加窗口resize事件监听
    window.addEventListener('resize', handleResize)
  })
})

// 窗口大小变化处理
const handleResize = () => {
  if (chartInstance) {
    chartInstance.resize()
  }
}

// 组件卸载时销毁图表和移除事件监听
onUnmounted(() => {
  if (chartInstance) {
    chartInstance.dispose()
  }
  window.removeEventListener('resize', handleResize)
})
</script>

<style scoped>
.admin-dashboard {
  padding: 20px;
}

.dashboard-main {
  display: grid;
  grid-template-columns: 280px 1px 1fr;
  gap: 24px;
  align-items: start;
  margin-bottom: 30px;
}

.action-cards {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.action-card {
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  border-left: 4px solid transparent;
}

.action-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.publish-card {
  border-left-color: #409EFF !important;
}

.publish-card .action-icon {
  color: #409EFF;
}

.draft-card {
  border-left-color: #E6A23C !important;
}

.draft-card .action-icon {
  color: #E6A23C;
}

.action-card .card-content {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 4px;
}

.action-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  border-radius: 8px;
  background: rgba(64, 158, 255, 0.1);
}

.draft-card .action-icon {
  background: rgba(230, 162, 60, 0.1);
}

.action-info {
  flex: 1;
}

.action-title {
  font-size: 18px;
  font-weight: 600;
  color: #303133;
  margin-bottom: 4px;
}

.action-desc {
  font-size: 14px;
  color: #909399;
}

.divider {
  width: 1px;
  background: #e4e7ed;
  height: 100%;
  min-height: 200px;
}

.stats-cards {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(2, auto);
  gap: 20px;
}

@media (max-width: 1200px) {
  .dashboard-main {
    grid-template-columns: 1fr;
    gap: 20px;
  }
  
  .divider {
    display: none;
  }
  
  .stats-cards {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media (max-width: 900px) {
  .stats-cards {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 600px) {
  .stats-cards {
    grid-template-columns: 1fr;
  }
  
  .action-card .card-content {
    padding: 16px;
  }
  
  .action-icon {
    width: 40px;
    height: 40px;
  }
}

.chart-section {
  margin-top: 30px;
}

.chart-card {
  border-radius: 8px;
}

.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 10px;
}

.chart-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.time-selector {
  display: flex;
  align-items: center;
}

.chart-container {
  padding: 10px;
}

.popular-articles {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 20px;
  margin-top: 30px;
}

@media (max-width: 900px) {
  .popular-articles {
    grid-template-columns: 1fr;
  }
}

.article-column {
  height: 100%;
}

.article-card {
  height: 100%;
  border-radius: 8px;
}

.card-header {
  display: flex;
  align-items: center;
  padding: 0 10px;
}

.card-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.rank-badge {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  font-size: 12px;
  font-weight: bold;
}

.article-title {
  font-size: 16px;
  color: #303133;
  line-height: 1.4;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-weight: 500;
}

.views-count,
.likes-count {
  font-weight: 600;
  color: #409EFF;
}

.article-date {
  font-size: 12px;
  color: #909399;
}

/* 表格表头样式 */
:deep(.el-table__header-wrapper) .el-table__cell {
  font-size: 14px;
  font-weight: 600;
  color: #303133;
}

/* 表格内容样式 */
:deep(.el-table__body-wrapper) .el-table__cell {
  font-size: 14px;
}

:deep(.el-table__row) {
  height: 48px;  /* 增加行高，改善行间距 */
}

.stat-card {
  position: relative;
  border-radius: 8px;
}

.card-content {
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.stat-info {
  flex: 1;
}

.stat-label {  
  font-size: 18px;
  font-weight: 600;
  color: #606266;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 22px;
  font-weight: bold;
  color: #303133;
  /* margin-bottom: 4px; */
}

</style>