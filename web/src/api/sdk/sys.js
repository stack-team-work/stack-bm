import request from '../../utils/request'

export function getSdkLogList(params) {
  return request.post('/sdk-logs/list', params)
}
