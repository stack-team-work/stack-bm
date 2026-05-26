package media

type MediaManager struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	MediaID       int    `gorm:"column:media_id;not null" json:"media_id"`
	Name          string `gorm:"size:50;not null" json:"name"`
	ApplicationID int    `gorm:"column:application_id;not null" json:"application_id"`
	Account       string `gorm:"size:50;not null" json:"account"`
	AccountID     string `gorm:"column:account_id;size:50;not null" json:"account_id"`
	AccountNum    int    `gorm:"column:account_num;default:0" json:"account_num"`
	AuthStatus    string `gorm:"column:auth_status;size:255;default:0" json:"auth_status"`
	Status        int8   `gorm:"default:1" json:"status"`
	IsDeleted     int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	Remark        string `gorm:"size:255;not null" json:"remark"`
	Extra         string `gorm:"type:text" json:"extra"`
	AdminID       int    `gorm:"column:admin_id;not null" json:"admin_id"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (MediaManager) TableName() string {
	return "media_manager"
}
