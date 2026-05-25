<template>
  <div>
    <n-card :title="isEdit ? '编辑游戏应用' : '新增游戏应用'" :bordered="false">
      <template #header-extra>
        <n-button @click="goBack">返回列表</n-button>
      </template>

      <n-form ref="formRef" :model="formData" :rules="rules" label-width="120" style="max-width: 800px">
        <n-space vertical>
          <n-space>
            <n-form-item path="pid" label="所属游戏" style="width: 280px">
              <n-select v-model:value="formData.pid" :options="gameOptions" placeholder="请选择父游戏" />
            </n-form-item>
            <n-form-item path="name" label="应用名称" style="width: 280px">
              <n-input v-model:value="formData.name" placeholder="请输入应用名称" />
            </n-form-item>
          </n-space>
          <n-space>
            <n-form-item path="package_name" label="包名" style="width: 280px">
              <n-input v-model:value="formData.package_name" placeholder="请输入包名" :disabled="isEdit" />
            </n-form-item>
            <n-form-item path="app_name" label="App名称" style="width: 280px">
              <n-input v-model:value="formData.app_name" placeholder="请输入App名称" />
            </n-form-item>
          </n-space>
          <n-space>
            <n-form-item path="os" label="操作系统" style="width: 280px">
              <n-select v-model:value="formData.os" :options="osOptions" />
            </n-form-item>
            <n-form-item path="sdk_ver" label="SDK版本" style="width: 280px">
              <n-input v-model:value="formData.sdk_ver" placeholder="请输入SDK版本" />
            </n-form-item>
          </n-space>
          <n-space>
            <n-form-item path="app_ver" label="应用版本" style="width: 280px">
              <n-input v-model:value="formData.app_ver" placeholder="请输入应用版本" />
            </n-form-item>
            <n-form-item path="age" label="年龄限制" style="width: 280px">
              <n-input-number v-model:value="formData.age" :min="0" style="width: 100%" />
            </n-form-item>
          </n-space>
          <n-space>
            <n-form-item path="callback_url" label="回调地址" style="width: 280px">
              <n-input v-model:value="formData.callback_url" placeholder="请输入回调地址" />
            </n-form-item>
            <n-form-item path="api_domain" label="API域名" style="width: 280px">
              <n-input v-model:value="formData.api_domain" placeholder="请输入API域名" />
            </n-form-item>
          </n-space>
          <n-form-item path="pay_domain" label="支付域名" style="width: 280px">
            <n-input v-model:value="formData.pay_domain" placeholder="请输入支付域名" />
          </n-form-item>
        </n-space>

        <n-divider>功能开关</n-divider>
        <n-space>
          <n-form-item path="is_verify" label="实名认证" label-placement="left">
            <n-switch v-model:value="formData.is_verify" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item path="is_open_charge" label="开启充值" label-placement="left">
            <n-switch v-model:value="formData.is_open_charge" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item path="is_open_register" label="开启注册" label-placement="left">
            <n-switch v-model:value="formData.is_open_register" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item path="is_alert_email" label="邮件提醒" label-placement="left">
            <n-switch v-model:value="formData.is_alert_email" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
        </n-space>
        <n-space>
          <n-form-item path="is_alert_phone" label="短信提醒" label-placement="left">
            <n-switch v-model:value="formData.is_alert_phone" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item path="is_alert_auth" label="认证提醒" label-placement="left">
            <n-switch v-model:value="formData.is_alert_auth" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
          <n-form-item path="is_open_float" label="悬浮窗" label-placement="left">
            <n-switch v-model:value="formData.is_open_float" :checked-value="1" :unchecked-value="0" />
          </n-form-item>
        </n-space>

        <n-divider>扩展参数</n-divider>
        <n-form-item path="cs_params" label="客服参数">
          <n-input v-model:value="formData.cs_params" type="textarea" placeholder="请输入客服参数(JSON)" />
        </n-form-item>
        <n-form-item path="pay_params" label="支付参数">
          <n-input v-model:value="formData.pay_params" type="textarea" placeholder="请输入支付参数(JSON)" />
        </n-form-item>
        <n-form-item path="h5_params" label="H5参数">
          <n-input v-model:value="formData.h5_params" type="textarea" placeholder="请输入H5参数(JSON)" />
        </n-form-item>

        <n-form-item path="status" label="状态">
          <n-switch v-model:value="formData.status" :checked-value="1" :unchecked-value="0" checked-text="启用" unchecked-text="禁用" />
        </n-form-item>

        <div v-if="isEdit && appKey" style="margin-bottom: 16px; padding: 12px; background: #f5f7fa; border-radius: 4px">
          <n-text depth="3">AppKey: </n-text><n-text code>{{ appKey }}</n-text>
          <n-text depth="3" style="margin-left: 24px">AppSecret: </n-text><n-text code>{{ appSecret }}</n-text>
          <n-text depth="3" style="margin-left: 12px; font-size: 12px">(自动生成，不可修改)</n-text>
        </div>

        <n-button type="primary" :loading="submitLoading" @click="handleSubmit" style="margin-top: 12px">
          {{ isEdit ? '保存修改' : '确认创建' }}
        </n-button>
      </n-form>
    </n-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useMessage } from 'naive-ui'
