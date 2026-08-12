<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索角色名称" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑角色' : '新增角色'" preset="card" style="width: 560px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="mark" label="标识">
            <n-input v-model:value="formData.mark" placeholder="请输入标识" :disabled="isEdit" />
          </n-form-item-gi>
          <n-form-item-gi path="name" label="角色名称">
            <n-input v-model:value="formData.name" placeholder="请输入角色名称" />
          </n-form-item-gi>
          <n-form-item-gi path="description" label="描述" :span="2">
            <n-input v-model:value="formData.description" type="textarea" :rows="2" placeholder="请输入描述" />
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
import { ref, reactive, h, onMounted } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm, NIcon, useMessage } from 'naive-ui'
import { CreateOutline, TrashOutline } from '@vicons/ionicons5'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { getAdminGroupList, createAdminGroup, updateAdminGroup, deleteAdminGroup } from '../../../api/bm/sys'
import { formatTime } from '../../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getAdminGroupList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()

const searchKeyword = ref('')
const formData = reactive({ mark: '', name: '', description: '', status: 1 })
function resetForm() { Object.assign(formData, { mark: '', name: '', description: '', status: 1 }) }
const rules = { name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '标识', key: 'mark', width: 90 },
  { title: '角色名称', key: 'name' },
  { title: '描述', key: 'description' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(CreateOutline) }), ' 编辑'] }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(TrashOutline) }), ' 删除'] }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.mark = row.mark; formData.name = row.name; formData.description = row.description; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createAdminGroup, updateAdminGroup)) search({ keyword: searchKeyword.value }) }
async function onDelete(id) { if (await doDelete(id, deleteAdminGroup)) search({ keyword: searchKeyword.value }) }

async function handleStatusChange(row, val) {
  try {
    await updateAdminGroup(row.id, { ...row, status: val ? 1 : 0 })
    row.status = val ? 1 : 0
    message.success('状态已更新')
  } catch {
    message.error('更新失败')
  }
}

onMounted(() => search({}))
</script>
