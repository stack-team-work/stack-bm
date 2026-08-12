package tt

type WordList struct {
	ID             int    `bson:"id" json:"id"`
	Name           string `bson:"name" json:"name"`
	MaxWordLen     int    `bson:"max_word_len" json:"max_word_len"`
	CreativeWordID int    `bson:"creative_word_id" json:"creative_word_id"`
}

func (WordList) CollectionName() string { return "tt_word_lists" }
