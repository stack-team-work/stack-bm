package game

type GamePlatform struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Mark      string `gorm:"size:50;uniqueIndex" json:"mark"`
	Name      string `gorm:"size:50;not null" json:"name"`
	Status    int8   `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GamePlatform) TableName() string {
	return "game_platform"
}
