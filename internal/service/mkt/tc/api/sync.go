package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	tcWalletURI    = "https://api.e.qq.com/v3.0/wallet/get"
	tcAdgroupList  = "https://api.e.qq.com/v3.0/adgroups/get"
)

// GetAccountWallet 获取腾讯账户钱包余额
func GetAccountWallet(accessToken string, accountID string) (float64, error) {
	body := map[string]interface{}{"account_id": accountID}
	data, err := utils.HTTPPostJSON(tcWalletURI, body, map[string]string{"Access-Token": accessToken})
	if err != nil {
		return 0, err
	}
	var res struct {
		Code int `json:"code"`
		Data struct {
			Wallet struct {
				Balance int64 `json:"balance"`
			} `json:"wallet"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return 0, err
	}
	if res.Code != 0 {
		return 0, fmt.Errorf("腾讯获取钱包失败: code=%d", res.Code)
	}
	return float64(res.Data.Wallet.Balance) / 100, nil
}

// ListAdgroups 拉取腾讯广告组列表
func ListAdgroups(accessToken string, accountID string, page, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{"account_id": accountID, "page": page, "page_size": pageSize}
	data, err := utils.HTTPPostJSON(tcAdgroupList, body, map[string]string{"Access-Token": accessToken})
	if err != nil {
		return nil, err
	}
	var res struct {
		Code int `json:"code"`
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("腾讯广告组列表获取失败: code=%d", res.Code)
	}
	return res.Data.List, nil
}