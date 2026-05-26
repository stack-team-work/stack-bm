import request from '../utils/request'

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
