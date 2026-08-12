<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑定向包模板' : '新增定向包模板' }}</n-text>
    </n-space>

    <n-card style="max-width: 1000px; margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top">
        <n-divider style="margin-top: 4px">基本信息</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="template_name" label="模板名称">
            <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="description" label="描述">
            <n-input v-model:value="formData.description" placeholder="请输入描述（1-50字）" />
          </n-form-item-gi>
        </n-grid>

        <n-divider>定向设置</n-divider>
        <n-form-item label="定向包类型" path="landing_type">
          <n-radio-group v-model:value="formData.landing_type">
            <n-radio-button v-for="o in dict('tt_audience_landing_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="地域定向类型" path="district">
          <n-radio-group v-model:value="formData.district">
            <n-radio-button v-for="o in dict('tt_district')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <template v-if="formData.district !== 'NONE'">
          <n-form-item label="位置类型" path="location_type">
            <n-radio-group v-model:value="formData.location_type">
              <n-radio-button v-for="o in dict('tt_location_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="城市ID列表">
            <n-input v-model:value="formData.city_text" type="textarea" :rows="2" placeholder="如 [110000,310000]（城市ID）" />
          </n-form-item>
        </template>

        <n-form-item label="性别" path="gender">
          <n-radio-group v-model:value="formData.gender">
            <n-radio-button v-for="o in dict('tt_gender')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="年龄" path="age">
          <n-checkbox-group v-model:value="formData.age">
            <n-checkbox v-for="o in dict('tt_age')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="职业" path="career">
          <n-checkbox-group v-model:value="formData.career">
            <n-checkbox v-for="o in dict('tt_career')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="投放平台" path="platform">
          <n-checkbox-group v-model:value="formData.platform">
            <n-checkbox v-for="o in dict('tt_platform_name')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="设备类型" path="device_type">
          <n-checkbox-group v-model:value="formData.device_type">
            <n-checkbox v-for="o in dict('tt_device_type')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="网络类型" path="ac">
          <n-checkbox-group v-model:value="formData.ac">
            <n-checkbox v-for="o in dict('tt_ac')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="运营商" path="carrier">
          <n-checkbox-group v-model:value="formData.carrier">
            <n-checkbox v-for="o in dict('tt_carrier')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="手机品牌" path="device_brand">
          <n-checkbox-group v-model:value="formData.device_brand">
            <n-checkbox v-for="o in dict('tt_device_brand')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-grid :cols="3" :x-gap="24">
          <n-form-item-gi label="最低安卓版本" path="android_osv">
            <n-select v-model:value="formData.android_osv" :options="dict('tt_android_osv')" clearable />
          </n-form-item-gi>
          <n-form-item-gi label="最低iOS版本" path="ios_osv">
            <n-select v-model:value="formData.ios_osv" :options="dict('tt_ios_osv')" clearable />
          </n-form-item-gi>
          <n-form-item-gi label="最低鸿蒙版本" path="harmony_osv">
            <n-select v-model:value="formData.harmony_osv" :options="dict('tt_harmony_osv')" clearable />
          </n-form-item-gi>
        </n-grid>

        <n-divider>行为兴趣</n-divider>
        <n-form-item label="行为兴趣模式" path="interest_action_mode">
          <n-radio-group v-model:value="formData.interest_action_mode">
            <n-radio-button v-for="o in dict('tt_interest_action_mode')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <template v-if="formData.interest_action_mode === 'CUSTOM'">
          <n-form-item label="行为场景" path="action_scene">
            <n-checkbox-group v-model:value="formData.action_scene">
              <n-checkbox v-for="o in dict('tt_action_scene')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
            </n-checkbox-group>
          </n-form-item>
          <n-form-item label="行为天数" path="action_days">
            <n-radio-group v-model:value="formData.action_days">
              <n-radio-button v-for="o in dict('tt_action_days')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
            </n-radio-group>
          </n-form-item>
          <n-form-item label="行为类目词">
            <n-input v-model:value="formData.action_categories_text" type="textarea" :rows="2" placeholder='如 [{"id":123,"name":"类目名"}]' />
          </n-form-item>
          <n-form-item label="行为关键词">
            <n-input v-model:value="formData.action_words_text" type="textarea" :rows="2" placeholder='如 [{"id":123,"name":"关键词"}]' />
          </n-form-item>
          <n-form-item label="兴趣类目词">
            <n-input v-model:value="formData.interest_categories_text" type="textarea" :rows="2" placeholder='如 [{"id":123,"name":"类目名"}]' />
          </n-form-item>
          <n-form-item label="兴趣关键词">
            <n-input v-model:value="formData.interest_words_text" type="textarea" :rows="2" placeholder='如 [{"id":123,"name":"关键词"}]' />
          </n-form-item>
        </template>

        <n-divider>抖音达人</n-divider>
        <n-form-item label="达人互动行为" path="aweme_fan_behaviors">
          <n-checkbox-group v-model:value="formData.aweme_fan_behaviors">
            <n-checkbox v-for="o in dict('tt_aweme_fan_behaviors')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="互动行为时间范围" path="aweme_fan_time_scope">
          <n-radio-group v-model:value="formData.aweme_fan_time_scope">
            <n-radio-button v-for="o in dict('tt_aweme_fan_time_scope')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="达人分类ID列表">
          <n-input v-model:value="formData.aweme_fan_categories_text" type="textarea" :rows="2" placeholder="如 [1,2,3]（达人分类ID）" />
        </n-form-item>
        <n-form-item label="达人ID列表">
          <n-input v-model:value="formData.aweme_fan_accounts_text" type="textarea" :rows="2" placeholder='如 [{"user_id":"123","name":"达人"}]' />
        </n-form-item>

        <n-divider>过滤与放量</n-divider>
        <n-form-item label="媒体定向" path="superior_popularity_type">
          <n-radio-group v-model:value="formData.superior_popularity_type">
            <n-radio-button v-for="o in dict('tt_superior_popularity_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="过滤已安装" path="hide_if_exists">
          <n-radio-group v-model:value="formData.hide_if_exists">
            <n-radio-button v-for="o in dict('tt_hide_if_exists')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="过滤已转化用户" path="hide_if_converted">
          <n-radio-group v-model:value="formData.hide_if_converted">
            <n-radio-button v-for="o in dict('tt_hide_if_converted')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.hide_if_converted && formData.hide_if_converted !== 'NO_EXCLUDE'" label="过滤时间范围" path="converted_time_duration">
          <n-radio-group v-model:value="formData.converted_time_duration">
            <n-radio-button v-for="o in dict('tt_convert_time_duration')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="过滤高活跃用户" path="filter_aweme_abnormal_active">
          <n-radio-group v-model:value="formData.filter_aweme_abnormal_active">
            <n-radio-button v-for="o in dict('tt_filter_aweme_abnormal')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="过滤自己的粉丝" path="filter_own_aweme_fans">
          <n-radio-group v-model:value="formData.filter_own_aweme_fans">
            <n-radio-button v-for="o in dict('tt_filter_own_aweme_fans')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="过滤高关注数用户" path="filter_aweme_fans_count">
          <n-radio-group v-model:value="formData.filter_aweme_fans_count">
            <n-radio-button v-for="o in dict('tt_filter_aweme_fans_count')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="新用户使用头条时间" path="activate_type">
          <n-checkbox-group v-model:value="formData.activate_type">
            <n-checkbox v-for="o in dict('tt_activate_type')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="智能放量" path="auto_extend_enabled">
          <n-radio-group v-model:value="formData.auto_extend_enabled">
            <n-radio-button v-for="o in dict('tt_auto_extend_enabled')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.auto_extend_enabled === 'ON'" label="可放开定向" path="auto_extend_targets">
          <n-checkbox-group v-model:value="formData.auto_extend_targets">
            <n-checkbox v-for="o in dict('tt_auto_extend_target')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>

        <n-divider>高级配置（JSON）</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi label="主体ID" path="subject_id">
            <n-input-number v-model:value="formData.subject_id" :min="0" style="width: 100%" />
          </n-form-item-gi>
          <n-form-item-gi label="定向人群包ID">
            <n-input v-model:value="formData.retargeting_tags_include_text" type="textarea" :rows="2" placeholder='如 ["tag1","tag2"]' />
          </n-form-item-gi>
          <n-form-item-gi label="排除人群包ID">
            <n-input v-model:value="formData.retargeting_tags_exclude_text" type="textarea" :rows="2" placeholder='如 ["tag1","tag2"]' />
          </n-form-item-gi>
          <n-form-item-gi label="定向关键词" path="retargeting_keywords">
            <n-input v-model:value="formData.retargeting_keywords" placeholder="请输入关键词" />
          </n-form-item-gi>
          <n-form-item-gi label="定向逻辑媒体包">
            <n-input v-model:value="formData.flow_package_text" type="textarea" :rows="2" placeholder="如 [1,2,3]" />
          </n-form-item-gi>
          <n-form-item-gi label="排除媒体包">
            <n-input v-model:value="formData.exclude_flow_package_text" type="textarea" :rows="2" placeholder="如 [1,2,3]" />
          </n-form-item-gi>
          <n-form-item-gi label="手机价格">
            <n-input v-model:value="formData.launch_price_text" type="textarea" :rows="2" placeholder="如 [1,2,3]（价位ID）" />
          </n-form-item-gi>
        </n-grid>

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
import { getTtAudienceTemplateDetail, createTtAudienceTemplate, updateTtAudienceTemplate } from '../../../api/mkt/tt'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const { load: loadDict, options } = useDict()

