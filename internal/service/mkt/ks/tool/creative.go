package tool

import (
	"stack-bm/internal/service/mkt/ks/api"
	"stack-bm/internal/service/mkt/oauth"
)

// CreativeTool 第三层级（创意）批量操作
type CreativeTool struct {
	ctx *toolCtx
}

func NewCreativeTool(auth *oauth.ManagerAuth) *CreativeTool {
	return &CreativeTool{ctx: &toolCtx{auth: auth, adDataRepo: newAdDataRepo()}}
}

func (t *CreativeTool) Open(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("creative", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, advID, api.KsStatusOpen, []int{id})
}

func (t *CreativeTool) Pause(id int) (bool, error) {
	advID, token, err := t.ctx.resolve("creative", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, advID, api.KsStatusPause, []int{id})
}

func (t *CreativeTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	op := api.KsStatusOpen
	if status != 1 {
		op = api.KsStatusPause
	}
	advID, token, err := t.ctx.resolve("creative", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, advID, op, ids)
}

func (t *CreativeTool) Preview(id int) (string, error) {
	return "", nil
}