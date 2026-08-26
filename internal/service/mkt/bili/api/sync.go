package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	biliCashURI      = "https://cm.bilibili.com/takumi/api/open_api/report/v2/cash"
	biliCampaignList = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/campaign/list_campaigns"
	biliUnitList     = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/unit/list_units"
	biliCreativeList = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/creative/list_creatives"
)

// GetAccountCash 获取B站账户余额
func GetAccountCash(accessToken string, accountID int) (float64, error) {
	data, err := utils.HTTPGet(biliCashURI, nil, map[string]string{"Authorization": "Bearer " + accessToken})
	if err != nil {
		return 0, err
	}
	var res struct {
		Status string `json:"status"`
		Data   struct {
			Cash float64 `json:"cash"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return 0, err
	}
	if res.Status != "success" {
		return 0, fmt.Errorf("B站获取余额失败: %s", res.Status)
	}
	return res.Data.Cash, nil
}

// ListCampaigns 拉取B站广告组列表
func ListCampaigns(accessToken string, accountID int, page int, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{"account_id": accountID, "page": page, "page_size": pageSize, "status": "ALL"}
	data, err := utils.HTTPPostJSON(biliCampaignList, body, map[string]string{"Authorization": "Bearer " + accessToken})
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
	return res.Data, nil
}

// ListUnits 拉取B站广告列表
func ListUnits(accessToken string, accountID int, page int, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{"account_id": accountID, "page": page, "page_size": pageSize, "status": "ALL"}
	data, err := utils.HTTPPostJSON(biliUnitList, body, map[string]string{"Authorization": "Bearer " + accessToken})
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
	return res.Data, nil
}

// ListCreatives 拉取B站创意列表
func ListCreatives(accessToken string, accountID int, page int, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{"account_id": accountID, "page": page, "page_size": pageSize, "status": "ALL"}
	data, err := utils.HTTPPostJSON(biliCreativeList, body, map[string]string{"Authorization": "Bearer " + accessToken})
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
	return res.Data, nil
}