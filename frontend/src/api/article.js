import request from '../utils/request'

// 创建并发布文章
export function createAndPublishArticle(data) {
  return request({
    url: '/api/v1/articles',
    method: 'post',
    data
  })
}

// 更新已发布文章
export function updatePublishedArticle(id, data) {
  return request({
    url: `/api/v1/articles/${id}`,
    method: 'put',
    data
  })
}

// 保存草稿
export function createDraft(data) {
  return request({
    url: '/api/v1/articles/drafts',
    method: 'post',
    data
  })
}

// 更新草稿
export function updateDraft(id, data) {
  return request({
    url: `/api/v1/articles/drafts/${id}`,
    method: 'put',
    data
  })
}

// 发布草稿
export function publishDraft(id, data) {
  return request({
    url: `/api/v1/articles/drafts/${id}/publish`,
    method: 'post',
    data
  })
}

// 获取文章详情
export function getArticleDetail(id) {
  return request({
    url: `/api/v1/articles/${id}`,
    method: 'get'
  })
}

// 获取文章列表
export function getArticleList(params) {
  return request({
    url: '/api/v1/articles',
    method: 'get',
    params
  })
}

// 删除文章
export function deleteArticle(id) {
  return request({
    url: `/api/v1/articles/${id}`,
    method: 'delete'
  })
}