<template>
  <n-layout has-sider style="height: 100vh">
    <n-layout-sider
        bordered
        collapse-mode="width"
        :collapsed-width="64"
        :width="220"
        :collapsed="collapsed"
        :inverted="true"
        :native-scrollbar="false"
      >
      <div class="layout-logo">
        <img src="/logo.png" class="layout-logo-img" />
        <span v-if="!collapsed" class="layout-logo-text">STACK-BM</span>
      </div>
      <n-menu
        :collapsed="collapsed"
        :collapsed-width="64"
        :collapsed-icon-size="22"
        :value="currentRoute"
        :options="sidebarMenu"
        :inverted="true"
        :root-indent="20"
        :indent="24"
        @update:value="handleMenuSelect"
      />
    </n-layout-sider>
    <n-layout>
      <n-layout-header bordered class="layout-header">
        <n-space :size="8" align="center">
          <n-button quaternary size="small" class="layout-collapse-btn" @click="collapsed = !collapsed">
            <n-icon :size="20"><menu-outline /></n-icon>
          </n-button>
          <n-menu
            mode="horizontal"
            :value="activeTopNav"
            :options="topNavOptions"
            @update:value="handleTopNav"
          />
        </n-space>
        <n-dropdown :options="userDropdownOptions" @select="handleUserAction">
          <div class="layout-user">
            <n-avatar round :size="28" class="layout-user-avatar">
              <n-icon><person-outline /></n-icon>
            </n-avatar>
            <span class="layout-user-name">{{ userInfo?.username || '' }}</span>
            <n-icon :size="14" class="layout-user-caret"><chevron-down-outline /></n-icon>
          </div>
        </n-dropdown>
      </n-layout-header>
      <n-layout-content class="layout-content">
        <router-view v-slot="{ Component }">
          <transition name="fade-slide" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<style>
.n-menu.n-menu--horizontal .n-menu-item-content {
  padding: 0 12px;
  border-radius: 4px;
}
.n-menu.n-menu--horizontal .n-menu-item-content.n-menu-item-content--selected {
  background: #f0f0f0 !important;
}
.n-menu.n-menu--horizontal .n-menu-item-content:hover {
  background: #f5f5f5 !important;
}

.layout-logo {
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}
.layout-logo-img {
  width: 32px;
  height: 32px;
  flex-shrink: 0;
}
.layout-logo-text {
  color: #fff;
  font-size: 16px;
  font-weight: bold;
  letter-spacing: 2px;
}

.layout-header {
  height: 56px;
  padding: 0 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
}
.layout-collapse-btn {
  font-size: 20px;
}

.layout-user {
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
}
.layout-user-avatar {
  background: #18a058;
}
.layout-user-name {
  font-size: 14px;
}
.layout-user-caret {
  color: #999;
}

.layout-content {
  padding: 16px 20px 24px;
  background: #f5f6fa;
  min-height: calc(100vh - 56px);
}

.layout-breadcrumb {
  margin-bottom: 16px;
}

.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}
.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(8px);
}
.fade-slide-leave-to {
  opacity: 0;
}
</style>

