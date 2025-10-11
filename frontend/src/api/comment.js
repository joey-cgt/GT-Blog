import request from '../utils/request'

// 提交评论
export function submitComment(data) {
  return request({
    url: '/api/v1/comments',
    method: 'post',
    data
  })
}

// 获取文章评论列表
export function getArticleComments(articleId, params = {}) {
  return request({
    url: `/api/v1/comments/article/${articleId}`,
    method: 'get',
    params
  })
}

// 删除评论
export function deleteComment(commentId) {
  return request({
    url: `/api/v1/comments/${commentId}`,
    method: 'delete'
  })
}