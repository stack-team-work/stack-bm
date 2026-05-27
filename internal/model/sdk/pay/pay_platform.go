package pay

type PayPlatform struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:64;not null" json:"name"`
	Mark      string `gorm:"size:16;not null" json:"mark"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (PayPlatform) TableName() string {
	return "pay_platform"
}
