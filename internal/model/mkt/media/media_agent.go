package media

type MediaAgent struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Mark      string `gorm:"size:50;not null" json:"mark"`
	SubjectID int    `gorm:"column:subject_id;not null" json:"subject_id"`
	Name      string `gorm:"size:50;not null" json:"name"`
	Status    int8   `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted" json:"is_deleted"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (MediaAgent) TableName() string {
	return "media_agent"
}
