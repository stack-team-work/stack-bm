package game

type GameGift struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:50;not null" json:"name"`
	GetType   int8   `gorm:"column:get_type;not null" json:"get_type"`
	IsCode    int8   `gorm:"column:is_code;not null" json:"is_code"`
	Type      int    `gorm:"default:1" json:"type"`
	Cond      string `gorm:"column:cond;type:text" json:"cond"`
	Desc      string `gorm:"column:desc;size:255" json:"desc"`
	Status    int8   `gorm:"default:1" json:"status"`
	Stime     int64  `gorm:"not null" json:"stime"`
	Etime     int64  `gorm:"not null" json:"etime"`
	AdminID   int    `gorm:"column:admin_id;not null" json:"admin_id"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameGift) TableName() string { return "game_gift" }
