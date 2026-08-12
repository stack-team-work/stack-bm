<template>
  <div>
    <n-space vertical :size="16">
      <n-space>
        <n-input v-model:value="searchKeyword" placeholder="搜索名称/字段/标识" clearable style="width: 200px" @keyup.enter="doSearch" />
        <n-select v-model:value="searchReportType" :options="reportTypeOptions" placeholder="报表类型" clearable style="width: 130px" @update:value="doSearch" />
        <n-select v-model:value="searchIndicatorType" :options="indicatorTypeOptions" placeholder="指标类型" clearable style="width: 130px" @update:value="doSearch" />
        <n-select v-model:value="searchStatus" :options="statusOptions" placeholder="状态" clearable style="width: 120px" @update:value="doSearch" />
        <n-button type="primary" size="small" @click="doSearch">搜索</n-button>
        <n-button type="success" size="small" @click="handleAdd">新增</n-button>
      </n-space>
      <n-data-table :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
    </n-space>
    <n-modal v-model:show="showModal" :title="isEdit ? '编辑指标' : '新增指标'" preset="card" style="width: 640px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-grid :cols="2" :x-gap="16">
          <n-form-item-gi path="report_type" label="报表类型">
            <n-select v-model:value="formData.report_type" :options="reportTypeOptions" placeholder="请选择报表类型" />
          </n-form-item-gi>
          <n-form-item-gi path="indicator_type" label="指标类型">
            <n-select v-model:value="formData.indicator_type" :options="indicatorTypeOptions" placeholder="请选择指标类型" />
          </n-form-item-gi>
          <n-form-item-gi path="name" label="列名">
            <n-input v-model:value="formData.name" placeholder="请输入列名" />
          </n-form-item-gi>
          <n-form-item-gi path="field" label="字段名">
            <n-input v-model:value="formData.field" placeholder="请输入字段名" :disabled="isEdit" />
          </n-form-item-gi>
          <n-form-item-gi path="mark" label="标识">
            <n-input v-model:value="formData.mark" placeholder="请输入标识（留空自动生成）" :disabled="isEdit" />
          </n-form-item-gi>
          <n-grid-item>
            <n-form-item path="default" label="默认选中" label-placement="left">
              <n-switch v-model:value="formData.default" :checked-value="1" :unchecked-value="0" checked-text="是" unchecked-text="否" />
            </n-form-item>
          </n-grid-item>
          <n-grid-item :span="2">
            <n-form-item path="status" label="状态" label-placement="left">
              <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="启用" unchecked-text="禁用" />
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
import { NButton, NSpace, NSwitch, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { getSysColumnList, createSysColumn, updateSysColumn, deleteSysColumn } from '../../../api/bm/sys'
import { formatTime } from '../../../utils/format'
import { useDict } from '../../../composables/useDict'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getSysColumnList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const message = useMessage()
const searchKeyword = ref('')
const searchReportType = ref(null)
const searchIndicatorType = ref(null)
const searchStatus = ref(null)
const statusOptions = computed(() => options('status'))
const reportTypeOptions = computed(() => options('sys_column_report_type'))
const indicatorTypeOptions = computed(() => options('sys_column_indicator_type'))
const reportTypeLabel = { 1: '投放报表' }
const indicatorTypeLabel = { 1: '属性指标', 2: '媒体指标', 3: 'BM指标', 4: 'N日指标' }
const formData = reactive({ report_type: 1, indicator_type: 1, name: '', field: '', mark: '', default: 0, status: 1 })
function resetForm() { Object.assign(formData, { report_type: 1, indicator_type: 1, name: '', field: '', mark: '', default: 0, status: 1 }) }
const rules = { name: [{ required: true, message: '请输入列名', trigger: 'blur' }], field: [{ required: true, message: '请输入字段名', trigger: 'blur' }] }
const columns = [
  { title: 'ID', key: 'id', width: 60 },
  { title: '报表类型', key: 'report_type', width: 90, render: (row) => reportTypeLabel[row.report_type] || row.report_type },
  { title: '指标类型', key: 'indicator_type', width: 90, render: (row) => indicatorTypeLabel[row.indicator_type] || row.indicator_type },
  { title: '列名', key: 'name' },
  { title: '字段名', key: 'field' },
  { title: '标识', key: 'mark' },
  { title: '默认', key: 'default', width: 60, render: (row) => row.default === 1 ? '是' : '否' },
  { title: '状态', key: 'status', width: 70, render: (row) => h(NSwitch, { value: row.status === 1, onUpdateValue: (val) => handleStatusChange(row, val), size: 'small' }) },
  { title: '操作', key: 'actions', width: 140, render: (row) => h(NSpace, null, { default: () => [
    h(NButton, { size: 'tiny', onClick: () => handleEdit(row) }, { default: () => '编辑' }),
    h(NPopconfirm, { onPositiveClick: () => onDelete(row.id) }, { default: () => '确认删除?', trigger: () => h(NButton, { size: 'tiny', type: 'error' }, { default: () => '删除' }) }),
  ]}) },
]
function doSearch() { search({ keyword: searchKeyword.value, report_type: searchReportType.value ?? 0, indicator_type: searchIndicatorType.value ?? 0, status: searchStatus.value ?? -1 }) }
function handleAdd() { resetForm(); open() }
function handleEdit(row) { resetForm(); formData.report_type = row.report_type; formData.indicator_type = row.indicator_type; formData.name = row.name; formData.field = row.field; formData.mark = row.mark; formData.default = row.default; formData.status = row.status; openEdit(row) }
async function handleSubmit() { if (await submit(formData, createSysColumn, updateSysColumn)) search({ keyword: searchKeyword.value, report_type: searchReportType.value ?? 0, indicator_type: searchIndicatorType.value ?? 0, status: searchStatus.value ?? -1 }) }
async function onDelete(id) { if (await doDelete(id, deleteSysColumn)) search({ keyword: searchKeyword.value, report_type: searchReportType.value ?? 0, indicator_type: searchIndicatorType.value ?? 0, status: searchStatus.value ?? -1 }) }
async function handleStatusChange(row, val) {
  try { await updateSysColumn(row.id, { ...row, status: val ? 1 : 0 }); row.status = val ? 1 : 0; message.success('状态已更新') } catch { message.error('更新失败') }
}

onMounted(async () => { await loadDict(); search({}) })
</script>
