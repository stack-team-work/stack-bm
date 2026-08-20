<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchUserId" placeholder="用户ID" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchAppId" :options="appOptions" placeholder="子游戏" clearable style="width: 160px" @update:value="doSearch" />
          <n-date-picker v-model:value="searchRange" type="datetimerange" clearable style="width: 340px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
  </div>
</template>

<script setup>
import { ref, h, onMounted } from 'vue'
import { NTag } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useOptions } from '../../composables/useOptions'
import { getUserInfoList } from '../../api/user'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getUserInfoList)
const searchUserId = ref('')
const searchAppId = ref(null)
const searchRange = ref(null)
const appOptions = ref([])
const { loadOptions } = useOptions()

const yesNo = (v) => (v === 1 ? '是' : '否')
const osLabel = { 1: '安卓', 2: 'iOS' }

const columns = [
  { title: 'ID', key: 'id', width: 70 },
  { title: '用户ID', key: 'user_id', width: 180, ellipsis: { tooltip: true } },
  { title: '父游戏', key: 'pid', width: 70 },
  { title: '子游戏', key: 'app_id', width: 70 },
  { title: '账号', key: 'account', width: 130, ellipsis: { tooltip: true } },
  { title: '昵称', key: 'nickname', width: 110, ellipsis: { tooltip: true } },
  { title: '手机', key: 'phone', width: 120 },
  { title: '认证', key: 'is_verify', width: 70, render: (row) => h(NTag, { type: row.is_verify === 1 ? 'success' : 'default', size: 'small' }, { default: () => yesNo(row.is_verify) }) },
  { title: '充值', key: 'is_charge', width: 70, render: (row) => h(NTag, { type: row.is_charge === 1 ? 'warning' : 'default', size: 'small' }, { default: () => yesNo(row.is_charge) }) },
  { title: 'VIP', key: 'vip_level', width: 60 },
  { title: '金币', key: 'coin', width: 80 },
  { title: '注册来源', key: 'reg_from', width: 80 },
  { title: '媒体ID', key: 'media_id', width: 70 },
  { title: 'IP', key: 'ip', width: 130 },
  { title: '省市', key: 'province', width: 120, render: (row) => `${row.province || ''} ${row.city || ''}`.trim() },
  { title: '系统', key: 'os', width: 70, render: (row) => osLabel[row.os] || row.os },
  { title: '注册时间', key: 'created_at', width: 160, render: (row) => formatTime(row.created_at) },
]

function doSearch() {
  const [start, end] = searchRange.value || []
  search({
    user_id: searchUserId.value,
    app_id: searchAppId.value ?? 0,
    start_at: start ? Math.floor(start / 1000) : 0,
    end_at: end ? Math.floor(end / 1000) : 0,
  })
}

onMounted(async () => {
  appOptions.value = await loadOptions('game_app')
  search({})
})
</script>