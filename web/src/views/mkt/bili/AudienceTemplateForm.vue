<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑定向包模板' : '新增定向包模板' }}</n-text>
    </n-space>

    <n-card style="max-width: 900px; margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="150">
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
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="age_list" label="年龄段">
            <n-select v-model:value="formData.age_list" :options="dict('bili_age')" multiple placeholder="选择年龄段" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="gender_list" label="性别">
            <n-select v-model:value="formData.gender_list" :options="dict('bili_gender')" multiple placeholder="选择性别" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="network_list" label="网络环境">
            <n-select v-model:value="formData.network_list" :options="dict('bili_network')" multiple placeholder="选择网络" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="phone_price_list" label="手机价位">
            <n-select v-model:value="formData.phone_price_list" :options="dict('bili_phone_price')" multiple placeholder="选择手机价位" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="os_list" label="操作系统">
            <n-select v-model:value="formData.os_list" :options="dict('bili_os')" multiple placeholder="选择系统" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="installed_user_filter" label="安装过滤">
            <n-select v-model:value="formData.installed_user_filter" :options="dict('bili_installed_type')" multiple placeholder="选择安装过滤" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="converted_user_filter" label="转化用户过滤">
            <n-select v-model:value="formData.converted_user_filter" :options="dict('bili_converted_user')" clearable />
          </n-form-item-gi>
          <n-form-item-gi path="area_type" label="地域类型">
            <n-select v-model:value="formData.area_type" :options="dict('bili_area_type')" clearable />
          </n-form-item-gi>
        </n-grid>

        <n-divider>高级配置（JSON）</n-divider>
        <n-grid :cols="2" :x-gap="24">
          <n-form-item-gi path="area_list" label="地域列表" :span="2">
            <n-input v-model:value="formData.area_list" type="textarea" :rows="2" placeholder="如 [1,2,3]（地理划分ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="area_level_list" label="地域层级" :span="2">
            <n-input v-model:value="formData.area_level_list" type="textarea" :rows="2" placeholder="如 [1,2]（发展划分ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="crowd_pack" label="人群包" :span="2">
            <n-input v-model:value="formData.crowd_pack" type="textarea" :rows="2" placeholder="如 [1,2]（DMP人群包ID）" />
          </n-form-item-gi>
          <n-form-item-gi path="archive_content" label="档案内容" :span="2">
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
