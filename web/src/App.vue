<template>
  <n-config-provider :theme-overrides="themeOverrides" :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <n-dialog-provider>
        <n-global-style />
        <router-view />
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import { zhCN, dateZhCN } from 'naive-ui'

const themeOverrides = {
  common: {
    primaryColor: '#18a058',
    primaryColorHover: '#36ad6a',
    primaryColorPressed: '#0c7a43',
    primaryColorSuppl: '#36ad6a',
    infoColor: '#2080f0',
    successColor: '#18a058',
    warningColor: '#f0a020',
    errorColor: '#d03050',
    borderRadius: '6px',
    fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif",
  },
  Button: {
    borderRadiusTiny: '4px',
    borderRadiusSmall: '4px',
    borderRadiusMedium: '6px',
    borderRadiusLarge: '6px',
    fontWeight: 500,
  },
  Card: {
    borderRadius: '8px',
  },
  Modal: {
    borderRadius: '8px',
  },
  Popover: {
    borderRadius: '6px',
  },
  Input: {
    borderRadius: '6px',
  },
  Select: {
    borderRadius: '6px',
  },
  DatePicker: {
    borderRadius: '6px',
  },
  DataTable: {
    borderRadius: '8px',
  },
  Tag: {
    borderRadius: '4px',
  },
}
</script>
