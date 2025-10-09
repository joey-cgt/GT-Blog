import request from '@/utils/request'

export function getColumnList(params) {
  return request({
    url: '/api/v1/columns',
    method: 'get',
    params
  })
}

export function getColumnDetail(id) {
  return request({
    url: `/api/v1/columns/${id}`,
    method: 'get'
  })
}

export function createColumn(data) {
  return request({
    url: '/api/v1/columns',
    method: 'post',
    data
  })
}

export function updateColumn(id, data) {
  return request({
    url: `/api/v1/columns/${id}`,
    method: 'put',
    data
  })
}

export function deleteColumn(id) {
  return request({
    url: `/api/v1/columns/${id}`,
    method: 'delete'
  })
}