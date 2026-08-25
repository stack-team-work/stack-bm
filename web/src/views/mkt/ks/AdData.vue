<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import columnsMap from './adDataColumns.js'
import { getKsAdDataList } from '../../../api/mkt/ks'

const fetchList = (level, params) => getKsAdDataList(level, params)
</script>