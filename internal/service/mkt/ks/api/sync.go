package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	ksFundURI       = "https://ad.e.kuaishou.com/rest/openapi/v1/advertiser/fund/get"
	ksCampaignList  = "https://ad.e.kuaishou.com/rest/openapi/gw/dsp/campaign/list"
	ksUnitList      = "https://ad.e.kuaishou.com/rest/openapi/gw/dsp/unit/list"
	ksCreativeList  = "https://ad.e.kuaishou.com/rest/openapi/gw/dsp/creative/list"
)

// GetAccountFund 获取快手账户余额（单位：厘 -> 元）
func GetAccountFund(accessToken string, advertiserID int) (float64, error) {
	data, err := utils.HTTPGet(ksFundURI+"?advertiser_id="+fmt.Sprint(advertiserID), nil, map[string]string{"Access-Token": accessToken})
	if err != nil {
		return 0, err
	}
	var res struct {
		Code int `json:"code"`
		Data struct {
			Balance int64 `json:"balance"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return 0, err
	}
	if res.Code != 0 {
		return 0, fmt.Errorf("快手获取余额失败: code=%d", res.Code)
	}
	return float64(res.Data.Balance) / 1000, nil
}

// ksList 快手列表拉取
func ksList(accessToken string, advertiserID int, uri string, page int, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{"advertiser_id": advertiserID, "page": page, "page_size": pageSize}
	data, err := utils.HTTPPostJSON(uri, body, map[string]string{"Access-Token": accessToken})
	if err != nil {
		return nil, err
	}
	var res struct {
		Code int                      `json:"code"`
		Data []map[string]interface{} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("快手列表获取失败: code=%d", res.Code)
	}
	return res.Data, nil
}

func ListCampaigns(accessToken string, advertiserID int, page, pageSize int) ([]map[string]interface{}, error) {
	return ksList(accessToken, advertiserID, ksCampaignList, page, pageSize)
}

func ListUnits(accessToken string, advertiserID int, page, pageSize int) ([]map[string]interface{}, error) {
	return ksList(accessToken, advertiserID, ksUnitList, page, pageSize)
}

func ListCreatives(accessToken string, advertiserID int, page, pageSize int) ([]map[string]interface{}, error) {
	return ksList(accessToken, advertiserID, ksCreativeList, page, pageSize)
}