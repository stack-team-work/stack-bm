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
        component: () => import('../views/system/Dashboard.vue'),
        meta: { title: '概况' },
      },
      {
        path: 'admin',
        name: 'SysAdmin',
        component: () => import('../views/system/SysAdmin.vue'),
        meta: { title: '后台账号' },
      },
      {
        path: 'admin-group',
        name: 'SysAdminGroup',
        component: () => import('../views/system/SysAdminGroup.vue'),
        meta: { title: '后台角色' },
      },
      {
        path: 'game',
        name: 'Game',
        component: () => import('../views/game/Game.vue'),
        meta: { title: '游戏管理' },
      },
      {
        path: 'game-app',
        name: 'GameApp',
        component: () => import('../views/game/GameApp.vue'),
        meta: { title: '子游戏' },
      },
      {
        path: 'game-app/create',
        name: 'GameAppCreate',
        component: () => import('../views/game/GameAppForm.vue'),
        meta: { title: '新增子游戏' },
      },
      {
        path: 'game-app/edit/:id',
        name: 'GameAppEdit',
        component: () => import('../views/game/GameAppForm.vue'),
        meta: { title: '编辑子游戏' },
      },
      {
        path: 'game-cp',
        name: 'GameCp',
        component: () => import('../views/game/GameCp.vue'),
        meta: { title: '游戏CP管理' },
      },
      {
        path: 'logs',
        name: 'SysLogs',
        component: () => import('../views/system/SysLogs.vue'),
        meta: { title: '操作日志' },
      },
      {
        path: 'menu',
        name: 'SysMenu',
        component: () => import('../views/system/SysMenu.vue'),
        meta: { title: '菜单管理' },
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
