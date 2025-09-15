<script setup>
import { computed } from 'vue'
import MarkdownIt from 'markdown-it'
import hljs from 'highlight.js'
import 'highlight.js/styles/github.css'
import katex from 'katex'
import 'katex/dist/katex.min.css'
import texmath from 'markdown-it-texmath'
import 'github-markdown-css'

const props = defineProps({
  content: {
    type: String,
    default: ''
  }
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
})

const renderedContent = computed(() => {
  return md.render(props.content)
})
</script>

<template>
  <div class="content markdown-body" v-html="renderedContent"></div>
</template>

<style scoped>
.content.markdown-body {
    text-align: left;
}
</style>