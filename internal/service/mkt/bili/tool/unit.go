package tool

import (
	"stack-bm/internal/service/mkt/bili/api"
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
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpOpen, []int{id})
}

func (t *UnitTool) Pause(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpStop, []int{id})
}

func (t *UnitTool) Delete(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpDelete, []int{id})
}

func (t *UnitTool) SetBudget(id int, budget float64) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBudget(token, accID, id, budget)
}

func (t *UnitTool) SetBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, accID, id, bid, false)
}

func (t *UnitTool) SetDeepBid(id int, bid float64) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, accID, id, bid, true)
}

func (t *UnitTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	op := api.BiliOpOpen
	if status != 1 {
		op = api.BiliOpStop
	}
	accID, token, err := t.ctx.resolve("unit", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, op, ids)
}

func (t *UnitTool) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	accID, token, err := t.ctx.resolve("unit", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpDelete, ids)
}

func (t *UnitTool) SetRaise(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.OpenAccelerate(token, accID, id, 0, 1, 6)
}

func (t *UnitTool) StopRaise(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.CloseAccelerate(token, accID, id)
}

func (t *UnitTool) SetBeginDate(id int, beginDate string) (bool, error) {
	accID, token, err := t.ctx.resolve("unit", id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpDate, []int{id})
}

func (t *UnitTool) BatchSetRaise(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	for _, id := range ids {
		if _, err := t.SetRaise(id); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (t *UnitTool) BatchStopRaise(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	for _, id := range ids {
		if _, err := t.StopRaise(id); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (t *UnitTool) Collect(id int) (bool, error) {
	// 收藏为本地功能，待收藏表就绪后实现
	return true, nil
}

func (t *UnitTool) CancelCollect(id int) (bool, error) {
	// 收藏为本地功能，待收藏表就绪后实现
	return true, nil
}

func (t *UnitTool) RaiseInfo(id int) (map[string]interface{}, error) {
	return map[string]interface{}{"status": "无起量信息"}, nil
}