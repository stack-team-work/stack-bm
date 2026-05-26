package game

type GameCp struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:64" json:"name"`
	Mark      string `gorm:"size:64" json:"mark"`
	Phone     string `gorm:"size:16" json:"phone"`
	Addr      string `gorm:"size:255" json:"addr"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	Status    int8   `gorm:"default:1" json:"status"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameCp) TableName() string {
	return "game_cp"
}
