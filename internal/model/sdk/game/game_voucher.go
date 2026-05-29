package game

type GameVoucher struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:50;not null" json:"name"`
	UseType   int    `gorm:"column:use_type;default:1" json:"use_type"`
	UseLimit  int8   `gorm:"column:use_limit;not null" json:"use_limit"`
	Total     int    `gorm:"not null" json:"total"`
	TotalFee  int    `gorm:"column:total_fee;not null" json:"total_fee"`
	Desc      string `gorm:"column:desc;size:255;default:''" json:"desc"`
	Stime     int64  `gorm:"not null" json:"stime"`
	Etime     int64  `gorm:"not null" json:"etime"`
	Status    int8   `gorm:"default:1" json:"status"`
	AdminID   int    `gorm:"column:admin_id;not null" json:"admin_id"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameVoucher) TableName() string { return "game_voucher" }
