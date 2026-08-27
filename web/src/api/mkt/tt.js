import request from '../../utils/request'

export function getTtAdTemplateList(params) {
  return request.post('/tt/v1/template/ad/list', params)
}

export function getTtAdTemplateDetail(id) {
  return request.post(`/tt/v1/template/ad/detail/${id}`)
}

export function createTtAdTemplate(data) {
  return request.post('/tt/v1/template/ad/create', data)
}

export function updateTtAdTemplate(id, data) {
  return request.post(`/tt/v1/template/ad/update/${id}`, data)
}

export function deleteTtAdTemplate(id) {
  return request.post(`/tt/v1/template/ad/delete/${id}`)
}

export function copyTtAdTemplate(data) {
  return request.post('/tt/v1/template/ad/copy', data)
}

export function getTtAudienceTemplateList(params) {
  return request.post('/tt/v1/template/audience/list', params)
}

export function getTtAudienceTemplateDetail(id) {
  return request.post(`/tt/v1/template/audience/detail/${id}`)
}

export function createTtAudienceTemplate(data) {
  return request.post('/tt/v1/template/audience/create', data)
}

export function updateTtAudienceTemplate(id, data) {
  return request.post(`/tt/v1/template/audience/update/${id}`, data)
}

export function deleteTtAudienceTemplate(id) {
  return request.post(`/tt/v1/template/audience/delete/${id}`)
}

export function copyTtAudienceTemplate(data) {
  return request.post('/tt/v1/template/audience/copy', data)
}

export function getTtTitleTemplateList(params) {
  return request.post('/tt/v1/template/title/list', params)
}

export function getTtTitleTemplateDetail(id) {
  return request.post(`/tt/v1/template/title/detail/${id}`)
}

export function createTtTitleTemplate(data) {
  return request.post('/tt/v1/template/title/create', data)
}

export function updateTtTitleTemplate(id, data) {
  return request.post(`/tt/v1/template/title/update/${id}`, data)
}

export function deleteTtTitleTemplate(id) {
  return request.post(`/tt/v1/template/title/delete/${id}`)
}

export function copyTtTitleTemplate(data) {
  return request.post('/tt/v1/template/title/copy', data)
}

export function getTtWordList() {
  return request.post('/tt/v1/template/word/list')
}

export function getTtAdDataList(level, params) {
  return request.post(`/tt/v1/ad/${level}/list`, params)
}

export function ttTool(level, action, data) {
  return request.post(`/tt/v1/ad/${level}/${action}`, data)
}
