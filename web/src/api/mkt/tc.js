import request from '../../utils/request'

export function getTcAdDataList(level, params) {
  return request.post(`/tc-ad-data/${level}/list`, params)
}