import request from '../../utils/request'

export function getBiliAdTemplateList(params) {
  return request.post('/bili-ad-template/list', params)
}

export function getBiliAdTemplateDetail(id) {
  return request.post(`/bili-ad-template/detail/${id}`)
}

export function createBiliAdTemplate(data) {
  return request.post('/bili-ad-template/create', data)
}

export function updateBiliAdTemplate(id, data) {
  return request.post(`/bili-ad-template/update/${id}`, data)
}

export function deleteBiliAdTemplate(id) {
  return request.post(`/bili-ad-template/delete/${id}`)
}

export function copyBiliAdTemplate(data) {
  return request.post('/bili-ad-template/copy', data)
}

export function getBiliAudienceTemplateList(params) {
  return request.post('/bili-audience-template/list', params)
}

export function getBiliAudienceTemplateDetail(id) {
  return request.post(`/bili-audience-template/detail/${id}`)
}

export function createBiliAudienceTemplate(data) {
  return request.post('/bili-audience-template/create', data)
}

export function updateBiliAudienceTemplate(id, data) {
  return request.post(`/bili-audience-template/update/${id}`, data)
}

export function deleteBiliAudienceTemplate(id) {
  return request.post(`/bili-audience-template/delete/${id}`)
}

export function copyBiliAudienceTemplate(data) {
  return request.post('/bili-audience-template/copy', data)
}

export function getBiliTitleTemplateList(params) {
  return request.post('/bili-title-template/list', params)
}

export function getBiliTitleTemplateDetail(id) {
  return request.post(`/bili-title-template/detail/${id}`)
}

export function createBiliTitleTemplate(data) {
  return request.post('/bili-title-template/create', data)
}

export function updateBiliTitleTemplate(id, data) {
  return request.post(`/bili-title-template/update/${id}`, data)
}

export function deleteBiliTitleTemplate(id) {
  return request.post(`/bili-title-template/delete/${id}`)
}

export function copyBiliTitleTemplate(data) {
  return request.post('/bili-title-template/copy', data)
}
