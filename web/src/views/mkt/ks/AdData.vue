<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" :levels="levels" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import { AD_DATA_LEVEL_TABS } from '../adDataFields'
import columnsMap from './adDataColumns.js'
import { getKsAdDataList } from '../../../api/mkt/ks'

const levels = ['account', 'campaign', 'unit', 'creative'].map((k) => AD_DATA_LEVEL_TABS.find((t) => t.key === k))
const fetchList = (level, params) => getKsAdDataList(level, params)
</script>