package pay

type PayMerchant struct {
	ID           uint    `gorm:"primarykey" json:"id"`
	Name         string  `gorm:"size:64;not null" json:"name"`
	ShowName     string  `gorm:"column:show_name;size:64;not null" json:"show_name"`
	Type         int8    `gorm:"default:0" json:"type"`
	PlatformMark int8    `gorm:"column:platform_mark;default:0" json:"platform_mark"`
	Mark         string  `gorm:"size:16;not null" json:"mark"`
	Status       int8    `gorm:"default:1" json:"status"`
	Weight       int     `gorm:"default:0" json:"weight"`
	URL          string  `gorm:"size:50" json:"url"`
	Rate         float64 `gorm:"type:decimal(10,4);default:0" json:"rate"`
	Config       string  `gorm:"column:config;type:mediumtext" json:"config"`
	CreatedAt    int64   `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    int64   `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (PayMerchant) TableName() string {
	return "pay_merchant"
}
