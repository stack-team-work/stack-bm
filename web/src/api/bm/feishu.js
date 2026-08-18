import request from '../../utils/request'

export function getFeishuAppList(params) {
  return request.post('/feishu-app/list', params)
}

export function getFeishuAppAll() {
  return request.post('/feishu-app/all')
}

export function createFeishuApp(data) {
  return request.post('/feishu-app/create', data)
}

export function updateFeishuApp(id, data) {
  return request.post(`/feishu-app/update/${id}`, data)
}

export function updateFeishuAppStatus(id, data) {
  return request.post(`/feishu-app/status/${id}`, data)
}

export function getFeishuChatList(params) {
  return request.post('/feishu-chat/list', params)
}

export function getFeishuChatAll() {
  return request.post('/feishu-chat/all')
}

export function createFeishuChat(data) {
  return request.post('/feishu-chat/create', data)
}

export function updateFeishuChat(id, data) {
  return request.post(`/feishu-chat/update/${id}`, data)
}

export function updateFeishuChatStatus(id, data) {
  return request.post(`/feishu-chat/status/${id}`, data)
}

export function getFeishuUserList(params) {
  return request.post('/feishu-user/list', params)
}

export function getFeishuUserAll() {
  return request.post('/feishu-user/all')
}

export function createFeishuUser(data) {
  return request.post('/feishu-user/create', data)
}

export function updateFeishuUser(id, data) {
  return request.post(`/feishu-user/update/${id}`, data)
}

export function updateFeishuUserStatus(id, data) {
  return request.post(`/feishu-user/status/${id}`, data)
}
