package sys

type SysColumn struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	ReportType    int8   `gorm:"column:report_type;default:1" json:"report_type"`
	IndicatorType int8   `gorm:"column:indicator_type;default:1" json:"indicator_type"`
	Name          string `gorm:"size:50;not null" json:"name"`
	Mark          string `gorm:"size:50" json:"mark"`
	Field         string `gorm:"size:50;not null" json:"field"`
	Default       int8   `gorm:"default:0" json:"default"`
	Status        int8   `gorm:"default:1" json:"status"`
	AdminID       int    `gorm:"column:admin_id;not null" json:"admin_id"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (SysColumn) TableName() string { return "sys_columns" }
