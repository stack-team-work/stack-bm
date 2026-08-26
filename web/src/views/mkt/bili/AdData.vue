<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" :tool-fn="toolFn" :levels="levels" storage-key="bili" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import { AD_DATA_LEVEL_TABS } from '../adDataFields'
import columnsMap from './adDataColumns.js'
import { getBiliAdDataList, biliTool } from '../../../api/mkt/bili'

const levels = ['account', 'campaign', 'unit', 'creative'].map((k) => AD_DATA_LEVEL_TABS.find((t) => t.key === k))
const fetchList = (level, params) => getBiliAdDataList(level, params)
const toolFn = (level, action, data) => biliTool(level, action, data)
</script>