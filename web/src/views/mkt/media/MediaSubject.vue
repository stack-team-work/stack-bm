<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/标识" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑主体' : '新增主体'" preset="card" style="width: 560px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="name" label="主体名称">
            <n-input v-model:value="formData.name" placeholder="请输入主体名称" />
          </n-form-item-gi>
          <n-form-item-gi path="mark" label="主体标识">
            <n-input v-model:value="formData.mark" placeholder="请输入主体标识（留空自动生成）" :disabled="isEdit" />
          </n-form-item-gi>
          <n-grid-item :span="2">
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
import { ref, reactive, h, computed, onMounted } from 'vue'
import { NSwitch, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { getMediaSubjectList, createMediaSubject, updateMediaSubject, deleteMediaSubject } from '../../../api/mkt/media'
import { formatTime } from '../../../utils/format'
import { useDict } from '../../../composables/useDict'
import TableActions from '../../../components/TableActions.vue'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getMediaSubjectList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()
const searchKeyword = ref('')
const searchStatus = ref(null)
const statusOptions = computed(() => options('status'))
const formData = reactive({ name: '', mark: '', status: 1 })
function resetForm() { Object.assign(formData, { name: '', mark: '', status: 1 }) }
const rules = { name: [{ required: true, message: '请输入主体名称', trigger: 'blur' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '主体名称', key: 'name' },
  { title: '主体标识', key: 'mark' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(TableActions, { row, edit: () => handleEdit(row), remove: () => onDelete(row.id) }) },
]
function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.mark = row.mark; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createMediaSubject, updateMediaSubject)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteMediaSubject)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateMediaSubject(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
