<template>
  <div class="login-page">
    <div class="login-card">
      <div class="login-header">
        <img src="/logo.png" class="login-logo" />
        <h1 class="login-title">游戏发行后台管理系统</h1>
        <p class="login-subtitle">STACK-BM · 管理控制台</p>
      </div>
      <n-form ref="formRef" :model="formData" :rules="rules">
        <n-form-item path="username" label="用户名">
          <n-input v-model:value="formData.username" placeholder="请输入用户名" />
        </n-form-item>
        <n-form-item path="password" label="密码">
          <n-input v-model:value="formData.password" type="password" placeholder="请输入密码" @keyup.enter="handleLogin" />
        </n-form-item>
        <n-form-item path="captcha" label="验证码">
          <n-space>
            <n-input v-model:value="formData.captcha" placeholder="验证码" style="width: 140px" @keyup.enter="handleLogin" />
            <n-button size="small" @click="refreshCaptcha" :loading="captchaLoading">{{ captchaText }}</n-button>
          </n-space>
        </n-form-item>
        <n-form-item>
          <n-button type="primary" size="large" :loading="loading" block @click="handleLogin">登 录</n-button>
        </n-form-item>
      </n-form>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  position: relative;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #f0f9f4 0%, #e8f5ee 40%, #eef2fb 100%);
  overflow: hidden;
}
.login-page::before {
  content: '';
  position: absolute;
  width: 480px;
  height: 480px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(24, 160, 88, 0.18), transparent 70%);
  top: -140px;
  right: -80px;
}
.login-page::after {
  content: '';
  position: absolute;
  width: 420px;
  height: 420px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(24, 160, 88, 0.12), transparent 70%);
  bottom: -120px;
  left: -60px;
}
.login-card {
  position: relative;
  z-index: 1;
  width: 400px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 20px 60px rgba(15, 80, 45, 0.12);
  padding: 40px 36px 28px;
}
.login-header {
  text-align: center;
  margin-bottom: 28px;
}
.login-logo {
  width: 56px;
  height: 56px;
  margin-bottom: 16px;
}
.login-title {
  margin: 0 0 8px;
  font-size: 20px;
  font-weight: 600;
  color: #1a1a2e;
}
.login-subtitle {
  margin: 0;
  font-size: 13px;
  color: #999;
  letter-spacing: 1px;
}
</style>

<script setup>
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import request from '../utils/request'

const router = useRouter()
const message = useMessage()
const loading = ref(false)
const captchaLoading = ref(false)
const formRef = ref(null)
const captchaID = ref('')
const captchaText = ref('点击获取')

const formData = reactive({
  username: '',
  password: '',
  captcha: '',
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
}

async function refreshCaptcha() {
  captchaLoading.value = true
  try {
    const res = await request.post('/captcha')
    captchaID.value = res.data.captcha_id
    captchaText.value = res.data.question
  } catch {
    captchaText.value = '获取失败'
  } finally {
    captchaLoading.value = false
  }
}

async function handleLogin() {
  try { await formRef.value?.validate() } catch { return }
  loading.value = true
  try {
    const res = await request.post('/login', {
      username: formData.username,
      password: formData.password,
      captcha_id: captchaID.value,
      captcha: formData.captcha,
    })
    localStorage.setItem('token', res.data.token)
    localStorage.setItem('userInfo', JSON.stringify(res.data.admin))
    message.success('登录成功')
    router.push('/')
  } catch (err) {
    message.error(err.message || '登录失败')
    refreshCaptcha()
  } finally {
    loading.value = false
  }
}

onMounted(() => refreshCaptcha())
</script>
