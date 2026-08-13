<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/账号" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchMediaId" :options="mediaOptions" placeholder="媒体渠道" clearable style="width: 150px" @update:value="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑管家' : '新增管家'" preset="card" style="width: 680px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="media_id" label="媒体渠道">
            <n-select v-model:value="formData.media_id" :options="mediaOptions" placeholder="请选择媒体渠道" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="application_id" label="所属应用">
            <n-select v-model:value="formData.application_id" :options="appOptions" placeholder="请选择应用" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="name" label="管家名称">
            <n-input v-model:value="formData.name" placeholder="请输入管家名称" />
          </n-form-item-gi>
          <n-form-item-gi path="account" label="管家账号">
            <n-input v-model:value="formData.account" placeholder="请输入管家账号" />
          </n-form-item-gi>
          <n-form-item-gi path="account_id" label="管家ID">
            <n-input v-model:value="formData.account_id" placeholder="请输入管家ID" />
          </n-form-item-gi>
          <n-form-item-gi path="account_num" label="绑定账户数">
            <n-input-number v-model:value="formData.account_num" placeholder="绑定账户数" :min="0" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="auth_status" label="授权状态">
            <n-select v-model:value="formData.auth_status" :options="authOptions" placeholder="授权状态" />
          </n-form-item-gi>
          <n-grid-item>
            <n-form-item path="status" label="状态" label-placement="left">
              <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="启用" unchecked-text="禁用" />
            </n-form-item>
          </n-grid-item>
          <n-form-item-gi path="remark" label="备注" :span="2">
            <n-input v-model:value="formData.remark" placeholder="请输入备注" />
          </n-form-item-gi>
          <n-form-item-gi path="extra" label="扩展信息" :span="2">
            <n-input v-model:value="formData.extra" type="textarea" :rows="2" placeholder="请输入扩展信息" />
          </n-form-item-gi>
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
import { NSwitch, NInputNumber, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { getMediaManagerList, createMediaManager, updateMediaManager, deleteMediaManager } from '../../../api/mkt/media'
import { formatTime } from '../../../utils/format'
import { useDict } from '../../../composables/useDict'
import { useOptions } from '../../../composables/useOptions'
import TableActions from '../../../components/TableActions.vue'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getMediaManagerList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()
const searchKeyword = ref('')
const searchMediaId = ref(null)
const searchStatus = ref(null)
const mediaOptions = ref([])
const appOptions = ref([])
const { loadOptions } = useOptions()
const statusOptions = computed(() => options('status'))
const authOptions = computed(() => options('media_manager_auth_status'))
const formData = reactive({ media_id: null, application_id: null, name: '', account: '', account_id: '', account_num: 0, auth_status: 0, remark: '', extra: '', status: 1 })
function resetForm() { Object.assign(formData, { media_id: null, application_id: null, name: '', account: '', account_id: '', account_num: 0, auth_status: 0, remark: '', extra: '', status: 1 }) }
const rules = {
  name: [{ required: true, message: '请输入管家名称', trigger: 'blur' }],
  media_id: [{ required: true, type: 'number', message: '请选择媒体渠道', trigger: 'change' }],
  application_id: [{ required: true, type: 'number', message: '请选择应用', trigger: 'change' }],
  account: [{ required: true, message: '请输入管家账号', trigger: 'blur' }],
  account_id: [{ required: true, message: '请输入管家ID', trigger: 'blur' }],
}
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '媒体渠道', key: 'media_id', width: 120, render: (row) => { const m = mediaOptions.value.find(o => o.value === row.media_id); return m ? m.label : '' } },
  { title: '管家名称', key: 'name', width: 120 },
  { title: '管家账号', key: 'account' },
  { title: '账号数', key: 'account_num', width: 70 },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(TableActions, { row, edit: () => handleEdit(row), remove: () => onDelete(row.id) }) },
]
function doSearch() { search({ keyword: searchKeyword.value, media_id: searchMediaId.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.media_id = row.media_id; formData.application_id = row.application_id; formData.name = row.name; formData.account = row.account; formData.account_id = row.account_id; formData.account_num = row.account_num; formData.auth_status = row.auth_status; formData.remark = row.remark; formData.extra = row.extra; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createMediaManager, updateMediaManager)) search({ keyword: searchKeyword.value, media_id: searchMediaId.value ?? 0, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteMediaManager)) search({ keyword: searchKeyword.value, media_id: searchMediaId.value ?? 0, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateMediaManager(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
async function loadData() {
  mediaOptions.value = await loadOptions('media')
  appOptions.value = await loadOptions('media_application')
}
onMounted(async () => { await loadDict(); loadData(); search({}) })
</script>
