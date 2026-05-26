package media

type MediaSubject struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Mark      string `gorm:"size:50" json:"mark"`
	Name      string `gorm:"size:50;not null" json:"name"`
	Status    int8   `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (MediaSubject) TableName() string {
	return "media_subject"
}
