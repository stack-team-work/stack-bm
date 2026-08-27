<template>
  <n-card :bordered="false">
    <n-space v-if="batchActions.length || searchable" :size="8" style="margin-bottom: 12px" align="center" wrap>
      <n-date-picker v-if="searchable" v-model:value="dateRange" type="daterange" clearable style="width: 240px" placeholder="报表日期" />
      <n-input v-if="searchable" v-model:value="searchKeyword" placeholder="名称/账号" clearable style="width: 160px" @keyup.enter="doSearch" />
      <n-select v-if="searchable" v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" />
      <n-button v-if="searchable" type="info" size="small" @click="doSearch">搜索</n-button>
      <n-button v-if="searchable" type="default" size="small" @click="resetSearch">重置</n-button>
      <n-checkbox v-if="batchActions.length" v-model:checked="allChecked" @update:checked="toggleAll">全选</n-checkbox>
      <n-button v-for="ba in batchActions" :key="ba.key + ba.label" size="small" :type="ba.buttonType || 'primary'" secondary :disabled="toolPending" @click="handleBatch(ba)">{{ ba.label }}</n-button>
      <span v-if="checkedRowKeys.length" style="font-size: 12px; color: #999">已选 {{ checkedRowKeys.length }} 项</span>
      <n-button size="small" secondary @click="showColumnModal = true">列设置</n-button>
    </n-space>

    <n-empty v-if="!columns.length" description="该层级暂无数据" style="padding: 40px 0" />
    <n-data-table
      v-else
      :bordered="false"
      :columns="tableColumns"
      :data="tableData"
      :loading="loading"
      :pagination="pagination"
      :row-key="rowKey"
      :checked-row-keys="checkedRowKeys"
      @update:checked-row-keys="onCheckedChange"
      @update:page="onPageChange"
      @update:page-size="onPageSizeChange"
    />

    <n-modal v-model:show="showInput" preset="dialog" :title="inputAction?.inputLabel || '请输入'" positive-text="确定" negative-text="取消" :positive-button-props="{ loading: toolPending }" @positive-click="confirmInput" @negative-click="showInput = false">
      <n-input v-if="inputAction?.inputType === 'number'" v-model:value="inputValue" type="number" :placeholder="inputAction?.inputLabel" />
      <n-input v-else v-model:value="inputValue" :placeholder="inputAction?.inputLabel" />
    </n-modal>

    <n-modal v-model:show="showColumnModal" preset="card" title="列设置" style="width: min(1100px, calc(100vw - 48px))">
      <!-- 搜索 + 已选计数 -->
      <div style="display: flex; align-items: center; gap: 8px; margin-bottom: 8px">
        <n-input v-model:value="colKeyword" size="small" clearable placeholder="搜索列名" style="flex: 1">
          <template #prefix>
            <n-icon><SearchOutline /></n-icon>
          </template>
        </n-input>
        <n-tag size="small" :bordered="false" type="info">已选 {{ selectedColumnKeys.length }} / {{ allColumnKeys.length }}</n-tag>
      </div>

      <!-- 滚动分组（可折叠，默认展开） -->
      <div style="max-height: 500px; overflow-y: auto; padding: 2px 12px 4px 2px">
        <n-collapse :default-expanded-names="allGroupNames">
          <n-collapse-item v-for="g in filteredColumnGroups" :key="g.name" :name="g.name" style="margin-bottom: 4px" :title="`${g.name} (${g.columns.length})`">
            <n-checkbox-group v-model:value="selectedColumnKeys">
              <div style="display: grid; grid-template-columns: repeat(auto-fill, minmax(190px, 1fr)); gap: 12px 24px">
                <n-checkbox v-for="col in g.columns" :key="col.key" :value="col.key" :label="col.title" />
              </div>
            </n-checkbox-group>
          </n-collapse-item>
        </n-collapse>
        <div v-if="!filteredColumnGroups.length" style="color: #999; padding: 16px 0; text-align: center">暂无匹配列</div>
      </div>

      <!-- 底部操作 -->
      <div style="margin-top: 10px; display: flex; justify-content: space-between; align-items: center">
        <div style="display: flex; gap: 8px">
          <n-button size="small" type="info" secondary @click="selectAllColumns">全选</n-button>
          <n-button size="small" type="warning" secondary @click="selectNoneColumns">全不选</n-button>
        </div>
        <div>
          <n-button size="small" type="primary" @click="applyColumns">确定</n-button>
          <n-button size="small" style="margin-left: 8px" @click="resetColumns">重置列</n-button>
        </div>
      </div>
    </n-modal>
  </n-card>
