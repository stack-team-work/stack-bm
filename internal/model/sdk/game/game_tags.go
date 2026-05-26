package game

type GameTag struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Type      int8   `gorm:"not null;default:1" json:"type"`
	Mark      string `gorm:"size:255;not null" json:"mark"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Status    int    `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameTag) TableName() string {
	return "game_tags"
}
