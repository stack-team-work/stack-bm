package ad

// ListRequest 广告数据列表通用入参（columns 为前端勾选的指标列）
type ListRequest struct {
	Page      int                    `json:"page"`
	Size      int                    `json:"size"`
	Columns   []string               `json:"columns"`
	Keyword   string                 `json:"keyword"`
	Status    int                    `json:"status"`
	AccountID string                 `json:"account_id"`
	StartDate string                 `json:"start_date"`
	EndDate   string                 `json:"end_date"`
	Extra     map[string]interface{} `json:"extra"`
}

// Normalize 兜底分页参数
func (r *ListRequest) Normalize() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.Size < 1 {
		r.Size = 10
	}
}

// Filters 组装查询筛选条件
func (r *ListRequest) Filters() map[string]interface{} {
	filters := map[string]interface{}{}
	if r.Keyword != "" {
		filters["keyword"] = r.Keyword
	}
	if r.Status != 0 {
		filters["status"] = r.Status
	}
	if r.AccountID != "" {
		filters["account_id"] = r.AccountID
	}
	if r.StartDate != "" {
		filters["start_date"] = r.StartDate
	}
	if r.EndDate != "" {
		filters["end_date"] = r.EndDate
	}
	for k, v := range r.Extra {
		filters[k] = v
	}
	return filters
}

// ActionRequest 层级操作通用入参（单条传 id，批量传 ids）
type ActionRequest struct {
	ID        int     `json:"id"`
	IDs       []int   `json:"ids"`
	Status    int     `json:"status"`
	Budget    float64 `json:"budget"`
	Bid       float64 `json:"bid"`
	DeepBid   float64 `json:"deep_bid"`
	BeginDate string  `json:"begin_date"`
}
