package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	projectListURI  = "https://ad.oceanengine.com/open_api/v3.0/project/get/"
	updateStatusURI = "https://ad.oceanengine.com/open_api/v3.0/project/status/update/"
	updateBudgetURI = "https://ad.oceanengine.com/open_api/v3.0/project/budget/update/"
	updateURI       = "https://ad.oceanengine.com/open_api/v3.0/project/update/"
	deleteURI       = "https://ad.oceanengine.com/open_api/v3.0/project/delete/"
)

// 头条V3 项目操作状态
const (
	TtOptStatusEnable  = "ENABLE"
	TtOptStatusDisable = "DISABLE"
	// 预算模式
	TtBudgetModeDay      = "BUDGET_MODE_DAY"
	TtBudgetModeInfinite = "BUDGET_MODE_INFINITE"
)

func headers(accessToken string) map[string]string {
	return map[string]string{"Access-Token": accessToken}
}

// post 头条 POST：注入 advertiser_id（字符串），成功 code==0
func post(accessToken string, advertiserID string, uri string, body map[string]interface{}) (bool, error) {
	body["advertiser_id"] = advertiserID
	data, err := utils.HTTPPostJSON(uri, body, headers(accessToken))
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

// ListProjects 分页拉取头条V3项目列表
func ListProjects(accessToken string, advertiserID string, page, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
		"filtering":     map[string]interface{}{"status": "ALL"},
	}
	data, err := utils.HTTPPostJSON(projectListURI, body, headers(accessToken))
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("头条项目列表获取失败: code=%d %s", res.Code, res.Message)
	}
	return res.Data.List, nil
}

// UpdateProjectStatus 项目启停（批量）
func UpdateProjectStatus(accessToken string, advertiserID string, optStatus string, projectIDs []int) (bool, error) {
	data := make([]map[string]interface{}, 0, len(projectIDs))
	for _, pid := range projectIDs {
		data = append(data, map[string]interface{}{"project_id": pid, "opt_status": optStatus})
	}
	return post(accessToken, advertiserID, updateStatusURI, map[string]interface{}{"data": data})
}

// DeleteProject 删除项目（批量）
func DeleteProject(accessToken string, advertiserID string, projectIDs []int) (bool, error) {
	return post(accessToken, advertiserID, deleteURI, map[string]interface{}{"project_ids": projectIDs})
}

// UpdateProjectBudget 项目预算
func UpdateProjectBudget(accessToken string, advertiserID string, projectID int, budget float64) (bool, error) {
	budgetMode := TtBudgetModeInfinite
	if budget > 0 {
		budgetMode = TtBudgetModeDay
	}
	return post(accessToken, advertiserID, updateBudgetURI, map[string]interface{}{
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
	return post(accessToken, advertiserID, updateURI, body)
}
