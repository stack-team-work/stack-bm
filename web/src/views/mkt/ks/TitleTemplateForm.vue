<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑快手标题包模板' : '新增快手标题包模板' }}</n-text>
    </n-space>

    <n-card style="max-width: 900px; margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-form-item path="template_name" label="模板名称">
          <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" style="max-width: 600px" />
        </n-form-item>

        <n-divider style="margin: 4px 0 12px">标题物料（1-30个字符，支持动态词包）</n-divider>
        <n-card v-for="(m, i) in formData.title_materials" :key="i" size="small" style="margin-bottom: 12px">
          <template #header>
            <n-space align="center" :size="8">
              <n-text>标题 {{ i + 1 }}</n-text>
              <n-button size="tiny" type="error" quaternary @click="removeMaterial(i)">移除</n-button>
            </n-space>
          </template>
          <n-input v-model:value="m.title" placeholder="请输入标题文本（{词包} 为动态词包占位）" />
          <n-text depth="3" style="font-size: 12px; margin-top: 8px; display: block">动态词包</n-text>
          <n-dynamic-tags v-model:value="m.smart_title" :max="10" style="margin-top: 8px" />
        </n-card>
        <n-button size="small" @click="addMaterial">+ 添加标题</n-button>

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
import { getKsTitleTemplateDetail, createKsTitleTemplate, updateKsTitleTemplate } from '../../../api/mkt/ks'

const router = useRouter()
const route = useRoute()
const message = useMessage()

const formRef = ref(null)
const submitLoading = ref(false)
const isEdit = ref(false)
const editId = ref(null)

const formData = reactive({ template_name: '', title_materials: [] })

function newMaterial() { return { title: '', smart_title: [] } }

const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function goBack() { router.push({ path: '/ks-ads', query: { tab: 'title' } }) }

function addMaterial() { formData.title_materials.push(newMaterial()) }
function removeMaterial(i) { formData.title_materials.splice(i, 1) }

function fill(data) {
  formData.template_name = data.template_name || ''
  formData.title_materials = (data.title_materials || []).map(m => ({
    title: m.title || '',
    smart_title: (m.smart_title || []).map(String),
  }))
  if (!formData.title_materials.length) formData.title_materials.push(newMaterial())
}

async function handleSubmit() {
  try { await formRef.value?.validate() } catch { return }
  submitLoading.value = true
  try {
    const data = {
      template_name: formData.template_name,
      title_materials: formData.title_materials
        .filter(m => m.title.trim())
        .map(m => ({ title: m.title, smart_title: (m.smart_title || []).filter(Boolean) })),
    }
    if (isEdit.value) { await updateKsTitleTemplate(editId.value, data); message.success('更新成功') }
    else { await createKsTitleTemplate(data); message.success('创建成功') }
    goBack()
  } catch (err) { message.error(err.message || '操作失败') }
  finally { submitLoading.value = false }
}

onMounted(async () => {
  const id = route.params.id
  if (id) {
    isEdit.value = true
    editId.value = id
    try {
      const res = await getKsTitleTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  }
})
</script>
