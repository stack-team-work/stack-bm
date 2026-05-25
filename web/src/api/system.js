import request from '../utils/request'

export function getDashboardStats() {
  return request.post('/dashboard/stats')
}

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

export function getMenuList(params) {
  return request.post('/menu/list', params)
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
