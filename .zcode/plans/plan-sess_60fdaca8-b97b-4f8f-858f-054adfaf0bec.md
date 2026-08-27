广告数据模块按功能目录重构（对齐 PHP 老项目规划，四渠道一次到位）

════ 总体原则 ════

1. 版本目录 v1 是本项目自定义版本号：头条 v1 ↔ 巨量引擎 V3 API、B站 v1 ↔ 当前开放 API、快手/腾讯 v1 ↔ 当前在用版本；后续平台出新版时建 v2 并存。
2. 一个功能一个文件、一个层级一个文件（照搬 bm2 的 Controller/Ad/Campaign + Template/Title 粒度）；四渠道各自独立实现，不抽跨渠道接口。
3. token 每渠道自持（对齐 bm2 的 XxTokenService），不再经共享 oauth.ManagerAuth 解析。
4. 旧参数化路由全部废弃、新旧一次切换，前端 web/src/api/mkt/*.js 同步改 URL。

════ 目标结构（以 B站为例，其余渠道同构差异见后） ════

─ handler（薄 HTTP 层，每层级一个文件，自带路由注册函数）─
internal/handler/mkt/bili/
  base.go                     保留（adminID 助手）
  v1/ad/account.go            账户列表
  v1/ad/campaign.go           计划（第一层级）：list + 开启/暂停/删除/预算/出价 + 批量
  v1/ad/unit.go               单元（第二层级）：list + 收藏/起量/时间 + 批量
  v1/ad/creative.go           创意：list + 开启/暂停/预览
  v1/template/ad.go           广告模板 CRUD
  v1/template/audience.go     定向模板 CRUD
  v1/template/title.go        标题包模板 CRUD

现 tool.go 的 ~75 case 大 switch 按层级消解进各文件。路由形态：
  POST /api/bili/v1/ad/campaign/list、/open、/pause、/set-budget、/batch-status…
（动作 token 与前端 AD_DATA_ACTIONS 现值保持一致 set-budget/set-deep-bid/…，避免前端业务改动）

─ service ─
internal/service/mkt/bili/v1/
  token/token.go              BiliTokenService：本渠道 access_token 解析（读写现有 mkt_account_manager_token 集合 + media_accounts），提供本渠道 api 所需的 (uid, token)
  api/auth.go                 BuildOauthURL / ExchangeToken（现 api/api.go 迁入）
  api/account.go              余额 cash、主体列表；campaign.go 计划列表+状态/预算/出价请求；unit.go；creative.go —— 现 api/sync.go + api/operation.go 按"接口族"拆入
  sync/advertiser.go          广告主同步（现 sync/advertiser.go 迁入）
  sync/balance.go             余额回填
  sync/campaign.go / unit.go / creative.go   各层级拉取入库，导出独立入口函数（如 SyncCampaign(accountID)），全量分页拉取
  ad/campaign.go / unit.go / creative.go    列表查询 + 操作编排（承接现 tool/campaign.go 等业务）；sync/ad_data.go 的 List 骨架并入各层级文件
  template/ad.go / audience.go / title.go   现根级三个模板服务迁入

─ repository（不加 v1：对齐 bm2 中 DAO/模型层不随 Controller/Service 版本化的惯例）─
internal/repository/mkt/bili/
  ad/campaign.go / unit.go / creative.go    该层级 Mongo List/Upsert/按层级ID解析account_id（现单文件 ad_data.go 拆开；统一写 int、读取兼容存量字符串）

─ model 保持现状不动（本就按实体一文件，属数据定义层）─

═ 渠道同构差异 ═
· ks：同 bili，四层级齐全
· tt：仅 account + campaign（v1=巨量V3 项目层级，字段即 project_*）；无模板外多一个 word 词包归 v1/template/word.go
· tc：account + campaign(adgroup)；unit 无数据源不建；无模板 → 只有 v1/ad/*
· 各渠道 api 文件保持各自协议差异（tt/tc/ks Access-Token + code==0；bili Bearer + status==success；ks data 为裸数组、金额厘；tc 元/分转换）

═ 定时同步入口 ═
· 废弃 internal/service/mkt/sync/runner.go 的 channelSyncer 接口分发
· cmd/sync/main.go 保留 --channel/--level/--account 外部 cron 约定不变，内部直接 switch 渠道调用各自的 v1/sync 入口函数（无共享接口）

═ 删除清单 ═
router.go 旧段 /<ch>-ad-data/:level/list、/<ch>-tool/:level/:action、/<ch>-{ad,audience,title}-template/*；
旧文件：handler/mkt/<ch>/{ad_data,tool,*_template}.go、service/mkt/<ch>/{api,sync,tool,模板服务}旧位置、repository/mkt/<ch>/ad_data.go 旧文件、internal/service/mkt/oauth 中 GetAccountContext 相关迁移后冗余部分（media 管家授权回调流程不动）

════ 重构中一并落地的已知问题（避免二次返工） ════
1. 同步全量分页（替代现在各处只拉第 1 页 100 条的静默截断）
2. account_id 类型统一为 int 写入 + 读取兼容存量字符串（修 tt/tc 工具操作解析必失败）
3. tc 空 unit 死页签：前端隐藏（levels 去掉 'unit'）
4. 前端工具操作失败提示 + 防重复提交、勾选跨页残留清理、日期本地格式化（修 UTC 偏移）、重置搜索保留 columns、「创建时间」→「报表日期」文案

════ 前端改动 ════
· web/src/api/mkt/{bili,ks,tt,tc}.js：广告数据 list/tool 与模板 CRUD 的 URL 换成新路径（函数签名/动作 token 不变，views 层组件无需动）
· 前端 Bug 修复集中在 web/src/views/mkt/AdDataTable.vue

════ 实施顺序 ════
1. 先建 B站 v1 全链路 → go build ./... 过 → 作为样式基准
2. 依样铺开 ks / tt / tc（含各自层级与模板差异）
3. 删旧实现与旧路由、清理 oauth 冗余、router 注册新路由
4. 前端 api 层换 URL + AdDataTable.vue 体验修复
5. go build ./... + npm run build 双验证

════ 风险与验收 ════
· 平台真实 API 与线上 Mongo 无法本地联调，本轮以两端编译通过 + 结构验收为准；同步/工具操作的真实调用列为人工联调项
· 建议开工前先把当前工作区未提交的前端改动（列设置弹窗等）提交一版，保证重构 diff 干净可回退