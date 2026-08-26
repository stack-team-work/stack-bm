package tool

import (
	"stack-bm/internal/service/mkt/ks/api"
	"stack-bm/internal/service/mkt/oauth"
)

// UnitTool 第二层级（广告）批量操作
type UnitTool struct {
	ctx *toolCtx
}

func NewUnitTool(auth *oauth.ManagerAuth) *UnitTool {
	return &UnitTool{ctx: &toolCtx{auth: auth, adDataRepo: newAdDataRepo()}}
}

func (t *UnitTool) Open(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, advID, api.KsStatusOpen, []int{id})
}

func (t *UnitTool) Pause(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, advID, api.KsStatusPause, []int{id})
}

func (t *UnitTool) Delete(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, advID, api.KsStatusDelete, []int{id})
}

func (t *UnitTool) SetBudget(id int, budget float64) (bool, error) {
	advID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBudget(token, advID, id, budget)
}

func (t *UnitTool) SetBid(id int, bid float64) (bool, error) {
	advID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, advID, id, bid, false)
}

func (t *UnitTool) SetDeepBid(id int, bid float64) (bool, error) {
	advID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, advID, id, bid, true)
}

func (t *UnitTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	op := api.KsStatusOpen
	if status != 1 {
		op = api.KsStatusPause
	}
	advID, token, err := t.ctx.resolve("unit", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, advID, op, ids)
}

func (t *UnitTool) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	advID, token, err := t.ctx.resolve("unit", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, advID, api.KsStatusDelete, ids)
}

func (t *UnitTool) SetBeginDate(id int, beginDate string) (bool, error) {
	// 快手广告投放日期更新接口暂未实现
	return true, nil
}

func (t *UnitTool) SetRaise(id int) (bool, error) { return true, nil }
func (t *UnitTool) StopRaise(id int) (bool, error) { return true, nil }
func (t *UnitTool) BatchSetRaise(ids []int) (bool, error) { return true, nil }
func (t *UnitTool) BatchStopRaise(ids []int) (bool, error) { return true, nil }
func (t *UnitTool) Collect(id int) (bool, error) { return true, nil }
func (t *UnitTool) CancelCollect(id int) (bool, error) { return true, nil }
func (t *UnitTool) RaiseInfo(id int) (map[string]interface{}, error) {
	return map[string]interface{}{"status": "无起量信息"}, nil
}