import { getGameAll, createGameApp, updateGameApp } from '../../api/game'

const router = useRouter()
const route = useRoute()
const message = useMessage()
const formRef = ref(null)
const submitLoading = ref(false)
const gameOptions = ref([])
const appKey = ref('')
const appSecret = ref('')

const isEdit = ref(false)
const editId = ref(null)

const osOptions = [
  { label: 'Android', value: 1 },
  { label: 'iOS', value: 2 },
  { label: 'H5', value: 3 },
]

const formData = reactive({
  pid: null,
  name: '',
  package_name: '',
  app_name: '',
  os: 1,
  sdk_ver: '',
  app_ver: '',
  age: 18,
  callback_url: '',
  api_domain: '',
  pay_domain: '',
  is_verify: 0,
  is_open_charge: 1,
  is_open_register: 1,
  is_alert_email: 1,
  is_alert_phone: 1,
  is_alert_auth: 0,
  is_open_float: 1,
  cs_params: '',
  pay_params: '',
  h5_params: '',
  status: 1,
})

const rules = {
  name: [{ required: true, message: '请输入应用名称', trigger: 'blur' }],
  pid: [{ required: true, type: 'number', min: 1, message: '请选择所属游戏', trigger: 'change' }],
}

function goBack() {
  router.push('/game-app')
}

async function loadGames() {
  try {
    const res = await getGameAll()
    gameOptions.value = (res.data || []).map(g => ({ label: g.name, value: g.id }))
  } catch { /* ignore */ }
}

function fillForm(app) {
  formData.pid = app.pid
  formData.name = app.name
  formData.package_name = app.package_name
  formData.app_name = app.app_name || ''
  formData.os = app.os || 1
  formData.sdk_ver = app.sdk_ver || ''
  formData.app_ver = app.app_ver || ''
  formData.age = app.age || 18
  formData.callback_url = app.callback_url || ''
  formData.api_domain = app.api_domain || ''
  formData.pay_domain = app.pay_domain || ''
  formData.is_verify = app.is_verify ?? 0
  formData.is_open_charge = app.is_open_charge ?? 1
  formData.is_open_register = app.is_open_register ?? 1
  formData.is_alert_email = app.is_alert_email ?? 1
  formData.is_alert_phone = app.is_alert_phone ?? 1
  formData.is_alert_auth = app.is_alert_auth ?? 0
  formData.is_open_float = app.is_open_float ?? 1
  formData.cs_params = app.cs_params || ''
  formData.pay_params = app.pay_params || ''
  formData.h5_params = app.h5_params || ''
  formData.status = app.status ?? 1
  appKey.value = app.app_key || ''
  appSecret.value = app.app_secret || ''
}

async function handleSubmit() {
  try {
    await formRef.value?.validate()
  } catch { return }

  submitLoading.value = true
  try {
    if (isEdit.value) {
      await updateGameApp(editId.value, { ...formData })
      message.success('更新成功')
    } else {
      const res = await createGameApp({ ...formData })
      message.success('创建成功')
      appKey.value = res.data?.app_key || ''
      appSecret.value = res.data?.app_secret || ''
    }
    if (!isEdit.value) {
      setTimeout(() => router.push('/game-app'), 500)
    }
  } catch (err) {
    message.error(err.message || '操作失败')
  } finally {
    submitLoading.value = false
  }
}

onMounted(async () => {
  await loadGames()
  const id = route.params.id
  if (id && id !== 'create') {
    isEdit.value = true
    editId.value = Number(id)
    const res = await fetch(`/api/game-app/detail/${id}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/x-www-form-urlencoded', 'Authorization': `Bearer ${localStorage.getItem('token')}` },
    })
    if (res.ok) {
      const data = await res.json()
      if (data.data) fillForm(data.data)
    }
  }
})
</script>
