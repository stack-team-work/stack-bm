import { getOptions } from '../api/index'

const cache = {}

export function useOptions() {
  async function loadOptions(type) {
    if (cache[type]) return cache[type]
    try {
      const res = await getOptions(type)
      cache[type] = res.data || []
      return cache[type]
    } catch { return [] }
  }

  return { loadOptions }
}
