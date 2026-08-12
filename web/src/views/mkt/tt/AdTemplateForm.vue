<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑广告模板' : '新增广告模板' }}</n-text>
    </n-space>

    <n-card style="margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top" :label-width="120">
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
        <n-form-item label="推广目的" path="landing_type">
          <n-radio-group v-model:value="formData.landing_type">
            <n-radio-button v-for="o in dict('tt_landing_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.landing_type === 'APP'" label="子目标" path="app_promotion_type">
          <n-radio-group v-model:value="formData.app_promotion_type">
            <n-radio-button v-for="o in dict('tt_app_promotion_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.landing_type === 'MICRO_GAME'" label="小程序类型" path="micro_promotion_type">
          <n-radio-group v-model:value="formData.micro_promotion_type">
            <n-radio-button v-for="o in dict('tt_micro_promotion_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="营销场景" path="marketing_goal">
          <n-radio-group v-model:value="formData.marketing_goal">
            <n-radio-button v-for="o in dict('tt_marketing_goal')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="投放模式" path="delivery_mode">
          <n-radio-group v-model:value="formData.delivery_mode">
            <n-radio-button v-for="o in dict('tt_delivery_mode')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="广告类型" path="ad_type">
          <n-radio-group v-model:value="formData.ad_type">
            <n-radio-button v-for="o in dict('tt_ad_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.landing_type === 'APP'" label="下载方式" path="download_type">
          <n-radio-group v-model:value="formData.download_type">
            <n-radio-button v-for="o in dict('tt_download_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.landing_type === 'APP'" label="下载模式" path="download_mode">
          <n-radio-group v-model:value="formData.download_mode">
            <n-radio-button v-for="o in dict('tt_download_mode')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="投放类型" path="mkt_ad_type">
          <n-radio-group v-model:value="formData.mkt_ad_type">
            <n-radio-button v-for="o in dict('tt_mkt_ad_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.landing_type === 'LINK'" label="投放内容" path="promotion_type">
          <n-radio-group v-model:value="formData.promotion_type">
            <n-radio-button v-for="o in dict('tt_promotion_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>

        <n-divider>广告位置</n-divider>
        <n-form-item label="广告位置大类" path="inventory_catalog">
          <n-radio-group v-model:value="formData.inventory_catalog">
            <n-radio-button v-for="o in dict('tt_inventory_catalog')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.inventory_catalog === 'MANUAL'" label="投放位置" path="inventory_type">
          <n-checkbox-group v-model:value="formData.inventory_type">
            <n-checkbox v-for="o in dict('tt_inventory_type')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item v-if="isOnlyUnionSlot" label="投放形式" path="union_video_type">
          <n-radio-group v-model:value="formData.union_video_type">
            <n-radio-button v-for="o in dict('tt_union_video_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.marketing_goal === 'LIVE'" label="素材类型" path="materials_type">
          <n-radio-group v-model:value="formData.materials_type">
            <n-radio-button v-for="o in dict('tt_materials_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>

        <n-divider>优化目标与出价</n-divider>
        <n-form-item label="竞价策略" path="bid_type">
          <n-radio-group v-model:value="formData.bid_type">
            <n-radio-button v-for="o in dict('tt_bid_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="优化目标" path="external_action">
          <n-radio-group v-model:value="formData.external_action">
            <n-radio-button v-for="o in dict('tt_ad_convert_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="深度优化目标" path="deep_external_action">
          <n-radio-group v-model:value="formData.deep_external_action">
            <n-radio-button v-for="o in dict('tt_deep_ad_convert_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.deep_external_action && formData.deep_external_action !== 'NONE'" label="深度出价方式" path="deep_bid_type">
          <n-radio-group v-model:value="formData.deep_bid_type">
            <n-radio-button v-for="o in dict('tt_deep_bid_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.bid_type === 'CUSTOM'" label="自定义出价方式" path="custom_bid_type">
          <n-radio-group v-model:value="formData.custom_bid_type">
            <n-radio-button v-for="o in dict('tt_custom_bid_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid v-if="formData.bid_type === 'CUSTOM'" :cols="2" :x-gap="24">
          <n-form-item-gi v-if="formData.custom_bid_type === 'CUSTOM_BID_TYPE_NORMAL'" label="出价" path="cpa_bid">
            <n-input-number v-model:value="formData.cpa_bid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <template v-if="formData.custom_bid_type === 'CUSTOM_BID_TYPE_LADDER' || formData.custom_bid_type === 'CUSTOM_BID_TYPE_RAND'">
            <n-form-item-gi label="最小出价" path="min_bid">
              <n-input-number v-model:value="formData.min_bid" :min="0" :step="0.01" style="width: 100%" />
            </n-form-item-gi>
            <n-form-item-gi label="最大出价" path="max_bid">
              <n-input-number v-model:value="formData.max_bid" :min="0" :step="0.01" style="width: 100%" />
            </n-form-item-gi>
          </template>
          <n-form-item-gi v-if="formData.deep_bid_type === 'DEEP_BID_MIN'" label="深度出价" path="deep_cpabid">
            <n-input-number v-model:value="formData.deep_cpabid" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi v-if="formData.deep_bid_type === 'ROI_COEFFICIENT'" label="深度转化ROI系数" path="roi_goal">
            <n-input-number v-model:value="formData.roi_goal" :min="0" :step="0.0001" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item label="预算择优分配" path="budget_optimize_switch">
          <n-radio-group v-model:value="formData.budget_optimize_switch">
            <n-radio-button v-for="o in dict('tt_budget_optimize_switch')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>

        <n-divider>预算</n-divider>
        <n-form-item label="项目预算类型" path="budget_mode">
          <n-radio-group v-model:value="formData.budget_mode">
            <n-radio-button v-for="o in dict('tt_budget_mode')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid v-if="formData.budget_mode === 'BUDGET_MODE_DAY'" :cols="2" :x-gap="24">
          <n-form-item-gi label="项目预算" path="project_budget">
            <n-input-number v-model:value="formData.project_budget" :min="0" :step="0.01" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item v-if="formData.delivery_mode === 'MANUAL'" label="广告预算" path="ad_budget">
          <n-input-number v-model:value="formData.ad_budget" :min="0" :step="0.01" style="width: 240px" />
        </n-form-item>

        <n-divider>投放时间</n-divider>
        <n-form-item label="投放时间" path="schedule_type">
          <n-radio-group v-model:value="formData.schedule_type">
            <n-radio-button v-for="o in dict('tt_schedule_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-grid v-if="formData.schedule_type === 'SCHEDULE_START_END'" :cols="2" :x-gap="24">
          <n-form-item-gi label="开始日期" path="start_time">
            <n-date-picker v-model:formatted-value="formData.start_time" type="date" value-format="yyyy-MM-dd" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="结束日期" path="end_time">
            <n-date-picker v-model:formatted-value="formData.end_time" type="date" value-format="yyyy-MM-dd" style="width: 100%" />
          </n-form-item-gi>
        </n-grid>
        <n-form-item label="投放时段类型" path="schedule_time_type">
          <n-radio-group v-model:value="formData.schedule_time_type">
            <n-radio-button v-for="o in dict('tt_schedule_time_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.schedule_time_type === 1" label="投放时段" path="schedule_time">
          <n-input v-model:value="formData.schedule_time" placeholder="如 00:00-08:00,12:00-14:00" style="max-width: 400px" />
        </n-form-item>

        <n-divider>产品信息</n-divider>
        <n-form-item label="产品名称" path="title">
          <n-input v-model:value="formData.title" placeholder="请输入产品名称（1-20字）" style="max-width: 400px" />
        </n-form-item>
        <n-form-item label="产品卖点">
          <n-dynamic-input v-model:value="formData.selling_points" placeholder="请输入产品卖点（6-9字）" style="max-width: 480px" />
        </n-form-item>
        <n-form-item label="行动号召">
          <n-dynamic-input v-model:value="formData.call_to_action_buttons" placeholder="请输入行动号召（2-6字）" style="max-width: 480px" />
        </n-form-item>

        <n-divider>原生广告</n-divider>
        <n-form-item label="原生锚点" path="anchor_related_type">
          <n-radio-group v-model:value="formData.anchor_related_type">
            <n-radio-button v-for="o in dict('tt_anchor_related_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <template v-if="formData.anchor_related_type === 'SELECT'">
          <n-grid :cols="2" :x-gap="24">
            <n-form-item-gi label="锚点标题" path="anchor_title">
              <n-input v-model:value="formData.anchor_title" placeholder="1-12字" />
            </n-form-item-gi>
            <n-form-item-gi label="引导文案" path="guide_text">
              <n-input v-model:value="formData.guide_text" placeholder="1-15字" />
            </n-form-item-gi>
            <n-form-item-gi label="游戏简介" path="game_description">
              <n-input v-model:value="formData.game_description" placeholder="1-45字" />
            </n-form-item-gi>
            <n-form-item-gi label="游戏特色" path="game_charatoristic">
              <n-input v-model:value="formData.game_charatoristic" placeholder="1-45字" />
            </n-form-item-gi>
            <n-form-item-gi label="游戏标签">
              <n-dynamic-input v-model:value="formData.app_tags" placeholder="1-6字，最多2个" style="max-width: 100%" />
            </n-form-item-gi>
            <n-form-item-gi label="头图素材ID" path="head_image_list">
              <n-input v-model:value="formData.head_image_list_text" type="textarea" :rows="2" placeholder="如 [1,2,3]（素材ID）" />
            </n-form-item-gi>
            <n-form-item-gi label="游戏图片ID" path="app_images">
              <n-input v-model:value="formData.app_images_text" type="textarea" :rows="2" placeholder="如 [1,2,3]（素材ID，3-8个）" />
            </n-form-item-gi>
          </n-grid>
        </template>
        <n-form-item label="主页隐藏广告视频" path="is_feed_and_fav_see">
          <n-switch v-model:value="formData.is_feed_and_fav_see" :checked-value="1" :unchecked-value="0" />
        </n-form-item>

        <n-divider v-if="formData.ad_type === 'SEARCH'">搜索广告</n-divider>
        <template v-if="formData.ad_type === 'SEARCH'">
          <n-form-item label="项目成本稳投" path="project_custom">
            <n-radio-group v-model:value="formData.project_custom">
              <n-radio-button v-for="o in dict('tt_project_custom')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
            </n-radio-group>
          </n-form-item>
          <n-form-item v-if="formData.project_custom === 'ON'" label="项目出价" path="project_cpa_bid">
            <n-input-number v-model:value="formData.project_cpa_bid" :min="0" :step="0.01" style="width: 240px" />
          </n-form-item>
          <n-form-item label="智能拓流" path="auto_extend_traffic">
            <n-radio-group v-model:value="formData.auto_extend_traffic">
              <n-radio-button v-for="o in dict('tt_auto_extend_traffic')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="文本摘要信息">
            <n-input v-model:value="formData.text_abstract_list_text" type="textarea" :rows="3" placeholder='如 [{"abstract_text":"摘要内容"}]（25-45字）' />
          </n-form-item>
          <n-form-item label="关键词列表">
            <n-input v-model:value="formData.keywords_text" type="textarea" :rows="3" placeholder='如 [{"word":"关键词","match_type":"PHRASE"}]' />
          </n-form-item>
        </template>

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
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useMessage } from 'naive-ui'
import { useDict } from '../../../composables/useDict'
import { useOptions } from '../../../composables/useOptions'
import { getTtAdTemplateDetail, createTtAdTemplate, updateTtAdTemplate } from '../../../api/mkt/tt'

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

