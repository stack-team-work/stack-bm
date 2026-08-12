package game

type GameAppTemplate struct {
	ID             uint   `gorm:"primarykey" json:"id"`
	Name           string `gorm:"size:50;not null" json:"name"`
	IsOpenRealname int8   `gorm:"column:is_open_realname;default:1" json:"is_open_realname"`
	IsOpenRegister int8  `gorm:"column:is_open_register;default:1" json:"is_open_register"`
	IsOpenCharge   int8  `gorm:"column:is_open_charge;default:1" json:"is_open_charge"`
	IsAlertEmail   int8  `gorm:"column:is_alert_email;default:1" json:"is_alert_email"`
	IsAlertPhone   int8  `gorm:"column:is_alert_phone;default:1" json:"is_alert_phone"`
	IsAlertAuth    int8  `gorm:"column:is_alert_auth;default:1" json:"is_alert_auth"`
	IsOpenFloat    int8  `gorm:"column:is_open_float;default:1" json:"is_open_float"`
	AllowAge       int   `gorm:"column:allow_age;default:18" json:"allow_age"`
	Status         int8  `gorm:"default:1" json:"status"`
	AdminID        int   `gorm:"column:admin_id;default:0" json:"admin_id"`
	PrivacyURL     string `gorm:"column:privacy_url;size:255" json:"privacy_url"`
	AgreementURL   string `gorm:"column:agreement_url;size:255" json:"agreement_url"`
	CreatedAt      int64 `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      int64 `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameAppTemplate) TableName() string {
	return "game_app_template"
}
