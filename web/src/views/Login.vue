<template>
  <div class="login-page">
    <div class="login-panel">
      <div class="login-brand">
        <img src="/logo.png" class="brand-logo" />
        <h1 class="brand-title">游戏发行后台管理系统</h1>
        <p class="brand-subtitle">STACK-BM · 一站式投放管理控制台</p>
        <ul class="brand-features">
          <li>多渠道广告模板统一配置</li>
          <li>游戏 / 礼包 / 支付全链路管理</li>
          <li>实时数据报表与日志追踪</li>
        </ul>
      </div>
      <div class="login-form-panel">
        <div class="login-card">
          <h2 class="login-form-title">账号登录</h2>
          <p class="login-form-subtitle">使用后台账号登录管理控制台</p>
          <n-form ref="formRef" :model="formData" :rules="rules" label-placement="top" size="large">
            <n-form-item path="username" label="用户名">
              <n-input v-model:value="formData.username" placeholder="请输入用户名" />
            </n-form-item>
            <n-form-item path="password" label="密码">
              <n-input v-model:value="formData.password" type="password" placeholder="请输入密码" show-password-on="click" @keyup.enter="handleLogin" />
            </n-form-item>
            <n-form-item path="captcha" label="验证码">
              <div class="captcha-row">
                <n-input v-model:value="formData.captcha" placeholder="验证码" @keyup.enter="handleLogin" />
                <n-button class="captcha-btn" :loading="captchaLoading" @click="refreshCaptcha">{{ captchaText }}</n-button>
              </div>
            </n-form-item>
            <n-form-item>
              <n-button type="primary" size="large" :loading="loading" block class="login-btn" @click="handleLogin">登 录</n-button>
            </n-form-item>
          </n-form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  background: linear-gradient(135deg, #ecfdf5 0%, #f0f4fb 60%, #eef2fb 100%);
  overflow: hidden;
}
.login-panel {
  width: 900px;
  max-width: 100%;
  display: flex;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 24px 64px rgba(15, 80, 45, 0.16);
  background: #fff;
}
.login-brand {
  flex: 1;
  padding: 48px 40px;
  color: #fff;
  background: linear-gradient(160deg, #059669 0%, #047857 55%, #064e3b 100%);
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
  overflow: hidden;
}
.login-brand::before {
  content: '';
  position: absolute;
  width: 320px;
  height: 320px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.14), transparent 70%);
  top: -120px;
  right: -100px;
}
.login-brand::after {
  content: '';
  position: absolute;
  width: 260px;
  height: 260px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.08), transparent 70%);
  bottom: -90px;
  left: -70px;
}
.brand-logo {
  width: 56px;
  height: 56px;
  margin-bottom: 20px;
  position: relative;
  z-index: 1;
}
.brand-title {
  margin: 0 0 8px;
  font-size: 24px;
  font-weight: 600;
  position: relative;
  z-index: 1;
}
.brand-subtitle {
  margin: 0 0 32px;
  font-size: 13px;
  opacity: 0.85;
  letter-spacing: 1px;
  position: relative;
  z-index: 1;
}
.brand-features {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 14px;
  position: relative;
  z-index: 1;
}
.brand-features li {
  font-size: 14px;
  opacity: 0.92;
  padding-left: 24px;
  position: relative;
}
.brand-features li::before {
  content: '✓';
  position: absolute;
  left: 0;
  top: 0;
  color: #a7f3d0;
  font-weight: 600;
}
.login-form-panel {
  width: 420px;
  display: flex;
  align-items: center;
  padding: 40px 44px;
}
.login-card {
  width: 100%;
}
.login-form-title {
  margin: 0 0 4px;
  font-size: 22px;
  font-weight: 600;
  color: var(--text-1);
}
.login-form-subtitle {
  margin: 0 0 28px;
  font-size: 13px;
  color: var(--text-3);
}
.captcha-row {
  display: flex;
  gap: 12px;
  width: 100%;
}
.captcha-btn {
  flex-shrink: 0;
  min-width: 120px;
}
.login-btn {
  margin-top: 8px;
}
@media (max-width: 720px) {
  .login-brand {
    display: none;
  }
  .login-form-panel {
    width: 100%;
  }
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
