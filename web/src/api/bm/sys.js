import request from '../../utils/request'

export function getAdminList(params) {
  return request.post('/admin/list', params)
}

export function createAdmin(data) {
  return request.post('/admin/create', data)
}

export function updateAdmin(id, data) {
  return request.post(`/admin/update/${id}`, data)
}

export function deleteAdmin(id) {
  return request.post(`/admin/delete/${id}`)
}

export function getAdminGroupList(params) {
  return request.post('/admin-group/list', params)
}

export function getAdminGroupAll() {
  return request.post('/admin-group/all')
}

export function createAdminGroup(data) {
  return request.post('/admin-group/create', data)
}

export function updateAdminGroup(id, data) {
  return request.post(`/admin-group/update/${id}`, data)
}

export function deleteAdminGroup(id) {
  return request.post(`/admin-group/delete/${id}`)
}

export function getLogList(params) {
  return request.post('/logs/list', params)
}

export function clearLogs() {
  return request.post('/logs/clear')
}

export function getDashboardStats() {
  return request.post('/dashboard/stats')
}

export function getMenuAll() {
  return request.post('/menu/all')
}

export function createMenu(data) {
  return request.post('/menu/create', data)
}

export function updateMenu(id, data) {
  return request.post(`/menu/update/${id}`, data)
}

export function deleteMenu(id) {
  return request.post(`/menu/delete/${id}`)
}

export function getSysColumnList(params) {
  return request.post('/sys-column/list', params)
}

export function getSysColumnAll() {
  return request.post('/sys-column/all')
}

export function createSysColumn(data) {
  return request.post('/sys-column/create', data)
}

export function updateSysColumn(id, data) {
  return request.post(`/sys-column/update/${id}`, data)
}

export function deleteSysColumn(id) {
  return request.post(`/sys-column/delete/${id}`)
}

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
