package sys

type SysTag struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Type      int8   `gorm:"column:type;not null" json:"type"`
	Name      string `gorm:"size:100;not null" json:"name"`
	Remark    string `gorm:"size:255" json:"remark"`
	AdminID   int    `gorm:"column:admin_id;not null" json:"admin_id"`
	Status    int8   `gorm:"default:1" json:"status"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (SysTag) TableName() string { return "sys_tags" }
