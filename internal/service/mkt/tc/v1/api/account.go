package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const walletURI = "https://api.e.qq.com/v3.0/wallet/get"

func headers(accessToken string) map[string]string {
	return map[string]string{"Access-Token": accessToken}
}

// GetAccountWallet 获取腾讯账户钱包余额（分 -> 元）
func GetAccountWallet(accessToken string, accountID string) (float64, error) {
	body := map[string]interface{}{"account_id": accountID}
	data, err := utils.HTTPPostJSON(walletURI, body, headers(accessToken))
	if err != nil {
		return 0, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Wallet struct {
				Balance int64 `json:"balance"`
			} `json:"wallet"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return 0, err
	}
	if res.Code != 0 {
		return 0, fmt.Errorf("腾讯获取钱包失败: code=%d %s", res.Code, res.Message)
	}
	return float64(res.Data.Wallet.Balance) / 100, nil
}
