package api

import (
	"fmt"
	"net/url"

	"stack-bm/pkg/utils"
)

const fundURI = "https://ad.e.kuaishou.com/rest/openapi/v1/advertiser/fund/get"

// GetAccountFund 获取快手账户余额（单位：厘 -> 元）
func GetAccountFund(accessToken string, advertiserID int) (float64, error) {
	query := url.Values{}
	query.Set("advertiser_id", fmt.Sprint(advertiserID))
	data, err := utils.HTTPGet(fundURI+"?"+query.Encode(), nil, headers(accessToken))
	if err != nil {
		return 0, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Balance int64 `json:"balance"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return 0, err
	}
	if res.Code != 0 {
		return 0, fmt.Errorf("快手获取余额失败: code=%d %s", res.Code, res.Message)
	}
	return float64(res.Data.Balance) / 1000, nil
}