const jsonFields = ['head_image_list', 'app_images', 'text_abstract_list', 'keywords']
const jsonTextFields = ['head_image_list_text', 'app_images_text', 'text_abstract_list_text', 'keywords_text']

const formData = reactive({
  template_name: '', app_id: [],
  landing_type: 'APP', app_promotion_type: 'DOWNLOAD', micro_promotion_type: 'WECHAT_GAME',
  marketing_goal: 'VIDEO_AND_IMAGE', ad_type: 'ALL', delivery_mode: 'MANUAL',
  download_type: 'DOWNLOAD_URL', download_mode: 'DEFAULT', mkt_ad_type: 1,
  inventory_catalog: 'MANUAL', inventory_type: [], union_video_type: 'ORIGINAL_VIDEO', materials_type: 'LIVE_MATERIALS',
  bid_type: 'CUSTOM', external_action: 'AD_CONVERT_TYPE_PAY', deep_external_action: 'NONE',
  deep_bid_type: 'NONE', custom_bid_type: 'CUSTOM_BID_TYPE_NORMAL',
  cpa_bid: 0, min_bid: 0, max_bid: 0, deep_cpabid: 0, roi_goal: 0,
  budget_optimize_switch: 'OFF',
  budget_mode: 'BUDGET_MODE_INFINITE', project_budget: 0, ad_budget: 0,
  project_custom: 'OFF', project_cpa_bid: 0,
  schedule_type: 'SCHEDULE_FROM_NOW', start_time: null, end_time: null,
  schedule_time_type: 0, schedule_time: '',
  title: '', selling_points: [], call_to_action_buttons: [],
  is_feed_and_fav_see: 0, anchor_related_type: 'OFF', anchor_title: '', app_tags: [],
  guide_text: '', game_description: '', game_charatoristic: '',
  head_image_list_text: '[]', app_images_text: '[]',
  auto_extend_traffic: 'OFF', text_abstract_list_text: '[]', keywords_text: '[]',
  promotion_type: 'LANDING_PAGE_LINK',
  source: '', yuntu_category_id: 0, brand_name_id: '',
  douyin_junior_switch: 0, guide_video_id_switch: 0, guide_video_id_fodder_id: 0,
  original_video_title: 1, material_source: 0, audience_extend: 1,
  intelligent_generation: 0, aigc_dynamic_creative_switch: 0, is_comment_switch: 0,
})