<script setup>
import { ref, computed, h, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { NIcon } from 'naive-ui'
import {
  HomeOutline, SettingsOutline, GameControllerOutline,
  BarChartOutline, GiftOutline, TicketOutline, CashOutline,
  ServerOutline, DocumentTextOutline, MenuOutline,
  PersonOutline, ChevronDownOutline,
} from '@vicons/ionicons5'

const router = useRouter()
const route = useRoute()
const collapsed = ref(false)

const userInfo = JSON.parse(localStorage.getItem('userInfo') || 'null')
const userDropdownOptions = [{ label: '退出登录', key: 'logout' }]

function handleUserAction(key) { if (key === 'logout') handleLogout() }

const renderIcon = (icon) => () => h(NIcon, null, { default: () => h(icon) })

const topNavOptions = [
  { label: '发行中心', key: 'publish' },
  { label: '游戏中心', key: 'game' },
  { label: '游戏运营', key: 'operation' },
  { label: '财务中心', key: 'finance' },
  { label: '研发数据', key: 'dev-data' },
  { label: '系统设置', key: 'system' },
]

const topNavRouteMap = {
  '/dashboard': 'publish',
  '/media': 'publish', '/media-sub': 'publish', '/media-agent': 'publish',
  '/media-application': 'publish', '/media-manager': 'publish', '/media-subject': 'publish',
  '/bili-ads': 'publish',
  '/bili-ads/ad-template/create': 'publish', '/bili-ads/ad-template/edit': 'publish',
  '/bili-ads/audience-template/create': 'publish', '/bili-ads/audience-template/edit': 'publish',
  '/bili-ads/title-template/create': 'publish', '/bili-ads/title-template/edit': 'publish',
  '/ks-ads': 'publish',
  '/ks-ads/ad-template/create': 'publish', '/ks-ads/ad-template/edit': 'publish',
  '/ks-ads/audience-template/create': 'publish', '/ks-ads/audience-template/edit': 'publish',
  '/ks-ads/title-template/create': 'publish', '/ks-ads/title-template/edit': 'publish',
  '/tt-ads': 'publish',
  '/tt-ads/ad-template/create': 'publish', '/tt-ads/ad-template/edit': 'publish',
  '/tt-ads/audience-template/create': 'publish', '/tt-ads/audience-template/edit': 'publish',
  '/tt-ads/title-template/create': 'publish', '/tt-ads/title-template/edit': 'publish',
  '/game': 'game', '/game-app': 'game', '/game-app/create': 'game', '/game-app/edit': 'game',
  '/game-app-template': 'game',
  '/game-cp': 'game', '/game-tag': 'game', '/game-variable': 'game',
  '/game-platform': 'game', '/sdk-logs': 'game',
  '/game-gift': 'operation', '/game-gift-code': 'operation', '/game-gift-user-code': 'operation',
  '/game-voucher': 'operation', '/game-voucher-use': 'operation',
  '/pay-platform': 'finance', '/pay-merchant': 'finance',
  '/admin': 'system', '/admin-group': 'system', '/menu': 'system', '/sys-column': 'system', '/logs': 'system',
  '/feishu-user': 'system', '/feishu-app': 'system', '/feishu-chat': 'system',
}

const activeTopNav = ref(topNavRouteMap[route.path] || 'publish')

watch(() => route.path, (p) => {
  const nav = topNavRouteMap[p]
  if (nav) activeTopNav.value = nav
})

const topNavFirstRoute = {
  publish: '/dashboard',
  game: '/game',
  operation: '/game-gift',
  finance: '/pay-platform',
  'dev-data': '',
  system: '/admin',
}

function handleTopNav(key) {
  activeTopNav.value = key
  const first = topNavFirstRoute[key]
  if (first) router.push(first)
}

const sidebarMenu = computed(() => {
  switch (activeTopNav.value) {
    case 'publish':
      return [
        { label: '概况', key: '/dashboard', icon: renderIcon(HomeOutline) },
        {
          label: '渠道管理', key: 'mkt-mgmt', icon: renderIcon(BarChartOutline),
          children: [
            { label: '媒体渠道', key: '/media' },
            { label: '媒体子渠道', key: '/media-sub' },
            { label: '渠道代理', key: '/media-agent' },
            { label: 'mkt应用', key: '/media-application' },
            { label: 'mkt管家', key: '/media-manager' },
            { label: '主体管理', key: '/media-subject' },
          ],
        },
        {
          label: 'B站广告', key: 'bili-ads', icon: renderIcon(GameControllerOutline),
          children: [
            { label: '模板列表', key: '/bili-ads' },
          ],
        },
        {
          label: '快手广告', key: 'ks-ads', icon: renderIcon(GameControllerOutline),
          children: [
            { label: '模板列表', key: '/ks-ads' },
          ],
        },
        {
          label: '头条广告', key: 'tt-ads', icon: renderIcon(GameControllerOutline),
          children: [
            { label: '模板列表', key: '/tt-ads' },
          ],
        },
      ]
    case 'game':
      return [{
        label: '游戏管理', key: 'game-mgmt', icon: renderIcon(GameControllerOutline),
        children: [
          { label: '父游戏', key: '/game' },
          { label: '子游戏', key: '/game-app' },
          { label: 'SDK模板', key: '/game-app-template' },
          { label: '游戏CP', key: '/game-cp' },
          { label: '游戏标签', key: '/game-tag' },
          { label: '游戏变量', key: '/game-variable' },
          { label: '平台管理', key: '/game-platform' },
        ],
      },
      { label: '游戏日志', key: '/sdk-logs', icon: renderIcon(DocumentTextOutline) },
      ]
    case 'operation':
      return [
        {
          label: '礼包管理', key: 'gift-mgmt', icon: renderIcon(GiftOutline),
          children: [
            { label: '礼包配置', key: '/game-gift' },
            { label: '礼包码配置', key: '/game-gift-code' },
            { label: '使用记录', key: '/game-gift-user-code' },
          ],
        },
        {
          label: '代金券', key: 'voucher-mgmt', icon: renderIcon(TicketOutline),
          children: [
            { label: '代金券配置', key: '/game-voucher' },
            { label: '使用记录', key: '/game-voucher-use' },
          ],
        },
      ]
    case 'finance':
      return [{
        label: '支付设置', key: 'pay-mgmt', icon: renderIcon(CashOutline),
        children: [
          { label: '支付平台', key: '/pay-platform' },
          { label: '平台商户', key: '/pay-merchant' },
        ],
      }]
    case 'dev-data':
      return [{ label: '暂无数据', key: 'empty', icon: renderIcon(ServerOutline) }]
    case 'system':
      return [
        {
          label: '系统管理', key: 'sys-mgmt', icon: renderIcon(SettingsOutline),
          children: [
            { label: '后台账号', key: '/admin' },
            { label: '后台角色', key: '/admin-group' },
            { label: '菜单管理', key: '/menu' },
            { label: '报表指标', key: '/sys-column' },
          ],
        },
        { label: '后台日志', key: '/logs', icon: renderIcon(DocumentTextOutline) },
        {
          label: '飞书管理', key: 'feishu-mgmt', icon: renderIcon(SettingsOutline),
          children: [
            { label: '飞书绑定', key: '/feishu-user' },
            { label: '飞书应用', key: '/feishu-app' },
            { label: '飞书聊天', key: '/feishu-chat' },
          ],
        },
      ]
    default:
      return []
  }
})

const currentRoute = computed(() => route.path)

function handleMenuSelect(key) {
  if (key !== 'empty') router.push(key)
}

function handleLogout() {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
  router.push('/login')
}
</script>
