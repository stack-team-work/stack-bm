<template>
  <n-layout has-sider style="height: 100vh">
    <n-layout-sider
      bordered
      collapse-mode="width"
      :collapsed-width="64"
      :width="220"
      :collapsed="collapsed"
      :inverted="true"
      show-trigger
      @collapse="collapsed = true"
      @expand="collapsed = false"
    >
      <div style="height: 56px; display: flex; align-items: center; justify-content: center; border-bottom: 1px solid rgba(255,255,255,0.08)">
        <span v-if="!collapsed" style="color: #fff; font-size: 16px; font-weight: bold; letter-spacing: 2px">Stack-BM</span>
        <span v-else style="color: #fff; font-size: 14px; font-weight: bold">BM</span>
      </div>
      <n-menu
        :collapsed="collapsed"
        :collapsed-width="64"
        :collapsed-icon-size="22"
        :value="currentRoute"
        :options="menuOptions"
        :inverted="true"
        @update:value="handleMenuSelect"
      />
    </n-layout-sider>
    <n-layout>
      <n-layout-header bordered style="height: 56px; padding: 0 24px; display: flex; align-items: center; justify-content: space-between; background: #fff">
        <span style="font-size: 16px; font-weight: 500">游戏发行后台管理系统</span>
        <n-space align="center">
          <n-text>{{ userInfo?.username || '' }}</n-text>
          <n-button size="small" @click="handleLogout">退出登录</n-button>
        </n-space>
      </n-layout-header>
      <n-layout-content style="padding: 24px; background: #fff; min-height: calc(100vh - 56px)">
        <router-view />
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<script setup>
import { ref, computed, h } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { NIcon } from 'naive-ui'
import { HomeOutline, SettingsOutline, GameControllerOutline } from '@vicons/ionicons5'

const router = useRouter()
const route = useRoute()
const collapsed = ref(false)

const userInfo = JSON.parse(localStorage.getItem('userInfo') || 'null')

const renderIcon = (icon) => () => h(NIcon, null, { default: () => h(icon) })

const menuOptions = [
  {
    label: '概况',
    key: '/dashboard',
    icon: renderIcon(HomeOutline),
  },
  {
    label: '系统管理',
    key: 'system',
    icon: renderIcon(SettingsOutline),
    children: [
      { label: '后台账号', key: '/admin' },
      { label: '后台角色', key: '/admin-group' },
      { label: '菜单管理', key: '/menu' },
      { label: '操作日志', key: '/logs' },
    ],
  },
  {
    label: '游戏管理',
    key: 'game-mgmt',
    icon: renderIcon(GameControllerOutline),
    children: [
      { label: '父游戏', key: '/game' },
      { label: '子游戏', key: '/game-app' },
      { label: '游戏CP', key: '/game-cp' },
    ],
  },
]

const currentRoute = computed(() => route.path)

function handleMenuSelect(key) {
  router.push(key)
}

function handleLogout() {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
  router.push('/login')
}
</script>
