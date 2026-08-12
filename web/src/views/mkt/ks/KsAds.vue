<template>
  <div>
    <n-tabs type="line" :value="activeTab" @update:value="onTabChange" style="margin-top: -8px">
      <n-tab-pane name="ad" tab="广告模板">
        <AdTemplate />
      </n-tab-pane>
      <n-tab-pane name="audience" tab="定向包模板">
        <AudienceTemplate />
      </n-tab-pane>
      <n-tab-pane name="title" tab="标题包模板">
        <TitleTemplate />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import AdTemplate from './AdTemplate.vue'
import AudienceTemplate from './AudienceTemplate.vue'
import TitleTemplate from './TitleTemplate.vue'

const route = useRoute()
const router = useRouter()

const tabs = ['ad', 'audience', 'title']
const activeTab = computed(() => tabs.includes(route.query.tab) ? route.query.tab : 'ad')

function onTabChange(key) {
  router.replace({ query: { tab: key } })
}
</script>
