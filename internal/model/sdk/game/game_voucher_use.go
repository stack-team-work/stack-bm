package game

type GameVoucherUse struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	VoucherID  int    `gorm:"column:voucher_id;not null" json:"voucher_id"`
	UserID     string `gorm:"column:user_id;size:50;not null" json:"user_id"`
	AppID      int    `gorm:"column:app_id;not null" json:"app_id"`
	RoleID     string `gorm:"column:role_id;size:50;not null" json:"role_id"`
	RoleName   string `gorm:"column:role_name;size:50;not null" json:"role_name"`
	ServerName string `gorm:"column:server_name;size:50" json:"server_name"`
	ServerID   string `gorm:"column:server_id;size:50;not null" json:"server_id"`
	IsUse      int8   `gorm:"column:is_use;default:0" json:"is_use"`
	Stime      int64  `gorm:"not null" json:"stime"`
	Etime      int64  `gorm:"not null" json:"etime"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameVoucherUse) TableName() string { return "game_voucher_use" }
