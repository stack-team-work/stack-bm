package bili

type AdTemplate struct {
	BaseTemplate `bson:",inline"`
	TemplateName              string        `bson:"template_name" json:"template_name"`
	AllowUserIDs              []int         `bson:"allow_user_ids,omitempty" json:"allow_user_ids,omitempty"`
	AppID                     []int         `bson:"app_id,omitempty" json:"app_id,omitempty"`
	PromotionPurposeType      int           `bson:"promotion_purpose_type,omitempty" json:"promotion_purpose_type,omitempty"`
	AdBudgetLimitType         int           `bson:"ad_budget_limit_type,omitempty" json:"ad_budget_limit_type,omitempty"`
	AdBudget                  float64       `bson:"ad_budget,omitempty" json:"ad_budget,omitempty"`
	AdType                    int           `bson:"ad_type,omitempty" json:"ad_type,omitempty"`
	SupportAuto               int           `bson:"support_auto,omitempty" json:"support_auto,omitempty"`
	PromotionContentType      int           `bson:"promotion_content_type,omitempty" json:"promotion_content_type,omitempty"`
	SpeedMode                 int           `bson:"speed_mode,omitempty" json:"speed_mode,omitempty"`
	UnitDateType              int           `bson:"unit_date_type,omitempty" json:"unit_date_type,omitempty"`
	LaunchBeginDate           string        `bson:"launch_begin_date,omitempty" json:"launch_begin_date,omitempty"`
	LaunchEndDate             string        `bson:"launch_end_date,omitempty" json:"launch_end_date,omitempty"`
	UnitTimeType              int           `bson:"unit_time_type,omitempty" json:"unit_time_type,omitempty"`
	LaunchTime                string        `bson:"launch_time,omitempty" json:"launch_time,omitempty"`
	UnitBudget                float64       `bson:"unit_budget,omitempty" json:"unit_budget,omitempty"`
	IsNoBid                   int           `bson:"is_no_bid,omitempty" json:"is_no_bid,omitempty"`
	BaseTarget                int           `bson:"base_target,omitempty" json:"base_target,omitempty"`
	BaseBid                   float64       `bson:"base_bid,omitempty" json:"base_bid,omitempty"`
	MinBaseBid                float64       `bson:"min_base_bid,omitempty" json:"min_base_bid,omitempty"`
	MaxBaseBid                float64       `bson:"max_base_bid,omitempty" json:"max_base_bid,omitempty"`
	CpaTarget                 int           `bson:"cpa_target,omitempty" json:"cpa_target,omitempty"`
	CpaBid                    float64       `bson:"cpa_bid,omitempty" json:"cpa_bid,omitempty"`
	MinCpaBid                 float64       `bson:"min_cpa_bid,omitempty" json:"min_cpa_bid,omitempty"`
	MaxCpaBid                 float64       `bson:"max_cpa_bid,omitempty" json:"max_cpa_bid,omitempty"`
	DeepCpaTarget             int           `bson:"deep_cpa_target,omitempty" json:"deep_cpa_target,omitempty"`
	DeepCpaBid                float64       `bson:"deep_cpa_bid,omitempty" json:"deep_cpa_bid,omitempty"`
	MinDeepCpaBid             float64       `bson:"min_deep_cpa_bid,omitempty" json:"min_deep_cpa_bid,omitempty"`
	MaxDeepCpaBid             float64       `bson:"max_deep_cpa_bid,omitempty" json:"max_deep_cpa_bid,omitempty"`
	DualBidTwoStageOptimization int         `bson:"dual_bid_two_stage_optimization,omitempty" json:"dual_bid_two_stage_optimization,omitempty"`
	AssistCpaTarget           int           `bson:"assist_cpa_target,omitempty" json:"assist_cpa_target,omitempty"`
	AssistCpaBid              float64       `bson:"assist_cpa_bid,omitempty" json:"assist_cpa_bid,omitempty"`
	IsBiliNative              int           `bson:"is_bili_native,omitempty" json:"is_bili_native,omitempty"`
	ChannelID                 int           `bson:"channel_id,omitempty" json:"channel_id,omitempty"`
	TagList                   []string      `bson:"tag_list,omitempty" json:"tag_list,omitempty"`
	IsSmartMaterial           int           `bson:"is_smart_material,omitempty" json:"is_smart_material,omitempty"`
	CreativeComponents        interface{}   `bson:"creative_components,omitempty" json:"creative_components,omitempty"`
	CustomBidType             string        `bson:"custom_bid_type,omitempty" json:"custom_bid_type,omitempty"`
	MiniGameType              int           `bson:"mini_game_type,omitempty" json:"mini_game_type,omitempty"`
}

func (AdTemplate) CollectionName() string { return "bili_ad_template" }
