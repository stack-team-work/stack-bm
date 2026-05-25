<template>
  <div>
    <n-card title="操作日志" :bordered="false">
      <n-space vertical :size="16">
        <n-space>
          <n-input v-model:value="searchKeyword" placeholder="搜索路径/用户/描述" clearable style="width: 250px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchLevel" :options="levelOptions" placeholder="级别" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="primary" @click="doSearch">搜索</n-button>
          <n-button @click="resetAll">重置</n-button>
        </n-space>

        <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
      </n-space>
    </n-card>
  </div>
</template>

<script setup>
import { ref, h, onMounted } from 'vue'
import { NTag } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { getLogList } from '../../api/system'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getLogList)

const searchKeyword = ref('')
const searchLevel = ref('')
const levelOptions = [
  { label: '全部', value: '' },
  { label: 'error', value: 'error' },
  { label: 'warn', value: 'warn' },
  { label: 'info', value: 'info' },
]

const columns = [
  { title: 'ID', key: 'id', width: 70 },
  { title: '级别', key: 'level', width: 80, render: (row) => h(NTag, { type: row.level === 'error' ? 'error' : 'default', size: 'small' }, { default: () => row.level }) },
  { title: '路径', key: 'path', width: 200 },
  { title: '用户', key: 'username', width: 100 },
  { title: 'IP', key: 'ip', width: 140 },
  { title: '描述', key: 'desc', ellipsis: { tooltip: true } },
  { title: '时间', key: 'created_at', width: 180 },
]

function doSearch() { search({ keyword: searchKeyword.value, level: searchLevel.value }) }
function resetAll() { searchKeyword.value = ''; searchLevel.value = ''; resetSearch() }

onMounted(() => search({}))
</script>
