<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑快手广告模板' : '新增快手广告模板' }}</n-text>
    </n-space>

    <n-card style="margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top" label-width="130">
        <n-divider style="margin-top: 4px">基本信息</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="template_name" label="模板名称">
            <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="app_id" label="适用应用">
            <n-select v-model:value="formData.app_id" :options="gameOptions" multiple placeholder="选择适用应用" clearable />
          </n-form-item-gi>
        </n-grid>

        <n-divider>投放设置</n-divider>
        <n-form-item label="投放目标" path="market_target">
          <n-radio-group v-model:value="formData.market_target">
            <n-radio-button v-for="o in dict('ks_market_target')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="广告场景" path="type">
          <n-radio-group v-model:value="formData.type">
            <n-radio-button v-for="o in dict('ks_scene')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="广告类型" path="ad_type">
          <n-radio-group v-model:value="formData.ad_type">
            <n-radio-button v-for="o in dict('ks_ad_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="投放方式" path="auto_manage">
          <n-radio-group v-model:value="formData.auto_manage">
            <n-radio-button v-for="o in dict('ks_auto_manager')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="投放位置" path="scene_id">
          <n-checkbox-group v-model:value="formData.scene_id">
            <n-checkbox v-for="o in dict('ks_scene_inventory')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="投放日期" path="unit_date_type">
          <n-radio-group v-model:value="formData.unit_date_type">
            <n-radio-button v-for="o in dict('ks_unit_date_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid v-if="formData.unit_date_type === 1" :cols="2" :x-gap="24">
          <n-form-item-gi label="开始日期" path="begin_time">
            <n-date-picker v-model:formatted-value="formData.begin_time" type="date" value-format="yyyy-MM-dd" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="结束日期" path="end_time">
            <n-date-picker v-model:formatted-value="formData.end_time" type="date" value-format="yyyy-MM-dd" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item label="投放时间段" path="unit_time_type">
          <n-radio-group v-model:value="formData.unit_time_type">
            <n-radio-button v-for="o in dict('ks_unit_time_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.unit_time_type === 1" label="时间段" path="schedule_time">
          <n-input v-model:value="formData.schedule_time" placeholder="如 00:00-08:00,12:00-14:00" style="max-width: 400px" />
        </n-form-item>
        <n-form-item label="投放速度" path="speed_type">
          <n-radio-group v-model:value="formData.speed_type">
            <n-radio-button v-for="o in dict('ks_speed_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="展现方式" path="show_mode">
          <n-radio-group v-model:value="formData.show_mode">
            <n-radio-button v-for="o in dict('ks_show_mode')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>

        <n-divider>预算</n-divider>
        <n-form-item label="预算类型" path="budget_type">
          <n-radio-group v-model:value="formData.budget_type">
            <n-radio-button v-for="o in dict('ks_budget_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.budget_type === 2" label="每日预算" path="day_budget">
          <n-input-number v-model:value="formData.day_budget" :min="0" style="width: 240px" />
        </n-form-item>
        <n-form-item label="单元预算类型" path="unit_budget_type">
          <n-radio-group v-model:value="formData.unit_budget_type">
            <n-radio-button v-for="o in dict('ks_budget_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.unit_budget_type === 2" label="单元每日预算" path="unit_day_budget">
          <n-input-number v-model:value="formData.unit_day_budget" :min="0" style="width: 240px" />
        </n-form-item>

        <n-divider>创意设置</n-divider>
        <n-form-item label="创意类型" path="creative_unit_type">
          <n-radio-group v-model:value="formData.creative_unit_type">
            <n-radio-button v-for="o in dict('ks_creative_unit')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="小游戏类型" path="mini_type">
          <n-radio-group v-model:value="formData.mini_type">
            <n-radio-button v-for="o in dict('ks_mini_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="素材优选" path="auto_photo_scope">
          <n-radio-group v-model:value="formData.auto_photo_scope">
            <n-radio-button v-for="o in dict('ks_auto_photo_scope')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="行动号召" path="action_bar_text">
          <n-input v-model:value="formData.action_bar_text" placeholder="请输入行动号召文案" style="max-width: 400px" />
        </n-form-item>
        <n-form-item label="资产挖矿" path="asset_mining">
          <n-switch v-model:value="formData.asset_mining" :checked-value="1" :unchecked-value="0" />
        </n-form-item>
        <n-form-item label="智能封面" path="smart_cover">
          <n-switch v-model:value="formData.smart_cover" :checked-value="1" :unchecked-value="0" />
        </n-form-item>

        <n-divider>出价与优化目标</n-divider>
        <n-form-item label="竞价策略" path="bid_type">
          <n-radio-group v-model:value="formData.bid_type">
            <n-radio-button v-for="o in dict('ks_bid_strategy')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="自定义出价" path="custom_bid_type">
          <n-radio-group v-model:value="formData.custom_bid_type">
            <n-radio-button v-for="o in dict('ks_bid_way')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="出价方式" path="base_target">
          <n-radio-group v-model:value="formData.base_target">
            <n-radio-button v-for="o in dict('ks_bid_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="手动/自动出价" path="smart_bid">
          <n-radio-group v-model:value="formData.smart_bid">
            <n-radio-button v-for="o in dict('ks_smart_bid')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid :cols="3" :x-gap="24">
          <n-form-item-gi label="基础出价" path="base_bid">
            <n-input-number v-model:value="formData.base_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="出价下限" path="min_base_bid">
            <n-input-number v-model:value="formData.min_base_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="出价上限" path="max_base_bid">
            <n-input-number v-model:value="formData.max_base_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item label="优化目标" path="cpa_target">
          <n-radio-group v-model:value="formData.cpa_target">
            <n-radio-button v-for="o in dict('ks_ocpx_action')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="目标行为" path="target_action">
          <n-radio-group v-model:value="formData.target_action">
            <n-radio-button v-for="o in dict('ks_target_action')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid :cols="3" :x-gap="24">
          <n-form-item-gi label="目标转化出价" path="cpa_bid">
            <n-input-number v-model:value="formData.cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="转化出价下限" path="min_cpa_bid">
            <n-input-number v-model:value="formData.min_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="转化出价上限" path="max_cpa_bid">
            <n-input-number v-model:value="formData.max_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item label="深度优化目标" path="deep_cpa_target">
          <n-radio-group v-model:value="formData.deep_cpa_target">
            <n-radio-button v-for="o in dict('ks_deep_conversion')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid v-if="formData.deep_cpa_target" :cols="3" :x-gap="24">
          <n-form-item-gi label="深度转化出价" path="deep_cpa_bid">
            <n-input-number v-model:value="formData.deep_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="深度出价下限" path="min_deep_cpa_bid">
            <n-input-number v-model:value="formData.min_deep_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="深度出价上限" path="max_deep_cpa_bid">
            <n-input-number v-model:value="formData.max_deep_cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item label="快速搜索" path="quick_search">
          <n-switch v-model:value="formData.quick_search" :checked-value="1" :unchecked-value="0" />
        </n-form-item>
        <n-form-item label="目标探索" path="target_explore">
          <n-switch v-model:value="formData.target_explore" :checked-value="1" :unchecked-value="0" />
        </n-form-item>

        <n-space style="margin-top: 16px">
          <n-button type="primary" :loading="submitLoading" @click="handleSubmit" size="medium">
            {{ isEdit ? '保存修改' : '确认创建' }}
          </n-button>
          <n-button @click="goBack" size="medium">取消</n-button>
        </n-space>
      </n-form>
    </n-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useMessage } from 'naive-ui'
