import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/api/v1/admin/login',
    method: 'post',
    data
  })
}

export function updatePassword(data) {
  return request({
    url: '/api/v1/admin/password',
    method: 'put',
    data
  })
}

export function getProfile() {
  return request({
    url: '/api/v1/admin/profile',
    method: 'get'
  })
}

export function updateProfile(data) {
  return request({
    url: '/api/v1/admin/profile',
    method: 'put',
    data
  })
}

