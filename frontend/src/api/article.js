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

export function getPopularArticles() {
  return request({
    url: '/api/v1/articles/popular',
    method: 'get',
  })
}

export function getLatestArticles() {
  return request({
    url: '/api/v1/articles/latest',
    method: 'get',
  })
}

export function getPinnedArticles() {
  return request({
    url: '/api/v1/articles/top',
    method: 'get',
  })
}

// 获取浏览量最高的前五篇文章
export function getMostViewedArticles() {
  return request({
    url: '/api/v1/articles/most-viewed',
    method: 'get',
  })
}

// 获取点赞量最高的前五篇文章
export function getMostLikedArticles() {
  return request({
    url: '/api/v1/articles/most-liked',
    method: 'get',
  })
}

export function getColumnArticles(columnId) {
  return request({
    url: `/api/v1/articles/aggregated?type=column&id=${columnId}`,
    method: 'get',
  })
}

export function getCategoryArticles(categoryId) {
  return request({
    url: `/api/v1/articles/aggregated?type=category&id=${categoryId}`,
    method: 'get',
  })
}

export function getTagArticles(tagId) {
  return request({
    url: `/api/v1/articles/aggregated?type=tag&id=${tagId}`,
    method: 'get',
  })
}

export function incrementArticleLike(articleId) {
  return request({
    url: `/api/v1/articles/${articleId}/like`,
    method: 'post',
  })
}

export function decrementArticleLike(articleId) {
  return request({
    url: `/api/v1/articles/${articleId}/like`,
    method: 'delete',
  })
}

// 记录文章浏览量
export function recordArticleView(articleId) {
  return request({
    url: `/api/v1/articles/${articleId}/view`,
    method: 'post',
  })
}