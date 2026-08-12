<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/标识" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑游戏' : '新增游戏'" preset="card" style="width: 640px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="name" label="游戏名称">
            <n-input v-model:value="formData.name" placeholder="请输入游戏名称" />
          </n-form-item-gi>
          <n-form-item-gi path="mark" label="游戏标识">
            <n-input v-model:value="formData.mark" placeholder="请输入游戏标识" :disabled="isEdit" />
          </n-form-item-gi>
          <n-form-item-gi path="web_name" label="Web名称">
            <n-input v-model:value="formData.web_name" placeholder="请输入Web显示名称" />
          </n-form-item-gi>
          <n-form-item-gi path="icon" label="图标">
            <n-input v-model:value="formData.icon" placeholder="请输入图标地址" />
          </n-form-item-gi>
          <n-form-item-gi path="cp_id" label="所属CP">
            <n-select v-model:value="formData.cp_id" :options="cpOptions" placeholder="请选择CP" clearable />
          </n-form-item-gi>
          <n-grid-item>
            <n-form-item path="status" label="状态" label-placement="left">
              <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="启用" unchecked-text="禁用" />
            </n-form-item>
          </n-grid-item>
        </n-grid>
      </n-form>
      <template #footer>
        <n-space justify="end">
          <n-button @click="showModal = false">取消</n-button>
          <n-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script setup>
import { ref, reactive, h, onMounted, computed } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { useDict } from '../../composables/useDict'
import { useOptions } from '../../composables/useOptions'
import { getGameList, createGame, updateGame, deleteGame } from '../../api/game'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getGameList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()

const searchKeyword = ref('')
const searchStatus = ref(null)
const cpOptions = ref([])
const { loadOptions } = useOptions()
const { load: loadDict, options } = useDict()
const statusOptions = computed(() => options('status'))

const formData = reactive({ name: '', mark: '', web_name: '', icon: '', cp_id: null, status: 1 })
function resetForm() { Object.assign(formData, { name: '', mark: '', web_name: '', icon: '', cp_id: null, status: 1 }) }
const rules = { name: [{ required: true, message: '请输入游戏名称', trigger: 'blur' }], mark: [{ required: true, message: '请输入游戏标识', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '游戏名称', key: 'name' },
  { title: '标识', key: 'mark' },
  { title: 'Web名称', key: 'web_name' },
  { title: 'CP', key: 'cp_id', width: 80, render: (row) => { const c = cpOptions.value.find(o => o.value === row.cp_id); return c ? c.label : '' } },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.mark = row.mark; formData.web_name = row.web_name; formData.icon = row.icon; formData.cp_id = row.cp_id; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createGame, updateGame)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteGame)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }

async function handleStatusChange(row, val) {
  try {
    await updateGame(row.id, { ...row, status: val ? 1 : 0 })
    row.status = val ? 1 : 0
    message.success('状态已更新')
  } catch {
    message.error('更新失败')
  }
}

onMounted(async () => { await loadDict(); cpOptions.value = await loadOptions('game_cp'); search({}) })
</script>
