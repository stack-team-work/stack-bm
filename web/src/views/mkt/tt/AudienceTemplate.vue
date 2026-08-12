<template>
  <div>
    <n-card :bordered="false">
      <div class="search-bar">
        <n-space :size="12" align="center" wrap>
          <n-input v-model:value="searchKeyword" placeholder="搜索模板名称" clearable style="width: 220px" @keyup.enter="doSearch" />
          <n-button type="info" size="small" @click="doSearch">搜索</n-button>
          <n-button type="primary" size="small" @click="handleAdd">新增</n-button>
        </n-space>
      </div>
      <n-data-table :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-card>
  </div>
</template>

<script setup>
import { ref, h, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { NButton, NSpace, NPopconfirm, NIcon, useDialog, useMessage } from 'naive-ui'
import { CreateOutline, CopyOutline, TrashOutline } from '@vicons/ionicons5'
import { useTable } from '../../../composables/useTable'
import { useDict } from '../../../composables/useDict'
import { getTtAudienceTemplateList, deleteTtAudienceTemplate, copyTtAudienceTemplate } from '../../../api/mkt/tt'

const router = useRouter()
const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getTtAudienceTemplateList)
const { load: loadDict, options } = useDict()
const message = useMessage()
const dialog = useDialog()

const searchKeyword = ref('')
const dictLabel = (key, val) => { const o = options(key).find(i => i.value === val); return o ? o.label : (val ?? '-') }
const arrLabel = (key, arr) => Array.isArray(arr) && arr.length ? arr.map(v => dictLabel(key, v)).join('、') : '-'

const columns = [
  { title: 'ID', key: 'id', width: 180, ellipsis: { tooltip: true } },
  { title: '模板名称', key: 'template_name' },
  { title: '定向类型', key: 'landing_type', width: 130, render: (row) => dictLabel('tt_audience_landing_type', row.landing_type) },
  { title: '性别', key: 'gender', width: 80, render: (row) => dictLabel('tt_gender', row.gender) },
  { title: '年龄', key: 'age', width: 140, ellipsis: { tooltip: true }, render: (row) => arrLabel('tt_age', row.age) },
  { title: '更新时间', key: 'updated_at', width: 170 },
  { title: '操作', key: 'actions', width: 180, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => router.push(`/tt-ads/audience-template/edit/${row.id}`) }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(CreateOutline) }), ' 编辑'] }),
    h(NButton, { size: 'tiny', onClick: () => handleCopy(row) }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(CopyOutline) }), ' 复制'] }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(TrashOutline) }), ' 删除'] }) }),
  ]}) },
]

function doSearch() { search({ keyword: searchKeyword.value }) }
function handleAdd() { router.push('/tt-ads/audience-template/create') }
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
      try { await copyTtAudienceTemplate({ id: row.id, template_name: name }); message.success('复制成功'); search({ keyword: searchKeyword.value }); return true }
      catch (err) { message.error(err.message || '复制失败'); return false }
    },
  })
}
async function onDelete(id) {
  try { await deleteTtAudienceTemplate(id); message.success('删除成功'); search({ keyword: searchKeyword.value }) }
  catch (err) { message.error(err.message || '删除失败') }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
