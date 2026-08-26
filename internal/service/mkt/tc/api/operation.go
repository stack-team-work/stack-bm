package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	tcUpdateAdgroupURI = "https://api.e.qq.com/v3.0/adgroups/update"
	tcDeleteAdgroupURI = "https://api.e.qq.com/v3.0/adgroups/delete"
)

// 腾讯V3 广告组 configured_status
const (
	TcAdStatusPause   = 1 // 暂停
	TcAdStatusNormal  = 2 // 启用
	TcAdStatusDeleted = 3 // 删除
)

// tcPost 腾讯 POST，Access-Token 头 + account_id 参数，成功 code==0
func tcPost(accessToken string, accountID string, uri string, body map[string]interface{}) (bool, error) {
	body["account_id"] = accountID
	data, err := utils.HTTPPostJSON(uri, body, map[string]string{"Access-Token": accessToken})
	if err != nil {
		return false, err
	}
	var res struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		MessageCn string `json:"message_cn"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return false, err
	}
	if res.Code != 0 {
		return false, fmt.Errorf("腾讯操作失败: code=%d %s%s", res.Code, res.Message, res.MessageCn)
	}
	return true, nil
}

// UpdateAdgroupStatus 广告组启停
func UpdateAdgroupStatus(accessToken string, accountID string, adgroupID int, configuredStatus int) (bool, error) {
	return tcPost(accessToken, accountID, tcUpdateAdgroupURI, map[string]interface{}{
		"adgroup_id":        adgroupID,
		"configured_status": configuredStatus,
	})
}

// UpdateAdgroupBudget 广告组预算（单位：分）
func UpdateAdgroupBudget(accessToken string, accountID string, adgroupID int, budget float64) (bool, error) {
	return tcPost(accessToken, accountID, tcUpdateAdgroupURI, map[string]interface{}{
		"adgroup_id":   adgroupID,
		"daily_budget": int(budget * 100),
	})
}

// UpdateAdgroupBid 广告组出价（单位：分）
func UpdateAdgroupBid(accessToken string, accountID string, adgroupID int, bid float64) (bool, error) {
	return tcPost(accessToken, accountID, tcUpdateAdgroupURI, map[string]interface{}{
		"adgroup_id": adgroupID,
		"bid_amount": int(bid * 100),
	})
}

// DeleteAdgroup 删除广告组
func DeleteAdgroup(accessToken string, accountID string, adgroupID int) (bool, error) {
	return tcPost(accessToken, accountID, tcDeleteAdgroupURI, map[string]interface{}{
		"adgroup_id": adgroupID,
	})
}