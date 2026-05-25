<template>
  <div style="display: flex; justify-content: center; align-items: center; min-height: 100vh; background: #f5f7fa">
    <n-card title="游戏发行后台管理系统" style="width: 400px">
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="username" label="用户名">
          <n-input v-model:value="formData.username" placeholder="请输入用户名" />
        </n-form-item>
        <n-form-item path="password" label="密码">
          <n-input v-model:value="formData.password" type="password" placeholder="请输入密码" @keyup.enter="handleLogin" />
        </n-form-item>
        <n-form-item>
          <n-button type="primary" :loading="loading" block @click="handleLogin">登 录</n-button>
        </n-form-item>
      </n-form>
    </n-card>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { login } from '../api'

const router = useRouter()
const message = useMessage()
const loading = ref(false)
const formRef = ref(null)

const formData = reactive({
  username: '',
  password: '',
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
}

async function handleLogin() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }

  loading.value = true
  try {
    const res = await login(formData.username, formData.password)
    localStorage.setItem('token', res.data.token)
    localStorage.setItem('userInfo', JSON.stringify(res.data.admin))
    message.success('登录成功')
    router.push('/')
  } catch (err) {
    message.error(err.message || '登录失败')
  } finally {
    loading.value = false
  }
}
</script>
