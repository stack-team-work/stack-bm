package sys

type SysAdminGroup struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	Mark         string `gorm:"size:11;uniqueIndex" json:"mark"`
	Name         string `gorm:"size:64" json:"name"`
	Description  string `gorm:"size:256" json:"description"`
	MenuPermit   string `gorm:"column:menu_permit;type:mediumtext" json:"menu_permit"`
	ColumnPermit string `gorm:"column:column_permit;type:mediumtext" json:"column_permit"`
	IsDeleted    int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	Status       int8   `gorm:"default:1" json:"status"`
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (SysAdminGroup) TableName() string {
	return "sys_admin_group"
}
