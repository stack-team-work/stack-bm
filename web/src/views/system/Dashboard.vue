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

      <n-card size="small" :bordered="false">
        <template #header>
          <n-space align="center" justify="space-between">
            <span>充值趋势（近15天）</span>
            <n-space :size="8">
              <n-button-group size="tiny">
                <n-button :type="rechargeView === 'line' ? 'primary' : 'default'" @click="rechargeView = 'line'">折线图</n-button>
                <n-button :type="rechargeView === 'bar' ? 'primary' : 'default'" @click="rechargeView = 'bar'">柱状图</n-button>
                <n-button :type="rechargeView === 'table' ? 'primary' : 'default'" @click="rechargeView = 'table'">表格</n-button>
              </n-button-group>
              <n-button size="tiny" v-if="rechargeView !== 'table'" @click="downloadChart('rechargeChart')">
                <template #icon><n-icon><download-outline /></n-icon></template>
              </n-button>
            </n-space>
          </n-space>
        </template>
        <div v-if="rechargeView !== 'table'" style="margin-bottom: 8px; display: flex; gap: 20px; flex-wrap: wrap">
          <span v-for="s in rechargeHighLow" :key="s.name" style="background: #f0f5ff; border: 1px solid #d6e4ff; border-radius: 4px; padding: 2px 10px; font-size: 12px; font-weight: 600; color: #1a1a2e">
            {{ s.name }} <span style="color: #e74c3c">▲ {{ s.maxVal }} ({{ s.maxDate }})</span> <span style="color: #3498db">▼ {{ s.minVal }} ({{ s.minDate }})</span>
          </span>
        </div>
        <v-chart v-if="rechargeView !== 'table'" ref="rechargeChart" :option="rechargeOption" style="height: 320px" autoresize />
        <n-data-table v-else :columns="rechargeTableColumns" :data="rechargeTableData" :pagination="false" size="small" :bordered="false" />
      </n-card>

      <n-card size="small" :bordered="false">
        <template #header>
          <n-space align="center" justify="space-between">
            <span>用户趋势（近15天）</span>
            <n-space :size="8">
              <n-button-group size="tiny">
                <n-button :type="userView === 'line' ? 'primary' : 'default'" @click="userView = 'line'">折线图</n-button>
                <n-button :type="userView === 'bar' ? 'primary' : 'default'" @click="userView = 'bar'">柱状图</n-button>
                <n-button :type="userView === 'table' ? 'primary' : 'default'" @click="userView = 'table'">表格</n-button>
              </n-button-group>
              <n-button size="tiny" v-if="userView !== 'table'" @click="downloadChart('userChart')">
                <template #icon><n-icon><download-outline /></n-icon></template>
              </n-button>
            </n-space>
          </n-space>
        </template>
        <div v-if="userView !== 'table'" style="margin-bottom: 8px; display: flex; gap: 20px; flex-wrap: wrap">
          <span v-for="s in userHighLow" :key="s.name" style="background: #f0f5ff; border: 1px solid #d6e4ff; border-radius: 4px; padding: 2px 10px; font-size: 12px; font-weight: 600; color: #1a1a2e">
            {{ s.name }} <span style="color: #e74c3c">▲ {{ s.maxVal }} ({{ s.maxDate }})</span> <span style="color: #3498db">▼ {{ s.minVal }} ({{ s.minDate }})</span>
          </span>
        </div>
        <v-chart v-if="userView !== 'table'" ref="userChart" :option="userOption" style="height: 320px" autoresize />
        <n-data-table v-else :columns="userTableColumns" :data="userTableData" :pagination="false" size="small" :bordered="false" />
      </n-card>
    </n-space>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { DownloadOutline } from '@vicons/ionicons5'
