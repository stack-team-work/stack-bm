package ks

type AdTemplate struct {
	BaseTemplate `bson:",inline"`
	TemplateName           string        `bson:"template_name" json:"template_name"`
	AllowUserIDs           []int         `bson:"allow_user_ids,omitempty" json:"allow_user_ids,omitempty"`
	AppID                  []int         `bson:"app_id,omitempty" json:"app_id,omitempty"`
	MarketTarget           int           `bson:"market_target,omitempty" json:"market_target,omitempty"`
	Type                   int           `bson:"type,omitempty" json:"type,omitempty"`
	AdType                 int           `bson:"ad_type,omitempty" json:"ad_type,omitempty"`
	AutoManage             int           `bson:"auto_manage,omitempty" json:"auto_manage,omitempty"`
	PeriodicDeliveryType   int           `bson:"periodic_delivery_type,omitempty" json:"periodic_delivery_type,omitempty"`
	RangeBudget            float64       `bson:"range_budget,omitempty" json:"range_budget,omitempty"`
	ContinuePeriodType     int           `bson:"continue_period_type,omitempty" json:"continue_period_type,omitempty"`
	PeriodicDays           int           `bson:"periodic_days,omitempty" json:"periodic_days,omitempty"`
	AutoPhotoScope         int           `bson:"auto_photo_scope,omitempty" json:"auto_photo_scope,omitempty"`
	BidType                int           `bson:"bid_type,omitempty" json:"bid_type,omitempty"`
	BudgetType             int           `bson:"budget_type,omitempty" json:"budget_type,omitempty"`
	DayBudget              float64       `bson:"day_budget,omitempty" json:"day_budget,omitempty"`
	DayBudgetSchedule      interface{}   `bson:"day_budget_schedule,omitempty" json:"day_budget_schedule,omitempty"`
	SceneID                []int         `bson:"scene_id,omitempty" json:"scene_id,omitempty"`
	UnitCreativeType       int           `bson:"unit_creative_type,omitempty" json:"unit_creative_type,omitempty"`
	AssetMining            int           `bson:"asset_mining,omitempty" json:"asset_mining,omitempty"`
	SmartCover             int           `bson:"smart_cover,omitempty" json:"smart_cover,omitempty"`
	MiniType               int           `bson:"mini_type,omitempty" json:"mini_type,omitempty"`
	UnitDateType           int           `bson:"unit_date_type,omitempty" json:"unit_date_type,omitempty"`
	BeginTime              string        `bson:"begin_time,omitempty" json:"begin_time,omitempty"`
	EndTime                string        `bson:"end_time,omitempty" json:"end_time,omitempty"`
	UnitTimeType           int           `bson:"unit_time_type,omitempty" json:"unit_time_type,omitempty"`
	ScheduleTime           string        `bson:"schedule_time,omitempty" json:"schedule_time,omitempty"`
	UnitBudgetType         int           `bson:"unit_budget_type,omitempty" json:"unit_budget_type,omitempty"`
	UnitDayBudget          float64       `bson:"unit_day_budget,omitempty" json:"unit_day_budget,omitempty"`
	UnitDayBudgetSchedule  interface{}   `bson:"unit_day_budget_schedule,omitempty" json:"unit_day_budget_schedule,omitempty"`
	CustomBidType          int           `bson:"custom_bid_type,omitempty" json:"custom_bid_type,omitempty"`
	BaseTarget             int           `bson:"base_target,omitempty" json:"base_target,omitempty"`
	BaseBid                float64       `bson:"base_bid,omitempty" json:"base_bid,omitempty"`
	MinBaseBid             float64       `bson:"min_base_bid,omitempty" json:"min_base_bid,omitempty"`
	MaxBaseBid             float64       `bson:"max_base_bid,omitempty" json:"max_base_bid,omitempty"`
	CpaTarget              int           `bson:"cpa_target,omitempty" json:"cpa_target,omitempty"`
	CpaBid                 float64       `bson:"cpa_bid,omitempty" json:"cpa_bid,omitempty"`
	MinCpaBid              float64       `bson:"min_cpa_bid,omitempty" json:"min_cpa_bid,omitempty"`
	MaxCpaBid              float64       `bson:"max_cpa_bid,omitempty" json:"max_cpa_bid,omitempty"`
	DeepCpaTarget          int           `bson:"deep_cpa_target,omitempty" json:"deep_cpa_target,omitempty"`
	DeepCpaBid             float64       `bson:"deep_cpa_bid,omitempty" json:"deep_cpa_bid,omitempty"`
	MinDeepCpaBid          float64       `bson:"min_deep_cpa_bid,omitempty" json:"min_deep_cpa_bid,omitempty"`
	MaxDeepCpaBid          float64       `bson:"max_deep_cpa_bid,omitempty" json:"max_deep_cpa_bid,omitempty"`
	QuickSearch            int           `bson:"quick_search,omitempty" json:"quick_search,omitempty"`
	TargetExplore          int           `bson:"target_explore,omitempty" json:"target_explore,omitempty"`
	NegativeWordParam      interface{}   `bson:"negative_word_param,omitempty" json:"negative_word_param,omitempty"`
	MicroChangeSwitch      int           `bson:"micro_change_switch,omitempty" json:"micro_change_switch,omitempty"`
	CreativeCategory       int           `bson:"creative_category,omitempty" json:"creative_category,omitempty"`
	CreativeTag            interface{}   `bson:"creative_tag,omitempty" json:"creative_tag,omitempty"`
	ActionBarText          string        `bson:"action_bar_text,omitempty" json:"action_bar_text,omitempty"`
	Recommendation         interface{}   `bson:"recommendation,omitempty" json:"recommendation,omitempty"`
	ExposeTag              interface{}   `bson:"expose_tag,omitempty" json:"expose_tag,omitempty"`
}

func (AdTemplate) CollectionName() string { return "ks_ad_template" }
