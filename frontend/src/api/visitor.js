import request from '@/utils/request'

// 获取作者信息
export function getAuthor(id) {
  return request({
    url: `/api/v1/public/authorinfo/${id}`,
    method: 'get', 
  })
}
