package tt

type AudienceTemplate struct {
	BaseTemplate              `bson:",inline"`
	TemplateName              string      `bson:"template_name" json:"template_name"`
	Description               string      `bson:"description,omitempty" json:"description,omitempty"`
	AllowUserIDs              []int       `bson:"allow_user_ids,omitempty" json:"allow_user_ids,omitempty"`
	LandingType               string      `bson:"landing_type,omitempty" json:"landing_type,omitempty"`
	District                  string      `bson:"district,omitempty" json:"district,omitempty"`
	City                      []int       `bson:"city,omitempty" json:"city,omitempty"`
	LocationType              string      `bson:"location_type,omitempty" json:"location_type,omitempty"`
	Gender                    string      `bson:"gender,omitempty" json:"gender,omitempty"`
	Age                       []string    `bson:"age,omitempty" json:"age,omitempty"`
	Career                    []string    `bson:"career,omitempty" json:"career,omitempty"`
	SubjectID                 int         `bson:"subject_id,omitempty" json:"subject_id,omitempty"`
	RetargetingTagsInclude    []string    `bson:"retargeting_tags_include,omitempty" json:"retargeting_tags_include,omitempty"`
	RetargetingTagsExclude    []string    `bson:"retargeting_tags_exclude,omitempty" json:"retargeting_tags_exclude,omitempty"`
	RetargetingKeywords       string      `bson:"retargeting_keywords,omitempty" json:"retargeting_keywords,omitempty"`
	InterestActionMode        string      `bson:"interest_action_mode,omitempty" json:"interest_action_mode,omitempty"`
	ActionScene               []string    `bson:"action_scene,omitempty" json:"action_scene,omitempty"`
	ActionDays                int         `bson:"action_days,omitempty" json:"action_days,omitempty"`
	ActionCategories          interface{} `bson:"action_categories,omitempty" json:"action_categories,omitempty"`
	ActionWords               interface{} `bson:"action_words,omitempty" json:"action_words,omitempty"`
	InterestCategories        interface{} `bson:"interest_categories,omitempty" json:"interest_categories,omitempty"`
	InterestWords             interface{} `bson:"interest_words,omitempty" json:"interest_words,omitempty"`
	AwemeFanBehaviors         []string    `bson:"aweme_fan_behaviors,omitempty" json:"aweme_fan_behaviors,omitempty"`
	AwemeFanTimeScope         string      `bson:"aweme_fan_time_scope,omitempty" json:"aweme_fan_time_scope,omitempty"`
	AwemeFanCategories        interface{} `bson:"aweme_fan_categories,omitempty" json:"aweme_fan_categories,omitempty"`
	AwemeFanAccounts          interface{} `bson:"aweme_fan_accounts,omitempty" json:"aweme_fan_accounts,omitempty"`
	SuperiorPopularityType    string      `bson:"superior_popularity_type,omitempty" json:"superior_popularity_type,omitempty"`
	FlowPackage               interface{} `bson:"flow_package,omitempty" json:"flow_package,omitempty"`
	ExcludeFlowPackage        interface{} `bson:"exclude_flow_package,omitempty" json:"exclude_flow_package,omitempty"`
	Platform                  []string    `bson:"platform,omitempty" json:"platform,omitempty"`
	AndroidOsv                string      `bson:"android_osv,omitempty" json:"android_osv,omitempty"`
	IosOsv                    string      `bson:"ios_osv,omitempty" json:"ios_osv,omitempty"`
	HarmonyOsv                string      `bson:"harmony_osv,omitempty" json:"harmony_osv,omitempty"`
	DeviceType                []string    `bson:"device_type,omitempty" json:"device_type,omitempty"`
	Ac                        []string    `bson:"ac,omitempty" json:"ac,omitempty"`
	Carrier                   []string    `bson:"carrier,omitempty" json:"carrier,omitempty"`
	HideIfExists              string      `bson:"hide_if_exists,omitempty" json:"hide_if_exists,omitempty"`
	HideIfConverted           string      `bson:"hide_if_converted,omitempty" json:"hide_if_converted,omitempty"`
	ConvertedTimeDuration     string      `bson:"converted_time_duration,omitempty" json:"converted_time_duration,omitempty"`
	FilterAwemeAbnormalActive string      `bson:"filter_aweme_abnormal_active,omitempty" json:"filter_aweme_abnormal_active,omitempty"`
	FilterOwnAwemeFans        string      `bson:"filter_own_aweme_fans,omitempty" json:"filter_own_aweme_fans,omitempty"`
	FilterAwemeFansCount      interface{} `bson:"filter_aweme_fans_count,omitempty" json:"filter_aweme_fans_count,omitempty"`
	ActivateType              []string    `bson:"activate_type,omitempty" json:"activate_type,omitempty"`
	DeviceBrand               []string    `bson:"device_brand,omitempty" json:"device_brand,omitempty"`
	LaunchPrice               interface{} `bson:"launch_price,omitempty" json:"launch_price,omitempty"`
	AutoExtendEnabled         string      `bson:"auto_extend_enabled,omitempty" json:"auto_extend_enabled,omitempty"`
	AutoExtendTargets         []string    `bson:"auto_extend_targets,omitempty" json:"auto_extend_targets,omitempty"`
}

func (AudienceTemplate) CollectionName() string { return "tt_audience_template" }
