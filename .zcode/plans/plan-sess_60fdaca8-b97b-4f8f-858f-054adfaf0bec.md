优化「列设置」弹出面板太挤的问题

背景：B站/快手/腾讯/头条四个广告数据页共用的 `web/src/views/mkt/AdDataTable.vue` 中，「列设置」是一个固定 720px 宽的 `n-popover` 气泡；创意层级最多约 88 个指标，导致复选框非常拥挤。

方案：将该气泡替换为项目通用的居中卡片弹窗（`n-modal preset="card"`），并使用响应式宽度 `min(1100px, calc(100vw - 48px))` —— 大屏时加宽到 1100px 容纳更多指标，小屏也不溢出。仅改动这一个文件：

1. 模板部分：
   - 触发按钮改为 `@click="showColumnModal = true"` 打开弹窗。
   - 原 popover 内部内容原样迁入 `<n-modal v-model:show="showColumnModal" preset="card" title="列设置" style="width: min(1100px, calc(100vw - 48px))">`：搜索框 + 已选计数、可折叠分组复选框网格、底部全选/全不选/确定/重置按钮，逻辑全部保持不变。
2. 脚本部分：新增 `const showColumnModal = ref(false)`；不改任何列选择/搜索/localStorage 保存逻辑。

无后端改动，NModal 在该文件已注册（现有输入弹窗已在用），不需要改 components.d.ts 或其他页面。