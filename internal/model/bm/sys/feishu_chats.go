package sys

type FeishuChat struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	Type          int    `gorm:"not null" json:"type"`
	ChatID        string `gorm:"column:chat_id;size:100;not null" json:"chat_id"`
	DefaultAtList string `gorm:"column:default_at_list;type:text;not null" json:"default_at_list"`
	AtList        string `gorm:"column:at_list;type:text" json:"at_list"`
	AtType        int    `gorm:"column:at_type;not null" json:"at_type"`
	FeishuAppID   int    `gorm:"column:feishu_app_id" json:"feishu_app_id"`
	CallAction    string `gorm:"column:call_action;size:100;not null" json:"call_action"`
	ActionTitle   string `gorm:"column:action_title;size:100" json:"action_title"`
	AdminID       int    `gorm:"column:admin_id;not null" json:"admin_id"`
	Status        int    `gorm:"default:1" json:"status"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (FeishuChat) TableName() string { return "sys_feishu_chats" }
