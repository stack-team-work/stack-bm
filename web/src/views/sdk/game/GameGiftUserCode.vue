<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索激活码/用户/角色" clearable style="width: 220px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchGiftId" :options="giftOptions" placeholder="礼包" clearable style="width: 160px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useTable } from '../../composables/useTable'
import { useOptions } from '../../composables/useOptions'
import { getGameGiftUserCodeList } from '../../api/game'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getGameGiftUserCodeList)
const searchKeyword = ref('')
const searchGiftId = ref(null)
const giftOptions = ref([])
const { loadOptions } = useOptions()
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '激活码', key: 'code', width: 130 },
  { title: '礼包ID', key: 'gift_id', width: 70 },
  { title: '用户ID', key: 'user_id', width: 100 },
  { title: '角色', key: 'role_name', width: 100 },
  { title: '服务器', key: 'server_name', width: 100 },
  { title: '子游戏ID', key: 'app_id', width: 80 },
  { title: '使用时间', key: 'created_at', width: 160, render: (row) => formatTime(row.created_at) },
]
function doSearch() { search({ keyword: searchKeyword.value, gift_id: searchGiftId.value ?? 0 }) }
onMounted(async () => { giftOptions.value = await loadOptions('game_gift'); search({}) })
</script>
