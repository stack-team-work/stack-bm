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
      <div style="height: 56px; display: flex; align-items: center; justify-content: center; border-bottom: 1px solid rgba(255,255,255,0.08); gap: 8px">
        <img src="/logo.png" style="width: 32px; height: 32px; flex-shrink: 0" />
        <span v-if="!collapsed" style="color: #fff; font-size: 16px; font-weight: bold; letter-spacing: 2px">STACK-BM</span>
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
        <n-breadcrumb>
          <n-breadcrumb-item v-for="item in breadcrumbs" :key="item.path" @click="item.path && router.push(item.path)">
            {{ item.label }}
          </n-breadcrumb-item>
        </n-breadcrumb>
        <n-dropdown :options="userDropdownOptions" @select="handleUserAction">
          <span style="cursor: pointer; display: flex; align-items: center; gap: 4px">
            <n-text>{{ userInfo?.username || '' }}</n-text>
            <span style="font-size: 12px; color: #999">▼</span>
          </span>
        </n-dropdown>
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
import { HomeOutline, SettingsOutline, GameControllerOutline, BarChartOutline } from '@vicons/ionicons5'

const router = useRouter()
const route = useRoute()
const collapsed = ref(false)

const userInfo = JSON.parse(localStorage.getItem('userInfo') || 'null')

const userDropdownOptions = [
  { label: '退出登录', key: 'logout' },
]

function handleUserAction(key) {
  if (key === 'logout') handleLogout()
}

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
      { label: '后台日志', key: '/logs' },
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
      { label: '游戏标签', key: '/game-tag' },
      { label: '游戏变量', key: '/game-variable' },
      { label: '平台管理', key: '/game-platform' },
      { label: '游戏日志', key: '/sdk-logs' },
    ],
  },
  {
    label: '渠道管理',
    key: 'mkt-mgmt',
    icon: renderIcon(BarChartOutline),
    children: [
      { label: '媒体渠道', key: '/media' },
      { label: '媒体子渠道', key: '/media-sub' },
      { label: '渠道代理', key: '/media-agent' },
      { label: 'mkt应用', key: '/media-application' },
      { label: 'mkt管家', key: '/media-manager' },
      { label: '主体管理', key: '/media-subject' },
    ],
  },
]

const currentRoute = computed(() => route.path)

const breadcrumbs = computed(() => {
  return route.matched
    .filter(r => r.meta?.title)
    .map(r => ({ label: r.meta.title, path: r.path }))
})

function handleMenuSelect(key) {
  router.push(key)
}

function handleLogout() {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
  router.push('/login')
}
</script>
