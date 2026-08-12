import request from '../../utils/request'

export function getPayPlatformList(params) {
  return request.post('/pay-platform/list', params)
}

export function getPayPlatformAll() {
  return request.post('/pay-platform/all')
}

export function createPayPlatform(data) {
  return request.post('/pay-platform/create', data)
}

export function updatePayPlatform(id, data) {
  return request.post(`/pay-platform/update/${id}`, data)
}

export function deletePayPlatform(id) {
  return request.post(`/pay-platform/delete/${id}`)
}

export function getPayMerchantList(params) {
  return request.post('/pay-merchant/list', params)
}

export function getPayMerchantAll() {
  return request.post('/pay-merchant/all')
}

export function createPayMerchant(data) {
  return request.post('/pay-merchant/create', data)
}

export function updatePayMerchant(id, data) {
  return request.post(`/pay-merchant/update/${id}`, data)
}

export function deletePayMerchant(id) {
  return request.post(`/pay-merchant/delete/${id}`)
}
