package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	ttUpdateProjectStatusURI = "https://ad.oceanengine.com/open_api/v3.0/project/status/update/"
	ttUpdateProjectBudgetURI = "https://ad.oceanengine.com/open_api/v3.0/project/budget/update/"
	ttUpdateProjectURI       = "https://ad.oceanengine.com/open_api/v3.0/project/update/"
	ttDeleteProjectURI       = "https://ad.oceanengine.com/open_api/v3.0/project/delete/"
)

// 头条V3 项目操作状态
const (
	TtOptStatusEnable  = "ENABLE"
	TtOptStatusDisable = "DISABLE"
	// 预算模式
	TtBudgetModeDay      = "BUDGET_MODE_DAY"
	TtBudgetModeInfinite = "BUDGET_MODE_INFINITE"
)

// ttPost 头条 POST，Access-Token 头 + advertiser_id 参数，成功 code==0
func ttPost(accessToken string, advertiserID string, uri string, body map[string]interface{}) (bool, error) {
	body["advertiser_id"] = advertiserID
	data, err := utils.HTTPPostJSON(uri, body, map[string]string{"Access-Token": accessToken})
	if err != nil {
		return false, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return false, err
	}
	if res.Code != 0 {
		return false, fmt.Errorf("头条操作失败: code=%d %s", res.Code, res.Message)
	}
	return true, nil
}

// UpdateProjectStatus 项目启停
func UpdateProjectStatus(accessToken string, advertiserID string, optStatus string, projectIDs []int) (bool, error) {
	data := make([]map[string]interface{}, 0, len(projectIDs))
	for _, pid := range projectIDs {
		data = append(data, map[string]interface{}{"project_id": pid, "opt_status": optStatus})
	}
	return ttPost(accessToken, advertiserID, ttUpdateProjectStatusURI, map[string]interface{}{"data": data})
}

// DeleteProject 删除项目
func DeleteProject(accessToken string, advertiserID string, projectIDs []int) (bool, error) {
	return ttPost(accessToken, advertiserID, ttDeleteProjectURI, map[string]interface{}{"project_ids": projectIDs})
}

// UpdateProjectBudget 项目预算
func UpdateProjectBudget(accessToken string, advertiserID string, projectID int, budget float64) (bool, error) {
	budgetMode := TtBudgetModeInfinite
	if budget > 0 {
		budgetMode = TtBudgetModeDay
	}
	return ttPost(accessToken, advertiserID, ttUpdateProjectBudgetURI, map[string]interface{}{
		"project_id":  projectID,
		"budget_mode": budgetMode,
		"budget":      budget,
	})
}

// UpdateProjectBid 项目出价（通过项目更新接口）
func UpdateProjectBid(accessToken string, advertiserID string, projectID int, bid float64, deepBid bool) (bool, error) {
	body := map[string]interface{}{"project_id": projectID}
	if deepBid {
		body["deep_bid"] = bid
	} else {
		body["bid"] = bid
	}
	return ttPost(accessToken, advertiserID, ttUpdateProjectURI, body)
}