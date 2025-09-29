import request from '@/utils/request'

/**
 * 获取标签列表
 * @param {Object} params - 查询参数
 * @returns {Promise}
 */
export function getTags(params = {}) {
  return request({
    url: '/api/v1/tags',
    method: 'get',
    params
  })
}

/**
 * 获取单个标签详情
 * @param {number} id - 标签ID
 * @returns {Promise}
 */
export function getTagDetail(id) {
  return request({
    url: `/api/v1/tags/${id}`,
    method: 'get'
  })
}

/**
 * 创建标签
 * @param {Object} data - 标签数据
 * @param {string} data.name - 标签名称
 * @param {string} [data.description] - 标签描述
 * @returns {Promise}
 */
export function createTag(data) {
  return request({
    url: '/api/v1/tags',
    method: 'post',
    data
  })
}

/**
 * 更新标签
 * @param {number} id - 标签ID
 * @param {Object} data - 标签数据
 * @param {string} data.name - 标签名称
 * @param {string} [data.description] - 标签描述
 * @returns {Promise}
 */
export function updateTag(id, data) {
  return request({
    url: `/api/v1/tags/${id}`,
    method: 'put',
    data
  })
}

/**
 * 删除标签
 * @param {number} id - 标签ID
 * @returns {Promise}
 */
export function deleteTag(id) {
  return request({
    url: `/api/v1/tags/${id}`,
    method: 'delete'
  })
}

/**
 * 批量删除标签
 * @param {Array} ids - 标签ID数组
 * @returns {Promise}
 */
export function batchDeleteTags(ids) {
  return request({
    url: '/api/v1/tags/batch',
    method: 'delete',
    data: {
      ids
    }
  })
}