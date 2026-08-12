<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索路径/用户/描述" clearable style="width: 250px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchLevel" :options="levelOptions" placeholder="级别" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-popconfirm @positive-click="handleClear">
          <template #trigger>
            <n-button type="error" size="small" :loading="clearing">清空日志</n-button>
          </template>
          确认清空所有操作日志？
        </n-popconfirm>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
  </div>
</template>

<script setup>
import { ref, h, onMounted, computed } from 'vue'
import { NTag, NPopconfirm, NButton, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { getLogList, clearLogs } from '../../../api/bm/sys'
import { formatTime } from '../../../utils/format'
import { useDict } from '../../../composables/useDict'

const message = useMessage()
const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getLogList)
const { load: loadDict, options } = useDict()

const searchKeyword = ref('')
const searchLevel = ref('')
const clearing = ref(false)
const levelOptions = computed(() => options('bm_log_level'))

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '级别', key: 'level', width: 70, render: (row) => h(NTag, { type: row.level === 2 ? 'error' : 'default', size: 'small' }, { default: () => row.level === 2 ? 'error' : 'info' }) },
  { title: '路径', key: 'path', width: 180 },
  { title: '用户', key: 'username', width: 90 },
  { title: 'IP', key: 'ip', width: 130 },
  { title: '描述', key: 'desc', ellipsis: { tooltip: true } },
  { title: '时间', key: 'created_at', width: 160, render: (row) => formatTime(row.created_at) },
]

function doSearch() { search({ keyword: searchKeyword.value, level: searchLevel.value }) }

async function handleClear() {
  clearing.value = true
  try {
    await clearLogs()
    message.success('日志已清空')
    search({})
  } catch {
    message.error('清空失败')
  } finally {
    clearing.value = false
  }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
