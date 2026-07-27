<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/备注" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchMediaId" :options="mediaOptions" placeholder="媒体渠道" clearable style="width: 150px" @update:value="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑应用' : '新增应用'" preset="card" style="width: 520px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="media_id" label="媒体渠道">
          <n-select v-model:value="formData.media_id" :options="mediaOptions" placeholder="请选择媒体渠道" clearable />
        </n-form-item>
        <n-form-item path="name" label="应用名称">
          <n-input v-model:value="formData.name" placeholder="请输入应用名称" />
        </n-form-item>
        <n-form-item path="app_id" label="App ID">
          <n-input-number v-model:value="formData.app_id" placeholder="请输入App ID" style="width: 100%" />
        </n-form-item>
        <n-form-item path="app_secret" label="App Secret">
          <n-input-number v-model:value="formData.app_secret" placeholder="请输入App Secret" style="width: 100%" />
        </n-form-item>
        <n-form-item path="remark" label="备注">
          <n-input v-model:value="formData.remark" placeholder="请输入备注" />
        </n-form-item>
        <n-form-item path="extra" label="扩展信息">
          <n-input v-model:value="formData.extra" type="textarea" placeholder="请输入扩展信息" />
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
import { ref, reactive, h, onMounted, computed } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm, NInputNumber, useMessage } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { getMediaApplicationList, createMediaApplication, updateMediaApplication, deleteMediaApplication } from '../../api/mkt'
import { formatTime } from '../../utils/format'
import { useDict } from '../../composables/useDict'
import { useOptions } from '../../composables/useOptions'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getMediaApplicationList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()
const searchKeyword = ref('')
const searchMediaId = ref(null)
const searchStatus = ref(null)
const mediaOptions = ref([])
const { loadOptions } = useOptions()
const statusOptions = computed(() => options('status'))
const formData = reactive({ media_id: null, name: '', app_id: null, app_secret: null, remark: '', extra: '', status: 1 })
function resetForm() { Object.assign(formData, { media_id: null, name: '', app_id: null, app_secret: null, remark: '', extra: '', status: 1 }) }
const rules = {
  name: [{ required: true, message: '请输入应用名称', trigger: 'blur' }],
  media_id: [{ required: true, type: 'number', message: '请选择媒体渠道', trigger: 'change' }],
  app_id: [{ required: true, type: 'number', message: '请输入App ID', trigger: 'blur' }],
}
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '媒体渠道', key: 'media_id', width: 120, render: (row) => { const m = mediaOptions.value.find(o => o.value === row.media_id); return m ? m.label : '' } },
  { title: '应用名称', key: 'name' },
  { title: 'App ID', key: 'app_id', width: 80 },
  { title: '备注', key: 'remark', width: 100, ellipsis: { tooltip: true } },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, media_id: searchMediaId.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.media_id = row.media_id; formData.name = row.name; formData.app_id = row.app_id; formData.app_secret = row.app_secret; formData.remark = row.remark; formData.extra = row.extra; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createMediaApplication, updateMediaApplication)) search({ keyword: searchKeyword.value, media_id: searchMediaId.value ?? 0, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteMediaApplication)) search({ keyword: searchKeyword.value, media_id: searchMediaId.value ?? 0, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateMediaApplication(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
onMounted(async () => { await loadDict(); mediaOptions.value = await loadOptions('media'); search({}) })
</script>
