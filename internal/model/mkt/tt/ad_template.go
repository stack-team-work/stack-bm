package tt

// AdTemplate 头条广告模板。
// 说明：头条"广告模板"在源项目实际为项目模板（TemplateProjectV2Service，
// 原集合 tt_template_project）。新项目对齐 b站/快手命名，统一为 tt_ad_template。
type AdTemplate struct {
	BaseTemplate              `bson:",inline"`
	TemplateName              string      `bson:"template_name" json:"template_name"`
	AllowUserIDs              []int       `bson:"allow_user_ids,omitempty" json:"allow_user_ids,omitempty"`
	AppID                     []int       `bson:"app_id,omitempty" json:"app_id,omitempty"`
	LandingType               string      `bson:"landing_type,omitempty" json:"landing_type,omitempty"`
	AppPromotionType          string      `bson:"app_promotion_type,omitempty" json:"app_promotion_type,omitempty"`
	MicroPromotionType        string      `bson:"micro_promotion_type,omitempty" json:"micro_promotion_type,omitempty"`
	DeliveryMode              string      `bson:"delivery_mode,omitempty" json:"delivery_mode,omitempty"`
	MarketingGoal             string      `bson:"marketing_goal,omitempty" json:"marketing_goal,omitempty"`
	AdType                    string      `bson:"ad_type,omitempty" json:"ad_type,omitempty"`
	DownloadType              string      `bson:"download_type,omitempty" json:"download_type,omitempty"`
	DownloadMode              string      `bson:"download_mode,omitempty" json:"download_mode,omitempty"`
	InventoryCatalog          string      `bson:"inventory_catalog,omitempty" json:"inventory_catalog,omitempty"`
	InventoryType             []string    `bson:"inventory_type,omitempty" json:"inventory_type,omitempty"`
	UnionVideoType            string      `bson:"union_video_type,omitempty" json:"union_video_type,omitempty"`
	MaterialsType             string      `bson:"materials_type,omitempty" json:"materials_type,omitempty"`
	BidType                   string      `bson:"bid_type,omitempty" json:"bid_type,omitempty"`
	ExternalAction            string      `bson:"external_action,omitempty" json:"external_action,omitempty"`
	DeepExternalAction        string      `bson:"deep_external_action,omitempty" json:"deep_external_action,omitempty"`
	DeepBidType               string      `bson:"deep_bid_type,omitempty" json:"deep_bid_type,omitempty"`
	BudgetOptimizeSwitch      string      `bson:"budget_optimize_switch,omitempty" json:"budget_optimize_switch,omitempty"`
	CustomBidType             string      `bson:"custom_bid_type,omitempty" json:"custom_bid_type,omitempty"`
	CpaBid                    float64     `bson:"cpa_bid,omitempty" json:"cpa_bid,omitempty"`
	MinBid                    float64     `bson:"min_bid,omitempty" json:"min_bid,omitempty"`
	MaxBid                    float64     `bson:"max_bid,omitempty" json:"max_bid,omitempty"`
	DeepCpaBid                float64     `bson:"deep_cpabid,omitempty" json:"deep_cpabid,omitempty"`
	MinDeepCpaBid             float64     `bson:"min_deep_cpabid,omitempty" json:"min_deep_cpabid,omitempty"`
	MaxDeepCpaBid             float64     `bson:"max_deep_cpabid,omitempty" json:"max_deep_cpabid,omitempty"`
	RoiGoal                   float64     `bson:"roi_goal,omitempty" json:"roi_goal,omitempty"`
	MinRoiGoal                float64     `bson:"min_roi_goal,omitempty" json:"min_roi_goal,omitempty"`
	MaxRoiGoal                float64     `bson:"max_roi_goal,omitempty" json:"max_roi_goal,omitempty"`
	BudgetMode                string      `bson:"budget_mode,omitempty" json:"budget_mode,omitempty"`
	ProjectBudget             float64     `bson:"project_budget,omitempty" json:"project_budget,omitempty"`
	AdBudget                  float64     `bson:"ad_budget,omitempty" json:"ad_budget,omitempty"`
	ProjectCustom             string      `bson:"project_custom,omitempty" json:"project_custom,omitempty"`
	ProjectCpaBid             float64     `bson:"project_cpa_bid,omitempty" json:"project_cpa_bid,omitempty"`
	ScheduleType              string      `bson:"schedule_type,omitempty" json:"schedule_type,omitempty"`
	StartTime                 string      `bson:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime                   string      `bson:"end_time,omitempty" json:"end_time,omitempty"`
	ScheduleTimeType          int         `bson:"schedule_time_type,omitempty" json:"schedule_time_type,omitempty"`
	ScheduleTime              string      `bson:"schedule_time,omitempty" json:"schedule_time,omitempty"`
	IsFeedAndFavSee           int         `bson:"is_feed_and_fav_see,omitempty" json:"is_feed_and_fav_see,omitempty"`
	AnchorRelatedType         string      `bson:"anchor_related_type,omitempty" json:"anchor_related_type,omitempty"`
	AnchorTitle               string      `bson:"anchor_title,omitempty" json:"anchor_title,omitempty"`
	AppTags                   []string    `bson:"app_tags,omitempty" json:"app_tags,omitempty"`
	GuideText                 string      `bson:"guide_text,omitempty" json:"guide_text,omitempty"`
	GameDescription           string      `bson:"game_description,omitempty" json:"game_description,omitempty"`
	GameCharatoristic         string      `bson:"game_charatoristic,omitempty" json:"game_charatoristic,omitempty"`
	HeadImageList             []int       `bson:"head_image_list,omitempty" json:"head_image_list,omitempty"`
	AppImages                 []int       `bson:"app_images,omitempty" json:"app_images,omitempty"`
	Title                     string      `bson:"title,omitempty" json:"title,omitempty"`
	SellingPoints             []string    `bson:"selling_points,omitempty" json:"selling_points,omitempty"`
	CallToActionButtons       []string    `bson:"call_to_action_buttons,omitempty" json:"call_to_action_buttons,omitempty"`
	Source                    string      `bson:"source,omitempty" json:"source,omitempty"`
	YuntuCategoryID           int         `bson:"yuntu_category_id,omitempty" json:"yuntu_category_id,omitempty"`
	BrandNameID               string      `bson:"brand_name_id,omitempty" json:"brand_name_id,omitempty"`
	TextAbstractList          interface{} `bson:"text_abstract_list,omitempty" json:"text_abstract_list,omitempty"`
	Keywords                  interface{} `bson:"keywords,omitempty" json:"keywords,omitempty"`
	AutoExtendTraffic         string      `bson:"auto_extend_traffic,omitempty" json:"auto_extend_traffic,omitempty"`
	PromotionType             string      `bson:"promotion_type,omitempty" json:"promotion_type,omitempty"`
	DouyinJuniorSwitch        int         `bson:"douyin_junior_switch,omitempty" json:"douyin_junior_switch,omitempty"`
	GuideVideoIDSwitch        int         `bson:"guide_video_id_switch,omitempty" json:"guide_video_id_switch,omitempty"`
	GuideVideoIDFodderID      int         `bson:"guide_video_id_fodder_id,omitempty" json:"guide_video_id_fodder_id,omitempty"`
	MktAdType                 int         `bson:"mkt_ad_type,omitempty" json:"mkt_ad_type,omitempty"`
	OriginalVideoTitle        int         `bson:"original_video_title,omitempty" json:"original_video_title,omitempty"`
	MaterialSource            int         `bson:"material_source,omitempty" json:"material_source,omitempty"`
	AudienceExtend            int         `bson:"audience_extend,omitempty" json:"audience_extend,omitempty"`
	IntelligentGeneration     int         `bson:"intelligent_generation,omitempty" json:"intelligent_generation,omitempty"`
	AigcDynamicCreativeSwitch int         `bson:"aigc_dynamic_creative_switch,omitempty" json:"aigc_dynamic_creative_switch,omitempty"`
	IsCommentSwitch           int         `bson:"is_comment_switch,omitempty" json:"is_comment_switch,omitempty"`
}

func (AdTemplate) CollectionName() string { return "tt_ad_template" }
