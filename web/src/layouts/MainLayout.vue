<template>
  <n-layout has-sider style="height: 100vh">
    <n-layout-sider
      bordered
      collapse-mode="width"
      :collapsed-width="64"
      :width="220"
      :collapsed="collapsed"
      show-trigger
      @collapse="collapsed = true"
      @expand="collapsed = false"
    >
      <n-menu
        :collapsed="collapsed"
        :collapsed-width="64"
        :collapsed-icon-size="22"
        :value="currentRoute"
        :options="menuOptions"
        @update:value="handleMenuSelect"
      />
    </n-layout-sider>
    <n-layout>
      <n-layout-header bordered style="height: 56px; padding: 0 24px; display: flex; align-items: center; justify-content: space-between">
        <n-space align="center">
          <span style="font-size: 18px; font-weight: bold">游戏发行后台管理系统</span>
        </n-space>
        <n-space align="center">
          <n-text>{{ userInfo?.username || '' }}</n-text>
          <n-button size="small" @click="handleLogout">退出登录</n-button>
        </n-space>
      </n-layout-header>
      <n-layout-content style="padding: 24px; background: #f5f7fa; min-height: calc(100vh - 56px)">
        <router-view />
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<script setup>
import { ref, computed, h } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { NIcon } from 'naive-ui'
import { PeopleOutline, AppsOutline, GameControllerOutline, SettingsOutline } from '@vicons/ionicons5'

const router = useRouter()
const route = useRoute()
const collapsed = ref(false)

const userInfo = JSON.parse(localStorage.getItem('userInfo') || 'null')

const renderIcon = (icon) => () => h(NIcon, null, { default: () => h(icon) })

const menuOptions = [
  {
    label: '系统管理',
    key: 'system',
    icon: renderIcon(SettingsOutline),
    children: [
      { label: '管理员管理', key: '/admin' },
      { label: '管理员分组', key: '/admin-group' },
    ],
  },
  {
    label: '游戏管理',
    key: 'game-mgmt',
    icon: renderIcon(GameControllerOutline),
    children: [
      { label: '父游戏', key: '/game' },
      { label: '子游戏管理', key: '/game-app' },
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
