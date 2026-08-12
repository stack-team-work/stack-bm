<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑快手定向包模板' : '新增快手定向包模板' }}</n-text>
    </n-space>

    <n-card style="margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top" label-width="160">
        <n-divider style="margin-top: 4px">基本信息</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="template_name" label="模板名称">
            <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="description" label="描述">
            <n-input v-model:value="formData.description" placeholder="请输入描述" />
          </n-form-item-gi>
        </n-grid>

        <n-divider>基础定向</n-divider>
        <n-form-item label="定向类型" path="target_type">
          <n-radio-group v-model:value="formData.target_type">
            <n-radio-button v-for="o in dict('ks_target_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="智能放量" path="intelli_extend_option">
          <n-radio-group v-model:value="formData.intelli_extend_option">
            <n-radio-button v-for="o in dict('ks_intelli_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="性别" path="gender">
          <n-radio-group v-model:value="formData.gender">
            <n-radio-button v-for="o in dict('ks_gender')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="网络环境" path="network">
          <n-radio-group v-model:value="formData.network">
            <n-radio-button v-for="o in dict('ks_network')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="运营商" path="operators">
          <n-checkbox-group v-model:value="formData.operators">
            <n-checkbox v-for="o in dict('ks_operators')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="操作系统" path="platform_os">
          <n-radio-group v-model:value="formData.platform_os">
            <n-radio-button v-for="o in dict('ks_platform_os')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="设备价位" path="device_price">
          <n-checkbox-group v-model:value="formData.device_price">
            <n-checkbox v-for="o in dict('ks_device_price')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="排除已安装" path="disable_installed_app_switch">
          <n-radio-group v-model:value="formData.disable_installed_app_switch">
            <n-radio-button v-for="o in dict('ks_installed_app')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="过滤已转化" path="filter_converted_level">
          <n-radio-group v-model:value="formData.filter_converted_level">
            <n-radio-button v-for="o in dict('ks_filter_converted')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item v-if="formData.filter_converted_level !== 0" label="转化时间" path="filter_time_range">
          <n-radio-group v-model:value="formData.filter_time_range">
            <n-radio-button v-for="o in dict('ks_filter_time')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="行为定向" path="behavior_type">
          <n-radio-group v-model:value="formData.behavior_type">
            <n-radio-button v-for="o in dict('ks_behavior_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="共享用户过滤" path="shared_user">
          <n-radio-group v-model:value="formData.shared_user">
            <n-radio-button v-for="o in dict('ks_share_user')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>

        <n-divider>高级配置（JSON）</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi label="设备品牌" :span="2">
            <n-input v-model:value="formData.device_brand_ids" type="textarea" :rows="2" placeholder="如 [1,2,3]（设备品牌ID）" />
          </n-form-item-gi>
          <n-form-item-gi label="安卓系统版本" :span="2">
            <n-input v-model:value="formData.android_osv" type="textarea" :rows="2" placeholder="JSON 数组或 ID" />
          </n-form-item-gi>
          <n-form-item-gi label="定向人群包" :span="2">
            <n-input v-model:value="formData.population" type="textarea" :rows="2" placeholder="如 [1,2]（DMP人群包ID）" />
          </n-form-item-gi>
          <n-form-item-gi label="排除人群包" :span="2">
            <n-input v-model:value="formData.exclude_population" type="textarea" :rows="2" placeholder="如 [1,2]" />
          </n-form-item-gi>
          <n-form-item-gi label="地域" :span="2">
            <n-input v-model:value="formData.region" type="textarea" :rows="2" placeholder="高级参数，留空即可" />
          </n-form-item-gi>
          <n-form-item-gi label="行为" :span="2">
            <n-input v-model:value="formData.behavior" type="textarea" :rows="2" placeholder="高级参数，留空即可" />
          </n-form-item-gi>
          <n-form-item-gi label="兴趣" :span="2">
            <n-input v-model:value="formData.interest" type="textarea" :rows="2" placeholder="高级参数，留空即可" />
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
import { getKsAudienceTemplateDetail, createKsAudienceTemplate, updateKsAudienceTemplate } from '../../../api/mkt/ks'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const { load: loadDict, options } = useDict()

const formRef = ref(null)
const submitLoading = ref(false)
const isEdit = ref(false)
const editId = ref(null)
const dict = (key) => options(key)

const jsonFields = ['device_brand_ids', 'android_osv', 'ios_osv', 'harmony_osv', 'population', 'exclude_population', 'seed_population', 'media', 'app_interest_ids', 'app_ids', 'region', 'behavior', 'interest', 'celebrity']

const formData = reactive({
  template_name: '', description: '',
  target_type: 0, intelli_extend_option: 0, gender: 0, network: 0, operators: [],
  disable_installed_app_switch: 0, filter_converted_level: 0, filter_time_range: 0,
  platform_os: 0, device_price: [], shared_user: 0, behavior_type: 0,
  device_brand_ids: '[]', android_osv: '[]', ios_osv: '[]', harmony_osv: '[]',
  population: '[]', exclude_population: '[]', seed_population: '[]', media: '[]',
  app_interest_ids: '[]', app_ids: '[]', region: '[]', behavior: '[]', interest: '[]', celebrity: '[]',
})
const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function toJSONText(v) { return v === undefined || v === null ? '[]' : JSON.stringify(v) }

function goBack() { router.push({ path: '/ks-ads', query: { tab: 'audience' } }) }

function fill(data) {
  Object.keys(formData).forEach(k => {
    if (data[k] === undefined || data[k] === null) return
    if (jsonFields.includes(k)) formData[k] = toJSONText(data[k])
    else formData[k] = data[k]
  })
}

async function handleSubmit() {
  try { await formRef.value?.validate() } catch { return }
  submitLoading.value = true
  try {
    const data = { ...formData }
    jsonFields.forEach(f => {
      try { const v = JSON.parse(data[f]); data[f] = Array.isArray(v) ? v : [] } catch { data[f] = [] }
    })
    data.operators = data.operators || []; data.device_price = data.device_price || []
    if (isEdit.value) { await updateKsAudienceTemplate(editId.value, data); message.success('更新成功') }
    else { await createKsAudienceTemplate(data); message.success('创建成功') }
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
      const res = await getKsAudienceTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  }
})
</script>
