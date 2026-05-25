import request from '../utils/request'

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
