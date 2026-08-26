package tool

import (
	"errors"

	ttRepo "stack-bm/internal/repository/mkt/tt"
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

func newAdDataRepo() *ttRepo.AdDataRepository {
	return ttRepo.NewAdDataRepository()
}

// toolCtx 共享依赖：渠道账户/管家/token 解析
type toolCtx struct {
	auth       *oauth.ManagerAuth
	adDataRepo *ttRepo.AdDataRepository
}

// resolve 操作ID(cpid/aid/cid) -> 渠道账户uid(string) + access_token
func (c *toolCtx) resolve(level string, id int) (advertiserID string, token string, err error) {
	if id <= 0 {
		return "", "", errors.New("id不能为空")
	}
	accountID, err := c.adDataRepo.FindAccountByLevelID(level, id)
	if err != nil {
		return "", "", err
	}
	advertiserID, token, err = c.auth.GetAccountContext(uint(accountID))
	if err != nil {
		return "", "", err
	}
	return advertiserID, token, nil
}