</template>

<script setup>
import { ref, computed, h, onMounted, watch } from 'vue'
import { NButton, useMessage, useDialog } from 'naive-ui'
import { SearchOutline } from '@vicons/ionicons5'
import { useTable } from '../../composables/useTable'
import { useDict } from '../../composables/useDict'
import { groupField } from './adDataFields'

const props = defineProps({
  fetchFn: { type: Function, required: true },
  columns: { type: Array, required: true },
  level: { type: String, required: true },
  storageKey: { type: String, required: true },
  searchable: { type: Boolean, default: true },
  actions: { type: Object, default: () => ({ idKey: 'id', row: [], batch: [] }) },
  toolFn: { type: Function, default: null },
})

const message = useMessage()
const dialog = useDialog()
const { options } = useDict()
const statusOptions = computed(() => options('status'))

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(props.fetchFn)

const idKey = computed(() => props.actions?.idKey || 'id')
const rowActions = computed(() => props.actions?.row || [])
const batchActions = computed(() => props.actions?.batch || [])

// 列选择
const allColumnKeys = computed(() => props.columns.map((c) => c.key))
const storageKey = `adData_${props.storageKey}_${props.level}`
const selectedColumnKeys = ref([])
const showColumnModal = ref(false)
const columnGroups = computed(() => {
  const groups = {}
  props.columns.forEach((c) => {
    const g = groupField(c.key)
    ;(groups[g] = groups[g] || []).push(c)
  })
  return Object.keys(groups).map((name) => ({ name, columns: groups[name] }))
})

// 搜索（只影响展示，不影响已选列）
const colKeyword = ref('')
const filteredColumnGroups = computed(() => {
  const kw = colKeyword.value.trim().toLowerCase()
  if (!kw) return columnGroups.value
  return columnGroups.value
    .map((g) => ({ ...g, columns: g.columns.filter((c) => (c.title || '').toLowerCase().includes(kw) || (c.key || '').toLowerCase().includes(kw)) }))
    .filter((g) => g.columns.length)
})
const allGroupNames = computed(() => filteredColumnGroups.value.map((g) => g.name))

function selectAllColumns() {
  selectedColumnKeys.value = allColumnKeys.value.slice()
}

function selectNoneColumns() {
  selectedColumnKeys.value = []
}

function initColumns() {
  let saved = null
  try { saved = JSON.parse(localStorage.getItem(storageKey)) } catch { /* */ }
  selectedColumnKeys.value = Array.isArray(saved) && saved.length ? saved : allColumnKeys.value.slice()
}
function applyColumns() {
  if (!selectedColumnKeys.value.length) {
    message.warning('至少保留一列')
    selectedColumnKeys.value = allColumnKeys.value.slice()
    return
  }
  localStorage.setItem(storageKey, JSON.stringify(selectedColumnKeys.value))
  message.success('列设置已保存')
}
function resetColumns() {
  selectedColumnKeys.value = allColumnKeys.value.slice()
  localStorage.setItem(storageKey, JSON.stringify(selectedColumnKeys.value))
  message.success('已重置列')
}

// 搜索
const searchKeyword = ref('')
const searchStatus = ref(null)
const dateRange = ref(null)

function fmtDate(ts) {
  const d = new Date(ts)
  const m = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  return `${d.getFullYear()}-${m}-${day}`
}

