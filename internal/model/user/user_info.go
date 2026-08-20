package user

type UserInfo struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	ClientID    string `gorm:"column:client_id;size:50" json:"client_id"`
	UserID      string `gorm:"column:user_id;size:50" json:"user_id"`
	Pid         int    `gorm:"column:pid" json:"pid"`
	AppID       int    `gorm:"column:app_id" json:"app_id"`
	Account     string `gorm:"column:account;size:50" json:"account"`
	Nickname    string `gorm:"column:nickname;size:50" json:"nickname"`
	Phone       string `gorm:"column:phone;size:255" json:"phone"`
	AuthType    int    `gorm:"column:auth_type" json:"auth_type"`
	IsVerify    int8   `gorm:"column:is_verify" json:"is_verify"`
	IsCharge    int8   `gorm:"column:is_charge" json:"is_charge"`
	IsLogin     int8   `gorm:"column:is_login" json:"is_login"`
	VipLevel    int    `gorm:"column:vip_level" json:"vip_level"`
	Coin        int    `gorm:"column:coin" json:"coin"`
	RegFrom     int8   `gorm:"column:reg_from" json:"reg_from"`
	MediaID     int    `gorm:"column:media_id" json:"media_id"`
	MediaSubID  int    `gorm:"column:media_sub_id" json:"media_sub_id"`
	IP          string `gorm:"column:ip;size:64" json:"ip"`
	Province    string `gorm:"column:province;size:50" json:"province"`
	City        string `gorm:"column:city;size:50" json:"city"`
	OS          int8   `gorm:"column:os" json:"os"`
	Unixtime    int64  `gorm:"column:unixtime" json:"unixtime"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (UserInfo) TableName() string { return "user_info" }