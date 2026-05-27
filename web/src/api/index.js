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
