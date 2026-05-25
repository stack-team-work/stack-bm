import { ref, reactive } from 'vue'
import { useMessage } from 'naive-ui'

export function useTable(fetchFn) {
  const message = useMessage()
  const loading = ref(false)
  const tableData = ref([])

  const pagination = reactive({
    page: 1,
    pageSize: 10,
    itemCount: 0,
    showSizePicker: true,
    pageSizes: [10, 20, 50],
  })

  const searchParams = reactive({})

  async function loadData() {
    loading.value = true
    try {
      const res = await fetchFn({
        page: pagination.page,
        size: pagination.pageSize,
        ...searchParams,
      })
      tableData.value = res.data.list || []
      pagination.itemCount = res.data.total || 0
    } catch (err) {
      message.error(err.message || '获取数据失败')
    } finally {
      loading.value = false
    }
  }

  function handlePageChange(page) {
    pagination.page = page
    loadData()
  }

  function handlePageSizeChange(size) {
    pagination.pageSize = size
    pagination.page = 1
    loadData()
  }

  function search(params) {
    Object.keys(searchParams).forEach(k => delete searchParams[k])
    Object.assign(searchParams, params)
    pagination.page = 1
    loadData()
  }

  function resetSearch() {
    Object.keys(searchParams).forEach(k => delete searchParams[k])
    pagination.page = 1
    loadData()
  }

  return { loading, tableData, pagination, searchParams, loadData, search, resetSearch, handlePageChange, handlePageSizeChange }
}
