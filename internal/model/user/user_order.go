package user

type UserOrder struct {
	ID             uint    `gorm:"primarykey" json:"id"`
	OrderNum       string  `gorm:"column:order_num;size:64" json:"order_num"`
	ThirdOrderNum  string  `gorm:"column:third_order_num;size:64" json:"third_order_num"`
	CpOrderNum     string  `gorm:"column:cp_order_num;size:64" json:"cp_order_num"`
	UserID         string  `gorm:"column:user_id;size:50" json:"user_id"`
	Pid            int     `gorm:"column:pid" json:"pid"`
	AppID          int     `gorm:"column:app_id" json:"app_id"`
	ServerID       string  `gorm:"column:server_id;size:50" json:"server_id"`
	ServerName     string  `gorm:"column:server_name;size:50" json:"server_name"`
	RoleID         string  `gorm:"column:role_id;size:50" json:"role_id"`
	RoleName       string  `gorm:"column:role_name;size:50" json:"role_name"`
	Currency       string  `gorm:"column:currency;size:16" json:"currency"`
	Discount       float64 `gorm:"column:discount" json:"discount"`
	OriginTotalFee float64 `gorm:"column:origin_total_fee" json:"origin_total_fee"`
	TotalFee       float64 `gorm:"column:total_fee" json:"total_fee"`
	Product        string  `gorm:"column:product;size:32" json:"product"`
	ProductID      string  `gorm:"column:product_id;size:32" json:"product_id"`
	IsFirst        int     `gorm:"column:is_first" json:"is_first"`
	PayStatus      int8    `gorm:"column:pay_status" json:"pay_status"`
	Status         int8    `gorm:"column:status" json:"status"`
	PayWay         int8    `gorm:"column:pay_way" json:"pay_way"`
	PayType        int     `gorm:"column:pay_type" json:"pay_type"`
	PayAt          int64   `gorm:"column:pay_at" json:"pay_at"`
	CallbackAt     int64   `gorm:"column:callback_at" json:"callback_at"`
	RegAt          int64   `gorm:"column:reg_at" json:"reg_at"`
	MediaID        int     `gorm:"column:media_id" json:"media_id"`
	MediaSubID     int     `gorm:"column:media_sub_id" json:"media_sub_id"`
	IP             string  `gorm:"column:ip;size:64" json:"ip"`
	IsTest         int8    `gorm:"column:is_test" json:"is_test"`
	CreatedAt      int64   `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      int64   `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (UserOrder) TableName() string { return "user_orders" }