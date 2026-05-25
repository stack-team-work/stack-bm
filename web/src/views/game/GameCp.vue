<template>
  <div>
    <n-card title="游戏CP管理" :bordered="false">
      <template #header-extra><n-button type="primary" @click="handleAdd">新增CP</n-button></template>

      <n-space vertical :size="16">
        <n-space>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/标识" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="primary" @click="doSearch">搜索</n-button>
          <n-button @click="resetAll">重置</n-button>
        </n-space>

        <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
      </n-space>
    </n-card>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑CP' : '新增CP'" preset="card" style="width: 500px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="name" label="CP名称">
          <n-input v-model:value="formData.name" placeholder="请输入CP名称" />
        </n-form-item>
        <n-form-item path="mark" label="标识">
          <n-input v-model:value="formData.mark" placeholder="请输入标识" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="phone" label="电话">
          <n-input v-model:value="formData.phone" placeholder="请输入电话" />
        </n-form-item>
        <n-form-item path="addr" label="地址">
          <n-input v-model:value="formData.addr" placeholder="请输入地址" />
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
import { ref, reactive, h, onMounted } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { getGameCpList, createGameCp, updateGameCp, deleteGameCp } from '../../api/game'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getGameCpList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()

const searchKeyword = ref('')
const searchStatus = ref(null)
const statusOptions = [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }]

const formData = reactive({ name: '', mark: '', phone: '', addr: '', status: 1 })
function resetForm() { Object.assign(formData, { name: '', mark: '', phone: '', addr: '', status: 1 }) }

const rules = { name: [{ required: true, message: '请输入CP名称', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 80 },
  { title: 'CP名称', key: 'name' },
  { title: '标识', key: 'mark' },
  { title: '电话', key: 'phone' },
  { title: '地址', key: 'addr', ellipsis: { tooltip: true } },
  { title: '状态', key: 'status', width: 80, render: (row) => h(NSwitch, { value: row.status === 1, readonly: true, size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 180 },
  {
    title: '操作', key: 'actions', width: 160,
    render: (row) => h(NSpace, null, {
      default: () => [
        h(NButton, { size: 'small', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
        h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, {
          default: () => '确认删除?', trigger: () => h(NButton, { size: 'small', type: 'error' }, { default: () => '删除' }),
        }),
      ],
    }),
  },
]

function doSearch() { search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 }) }
function resetAll() { searchKeyword.value = ''; searchStatus.value = null; resetSearch() }

function handleAdd() { resetForm(); open() }
function handleEdit(row) {
  resetForm()
  formData.name = row.name; formData.mark = row.mark; formData.phone = row.phone || ''; formData.addr = row.addr || ''; formData.status = row.status
  openEdit(row)
}

async function handleSubmit() {
  const ok = await submit(formData)
  if (ok) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 })
}

async function onDelete(id) {
  const ok = await doDelete(id, deleteGameCp)
  if (ok) search({ keyword: searchKeyword.value, status: searchStatus.value ?? -1 })
}

onMounted(() => search({}))
</script>
