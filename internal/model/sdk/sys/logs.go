package sys

type SysLog struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Pid        int    `gorm:"column:pid" json:"pid"`
	AppID      int    `gorm:"column:app_id" json:"app_id"`
	Type       int    `gorm:"default:1" json:"type"`
	Level      int    `gorm:"default:1" json:"level"`
	IP         string `gorm:"column:ip;size:255" json:"ip"`
	Desc       string `gorm:"column:desc;type:mediumtext" json:"desc"`
	CreateTime int64  `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (SysLog) TableName() string {
	return "sys_logs"
}
