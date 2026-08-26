package tool

import (
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tt/api"
)

// CampaignTool 第一层级（项目）批量操作
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
	return api.UpdateProjectStatus(token, advID, api.TtOptStatusEnable, []int{id})
}

func (t *CampaignTool) Pause(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateProjectStatus(token, advID, api.TtOptStatusDisable, []int{id})
}

func (t *CampaignTool) Delete(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.DeleteProject(token, advID, []int{id})
}

func (t *CampaignTool) SetBudget(id int, budget float64) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateProjectBudget(token, advID, id, budget)
}

func (t *CampaignTool) SetBid(id int, bid float64) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateProjectBid(token, advID, id, bid, false)
}

func (t *CampaignTool) SetDeepBid(id int, bid float64) (bool, error) {
	advID, token, err := t.ctx.resolve("campaign", id)
	if err != nil {
		return false, err
	}
	return api.UpdateProjectBid(token, advID, id, bid, true)
}

func (t *CampaignTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	opt := api.TtOptStatusEnable
	if status != 1 {
		opt = api.TtOptStatusDisable
	}
	advID, token, err := t.ctx.resolve("campaign", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateProjectStatus(token, advID, opt, ids)
}

func (t *CampaignTool) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	advID, token, err := t.ctx.resolve("campaign", ids[0])
	if err != nil {
		return false, err
	}
	return api.DeleteProject(token, advID, ids)
}