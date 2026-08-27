import request from '../../utils/request'

export function getKsAdTemplateList(params) {
  return request.post('/ks/v1/template/ad/list', params)
}

export function getKsAdTemplateDetail(id) {
  return request.post(`/ks/v1/template/ad/detail/${id}`)
}

export function createKsAdTemplate(data) {
  return request.post('/ks/v1/template/ad/create', data)
}

export function updateKsAdTemplate(id, data) {
  return request.post(`/ks/v1/template/ad/update/${id}`, data)
}

export function deleteKsAdTemplate(id) {
  return request.post(`/ks/v1/template/ad/delete/${id}`)
}

export function copyKsAdTemplate(data) {
  return request.post('/ks/v1/template/ad/copy', data)
}

export function getKsAudienceTemplateList(params) {
  return request.post('/ks/v1/template/audience/list', params)
}

export function getKsAudienceTemplateDetail(id) {
  return request.post(`/ks/v1/template/audience/detail/${id}`)
}

export function createKsAudienceTemplate(data) {
  return request.post('/ks/v1/template/audience/create', data)
}

export function updateKsAudienceTemplate(id, data) {
  return request.post(`/ks/v1/template/audience/update/${id}`, data)
}

export function deleteKsAudienceTemplate(id) {
  return request.post(`/ks/v1/template/audience/delete/${id}`)
}

export function copyKsAudienceTemplate(data) {
  return request.post('/ks/v1/template/audience/copy', data)
}

export function getKsTitleTemplateList(params) {
  return request.post('/ks/v1/template/title/list', params)
}

export function getKsTitleTemplateDetail(id) {
  return request.post(`/ks/v1/template/title/detail/${id}`)
}

export function createKsTitleTemplate(data) {
  return request.post('/ks/v1/template/title/create', data)
}

export function updateKsTitleTemplate(id, data) {
  return request.post(`/ks/v1/template/title/update/${id}`, data)
}

export function deleteKsTitleTemplate(id) {
  return request.post(`/ks/v1/template/title/delete/${id}`)
}

export function copyKsTitleTemplate(data) {
  return request.post('/ks/v1/template/title/copy', data)
}

export function getKsAdDataList(level, params) {
  return request.post(`/ks/v1/ad/${level}/list`, params)
}

export function ksTool(level, action, data) {
  return request.post(`/ks/v1/ad/${level}/${action}`, data)
}
