package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const unitListURI = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/unit/list_units"

// ListUnits 分页拉取B站单元列表
func ListUnits(accessToken string, accountID, page, pageSize int) ([]map[string]interface{}, error) {
	return listAll(accessToken, unitListURI, map[string]interface{}{
		"account_id": accountID,
		"page":       page,
		"page_size":  pageSize,
		"status":     "ALL",
	})
}

// UpdateUnitStatus 单元启停/删除/改投放时间操作类型（批量）
func UpdateUnitStatus(accessToken string, accountID int, operationType int, unitIDs []int) (bool, error) {
	body := map[string]interface{}{
		"operation_type": operationType,
		"operations":     []map[string]interface{}{{"unit_ids": unitIDs}},
	}
	return updateGeneric(accessToken, accountID, updateUnitURI, body)
}

// UpdateUnitBudget 修改单元预算
func UpdateUnitBudget(accessToken string, accountID, unitID int, budget float64) (bool, error) {
	body := map[string]interface{}{
		"operation_type": BiliOpBudget,
		"operations":     []map[string]interface{}{{"unit_ids": []int{unitID}, "budget": budgetOp(budget)}},
	}
	return updateGeneric(accessToken, accountID, updateUnitURI, body)
}

// UpdateUnitBid 修改单元出价
func UpdateUnitBid(accessToken string, accountID, unitID int, bid float64, deepBid bool) (bool, error) {
	body := map[string]interface{}{
		"operation_type": BiliOpBid,
		"operations":     []map[string]interface{}{{"unit_ids": []int{unitID}, "bid": bidOp(bid, deepBid)}},
	}
	return updateGeneric(accessToken, accountID, updateUnitURI, body)
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
	data, err := utils.HTTPPostJSON(accelerateURI, body, bearerHeaders(accessToken))
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
	data, err := utils.HTTPPostJSON(finishAccelerateURI, body, bearerHeaders(accessToken))
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
