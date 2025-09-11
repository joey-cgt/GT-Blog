<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'

const props = defineProps({
  contentSelector: {
    type: String,
    default: '.article-body'
  }
})

const headings = ref([])
const activeHeadingId = ref('')
const observer = ref(null)

// 提取标题并生成目录结构
const extractHeadings = () => {
  const contentElement = document.querySelector(props.contentSelector)
  if (!contentElement) return []

  const headingElements = contentElement.querySelectorAll('h1, h2, h3, h4, h5, h6')
  const headingData = []

  headingElements.forEach((heading, index) => {
    // 为标题添加ID（如果还没有）
    let id = heading.id
    if (!id) {
      id = `heading-${index}`
      heading.id = id
    }

    headingData.push({
      id: id,
      text: heading.textContent,
      level: parseInt(heading.tagName.substring(1)),
      element: heading
    })
  })

  return headingData
}

// 生成嵌套的目录结构
const generateTocStructure = (headings) => {
  if (headings.length === 0) return []

  const toc = []
  const stack = []

  headings.forEach(heading => {
    const item = {
      id: heading.id,
      text: heading.text,
      level: heading.level,
      children: []
    }

    // 处理嵌套关系
    while (stack.length > 0 && stack[stack.length - 1].level >= heading.level) {
      stack.pop()
    }

    if (stack.length === 0) {
      toc.push(item)
    } else {
      stack[stack.length - 1].children.push(item)
    }

    stack.push(item)
  })

  return toc
}

// 平滑滚动到指定标题（考虑Header高度偏移）
const scrollToHeading = (id) => {
  const element = document.getElementById(id)
  if (element) {
    // 获取Header高度（从header元素获取实际高度）
    const header = document.querySelector('header')
    const headerHeight = header ? header.offsetHeight : 70 // 默认70px，与Header组件一致
    
    // 计算目标位置（考虑Header高度和额外间距）
    const targetPosition = element.offsetTop - headerHeight - 16 // 额外16px间距
    
    window.scrollTo({
      top: targetPosition,
      behavior: 'smooth'
    })
    // 更新URL hash
    window.history.replaceState(null, null, `#${id}`)
  }
}

// 设置IntersectionObserver来检测当前可见的标题
const setupObserver = () => {
  if (observer.value) {
    observer.value.disconnect()
  }

  observer.value = new IntersectionObserver(
    (entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          activeHeadingId.value = entry.target.id
        }
      })
    },
    {
      rootMargin: '-20% 0px -80% 0px',
      threshold: 0
    }
  )

  // 观察所有标题元素
  headings.value.forEach(heading => {
    const element = document.getElementById(heading.id)
    if (element) {
      observer.value.observe(element)
    }
  })
}



// 初始化目录
const initToc = () => {
  const extractedHeadings = extractHeadings()
  headings.value = extractedHeadings
  
  if (extractedHeadings.length > 0) {
    setupObserver()
  }
}

// 监听文章内容变化
watch(() => props.contentSelector, initToc)

onMounted(() => {
  // 使用nextTick确保DOM已经渲染完成
  setTimeout(initToc, 100)
})

onUnmounted(() => {
  if (observer.value) {
    observer.value.disconnect()
  }
})

// 暴露方法给模板使用
defineExpose({
  scrollToHeading
})
</script>

<template>
  <div v-if="headings.length > 0" class="table-of-contents">
    <div class="toc-header">
      <h3 class="toc-title">目录</h3>
    </div>
    <nav class="toc-nav">
      <ul class="toc-list">
        <li v-for="item in generateTocStructure(headings)" :key="item.id" 
            :class="['toc-item', `level-${item.level}`, { active: activeHeadingId === item.id }]">
          <a href="#" class="toc-link" @click.prevent="scrollToHeading(item.id)">
            {{ item.text }}
          </a>
          <ul v-if="item.children && item.children.length > 0" class="toc-children">
            <li v-for="child in item.children" :key="child.id" 
                :class="['toc-item', `level-${child.level}`, { active: activeHeadingId === child.id }]">
              <a href="#" class="toc-link" @click.prevent="scrollToHeading(child.id)">
                {{ child.text }}
              </a>
            </li>
          </ul>
        </li>
      </ul>
    </nav>
  </div>
</template>

<style scoped>
.table-of-contents {
  background: #fff;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 20px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.toc-header {
  margin-bottom: 12px;
  padding-bottom: 8px;
  border-bottom: 1px solid #eaeaea;
}

.toc-title {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.toc-nav {
  max-height: 500px;
  overflow-y: auto;
}

.toc-list {
  list-style: none;
  padding: 0;
  margin: 0;
  text-align: left;
}

.toc-item {
  margin: 4px 0;
  line-height: 1.4;
}

.toc-link {
  display: block;
  padding: 6px 8px;
  text-decoration: none;
  color: #555;
  font-size: 14px;
  border-radius: 4px;
  transition: all 0.2s ease;
  border-left: 3px solid transparent;
}

.toc-link:hover {
  color: #3498db;
  background-color: #f8f9fa;
}

.toc-item.active .toc-link {
  color: #3498db;
  background-color: #f0f0f0;
  /* border-left-color: #3498db; */
  font-weight: 500;
}

/* 层级缩进 */
.toc-item.level-1 { padding-left: 0px; }
.toc-item.level-2 { padding-left: 12px; }
.toc-item.level-3 { padding-left: 24px; }
.toc-item.level-4 { padding-left: 36px; }
.toc-item.level-5 { padding-left: 48px; }
.toc-item.level-6 { padding-left: 60px; }

.toc-children {
  list-style: none;
  padding-left: 16px;
  margin: 4px 0;
}

/* 滚动条样式 */
.toc-nav::-webkit-scrollbar {
  width: 4px;
}

.toc-nav::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 2px;
}

.toc-nav::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 2px;
}

.toc-nav::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .table-of-contents {
    padding: 12px;
  }
  
  .toc-nav {
    max-height: 300px;
  }
}
</style>