<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/标识" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchType" :options="typeOptions" placeholder="标签类型" clearable style="width: 130px" @update:value="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 100px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑标签' : '新增标签'" preset="card" style="width: 500px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="name" label="标签名称">
          <n-input v-model:value="formData.name" placeholder="请输入标签名称" />
        </n-form-item>
        <n-form-item path="mark" label="标签标识">
          <n-input v-model:value="formData.mark" placeholder="请输入标签标识" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="type" label="标签类型">
          <n-select v-model:value="formData.type" :options="typeOptions" placeholder="请选择标签类型" :disabled="isEdit" />
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
import { NButton, NSpace, NSwitch, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { getGameTagList, createGameTag, updateGameTag, deleteGameTag } from '../../api/game'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getGameTagList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()

const searchKeyword = ref('')
const searchType = ref(null)
const searchStatus = ref(null)
const typeOptions = [{ label: '游戏风格(style_id)', value: 1 }, { label: '游戏类型(type_id)', value: 2 }]
const statusOptions = [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }]

const formData = reactive({ name: '', mark: '', type: 1, status: 1 })
function resetForm() { Object.assign(formData, { name: '', mark: '', type: 1, status: 1 }) }
const rules = { name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }], mark: [{ required: true, message: '请输入标签标识', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '标签名称', key: 'name' },
  { title: '标识', key: 'mark' },
  { title: '类型', key: 'type', width: 150, render: (row) => row.type === 1 ? '游戏风格(style_id)' : '游戏类型(type_id)' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value, type: searchType.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.mark = row.mark; formData.type = row.type; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createGameTag, updateGameTag)) search({ keyword: searchKeyword.value, type: searchType.value ?? 0, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteGameTag)) search({ keyword: searchKeyword.value, type: searchType.value ?? 0, status: searchStatus.value ?? -1 }) }

async function handleStatusChange(row, val) {
  try {
    await updateGameTag(row.id, { ...row, status: val ? 1 : 0 })
    row.status = val ? 1 : 0
    message.success('状态已更新')
  } catch {
    message.error('更新失败')
  }
}

onMounted(() => search({}))
</script>
