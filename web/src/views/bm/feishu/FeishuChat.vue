<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索ChatID/CallAction" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchFeishuAppId" :options="feishuAppOptions" placeholder="飞书应用" clearable style="width: 160px" @update:value="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑飞书聊天' : '新增飞书聊天'" preset="card" style="width: 720px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="type" label="机器人类型">
            <n-select v-model:value="formData.type" :options="chatTypeOptions" placeholder="请选择机器人类型" />
          </n-form-item-gi>
          <n-form-item-gi path="chat_id" label="群聊天ID">
            <n-input v-model:value="formData.chat_id" placeholder="请输入群聊天ID" />
          </n-form-item-gi>
          <n-form-item-gi path="feishu_app_id" label="关联飞书应用">
            <n-select v-model:value="formData.feishu_app_id" :options="feishuAppOptions" placeholder="请选择飞书应用" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="call_action" label="Call Action">
            <n-input v-model:value="formData.call_action" placeholder="请输入Call Action" />
          </n-form-item-gi>
          <n-form-item-gi path="action_title" label="对话标题">
            <n-input v-model:value="formData.action_title" placeholder="请输入对话标题" />
          </n-form-item-gi>
          <n-form-item-gi path="at_type" label="艾特方式">
            <n-select v-model:value="formData.at_type" :options="atTypeOptions" placeholder="请选择艾特方式" />
          </n-form-item-gi>
          <n-form-item-gi path="default_at_list" label="默认艾特列表" :span="2">
            <n-select v-model:value="formData.default_at_list" :options="atUserOptions" multiple placeholder="请选择默认艾特用户" />
          </n-form-item-gi>
          <n-form-item-gi path="at_list" label="选择艾特列表" :span="2">
            <n-select v-model:value="formData.at_list" :options="atUserOptions" multiple placeholder="请选择艾特用户" />
          </n-form-item-gi>
          <n-grid-item :span="2">
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
import { ref, reactive, h, computed, onMounted } from 'vue'
import { NButton, NSpace, NSwitch, NIcon, useMessage } from 'naive-ui'
import { CreateOutline } from '@vicons/ionicons5'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { useOptions } from '../../../composables/useOptions'
import { getFeishuChatList, createFeishuChat, updateFeishuChat, updateFeishuChatStatus, getFeishuUserAll } from '../../../api/bm/feishu'
import { formatTime } from '../../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getFeishuChatList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit } = useModal()
const { load: loadDict, options } = useDict()
const { loadOptions } = useOptions()
const message = useMessage()
const searchKeyword = ref('')
const searchFeishuAppId = ref(null)
const searchStatus = ref(null)
const feishuAppOptions = ref([])
const atUserOptions = ref([])
const feishuUserIdMap = ref({})
const statusOptions = computed(() => options('status'))
const chatTypeOptions = computed(() => options('feishu_chat_type'))
const atTypeOptions = computed(() => options('feishu_chat_at_type'))
const formData = reactive({ type: 1, chat_id: '', feishu_app_id: null, call_action: '', action_title: '', at_type: 1, default_at_list: [], at_list: [], status: 1 })
function resetForm() { Object.assign(formData, { type: 1, chat_id: '', feishu_app_id: null, call_action: '', action_title: '', at_type: 1, default_at_list: [], at_list: [], status: 1 }) }
const rules = {
  chat_id: [{ required: true, message: '请输入群聊天ID', trigger: 'blur' }],
  call_action: [{ required: true, message: '请输入Call Action', trigger: 'blur' }],
  default_at_list: [{ required: true, type: 'array', min: 1, message: '请至少选择一个默认艾特用户', trigger: 'change' }],
}
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: 'Chat ID', key: 'chat_id', width: 140 },
  { title: '类型', key: 'type', width: 90, render: (row) => { const t = chatTypeOptions.value.find(o => o.value === row.type); return t ? t.label : row.type } },
  { title: '应用', key: 'feishu_app_id', width: 110, render: (row) => { const a = feishuAppOptions.value.find(o => o.value === row.feishu_app_id); return a ? a.label : '' } },
  { title: 'Call Action', key: 'call_action', width: 120 },
  { title: '标题', key: 'action_title', width: 100, ellipsis: { tooltip: true } },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 80, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(CreateOutline) }), ' 编辑'] }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, feishu_app_id: searchFeishuAppId.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) {
  resetForm()
  formData.type = row.type; formData.chat_id = row.chat_id; formData.feishu_app_id = row.feishu_app_id
  formData.call_action = row.call_action; formData.action_title = row.action_title; formData.at_type = row.at_type
  formData.status = row.status
  try { formData.default_at_list = Object.keys(JSON.parse(row.default_at_list || '{}')).map(Number) } catch { formData.default_at_list = [] }
  try { formData.at_list = Object.keys(JSON.parse(row.at_list || '{}')).map(Number) } catch { formData.at_list = [] }
  openEdit(row)
}
async function handleSubmit() {
  const data = { ...formData }
  data.default_at_list = JSON.stringify(Object.fromEntries(
    (formData.default_at_list || []).map(id => [id, feishuUserIdMap.value[id] || ''])
  ))
  data.at_list = JSON.stringify(Object.fromEntries(
    (formData.at_list || []).map(id => [id, feishuUserIdMap.value[id] || ''])
  ))
  if (await submit(data, createFeishuChat, updateFeishuChat)) search({ keyword: searchKeyword.value, feishu_app_id: searchFeishuAppId.value ?? 0, status: searchStatus.value ?? -1 })
}
async function handleStatusChange(row, val) {
  try { await updateFeishuChatStatus(row.id, { status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
async function loadAtUsers() {
  try {
    const res = await getFeishuUserAll()
    const users = res.data || []
    feishuUserIdMap.value = {}
    atUserOptions.value = []
    users.filter(u => u.status === 1).forEach(u => {
      feishuUserIdMap.value[u.admin_id] = u.feishu_user_id
      atUserOptions.value.push({ label: u.admin_name || `用户${u.admin_id}`, value: u.admin_id })
    })
  } catch { /* */ }
}
onMounted(async () => { await loadDict(); feishuAppOptions.value = await loadOptions('feishu_app'); await loadAtUsers(); search({}) })
</script>
