<template>
  <n-card :bordered="false">
    <n-empty v-if="!columns.length" description="该层级暂无数据" style="padding: 40px 0" />
    <n-data-table v-else :bordered="false" :columns="columns" :data="tableData" :loading="loading" :pagination="pagination" @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />
  </n-card>
</template>

<script setup>
import { onMounted } from 'vue'
import { useTable } from '../../composables/useTable'

const props = defineProps({
  fetchFn: { type: Function, required: true },
  columns: { type: Array, required: true },
})

const { loading, tableData, pagination, search, handlePageChange, handlePageSizeChange } = useTable(props.fetchFn)
onMounted(() => search({}))
</script>