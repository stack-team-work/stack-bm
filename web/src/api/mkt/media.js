import request from '../../utils/request'

export function getMediaList(params) {
  return request.post('/media/list', params)
}

export function getMediaAll() {
  return request.post('/media/all')
}

export function createMedia(data) {
  return request.post('/media/create', data)
}

export function updateMedia(id, data) {
  return request.post(`/media/update/${id}`, data)
}

export function deleteMedia(id) {
  return request.post(`/media/delete/${id}`)
}

export function getMediaSubList(params) {
  return request.post('/media-sub/list', params)
}

export function getMediaSubAll() {
  return request.post('/media-sub/all')
}

export function createMediaSub(data) {
  return request.post('/media-sub/create', data)
}

export function updateMediaSub(id, data) {
  return request.post(`/media-sub/update/${id}`, data)
}

export function deleteMediaSub(id) {
  return request.post(`/media-sub/delete/${id}`)
}

export function getMediaAgentList(params) {
  return request.post('/media-agent/list', params)
}

export function getMediaAgentAll() {
  return request.post('/media-agent/all')
}

export function createMediaAgent(data) {
  return request.post('/media-agent/create', data)
}

export function updateMediaAgent(id, data) {
  return request.post(`/media-agent/update/${id}`, data)
}

export function deleteMediaAgent(id) {
  return request.post(`/media-agent/delete/${id}`)
}

export function getMediaApplicationList(params) {
  return request.post('/media-application/list', params)
}

export function getMediaApplicationAll() {
  return request.post('/media-application/all')
}

export function createMediaApplication(data) {
  return request.post('/media-application/create', data)
}

export function updateMediaApplication(id, data) {
  return request.post(`/media-application/update/${id}`, data)
}

export function deleteMediaApplication(id) {
  return request.post(`/media-application/delete/${id}`)
}

export function getMediaManagerList(params) {
  return request.post('/media-manager/list', params)
}

export function getMediaManagerAll() {
  return request.post('/media-manager/all')
}

export function createMediaManager(data) {
  return request.post('/media-manager/create', data)
}

export function updateMediaManager(id, data) {
  return request.post(`/media-manager/update/${id}`, data)
}

export function deleteMediaManager(id) {
  return request.post(`/media-manager/delete/${id}`)
}

export function getMediaSubjectList(params) {
  return request.post('/media-subject/list', params)
}

export function getMediaSubjectAll() {
  return request.post('/media-subject/all')
}

export function createMediaSubject(data) {
  return request.post('/media-subject/create', data)
}

export function updateMediaSubject(id, data) {
  return request.post(`/media-subject/update/${id}`, data)
}

export function deleteMediaSubject(id) {
  return request.post(`/media-subject/delete/${id}`)
}
