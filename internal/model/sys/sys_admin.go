package sys

type SysAdmin struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	Username      string `gorm:"size:64;uniqueIndex" json:"username"`
	Password      string `gorm:"size:64" json:"password,omitempty"`
	Salt          string `gorm:"size:64" json:"-"`
	UserType      int    `gorm:"column:user_type;default:1" json:"user_type"`
	GroupID       uint   `gorm:"column:group_id;default:0" json:"group_id"`
	Name          string `gorm:"size:64" json:"name"`
	Phone         string `gorm:"size:16" json:"phone"`
	DepartmentID  string `gorm:"column:department_id;size:64" json:"department_id"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	Expiration    int64  `gorm:"default:0" json:"expiration"`
	LoginNum      int    `gorm:"column:login_num;default:0" json:"login_num"`
	LastLoginTime int64  `gorm:"column:last_login_time;default:0" json:"last_login_time"`
	LastLoginIP   string `gorm:"column:last_login_ip;size:64" json:"last_login_ip"`
	GameAppPermit string `gorm:"column:game_app_permit;type:mediumtext" json:"game_app_permit"`
	GamePermit    string `gorm:"column:game_permit;type:mediumtext" json:"game_permit"`
	IsDeleted     int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	Status        int8   `gorm:"default:1" json:"status"`
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

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

type SysLog struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Level     string `gorm:"size:50" json:"level"`
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
