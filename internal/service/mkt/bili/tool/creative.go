package tool

import (
	"stack-bm/internal/service/mkt/bili/api"
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
	accID, token, err := t.ctx.resolve("creative", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, accID, api.BiliOpOpen, []int{id})
}

func (t *CreativeTool) Pause(id int) (bool, error) {
	accID, token, err := t.ctx.resolve("creative", id)
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, accID, api.BiliOpStop, []int{id})
}

func (t *CreativeTool) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	op := api.BiliOpOpen
	if status != 1 {
		op = api.BiliOpStop
	}
	accID, token, err := t.ctx.resolve("creative", ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, accID, op, ids)
}

func (t *CreativeTool) Preview(id int) (string, error) {
	// 创意预览URL待媒体接口确认后实现
	return "", nil
}