import { buildColumns } from '../adDataFields'

// 腾讯媒体报表 + 投放报表 指标字段（公共尾部）
const reportFields = [
  'mkt_total_cost', 'mkt_acquisition_cost', 'mkt_total_fodder_impression', 'mkt_thousand_display_price',
  'mkt_total_fodder_click', 'mkt_cpc', 'mkt_ctr', 'mkt_mini_game_register_users_cnt', 'mkt_mini_game_register_cost',
  'mkt_mini_game_paying_users_d1', 'mkt_mini_game_paying_amount_d1', 'mkt_mini_game_first_day_paying_roi',
  'mkt_mini_game_paying_amount_d7', 'mkt_mini_game_pay_d7_roi', 'mkt_mini_game_paying_amount_d30', 'mkt_mini_game_pay_d30_roi',
  'mkt_mini_game_bf_uv', 'mkt_mini_game_bf_cost', 'mkt_mini_game_bf_purchase_uv', 'mkt_mini_game_bf_purchase_amount',
  'mkt_mini_game_bf_purchase_roi', 'mkt_mini_game_first_day_ad_monetization_users', 'mkt_income_pv_24h_pla',
  'mkt_mini_game_first_day_ad_paying_cost', 'mkt_mini_game_first_day_ad_monetization_amount', 'mkt_mini_game_first_day_ad_paying_arpu',
  'mkt_mini_game_income_roi_1', 'mkt_minigame_3d_income_uv', 'mkt_minigame_3d_income_count', 'mkt_mini_game_ad_monetization_amount_d3',
  'mkt_minigame_3d_income_roi', 'mkt_income_roi_1_24h_pla', 'mkt_mini_game_ad_monetization_amount_d7', 'mkt_minigame_7d_income_roi',
  'mkt_last_absolute_cost', 'mkt_last_diff_cost', 'mkt_last_absolute_impression', 'mkt_last_diff_impression', 'mkt_force_close_rate',
  'total_cost', 'total_active', 'total_reg', 'ratio_reg_cost', 'total_new_num', 'rate_new_num', 'ratio_fee_cost', 'ratio_new_arppu',
  'total_fee_num', 'total_fee_count', 'total_purchase', 'total_roi_profit', 'rate_roi_profit_1', 'rate_new_remain_eq_2', 'ratio_roi_profit_x_2',
  'ratio_role_cost', 'ratio_reg_coin_cost', 'total_fee_100', 'ratio_fee_100_cost', 'total_new_sum', 'mkt_exchange_cash_ecpm', 'rate_role_reg', 'rate_roi_1',
]

const accountFields = [
  'subject_id', 'subject_name', 'mkt_balance',
  ...reportFields,
]

const campaignFields = [
  'site_set_name', 'cpid', 'adgroup_name', 'marketing_sub_goal', 'marketing_sub_goal_name',
  'auto_acquisition_enabled', 'auto_acquisition_enabled_name', 'daily_budget', 'bid_mode', 'bid_mode_name',
  'bid_scene', 'bid_scene_name', 'deep_roi_goal', 'deep_roi_goal_name', 'system_status', 'system_status_name',
  'configured_status', 'configured_status_name', 'is_deleted', 'is_deleted_name', 'optimization_goal', 'optimization_goal_name',
  'subject_id', 'subject_name', 'auto_acquisition_status', 'auto_acquisition_status_name', 'demand_classify_id',
  ...reportFields.slice(0, 1),
  'mkt_mini_game_pay_in_24h_amount_roi', 'mkt_mini_game_pay_in_24h_arrive_coefficient',
  'mkt_purchase_pla_clk_1d_amount_roi', 'mkt_purchase_pla_clk_1d_arrive_coefficient',
  ...reportFields.slice(1),
  'ad_id',
]

const unitFields = []

const creativeFields = [
  'auto_derived_program_creative_switch_name', 'cpid', 'aid', 'dynamic_creative_name', 'configured_status', 'configured_status_name',
  'is_deleted', 'optimization_goal', 'optimization_goal_name', 'subject_id', 'subject_name', 'deep_roi_goal', 'deep_roi_goal_name',
  'adgroup_configured_status', 'adgroup_configured_status_name',
  ...reportFields,
]

export default {
  account: buildColumns(accountFields),
  campaign: buildColumns(campaignFields),
  unit: buildColumns(unitFields),
  creative: buildColumns(creativeFields),
}