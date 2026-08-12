<template>
  <n-config-provider :theme-overrides="themeOverrides">
    <n-message-provider>
      <router-view />
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
const themeOverrides = {
  common: {
    borderRadius: '6px',
  },
  Button: {
    borderRadiusTiny: '6px',
    borderRadiusSmall: '6px',
    borderRadiusMedium: '6px',
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
}
</script>
