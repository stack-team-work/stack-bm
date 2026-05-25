<template>
  <div>
    <n-space vertical :size="24">
      <n-grid :cols="4" :x-gap="24">
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="今日充值人数" :value="statRecharge" />
          </n-card>
        </n-gi>
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="今日注册人数" :value="statRegister" />
          </n-card>
        </n-gi>
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="今日登录人数" :value="statLogin" />
          </n-card>
        </n-gi>
        <n-gi>
          <n-card size="small" hoverable>
            <n-statistic label="今日激活人数" :value="statActivate" />
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
import { getDashboardStats, getLogList } from '../../api/system'
import { NTag } from 'naive-ui'
import { formatTime } from '../../utils/format'
import { h } from 'vue'

const loading = ref(false)
const statRecharge = ref(0)
const statRegister = ref(0)
const statLogin = ref(0)
const statActivate = ref(0)
const recentLogs = ref([])

const logColumns = [
  { title: '级别', key: 'level', width: 60, render: (row) => h(NTag, { type: row.level === 'error' ? 'error' : 'info', size: 'small' }, { default: () => row.level }) },
  { title: '路径', key: 'path', width: 160 },
  { title: '用户', key: 'username', width: 80 },
  { title: '描述', key: 'desc', width: 200, ellipsis: { tooltip: true } },
  { title: '时间', key: 'created_at', width: 150, render: (row) => formatTime(row.created_at) },
]

onMounted(async () => {
  loading.value = true
  try {
    const [stats, logs] = await Promise.all([
      getDashboardStats(),
      getLogList({ page: 1, size: 10, level: '', keyword: '' }),
    ])
    statRecharge.value = stats.data.recharge_users || 0
    statRegister.value = stats.data.register_users || 0
    statLogin.value = stats.data.login_users || 0
    statActivate.value = stats.data.activate_users || 0
    recentLogs.value = logs.data.list || []
  } catch { /* */ }
  finally { loading.value = false }
})
</script>
