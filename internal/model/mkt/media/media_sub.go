package media

type MediaSub struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	MediaID   int    `gorm:"column:media_id;not null" json:"media_id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Mark      string `gorm:"size:255" json:"mark"`
	Status    int8   `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (MediaSub) TableName() string {
	return "media_sub"
}
