<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索飞书用户ID" clearable style="width: 200px" @keyup.enter="doSearch" />
          <n-select v-model:value="searchAdminId" :options="adminOptions" placeholder="管理员" clearable style="width: 160px" @update:value="doSearch" />
          <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑飞书绑定' : '新增飞书绑定'" preset="card" style="width: 560px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="admin_id" label="管理员">
            <n-select v-model:value="formData.admin_id" :options="adminOptions" placeholder="请选择管理员" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="feishu_user_id" label="飞书用户ID">
            <n-input v-model:value="formData.feishu_user_id" placeholder="请输入飞书用户ID" :disabled="isEdit" />
          </n-form-item-gi>
          <n-grid-item :span="2">
            <n-form-item path="status" label="状态" label-placement="left">
              <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="正常" unchecked-text="异常" />
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
import { NButton, NSpace, NSwitch, NIcon, useMessage } from 'naive-ui'
import { CreateOutline } from '@vicons/ionicons5'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { useOptions } from '../../../composables/useOptions'
import { getFeishuUserList, createFeishuUser, updateFeishuUser, updateFeishuUserStatus } from '../../../api/bm/sys'
import { formatTime } from '../../../utils/format'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getFeishuUserList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const { loadOptions } = useOptions()
const message = useMessage()
const searchKeyword = ref('')
const searchAdminId = ref(null)
const searchStatus = ref(null)
const adminOptions = ref([])
const statusOptions = computed(() => options('status'))
const formData = reactive({ admin_id: null, feishu_user_id: '', status: 1 })
function resetForm() { Object.assign(formData, { admin_id: null, feishu_user_id: '', status: 1 }) }
const rules = { feishu_user_id: [{ required: true, message: '请输入飞书用户ID', trigger: 'blur' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '管理员', key: 'admin_id', width: 120, render: (row) => { const a = adminOptions.value.find(o => o.value === row.admin_id); return a ? a.label : row.admin_id } },
  { title: '飞书用户ID', key: 'feishu_user_id' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170, render: (row) => formatTime(row.created_at) },
  { title: '操作', key: 'actions', width: 80, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(CreateOutline) }), ' 编辑'] }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, admin_id: searchAdminId.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.admin_id = row.admin_id; formData.feishu_user_id = row.feishu_user_id; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createFeishuUser, updateFeishuUser)) search({ keyword: searchKeyword.value, admin_id: searchAdminId.value ?? 0, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateFeishuUserStatus(row.id, { status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}
onMounted(async () => { await loadDict(); adminOptions.value = await loadOptions('admin'); search({}) })
</script>
