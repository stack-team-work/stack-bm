<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索菜单名称" clearable style="width: 180px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchParent" :options="parentSearchOptions" placeholder="父级菜单" clearable style="width: 150px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="false" />
    </n-space>

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
import { useModal } from '../../composables/useModal'
import { getMenuAll, createMenu, updateMenu, deleteMenu } from '../../api/system'

const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const loading = ref(false)
const tableData = ref([])
const searchKeyword = ref('')
const searchParent = ref(null)
const allMenus = ref([])
const parentSearchOptions = computed(() => {
  return [{ label: '全部菜单', value: null }, ...allMenus.value.filter(m => m.parent === 0).map(m => ({ label: m.name, value: m.id }))]
})

const parentOptions = computed(() => [{ label: '顶级菜单', value: 0 }, ...allMenus.value.filter(m => m.id !== editId.value && m.is_deleted !== 1).map(m => ({ label: m.name, value: m.id }))])

const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '类型', key: 'type', width: 60, render: (row) => h(NTag, { type: row.type === 1 ? 'info' : 'default', size: 'small' }, { default: () => row.type === 1 ? '菜单' : '按钮' }) },
  { title: '名称', key: 'name' },
  { title: '路径', key: 'path', width: 160 },
  { title: '父级', key: 'parent', width: 60 },
  { title: '图标', key: 'icon', width: 80 },
  { title: '排序', key: 'sort', width: 60 },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, readonly: true, size: 'small' }) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() {
  let list = allMenus.value || []
  if (searchKeyword.value) {
    const kw = searchKeyword.value.toLowerCase()
    list = list.filter(m => m.name.toLowerCase().includes(kw))
  }
  if (searchParent.value !== null && searchParent.value !== undefined) {
    list = list.filter(m => m.parent === searchParent.value)
  }
  tableData.value = list
}

function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.type = row.type; formData.name = row.name; formData.path = row.path; formData.parent = row.parent; formData.icon = row.icon || ''; formData.sort = row.sort; formData.author = row.author || ''; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createMenu, updateMenu)) loadAll() }
async function onDelete(id) { if (await doDelete(id, deleteMenu)) loadAll() }

async function loadAll() { try { const res = await getMenuAll(); allMenus.value = res.data || []; doSearch() } catch { /* */ } }
onMounted(() => { loadAll() })
</script>
