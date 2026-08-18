package feishu

type FeishuUser struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	AdminID      int    `gorm:"column:admin_id;default:0" json:"admin_id"`
	FeishuUserID string `gorm:"column:feishu_user_id;size:100;not null" json:"feishu_user_id"`
	Status       int    `gorm:"default:1" json:"status"`
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (FeishuUser) TableName() string { return "sys_feishu_users" }
