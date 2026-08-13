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
import { NButton, NIcon, useDialog, useMessage } from 'naive-ui'
import { CopyOutline } from '@vicons/ionicons5'
import { useTable } from '../../../composables/useTable'
import { getKsTitleTemplateList, deleteKsTitleTemplate, copyKsTitleTemplate } from '../../../api/mkt/ks'
import TableActions from '../../../components/TableActions.vue'

const router = useRouter()
const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getKsTitleTemplateList)
const message = useMessage()
const dialog = useDialog()

const searchKeyword = ref('')

const columns = [
  { title: 'ID', key: 'id', width: 180, ellipsis: { tooltip: true } },
  { title: '模板名称', key: 'template_name' },
  { title: '标题数', key: 'title_num', width: 70 },
  { title: '更新时间', key: 'updated_at', width: 170 },
  { title: '操作', key: 'actions', width: 180, render: (row) => h(TableActions, { row, edit: () => router.push(`/ks-ads/title-template/edit/${row.id}`), remove: () => onDelete(row.id) }, { extra: () => h(NButton, { size: 'tiny', onClick: () => handleCopy(row) }, { default: () => [h(NIcon, { size: 14 }, { default: () => h(CopyOutline) }), ' 复制'] }) }) },
]

function doSearch() { search({ keyword: searchKeyword.value }) }
function handleAdd() { router.push('/ks-ads/title-template/create') }
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
      try { await copyKsTitleTemplate({ id: row.id, template_name: name }); message.success('复制成功'); search({ keyword: searchKeyword.value }); return true }
      catch (err) { message.error(err.message || '复制失败'); return false }
    },
  })
}
async function onDelete(id) {
  try { await deleteKsTitleTemplate(id); message.success('删除成功'); search({ keyword: searchKeyword.value }) }
  catch (err) { message.error(err.message || '删除失败') }
}

onMounted(() => search({}))
</script>
