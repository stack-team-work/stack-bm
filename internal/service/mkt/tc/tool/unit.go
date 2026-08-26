package tool

import (
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tc/api"
)

// UnitTool 第二层级（计划）批量操作
type UnitTool struct {
	ctx *toolCtx
}

func NewUnitTool(auth *oauth.ManagerAuth) *UnitTool {
	return &UnitTool{ctx: &toolCtx{auth: auth, adDataRepo: newAdDataRepo()}}
}

func (t *UnitTool) Open(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupStatus(token, accID, id, api.TcAdStatusNormal)
}

func (t *UnitTool) Pause(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupStatus(token, accID, id, api.TcAdStatusPause)
}

func (t *UnitTool) Delete(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.DeleteAdgroup(token, accID, id)
}

func (t *UnitTool) SetBudget(id int, budget float64) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupBudget(token, accID, id, budget)
}

func (t *UnitTool) SetBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupBid(token, accID, id, bid)
}

func (t *UnitTool) SetDeepBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateAdgroupBid(token, accID, id, bid)
}

func (t *UnitTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	cs := api.TcAdStatusNormal
	if status != 1 {
		cs = api.TcAdStatusPause
	}
	for _, id := range ids {
		accID, token, err := t.ctx.resolve("unit", id)
		if err != nil {
			return false, err
		}
		if _, err := api.UpdateAdgroupStatus(token, accID, id, cs); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (t *UnitTool) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	for _, id := range ids {
		accID, token, err := t.ctx.resolve("unit", id)
		if err != nil {
			return false, err
		}
		if _, err := api.DeleteAdgroup(token, accID, id); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (t *UnitTool) SetBeginDate(id int, beginDate string) (bool, error) { return true, nil }
func (t *UnitTool) SetRaise(id int) (bool, error) { return true, nil }
func (t *UnitTool) StopRaise(id int) (bool, error) { return true, nil }
func (t *UnitTool) BatchSetRaise(ids []int) (bool, error) { return true, nil }
func (t *UnitTool) BatchStopRaise(ids []int) (bool, error) { return true, nil }
func (t *UnitTool) Collect(id int) (bool, error) { return true, nil }
func (t *UnitTool) CancelCollect(id int) (bool, error) { return true, nil }
func (t *UnitTool) RaiseInfo(id int) (map[string]interface{}, error) {
	return map[string]interface{}{"status": "无起量信息"}, nil
}