function buildFilters() {
  const f = { columns: selectedColumnKeys.value }
  if (searchKeyword.value) f.keyword = searchKeyword.value
  if (searchStatus.value != null) f.status = searchStatus.value
  if (Array.isArray(dateRange.value) && dateRange.value.length === 2) {
    f.start_date = fmtDate(dateRange.value[0])
    f.end_date = fmtDate(dateRange.value[1])
  }
  return f
}
function doSearch() { clearChecked(); search(buildFilters()) }
function resetSearch() {
  searchKeyword.value = ''
  searchStatus.value = null
  dateRange.value = null
  clearChecked()
  search(buildFilters())
}

watch(() => props.columns, initColumns)

// 行选择
const checkedRowKeys = ref([])
const allChecked = ref(false)
function rowKey(row) { return row[idKey.value] ?? row.id ?? '' }
function clearChecked() { checkedRowKeys.value = []; allChecked.value = false }
function onCheckedChange(keys) {
  checkedRowKeys.value = keys
  allChecked.value = tableData.value.length > 0 && keys.length === tableData.value.length
}
function toggleAll(checked) {
  checkedRowKeys.value = checked ? tableData.value.map((r) => rowKey(r)).filter((k) => k !== '') : []
}
function onPageChange(page) { clearChecked(); handlePageChange(page) }
function onPageSizeChange(size) { clearChecked(); handlePageSizeChange(size) }

// 操作
const toolPending = ref(false)
function doTool(level, action, payload) {
  if (!props.toolFn) { message.warning('未配置操作服务'); return Promise.reject(new Error('未配置操作服务')) }
  if (toolPending.value) return Promise.resolve()
  toolPending.value = true
  return props.toolFn(level, action, payload)
    .then(() => { message.success('操作成功') })
    .catch((err) => { message.error(err?.message || '操作失败'); throw err })
    .finally(() => { toolPending.value = false })
}
function handleRowAction(row, act) {
  const id = row[idKey.value]
  if (act.type === 'input') {
    pendingRow.value = row
    inputAction.value = act
    inputValue.value = ''
    showInput.value = true
    return
  }
  const run = () => doTool(props.level, act.key, { id })
  if (act.confirm) {
    dialog.warning({ title: '确认', content: `确定执行「${act.label}」吗？`, positiveText: '确定', negativeText: '取消', onPositiveClick: () => run().catch(() => {}) })
  } else {
    run().catch(() => {})
  }
}
function confirmInput() {
  if (!pendingRow.value || !inputAction.value || toolPending.value) return
  const payload = { id: pendingRow.value[idKey.value] }
  payload[inputAction.value.field] = inputAction.value.inputType === 'number' ? Number(inputValue.value) : inputValue.value
  const action = inputAction.value
  return doTool(props.level, action.key, payload)
    .then(() => { showInput.value = false })
    .finally(() => { pendingRow.value = null; inputAction.value = null })
}
function handleBatch(ba) {
  const ids = checkedRowKeys.value
  if (!ids.length) { message.warning('请先勾选记录'); return }
  const run = () => {
    const payload = { ids }
    if (ba.status !== undefined) payload.status = ba.status
    return doTool(props.level, ba.key, payload)
  }
  if (ba.confirm) {
    dialog.warning({ title: '确认', content: `确定批量执行「${ba.label}」吗？`, positiveText: '确定', negativeText: '取消', onPositiveClick: () => run().catch(() => {}) })
  } else {
    run().catch(() => {})
  }
}

const actionColumn = computed(() => {
  if (!rowActions.value.length) return null
  return {
    title: '操作', key: '__actions',
    width: Math.max(140, rowActions.value.length * 56),
    render: (row) => h('div', { style: 'display:flex;gap:4px;flex-wrap:wrap' },
      rowActions.value.map((act) => h(NButton, { size: 'tiny', secondary: true, onClick: () => handleRowAction(row, act) }, { default: () => act.label }))),
  }
})

const tableColumns = computed(() => {
  const cols = props.columns.filter((c) => selectedColumnKeys.value.includes(c.key)).map((c) => ({ ...c }))
  if (actionColumn.value) cols.push(actionColumn.value)
  return cols
})

const showInput = ref(false)
const inputAction = ref(null)
const inputValue = ref('')
const pendingRow = ref(null)

onMounted(() => { initColumns(); search(buildFilters()) })
</script>