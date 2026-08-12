package game

type GameApp struct {
	ID             uint   `gorm:"primarykey" json:"id"`
	Pid            int    `gorm:"not null" json:"pid"`
	Name           string `gorm:"size:50;not null" json:"name"`
	PackageName    string `gorm:"column:package_name;size:50" json:"package_name"`
	AppName        string `gorm:"column:app_name;size:50" json:"app_name"`
	AppTemplateID  int    `gorm:"column:app_template_id;not null" json:"app_template_id"`
	Os             int    `gorm:"default:1" json:"os"`
	SdkVer         string `gorm:"column:sdk_ver;size:50" json:"sdk_ver"`
	AppVer         string `gorm:"column:app_ver;size:50" json:"app_ver"`
	AppKey         string `gorm:"column:app_key;size:50;not null" json:"app_key"`
	AppSecret      string `gorm:"column:app_secret;size:50;not null" json:"app_secret"`
	CallbackURL    string `gorm:"column:callback_url;size:50" json:"callback_url"`
	ApiDomain      string `gorm:"column:api_domain;size:50" json:"api_domain"`
	PayDomain      string `gorm:"column:pay_domain;size:50" json:"pay_domain"`
	Status         int8   `gorm:"default:1" json:"status"`
	CsParams       string `gorm:"column:cs_params;type:mediumtext" json:"cs_params"`
	PayParams      string `gorm:"column:pay_params;type:mediumtext" json:"pay_params"`
	H5Params       string `gorm:"column:h5_params;type:mediumtext" json:"h5_params"`
	CreatedAt      int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameApp) TableName() string {
	return "game_app"
}
