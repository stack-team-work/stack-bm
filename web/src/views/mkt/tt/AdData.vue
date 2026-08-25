<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import columnsMap from './adDataColumns.js'
import { getTtAdDataList } from '../../../api/mkt/tt'

const fetchList = (level, params) => getTtAdDataList(level, params)
</script>