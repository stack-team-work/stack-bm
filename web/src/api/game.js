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

export function createGameApp(data) {
  return request.post('/game-app/create', data)
}

export function updateGameApp(id, data) {
  return request.post(`/game-app/update/${id}`, data)
}

export function deleteGameApp(id) {
  return request.post(`/game-app/delete/${id}`)
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
