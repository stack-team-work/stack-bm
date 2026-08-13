<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索名称/标识" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑支付平台' : '新增支付平台'" preset="card" style="width: 560px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="name" label="平台名称">
            <n-input v-model:value="formData.name" placeholder="请输入平台名称" />
          </n-form-item-gi>
          <n-form-item-gi path="mark" label="平台标识">
            <n-input v-model:value="formData.mark" placeholder="请输入平台标识（留空自动生成）" :disabled="isEdit" />
          </n-form-item-gi>
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
import { ref, reactive, h } from 'vue'
import { NButton, NSpace, useMessage } from 'naive-ui'
import TableActions from '../../../components/TableActions.vue'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { getPayPlatformList, createPayPlatform, updatePayPlatform, deletePayPlatform } from '../../../api/sdk/pay'
import { formatTime } from '../../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getPayPlatformList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const message = useMessage()
const searchKeyword = ref('')
const formData = reactive({ name: '', mark: '' })
function resetForm() { Object.assign(formData, { name: '', mark: '' }) }
const rules = { name: [{ required: true, message: '请输入平台名称', trigger: 'blur' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '平台名称', key: 'name' },
  { title: '平台标识', key: 'mark' },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(TableActions, { row, edit: () => handleEdit(row), remove: () => onDelete(row.id) }) },
]
function doSearch() { search({ keyword: searchKeyword.value }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.name = row.name; formData.mark = row.mark; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createPayPlatform, updatePayPlatform)) search({ keyword: searchKeyword.value }) }
async function onDelete(id) { if (await doDelete(id, deletePayPlatform)) search({ keyword: searchKeyword.value }) }
</script>
