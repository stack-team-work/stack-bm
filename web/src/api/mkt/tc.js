import request from '../../utils/request'

export function getTcAdDataList(level, params) {
  return request.post(`/tc/v1/ad/${level}/list`, params)
}

export function tcTool(level, action, data) {
  return request.post(`/tc/v1/ad/${level}/${action}`, data)
}