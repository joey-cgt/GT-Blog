<template>
  <div class="admin-dashboard">    
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
  </div>
</template>

<script setup>
import { ref, onMounted, watch, nextTick, onUnmounted } from 'vue'
import { Document, Star, View, Top, Collection, Folder, PriceTag, Right } from '@element-plus/icons-vue'
import * as echarts from 'echarts'

const timeRange = ref('7d')
const chartRef = ref(null)
let chartInstance = null

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
  })
})

// 组件卸载时销毁图表
onUnmounted(() => {
  if (chartInstance) {
    chartInstance.dispose()
  }
})
</script>

<style scoped>
.admin-dashboard {
  padding: 20px;
}

.stats-cards {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(2, auto);
  gap: 20px;
  margin-top: 20px;
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
  width: 48px;
  height: 48px;
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
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 4px;
}

</style>