<template>
  <div>
    <n-card title="管理员分组" :bordered="false">
      <template #header-extra><n-button type="primary" @click="handleAdd">新增分组</n-button></template>

      <n-space vertical :size="16">
        <n-space>
          <n-input v-model:value="searchKeyword" placeholder="搜索分组名称" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-button type="primary" @click="doSearch">搜索</n-button>
          <n-button @click="resetAll">重置</n-button>
        </n-space>

        <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
      </n-space>
    </n-card>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑分组' : '新增分组'" preset="card" style="width: 500px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="mark" label="标识">
          <n-input v-model:value="formData.mark" placeholder="请输入标识" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="name" label="分组名称">
          <n-input v-model:value="formData.name" placeholder="请输入分组名称" />
        </n-form-item>
        <n-form-item path="description" label="描述">
          <n-input v-model:value="formData.description" type="textarea" placeholder="请输入描述" />
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
import { getAdminGroupList, createAdminGroup, updateAdminGroup, deleteAdminGroup } from '../../api/system'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getAdminGroupList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()

const searchKeyword = ref('')

const formData = reactive({ mark: '', name: '', description: '', status: 1 })
function resetForm() { Object.assign(formData, { mark: '', name: '', description: '', status: 1 }) }

const rules = { name: [{ required: true, message: '请输入分组名称', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 80 },
  { title: '标识', key: 'mark', width: 100 },
  { title: '分组名称', key: 'name' },
  { title: '描述', key: 'description' },
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

function doSearch() { search({ keyword: searchKeyword.value }) }
function resetAll() { searchKeyword.value = ''; resetSearch() }

function handleAdd() { resetForm(); open() }
function handleEdit(row) {
  resetForm()
  formData.mark = row.mark; formData.name = row.name; formData.description = row.description; formData.status = row.status
  openEdit(row)
}

async function handleSubmit() {
  const ok = await submit(formData)
  if (ok) search({ keyword: searchKeyword.value })
}

async function onDelete(id) {
  const ok = await doDelete(id, deleteAdminGroup)
  if (ok) search({ keyword: searchKeyword.value })
}

onMounted(() => search({}))
</script>