import { use } from 'echarts/core'
import { LineChart, BarChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, LegendComponent, ToolboxComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import VChart from 'vue-echarts'
import { getDashboardStats } from '../../api/system'

use([LineChart, BarChart, GridComponent, TooltipComponent, LegendComponent, ToolboxComponent, CanvasRenderer])

const statRecharge = ref(0)
const statRegister = ref(0)
const statLogin = ref(0)
const statActivate = ref(0)
const rechargeView = ref('line')
const userView = ref('line')
const rechargeChart = ref(null)
const userChart = ref(null)

function mockDates(count) {
  const dates = []
  const now = new Date()
  for (let i = count - 1; i >= 0; i--) {
    const d = new Date(now)
    d.setDate(d.getDate() - i)
    dates.push(`${d.getMonth() + 1}/${d.getDate()}`)
  }
  return dates
}

function mockValues(count, base, range) {
  return Array.from({ length: count }, () => Math.floor(base + Math.random() * range))
}

const dates = mockDates(15)
const rechargeSeries = [
  { name: '充值金额', data: mockValues(15, 2000, 3000) },
  { name: '新增充值', data: mockValues(15, 500, 1000) },
  { name: '成本', data: mockValues(15, 300, 500) },
]
const userSeries = [
  { name: '注册用户', data: mockValues(15, 100, 200) },
  { name: '登录用户', data: mockValues(15, 80, 150) },
  { name: '激活用户', data: mockValues(15, 50, 100) },
]

function makeOption(seriesData, chartType, colors) {
  return {
    color: colors,
    tooltip: { trigger: 'axis' },
    legend: { data: seriesData.map(s => s.name), top: 0 },
    toolbox: {
      feature: {
        magicType: { type: ['line', 'bar'] },
        saveAsImage: {},
      },
    },
    grid: { left: '3%', right: '4%', top: 40, bottom: '3%', containLabel: true },
    xAxis: { type: 'category', data: dates },
    yAxis: { type: 'value' },
    series: seriesData.map(s => ({
      name: s.name,
      type: chartType,
      smooth: chartType === 'line',
      data: s.data,
      markPoint: {
        data: [
          { type: 'max', name: '最大值', label: { fontSize: 10 } },
          { type: 'min', name: '最小值', label: { fontSize: 10 } },
        ],
      },
    })),
  }
}

const rechargeOption = computed(() => makeOption(rechargeSeries, rechargeView.value, ['#00b894', '#e17055', '#fdcb6e']))
const userOption = computed(() => makeOption(userSeries, userView.value, ['#00b894', '#e17055', '#fdcb6e']))

function highLow(s) {
  let maxVal = -Infinity, minVal = Infinity, maxIdx = 0, minIdx = 0
  s.data.forEach((v, i) => {
    if (v > maxVal) { maxVal = v; maxIdx = i }
    if (v < minVal) { minVal = v; minIdx = i }
  })
  return { name: s.name, maxVal, maxDate: dates[maxIdx], minVal, minDate: dates[minIdx] }
}

const rechargeHighLow = computed(() => rechargeSeries.map(highLow))
const userHighLow = computed(() => userSeries.map(highLow))

function makeTableColumns(seriesData) {
  const cols = [{ title: '日期', key: 'date', width: 90 }]
  seriesData.forEach(s => cols.push({ title: s.name, key: s.name }))
  return cols
}

function makeTableData(seriesData) {
  return dates.map((d, i) => {
    const row = { date: d }
    seriesData.forEach(s => { row[s.name] = s.data[i] })
    return row
  })
}

const rechargeTableColumns = computed(() => makeTableColumns(rechargeSeries))
const rechargeTableData = computed(() => makeTableData(rechargeSeries))
const userTableColumns = computed(() => makeTableColumns(userSeries))
const userTableData = computed(() => makeTableData(userSeries))

function downloadChart(refName) {
  const chart = refName === 'rechargeChart' ? rechargeChart.value : userChart.value
  if (chart) {
    const url = chart.getDataURL({ type: 'png', pixelRatio: 2, backgroundColor: '#fff' })
    const a = document.createElement('a')
    a.href = url
    a.download = `${refName}.png`
    a.click()
  }
}

onMounted(async () => {
  try {
    const stats = await getDashboardStats()
    statRecharge.value = stats.data.recharge_users || 0
    statRegister.value = stats.data.register_users || 0
    statLogin.value = stats.data.login_users || 0
    statActivate.value = stats.data.activate_users || 0
  } catch { /* */ }
})
</script>
