package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	adgroupListURI   = "https://api.e.qq.com/v3.0/adgroups/get"
	updateAdgroupURI = "https://api.e.qq.com/v3.0/adgroups/update"
	deleteAdgroupURI = "https://api.e.qq.com/v3.0/adgroups/delete"
)

// 腾讯V3 广告组 configured_status
const (
	TcAdStatusPause   = 1 // 暂停
	TcAdStatusNormal  = 2 // 启用
	TcAdStatusDeleted = 3 // 删除
)

// post 腾讯 POST：注入 account_id（字符串），成功 code==0
func post(accessToken string, accountID string, uri string, body map[string]interface{}) (bool, error) {
	body["account_id"] = accountID
	data, err := utils.HTTPPostJSON(uri, body, headers(accessToken))
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

// ListAdgroups 分页拉取腾讯广告组列表
func ListAdgroups(accessToken string, accountID string, page, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{"account_id": accountID, "page": page, "page_size": pageSize}
	data, err := utils.HTTPPostJSON(adgroupListURI, body, headers(accessToken))
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
		return nil, fmt.Errorf("腾讯广告组列表获取失败: code=%d %s", res.Code, res.Message)
	}
	return res.Data.List, nil
}

// UpdateAdgroupStatus 广告组启停
func UpdateAdgroupStatus(accessToken string, accountID string, adgroupID int, configuredStatus int) (bool, error) {
	return post(accessToken, accountID, updateAdgroupURI, map[string]interface{}{
		"adgroup_id":        adgroupID,
		"configured_status": configuredStatus,
	})
}

// UpdateAdgroupBudget 广告组预算（元 -> 分）
func UpdateAdgroupBudget(accessToken string, accountID string, adgroupID int, budget float64) (bool, error) {
	return post(accessToken, accountID, updateAdgroupURI, map[string]interface{}{
		"adgroup_id":   adgroupID,
		"daily_budget": int(budget * 100),
	})
}

// UpdateAdgroupBid 广告组出价（元 -> 分）
func UpdateAdgroupBid(accessToken string, accountID string, adgroupID int, bid float64) (bool, error) {
	return post(accessToken, accountID, updateAdgroupURI, map[string]interface{}{
		"adgroup_id": adgroupID,
		"bid_amount": int(bid * 100),
	})
}

// DeleteAdgroup 删除广告组
func DeleteAdgroup(accessToken string, accountID string, adgroupID int) (bool, error) {
	return post(accessToken, accountID, deleteAdgroupURI, map[string]interface{}{
		"adgroup_id": adgroupID,
	})
}
