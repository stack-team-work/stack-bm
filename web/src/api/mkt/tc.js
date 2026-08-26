import request from '../../utils/request'

export function getTcAdDataList(level, params) {
  return request.post(`/tc-ad-data/${level}/list`, params)
}

export function tcTool(level, action, data) {
  return request.post(`/tc-tool/${level}/${action}`, data)
}