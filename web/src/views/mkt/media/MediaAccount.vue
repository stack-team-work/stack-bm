<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/账号/UID" clearable style="width: 220px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchMediaSubId" :options="mediaSubOptions" placeholder="子渠道" clearable style="width: 150px" @update:value="doSearch" />
          <n-select v-model:value="searchSubjectId" :options="subjectOptions" placeholder="主体" clearable style="width: 150px" @update:value="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑渠道账户' : '新增渠道账户'" preset="card" style="width: 680px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="110">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="name" label="账户别名">
            <n-input v-model:value="formData.name" placeholder="请输入账户别名" />
          </n-form-item-gi>
          <n-form-item-gi path="uid" label="平台UID">
            <n-input v-model:value="formData.uid" placeholder="请输入平台UID" :disabled="isEdit" />
          </n-form-item-gi>
          <n-form-item-gi path="username" label="媒体账号">
            <n-input v-model:value="formData.username" placeholder="请输入媒体渠道账号" />
          </n-form-item-gi>
          <n-form-item-gi path="media_sub_id" label="子渠道">
            <n-select v-model:value="formData.media_sub_id" :options="mediaSubOptions" placeholder="请选择子渠道" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="agent_id" label="代理">
            <n-select v-model:value="formData.agent_id" :options="agentOptions" placeholder="请选择代理" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="subject_id" label="主体">
            <n-select v-model:value="formData.subject_id" :options="subjectOptions" placeholder="请选择主体" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="media_manager_manager_id" label="关联管家">
            <n-select v-model:value="formData.media_manager_manager_id" :options="managerOptions" placeholder="请选择管家" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="rebate" label="返点">
            <n-input-number v-model:value="formData.rebate" placeholder="返点" :min="0" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="balance" label="余额">
            <n-input-number v-model:value="formData.balance" placeholder="余额" :min="0" style="width: 100%" />
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
import { NSwitch, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { getMediaAccountList, createMediaAccount, updateMediaAccount, deleteMediaAccount } from '../../../api/mkt/media'
import { formatTime } from '../../../utils/format'
import { useDict } from '../../../composables/useDict'
import { useOptions } from '../../../composables/useOptions'
import TableActions from '../../../components/TableActions.vue'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getMediaAccountList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()
const searchKeyword = ref('')
const searchMediaSubId = ref(null)
const searchSubjectId = ref(null)
const searchStatus = ref(null)
const mediaSubOptions = ref([])
const agentOptions = ref([])
const subjectOptions = ref([])
const managerOptions = ref([])
const { loadOptions } = useOptions()
const statusOptions = computed(() => options('status'))
const useTypeOptions = computed(() => options('status'))
const formData = reactive({ name: '', uid: '', username: '', media_sub_id: null, agent_id: null, subject_id: null, media_manager_manager_id: 0, rebate: 0, balance: 0, status: 1 })
function resetForm() { Object.assign(formData, { name: '', uid: '', username: '', media_sub_id: null, agent_id: null, subject_id: null, media_manager_manager_id: 0, rebate: 0, balance: 0, status: 1 }) }
const rules = {
  name: [{ required: true, message: '请输入账户别名', trigger: 'blur' }],
  uid: [{ required: true, message: '请输入平台UID', trigger: 'blur' }],
}
const columns = [
  { title: 'ID', key: 'id', width: 70 },
  { title: '账户别名', key: 'name', width: 130 },
  { title: '媒体账号', key: 'username', width: 130 },
  { title: '平台UID', key: 'uid', width: 140 },
  { title: '子渠道', key: 'media_sub_id', width: 100, render: (row) => { const o = mediaSubOptions.value.find(x => x.value === row.media_sub_id); return o ? o.label : '-' } },
  { title: '主体', key: 'subject_id', width: 100, render: (row) => { const o = subjectOptions.value.find(x => x.value === row.subject_id); return o ? o.label : '-' } },
  { title: '余额', key: 'balance', width: 100, render: (row) => row.balance ?? 0 },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(TableActions, { row, edit: () => handleEdit(row), remove: () => onDelete(row.id) }) },
]
function doSearch() { search({ keyword: searchKeyword.value, media_sub_id: searchMediaSubId.value ?? 0, subject_id: searchSubjectId.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); Object.assign(formData, { name: row.name, uid: row.uid, username: row.username, media_sub_id: row.media_sub_id, agent_id: row.agent_id, subject_id: row.subject_id, media_manager_manager_id: row.media_manager_manager_id, rebate: row.rebate, balance: row.balance, status: row.status }); openEdit(row) }
async function handleSubmit() { if (await submit(formData, createMediaAccount, updateMediaAccount)) doSearch() }
async function onDelete(id) { if (await doDelete(id, deleteMediaAccount)) doSearch() }
async function handleStatusChange(row, val) {
  try { await updateMediaAccount(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
async function loadData() {
  mediaSubOptions.value = await loadOptions('media_sub')
  agentOptions.value = await loadOptions('media_agent')
  subjectOptions.value = await loadOptions('media_subject')
  managerOptions.value = await loadOptions('media_manager')
}
onMounted(async () => { await loadDict(); loadData(); search({}) })
</script>