import request from '@/utils/request'

// 获取作者信息
export function getAuthor(id) {
  return request({
    url: `/api/v1/public/authorinfo/${id}`,
    method: 'get', 
  })
}

/**
 * 获取推荐文章列表
 * @param {Object} params - 查询参数
 * @param {string|number} [params.id] - 当前阅读的文章ID
 * @param {number} [params.limit=5] - 返回的文章数量
 * @returns {Promise} - 返回推荐文章列表
 */
export function getRecommendedArticles(params = {}) {
  // 设置默认参数
  const defaultParams = {
    limit: 5
  }
  
  // 合并默认参数和传入参数
  const queryParams = {
    ...defaultParams,
    ...params
  }
  
  return request({
    url: '/api/v1/articles/recommended',
    method: 'get',
    params: queryParams
  })
}

