<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { articles } from '../store/blog.js'
import MarkdownIt from 'markdown-it'
import hljs from 'highlight.js'
import 'highlight.js/styles/github.css'
import katex from 'katex'
import 'katex/dist/katex.min.css'
import texmath from 'markdown-it-texmath'
import '../assets/styles/themes/github.css'

const route = useRoute()
const articleId = computed(() => parseInt(route.params.id))
const article = computed(() => {
  return articles.find(a => a.id === articleId.value) || null
})



// 初始化markdown-it实例
const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true,
  highlight: function (str, lang) {
    if (lang && hljs.getLanguage(lang)) {
      try {
        return hljs.highlight(str, { language: lang }).value;
      } catch (__) {}
    }
    return ''; // 使用默认的转义
  }
})

// 添加数学公式支持
md.use(texmath, {
  engine: katex,
  delimiters: 'dollars',
  katexOptions: { macros: { "\\RR": "\\mathbb{R}" } }
});

// 模拟文章内容（实际项目中可能从API获取）
const content = ref('')
const renderedContent = computed(() => {
  return md.render(content.value)
})

onMounted(() => {
  // 模拟文章内容（实际项目中应该从API获取）
  if (article.value) {
    content.value = generateDummyContent(article.value.title)
  }
})

// 生成模拟Markdown内容的辅助函数
function generateDummyContent(title) {
  return `# ${title}

这是一篇关于${title}的详细文章。在实际应用中，这里应该是从后端API获取的文章详细内容。

## 第一部分

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam euismod, nisl eget aliquam ultricies, 
nunc nisl aliquet nunc, vitae aliquam nisl nunc vitae nisl. Nullam euismod, nisl eget aliquam ultricies, 
nunc nisl aliquet nunc, vitae aliquam nisl nunc vitae nisl.

## 第二部分

Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor 
in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

### 代码示例

\`\`\`javascript
function example() {
  console.log("这是一个代码示例");
  return "Hello Markdown!";
}
\`\`\`

## 数学公式支持

行内公式: $E = mc^2$

独立公式:

$$
\\frac{d}{dx}\\left( \\int_{0}^{x} f(u)\\,du\
ight)=f(x)
$$

## 第三部分

Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, 
totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.

## 总结

Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni 
dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor 
sit amet, consectetur, adipisci velit.
`
}
</script>

<template>
  <div v-if="article" class="article-content-page">
    <!-- 文章头部信息 -->
    <div class="article-header">
      <div class="article-cover">
        <img :src="article.cover" :alt="article.title" />
      </div>
      
      <h1 class="article-title">{{ article.title }}</h1>
      
      <div class="article-meta">
        <span class="meta-item date">{{ article.date }}</span>
        <span class="meta-item category">{{ article.category }}</span>
        <span class="meta-item views">{{ article.views }} 阅读</span>
      </div>
      
      <div class="article-tags">
        <span v-for="tag in article.tags" :key="tag" class="tag">{{ tag }}</span>
      </div>
    </div>
    
    <!-- 文章内容 -->
    <div class="article-body" v-html="renderedContent"></div>
  </div>
  
  <div v-else class="article-not-found">
    <h2>文章不存在</h2>
    <p>抱歉，您请求的文章不存在或已被删除。</p>
    <router-link to="/articles" class="back-link">返回文章列表</router-link>
  </div>
</template>

<style scoped>
.article-content-page {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  padding: 30px;
  margin-bottom: 30px;
}

.article-header {
  margin-bottom: 30px;
  text-align: left;
}

.article-cover {
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 20px;
}

.article-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.article-title {
  font-size: 32px;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 16px;
  line-height: 1.3;
}

.article-meta {
  display: flex;
  gap: 16px;
  margin-bottom: 16px;
  color: #7f8c8d;
  font-size: 14px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.article-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 20px;
}

.tag {
  background-color: #f0f0f0;
  color: #555;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
}

.article-body {
  line-height: 1.8;
  color: #333;
  font-size: 16px;
  text-align: left;
}

.article-body h1,
.article-body h2,
.article-body h3 {
  margin-top: 30px;
  margin-bottom: 15px;
  color: #2c3e50;
}

.article-body p {
  margin-bottom: 20px;
}

/* Markdown 样式 */
.article-body pre {
  background-color: #f5f5f5;
  padding: 16px;
  border-radius: 4px;
  overflow-x: auto;
  margin-bottom: 20px;
}

.article-body code {
  font-family: 'Courier New', Courier, monospace;
  background-color: #f5f5f5;
  padding: 2px 4px;
  border-radius: 3px;
  font-size: 0.9em;
}

.article-body pre code {
  padding: 0;
  background-color: transparent;
}

.article-body blockquote {
  border-left: 4px solid #ddd;
  padding-left: 16px;
  color: #666;
  margin: 0 0 20px;
}

.article-body img {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 20px 0;
}

.article-body table {
  border-collapse: collapse;
  width: 100%;
  margin-bottom: 20px;
}

.article-body table, .article-body th, .article-body td {
  border: 1px solid #ddd;
}

.article-body th, .article-body td {
  padding: 8px 12px;
  text-align: left;
}

.article-body th {
  background-color: #f5f5f5;
}

/* KaTeX 样式调整 */
.article-body .katex-display {
  margin: 1.5em 0;
  overflow-x: auto;
  overflow-y: hidden;
}

.article-not-found {
  text-align: center;
  padding: 50px 20px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.back-link {
  display: inline-block;
  margin-top: 20px;
  color: #3498db;
  text-decoration: none;
}

.back-link:hover {
  text-decoration: underline;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .article-content-page {
    padding: 20px;
  }
  
  .article-cover {
    height: 200px;
  }
  
  .article-title {
    font-size: 24px;
  }
}
</style>