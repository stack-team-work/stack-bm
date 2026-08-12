<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑快手定向包模板' : '新增快手定向包模板' }}</n-text>
    </n-space>

    <n-card style="max-width: 900px; margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="160">
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
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="target_type" label="定向类型">
            <n-select v-model:value="formData.target_type" :options="dict('ks_target_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="intelli_extend_option" label="智能放量">
            <n-select v-model:value="formData.intelli_extend_option" :options="dict('ks_intelli_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="gender" label="性别">
            <n-select v-model:value="formData.gender" :options="dict('ks_gender')" />
          </n-form-item-gi>
          <n-form-item-gi path="network" label="网络环境">
            <n-select v-model:value="formData.network" :options="dict('ks_network')" />
          </n-form-item-gi>
          <n-form-item-gi path="operators" label="运营商">
            <n-select v-model:value="formData.operators" :options="dict('ks_operators')" multiple clearable />
          </n-form-item-gi>
          <n-form-item-gi path="platform_os" label="操作系统">
            <n-select v-model:value="formData.platform_os" :options="dict('ks_platform_os')" />
          </n-form-item-gi>
          <n-form-item-gi path="device_price" label="设备价位">
            <n-select v-model:value="formData.device_price" :options="dict('ks_device_price')" multiple clearable />
          </n-form-item-gi>
          <n-form-item-gi path="disable_installed_app_switch" label="排除已安装">
            <n-select v-model:value="formData.disable_installed_app_switch" :options="dict('ks_installed_app')" />
          </n-form-item-gi>
          <n-form-item-gi path="filter_converted_level" label="过滤已转化">
            <n-select v-model:value="formData.filter_converted_level" :options="dict('ks_filter_converted')" />
          </n-form-item-gi>
          <n-form-item-gi path="filter_time_range" label="转化时间">
            <n-select v-model:value="formData.filter_time_range" :options="dict('ks_filter_time')" />
          </n-form-item-gi>
          <n-form-item-gi path="behavior_type" label="行为定向">
            <n-select v-model:value="formData.behavior_type" :options="dict('ks_behavior_type')" />
          </n-form-item-gi>
          <n-form-item-gi path="shared_user" label="共享用户过滤">
            <n-select v-model:value="formData.shared_user" :options="dict('ks_share_user')" />
          </n-form-item-gi>
        </n-grid>

        <n-divider>高级配置（JSON）</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="device_brand_ids" label="设备品牌" :span="2">
            <n-input v-model:value="formData.device_brand_ids" type="textarea" :rows="2" placeholder="如 [1,2,3]（设备品牌ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="android_osv" label="安卓系统版本" :span="2">
            <n-input v-model:value="formData.android_osv" type="textarea" :rows="2" placeholder="JSON 数组或 ID" />
          </n-form-item-gi>
          <n-form-item-gi path="population" label="定向人群包" :span="2">
            <n-input v-model:value="formData.population" type="textarea" :rows="2" placeholder="如 [1,2]（DMP人群包ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="exclude_population" label="排除人群包" :span="2">
            <n-input v-model:value="formData.exclude_population" type="textarea" :rows="2" placeholder="如 [1,2]" />
          </n-form-item-gi>
          <n-form-item-gi path="region" label="地域" :span="2">
            <n-input v-model:value="formData.region" type="textarea" :rows="2" placeholder="高级参数，留空即可" />
          </n-form-item-gi>
          <n-form-item-gi path="behavior" label="行为" :span="2">
            <n-input v-model:value="formData.behavior" type="textarea" :rows="2" placeholder="高级参数，留空即可" />
          </n-form-item-gi>
          <n-form-item-gi path="interest" label="兴趣" :span="2">
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
