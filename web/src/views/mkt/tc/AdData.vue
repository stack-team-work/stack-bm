<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" :tool-fn="toolFn" :levels="levels" storage-key="tc" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import { AD_DATA_LEVEL_TABS } from '../adDataFields'
import columnsMap from './adDataColumns.js'
import { getTcAdDataList, tcTool } from '../../../api/mkt/tc'

// 腾讯V3暂无 unit 层级数据源，仅暴露 账户 / 广告组 两层
const levels = ['account', 'campaign'].map((k) => AD_DATA_LEVEL_TABS.find((t) => t.key === k))
const fetchList = (level, params) => getTcAdDataList(level, params)
const toolFn = (level, action, data) => tcTool(level, action, data)
</script>