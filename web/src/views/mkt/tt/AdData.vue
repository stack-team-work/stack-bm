<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" :levels="levels" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import { AD_DATA_LEVEL_TABS } from '../adDataFields'
import columnsMap from './adDataColumns.js'
import { getTtAdDataList } from '../../../api/mkt/tt'

const levels = ['account', 'campaign'].map((k) => AD_DATA_LEVEL_TABS.find((t) => t.key === k))
const fetchList = (level, params) => getTtAdDataList(level, params)
</script>