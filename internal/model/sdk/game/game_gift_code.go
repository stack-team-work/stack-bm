package game

type GameGiftCode struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Code      string `gorm:"size:16;not null;uniqueIndex:uk_code_gift" json:"code"`
	GiftID    int    `gorm:"column:gift_id;default:0;uniqueIndex:uk_code_gift" json:"gift_id"`
	Status    int8   `gorm:"default:0" json:"status"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameGiftCode) TableName() string { return "game_gift_code" }
