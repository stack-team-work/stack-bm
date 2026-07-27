<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/描述" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑代金券' : '新增代金券'" preset="card" style="width: 560px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="name" label="券名称">
          <n-input v-model:value="formData.name" placeholder="请输入券名称" />
        </n-form-item>
        <n-form-item path="use_type" label="使用类型">
          <n-select v-model:value="formData.use_type" :options="useTypeOptions" placeholder="请选择使用类型" />
        </n-form-item>
        <n-form-item path="use_limit" label="可领取次数">
          <n-input-number v-model:value="formData.use_limit" :min="1" style="width: 100%" placeholder="可领取次数" />
        </n-form-item>
        <n-form-item path="total" label="库存">
          <n-input-number v-model:value="formData.total" :min="0" style="width: 100%" placeholder="库存数量" />
        </n-form-item>
        <n-form-item path="total_fee" label="价值">
          <n-input-number v-model:value="formData.total_fee" :min="0" style="width: 100%" placeholder="价值金额" />
        </n-form-item>
        <n-form-item path="desc" label="描述">
          <n-input v-model:value="formData.desc" placeholder="请输入描述" />
        </n-form-item>
        <n-form-item path="stime" label="开始时间">
          <n-date-picker v-model:formatted-value="formData.stime" type="datetime" value-format="x" :default-value="Date.now()" style="width: 100%" />
        </n-form-item>
        <n-form-item path="etime" label="结束时间">
          <n-date-picker v-model:formatted-value="formData.etime" type="datetime" value-format="x" :default-value="Date.now()" style="width: 100%" />
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
import { ref, reactive, h, toRaw, computed, onMounted } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm, NInputNumber, useMessage } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { useDict } from '../../composables/useDict'
import { getGameVoucherList, createGameVoucher, updateGameVoucher, deleteGameVoucher } from '../../api/game'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getGameVoucherList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()
const searchKeyword = ref('')
const searchStatus = ref(null)
const { load: loadDict, options } = useDict()
const statusOptions = computed(() => options('status'))
const useTypeOptions = computed(() => options('game_voucher_use_type'))
const useTypeLabel = { 1: '玩家角色', 2: 'SDK账户' }
const formData = reactive({ name: '', use_type: 1, use_limit: 1, total: 0, total_fee: 0, desc: '', stime: null, etime: null, status: 1 })
function resetForm() { Object.assign(formData, { name: '', use_type: 1, use_limit: 1, total: 0, total_fee: 0, desc: '', stime: null, etime: null, status: 1 }) }
const rules = { name: [{ required: true, message: '请输入券名称', trigger: 'blur' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '券名称', key: 'name' },
  { title: '使用类型', key: 'use_type', width: 90, render: (row) => useTypeLabel[row.use_type] || row.use_type },
  { title: '可领次数', key: 'use_limit', width: 80 },
  { title: '库存', key: 'total', width: 70 },
  { title: '价值', key: 'total_fee', width: 70 },
  { title: '开始时间', key: 'stime', width: 160, render: (row) => row.stime ? formatTime(row.stime) : '' },
  { title: '结束时间', key: 'etime', width: 160, render: (row) => row.etime ? formatTime(row.etime) : '' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.use_type = row.use_type; formData.use_limit = row.use_limit; formData.total = row.total; formData.total_fee = row.total_fee; formData.desc = row.desc; formData.stime = row.stime ? row.stime * 1000 : null; formData.etime = row.etime ? row.etime * 1000 : null; formData.status = row.status; openEdit(row) }
async function handleSubmit() {
  const data = toRaw(formData)
  if (data.stime) data.stime = Math.floor(Number(data.stime) / 1000)
  if (data.etime) data.etime = Math.floor(Number(data.etime) / 1000)
  if (await submit(data, createGameVoucher, updateGameVoucher)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 })
}
async function onDelete(id) { if (await doDelete(id, deleteGameVoucher)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateGameVoucher(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新')   } catch { message.error('更新失败') }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
