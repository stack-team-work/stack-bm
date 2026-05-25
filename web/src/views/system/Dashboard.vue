<template>
  <div>
    <n-space vertical :size="24">
      <n-grid :cols="4" :x-gap="24">
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="后台账号" :value="statAdmin" />
          </n-card>
        </n-gi>
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="后台角色" :value="statGroup" />
          </n-card>
        </n-gi>
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="父游戏" :value="statGame" />
          </n-card>
        </n-gi>
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="子游戏" :value="statApp" />
          </n-card>
        </n-gi>
      </n-grid>

      <n-card title="最近操作日志" size="small" :bordered="false">
        <n-data-table :columns="logColumns" :data="recentLogs" :loading="loading" :pagination="false" size="small" />
      </n-card>
    </n-space>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getAdminList, getAdminGroupList, getLogList } from '../../api/system'
import { getGameList, getGameAppList } from '../../api/game'
import { NTag } from 'naive-ui'
import { h } from 'vue'

const loading = ref(false)
const statAdmin = ref(0)
const statGroup = ref(0)
const statGame = ref(0)
const statApp = ref(0)
const recentLogs = ref([])

const logColumns = [
  { title: '级别', key: 'level', width: 60, render: (row) => h(NTag, { type: row.level === 'error' ? 'error' : 'info', size: 'small' }, { default: () => row.level }) },
  { title: '路径', key: 'path', width: 160 },
  { title: '用户', key: 'username', width: 80 },
  { title: '描述', key: 'desc', width: 200, ellipsis: { tooltip: true } },
  { title: '时间', key: 'created_at', width: 150 },
]

onMounted(async () => {
  loading.value = true
  try {
    const [ar, gr, g, a, logs] = await Promise.all([
      getAdminList({ page: 1, size: 1 }),
      getAdminGroupList({ page: 1, size: 1 }),
      getGameList({ page: 1, size: 1 }),
      getGameAppList({ page: 1, size: 1 }),
      getLogList({ page: 1, size: 10, level: '', keyword: '' }),
    ])
    statAdmin.value = ar.data.total || 0
    statGroup.value = gr.data.total || 0
    statGame.value = g.data.total || 0
    statApp.value = a.data.total || 0
    recentLogs.value = logs.data.list || []
  } catch { /* */ }
  finally { loading.value = false }
})
</script>
