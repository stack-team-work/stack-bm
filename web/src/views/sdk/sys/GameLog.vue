<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索IP/描述" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchType" :options="typeOptions" placeholder="类型" clearable style="width: 120px" @update:value="doSearch" />
        <n-select v-model:value="searchLevel" :options="levelOptions" placeholder="级别" clearable style="width: 100px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
  </div>
</template>

<script setup>
import { ref, h, onMounted, computed } from 'vue'
import { NTag } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useDict } from '../../composables/useDict'
import { getSdkLogList } from '../../api/game'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getSdkLogList)

const searchKeyword = ref('')
const searchType = ref(null)
const searchLevel = ref(null)
const { load: loadDict, options } = useDict()

const typeOptions = computed(() => options('sdk_log_type'))
const levelOptions = computed(() => options('sdk_log_level'))

const typeLabel = { 1: '注册日志', 2: '登录日志', 3: '支付日志' }
const levelLabel = { 1: 'info', 2: 'warning', 3: 'error' }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '父游戏', key: 'pid', width: 70 },
  { title: '子游戏', key: 'app_id', width: 70 },
  { title: '类型', key: 'type', width: 90, render: (row) => h(NTag, { size: 'small' }, { default: () => typeLabel[row.type] || row.type }) },
  { title: '级别', key: 'level', width: 80, render: (row) => h(NTag, { type: row.level >= 3 ? 'error' : row.level >= 2 ? 'warning' : 'default', size: 'small' }, { default: () => levelLabel[row.level] || row.level }) },
  { title: 'IP', key: 'ip', width: 130 },
  { title: '描述', key: 'desc', ellipsis: { tooltip: true } },
  { title: '时间', key: 'create_time', width: 160, render: (row) => formatTime(row.create_time) },
]

function doSearch() {
  search({ keyword: searchKeyword.value, type: searchType.value ?? 0, level: searchLevel.value ?? 0 })
}

onMounted(async () => { await loadDict(); search({}) })
</script>
