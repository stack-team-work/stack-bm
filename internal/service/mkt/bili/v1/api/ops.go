package api

import (
	"errors"
	"fmt"
	"strconv"

	"stack-bm/pkg/utils"
)

// B站开放API 协议常量与公共封装（对齐源 BilibiliConstants）

const (
	updateUnitURI       = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/batch/update_unit"
	updateCampaignURI   = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/batch/update_campaign"
	updateCreativeURI   = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/batch/update_creative"
	accelerateURI       = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/unit/accelerate"
	finishAccelerateURI = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/unit/finish_accelerate"
)

// 操作类型
const (
	BiliOpOpen     = 1
	BiliOpStop     = 2
	BiliOpDelete   = 3
	BiliOpBudget   = 4
	BiliOpBid      = 6
	BiliOpDate     = 7
	BiliOpDateTime = 8
)

const (
	BiliBidTypeOcpm     = 2
	BiliBidTypeDeepOcpm = 3
	BiliBudgetSpecific  = 1
	BiliBudgetUnlimited = 2
)

func bearerHeaders(accessToken string) map[string]string {
	return map[string]string{"Authorization": "Bearer " + accessToken}
}

// checkStatus B站响应 status != "success" 视为失败
func checkStatus(status, action string) error {
	if status != "success" {
		return fmt.Errorf("B站%s失败: %s", action, status)
	}
	return nil
}

// listAll 分页拉取层级列表（B站返回 data 为裸数组），补上原实现缺失的 status 校验
func listAll(accessToken string, uri string, body map[string]interface{}) ([]map[string]interface{}, error) {
	data, err := utils.HTTPPostJSON(uri, body, bearerHeaders(accessToken))
	if err != nil {
		return nil, err
	}
	var res struct {
		Status string                   `json:"status"`
		Data   []map[string]interface{} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Status != "success" {
		return nil, errors.New("B站列表获取失败: " + res.Status)
	}
	return res.Data, nil
}

// updateGeneric 通用批量更新（campaign/unit/creative 共用端点封装，uri 由调用方指定）
func updateGeneric(accessToken string, accountID int, uri string, body map[string]interface{}) (bool, error) {
	body["account_id"] = accountID
	data, err := utils.HTTPPostJSON(uri, body, bearerHeaders(accessToken))
	if err != nil {
		return false, err
	}
	var res struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		Message string `json:"message"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return false, err
	}
	if res.Status != "success" {
		return false, fmt.Errorf("B站操作失败: code=%d %s%s", res.Code, res.Msg, res.Message)
	}
	return true, nil
}

// budgetOp 组装预算操作体（budget<=0 表示不限预算）
func budgetOp(budget float64) map[string]interface{} {
	op := map[string]interface{}{"budget_limit_type": BiliBudgetUnlimited, "is_repeat": 0}
	if budget > 0 {
		op["budget_limit_type"] = BiliBudgetSpecific
		op["budget"] = strconv.FormatFloat(budget, 'f', -1, 64)
	}
	return op
}

// bidOp 组装出价操作体（deepBid 为深度出价类型）
func bidOp(bid float64, deepBid bool) map[string]interface{} {
	bidType := BiliBidTypeOcpm
	if deepBid {
		bidType = BiliBidTypeDeepOcpm
	}
	return map[string]interface{}{"bid": bid, "bid_type": bidType}
}
