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

// 设计 token（与 src/styles/tokens.css 保持一致，作为单一来源）
const tokens = {
  brand: '#10b981',
  brandHover: '#059669',
  brandPressed: '#047857',
  info: '#2563eb',
  success: '#16a34a',
  warning: '#f59e0b',
  error: '#ef4444',
  border: '#eef0f4',
  bgPage: '#f5f7fa',
  text2: '#646a73',
}

const themeOverrides = {
  common: {
    primaryColor: tokens.brand,
    primaryColorHover: tokens.brandHover,
    primaryColorPressed: tokens.brandPressed,
    primaryColorSuppl: tokens.brand,
    infoColor: tokens.info,
    successColor: tokens.success,
    warningColor: tokens.warning,
    errorColor: tokens.error,
    bodyColor: tokens.bgPage,
    borderRadius: '8px',
    fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif",
  },
  Button: {
    borderRadiusTiny: '4px',
    borderRadiusSmall: '6px',
    borderRadiusMedium: '6px',
    borderRadiusLarge: '8px',
    fontWeight: 500,
  },
  Card: {
    borderRadius: '8px',
  },
  Modal: {
    borderRadius: '12px',
    boxShadow: '0 12px 32px rgba(16, 24, 40, 0.12)',
  },
  Popover: {
    borderRadius: '8px',
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
    thColor: '#f7f8fa',
    thTextColor: tokens.text2,
    thFontWeight: '600',
    borderColor: tokens.border,
    tdColor: '#ffffff',
    tdColorHover: '#f7f8fa',
  },
  Tag: {
    borderRadius: '4px',
  },
}
</script>
