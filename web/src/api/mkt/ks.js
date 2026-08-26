import request from '../../utils/request'

export function getKsAdTemplateList(params) {
  return request.post('/ks-ad-template/list', params)
}

export function getKsAdTemplateDetail(id) {
  return request.post(`/ks-ad-template/detail/${id}`)
}

export function createKsAdTemplate(data) {
  return request.post('/ks-ad-template/create', data)
}

export function updateKsAdTemplate(id, data) {
  return request.post(`/ks-ad-template/update/${id}`, data)
}

export function deleteKsAdTemplate(id) {
  return request.post(`/ks-ad-template/delete/${id}`)
}

export function copyKsAdTemplate(data) {
  return request.post('/ks-ad-template/copy', data)
}

export function getKsAudienceTemplateList(params) {
  return request.post('/ks-audience-template/list', params)
}

export function getKsAudienceTemplateDetail(id) {
  return request.post(`/ks-audience-template/detail/${id}`)
}

export function createKsAudienceTemplate(data) {
  return request.post('/ks-audience-template/create', data)
}

export function updateKsAudienceTemplate(id, data) {
  return request.post(`/ks-audience-template/update/${id}`, data)
}

export function deleteKsAudienceTemplate(id) {
  return request.post(`/ks-audience-template/delete/${id}`)
}

export function copyKsAudienceTemplate(data) {
  return request.post('/ks-audience-template/copy', data)
}

export function getKsTitleTemplateList(params) {
  return request.post('/ks-title-template/list', params)
}

export function getKsTitleTemplateDetail(id) {
  return request.post(`/ks-title-template/detail/${id}`)
}

export function createKsTitleTemplate(data) {
  return request.post('/ks-title-template/create', data)
}

export function updateKsTitleTemplate(id, data) {
  return request.post(`/ks-title-template/update/${id}`, data)
}

export function deleteKsTitleTemplate(id) {
  return request.post(`/ks-title-template/delete/${id}`)
}

export function copyKsTitleTemplate(data) {
  return request.post('/ks-title-template/copy', data)
}

export function getKsAdDataList(level, params) {
  return request.post(`/ks-ad-data/${level}/list`, params)
}

export function ksTool(level, action, data) {
  return request.post(`/ks-tool/${level}/${action}`, data)
}
