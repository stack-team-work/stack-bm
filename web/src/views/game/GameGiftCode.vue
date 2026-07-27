<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索激活码" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchGiftId" :options="giftOptions" placeholder="礼包" clearable style="width: 160px" @update:value="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑礼包码' : '新增礼包码'" preset="card" style="width: 500px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="gift_id" label="所属礼包">
          <n-select v-model:value="formData.gift_id" :options="giftOptions" placeholder="请选择礼包" clearable />
        </n-form-item>
        <n-form-item path="code" label="激活码">
          <n-input v-model:value="formData.code" placeholder="请输入激活码" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="status" label="状态">
          <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="已使用" unchecked-text="未使用" />
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
import { NButton, NSpace, NSwitch, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { useDict } from '../../composables/useDict'
import { getGameGiftCodeList, createGameGiftCode, updateGameGiftCode, deleteGameGiftCode, getGameGiftAll } from '../../api/game'
import { formatTime } from '../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getGameGiftCodeList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()
const searchKeyword = ref('')
const searchGiftId = ref(null)
const searchStatus = ref(null)
const giftOptions = ref([])
const { load: loadDict, options } = useDict()
const statusOptions = computed(() => options('status'))
const formData = reactive({ gift_id: null, code: '', status: 0 })
function resetForm() { Object.assign(formData, { gift_id: null, code: '', status: 0 }) }
const rules = { code: [{ required: true, message: '请输入激活码', trigger: 'blur' }], gift_id: [{ required: true, type: 'number', message: '请选择礼包', trigger: 'change' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '激活码', key: 'code' },
  { title: '礼包', key: 'gift_id', width: 120, render: (row) => { const g = giftOptions.value.find(o => o.value === row.gift_id); return g ? g.label : '' } },
  { title: '状态', key: 'status', width: 80, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, gift_id: searchGiftId.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.gift_id = row.gift_id; formData.code = row.code; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createGameGiftCode, updateGameGiftCode)) search({ keyword: searchKeyword.value, gift_id: searchGiftId.value ?? 0, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteGameGiftCode)) search({ keyword: searchKeyword.value, gift_id: searchGiftId.value ?? 0, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateGameGiftCode(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
async function loadGifts() { try { const res = await getGameGiftAll(); giftOptions.value = (res.data || []).map(g => ({ label: g.name, value: g.id })) } catch { /* */ } }
onMounted(async () => { await loadDict(); loadGifts(); search({}) })
</script>
