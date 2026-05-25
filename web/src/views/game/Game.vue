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

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑游戏' : '新增游戏'" preset="card" style="width: 500px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="name" label="游戏名称">
          <n-input v-model:value="formData.name" placeholder="请输入游戏名称" />
        </n-form-item>
        <n-form-item path="mark" label="游戏标识">
          <n-input v-model:value="formData.mark" placeholder="请输入游戏标识" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="web_name" label="Web名称">
          <n-input v-model:value="formData.web_name" placeholder="请输入Web显示名称" />
        </n-form-item>
        <n-form-item path="icon" label="图标">
          <n-input v-model:value="formData.icon" placeholder="请输入图标地址" />
        </n-form-item>
        <n-form-item path="cp_id" label="所属CP">
          <n-select v-model:value="formData.cp_id" :options="cpOptions" placeholder="请选择CP" clearable />
        </n-form-item>
        <n-form-item path="status" label="状态">
          <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="启用" unchecked-text="禁用" />
        </n-form-item>
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
import { ref, reactive, h, onMounted } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { getGameList, createGame, updateGame, deleteGame, getGameCpAll } from '../../api/game'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getGameList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()

const searchKeyword = ref('')
const searchStatus = ref(null)
const cpOptions = ref([])
const statusOptions = [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }]

const formData = reactive({ name: '', mark: '', web_name: '', icon: '', cp_id: null, status: 1 })
function resetForm() { Object.assign(formData, { name: '', mark: '', web_name: '', icon: '', cp_id: null, status: 1 }) }
const rules = { name: [{ required: true, message: '请输入游戏名称', trigger: 'blur' }], mark: [{ required: true, message: '请输入游戏标识', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '游戏名称', key: 'name' },
  { title: '标识', key: 'mark' },
  { title: 'Web名称', key: 'web_name' },
  { title: 'CP', key: 'cp_id', width: 80, render: (row) => { const c = cpOptions.value.find(o => o.value === row.cp_id); return c ? c.label : '' } },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, readonly: true, size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170 },
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

async function loadCps() { try { const res = await getGameCpAll(); cpOptions.value = (res.data || []).map(c => ({ label: c.name, value: c.id })) } catch { /* */ } }
onMounted(() => { loadCps(); search({}) })
</script>
