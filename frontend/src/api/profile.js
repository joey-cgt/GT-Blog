import request from '@/utils/request'

export function getProfile(userId) {
  return request({
    url: `/api/profiles/${userId}`,
    method: 'get'
  })
}

export function updateProfile(data) {
  return request({
    url: '/api/profiles/update',
    method: 'post',
    data
  })
}