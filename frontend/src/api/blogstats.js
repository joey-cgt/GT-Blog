import request from '../utils/request'

export function getBlogStatistics() {
  return request({
    url: '/api/v1/statistics/overview',
    method: 'get', 
  })
}

export function getViewTrend(timeRange = '7d') {
  return request({
    url: '/api/v1/statistics/view-trend',
    method: 'get', 
    params: {
      range: timeRange
    }
  })
}