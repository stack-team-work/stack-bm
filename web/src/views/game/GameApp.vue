<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/包名" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchGameId" :options="gameSearchOptions" placeholder="父游戏" clearable style="width: 150px" @update:value="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
  </div>
</template>

<script setup>
import { ref, h, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { NButton, NSpace, NSwitch, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { useDict } from '../../composables/useDict'
import { getGameAppList, deleteGameApp, updateGameApp, getGameAll } from '../../api/game'
import { formatTime } from '../../utils/format'

const router = useRouter()
const message = useMessage()
const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getGameAppList)
const { handleDelete: doDelete } = useModal()

const searchKeyword = ref('')
const searchGameId = ref(null)
const searchStatus = ref(null)
const gameOptions = ref([])
const { load: loadDict, options } = useDict()
const statusOptions = computed(() => options('status'))
const gameSearchOptions = computed(() => [{ label: '全部游戏', value: 0 }, ...gameOptions.value])

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '父游戏', key: 'pid', width: 90, render: (row) => { const g = gameOptions.value.find(o => o.value === row.pid); return g ? g.label : row.pid } },
  { title: '子游戏名称', key: 'name' },
  { title: '包名', key: 'package_name' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 180, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => router.push(`/game-app/edit/${row.id}`) }, { default: () => '编辑' }),
    h(NButton, { size: 'tiny', onClick: () => handleExport(row) }, { default: () => '导出' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value, game_id: searchGameId.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { router.push('/game-app/create') }
async function onDelete(id) { if (await doDelete(id, deleteGameApp)) search({ keyword: searchKeyword.value, game_id: searchGameId.value ?? 0, status: searchStatus.value ?? -1 }) }

function handleExport(row) {
  const content = `子游戏ID: ${row.id}\n子游戏名称: ${row.name}\nAppKey: ${row.app_key}\nAppSecret: ${row.app_secret}`
  const bom = '\uFEFF'
  const blob = new Blob([bom + content], { type: 'text/plain;charset=utf-8;' })
  const url = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `game_app_${row.id}.txt`
  link.click()
  URL.revokeObjectURL(url)
  message.success('导出成功')
}

async function handleStatusChange(row, val) {
  try {
    await updateGameApp(row.id, { ...row, status: val ? 1 : 0 })
    row.status = val ? 1 : 0
    message.success('状态已更新')
  } catch {
    message.error('更新失败')
  }
}

async function loadGames() { try { const res = await getGameAll(); gameOptions.value = (res.data || []).map(g => ({ label: g.name, value: g.id })) } catch { /* */ } }
onMounted(async () => { await loadDict(); loadGames(); search({}) })
</script>