const formRef = ref(null)
const submitLoading = ref(false)
const isEdit = ref(false)
const editId = ref(null)
const dict = (key) => options(key)

const jsonFields = [
  'action_categories', 'action_words', 'interest_categories', 'interest_words',
  'aweme_fan_categories', 'aweme_fan_accounts', 'retargeting_tags_include',
  'retargeting_tags_exclude', 'flow_package', 'exclude_flow_package', 'launch_price', 'city',
]
const jsonTextFields = [
  'action_categories_text', 'action_words_text', 'interest_categories_text', 'interest_words_text',
  'aweme_fan_categories_text', 'aweme_fan_accounts_text', 'retargeting_tags_include_text',
  'retargeting_tags_exclude_text', 'flow_package_text', 'exclude_flow_package_text', 'launch_price_text', 'city_text',
]

const formData = reactive({
  template_name: '', description: '',
  landing_type: 'APP_ANDROID', district: 'NONE', location_type: 'ALL', city_text: '[]',
  gender: 'NONE', age: [], career: [], platform: [], device_type: [], ac: [], carrier: [], device_brand: [],
  android_osv: null, ios_osv: null, harmony_osv: null,
  subject_id: 0,
  retargeting_tags_include_text: '[]', retargeting_tags_exclude_text: '[]', retargeting_keywords: '',
  interest_action_mode: 'UNLIMITED', action_scene: [], action_days: 7,
  action_categories_text: '[]', action_words_text: '[]', interest_categories_text: '[]', interest_words_text: '[]',
  aweme_fan_behaviors: [], aweme_fan_time_scope: null,
  aweme_fan_categories_text: '[]', aweme_fan_accounts_text: '[]',
  superior_popularity_type: 'NONE',
  flow_package_text: '[]', exclude_flow_package_text: '[]',
  hide_if_exists: 'UNLIMITED', hide_if_converted: 'NO_EXCLUDE', converted_time_duration: 'NONE',
  filter_aweme_abnormal_active: 'NONE', filter_own_aweme_fans: 'NONE', filter_aweme_fans_count: 0,
  activate_type: [], launch_price_text: '[]',
  auto_extend_enabled: 'OFF', auto_extend_targets: [],
})

