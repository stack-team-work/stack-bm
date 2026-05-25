<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/Key" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 100px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑变量' : '新增变量'" preset="card" style="width: 600px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="name" label="变量名称">
          <n-input v-model:value="formData.name" placeholder="请输入变量名称" />
        </n-form-item>
        <n-form-item path="key" label="变量Key">
          <n-input v-model:value="formData.key" placeholder="请输入变量Key" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="value" label="变量值">
          <n-input v-model:value="formData.value" type="textarea" :rows="4" placeholder="请输入变量值，多行输入自动保存为JSON数组" />
        </n-form-item>
        <n-form-item path="mark" label="备注">
          <n-input v-model:value="formData.mark" placeholder="请输入备注" />
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
import { getGameVariableList, createGameVariable, updateGameVariable, deleteGameVariable } from '../../api/game'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getGameVariableList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()

const searchKeyword = ref('')
const searchStatus = ref(null)
const statusOptions = [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }]

const formData = reactive({ name: '', key: '', value: '', mark: '', status: 1 })
function resetForm() { Object.assign(formData, { name: '', key: '', value: '', mark: '', status: 1 }) }
const rules = { name: [{ required: true, message: '请输入变量名称', trigger: 'blur' }], key: [{ required: true, message: '请输入变量Key', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '变量名称', key: 'name', width: 120 },
  { title: 'Key', key: 'key', width: 150 },
  { title: '变量值', key: 'value', ellipsis: { tooltip: true }, render: (row) => row.value ? (row.value.length > 50 ? row.value.substring(0, 50) + '...' : row.value) : '' },
  { title: '备注', key: 'mark', width: 120, ellipsis: { tooltip: true } },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, readonly: true, size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170 },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.key = row.key; formData.value = row.value; formData.mark = row.mark || ''; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createGameVariable, updateGameVariable)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteGameVariable)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }

onMounted(() => search({}))
</script>
