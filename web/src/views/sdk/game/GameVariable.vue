<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/Key" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 100px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑变量' : '新增变量'" preset="card" style="width: 640px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="name" label="变量名称">
            <n-input v-model:value="formData.name" placeholder="请输入变量名称" />
          </n-form-item-gi>
          <n-form-item-gi path="key" label="变量Key">
            <n-input v-model:value="formData.key" placeholder="请输入变量Key" :disabled="isEdit" />
          </n-form-item-gi>
          <n-form-item-gi path="value" label="变量值" :span="2">
            <n-input v-model:value="formData.value" type="textarea" :rows="4" placeholder="请输入变量值，多行输入自动保存为JSON数组" />
          </n-form-item-gi>
          <n-form-item-gi path="mark" label="备注">
            <n-input v-model:value="formData.mark" placeholder="请输入备注" />
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
import { NButton, NSpace, NSwitch, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { getGameVariableList, createGameVariable, updateGameVariable, deleteGameVariable } from '../../../api/sdk/game'
import { formatTime } from '../../../utils/format'
import TableActions from '../../../components/TableActions.vue'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getGameVariableList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()

const searchKeyword = ref('')
const searchStatus = ref(null)
const { load: loadDict, options } = useDict()
const statusOptions = computed(() => options('status'))

const formData = reactive({ name: '', key: '', value: '', mark: '', status: 1 })
function resetForm() { Object.assign(formData, { name: '', key: '', value: '', mark: '', status: 1 }) }
const rules = { name: [{ required: true, message: '请输入变量名称', trigger: 'blur' }], key: [{ required: true, message: '请输入变量Key', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '变量名称', key: 'name', width: 120 },
  { title: 'Key', key: 'key', width: 150 },
  { title: '变量值', key: 'value', ellipsis: { tooltip: true }, render: (row) => row.value ? (row.value.length > 50 ? row.value.substring(0, 50) + '...' : row.value) : '' },
  { title: '备注', key: 'mark', width: 120, ellipsis: { tooltip: true } },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(TableActions, { row, edit: () => handleEdit(row), remove: () => onDelete(row.id) }) },
]

function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.key = row.key; formData.value = row.value; formData.mark = row.mark || ''; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createGameVariable, updateGameVariable)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteGameVariable)) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }

async function handleStatusChange(row, val) {
  try {
    await updateGameVariable(row.id, { ...row, status: val ? 1 : 0 })
    row.status = val ? 1 : 0
    message.success('状态已更新')
  } catch {
    message.error('更新失败')
  }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
