package game

type GameVariable struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Key       string `gorm:"size:32;uniqueIndex;not null" json:"key"`
	Value     string `gorm:"type:mediumtext;not null" json:"value"`
	Mark      string `gorm:"size:255" json:"mark"`
	Status    int8   `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (GameVariable) TableName() string {
	return "game_variable"
}
