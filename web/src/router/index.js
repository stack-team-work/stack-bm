import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: { title: '登录' },
  },
  {
    path: '/',
    component: () => import('../layouts/MainLayout.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../views/dashboard/Dashboard.vue'),
        meta: { title: '概况' },
      },
      {
        path: 'admin',
        name: 'SysAdmin',
        component: () => import('../views/bm/sys/SysAdmin.vue'),
        meta: { title: '后台账号' },
      },
      {
        path: 'admin-group',
        name: 'SysAdminGroup',
        component: () => import('../views/bm/sys/SysAdminGroup.vue'),
        meta: { title: '后台角色' },
      },
      {
        path: 'game',
        name: 'Game',
        component: () => import('../views/sdk/game/Game.vue'),
        meta: { title: '游戏管理' },
      },
      {
        path: 'game-app',
        name: 'GameApp',
        component: () => import('../views/sdk/game/GameApp.vue'),
        meta: { title: '子游戏' },
      },
      {
        path: 'game-app/create',
        name: 'GameAppCreate',
        component: () => import('../views/sdk/game/GameAppForm.vue'),
        meta: { title: '新增子游戏' },
      },
      {
        path: 'game-app/edit/:id',
        name: 'GameAppEdit',
        component: () => import('../views/sdk/game/GameAppForm.vue'),
        meta: { title: '编辑子游戏' },
      },
      {
        path: 'game-app-template',
        name: 'GameAppTemplate',
        component: () => import('../views/sdk/game/GameAppTemplate.vue'),
        meta: { title: 'SDK模板' },
      },
      {
        path: 'game-cp',
        name: 'GameCp',
        component: () => import('../views/sdk/game/GameCp.vue'),
        meta: { title: '游戏CP管理' },
      },
      {
        path: 'game-tag',
        name: 'GameTag',
        component: () => import('../views/sdk/game/GameTag.vue'),
        meta: { title: '游戏标签' },
      },
      {
        path: 'game-variable',
        name: 'GameVariable',
        component: () => import('../views/sdk/game/GameVariable.vue'),
        meta: { title: '游戏变量' },
      },
      {
        path: 'game-platform',
        name: 'GamePlatform',
        component: () => import('../views/sdk/game/GamePlatform.vue'),
        meta: { title: '平台管理' },
      },
      {
        path: 'sdk-logs',
        name: 'GameLog',
        component: () => import('../views/sdk/sys/GameLog.vue'),
        meta: { title: '游戏日志' },
      },
      {
        path: 'game-gift',
        name: 'GameGift',
        component: () => import('../views/sdk/game/GameGift.vue'),
        meta: { title: '礼包配置' },
      },
      {
        path: 'game-gift-code',
        name: 'GameGiftCode',
        component: () => import('../views/sdk/game/GameGiftCode.vue'),
        meta: { title: '礼包码配置' },
      },
      {
        path: 'game-gift-user-code',
        name: 'GameGiftUserCode',
        component: () => import('../views/sdk/game/GameGiftUserCode.vue'),
        meta: { title: '领取记录' },
      },
      {
        path: 'game-voucher',
        name: 'GameVoucher',
        component: () => import('../views/sdk/game/GameVoucher.vue'),
        meta: { title: '代金券配置' },
      },
      {
        path: 'game-voucher-use',
        name: 'GameVoucherUse',
        component: () => import('../views/sdk/game/GameVoucherUse.vue'),
        meta: { title: '领取记录' },
      },
      {
        path: 'media',
        name: 'Media',
        component: () => import('../views/mkt/media/Media.vue'),
        meta: { title: '媒体渠道' },
      },
      {
        path: 'media-sub',
        name: 'MediaSub',
        component: () => import('../views/mkt/media/MediaSub.vue'),
        meta: { title: '媒体子渠道' },
      },
      {
        path: 'media-agent',
        name: 'MediaAgent',
        component: () => import('../views/mkt/media/MediaAgent.vue'),
        meta: { title: '渠道代理' },
      },
      {
        path: 'media-application',
        name: 'MediaApplication',
        component: () => import('../views/mkt/media/MediaApplication.vue'),
        meta: { title: 'mkt应用' },
      },
      {
        path: 'media-manager',
        name: 'MediaManager',
        component: () => import('../views/mkt/media/MediaManager.vue'),
        meta: { title: 'mkt管家' },
      },
      {
        path: 'media-subject',
        name: 'MediaSubject',
        component: () => import('../views/mkt/media/MediaSubject.vue'),
        meta: { title: '主体管理' },
      },
      {
        path: 'bili-ads',
        name: 'BiliAds',
        component: () => import('../views/mkt/bili/BiliAds.vue'),
        meta: { title: 'B站Ads' },
      },
      {
        path: 'bili-ads/ad-template/create',
        name: 'BiliAdTemplateCreate',
        component: () => import('../views/mkt/bili/AdTemplateForm.vue'),
        meta: { title: '新增B站广告模板' },
      },
      {
        path: 'bili-ads/ad-template/edit/:id',
        name: 'BiliAdTemplateEdit',
        component: () => import('../views/mkt/bili/AdTemplateForm.vue'),
        meta: { title: '编辑B站广告模板' },
      },
      {
        path: 'bili-ads/audience-template/create',
        name: 'BiliAudienceTemplateCreate',
        component: () => import('../views/mkt/bili/AudienceTemplateForm.vue'),
        meta: { title: '新增B站定向包模板' },
      },
      {
        path: 'bili-ads/audience-template/edit/:id',
        name: 'BiliAudienceTemplateEdit',
        component: () => import('../views/mkt/bili/AudienceTemplateForm.vue'),
        meta: { title: '编辑B站定向包模板' },
      },
      {
        path: 'bili-ads/title-template/create',
        name: 'BiliTitleTemplateCreate',
        component: () => import('../views/mkt/bili/TitleTemplateForm.vue'),
        meta: { title: '新增B站标题包模板' },
      },
      {
        path: 'bili-ads/title-template/edit/:id',
        name: 'BiliTitleTemplateEdit',
        component: () => import('../views/mkt/bili/TitleTemplateForm.vue'),
        meta: { title: '编辑B站标题包模板' },
      },
      {
        path: 'ks-ads',
        name: 'KsAds',
        component: () => import('../views/mkt/ks/KsAds.vue'),
        meta: { title: '快手Ads' },
      },
      {
        path: 'ks-ads/ad-template/create',
        name: 'KsAdTemplateCreate',
        component: () => import('../views/mkt/ks/AdTemplateForm.vue'),
        meta: { title: '新增快手广告模板' },
      },
      {
        path: 'ks-ads/ad-template/edit/:id',
        name: 'KsAdTemplateEdit',
        component: () => import('../views/mkt/ks/AdTemplateForm.vue'),
        meta: { title: '编辑快手广告模板' },
      },
      {
        path: 'ks-ads/audience-template/create',
        name: 'KsAudienceTemplateCreate',
        component: () => import('../views/mkt/ks/AudienceTemplateForm.vue'),
        meta: { title: '新增快手定向包模板' },
      },
      {
        path: 'ks-ads/audience-template/edit/:id',
        name: 'KsAudienceTemplateEdit',
        component: () => import('../views/mkt/ks/AudienceTemplateForm.vue'),
        meta: { title: '编辑快手定向包模板' },
      },
      {
        path: 'ks-ads/title-template/create',
        name: 'KsTitleTemplateCreate',
        component: () => import('../views/mkt/ks/TitleTemplateForm.vue'),
        meta: { title: '新增快手标题包模板' },
      },
      {
        path: 'ks-ads/title-template/edit/:id',
        name: 'KsTitleTemplateEdit',
        component: () => import('../views/mkt/ks/TitleTemplateForm.vue'),
        meta: { title: '编辑快手标题包模板' },
      },
      {
        path: 'tt-ads',
        name: 'TtAds',
        component: () => import('../views/mkt/tt/TtAds.vue'),
        meta: { title: '头条Ads' },
      },
      {
        path: 'tt-ads/ad-template/create',
        name: 'TtAdTemplateCreate',
        component: () => import('../views/mkt/tt/AdTemplateForm.vue'),
        meta: { title: '新增头条广告模板' },
      },
      {
        path: 'tt-ads/ad-template/edit/:id',
        name: 'TtAdTemplateEdit',
        component: () => import('../views/mkt/tt/AdTemplateForm.vue'),
        meta: { title: '编辑头条广告模板' },
      },
      {
        path: 'tt-ads/audience-template/create',
        name: 'TtAudienceTemplateCreate',
        component: () => import('../views/mkt/tt/AudienceTemplateForm.vue'),
        meta: { title: '新增头条定向包模板' },
      },
      {
        path: 'tt-ads/audience-template/edit/:id',
        name: 'TtAudienceTemplateEdit',
        component: () => import('../views/mkt/tt/AudienceTemplateForm.vue'),
        meta: { title: '编辑头条定向包模板' },
      },
      {
        path: 'tt-ads/title-template/create',
        name: 'TtTitleTemplateCreate',
        component: () => import('../views/mkt/tt/TitleTemplateForm.vue'),
        meta: { title: '新增头条标题包模板' },
      },
      {
        path: 'tt-ads/title-template/edit/:id',
        name: 'TtTitleTemplateEdit',
        component: () => import('../views/mkt/tt/TitleTemplateForm.vue'),
        meta: { title: '编辑头条标题包模板' },
      },
      {
        path: 'pay-platform',
        name: 'PayPlatform',
        component: () => import('../views/sdk/pay/PayPlatform.vue'),
        meta: { title: '支付平台' },
      },
      {
        path: 'pay-merchant',
        name: 'PayMerchant',
        component: () => import('../views/sdk/pay/PayMerchant.vue'),
        meta: { title: '平台商户' },
      },
      {
        path: 'logs',
        name: 'SysLogs',
        component: () => import('../views/bm/sys/SysLogs.vue'),
        meta: { title: '后台日志' },
      },
      {
        path: 'menu',
        name: 'SysMenu',
        component: () => import('../views/bm/sys/SysMenu.vue'),
        meta: { title: '菜单管理' },
      },
      {
        path: 'sys-column',
        name: 'SysColumn',
        component: () => import('../views/bm/sys/SysColumn.vue'),
        meta: { title: '报表指标' },
      },
      {
        path: 'feishu-user',
        name: 'FeishuUser',
        component: () => import('../views/bm/sys/FeishuUser.vue'),
        meta: { title: '飞书绑定' },
      },
      {
        path: 'feishu-app',
        name: 'FeishuApp',
        component: () => import('../views/bm/sys/FeishuApp.vue'),
        meta: { title: '飞书应用' },
      },
      {
        path: 'feishu-chat',
        name: 'FeishuChat',
        component: () => import('../views/bm/sys/FeishuChat.vue'),
        meta: { title: '飞书聊天' },
      },
    ],
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.path !== '/login' && !token) {
    next('/login')
  } else if (to.path === '/login' && token) {
    next('/')
  } else {
    next()
  }
})

export default router
