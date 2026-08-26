<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" :tool-fn="toolFn" :levels="levels" storage-key="tt" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import { AD_DATA_LEVEL_TABS } from '../adDataFields'
import columnsMap from './adDataColumns.js'
import { getTtAdDataList, ttTool } from '../../../api/mkt/tt'

const levels = ['account', 'campaign'].map((k) => AD_DATA_LEVEL_TABS.find((t) => t.key === k))
const fetchList = (level, params) => getTtAdDataList(level, params)
const toolFn = (level, action, data) => ttTool(level, action, data)
</script>