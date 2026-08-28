所有功能的新增/编辑表单去掉状态开关：新增固定默认启用，编辑不改状态，状态只在列表操作

探索结论：全站共 27 个视图的弹窗/表单含状态控件（全部是 `n-form-item path="status"` + n-switch），且每个列表页都已有状态开关列。改动统一为：删除表单中的状态控件；保留 formData 默认 `status: 1`（新增提交即默认开启）；保留 handleEdit 里 `formData.status = row.status`（编辑提交原值、不改状态）——后端无需改动，列表开关行为不变。

═══ 批量处理 25 个文件（脚本正则删除状态表单块，含 n-grid-item 包裹形态）═══

bm/sys：SysAdmin、SysAdminGroup、SysColumn（保留其 default 开关，仅删 status）、SysTag
bm/feishu：FeishuApp、FeishuChat、FeishuUser
mkt/media：Media、MediaDep、MediaManager、MediaAccount、MediaSub、MediaAgent、MediaSubject、MediaApplication
sdk/game：Game、GameAppForm（独立表单页，删状态项后新增仍默认 1、编辑仍回填原状态）、GameAppTemplate、GameCp、GameGift、GamePlatform、GameTag、GameVariable、GameVoucher
sdk/pay：PayMerchant

═══ 两个特例单独处理 ═══

1. SysMenu.vue：删表单状态块，同时把列表里 readonly 的状态开关改为可操作（onUpdateValue → updateMenu({...row, status}) + 成功提示），否则删掉表单字段后菜单状态将无处修改。
2. GameGiftCode.vue：不改动。其 status 语义是「已使用/未使用」（新增默认 0），不属于"启用/禁用"范畴，保持现状。

═══ 验证 ═══
· 批处理后校验：25 个文件中不再有 path="status" 的表单项（搜索栏状态下拉保留不动）
· cd web && npm run build 通过
· 纯前端改动，后端无需重启