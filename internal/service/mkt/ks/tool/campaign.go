package tool

import (
	"stack-bm/internal/service/mkt/ks/api"
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
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, advID, api.KsStatusOpen, []int{id})
}

func (t *CampaignTool) Pause(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, advID, api.KsStatusPause, []int{id})
}

func (t *CampaignTool) Delete(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, advID, api.KsStatusDelete, []int{id})
}

func (t *CampaignTool) SetBudget(id int, budget float64) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBudget(token, advID, id, budget)
}

func (t *CampaignTool) SetBid(id int, bid float64) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, advID, id, bid, false)
}

func (t *CampaignTool) SetDeepBid(id int, bid float64) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, advID, id, bid, true)
}

func (t *CampaignTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	advID, token, err := t.ctx.resolve("campaign", ids[0])
	if err != nil {
		return false, err
	}
	op := api.KsStatusOpen
	if status != 1 {
		op = api.KsStatusPause
	}
	return api.UpdateCampaignStatus(token, advID, op, ids)
}

func (t *CampaignTool) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	advID, token, err := t.ctx.resolve("campaign", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, advID, api.KsStatusDelete, ids)
}