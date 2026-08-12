package tt

type TtTitleMaterial struct {
	Title    string `bson:"title" json:"title"`
	WordList []int  `bson:"word_list,omitempty" json:"word_list,omitempty"`
}

type TitleTemplate struct {
	BaseTemplate   `bson:",inline"`
	TemplateName   string            `bson:"template_name" json:"template_name"`
	TitleMaterials []TtTitleMaterial `bson:"title_materials,omitempty" json:"title_materials,omitempty"`
	TitleNum       int               `bson:"title_num,omitempty" json:"title_num,omitempty"`
}

func (TitleTemplate) CollectionName() string { return "tt_title_template" }
