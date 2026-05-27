import { ref } from 'vue'
import { getDict } from '../api/index'

const dict = ref({})
let loaded = false

export function useDict() {
  async function load() {
    if (loaded) return
    try {
      const res = await getDict()
      dict.value = res.data || {}
      loaded = true
    } catch { /* */ }
  }

  function options(key) {
    return dict.value[key] || []
  }

  return { dict, load, options }
}
