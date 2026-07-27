<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索应用名称/ID" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑飞书应用' : '新增飞书应用'" preset="card" style="width: 520px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="app_id" label="App ID">
          <n-input v-model:value="formData.app_id" placeholder="请输入App ID" />
        </n-form-item>
        <n-form-item path="app_secret" label="App Secret">
          <n-input v-model:value="formData.app_secret" placeholder="请输入App Secret" />
        </n-form-item>
        <n-form-item path="app_name" label="应用名称">
          <n-input v-model:value="formData.app_name" placeholder="请输入应用名称" />
        </n-form-item>
        <n-form-item path="mark" label="应用标识">
          <n-input v-model:value="formData.mark" placeholder="请输入应用标识（留空自动生成）" :disabled="isEdit" />
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
import { ref, reactive, h, computed, onMounted } from 'vue'
import { NButton, NSpace, NSwitch, useMessage } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { useDict } from '../../composables/useDict'
import { getFeishuAppList, createFeishuApp, updateFeishuApp, updateFeishuAppStatus } from '../../api/system'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getFeishuAppList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()
const searchKeyword = ref('')
const searchStatus = ref(null)
const statusOptions = computed(() => options('status'))
const formData = reactive({ app_id: '', app_secret: '', app_name: '', mark: '', status: 1 })
function resetForm() { Object.assign(formData, { app_id: '', app_secret: '', app_name: '', mark: '', status: 1 }) }
const rules = { app_id: [{ required: true, message: '请输入App ID', trigger: 'blur' }], app_secret: [{ required: true, message: '请输入App Secret', trigger: 'blur' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: 'App ID', key: 'app_id' },
  { title: '应用名称', key: 'app_name' },
  { title: '标识', key: 'mark' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 80, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.app_id = row.app_id; formData.app_secret = row.app_secret; formData.app_name = row.app_name; formData.mark = row.mark; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createFeishuApp, updateFeishuApp)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateFeishuAppStatus(row.id, { status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
onMounted(async () => { await loadDict(); search({}) })
</script>
