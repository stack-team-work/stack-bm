import request from '../utils/request'

export function login(username, password) {
  return request.post('/login', { username, password })
}

export function getUserInfo() {
  return request.post('/user/info')
}

export function getDict() {
  return request.post('/dict')
}

export function getDictByKey(key) {
  return request.post(`/dict/${key}`)
}

export function getOptions(type) {
  return request.post('/options', { type })
}
