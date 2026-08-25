import { buildColumns } from '../adDataFields'

// 头条V3 媒体报表 + 投放报表 指标字段（公共尾部）
const reportFields = [
  'rate_roi_profit_1', 'mkt_total_cost', 'mkt_total_fodder_impression', 'mkt_total_fodder_click',
  'mkt_total_fodder_app_addiction', 'mkt_total_register', 'mkt_total_dy_like_num', 'mkt_total_dy_comment_num',
  'mkt_total_dy_share_num', 'mkt_total_dislike_num', 'mkt_total_report_num', 'total_cost', 'total_active', 'total_reg',
  'ratio_reg_cost', 'total_new_num', 'rate_new_num', 'ratio_fee_cost', 'ratio_new_arppu', 'total_fee_num', 'total_fee_count',
  'total_purchase', 'total_roi_profit', 'rate_roi_profit', 'mkt_total_game_pay_count', 'mkt_attribution_active_pay_7d_per_count',
  'mkt_ctr', 'mkt_total_active_pay', 'mkt_total_game_addiction', 'mkt_game_addiction_cost', 'mkt_mini_game_first_day_ad_monetization_amount',
  'mkt_mini_game_ad_monetization_amount_d3', 'mkt_mini_game_ad_monetization_amount_d7', 'mkt_mini_game_income_roi_1', 'mkt_minigame_3d_income_roi',
  'mkt_minigame_7d_income_roi', 'mkt_game_addiction_rate', 'total_reg_behavior_num', 'mkt_total_live_component_click_num',
  'mkt_live_component_click_cost', 'mkt_live_component_click_rate', 'mkt_luban_live_enter_cnt_rate', 'mkt_live_watch_one_minute_count_rate',
  'mkt_luban_live_comment_cnt_rate', 'total_new_sum', 'rate_role_reg', 'mkt_convert_cnt', 'mkt_convert_cnt_cost', 'subject_id',
  'mkt_convert_cnt_rate', 'mkt_thousand_impression_cost', 'mkt_pay_count_cost', 'pay_count_cost', 'mkt_pre_convert_count', 'mkt_pre_convert_cost',
  'mkt_pre_convert_rate', 'mkt_attribution_game_pay_7d_count', 'mkt_attribution_game_in_app_ltv_8days', 'mkt_attribution_game_pay_7d_arppu',
  'total_bm_1day_pay_arppu', 'total_bm_7day_pay_arppu', 'total_behavior_reg_ecpm_count', 'ratio_behavior_ecpm',
  'rate_reg_behavior_ecpm_roi', 'rate_day_behavior_ecpm_roi', 'mkt_out_minigame_monetize_ad_income', 'mkt_out_minigame_monetize_ad_income_rate_roi',
  'mkt_ratio_game_pay_cost', 'mkt_total_click_start_cnt', 'mkt_ratio_cost_click_start', 'mkt_rate_click_start', 'mkt_total_download_finish_cnt',
  'mkt_ratio_cost_download_finish', 'mkt_rate_download_finish', 'mkt_total_install_finish_cnt', 'rate_new_remain_eq_2', 'ratio_roi_profit_x_2',
  'mkt_last_absolute_cost', 'mkt_last_diff_cost', 'mkt_last_absolute_impression', 'mkt_last_diff_impression', 'rate_roi_1',
]

const accountFields = [
  'account_id', 'uid', 'channel_id', 'sys_user_id', 'mkt_balance', 'note', 'budget', 'created_at', 'account_name',
  ...reportFields,
]

const campaignFields = [
  'cpid', 'project_name', 'account_id', 'uid', 'channel_id', 'sys_user_id', 'status', 'opt_status', 'landing_type', 'pricing',
  'inventory_catalog', 'inventory_type', 'platform', 'created_at', 'marketing_goal', 'ad_type', 'delivery_mode', 'deep_bid_type',
  'bid_type', 'bid_speed', 'budget_mode', 'budget', 'is_deleted', 'updated_at', 'roi_goal', 'app_id', 'anchor_id', 'cpa_bid', 'compensation',
  ...reportFields,
]

const unitFields = [
  'aid', 'promotion_name', 'cpid', 'project_name', 'account_id', 'uid', 'channel_id', 'sys_user_id', 'status', 'opt_status',
  'learning_phase', 'budget_mode', 'budget', 'project_status', 'landing_type', 'pricing', 'inventory_catalog', 'inventory_type',
  'platform', 'marketing_goal', 'ad_type', 'delivery_mode', 'deep_bid_type', 'is_deleted',
  'mkt_estimate_ecpm', 'rate_fee_fodder_click', 'rate_roi_profit_1', 'total_fee_100', 'ratio_fee_100_cost', 'rate_fodder_click',
  'app_id', 'project_opt_status', 'project_budget', 'account_budget', 'is_copy_ad', 'anchor_id', 'trusteeship_media_report_id',
  ...reportFields.filter((f) => f !== 'rate_roi_profit_1'),
]

const creativeFields = []

export default {
  account: buildColumns(accountFields),
  campaign: buildColumns(campaignFields),
  unit: buildColumns(unitFields),
  creative: buildColumns(creativeFields),
}