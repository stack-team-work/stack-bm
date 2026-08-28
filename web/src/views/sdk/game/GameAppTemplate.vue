<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索模板名称" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑SDK模板' : '新增SDK模板'" preset="card" style="width: 680px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="name" label="模板名称">
            <n-input v-model:value="formData.name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="allow_age" label="可玩年龄">
            <n-input-number v-model:value="formData.allow_age" :min="0" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="privacy_url" label="隐私条款">
            <n-input v-model:value="formData.privacy_url" placeholder="请输入隐私条款URL" />
          </n-form-item-gi>
          <n-form-item-gi path="agreement_url" label="用户协议">
            <n-input v-model:value="formData.agreement_url" placeholder="请输入用户协议URL" />
          </n-form-item-gi>
          <n-grid-item>
            <n-form-item path="is_open_realname" label="实名认证" label-placement="left">
              <n-switch v-model:value="formData.is_open_realname" :checked-value="1" :unchecked-value="0" />
            </n-form-item>
          </n-grid-item>
          <n-grid-item>
            <n-form-item path="is_open_register" label="开启注册" label-placement="left">
              <n-switch v-model:value="formData.is_open_register" :checked-value="1" :unchecked-value="0" />
            </n-form-item>
          </n-grid-item>
          <n-grid-item>
            <n-form-item path="is_open_charge" label="开启充值" label-placement="left">
              <n-switch v-model:value="formData.is_open_charge" :checked-value="1" :unchecked-value="0" />
            </n-form-item>
          </n-grid-item>
          <n-grid-item>
            <n-form-item path="is_open_float" label="悬浮窗" label-placement="left">
              <n-switch v-model:value="formData.is_open_float" :checked-value="1" :unchecked-value="0" />
            </n-form-item>
          </n-grid-item>
          <n-grid-item>
            <n-form-item path="is_alert_email" label="绑定邮箱提醒" label-placement="left">
              <n-switch v-model:value="formData.is_alert_email" :checked-value="1" :unchecked-value="0" />
            </n-form-item>
          </n-grid-item>
          <n-grid-item>
            <n-form-item path="is_alert_phone" label="绑定手机提醒" label-placement="left">
              <n-switch v-model:value="formData.is_alert_phone" :checked-value="1" :unchecked-value="0" />
            </n-form-item>
          </n-grid-item>
          <n-grid-item>
            <n-form-item path="is_alert_auth" label="自定义授权" label-placement="left">
              <n-switch v-model:value="formData.is_alert_auth" :checked-value="1" :unchecked-value="0" />
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
import { NButton, NSpace, NSwitch, useMessage } from 'naive-ui'
import TableActions from '../../../components/TableActions.vue'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { getGameAppTemplateList, createGameAppTemplate, updateGameAppTemplate, deleteGameAppTemplate } from '../../../api/sdk/game'
import { formatTime } from '../../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getGameAppTemplateList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()

const searchKeyword = ref('')
const searchStatus = ref(null)
const { load: loadDict, options } = useDict()
const statusOptions = computed(() => options('status'))

const formData = reactive({
  name: '', privacy_url: '', agreement_url: '', is_open_realname: 1, is_open_register: 1, is_open_charge: 1, is_alert_email: 1,
  is_alert_phone: 1, is_alert_auth: 1, is_open_float: 1, allow_age: 18, status: 1,
})
function resetForm() { Object.assign(formData, {
  name: '', privacy_url: '', agreement_url: '', is_open_realname: 1, is_open_register: 1, is_open_charge: 1, is_alert_email: 1,
  is_alert_phone: 1, is_alert_auth: 1, is_open_float: 1, allow_age: 18, status: 1,
}) }
const rules = { name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }], allow_age: [{ required: true, type: 'number', min: 0, message: '请输入可玩年龄', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '模板名称', key: 'name' },
  { title: '实名认证', key: 'is_open_realname', width: 100, render: (row) => row.is_open_realname === 1 ? '开启' : '关闭' },
  { title: '注册', key: 'is_open_register', width: 80, render: (row) => row.is_open_register === 1 ? '开启' : '关闭' },
  { title: '充值', key: 'is_open_charge', width: 80, render: (row) => row.is_open_charge === 1 ? '开启' : '关闭' },
  { title: '悬浮窗', key: 'is_open_float', width: 80, render: (row) => row.is_open_float === 1 ? '开启' : '关闭' },
  { title: '可玩年龄', key: 'allow_age', width: 80 },
  { title: '隐私条款', key: 'privacy_url', width: 160, ellipsis: { tooltip: true }, render: (row) => row.privacy_url || '-' },
  { title: '用户协议', key: 'agreement_url', width: 160, ellipsis: { tooltip: true }, render: (row) => row.agreement_url || '-' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(TableActions, { row, edit: () => handleEdit(row), remove: () => onDelete(row.id) }) },
]

function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) {
  resetForm()
  formData.name = row.name
  formData.privacy_url = row.privacy_url || ''
  formData.agreement_url = row.agreement_url || ''
  formData.is_open_realname = row.is_open_realname; formData.is_open_register = row.is_open_register
  formData.is_open_charge = row.is_open_charge; formData.is_alert_email = row.is_alert_email
  formData.is_alert_phone = row.is_alert_phone; formData.is_alert_auth = row.is_alert_auth
  formData.is_open_float = row.is_open_float; formData.allow_age = row.allow_age; formData.status = row.status
  openEdit(row)
}
async function handleSubmit() { if (await submit(formData, createGameAppTemplate, updateGameAppTemplate)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteGameAppTemplate)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }

async function handleStatusChange(row, val) {
  try {
    await updateGameAppTemplate(row.id, { ...row, status: val ? 1 : 0 })
    row.status = val ? 1 : 0
    message.success('状态已更新')
  } catch {
    message.error('更新失败')
  }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
