<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索用户名/姓名" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchGroupId" :options="groupSearchOptions" placeholder="分组" clearable style="width: 150px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑账号' : '新增账号'" preset="card" style="width: 500px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="username" label="用户名">
          <n-input v-model:value="formData.username" placeholder="请输入用户名" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="password" label="密码" :required="!isEdit">
          <n-input v-model:value="formData.password" type="password" :placeholder="isEdit ? '留空则不修改' : '请输入密码'" />
        </n-form-item>
        <n-form-item path="name" label="姓名">
          <n-input v-model:value="formData.name" placeholder="请输入姓名" />
        </n-form-item>
        <n-form-item path="phone" label="手机号">
          <n-input v-model:value="formData.phone" placeholder="请输入手机号" />
        </n-form-item>
        <n-form-item path="group_id" label="所属角色">
          <n-select v-model:value="formData.group_id" :options="groupOptions" placeholder="请选择角色" />
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
import { ref, reactive, h, onMounted, computed } from 'vue'
import { NButton, NSpace, NSwitch, NPopconfirm } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { getAdminList, createAdmin, updateAdmin, deleteAdmin, getAdminGroupAll } from '../../api/system'

const { loading, tableData, pagination, search, resetSearch, handlePageChange, handlePageSizeChange } = useTable(getAdminList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()

const searchKeyword = ref('')
const searchGroupId = ref(null)
const groupOptions = ref([])
const groupSearchOptions = computed(() => [{ label: '全部角色', value: 0 }, ...groupOptions.value])

const formData = reactive({ username: '', password: '', name: '', phone: '', group_id: 0, status: 1 })
function resetForm() { Object.assign(formData, { username: '', password: '', name: '', phone: '', group_id: 0, status: 1 }) }
const rules = { username: [{ required: true, message: '请输入用户名', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '用户名', key: 'username' },
  { title: '姓名', key: 'name' },
  { title: '手机号', key: 'phone' },
  { title: '角色', key: 'group_id', width: 100, render: (row) => { const g = groupOptions.value.find(o => o.value === row.group_id); return g ? g.label : row.group_id } },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, readonly: true, size: 'small' }) },
  { title: '创建时间', key: 'created_at', width: 170 },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value, group_id: searchGroupId.value ?? 0 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.username = row.username; formData.name = row.name; formData.phone = row.phone; formData.group_id = row.group_id; formData.status = row.status; openEdit(row) }

async function handleSubmit() {
  if (!isEdit.value && !formData.password) { const { useMessage } = await import('naive-ui'); useMessage().error('请输入密码'); return }
  if (await submit(formData, createAdmin, updateAdmin)) search({ keyword: searchKeyword.value, group_id: searchGroupId.value ?? 0 })
}
async function onDelete(id) { if (await doDelete(id, deleteAdmin)) search({ keyword: searchKeyword.value, group_id: searchGroupId.value ?? 0 }) }

async function loadGroups() { try { const res = await getAdminGroupAll(); groupOptions.value = (res.data || []).map(g => ({ label: g.name, value: g.id })) } catch { /* */ } }
onMounted(() => { loadGroups(); search({}) })
</script>
