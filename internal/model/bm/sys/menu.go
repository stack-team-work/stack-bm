package sys

type SysMenu struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Type      int8   `gorm:"default:1" json:"type"`
	Author    string `gorm:"size:64;not null" json:"author"`
	Name      string `gorm:"size:64;not null" json:"name"`
	Path      string `gorm:"size:100;not null;index" json:"path"`
	Parent    int    `gorm:"default:0" json:"parent"`
	Icon      string `gorm:"size:128" json:"icon"`
	Sort      int    `gorm:"default:0" json:"sort"`
	Status    int8   `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
