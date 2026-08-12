import request from '../../utils/request'

export function getTtAdTemplateList(params) {
  return request.post('/tt-ad-template/list', params)
}

export function getTtAdTemplateDetail(id) {
  return request.post(`/tt-ad-template/detail/${id}`)
}

export function createTtAdTemplate(data) {
  return request.post('/tt-ad-template/create', data)
}

export function updateTtAdTemplate(id, data) {
  return request.post(`/tt-ad-template/update/${id}`, data)
}

export function deleteTtAdTemplate(id) {
  return request.post(`/tt-ad-template/delete/${id}`)
}

export function copyTtAdTemplate(data) {
  return request.post('/tt-ad-template/copy', data)
}

export function getTtAudienceTemplateList(params) {
  return request.post('/tt-audience-template/list', params)
}

export function getTtAudienceTemplateDetail(id) {
  return request.post(`/tt-audience-template/detail/${id}`)
}

export function createTtAudienceTemplate(data) {
  return request.post('/tt-audience-template/create', data)
}

export function updateTtAudienceTemplate(id, data) {
  return request.post(`/tt-audience-template/update/${id}`, data)
}

export function deleteTtAudienceTemplate(id) {
  return request.post(`/tt-audience-template/delete/${id}`)
}

export function copyTtAudienceTemplate(data) {
  return request.post('/tt-audience-template/copy', data)
}

export function getTtTitleTemplateList(params) {
  return request.post('/tt-title-template/list', params)
}

export function getTtTitleTemplateDetail(id) {
  return request.post(`/tt-title-template/detail/${id}`)
}

export function createTtTitleTemplate(data) {
  return request.post('/tt-title-template/create', data)
}

export function updateTtTitleTemplate(id, data) {
  return request.post(`/tt-title-template/update/${id}`, data)
}

export function deleteTtTitleTemplate(id) {
  return request.post(`/tt-title-template/delete/${id}`)
}

export function copyTtTitleTemplate(data) {
  return request.post('/tt-title-template/copy', data)
}

export function getTtWordList() {
  return request.post('/tt-word-list/list')
}
