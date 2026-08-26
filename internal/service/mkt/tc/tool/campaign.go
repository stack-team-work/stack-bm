package tool

import (
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tc/api"
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
	return api.UpdateAdgroupStatus(token, accID, id, api.TcAdStatusNormal)
}

func (t *CampaignTool) Pause(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupStatus(token, accID, id, api.TcAdStatusPause)
}

func (t *CampaignTool) Delete(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.DeleteAdgroup(token, accID, id)
}

func (t *CampaignTool) SetBudget(id int, budget float64) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupBudget(token, accID, id, budget)
}

func (t *CampaignTool) SetBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupBid(token, accID, id, bid)
}

func (t *CampaignTool) SetDeepBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupBid(token, accID, id, bid)
}

func (t *CampaignTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	cs := api.TcAdStatusNormal
	if status != 1 {
		cs = api.TcAdStatusPause
	}
	for _, id := range ids {
		accID, token, err := t.ctx.resolve("campaign", id)
		if err != nil {
			return false, err
		}
		if _, err := api.UpdateAdgroupStatus(token, accID, id, cs); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (t *CampaignTool) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	for _, id := range ids {
		accID, token, err := t.ctx.resolve("campaign", id)
		if err != nil {
			return false, err
		}
		if _, err := api.DeleteAdgroup(token, accID, id); err != nil {
			return false, err
		}
	}
	return true, nil
}