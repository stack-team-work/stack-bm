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

    <n-modal v-model:show="showModal" :title="isEdit ? '编辑广告模板' : '新增广告模板'" preset="card" style="width: 900px" :mask-closable="false">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="120">
        <n-grid :cols="2" :x-gap="20">
          <n-form-item-gi path="template_name" label="模板名称">
            <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="app_id" label="适用游戏">
            <n-select v-model:value="formData.app_id" :options="gameOptions" multiple placeholder="选择适用游戏" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="promotion_purpose_type" label="推广目标">
            <n-select v-model:value="formData.promotion_purpose_type" :options="dict('bili_promotion_purpose')" />
          </n-form-item-gi>
          <n-form-item-gi path="ad_type" label="广告类型">
            <n-select v-model:value="formData.ad_type" :options="dict('bili_ad_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="support_auto" label="投放类型">
            <n-select v-model:value="formData.support_auto" :options="dict('bili_delivery_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="speed_mode" label="投放方式">
            <n-select v-model:value="formData.speed_mode" :options="dict('bili_speed_mode')" />
          </n-form-item-gi>
          <n-form-item-gi path="promotion_content_type" label="推广内容">
            <n-select v-model:value="formData.promotion_content_type" :options="dict('bili_promotion_content')" />
          </n-form-item-gi>
          <n-form-item-gi path="mini_game_type" label="小游戏类型">
            <n-select v-model:value="formData.mini_game_type" :options="dict('bili_mini_game')" />
          </n-form-item-gi>
          <n-form-item-gi path="ad_budget_limit_type" label="预算类型">
            <n-select v-model:value="formData.ad_budget_limit_type" :options="dict('bili_budget_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="ad_budget" label="项目预算">
            <n-input-number v-model:value="formData.ad_budget" :min="0" style="width: 100%" placeholder="不限预算时忽略" />
          </n-form-item-gi>
          <n-form-item-gi path="unit_budget" label="单元预算">
            <n-input-number v-model:value="formData.unit_budget" :min="0" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="unit_date_type" label="投放日期">
            <n-select v-model:value="formData.unit_date_type" :options="dict('bili_unit_date_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="launch_begin_date" label="开始日期">
            <n-date-picker v-model:formatted-value="formData.launch_begin_date" type="date" value-format="yyyy-MM-dd" :default-value="Date.now()" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="launch_end_date" label="结束日期">
            <n-date-picker v-model:formatted-value="formData.launch_end_date" type="date" value-format="yyyy-MM-dd" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="unit_time_type" label="投放时间段">
            <n-select v-model:value="formData.unit_time_type" :options="dict('bili_unit_time_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="is_no_bid" label="竞价策略">
            <n-select v-model:value="formData.is_no_bid" :options="dict('bili_strategy_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="base_target" label="出价方式">
            <n-select v-model:value="formData.base_target" :options="dict('bili_bid_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="base_bid" label="基础出价">
            <n-input-number v-model:value="formData.base_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="custom_bid_type" label="自定义出价">
            <n-select v-model:value="formData.custom_bid_type" :options="dict('bili_custom_bid_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="min_base_bid" label="出价下限">
            <n-input-number v-model:value="formData.min_base_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="max_base_bid" label="出价上限">
            <n-input-number v-model:value="formData.max_base_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="cpa_target" label="优化目标">
            <n-select v-model:value="formData.cpa_target" :options="dict('bili_optimize_gold')" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="cpa_bid" label="目标转化出价">
            <n-input-number v-model:value="formData.cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="min_cpa_bid" label="转化出价下限">
            <n-input-number v-model:value="formData.min_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="max_cpa_bid" label="转化出价上限">
            <n-input-number v-model:value="formData.max_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="deep_cpa_target" label="深度优化目标">
            <n-select v-model:value="formData.deep_cpa_target" :options="dict('bili_deep_optimize_gold')" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="deep_cpa_bid" label="深度转化出价">
            <n-input-number v-model:value="formData.deep_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="min_deep_cpa_bid" label="深度出价下限">
            <n-input-number v-model:value="formData.min_deep_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="max_deep_cpa_bid" label="深度出价上限">
            <n-input-number v-model:value="formData.max_deep_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi path="dual_bid_two_stage_optimization" label="深度优化方式">
            <n-select v-model:value="formData.dual_bid_two_stage_optimization" :options="dict('bili_deep_gold_type')" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="is_bili_native" label="投放模式">
            <n-select v-model:value="formData.is_bili_native" :options="dict('bili_launch_type')" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="channel_id" label="流量类型">
            <n-select v-model:value="formData.channel_id" :options="dict('bili_network_type')" clearable />
          </n-form-item-gi>
          <n-grid-item>
            <n-form-item path="is_smart_material" label="智能素材" label-placement="left">
              <n-switch v-model:value="formData.is_smart_material" :checked-value="1" :unchecked-value="0" />
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
import { ref, reactive, h, onMounted } from 'vue'
import { NButton, NSpace, NPopconfirm, useMessage } from 'naive-ui'
import { useTable } from '../../../composables/useTable'
import { useModal } from '../../../composables/useModal'
import { useDict } from '../../../composables/useDict'
import { useOptions } from '../../../composables/useOptions'
import { getBiliAdTemplateList, createBiliAdTemplate, updateBiliAdTemplate, deleteBiliAdTemplate, copyBiliAdTemplate } from '../../../api/mkt/bili'

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(getBiliAdTemplateList)
const { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, submit, handleDelete: doDelete } = useModal()
const { load: loadDict, options } = useDict()
const { loadOptions } = useOptions()
const message = useMessage()

