import request from '../utils/request'

export function getUserInfoList(params) {
  return request.post('/user/info/list', params)
}

export function getUserOrderList(params) {
  return request.post('/user/orders/list', params)
}

export function getUserLoginList(params) {
  return request.post('/user/logins/list', params)
}

export function getUserActiveList(params) {
  return request.post('/user/actives/list', params)
}