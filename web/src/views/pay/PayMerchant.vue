<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/标识" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchType" :options="typeOptions" placeholder="支付类型" clearable style="width: 120px" @update:value="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑商户' : '新增商户'" preset="card" style="width: 560px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="name" label="对内名称">
          <n-input v-model:value="formData.name" placeholder="请输入对内名称" />
        </n-form-item>
        <n-form-item path="show_name" label="对外名称">
          <n-input v-model:value="formData.show_name" placeholder="请输入对外名称" />
        </n-form-item>
        <n-form-item path="type" label="支付类型">
          <n-select v-model:value="formData.type" :options="typeOptions" placeholder="请选择支付类型" />
        </n-form-item>
        <n-form-item path="platform_mark" label="支付平台">
          <n-select v-model:value="formData.platform_mark" :options="platformOptions" placeholder="请选择支付平台" />
        </n-form-item>
        <n-form-item path="mark" label="标识">
          <n-input v-model:value="formData.mark" placeholder="请输入标识（留空自动生成）" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="url" label="支付URL">
          <n-input v-model:value="formData.url" placeholder="请输入支付URL" />
        </n-form-item>
        <n-form-item path="rate" label="费率">
          <n-input-number v-model:value="formData.rate" :min="0" :step="0.0001" style="width: 100%" placeholder="请输入费率" />
        </n-form-item>
        <n-form-item path="weight" label="权重">
          <n-input-number v-model:value="formData.weight" :min="0" style="width: 100%" placeholder="请输入权重" />
        </n-form-item>
        <n-form-item path="config" label="支付配置">
          <n-input v-model:value="formData.config" type="textarea" placeholder="请输入支付参数配置" />
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
import { getPayMerchantList, createPayMerchant, updatePayMerchant, deletePayMerchant, getPayPlatformAll } from '../../api/pay'
import { formatTime } from '../../utils/format'
import { useDict } from '../../composables/useDict'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getPayMerchantList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()
const searchKeyword = ref('')
const searchType = ref(null)
const searchStatus = ref(null)
const platformOptions = ref([])
const typeOptions = computed(() => options('pay_merchant_type'))
const statusOptions = computed(() => options('status'))
const formData = reactive({ name: '', show_name: '', type: 1, platform_mark: 0, mark: '', url: '', rate: 0, weight: 0, config: '', status: 1 })
function resetForm() { Object.assign(formData, { name: '', show_name: '', type: 1, platform_mark: 0, mark: '', url: '', rate: 0, weight: 0, config: '', status: 1 }) }
const rules = {
  name: [{ required: true, message: '请输入对内名称', trigger: 'blur' }],
  show_name: [{ required: true, message: '请输入对外名称', trigger: 'blur' }],
  type: [{ required: true, type: 'number', message: '请选择支付类型', trigger: 'change' }],
}
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '对内名称', key: 'name' },
  { title: '对外名称', key: 'show_name' },
  { title: '类型', key: 'type', width: 80, render: (row) => { const t = typeOptions.value.find(o => o.value === row.type); return t ? t.label : row.type } },
  { title: '平台', key: 'platform_mark', width: 100, render: (row) => { const p = platformOptions.value.find(o => o.value === row.platform_mark); return p ? p.label : '' } },
  { title: '费率', key: 'rate', width: 70 },
  { title: '权重', key: 'weight', width: 60 },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, type: searchType.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.show_name = row.show_name; formData.type = row.type; formData.platform_mark = row.platform_mark; formData.mark = row.mark; formData.url = row.url; formData.rate = row.rate; formData.weight = row.weight; formData.config = row.config; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createPayMerchant, updatePayMerchant)) search({ keyword: searchKeyword.value, type: searchType.value ?? 0, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deletePayMerchant)) search({ keyword: searchKeyword.value, type: searchType.value ?? 0, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updatePayMerchant(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
async function loadPlatforms() { try { const res = await getPayPlatformAll(); platformOptions.value = (res.data || []).map(p => ({ label: p.name, value: p.id })) } catch { /* */ } }
onMounted(async () => { await loadDict(); loadPlatforms(); search({}) })
</script>
