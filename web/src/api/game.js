import request from '../utils/request'

export function getGameList(params) {
  return request.post('/game/list', params)
}

export function getGameAll() {
  return request.post('/game/all')
}

export function createGame(data) {
  return request.post('/game/create', data)
}

export function updateGame(id, data) {
  return request.post(`/game/update/${id}`, data)
}

export function deleteGame(id) {
  return request.post(`/game/delete/${id}`)
}

export function getGameAppList(params) {
  return request.post('/game-app/list', params)
}

export function getGameAppAll() {
  return request.post('/game-app/all')
}

export function createGameApp(data) {
  return request.post('/game-app/create', data)
}

export function updateGameApp(id, data) {
  return request.post(`/game-app/update/${id}`, data)
}

export function deleteGameApp(id) {
  return request.post(`/game-app/delete/${id}`)
}

export function getGameAppTemplateList(params) {
  return request.post('/game-app-template/list', params)
}

export function createGameAppTemplate(data) {
  return request.post('/game-app-template/create', data)
}

export function updateGameAppTemplate(id, data) {
  return request.post(`/game-app-template/update/${id}`, data)
}

export function deleteGameAppTemplate(id) {
  return request.post(`/game-app-template/delete/${id}`)
}

export function getGameCpList(params) {
  return request.post('/game-cp/list', params)
}

export function getGameCpAll() {
  return request.post('/game-cp/all')
}

export function createGameCp(data) {
  return request.post('/game-cp/create', data)
}

export function updateGameCp(id, data) {
  return request.post(`/game-cp/update/${id}`, data)
}

export function deleteGameCp(id) {
  return request.post(`/game-cp/delete/${id}`)
}

export function getGameTagList(params) {
  return request.post('/game-tag/list', params)
}

export function getGameTagAll(params) {
  return request.post('/game-tag/all', params)
}

export function createGameTag(data) {
  return request.post('/game-tag/create', data)
}

export function updateGameTag(id, data) {
  return request.post(`/game-tag/update/${id}`, data)
}

export function deleteGameTag(id) {
  return request.post(`/game-tag/delete/${id}`)
}

export function getGameVariableList(params) {
  return request.post('/game-variable/list', params)
}

export function createGameVariable(data) {
  return request.post('/game-variable/create', data)
}

export function updateGameVariable(id, data) {
  return request.post(`/game-variable/update/${id}`, data)
}

export function deleteGameVariable(id) {
  return request.post(`/game-variable/delete/${id}`)
}

export function getGamePlatformList(params) {
  return request.post('/game-platform/list', params)
}

export function getGamePlatformAll() {
  return request.post('/game-platform/all')
}

export function createGamePlatform(data) {
  return request.post('/game-platform/create', data)
}

export function updateGamePlatform(id, data) {
  return request.post(`/game-platform/update/${id}`, data)
}

export function deleteGamePlatform(id) {
  return request.post(`/game-platform/delete/${id}`)
}

export function getSdkLogList(params) {
  return request.post('/sdk-logs/list', params)
}

export function getGameGiftList(params) {
  return request.post('/game-gift/list', params)
}

export function getGameGiftAll() {
  return request.post('/game-gift/all')
}

export function createGameGift(data) {
  return request.post('/game-gift/create', data)
}

export function updateGameGift(id, data) {
  return request.post(`/game-gift/update/${id}`, data)
}

export function deleteGameGift(id) {
  return request.post(`/game-gift/delete/${id}`)
}

export function getGameGiftCodeList(params) {
  return request.post('/game-gift-code/list', params)
}

export function createGameGiftCode(data) {
  return request.post('/game-gift-code/create', data)
}

export function updateGameGiftCode(id, data) {
  return request.post(`/game-gift-code/update/${id}`, data)
}

export function deleteGameGiftCode(id) {
  return request.post(`/game-gift-code/delete/${id}`)
}

export function getGameGiftUserCodeList(params) {
  return request.post('/game-gift-user-code/list', params)
}

export function getGameVoucherList(params) {
  return request.post('/game-voucher/list', params)
}

export function getGameVoucherAll() {
  return request.post('/game-voucher/all')
}

export function createGameVoucher(data) {
  return request.post('/game-voucher/create', data)
}

export function updateGameVoucher(id, data) {
  return request.post(`/game-voucher/update/${id}`, data)
}

export function deleteGameVoucher(id) {
  return request.post(`/game-voucher/delete/${id}`)
}

export function getGameVoucherUseList(params) {
  return request.post('/game-voucher-use/list', params)
}
