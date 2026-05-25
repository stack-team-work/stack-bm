<template>
  <div>
    <n-card title="子游戏管理" :bordered="false">
      <template #header-extra><n-button type="primary" @click="handleAdd">新增应用</n-button></template>

      <n-space vertical :size="16">
        <n-space>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/包名" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchGameId" :options="gameSearchOptions" placeholder="游戏" clearable style="width: 150px" @update:value="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="primary" @click="doSearch">搜索</n-button>
          <n-button @click="resetAll">重置</n-button>
        </n-space>

        <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
      </n-space>
    </n-card>
  </div>
</template>

<script setup>
import { ref, h, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { NButton, NSpace, NSwitch, NPopconfirm } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { getGameAppList, deleteGameApp, getGameAll } from '../../api/game'

const router = useRouter()
const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getGameAppList)
const { handleDelete: doDelete } = useModal()

const searchKeyword = ref('')
const searchGameId = ref(null)
const searchStatus = ref(null)
const gameOptions = ref([])
const statusOptions = [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }]
const gameSearchOptions = computed(() => [{ label: '全部游戏', value: 0 }, ...gameOptions.value])

const columns = [
  { title: 'ID', key: 'id', width: 80 },
  { title: '父游戏', key: 'pid', width: 100, render: (row) => {
    const g = gameOptions.value.find(o => o.value === row.pid); return g ? g.label : row.pid
  } },
  { title: '应用名称', key: 'name' },
  { title: '包名', key: 'package_name' },
  { title: 'AppKey', key: 'app_key' },
  { title: '状态', key: 'status', width: 80, render: (row) => h(NSwitch, { value: row.status === 1, readonly: true, size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 180 },
  {
    title: '操作', key: 'actions', width: 160,
    render: (row) => h(NSpace, null, {
      default: () => [
        h(NButton, { size: 'small', onClick: () => router.push(`/game-app/edit/${row.id}`) }, { default: () => '编辑' }),
        h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, {
          default: () => '确认删除?', trigger: () => h(NButton, { size: 'small', type: 'error' }, { default: () => '删除' }),
        }),
      ],
    }),
  },
]

function doSearch() { search({ keyword: searchKeyword.value, game_id: searchGameId.value ?? 0, status: searchStatus.value ?? -1 }) }
function resetAll() { searchKeyword.value = ''; searchGameId.value = null; searchStatus.value = null; resetSearch() }
function handleAdd() { router.push('/game-app/create') }

async function onDelete(id) {
  const ok = await doDelete(id, deleteGameApp)
  if (ok) search({ keyword: searchKeyword.value, game_id: searchGameId.value ?? 0, status: searchStatus.value ?? -1 })
}

async function loadGames() {
  try { const res = await getGameAll(); gameOptions.value = (res.data || []).map(g => ({ label: g.name, value: g.id })) } catch { /* */ }
}

onMounted(() => { loadGames(); search({}) })
</script>