const isOnlyUnionSlot = computed(() =>
  formData.inventory_catalog === 'MANUAL' &&
  Array.isArray(formData.inventory_type) &&
  formData.inventory_type.length === 1 &&
  formData.inventory_type[0] === 'INVENTORY_UNION_SLOT'
)

const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function goBack() { router.push({ path: '/tt-ads', query: { tab: 'ad' } }) }

function toJSONText(v) { return v === undefined || v === null ? '[]' : JSON.stringify(v) }

function fill(data) {
  Object.keys(formData).forEach(k => {
    if (jsonTextFields.includes(k)) return
    if (data[k] === undefined || data[k] === null) return
    if (jsonFields.includes(k)) formData[jsonTextFields[jsonFields.indexOf(k)]] = toJSONText(data[k])
    else formData[k] = data[k]
  })
  formData.app_id = data.app_id || []
  formData.selling_points = data.selling_points || []
  formData.call_to_action_buttons = data.call_to_action_buttons || []
  formData.app_tags = data.app_tags || []
}

function parseJSON(text) {
  try { const v = JSON.parse(text); return Array.isArray(v) ? v : [] } catch { return [] }
}

async function handleSubmit() {
  try { await formRef.value?.validate() } catch { return }
  submitLoading.value = true
  try {
    const data = { ...formData }
    delete data.head_image_list_text; delete data.app_images_text
    delete data.text_abstract_list_text; delete data.keywords_text
    data.head_image_list = parseJSON(formData.head_image_list_text)
    data.app_images = parseJSON(formData.app_images_text)
    data.text_abstract_list = parseJSON(formData.text_abstract_list_text)
    data.keywords = parseJSON(formData.keywords_text)
    if (isEdit.value) { await updateTtAdTemplate(editId.value, data); message.success('更新成功') }
    else { await createTtAdTemplate(data); message.success('创建成功') }
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
      const res = await getTtAdTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  }
})
</script>
