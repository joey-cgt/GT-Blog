<template>
  <div class="comments-section">
    <h3 class="comments-title">评论区</h3>
    <button 
      @click="showCommentDialog = true" 
      class="add-comment-btn"
    >
      写下评论
    </button>
    
    <!-- 评论弹窗 -->
    <el-dialog 
      v-model="showCommentDialog" 
      title="发表评论"
      width="500px"
    >
      <el-form 
        ref="commentFormRef" 
        :model="commentForm" 
        :rules="commentRules"
        label-width="80px"
      >
        <el-form-item label="昵称" prop="nickname">
          <el-input v-model="commentForm.nickname" placeholder="请输入您的昵称" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="commentForm.email" placeholder="请输入您的邮箱" type="email" />
        </el-form-item>
        <el-form-item label="评论内容" prop="content">
          <el-input
            v-model="commentForm.content"
            type="textarea"
            placeholder="请输入评论内容"
            :rows="4"
          />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showCommentDialog = false">取消</el-button>
          <el-button type="primary" @click="handleSubmitComment">提交评论</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 评论列表 -->
    <div class="comments-list" v-if="comments.length > 0">
      <div class="comment-item" v-for="comment in comments" :key="comment.id">
        <div class="comment-header">
          <div class="comment-author-info">
            <span class="comment-author">{{ comment.nickname }}</span>
            <span class="comment-email">{{ comment.email }}</span>
          </div>
          <span class="comment-time">{{ formatDate(comment.createdAt) }}</span>
        </div>
        <div class="comment-content">{{ comment.content }}</div>
      </div>
    </div>
    <div class="no-comments" v-else>
      暂无评论，快来抢沙发吧！
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { submitComment, getArticleComments } from '../../api/comment.js'

const props = defineProps({
  articleId: {
    type: Number,
    required: true
  }
})

// 评论表单相关
const showCommentDialog = ref(false)
const commentFormRef = ref(null)
const commentForm = reactive({
  nickname: '',
  email: '',
  content: ''
})

const commentRules = {
  nickname: [
    { required: true, message: '请输入昵称', trigger: 'blur' },
    { min: 2, max: 20, message: '昵称长度在 2 到 20 个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
  ],
  content: [
    { required: true, message: '请输入评论内容', trigger: 'blur' },
    { min: 1, message: '评论内容至少 1 个字符', trigger: 'blur' }
  ]
}

// 评论列表
const comments = ref([])

// 格式化日期
const formatDate = (dateString) => {
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

// 提交评论
const handleSubmitComment = async () => {
  try {
    await commentFormRef.value.validate()
    
    // 准备评论数据
    const commentData = {
      articleId: props.articleId,
      nickname: commentForm.nickname,
      email: commentForm.email,
      content: commentForm.content
    }
    
    // 调用API提交评论
    await submitComment(commentData)
    
    // 重置表单并关闭弹窗
    commentFormRef.value.resetFields()
    showCommentDialog.value = false
    
    // 重新加载评论列表
    loadComments()
    
    ElMessage.success('评论发表成功')
  } catch (error) {
    console.error('评论提交失败:', error)
    ElMessage.error('评论提交失败，请稍后重试')
  }
}

// 加载评论列表
const loadComments = async () => {
  try {
    const response = await getArticleComments(props.articleId)
    comments.value = response.data.comments || []
  } catch (error) {
    console.error('加载评论列表失败:', error)
    ElMessage.error('加载评论列表失败')
  }
}

// 监听文章ID变化，重新加载评论
watch(() => props.articleId, () => {
  if (props.articleId) {
    loadComments()
  }
})

// 组件挂载时加载评论
loadComments()
</script>

<style scoped>
.comments-section {
  margin-top: 30px;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
}

.comments-title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 20px;
}

.add-comment-btn {
  padding: 10px 24px;
  background-color: #409eff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  margin-bottom: 20px;
}

.add-comment-btn:hover {
  background-color: #66b1ff;
}

.comments-list {
  margin-top: 20px;
}

.comment-item {
  padding: 15px;
  background-color: white;
  border-radius: 4px;
  margin-bottom: 15px;
  border: 1px solid #eee;
}

.comment-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
}

.comment-author-info {
  display: flex;
  align-items: center;
}

.comment-author {
  font-weight: 600;
  color: #303133;
}

.comment-email {
  color: #606266;
  font-size: 12px;
  margin-left: 8px;
}

.comment-time {
  color: #909399;
  font-size: 12px;
}

.comment-content {
  font-size: 14px;
  line-height: 1.6;
  color: #606266;
  text-align: left;
}

.no-comments {
  text-align: center;
  color: #999;
  padding: 40px;
  font-size: 14px;
}
</style>