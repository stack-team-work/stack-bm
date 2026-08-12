<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索模板名称" clearable style="width: 220px" @keyup.enter="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>

      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑标题包模板' : '新增标题包模板'" preset="card" style="width: 860px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-form-item path="template_name" label="模板名称">
          <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
        </n-form-item>

        <n-divider style="margin: 4px 0 12px">描述（2-10个字符）</n-divider>
        <n-space v-for="(item, i) in formData.description" :key="i" align="center" style="margin-bottom: 8px">
          <n-input v-model:value="formData.description[i]" placeholder="请输入描述" style="width: 360px" />
          <n-button size="tiny" type="error" quaternary @click="removeDescription(i)">移除</n-button>
        </n-space>
        <n-button size="small" @click="addDescription">+ 添加描述</n-button>

        <n-divider style="margin-top: 16px">标题物料（2-40个字符，支持智能词包）</n-divider>
        <n-card v-for="(m, i) in formData.title_materials" :key="i" size="small" style="margin-bottom: 12px">
          <template #header>
            <n-space align="center" :size="8">
              <n-text>标题 {{ i + 1 }}</n-text>
              <n-button size="tiny" type="error" quaternary @click="removeMaterial(i)">移除</n-button>
            </n-space>
          </template>
          <n-input v-model:value="m.title" placeholder="请输入标题文本（{词包} 为智能词包占位）" />
          <n-space v-for="(s, j) in m.smart_title_list" :key="j" align="center" style="margin-top: 8px">
            <n-select v-model:value="s.type" :options="dict('bili_title_word_type')" placeholder="词包类型" style="width: 160px" />
            <n-input v-model:value="s.default_value" placeholder="默认值" style="width: 240px" />
            <n-button size="tiny" type="error" quaternary @click="removeSmartWord(m, j)">移除</n-button>
          </n-space>
          <n-button size="tiny" style="margin-top: 8px" @click="addSmartWord(m)">+ 智能词包</n-button>
        </n-card>
        <n-button size="small" @click="addMaterial">+ 添加标题</n-button>
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
import { NButton, NSpace, NText, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { getBiliTitleTemplateList, createBiliTitleTemplate, updateBiliTitleTemplate, deleteBiliTitleTemplate, copyBiliTitleTemplate } from '../../../api/mkt/bili'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getBiliTitleTemplateList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()

const searchKeyword = ref('')
const dict = (key) => options(key)

const formData = reactive({ template_name: '', description: [], title_materials: [] })

function newMaterial() { return { title: '', smart_title_list: [] } }
function newSmartWord() { return { type: 1, default_value: '' } }

function resetForm() {
  Object.assign(formData, {
    template_name: '',
    description: [''],
    title_materials: [{ title: '', smart_title_list: [{ type: 1, default_value: '' }] }],
  })
}
const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 180, ellipsis: { tooltip: true } },
  { title: '模板名称', key: 'template_name' },
  { title: '标题数', key: 'title_num', width: 70 },
  { title: '描述数', key: 'description_num', width: 70 },
  { title: '更新时间', key: 'updated_at', width: 170 },
  { title: '操作', key: 'actions', width: 180, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NButton, { size: 'tiny', onClick: () => handleCopy(row) }, { default: () => '复制' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function addDescription() { formData.description.push('') }
function removeDescription(i) { formData.description.splice(i, 1) }
function addMaterial() { formData.title_materials.push(newMaterial()) }
function removeMaterial(i) { formData.title_materials.splice(i, 1) }
function addSmartWord(m) { m.smart_title_list.push(newSmartWord()) }
function removeSmartWord(m, j) { m.smart_title_list.splice(j, 1) }

function doSearch() { search({ keyword: searchKeyword.value }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) {
  resetForm()
  formData.template_name = row.template_name || ''
  formData.description = (row.description || []).filter(Boolean)
  formData.title_materials = (row.title_materials || []).map(m => ({
    title: m.title || '',
    smart_title_list: (m.smart_title_list || []).map(s => ({ type: s.type, default_value: s.default_value || '' })),
  }))
  if (!formData.description.length) formData.description.push('')
  if (!formData.title_materials.length) formData.title_materials.push(newMaterial())
  openEdit(row)
}
async function handleCopy(row) {
  const name = prompt('请输入新模板名称', row.template_name + '(副本)')
  if (!name) return
  try { await copyBiliTitleTemplate({ id: row.id, template_name: name }); message.success('复制成功'); search({ keyword: searchKeyword.value }) }
  catch (err) { message.error(err.message || '复制失败') }
}
async function handleSubmit() {
  const data = {
    template_name: formData.template_name,
    description: formData.description.filter(Boolean),
    title_materials: formData.title_materials
      .filter(m => m.title.trim())
      .map(m => ({ title: m.title, smart_title_list: (m.smart_title_list || []).filter(s => s.type) })),
  }
  if (await submit(data, createBiliTitleTemplate, updateBiliTitleTemplate)) search({ keyword: searchKeyword.value })
}
async function onDelete(id) { if (await doDelete(id, deleteBiliTitleTemplate)) search({ keyword: searchKeyword.value }) }

onMounted(async () => { await loadDict(); search({}) })
</script>
