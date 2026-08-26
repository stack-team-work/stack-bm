<template>
  <div>
    <n-tabs type="line" :value="activeTab" @update:value="onTabChange">
      <n-tab-pane v-for="t in levels" :key="t.key" :name="t.key" :tab="t.label">
        <AdDataTable :fetch-fn="(params) => fetchList(t.key, params)" :columns="columnsMap[t.key] || []" :level="t.key" :storage-key="storageKey" :actions="actionsMap[t.key] || { idKey: 'id', row: [], batch: [] }" :tool-fn="toolFn" />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import AdDataTable from './AdDataTable.vue'
import { AD_DATA_LEVEL_TABS, AD_DATA_ACTIONS } from './adDataFields'

const props = defineProps({
  columnsMap: { type: Object, required: true },
  fetchList: { type: Function, required: true },
  toolFn: { type: Function, default: null },
  levels: { type: Array, default: () => AD_DATA_LEVEL_TABS },
  actionsMap: { type: Object, default: () => AD_DATA_ACTIONS },
  storageKey: { type: String, required: true },
})

const activeTab = ref(props.levels[0]?.key || 'account')

function onTabChange(key) {
  activeTab.value = key
}
</script>