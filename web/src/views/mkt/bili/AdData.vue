<template>
  <AdDataTabs :columns-map="columnsMap" :fetch-list="fetchList" />
</template>

<script setup>
import AdDataTabs from '../AdDataTabs.vue'
import columnsMap from './adDataColumns.js'
import { getBiliAdDataList } from '../../../api/mkt/bili'

const fetchList = (level, params) => getBiliAdDataList(level, params)
</script>