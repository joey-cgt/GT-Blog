<template>
  <div id="comments" class="comments-section">
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
      :title="replyToCommentId ? `回复 @${replyToCommentNickname}` : '发表评论'"
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
                <span v-if="comment.replyToNickname" class="reply-to-text">回复</span>
                <span v-if="comment.replyToNickname" class="reply-to-author">{{ comment.replyToNickname }}</span>
                <span v-if="comment.replyToEmail" class="reply-to-email">{{ comment.replyToEmail }}</span>
              </div>
              <span class="comment-time">{{ comment.createdAt }}</span>
            </div>
        <div class="comment-content" @click="handleReplyComment(comment)">{{ comment.content }}</div>
        
        <!-- 子评论列表 -->
        <div class="child-comment-list" v-if="comment.children && comment.children.length > 0">
          <div class="child-comment-item" v-for="childComment in sortedChildComments(comment.children)" :key="childComment.id">
            <div class="child-comment-header">
              <div class="child-comment-author-info">
                <span class="child-comment-author">{{ childComment.nickname }}</span>
                <span class="child-comment-email">{{ childComment.email }}</span>
                <span v-if="childComment.replyToNickname" class="child-reply-to-text">回复</span>
                <span v-if="childComment.replyToNickname" class="child-reply-to-author">{{ childComment.replyToNickname }}</span>
                <span v-if="childComment.replyToEmail" class="child-reply-to-email">{{ childComment.replyToEmail }}</span>
                <span class="child-comment-time">{{ childComment.createdAt }}</span>
              </div>
            </div>
            <div class="child-comment-content" @click="handleReplyComment(childComment)">{{ childComment.content }}</div>
            
            <!-- 所有层级的子评论都在顶级评论下平铺展开，这里不需要递归渲染 -->
          </div>
        </div>
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
const replyToCommentId = ref(null)
const replyToCommentNickname = ref('')
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

// 获取所有层级的子评论并平铺展开
const getAllChildComments = (children) => {
  let allComments = []
  
  const collectComments = (comments) => {
    if (!comments || !comments.length) return
    
    comments.forEach(comment => {
      allComments.push(comment)
      if (comment.children && comment.children.length) {
        collectComments(comment.children)
      }
    })
  }
  
  collectComments(children)
  return allComments
}

// 对子评论按照时间顺序排序
const sortedChildComments = (children) => {
  return getAllChildComments(children).sort((a, b) => new Date(a.createdAt) - new Date(b.createdAt))
}

// 处理回复评论
const handleReplyComment = (comment) => {
  replyToCommentId.value = comment.id
  replyToCommentNickname.value = comment.nickname
  showCommentDialog.value = true
}

// 提交评论
const handleSubmitComment = async () => {
  try {
    await commentFormRef.value.validate()
    
    // 准备评论数据，确保每个评论都提交ParentID字段，字段名与后端保持一致
    const commentData = {
      articleID: props.articleId,
      nickname: commentForm.nickname,
      email: commentForm.email, // email字段保留用于后端处理，但不会返回给前端
      content: commentForm.content,
      parentID: replyToCommentId.value || null
    }
    
    // 调用API提交评论
    await submitComment(commentData)
    
    // 重置表单并关闭弹窗
    commentFormRef.value.resetFields()
    showCommentDialog.value = false
    replyToCommentId.value = null
    replyToCommentNickname.value = ''
    
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
    // 如果没有真实数据，使用模拟数据
    let commentsData = response.data.comments && response.data.comments.length > 0 
      ? response.data.comments 
      : []
    
    // 预处理评论数据，为子评论添加回复目标的nickname和email
    commentsData = processCommentsForReplyInfo(commentsData)
    
    comments.value = commentsData
  } catch (error) {
    // 出错时使用模拟数据
    comments.value = []
    ElMessage.error('加载评论列表失败')
  }
}

// 预处理评论数据，为每个子评论添加回复目标的nickname和email
const processCommentsForReplyInfo = (comments) => {
  // 创建评论ID到评论对象的映射
  const commentMap = new Map()
  
  // 收集所有评论到映射中
  const collectComments = (commentList) => {
    commentList.forEach(comment => {
      commentMap.set(comment.id, comment)
      
      // 确保children数组存在
      if (!comment.children) {
        comment.children = []
      }
      
      // 递归处理子评论
      if (comment.children && comment.children.length > 0) {
        collectComments(comment.children)
      }
    })
  }
  
  collectComments(comments)
  
  // 为每个子评论添加回复信息
  const addReplyInfo = (commentList) => {
    commentList.forEach(comment => {
      // 如果评论有parentId，查找父评论并添加回复信息
      if (comment.parentId) {
        const parentComment = commentMap.get(comment.parentId)
        if (parentComment) {
          comment.replyToNickname = parentComment.nickname
          comment.replyToEmail = parentComment.email
        }
      }
      
      // 递归处理子评论
      if (comment.children && comment.children.length > 0) {
        addReplyInfo(comment.children)
      }
    })
  }
  
  addReplyInfo(comments)
  
  return comments
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
  margin-bottom: 10px;
  cursor: pointer;
  transition: color 0.3s;
}

.comment-content:hover {
  color: #409eff;
}

.no-comments {
  text-align: center;
  color: #999;
  padding: 40px;
  font-size: 14px;
}

/* 子评论列表样式 */
.child-comment-list {
  margin-top: 15px;
  padding-left: 20px;
  border-left: 2px solid #e6e6e6;
}

/* 所有层级子评论使用相同样式 */
.child-comment-list .child-comment-list {
  padding-left: 20px;
  border-left: 2px solid #e6e6e6;
}

.child-comment-item {
  padding: 10px;
  background-color: #f8f9fa;
  border-radius: 4px;
  margin-bottom: 10px;
  border: 1px solid #e9ecef;
}

.child-comment-header {
  display: flex;
  margin-bottom: 8px;
}

.child-comment-author-info {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  width: 100%;
}

.child-comment-author {
  font-weight: 600;
  color: #303133;
  font-size: 13px;
}

.child-comment-email {
  color: #606266;
  font-size: 11px;
}

.child-reply-to-text {
  color: #909399;
  font-size: 11px;
}

.child-reply-to-author {
  color: #409eff;
  font-size: 11px;
}

.child-reply-to-email {
  color: #606266;
  font-size: 11px;
}

.child-comment-time {
  color: #909399;
  font-size: 11px;
  margin-left: auto;
}

.child-comment-content {
  font-size: 13px;
  line-height: 1.5;
  color: #606266;
  text-align: left;
  cursor: pointer;
  transition: color 0.3s;
}

.child-comment-content:hover {
  color: #409eff;
}

/* 顶级评论样式保持不变 */
.comment-author-info {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
}

.comment-email {
  color: #606266;
  font-size: 12px;
}

.reply-to-text {
  color: #909399;
  font-size: 12px;
}

.reply-to-author {
  color: #409eff;
  font-size: 12px;
}

.reply-to-email {
  color: #606266;
  font-size: 12px;
  font-size: 13px;
}

.child-comment-time {
  color: #909399;
  font-size: 11px;
}

.child-comment-content {
  font-size: 13px;
  line-height: 1.5;
  color: #606266;
  text-align: left;
  cursor: pointer;
  transition: color 0.3s;
}

.child-comment-content:hover {
  color: #409eff;
}
</style>