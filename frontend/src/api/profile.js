import request from '@/utils/request'

export function getProfile(userId) {
  return request({
    url: `/api/profiles/get/${userId}`,
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

// 上传接口
export const uploadAvatar = (file, userId) => {
  const formData = new FormData()
  formData.append('file', file)
  formData.append('userId', userId)
  
  return request({
    url: '/api/profiles/avatar/upload',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}