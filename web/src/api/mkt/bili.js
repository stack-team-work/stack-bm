import request from '../../utils/request'

export function getBiliAdTemplateList(params) {
  return request.post('/bili/v1/template/ad/list', params)
}

export function getBiliAdTemplateDetail(id) {
  return request.post(`/bili/v1/template/ad/detail/${id}`)
}

export function createBiliAdTemplate(data) {
  return request.post('/bili/v1/template/ad/create', data)
}

export function updateBiliAdTemplate(id, data) {
  return request.post(`/bili/v1/template/ad/update/${id}`, data)
}

export function deleteBiliAdTemplate(id) {
  return request.post(`/bili/v1/template/ad/delete/${id}`)
}

export function copyBiliAdTemplate(data) {
  return request.post('/bili/v1/template/ad/copy', data)
}

export function getBiliAudienceTemplateList(params) {
  return request.post('/bili/v1/template/audience/list', params)
}

export function getBiliAudienceTemplateDetail(id) {
  return request.post(`/bili/v1/template/audience/detail/${id}`)
}

export function createBiliAudienceTemplate(data) {
  return request.post('/bili/v1/template/audience/create', data)
}

export function updateBiliAudienceTemplate(id, data) {
  return request.post(`/bili/v1/template/audience/update/${id}`, data)
}

export function deleteBiliAudienceTemplate(id) {
  return request.post(`/bili/v1/template/audience/delete/${id}`)
}

export function copyBiliAudienceTemplate(data) {
  return request.post('/bili/v1/template/audience/copy', data)
}

export function getBiliTitleTemplateList(params) {
  return request.post('/bili/v1/template/title/list', params)
}

export function getBiliTitleTemplateDetail(id) {
  return request.post(`/bili/v1/template/title/detail/${id}`)
}

export function createBiliTitleTemplate(data) {
  return request.post('/bili/v1/template/title/create', data)
}

export function updateBiliTitleTemplate(id, data) {
  return request.post(`/bili/v1/template/title/update/${id}`, data)
}

export function deleteBiliTitleTemplate(id) {
  return request.post(`/bili/v1/template/title/delete/${id}`)
}

export function copyBiliTitleTemplate(data) {
  return request.post('/bili/v1/template/title/copy', data)
}

export function getBiliAdDataList(level, params) {
  return request.post(`/bili/v1/ad/${level}/list`, params)
}

export function biliTool(level, action, data) {
  return request.post(`/bili/v1/ad/${level}/${action}`, data)
}
