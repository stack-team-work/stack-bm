package game

type GameGiftUserCode struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Code       string `gorm:"size:11;not null" json:"code"`
	GiftID     int    `gorm:"column:gift_id;not null" json:"gift_id"`
	UserID     string `gorm:"column:user_id;size:50;not null" json:"user_id"`
	RoleID     string `gorm:"column:role_id;size:50;not null" json:"role_id"`
	RoleName   string `gorm:"column:role_name;size:50;not null" json:"role_name"`
	ServerID   string `gorm:"column:server_id;size:50;not null" json:"server_id"`
	ServerName string `gorm:"column:server_name;size:50" json:"server_name"`
	AppID      int    `gorm:"column:app_id;not null" json:"app_id"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameGiftUserCode) TableName() string { return "game_gift_user_code" }
