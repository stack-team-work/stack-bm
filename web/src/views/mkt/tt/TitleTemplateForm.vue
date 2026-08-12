<template>
  <div>
    <n-space align="center" :size="12">
      <n-text style="font-size: 18px; font-weight: 600">{{ isEdit ? '编辑标题包模板' : '新增标题包模板' }}</n-text>
    </n-space>

    <n-card style="max-width: 1000px; margin-top: 16px">
      <n-form ref="formRef" :model="formData" :rules="rules" label-placement="left" label-width="100">
        <n-form-item path="template_name" label="模板名称">
          <n-input v-model:value="formData.template_name" placeholder="请输入模板名称" style="max-width: 600px" />
        </n-form-item>

        <n-divider style="margin: 4px 0 12px">标题物料（5-30字，支持 {词包名} 智能词包）</n-divider>
        <n-card v-for="(m, i) in formData.title_materials" :key="i" size="small" style="margin-bottom: 12px">
          <template #header>
            <n-space align="center" :size="8">
              <n-text>标题 {{ i + 1 }}</n-text>
              <n-button size="tiny" type="error" quaternary @click="removeMaterial(i)">移除</n-button>
            </n-space>
          </template>
          <n-space vertical :size="8">
            <n-input v-model:value="m.title" placeholder="请输入标题文本（{词包} 为智能词包占位）" />
            <n-space>
              <n-button size="tiny" @click="openWordDialog(i)">插入词包</n-button>
              <n-tag v-for="(w, j) in m.insertedWords" :key="j" size="small" closable @close="removeWord(m, j)">
                {{ w }}
              </n-tag>
            </n-space>
          </n-space>
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

    <n-modal v-model:show="wordModal.show" preset="card" title="选择词包" style="width: 480px">
      <n-data-table :columns="wordColumns" :data="wordList" :pagination="false" :max-height="360" :scroll-x="440" />
    </n-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, h } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { NButton, useMessage } from 'naive-ui'
import { getTtTitleTemplateDetail, createTtTitleTemplate, updateTtTitleTemplate, getTtWordList } from '../../../api/mkt/tt'

const router = useRouter()
const route = useRoute()
const message = useMessage()

const formRef = ref(null)
const submitLoading = ref(false)
const isEdit = ref(false)
const editId = ref(null)

const wordModal = reactive({ show: false, materialIndex: -1 })
const wordList = ref([])

const wordColumns = [
  { title: '词包名', key: 'name' },
  { title: '最大字数', key: 'max_word_len', width: 90 },
  {
    title: '操作', key: 'actions', width: 90,
    render: (row) => h(NButton, { size: 'tiny', type: 'primary', onClick: () => insertWord(row.name) }, { default: () => '插入' }),
  },
]

const formData = reactive({ template_name: '', title_materials: [] })

function newMaterial() { return { title: '', insertedWords: [] } }

const rules = { template_name: [{ required: true, message: '请输入模板名称', trigger: 'blur' }] }

function goBack() { router.push({ path: '/tt-ads', query: { tab: 'title' } }) }

function addMaterial() { formData.title_materials.push(newMaterial()) }
function removeMaterial(i) { formData.title_materials.splice(i, 1) }
function removeWord(m, j) {
  const w = m.insertedWords[j]
  m.insertedWords.splice(j, 1)
  m.title = m.title.replace(`{${w}}`, '')
}

async function openWordDialog(i) {
  wordModal.materialIndex = i
  wordModal.show = true
  if (!wordList.value.length) {
    try {
      const res = await getTtWordList()
      wordList.value = res.data || []
    } catch { message.error('加载词包失败') }
  }
}

function insertWord(name) {
  const m = formData.title_materials[wordModal.materialIndex]
  if (!m) return
  m.title = (m.title || '') + `{${name}}`
  if (!m.insertedWords.includes(name)) m.insertedWords.push(name)
  wordModal.show = false
}

function fill(data) {
  formData.template_name = data.template_name || ''
  formData.title_materials = (data.title_materials || []).map(m => ({
    title: m.title || '',
    insertedWords: (m.title.match(/\{([^\{\}]+)\}/g) || []).map(w => w.replace(/[{}]/g, '')),
  }))
  if (!formData.title_materials.length) formData.title_materials.push(newMaterial())
}

async function handleSubmit() {
  try { await formRef.value?.validate() } catch { return }
  submitLoading.value = true
  try {
    const data = {
      template_name: formData.template_name,
      title_materials: formData.title_materials.filter(m => m.title.trim()).map(m => ({ title: m.title.trim() })),
    }
    if (isEdit.value) { await updateTtTitleTemplate(editId.value, data); message.success('更新成功') }
    else { await createTtTitleTemplate(data); message.success('创建成功') }
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
      const res = await getTtTitleTemplateDetail(id)
      if (res.data) fill(res.data)
    } catch { message.error('加载模板失败') }
  } else {
    formData.title_materials.push(newMaterial())
  }
})
</script>
