package ks

type KsTitleMaterial struct {
	Title      string   `bson:"title" json:"title"`
	SmartTitle []string `bson:"smart_title,omitempty" json:"smart_title,omitempty"`
}

type TitleTemplate struct {
	BaseTemplate `bson:",inline"`
	TemplateName  string            `bson:"template_name" json:"template_name"`
	TitleMaterials []KsTitleMaterial `bson:"title_materials,omitempty" json:"title_materials,omitempty"`
	TitleNum       int               `bson:"title_num,omitempty" json:"title_num,omitempty"`
}

func (TitleTemplate) CollectionName() string { return "ks_title_template" }
