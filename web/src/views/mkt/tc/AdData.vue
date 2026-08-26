<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" :tool-fn="toolFn" :levels="levels" storage-key="tc" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import { AD_DATA_LEVEL_TABS } from '../adDataFields'
import columnsMap from './adDataColumns.js'
import { getTcAdDataList, tcTool } from '../../../api/mkt/tc'

const levels = ['account', 'campaign', 'unit'].map((k) => AD_DATA_LEVEL_TABS.find((t) => t.key === k))
const fetchList = (level, params) => getTcAdDataList(level, params)
const toolFn = (level, action, data) => tcTool(level, action, data)
</script>