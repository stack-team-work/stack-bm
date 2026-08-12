package bili

type SmartTitle struct {
	Type         int    `bson:"type,omitempty" json:"type,omitempty"`
	DefaultValue string `bson:"default_value,omitempty" json:"default_value,omitempty"`
}

type TitleMaterial struct {
	Title          string       `bson:"title" json:"title"`
	SmartTitleList []SmartTitle `bson:"smart_title_list,omitempty" json:"smart_title_list,omitempty"`
}

type TitleTemplate struct {
	BaseTemplate `bson:",inline"`
	TemplateName   string          `bson:"template_name" json:"template_name"`
	TitleMaterials []TitleMaterial `bson:"title_materials,omitempty" json:"title_materials,omitempty"`
	TitleNum       int             `bson:"title_num,omitempty" json:"title_num,omitempty"`
	Description    []string        `bson:"description,omitempty" json:"description,omitempty"`
	DescriptionNum int             `bson:"description_num,omitempty" json:"description_num,omitempty"`
}

func (TitleTemplate) CollectionName() string { return "bili_title_template" }
