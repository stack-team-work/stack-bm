package game

type Game struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:64;not null" json:"name"`
	WebName   string `gorm:"column:web_name;size:64;not null" json:"web_name"`
	Icon      string `gorm:"size:64" json:"icon"`
	IsWebShow int8   `gorm:"column:is_web_show;default:0" json:"is_web_show"`
	TypeID    int    `gorm:"column:type_id" json:"type_id"`
	StyleID   int    `gorm:"column:style_id" json:"style_id"`
	CpID      int    `gorm:"column:cp_id;not null" json:"cp_id"`
	ServerURL string `gorm:"column:server_url;size:64" json:"server_url"`
	RoleURL   string `gorm:"column:role_url;size:64" json:"role_url"`
	AuthName  string `gorm:"column:auth_name;size:64" json:"auth_name"`
	Author    string `gorm:"type:mediumtext" json:"author"`
	Mark      string `gorm:"size:64;uniqueIndex" json:"mark"`
	Status    int    `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Game) TableName() string {
	return "game"
}
