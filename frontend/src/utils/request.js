import axios from 'axios'
import { ElMessage } from 'element-plus'

// 创建axios实例
const service = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
  timeout: 10000
})

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 从localStorage获取token并添加到请求头
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data
    
    // 根据后端API结构调整判断逻辑
    if (res.code && res.code !== 200) {
      ElMessage.error(res.message || 'Error')
      return Promise.reject(new Error(res.message || 'Error'))
    }
    return res
  },
  error => {
    ElMessage.error(error.message)
    return Promise.reject(error)
  }
)

export default service