import { useDict } from '../../../composables/useDict'
import { useOptions } from '../../../composables/useOptions'
import { getKsAdTemplateDetail, createKsAdTemplate, updateKsAdTemplate } from '../../../api/mkt/ks'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const { load: loadDict, options } = useDict()
const { loadOptions } = useOptions()

const formRef = ref(null)
const submitLoading = ref(false)
const isEdit = ref(false)
const editId = ref(null)
const gameOptions = ref([])
const dict = (key) => options(key)

const formData = reactive({
  template_name: '', app_id: [], market_target: 4, type: 2, ad_type: 0, auto_manage: 0,
  budget_type: 1, day_budget: 0, unit_budget_type: 1, unit_day_budget: 0,
  bid_type: 0, auto_photo_scope: 0, creative_unit_type: 7, mini_type: 2,
  unit_date_type: 0, begin_time: null, end_time: null, unit_time_type: 0, schedule_time: '',
  scene_id: [], custom_bid_type: 0, base_target: 10, base_bid: 0, min_base_bid: 0, max_base_bid: 0,
  cpa_target: null, target_action: null, cpa_bid: 0, min_cpa_bid: 0, max_cpa_bid: 0,
  deep_cpa_target: null, deep_cpa_bid: 0, min_deep_cpa_bid: 0, max_deep_cpa_bid: 0,
  smart_bid: 0, speed_type: 1, show_mode: 1, quick_search: 1, target_explore: 1,
  asset_mining: 0, smart_cover: 0, action_bar_text: '',
})
const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function goBack() { router.push({ path: '/ks-ads', query: { tab: 'ad' } }) }

function fill(data) {
  Object.keys(formData).forEach(k => { if (data[k] !== undefined && data[k] !== null) formData[k] = data[k] })
  formData.app_id = data.app_id || []
  formData.scene_id = data.scene_id || []
}

async function handleSubmit() {
  try { await formRef.value?.validate() } catch { return }
  submitLoading.value = true
  try {
    if (isEdit.value) { await updateKsAdTemplate(editId.value, { ...formData }); message.success('更新成功') }
    else { await createKsAdTemplate({ ...formData }); message.success('创建成功') }
    goBack()
  } catch (err) { message.error(err.message || '操作失败') }
  finally { submitLoading.value = false }
}

onMounted(async () => {
  await loadDict()
  gameOptions.value = await loadOptions('game')
  const id = route.params.id
  if (id) {
    isEdit.value = true
    editId.value = id
    try {
      const res = await getKsAdTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  }
})
</script>
