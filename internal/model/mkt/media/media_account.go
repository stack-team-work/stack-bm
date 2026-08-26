package media

type MediaAccount struct {
	ID                     uint    `gorm:"primarykey" json:"id"`
	Name                   string  `gorm:"size:50;not null;default:''" json:"name"`
	AgentID                int     `gorm:"column:agent_id;default:0" json:"agent_id"`
	MediaSubID             int     `gorm:"column:media_sub_id;default:0" json:"media_sub_id"`
	AdminID                int     `gorm:"column:admin_id;default:0" json:"admin_id"`
	Username               string  `gorm:"size:50;not null;default:''" json:"username"`
	SubjectID              int     `gorm:"column:subject_id;default:0" json:"subject_id"`
	UID                    string  `gorm:"size:50;not null" json:"uid"`
	Rebate                 float64 `gorm:"type:decimal(16,4);default:0" json:"rebate"`
	Balance                float64 `gorm:"type:decimal(16,4);default:0" json:"balance"`
	Status                 int8    `gorm:"default:1" json:"status"`
	UseType                int8    `gorm:"column:use_type;default:1" json:"use_type"`
	MediaManagerManagerID  int     `gorm:"column:media_manager_manager_id;default:0" json:"media_manager_manager_id"`
	UpdatedAt              int64   `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt              int64   `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (MediaAccount) TableName() string {
	return "media_accounts"
}