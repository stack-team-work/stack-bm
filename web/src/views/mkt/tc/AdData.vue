<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import columnsMap from './adDataColumns.js'
import { getTcAdDataList } from '../../../api/mkt/tc'

const fetchList = (level, params) => getTcAdDataList(level, params)
</script>