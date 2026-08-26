package api

import (
	"fmt"
	"strconv"

	"stack-bm/pkg/utils"
)

const (
	biliUpdateUnitURI        = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/batch/update_unit"
	biliUpdateCampaignURI    = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/batch/update_campaign"
	biliUpdateCreativeURI    = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/batch/update_creative"
	biliAccelerateURI        = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/unit/accelerate"
	biliFinishAccelerateURI  = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/unit/finish_accelerate"
)

// B站操作类型（对齐源 BilibiliConstants）
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

// biliUpdateGeneric 通用批量更新（campaign/unit/creative），uri 由调用方指定
func biliUpdateGeneric(accessToken string, accountID int, uri string, body map[string]interface{}) (bool, error) {
	body["account_id"] = accountID
	data, err := utils.HTTPPostJSON(uri, body, map[string]string{"Authorization": "Bearer " + accessToken})
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

// UpdateUnitStatus 单元启停/删除
func UpdateUnitStatus(accessToken string, accountID int, operationType int, unitIDs []int) (bool, error) {
	body := map[string]interface{}{
		"operation_type": operationType,
		"operations":     []map[string]interface{}{{"unit_ids": unitIDs}},
	}
	return biliUpdateGeneric(accessToken, accountID, biliUpdateUnitURI, body)
}

// UpdateCampaignStatus 广告组启停/删除
func UpdateCampaignStatus(accessToken string, accountID int, operationType int, campaignIDs []int) (bool, error) {
	body := map[string]interface{}{
		"operation_type": operationType,
		"operations":     []map[string]interface{}{{"campaign_ids": campaignIDs}},
	}
	return biliUpdateGeneric(accessToken, accountID, biliUpdateCampaignURI, body)
}

// UpdateCreativeStatus 创意启停/删除
func UpdateCreativeStatus(accessToken string, accountID int, operationType int, creativeIDs []int) (bool, error) {
	body := map[string]interface{}{
		"operation_type": operationType,
		"operations":     []map[string]interface{}{{"creative_ids": creativeIDs}},
	}
	return biliUpdateGeneric(accessToken, accountID, biliUpdateCreativeURI, body)
}

// UpdateUnitBudget 修改单元预算
func UpdateUnitBudget(accessToken string, accountID, unitID int, budget float64) (bool, error) {
	budgetOp := map[string]interface{}{"budget_limit_type": BiliBudgetUnlimited, "is_repeat": 0}
	if budget > 0 {
		budgetOp["budget_limit_type"] = BiliBudgetSpecific
		budgetOp["budget"] = strconv.FormatFloat(budget, 'f', -1, 64)
	}
	body := map[string]interface{}{
		"operation_type": BiliOpBudget,
		"operations":     []map[string]interface{}{{"unit_ids": []int{unitID}, "budget": budgetOp}},
	}
	return biliUpdateGeneric(accessToken, accountID, biliUpdateUnitURI, body)
}

// UpdateUnitBid 修改单元出价
func UpdateUnitBid(accessToken string, accountID, unitID int, bid float64, deepBid bool) (bool, error) {
	bidType := BiliBidTypeOcpm
	if deepBid {
		bidType = BiliBidTypeDeepOcpm
	}
	body := map[string]interface{}{
		"operation_type": BiliOpBid,
		"operations": []map[string]interface{}{{
			"unit_ids": []int{unitID},
			"bid":      map[string]interface{}{"bid": bid, "bid_type": bidType},
		}},
	}
	return biliUpdateGeneric(accessToken, accountID, biliUpdateUnitURI, body)
}

// UpdateCampaignBudget 修改广告组预算
func UpdateCampaignBudget(accessToken string, accountID, campaignID int, budget float64) (bool, error) {
	budgetOp := map[string]interface{}{"budget_limit_type": BiliBudgetUnlimited, "is_repeat": 0}
	if budget > 0 {
		budgetOp["budget_limit_type"] = BiliBudgetSpecific
		budgetOp["budget"] = strconv.FormatFloat(budget, 'f', -1, 64)
	}
	body := map[string]interface{}{
		"operation_type": BiliOpBudget,
		"operations":     []map[string]interface{}{{"campaign_ids": []int{campaignID}, "budget": budgetOp}},
	}
	return biliUpdateGeneric(accessToken, accountID, biliUpdateCampaignURI, body)
}

// UpdateCampaignBid 修改广告组出价
func UpdateCampaignBid(accessToken string, accountID, campaignID int, bid float64, deepBid bool) (bool, error) {
	bidType := BiliBidTypeOcpm
	if deepBid {
		bidType = BiliBidTypeDeepOcpm
	}
	body := map[string]interface{}{
		"operation_type": BiliOpBid,
		"operations": []map[string]interface{}{{
			"campaign_ids": []int{campaignID},
			"bid":          map[string]interface{}{"bid": bid, "bid_type": bidType},
		}},
	}
	return biliUpdateGeneric(accessToken, accountID, biliUpdateCampaignURI, body)
}

// OpenAccelerate 开启一键起量
func OpenAccelerate(accessToken string, accountID, unitID int, accelerateBudget float64, accelerateType int, accelerateDuration int) (bool, error) {
	body := map[string]interface{}{
		"account_id": accountID,
		"unit_accelerate_list": []map[string]interface{}{{
			"unit_id":             unitID,
			"accelerate_budget":   accelerateBudget,
			"accelerate_type":     accelerateType,
			"accelerate_duration": accelerateDuration,
		}},
	}
	data, err := utils.HTTPPostJSON(biliAccelerateURI, body, map[string]string{"Authorization": "Bearer " + accessToken})
	if err != nil {
		return false, err
	}
	var res struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return false, err
	}
	if res.Status != "success" {
		return false, fmt.Errorf("B站开启一键起量失败: %s", res.Msg)
	}
	return true, nil
}

// CloseAccelerate 关闭一键起量
func CloseAccelerate(accessToken string, accountID, unitID int) (bool, error) {
	body := map[string]interface{}{
		"account_id": accountID,
		"unit_ids":   []int{unitID},
	}
	data, err := utils.HTTPPostJSON(biliFinishAccelerateURI, body, map[string]string{"Authorization": "Bearer " + accessToken})
	if err != nil {
		return false, err
	}
	var res struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return false, err
	}
	if res.Status != "success" {
		return false, fmt.Errorf("B站关闭一键起量失败: %s", res.Msg)
	}
	return true, nil
}