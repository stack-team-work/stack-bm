<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/描述" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑礼包' : '新增礼包'" preset="card" style="width: 720px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="name" label="礼包名称">
            <n-input v-model:value="formData.name" placeholder="请输入礼包名称" />
          </n-form-item-gi>
          <n-form-item-gi path="get_type" label="领取类型">
            <n-select v-model:value="formData.get_type" :options="getTypeOptions" placeholder="请选择领取类型" />
          </n-form-item-gi>
          <n-form-item-gi path="type" label="礼包类型">
            <n-select v-model:value="formData.type" :options="typeOptions" placeholder="请选择礼包类型" />
          </n-form-item-gi>
          <n-grid-item>
            <n-form-item path="is_code" label="需要激活码" label-placement="left">
              <n-switch v-model:value="formData.is_code" :checked-value="1" :unchecked-value="0" checked-text="需要" unchecked-text="不需要" />
            </n-form-item>
          </n-grid-item>
          <n-form-item-gi path="stime" label="开始时间">
            <n-date-picker v-model:formatted-value="formData.stime" type="datetime" value-format="x" :default-value="Date.now()" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="etime" label="结束时间">
            <n-date-picker v-model:formatted-value="formData.etime" type="datetime" value-format="x" :default-value="Date.now()" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="cond" label="领取条件" :span="2">
            <n-input v-model:value="formData.cond" type="textarea" :rows="2" placeholder="请输入领取条件" />
          </n-form-item-gi>
          <n-form-item-gi path="desc" label="描述">
            <n-input v-model:value="formData.desc" placeholder="请输入描述" />
          </n-form-item-gi>
          <n-grid-item>
            <n-form-item path="status" label="状态" label-placement="left">
              <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="有效" unchecked-text="无效" />
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
import { ref, reactive, h, toRaw, computed, onMounted } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm, NIcon, NInputNumber, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { getGameGiftList, createGameGift, updateGameGift, deleteGameGift } from '../../../api/sdk/game'
import { formatTime } from '../../../utils/format'
import { CreateOutline, TrashOutline } from '@vicons/ionicons5'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getGameGiftList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()
const searchKeyword = ref('')
const searchStatus = ref(null)
const { load: loadDict, options } = useDict()
const statusOptions = computed(() => options('status'))
const getTypeOptions = computed(() => options('game_gift_get_type'))
const getTypeLabel = { 1: '单次', 2: '每日' }
const typeOptions = computed(() => options('game_gift_type'))
const formData = reactive({ name: '', get_type: 1, is_code: 0, type: 1, cond: '', desc: '', stime: null, etime: null, status: 1 })
function resetForm() { Object.assign(formData, { name: '', get_type: 1, is_code: 0, type: 1, cond: '', desc: '', stime: null, etime: null, status: 1 }) }
const rules = { name: [{ required: true, message: '请输入礼包名称', trigger: 'blur' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '礼包名称', key: 'name' },
  { title: '领取类型', key: 'get_type', width: 80, render: (row) => getTypeLabel[row.get_type] || row.get_type },
  { title: '需要激活码', key: 'is_code', width: 90, render: (row) => row.is_code === 1 ? '是' : '否' },
  { title: '描述', key: 'desc', width: 120, ellipsis: { tooltip: true } },
  { title: '开始时间', key: 'stime', width: 160, render: (row) => row.stime ? formatTime(row.stime) : '' },
  { title: '结束时间', key: 'etime', width: 160, render: (row) => row.etime ? formatTime(row.etime) : '' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(CreateOutline) }), ' 编辑'] }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(TrashOutline) }), ' 删除'] }) }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.get_type = row.get_type; formData.is_code = row.is_code; formData.type = row.type; formData.cond = row.cond; formData.desc = row.desc; formData.stime = row.stime ? row.stime * 1000 : null; formData.etime = row.etime ? row.etime * 1000 : null; formData.status = row.status; openEdit(row) }
async function handleSubmit() {
  const data = toRaw(formData)
  if (data.stime) data.stime = Math.floor(Number(data.stime) / 1000)
  if (data.etime) data.etime = Math.floor(Number(data.etime) / 1000)
  if (await submit(data, createGameGift, updateGameGift)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 })
}
async function onDelete(id) { if (await doDelete(id, deleteGameGift)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateGameGift(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新')   } catch { message.error('更新失败') }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