const searchKeyword = ref('')
const gameOptions = ref([])

const dict = (key) => options(key)
const dictLabel = (key, val) => { const o = options(key).find(i => i.value === val); return o ? o.label : (val ?? '-') }

const formData = reactive({
  template_name: '', app_id: [], promotion_purpose_type: 2, ad_type: 0,
  ad_budget_limit_type: 2, ad_budget: 0, unit_budget: 0,
  support_auto: 0, speed_mode: 1, promotion_content_type: 2, mini_game_type: 5,
  unit_date_type: 1, launch_begin_date: '', launch_end_date: '',
  unit_time_type: 0, is_no_bid: 0,
  base_target: 1, base_bid: 0, min_base_bid: 0, max_base_bid: 0,
  cpa_target: null, cpa_bid: 0, min_cpa_bid: 0, max_cpa_bid: 0,
  deep_cpa_target: null, deep_cpa_bid: 0, min_deep_cpa_bid: 0, max_deep_cpa_bid: 0,
  dual_bid_two_stage_optimization: null, is_bili_native: null, channel_id: null,
  custom_bid_type: 'CUSTOM_BID_TYPE_NORMAL', is_smart_material: 1,
})
function resetForm() {
  Object.assign(formData, {
    template_name: '', app_id: [], promotion_purpose_type: 2, ad_type: 0,
    ad_budget_limit_type: 2, ad_budget: 0, unit_budget: 0,
    support_auto: 0, speed_mode: 1, promotion_content_type: 2, mini_game_type: 5,
    unit_date_type: 1, launch_begin_date: '', launch_end_date: '',
    unit_time_type: 0, is_no_bid: 0,
    base_target: 1, base_bid: 0, min_base_bid: 0, max_base_bid: 0,
    cpa_target: null, cpa_bid: 0, min_cpa_bid: 0, max_cpa_bid: 0,
    deep_cpa_target: null, deep_cpa_bid: 0, min_deep_cpa_bid: 0, max_deep_cpa_bid: 0,
    dual_bid_two_stage_optimization: null, is_bili_native: null, channel_id: null,
    custom_bid_type: 'CUSTOM_BID_TYPE_NORMAL', is_smart_material: 1,
  })
}
const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

const columns = [
  { title: 'ID', key: 'id', width: 180, ellipsis: { tooltip: true } },
  { title: '模板名称', key: 'template_name' },
  { title: '推广目标', key: 'promotion_purpose_type', width: 120, render: (row) => dictLabel('bili_promotion_purpose', row.promotion_purpose_type) },
  { title: '出价方式', key: 'base_target', width: 80, render: (row) => dictLabel('bili_bid_type', row.base_target) },
  { title: '单元预算', key: 'unit_budget', width: 90, render: (row) => row.unit_budget ?? '-' },
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
  Object.keys(formData).forEach(k => { if (row[k] !== undefined && row[k] !== null) formData[k] = row[k] })
  formData.app_id = row.app_id || []
  openEdit(row)
}
async function handleCopy(row) {
  const name = prompt('请输入新模板名称', row.template_name + '(副本)')
  if (!name) return
  try { await copyBiliAdTemplate({ id: row.id, template_name: name }); message.success('复制成功'); search({ keyword: searchKeyword.value }) }
  catch (err) { message.error(err.message || '复制失败') }
}
async function handleSubmit() {
  const data = { ...formData }
  if (await submit(data, createBiliAdTemplate, updateBiliAdTemplate)) search({ keyword: searchKeyword.value })
}
async function onDelete(id) { if (await doDelete(id, deleteBiliAdTemplate)) search({ keyword: searchKeyword.value }) }

onMounted(async () => { await loadDict(); gameOptions.value = await loadOptions('game'); search({}) })
</script>
