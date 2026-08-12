package bili

type AudienceTemplate struct {
	BaseTemplate `bson:",inline"`
	TemplateName                      string      `bson:"template_name" json:"template_name"`
	Description                       string      `bson:"description,omitempty" json:"description,omitempty"`
	AllowUserIDs                      []int       `bson:"allow_user_ids,omitempty" json:"allow_user_ids,omitempty"`
	AgeList                           []int       `bson:"age_list,omitempty" json:"age_list,omitempty"`
	GenderList                        []int       `bson:"gender_list,omitempty" json:"gender_list,omitempty"`
	NetworkList                       []int       `bson:"network_list,omitempty" json:"network_list,omitempty"`
	ConvertedUserFilter               int         `bson:"converted_user_filter,omitempty" json:"converted_user_filter,omitempty"`
	PhonePriceList                    []int       `bson:"phone_price_list,omitempty" json:"phone_price_list,omitempty"`
	AreaType                          int         `bson:"area_type,omitempty" json:"area_type,omitempty"`
	AreaList                          []int       `bson:"area_list,omitempty" json:"area_list,omitempty"`
	AreaLevelList                     []int       `bson:"area_level_list,omitempty" json:"area_level_list,omitempty"`
	OsList                            []int       `bson:"os_list,omitempty" json:"os_list,omitempty"`
	InstalledUserFilter               []int       `bson:"installed_user_filter,omitempty" json:"installed_user_filter,omitempty"`
	ArchiveContent                    interface{} `bson:"archive_content,omitempty" json:"archive_content,omitempty"`
	ProfessionInterestCrowdPackIDList []int       `bson:"profession_interest_crowd_pack_id_list,omitempty" json:"profession_interest_crowd_pack_id_list,omitempty"`
	VideoSecondPartitionList          []int       `bson:"video_second_partition_list,omitempty" json:"video_second_partition_list,omitempty"`
	ArcTagInterest                    []int       `bson:"arc_tag_interest,omitempty" json:"arc_tag_interest,omitempty"`
	AppCategoryList                   []int       `bson:"app_category_list,omitempty" json:"app_category_list,omitempty"`
	CrowdPack                         []int       `bson:"crowd_pack,omitempty" json:"crowd_pack,omitempty"`
	IntelligentMass                   []int       `bson:"intelligent_mass,omitempty" json:"intelligent_mass,omitempty"`
}

func (AudienceTemplate) CollectionName() string { return "bili_audience_template" }
