package sys

type SysLog struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Level     int    `gorm:"default:1" json:"level"`
	Path      string `gorm:"size:255" json:"path"`
	Username  string `gorm:"size:32;index" json:"username"`
	IP        string `gorm:"size:255" json:"ip"`
	Desc      string `gorm:"type:mediumtext" json:"desc"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (SysLog) TableName() string {
	return "sys_logs"
}
