<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索用户/角色" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchVoucherId" :options="voucherOptions" placeholder="代金券" clearable style="width: 160px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
  </div>
</template>

<script setup>
import { ref, h, onMounted } from 'vue'
import { NTag } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useOptions } from '../../../composables/useOptions'
import { getGameVoucherUseList } from '../../../api/sdk/game'
import { formatTime } from '../../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getGameVoucherUseList)
const searchKeyword = ref('')
const searchVoucherId = ref(null)
const voucherOptions = ref([])
const { loadOptions } = useOptions()
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '券ID', key: 'voucher_id', width: 70 },
  { title: '用户ID', key: 'user_id', width: 100 },
  { title: '角色', key: 'role_name', width: 100 },
  { title: '服务器', key: 'server_name', width: 100 },
  { title: '是否使用', key: 'is_use', width: 80, render: (row) => h(NTag, { type: row.is_use === 1 ? 'success' : 'default', size: 'small' }, { default: () => row.is_use === 1 ? '已用' : '未用' }) },
  { title: '开始时间', key: 'stime', width: 160, render: (row) => formatTime(row.stime) },
  { title: '结束时间', key: 'etime', width: 160, render: (row) => formatTime(row.etime) },
]
function doSearch() { search({ keyword: searchKeyword.value, voucher_id: searchVoucherId.value ?? 0 }) }
onMounted(async () => { voucherOptions.value = await loadOptions('game_voucher'); search({}) })
</script>
