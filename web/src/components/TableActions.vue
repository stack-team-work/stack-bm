<template>
  <n-space :size="4" align="center">
    <n-button v-if="edit" size="tiny" secondary @click="edit(row)">
      <template #icon><n-icon :size="14"><create-outline /></n-icon></template>
      编辑
    </n-button>
    <slot name="extra" :row="row" />
    <n-popconfirm v-if="remove" :positive-text="positiveText" negative-text="取消" @positive-click="remove(row)">
      <template #trigger>
        <n-button size="tiny" type="error" secondary>
          <template #icon><n-icon :size="14"><trash-outline /></n-icon></template>
          删除
        </n-button>
      </template>
      {{ confirmText }}
    </n-popconfirm>
  </n-space>
</template>

<script setup>
import { NButton, NSpace, NIcon, NPopconfirm } from 'naive-ui'
import { CreateOutline, TrashOutline } from '@vicons/ionicons5'

defineProps({
  row: { type: Object, required: true },
  edit: { type: Function, default: null },
  remove: { type: Function, default: null },
  confirmText: { type: String, default: '确定删除该记录吗？' },
  positiveText: { type: String, default: '删除' },
})
</script>