const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function goBack() { router.push({ path: '/tt-ads', query: { tab: 'audience' } }) }

function toJSONText(v) { return v === undefined || v === null ? '[]' : JSON.stringify(v) }

function fill(data) {
  Object.keys(formData).forEach(k => {
    if (jsonTextFields.includes(k)) return
    if (data[k] === undefined || data[k] === null) return
    if (jsonFields.includes(k)) formData[jsonTextFields[jsonFields.indexOf(k)]] = toJSONText(data[k])
    else formData[k] = data[k]
  })
}

function parseJSON(text) {
  try { const v = JSON.parse(text); return Array.isArray(v) ? v : [] } catch { return [] }
}

async function handleSubmit() {
  try { await formRef.value?.validate() } catch { return }
  submitLoading.value = true
  try {
    const data = { ...formData }
    jsonTextFields.forEach(t => delete data[t])
    jsonFields.forEach((f, i) => { data[f] = parseJSON(formData[jsonTextFields[i]]) })
    if (isEdit.value) { await updateTtAudienceTemplate(editId.value, data); message.success('更新成功') }
    else { await createTtAudienceTemplate(data); message.success('创建成功') }
    goBack()
  } catch (err) { message.error(err.message || '操作失败') }
  finally { submitLoading.value = false }
}

onMounted(async () => {
  await loadDict()
  const id = route.params.id
  if (id) {
    isEdit.value = true
    editId.value = id
    try {
      const res = await getTtAudienceTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  }
})
</script>
