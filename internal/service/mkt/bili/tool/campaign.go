package tool

import (
	"stack-bm/internal/service/mkt/bili/api"
	"stack-bm/internal/service/mkt/oauth"
)

// CampaignTool 第一层级（广告组）批量操作
type CampaignTool struct {
	ctx *toolCtx
}

func NewCampaignTool(auth *oauth.ManagerAuth) *CampaignTool {
	return &CampaignTool{ctx: &toolCtx{auth: auth, adDataRepo: newAdDataRepo()}}
}

func (t *CampaignTool) Open(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpOpen, []int{id})
}

func (t *CampaignTool) Pause(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpStop, []int{id})
}

func (t *CampaignTool) Delete(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpDelete, []int{id})
}

func (t *CampaignTool) SetBudget(id int, budget float64) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignBudget(token, accID, id, budget)
}

func (t *CampaignTool) SetBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignBid(token, accID, id, bid, false)
}

func (t *CampaignTool) SetDeepBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignBid(token, accID, id, bid, true)
}

func (t *CampaignTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	op := api.BiliOpOpen
	if status != 1 {
		op = api.BiliOpStop
	}
	// 取首个ID解析账户/token（同一批应属同一账户）
	accID, token, err := t.ctx.resolve("campaign", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, op, ids)
}

func (t *CampaignTool) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	accID, token, err := t.ctx.resolve("campaign", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpDelete, ids)
}