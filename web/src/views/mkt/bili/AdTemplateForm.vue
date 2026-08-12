<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑广告模板' : '新增广告模板' }}</n-text>
    </n-space>

    <n-card style="max-width: 1100px; margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top" :label-width="130">
        <n-divider style="margin-top: 4px">基本信息</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="template_name" label="模板名称">
            <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="app_id" label="适用游戏">
            <n-select v-model:value="formData.app_id" :options="gameOptions" multiple placeholder="选择适用游戏" clearable />
          </n-form-item-gi>
        </n-grid>

        <n-divider>推广设置</n-divider>
        <n-form-item label="推广目标" path="promotion_purpose_type">
          <n-radio-group v-model:value="formData.promotion_purpose_type">
            <n-radio-button v-for="o in dict('bili_promotion_purpose')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="广告类型" path="ad_type">
          <n-radio-group v-model:value="formData.ad_type">
            <n-radio-button v-for="o in dict('bili_ad_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="投放类型" path="support_auto">
          <n-radio-group v-model:value="formData.support_auto">
            <n-radio-button v-for="o in dict('bili_delivery_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="投放方式" path="speed_mode">
          <n-radio-group v-model:value="formData.speed_mode">
            <n-radio-button v-for="o in dict('bili_speed_mode')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="推广内容" path="promotion_content_type">
          <n-radio-group v-model:value="formData.promotion_content_type">
            <n-radio-button v-for="o in dict('bili_promotion_content')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.promotion_content_type === 65" label="小游戏类型" path="mini_game_type">
          <n-radio-group v-model:value="formData.mini_game_type">
            <n-radio-button v-for="o in dict('bili_mini_game')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="投放日期" path="unit_date_type">
          <n-radio-group v-model:value="formData.unit_date_type">
            <n-radio-button v-for="o in dict('bili_unit_date_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid v-if="formData.unit_date_type === 1" :cols="2" :x-gap="24">
          <n-form-item-gi label="开始日期" path="launch_begin_date">
            <n-date-picker v-model:formatted-value="formData.launch_begin_date" type="date" value-format="yyyy-MM-dd" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="结束日期" path="launch_end_date">
            <n-date-picker v-model:formatted-value="formData.launch_end_date" type="date" value-format="yyyy-MM-dd" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item label="投放时间段" path="unit_time_type">
          <n-radio-group v-model:value="formData.unit_time_type">
            <n-radio-button v-for="o in dict('bili_unit_time_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.unit_time_type === 1" label="时间段" path="launch_time">
          <n-input v-model:value="formData.launch_time" placeholder="如 00:00-08:00,12:00-14:00" style="max-width: 400px" />
        </n-form-item>
        <n-form-item label="竞价策略" path="is_no_bid">
          <n-radio-group v-model:value="formData.is_no_bid">
            <n-radio-button v-for="o in dict('bili_strategy_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="智能素材" path="is_smart_material">
          <n-switch v-model:value="formData.is_smart_material" :checked-value="1" :unchecked-value="0" />
        </n-form-item>

        <n-divider>预算</n-divider>
        <n-form-item label="预算类型" path="ad_budget_limit_type">
          <n-radio-group v-model:value="formData.ad_budget_limit_type">
            <n-radio-button v-for="o in dict('bili_budget_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="单元预算" path="unit_budget">
          <n-input-number v-model:value="formData.unit_budget" :min="0" style="width: 240px" />
        </n-form-item>
        <n-form-item v-if="formData.ad_budget_limit_type !== 2" label="项目预算" path="ad_budget">
          <n-input-number v-model:value="formData.ad_budget" :min="0" style="width: 240px" placeholder="不限预算时忽略" />
        </n-form-item>

        <n-divider>出价与优化目标</n-divider>
        <n-form-item label="出价方式" path="base_target">
          <n-radio-group v-model:value="formData.base_target">
            <n-radio-button v-for="o in dict('bili_bid_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="自定义出价" path="custom_bid_type">
          <n-radio-group v-model:value="formData.custom_bid_type">
            <n-radio-button v-for="o in dict('bili_custom_bid_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
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
            <n-radio-button v-for="o in dict('bili_optimize_gold')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
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
            <n-radio-button v-for="o in dict('bili_deep_optimize_gold')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.deep_cpa_target" label="深度优化方式" path="dual_bid_two_stage_optimization">
          <n-radio-group v-model:value="formData.dual_bid_two_stage_optimization">
            <n-radio-button v-for="o in dict('bili_deep_gold_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
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
        <n-form-item label="投放模式" path="is_bili_native">
          <n-radio-group v-model:value="formData.is_bili_native">
            <n-radio-button v-for="o in dict('bili_launch_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="流量类型" path="channel_id">
          <n-radio-group v-model:value="formData.channel_id">
            <n-radio-button v-for="o in dict('bili_network_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
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
import { getBiliAdTemplateDetail, createBiliAdTemplate, updateBiliAdTemplate } from '../../../api/mkt/bili'

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
  template_name: '', app_id: [], promotion_purpose_type: 2, ad_type: 0,
  ad_budget_limit_type: 2, ad_budget: 0, unit_budget: 0,
  support_auto: 0, speed_mode: 1, promotion_content_type: 2, mini_game_type: 5,
  unit_date_type: 1, launch_begin_date: null, launch_end_date: null,
  unit_time_type: 0, launch_time: '', is_no_bid: 0,
  base_target: 1, base_bid: 0, min_base_bid: 0, max_base_bid: 0,
  cpa_target: null, cpa_bid: 0, min_cpa_bid: 0, max_cpa_bid: 0,
  deep_cpa_target: null, deep_cpa_bid: 0, min_deep_cpa_bid: 0, max_deep_cpa_bid: 0,
  dual_bid_two_stage_optimization: null, is_bili_native: null, channel_id: null,
  custom_bid_type: 'CUSTOM_BID_TYPE_NORMAL', is_smart_material: 1,
})
const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function goBack() { router.push({ path: '/bili-ads', query: { tab: 'ad' } }) }

function fill(data) {
  Object.keys(formData).forEach(k => { if (data[k] !== undefined && data[k] !== null) formData[k] = data[k] })
  formData.app_id = data.app_id || []
}

async function handleSubmit() {
  try { await formRef.value?.validate() } catch { return }
  submitLoading.value = true
  try {
    if (isEdit.value) { await updateBiliAdTemplate(editId.value, { ...formData }); message.success('更新成功') }
    else { await createBiliAdTemplate({ ...formData }); message.success('创建成功') }
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
      const res = await getBiliAdTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  }
})
</script>
