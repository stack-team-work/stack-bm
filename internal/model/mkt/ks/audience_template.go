package ks

type AudienceTemplate struct {
	BaseTemplate `bson:",inline"`
	TemplateName              string      `bson:"template_name" json:"template_name"`
	Description               string      `bson:"description,omitempty" json:"description,omitempty"`
	AllowUserIDs              []int       `bson:"allow_user_ids,omitempty" json:"allow_user_ids,omitempty"`
	TargetType                int         `bson:"target_type,omitempty" json:"target_type,omitempty"`
	IntelliExtendOption       int         `bson:"intelli_extend_option,omitempty" json:"intelli_extend_option,omitempty"`
	Network                   int         `bson:"network,omitempty" json:"network,omitempty"`
	Gender                    int         `bson:"gender,omitempty" json:"gender,omitempty"`
	Operators                 []int       `bson:"operators,omitempty" json:"operators,omitempty"`
	DisableInstalledAppSwitch int         `bson:"disable_installed_app_switch,omitempty" json:"disable_installed_app_switch,omitempty"`
	FilterConvertedLevel      int         `bson:"filter_converted_level,omitempty" json:"filter_converted_level,omitempty"`
	FilterTimeRange           int         `bson:"filter_time_range,omitempty" json:"filter_time_range,omitempty"`
	PlatformOs                int         `bson:"platform_os,omitempty" json:"platform_os,omitempty"`
	DevicePrice               []int       `bson:"device_price,omitempty" json:"device_price,omitempty"`
	DeviceBrandIds            []int       `bson:"device_brand_ids,omitempty" json:"device_brand_ids,omitempty"`
	AndroidOsv                int         `bson:"android_osv,omitempty" json:"android_osv,omitempty"`
	IosOsv                    int         `bson:"ios_osv,omitempty" json:"ios_osv,omitempty"`
	HarmonyOsv                int         `bson:"harmony_osv,omitempty" json:"harmony_osv,omitempty"`
	Population                []int       `bson:"population,omitempty" json:"population,omitempty"`
	ExcludePopulation         []int       `bson:"exclude_population,omitempty" json:"exclude_population,omitempty"`
	SeedPopulation            []int       `bson:"seed_population,omitempty" json:"seed_population,omitempty"`
	AutoPopulation            int         `bson:"auto_population,omitempty" json:"auto_population,omitempty"`
	SharedUser                int         `bson:"shared_user,omitempty" json:"shared_user,omitempty"`
	MediaSourceType           int         `bson:"media_source_type,omitempty" json:"media_source_type,omitempty"`
	Media                     []int       `bson:"media,omitempty" json:"media,omitempty"`
	AppInterestIds            []int       `bson:"app_interest_ids,omitempty" json:"app_interest_ids,omitempty"`
	AppIds                    []int       `bson:"app_ids,omitempty" json:"app_ids,omitempty"`
	Region                    interface{} `bson:"region,omitempty" json:"region,omitempty"`
	UserType                  int         `bson:"user_type,omitempty" json:"user_type,omitempty"`
	BehaviorType              int         `bson:"behavior_type,omitempty" json:"behavior_type,omitempty"`
	Behavior                  interface{} `bson:"behavior,omitempty" json:"behavior,omitempty"`
	Interest                  interface{} `bson:"interest,omitempty" json:"interest,omitempty"`
	Celebrity                 []int       `bson:"celebrity,omitempty" json:"celebrity,omitempty"`
}

func (AudienceTemplate) CollectionName() string { return "ks_audience_template" }
