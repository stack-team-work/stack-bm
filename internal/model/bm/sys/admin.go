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
