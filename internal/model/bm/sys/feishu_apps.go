package sys

type FeishuApp struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	AppID     string `gorm:"column:app_id;size:100;not null" json:"app_id"`
	AppSecret string `gorm:"column:app_secret;size:100;not null" json:"app_secret"`
	AppName   string `gorm:"column:app_name;size:100" json:"app_name"`
	AdminID   int    `gorm:"column:admin_id;not null" json:"admin_id"`
	Mark      string `gorm:"size:100;not null" json:"mark"`
	Status    int    `gorm:"default:1" json:"status"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (FeishuApp) TableName() string { return "sys_feishu_apps" }
