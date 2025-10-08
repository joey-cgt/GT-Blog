import axios from 'axios'
import { ElMessage } from 'element-plus'
import router from '../router'

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
    // 处理401 JWT认证失败错误
    if (error.response && error.response.status === 401) {
      // 清除登录状态
      localStorage.removeItem('token')
      localStorage.removeItem('isLoggedIn')
      // 显示错误消息
      ElMessage.error('登录已过期，请重新登录')
      // 跳转到登录页面
      router.push('/login')
    } else {
      ElMessage.error(error.message)
    }
    return Promise.reject(error)
  }
)

export default service