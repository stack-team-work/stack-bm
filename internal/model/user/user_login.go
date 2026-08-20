package user

type UserLogin struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	ClientID   string `gorm:"column:client_id;size:50" json:"client_id"`
	UserID     string `gorm:"column:user_id;size:50" json:"user_id"`
	Pid        int    `gorm:"column:pid" json:"pid"`
	AppID      int    `gorm:"column:app_id" json:"app_id"`
	IsReg      int8   `gorm:"column:is_reg" json:"is_reg"`
	AuthType   int8   `gorm:"column:auth_type" json:"auth_type"`
	MediaID    int    `gorm:"column:media_id" json:"media_id"`
	MediaSubID int    `gorm:"column:media_sub_id" json:"media_sub_id"`
	IP         string `gorm:"column:ip;size:64" json:"ip"`
	OS         int8   `gorm:"column:os" json:"os"`
	Unixtime   int64  `gorm:"column:unixtime" json:"unixtime"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (UserLogin) TableName() string { return "user_logins" }