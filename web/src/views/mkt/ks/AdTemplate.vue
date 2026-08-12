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
  </div>
</template>

<script setup>
import { ref, h, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { NButton, NSpace, NPopconfirm, useDialog, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useDict } from '../../../composables/useDict'
import { getKsAdTemplateList, deleteKsAdTemplate, copyKsAdTemplate } from '../../../api/mkt/ks'

const router = useRouter()
const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getKsAdTemplateList)
const { load: loadDict, options } = useDict()
const message = useMessage()
const dialog = useDialog()

const searchKeyword = ref('')
const dictLabel = (key, val) => { const o = options(key).find(i => i.value === val); return o ? o.label : (val ?? '-') }

const columns = [
  { title: 'ID', key: 'id', width: 180, ellipsis: { tooltip: true } },
  { title: '模板名称', key: 'template_name' },
  { title: '广告场景', key: 'type', width: 140, render: (row) => dictLabel('ks_scene', row.type) },
  { title: '出价方式', key: 'base_target', width: 90, render: (row) => dictLabel('ks_bid_type', row.base_target) },
  { title: '单元预算', key: 'unit_day_budget', width: 100, render: (row) => row.unit_day_budget ?? '-' },
  { title: '更新时间', key: 'updated_at', width: 170 },
  { title: '操作', key: 'actions', width: 180, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => router.push(`/ks-ads/ad-template/edit/${row.id}`) }, { default: () => '编辑' }),
    h(NButton, { size: 'tiny', onClick: () => handleCopy(row) }, { default: () => '复制' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value }) }
function handleAdd() { router.push('/ks-ads/ad-template/create') }
function handleCopy(row) {
  dialog.warning({
    title: '复制模板',
    content: '请输入新模板名称',
    positiveText: '确定',
    negativeText: '取消',
    inputValue: row.template_name + '(副本)',
    onPositiveClick: async (d) => {
      const name = d?.inputValue || ''
      if (!name) { message.warning('请输入模板名称'); return false }
      try { await copyKsAdTemplate({ id: row.id, template_name: name }); message.success('复制成功'); search({ keyword: searchKeyword.value }); return true }
      catch (err) { message.error(err.message || '复制失败'); return false }
    },
  })
}
async function onDelete(id) {
  try { await deleteKsAdTemplate(id); message.success('删除成功'); search({ keyword: searchKeyword.value }) }
  catch (err) { message.error(err.message || '删除失败') }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
