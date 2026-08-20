package user

type UserActive struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	ClientID   string `gorm:"column:client_id;size:50" json:"client_id"`
	Pid        int    `gorm:"column:pid" json:"pid"`
	AppID      int    `gorm:"column:app_id" json:"app_id"`
	AdID       int    `gorm:"column:ad_id" json:"ad_id"`
	MediaID    int    `gorm:"column:media_id" json:"media_id"`
	MediaSubID int    `gorm:"column:media_sub_id" json:"media_sub_id"`
	PackageName string `gorm:"column:package_name;size:255" json:"package_name"`
	IP         string `gorm:"column:ip;size:64" json:"ip"`
	OS         int8   `gorm:"column:os" json:"os"`
	Unixtime   int64  `gorm:"column:unixtime" json:"unixtime"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (UserActive) TableName() string { return "user_actives" }