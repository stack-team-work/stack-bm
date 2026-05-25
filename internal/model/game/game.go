package game

type Game struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:64;not null" json:"name"`
	WebName   string `gorm:"column:web_name;size:64;not null" json:"web_name"`
	Icon      string `gorm:"size:64" json:"icon"`
	IsWebShow int8   `gorm:"column:is_web_show;default:0" json:"is_web_show"`
	TypeID    int    `gorm:"column:type_id" json:"type_id"`
	StyleID   int    `gorm:"column:style_id" json:"style_id"`
	CpID      int    `gorm:"column:cp_id;not null" json:"cp_id"`
	ServerURL string `gorm:"column:server_url;size:64" json:"server_url"`
	RoleURL   string `gorm:"column:role_url;size:64" json:"role_url"`
	AuthName  string `gorm:"column:auth_name;size:64" json:"auth_name"`
	Author    string `gorm:"type:mediumtext" json:"author"`
	Mark      string `gorm:"size:64;uniqueIndex" json:"mark"`
	Status    int    `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Game) TableName() string {
	return "game"
}

type GameApp struct {
	ID             uint   `gorm:"primarykey" json:"id"`
	Pid            int    `gorm:"not null" json:"pid"`
	Name           string `gorm:"size:50;not null" json:"name"`
	PackageName    string `gorm:"column:package_name;size:50" json:"package_name"`
	AppName        string `gorm:"column:app_name;size:50" json:"app_name"`
	Os             int    `gorm:"default:1" json:"os"`
	IsVerify       int8   `gorm:"column:is_verify;default:0" json:"is_verify"`
	Age            int    `gorm:"default:18" json:"age"`
	IsOpenCharge   int8   `gorm:"column:is_open_charge;default:1" json:"is_open_charge"`
	IsOpenRegister int8   `gorm:"column:is_open_register;default:1" json:"is_open_register"`
	IsAlertEmail   int8   `gorm:"column:is_alert_email;default:1" json:"is_alert_email"`
	IsAlertPhone   int8   `gorm:"column:is_alert_phone;default:1" json:"is_alert_phone"`
	IsAlertAuth    int8   `gorm:"column:is_alert_auth;default:0" json:"is_alert_auth"`
	IsOpenFloat    int8   `gorm:"column:is_open_float;default:1" json:"is_open_float"`
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
	IsDeleted      int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt      int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameApp) TableName() string {
	return "game_app"
}

type GameCp struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:64" json:"name"`
	Mark      string `gorm:"size:64" json:"mark"`
	Phone     string `gorm:"size:16" json:"phone"`
	Addr      string `gorm:"size:255" json:"addr"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	Status    int8   `gorm:"default:1" json:"status"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameCp) TableName() string {
	return "game_cp"
}

type GameTag struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Type      int8   `gorm:"not null;default:1" json:"type"`
	Mark      string `gorm:"size:255;not null" json:"mark"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Status    int    `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (GameTag) TableName() string {
	return "game_tags"
}

type GameVariable struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Key       string `gorm:"size:32;uniqueIndex;not null" json:"key"`
	Value     string `gorm:"type:mediumtext;not null" json:"value"`
	Mark      string `gorm:"size:255" json:"mark"`
	Status    int8   `gorm:"default:1" json:"status"`
	IsDeleted int8   `gorm:"column:is_deleted;default:0" json:"is_deleted"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (GameVariable) TableName() string {
	return "game_variable"
}
