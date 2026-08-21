package media

type MediaApplication struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	MediaID   int    `gorm:"column:media_id;not null" json:"media_id"`
	Name      string `gorm:"size:50;not null" json:"name"`
	AppID     string `gorm:"column:app_id;size:50;not null" json:"app_id"`
	AppSecret string `gorm:"column:app_secret;size:50" json:"app_secret"`
	Status    int8   `gorm:"not null" json:"status"`
	Remark    string `gorm:"size:255" json:"remark"`
	Extra     string `gorm:"type:text" json:"extra"`
	AdminID   int    `gorm:"column:admin_id;not null" json:"admin_id"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (MediaApplication) TableName() string {
	return "media_application"
}
