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

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑定向包模板' : '新增定向包模板'" preset="card" style="width: 760px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="140">
        <n-grid :cols="2" :x-gap="20">
          <n-form-item-gi path="template_name" label="模板名称" :span="2">
            <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="description" label="描述" :span="2">
            <n-input v-model:value="formData.description" placeholder="请输入描述" />
          </n-form-item-gi>
          <n-form-item-gi path="age_list" label="年龄段">
            <n-select v-model:value="formData.age_list" :options="dict('bili_age')" multiple placeholder="选择年龄段" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="gender_list" label="性别">
            <n-select v-model:value="formData.gender_list" :options="dict('bili_gender')" multiple placeholder="选择性别" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="network_list" label="网络环境">
            <n-select v-model:value="formData.network_list" :options="dict('bili_network')" multiple placeholder="选择网络" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="phone_price_list" label="手机价位">
            <n-select v-model:value="formData.phone_price_list" :options="dict('bili_phone_price')" multiple placeholder="选择手机价位" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="os_list" label="操作系统">
            <n-select v-model:value="formData.os_list" :options="dict('bili_os')" multiple placeholder="选择系统" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="installed_user_filter" label="安装过滤">
            <n-select v-model:value="formData.installed_user_filter" :options="dict('bili_installed_type')" multiple placeholder="选择安装过滤" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="area_type" label="地域类型">
            <n-select v-model:value="formData.area_type" :options="dict('bili_area_type')" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="converted_user_filter" label="转化用户过滤">
            <n-select v-model:value="formData.converted_user_filter" :options="dict('bili_converted_user')" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="area_list" label="地域列表(JSON)" :span="2">
            <n-input v-model:value="formData.area_list" type="textarea" :rows="2" placeholder="如 [1,2,3]（地理划分ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="area_level_list" label="地域层级(JSON)" :span="2">
            <n-input v-model:value="formData.area_level_list" type="textarea" :rows="2" placeholder="如 [1,2]（发展划分ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="crowd_pack" label="人群包(JSON)" :span="2">
            <n-input v-model:value="formData.crowd_pack" type="textarea" :rows="2" placeholder="如 [1,2]（DMP人群包ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="archive_content" label="档案内容(JSON)" :span="2">
            <n-input v-model:value="formData.archive_content" type="textarea" :rows="2" placeholder="高级参数，留空即可" />
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
import { ref, reactive, h, onMounted } from 'vue'
import { NButton, NSpace, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { getBiliAudienceTemplateList, createBiliAudienceTemplate, updateBiliAudienceTemplate, deleteBiliAudienceTemplate, copyBiliAudienceTemplate } from '../../../api/mkt'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getBiliAudienceTemplateList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()

const searchKeyword = ref('')
const dict = (key) => options(key)
const dictLabel = (key, val) => { const o = options(key).find(i => i.value === val); return o ? o.label : (val ?? '-') }
const arrLabel = (key, arr) => Array.isArray(arr) && arr.length ? arr.map(v => dictLabel(key, v)).join('、') : '-'

const jsonFields = ['area_list', 'area_level_list', 'crowd_pack', 'archive_content']

const formData = reactive({
  template_name: '', description: '',
  age_list: [], gender_list: [], network_list: [], phone_price_list: [], os_list: [], installed_user_filter: [],
  area_type: null, converted_user_filter: 0,
  area_list: '[]', area_level_list: '[]', crowd_pack: '[]', archive_content: '[]',
})
function resetForm() {
  Object.assign(formData, {
    template_name: '', description: '',
    age_list: [], gender_list: [], network_list: [], phone_price_list: [], os_list: [], installed_user_filter: [],
    area_type: null, converted_user_filter: 0,
    area_list: '[]', area_level_list: '[]', crowd_pack: '[]', archive_content: '[]',
  })
}
const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function toJSONText(v) { return v === undefined || v === null ? '[]' : JSON.stringify(v) }

const columns = [
  { title: 'ID', key: 'id', width: 180, ellipsis: { tooltip: true } },
  { title: '模板名称', key: 'template_name' },
  { title: '年龄段', key: 'age_list', width: 140, ellipsis: { tooltip: true }, render: (row) => arrLabel('bili_age', row.age_list) },
  { title: '性别', key: 'gender_list', width: 90, render: (row) => arrLabel('bili_gender', row.gender_list) },
  { title: '操作系统', key: 'os_list', width: 120, render: (row) => arrLabel('bili_os', row.os_list) },
  { title: '更新时间', key: 'updated_at', width: 170 },
  { title: '操作', key: 'actions', width: 180, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NButton, { size: 'tiny', onClick: () => handleCopy(row) }, { default: () => '复制' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) {
  resetForm()
  Object.keys(formData).forEach(k => {
    if (row[k] === undefined || row[k] === null) return
    if (jsonFields.includes(k)) formData[k] = toJSONText(row[k])
    else formData[k] = row[k]
  })
  openEdit(row)
}
async function handleCopy(row) {
  const name = prompt('请输入新模板名称', row.template_name + '(副本)')
  if (!name) return
  try { await copyBiliAudienceTemplate({ id: row.id, template_name: name }); message.success('复制成功'); search({ keyword: searchKeyword.value }) }
  catch (err) { message.error(err.message || '复制失败') }
}
async function handleSubmit() {
  const data = { ...formData }
  jsonFields.forEach(f => {
    try { const v = JSON.parse(data[f]); data[f] = Array.isArray(v) ? v : [] } catch { data[f] = [] }
  })
  data.age_list = data.age_list || []; data.gender_list = data.gender_list || []; data.network_list = data.network_list || []
  data.phone_price_list = data.phone_price_list || []; data.os_list = data.os_list || []; data.installed_user_filter = data.installed_user_filter || []
  if (await submit(data, createBiliAudienceTemplate, updateBiliAudienceTemplate)) search({ keyword: searchKeyword.value })
}
async function onDelete(id) { if (await doDelete(id, deleteBiliAudienceTemplate)) search({ keyword: searchKeyword.value }) }

onMounted(async () => { await loadDict(); search({}) })
</script>
