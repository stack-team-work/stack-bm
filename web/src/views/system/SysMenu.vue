<template>
  <div>
    <n-card title="菜单管理" :bordered="false">
      <template #header-extra><n-button type="primary" @click="handleAdd">新增菜单</n-button></template>

      <n-space vertical :size="16">
        <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="false" />
      </n-space>
    </n-card>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑菜单' : '新增菜单'" preset="card" style="width: 550px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="type" label="类型">
          <n-select v-model:value="formData.type" :options="typeOptions" />
        </n-form-item>
        <n-form-item path="name" label="菜单名称">
          <n-input v-model:value="formData.name" placeholder="请输入菜单名称" />
        </n-form-item>
        <n-form-item path="path" label="路由路径">
          <n-input v-model:value="formData.path" placeholder="如 /admin" :disabled="isEdit" />
        </n-form-item>
        <n-form-item path="parent" label="父级菜单">
          <n-select v-model:value="formData.parent" :options="parentOptions" placeholder="顶级菜单" clearable />
        </n-form-item>
        <n-form-item path="icon" label="图标">
          <n-input v-model:value="formData.icon" placeholder="请输入图标名称" />
        </n-form-item>
        <n-form-item path="sort" label="排序">
          <n-input-number v-model:value="formData.sort" :min="0" style="width: 100%" />
        </n-form-item>
        <n-form-item path="author" label="作者">
          <n-input v-model:value="formData.author" placeholder="请输入作者" />
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
import { NButton, NSpace, NSwitch, NPopconfirm, NTag } from 'naive-ui'
import { useTable } from '../../composables/useTable'
import { useModal } from '../../composables/useModal'
import { getMenuList, getMenuAll, createMenu, updateMenu, deleteMenu } from '../../api/system'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getMenuList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()

const formData = reactive({ type: 1, name: '', path: '', parent: 0, icon: '', sort: 0, author: '', status: 1 })
function resetForm() { Object.assign(formData, { type: 1, name: '', path: '', parent: 0, icon: '', sort: 0, author: '', status: 1 }) }

const rules = { name: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }], path: [{ required: true, message: '请输入路由路径', trigger: 'blur' }] }

const typeOptions = [
  { label: '菜单', value: 1 },
  { label: '按钮', value: 2 },
]

const allMenus = ref([])
const parentOptions = computed(() => {
  return [{ label: '顶级菜单', value: 0 }, ...allMenus.value.filter(m => m.id !== editId.value).map(m => ({ label: m.name, value: m.id }))]
})

const columns = [
  { title: 'ID', key: 'id', width: 70 },
  { title: '类型', key: 'type', width: 70, render: (row) => h(NTag, { type: row.type === 1 ? 'info' : 'default', size: 'small' }, { default: () => row.type === 1 ? '菜单' : '按钮' }) },
  { title: '名称', key: 'name' },
  { title: '路径', key: 'path', width: 180 },
  { title: '父级ID', key: 'parent', width: 70 },
  { title: '图标', key: 'icon', width: 80 },
  { title: '排序', key: 'sort', width: 60 },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, readonly: true, size: 'small' }) },
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

function handleAdd() { resetForm(); open() }
function handleEdit(row) {
  resetForm()
  formData.type = row.type; formData.name = row.name; formData.path = row.path; formData.parent = row.parent; formData.icon = row.icon || ''; formData.sort = row.sort; formData.author = row.author || ''; formData.status = row.status
  openEdit(row)
}

async function handleSubmit() {
  const ok = await submit(formData)
  if (ok) loadAll()
}

async function onDelete(id) {
  const ok = await doDelete(id, deleteMenu)
  if (ok) loadAll()
}

async function loadAll() {
  try { const res = await getMenuAll(); allMenus.value = res.data || []; tableData.value = allMenus.value } catch { /* */ }
}

onMounted(() => { loadAll() })
</script>
