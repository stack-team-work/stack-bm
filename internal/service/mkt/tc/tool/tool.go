package tool

import (
	"errors"

	tcRepo "stack-bm/internal/repository/mkt/tc"
	"stack-bm/internal/service/mkt/oauth"
)

var errEmptyIDs = errors.New("ids不能为空")

// ToolRequest 批量操作通用入参
type ToolRequest struct {
	ID        int     `json:"id"`
	IDs       []int   `json:"ids"`
	Status    int     `json:"status"`
	Budget    float64 `json:"budget"`
	Bid       float64 `json:"bid"`
	DeepBid   float64 `json:"deep_bid"`
	BeginDate string  `json:"begin_date"`
}

func newAdDataRepo() *tcRepo.AdDataRepository {
	return tcRepo.NewAdDataRepository()
}

// toolCtx 共享依赖：渠道账户/管家/token 解析
type toolCtx struct {
	auth       *oauth.ManagerAuth
	adDataRepo *tcRepo.AdDataRepository
}

// resolve 操作ID(cpid/aid/cid) -> 渠道账户uid(string) + access_token
func (c *toolCtx) resolve(level string, id int) (accountID string, token string, err error) {
	if id <= 0 {
		return "", "", errors.New("id不能为空")
	}
	accID, err := c.adDataRepo.FindAccountByLevelID(level, id)
	if err != nil {
		return "", "", err
	}
	accountID, token, err = c.auth.GetAccountContext(uint(accID))
	if err != nil {
		return "", "", err
	}
	return accountID, token, nil
}