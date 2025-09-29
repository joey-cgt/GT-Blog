import request from '@/utils/request'

// 获取分类列表
export function getCategoryList(params) {
  return request({
    url: '/api/v1/categories',
    method: 'get',
    params
  })
}

// 获取分类详情
export function getCategoryDetail(id) {
  return request({
    url: `/api/v1/categories/${id}`,
    method: 'get'
  })
}

// 创建分类
export function createCategory(data) {
  return request({
    url: '/api/v1/categories',
    method: 'post',
    data
  })
}

// 更新分类
export function updateCategory(id, data) {
  return request({
    url: `/api/v1/categories/${id}`,
    method: 'put',
    data
  })
}

// 删除分类
export function deleteCategory(id) {
  return request({
    url: `/api/v1/categories/${id}`,
    method: 'delete'
  })
}

