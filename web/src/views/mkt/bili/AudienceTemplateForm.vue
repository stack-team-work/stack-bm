<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑定向包模板' : '新增定向包模板' }}</n-text>
    </n-space>

    <n-card style="margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top" label-width="150">
        <n-divider style="margin-top: 4px">基本信息</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="template_name" label="模板名称">
            <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" />
          </n-form-item-gi>
          <n-form-item-gi path="description" label="描述">
            <n-input v-model:value="formData.description" placeholder="请输入描述" />
          </n-form-item-gi>
        </n-grid>

        <n-divider>人群定向</n-divider>
        <n-form-item label="年龄段" path="age_list">
          <n-checkbox-group v-model:value="formData.age_list">
            <n-checkbox v-for="o in dict('bili_age')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="性别" path="gender_list">
          <n-checkbox-group v-model:value="formData.gender_list">
            <n-checkbox v-for="o in dict('bili_gender')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="网络环境" path="network_list">
          <n-checkbox-group v-model:value="formData.network_list">
            <n-checkbox v-for="o in dict('bili_network')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="手机价位" path="phone_price_list">
          <n-checkbox-group v-model:value="formData.phone_price_list">
            <n-checkbox v-for="o in dict('bili_phone_price')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="操作系统" path="os_list">
          <n-checkbox-group v-model:value="formData.os_list">
            <n-checkbox v-for="o in dict('bili_os')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="安装过滤" path="installed_user_filter">
          <n-checkbox-group v-model:value="formData.installed_user_filter">
            <n-checkbox v-for="o in dict('bili_installed_type')" :key="o.value" :value="o.value" :label="o.label" style="margin-right: 12px" />
          </n-checkbox-group>
        </n-form-item>
        <n-form-item label="转化用户过滤" path="converted_user_filter">
          <n-radio-group v-model:value="formData.converted_user_filter">
            <n-radio-button v-for="o in dict('bili_converted_user')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>
        <n-form-item label="地域类型" path="area_type">
          <n-radio-group v-model:value="formData.area_type">
            <n-radio-button v-for="o in dict('bili_area_type')" :key="o.value" :value="o.value">{{ o.label }}</n-radio-button>
          </n-radio-group>
        </n-form-item>

        <n-divider>高级配置（JSON）</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi label="地域列表" :span="2">
            <n-input v-model:value="formData.area_list" type="textarea" :rows="2" placeholder="如 [1,2,3]（地理划分ID）" />
          </n-form-item-gi>
          <n-form-item-gi label="地域层级" :span="2">
            <n-input v-model:value="formData.area_level_list" type="textarea" :rows="2" placeholder="如 [1,2]（发展划分ID）" />
          </n-form-item-gi>
          <n-form-item-gi label="人群包" :span="2">
            <n-input v-model:value="formData.crowd_pack" type="textarea" :rows="2" placeholder="如 [1,2]（DMP人群包ID）" />
          </n-form-item-gi>
          <n-form-item-gi label="档案内容" :span="2">
            <n-input v-model:value="formData.archive_content" type="textarea" :rows="2" placeholder="高级参数，留空即可" />
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
import { getBiliAudienceTemplateDetail, createBiliAudienceTemplate, updateBiliAudienceTemplate } from '../../../api/mkt/bili'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const { load: loadDict, options } = useDict()

const formRef = ref(null)
const submitLoading = ref(false)
const isEdit = ref(false)
const editId = ref(null)
const dict = (key) => options(key)

const jsonFields = ['area_list', 'area_level_list', 'crowd_pack', 'archive_content']

const formData = reactive({
  template_name: '', description: '',
  age_list: [], gender_list: [], network_list: [], phone_price_list: [], os_list: [], installed_user_filter: [],
  area_type: null, converted_user_filter: 0,
  area_list: '[]', area_level_list: '[]', crowd_pack: '[]', archive_content: '[]',
})
const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function toJSONText(v) { return v === undefined || v === null ? '[]' : JSON.stringify(v) }

function goBack() { router.push({ path: '/bili-ads', query: { tab: 'audience' } }) }

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
    data.age_list = data.age_list || []; data.gender_list = data.gender_list || []; data.network_list = data.network_list || []
    data.phone_price_list = data.phone_price_list || []; data.os_list = data.os_list || []; data.installed_user_filter = data.installed_user_filter || []
    if (isEdit.value) { await updateBiliAudienceTemplate(editId.value, data); message.success('更新成功') }
    else { await createBiliAudienceTemplate(data); message.success('创建成功') }
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
      const res = await getBiliAudienceTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  }
})
</script>
