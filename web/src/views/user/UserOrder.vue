<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchUserId" placeholder="用户ID" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchAppId" :options="appOptions" placeholder="子游戏" clearable style="width: 160px" @update:value="doSearch" />
          <n-select v-model:value="searchPayStatus" :options="payStatusOptions" placeholder="支付状态" clearable style="width: 130px" @update:value="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="订单状态" clearable style="width: 130px" @update:value="doSearch" />
          <n-date-picker v-model:value="searchRange" type="datetimerange" clearable style="width: 340px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
  </div>
</template>

<script setup>
import { ref, h, onMounted, computed } from 'vue'
import { NTag } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useOptions } from '../../composables/useOptions'
import { useDict } from '../../composables/useDict'
import { getUserOrderList } from '../../api/user'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getUserOrderList)
const searchUserId = ref('')
const searchAppId = ref(null)
const searchPayStatus = ref(null)
const searchStatus = ref(null)
const searchRange = ref(null)
const appOptions = ref([])
const { loadOptions } = useOptions()
const { load: loadDict, options } = useDict()

const payStatusOptions = computed(() => options('user_order_pay_status'))
const statusOptions = computed(() => options('user_order_status'))
const payStatusLabel = { 1: '待支付', 2: '支付成功', 3: '支付失败' }
const statusLabel = { 1: '待同步', 2: '已同步', 3: '同步失败' }
const payWayLabel = { 1: 'SDK', 2: '支付宝', 3: '微信', 4: 'H5' }

const columns = [
  { title: 'ID', key: 'id', width: 70 },
  { title: '订单号', key: 'order_num', width: 180, ellipsis: { tooltip: true } },
  { title: '三方订单号', key: 'third_order_num', width: 170, ellipsis: { tooltip: true } },
  { title: '用户ID', key: 'user_id', width: 180, ellipsis: { tooltip: true } },
  { title: '子游戏', key: 'app_id', width: 70 },
  { title: '服务器', key: 'server_name', width: 100, ellipsis: { tooltip: true } },
  { title: '角色', key: 'role_name', width: 100, ellipsis: { tooltip: true } },
  { title: '商品', key: 'product', width: 90 },
  { title: '金额', key: 'total_fee', width: 100, render: (row) => row.total_fee + (row.currency || '') },
  { title: '支付状态', key: 'pay_status', width: 90, render: (row) => h(NTag, { type: row.pay_status === 2 ? 'success' : row.pay_status === 3 ? 'error' : 'default', size: 'small' }, { default: () => payStatusLabel[row.pay_status] || row.pay_status }) },
  { title: '状态', key: 'status', width: 90, render: (row) => h(NTag, { type: row.status === 2 ? 'success' : row.status === 3 ? 'error' : 'default', size: 'small' }, { default: () => statusLabel[row.status] || row.status }) },
  { title: '支付方式', key: 'pay_way', width: 80, render: (row) => payWayLabel[row.pay_way] || row.pay_way },
  { title: '首充', key: 'is_first', width: 70, render: (row) => (row.is_first === 1 ? '是' : '否') },
  { title: '支付时间', key: 'pay_at', width: 160, render: (row) => formatTime(row.pay_at) },
  { title: '创建时间', key: 'created_at', width: 160, render: (row) => formatTime(row.created_at) },
]

function doSearch() {
  const [start, end] = searchRange.value || []
  search({
    user_id: searchUserId.value,
    app_id: searchAppId.value ?? 0,
    pay_status: searchPayStatus.value ?? 0,
    status: searchStatus.value ?? 0,
    start_at: start ? Math.floor(start / 1000) : 0,
    end_at: end ? Math.floor(end / 1000) : 0,
  })
}

onMounted(async () => {
  await loadDict()
  appOptions.value = await loadOptions('game_app')
  search({})
})
</script>