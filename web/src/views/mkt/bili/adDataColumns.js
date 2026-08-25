import { buildColumns } from '../adDataFields'

// 媒体报表 + 投放报表 指标字段（各层级公共尾部）
const reportFields = [
  'mkt_total_cost', 'mkt_total_fodder_click', 'mkt_total_fodder_app_addiction', 'mkt_total_register',
  'mkt_mini_game_first_day_ad_monetization_amount', 'mkt_mini_game_income_roi_1', 'mkt_mini_game_ad_monetization_amount_d7', 'mkt_minigame_7d_income_roi',
  'mkt_total_fodder_impression', 'mkt_thousand_impression_cost', 'mkt_minigame_24h_income_roi', 'mkt_minigame_14d_income_roi',
  'mkt_minigame_30d_income_roi', 'mkt_minigame_24h_mix_income_roi', 'mkt_minigame_1d_mix_income_roi', 'mkt_minigame_7d_mix_income_roi',
  'mkt_minigame_14d_mix_income_roi', 'mkt_minigame_30d_mix_income_roi',
  'total_cost', 'total_active', 'total_reg', 'ratio_reg_cost', 'total_new_num', 'rate_new_num', 'ratio_fee_cost', 'ratio_new_arppu',
  'total_fee_num', 'total_fee_count', 'total_purchase', 'total_roi_profit', 'rate_roi_profit_1', 'ratio_role_cost', 'total_fee_100',
  'ratio_fee_100_cost', 'ratio_reg_coin_cost', 'total_new_sum', 'rate_role_reg', 'rate_day_1_hybrid_revenue_roi',
  'rate_hour_24_hybrid_revenue_roi', 'rate_roi_1', 'mkt_mini_game_pay_in_24h_amount_roi', 'mkt_mini_game_pay_in_24h_arrive_coefficient',
  'mkt_ratio_cost_click', 'mkt_ctr', 'mkt_mini_game_register_users_cnt', 'mkt_mini_game_register_users_cnt_rate', 'mkt_mini_game_register_cost',
  'mkt_mini_game_pay_in_24h_ratio', 'mkt_mini_game_pay_in_24h_count', 'mkt_mini_game_pay_in_7d_ratio', 'mkt_mini_game_pay_in_7d_count',
  'mkt_mini_game_paying_amount_d7_roi', 'mkt_mini_game_pay_in_14d_ratio', 'mkt_mini_game_pay_in_14d_count', 'mkt_mini_game_paying_amount_d14_roi',
  'mkt_mini_game_pay_in_30d_ratio', 'mkt_mini_game_pay_in_30d_count', 'mkt_mini_game_paying_amount_d30_roi',
]

const accountFields = [
  'created_at', 'account_id', 'uid', 'channel_id', 'sys_user_id', 'subject_id', 'budget', 'account_name', 'mkt_balance',
  ...reportFields,
]

const campaignFields = [
  'created_at', 'channel_id', 'account_id', 'uid', 'sys_user_id', 'sys_sub_dep_id', 'pid', 'app_id', 'ad_id', 'subject_id',
  'cpid', 'campaign_name', 'campaign_status', 'promotion_purpose_type', 'ad_type', 'support_auto', 'budget_limit_type', 'budget', 'is_deleted',
  ...reportFields,
]

const unitFields = [
  'created_at', 'channel_id', 'account_id', 'uid', 'sys_user_id', 'sys_sub_dep_id', 'pid', 'app_id', 'ad_id', 'subject_id', 'aid',
  'cpid', 'campaign_name', 'campaign_status', 'campaign_budget', 'unit_name', 'unit_status', 'promotion_content_type', 'unit_budget',
  'speed_mode', 'base_target', 'base_bid', 'cpa_target', 'cpa_bid', 'deep_cpa_target', 'deep_cpa_bid', 'is_no_bid', 'is_deleted',
  'demand_classify_id', 'demand_routine_id',
  ...reportFields,
]

const creativeFields = [
  'created_at', 'channel_id', 'account_id', 'uid', 'sys_user_id', 'sys_sub_dep_id', 'pid', 'app_id', 'ad_id', 'subject_id', 'aid',
  'cpid', 'campaign_name', 'campaign_status', 'campaign_budget', 'unit_name', 'unit_status', 'promotion_content_type', 'unit_budget',
  'speed_mode', 'base_target', 'base_bid', 'cpa_target', 'cpa_bid', 'deep_cpa_target', 'deep_cpa_bid', 'is_no_bid', 'is_deleted',
  'cid', 'creative_status', 'creative_name',
  ...reportFields,
]

export default {
  account: buildColumns(accountFields),
  campaign: buildColumns(campaignFields),
  unit: buildColumns(unitFields),
  creative: buildColumns(creativeFields),
}