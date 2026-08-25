<template>
  <div>
    <n-tabs type="line" :value="activeTab" @update:value="onTabChange">
      <n-tab-pane v-for="t in levels" :key="t.key" :name="t.key" :tab="t.label">
        <AdDataTable :fetch-fn="(params) => fetchList(t.key, params)" :columns="columnsMap[t.key] || []" />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import AdDataTable from './AdDataTable.vue'
import { AD_DATA_LEVEL_TABS } from './adDataFields'

const props = defineProps({
  columnsMap: { type: Object, required: true },
  fetchList: { type: Function, required: true },
  levels: { type: Array, default: () => AD_DATA_LEVEL_TABS },
})

const activeTab = ref(props.levels[0]?.key || 'account')

function onTabChange(key) {
  activeTab.value = key
}
</script>