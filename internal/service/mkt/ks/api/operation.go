package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	ksUpdateCampaignStatusURI = "https://ad.e.kuaishou.com/rest/openapi/v1/campaign/update/status"
	ksUpdateUnitStatusURI     = "https://ad.e.kuaishou.com/rest/openapi/v1/ad_unit/update/status"
	ksUpdateUnitBudgetURI     = "https://ad.e.kuaishou.com/rest/openapi/v1/ad_unit/update/day_budget"
	ksUpdateUnitBidURI        = "https://ad.e.kuaishou.com/rest/openapi/v1/ad_unit/update/bid"
	ksUpdateCreativeStatusURI = "https://ad.e.kuaishou.com/rest/openapi/v1/creative/update/status"
)

// 快手操作状态（对齐源 KuaiShouConstants）
const (
	KsStatusOpen   = 1
	KsStatusPause  = 2
	KsStatusDelete = 3
)

// ksPost 快手 POST，Access-Token 头 + advertiser_id 参数，成功 code==0
func ksPost(accessToken string, advertiserID int, uri string, body map[string]interface{}) (bool, error) {
	body["advertiser_id"] = advertiserID
	data, err := utils.HTTPPostJSON(uri, body, map[string]string{"Content-Type": "application/json", "Access-Token": accessToken})
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
		return false, fmt.Errorf("快手操作失败: code=%d %s", res.Code, res.Message)
	}
	return true, nil
}

// UpdateCampaignStatus 广告组启停/删除
func UpdateCampaignStatus(accessToken string, advertiserID int, status int, campaignIDs []int) (bool, error) {
	body := map[string]interface{}{"campaign_id": campaignIDs[0], "put_status": status}
	return ksPost(accessToken, advertiserID, ksUpdateCampaignStatusURI, body)
}

// UpdateUnitStatus 广告启停/删除
func UpdateUnitStatus(accessToken string, advertiserID int, status int, unitIDs []int) (bool, error) {
	body := map[string]interface{}{"unit_id": unitIDs[0], "put_status": status}
	return ksPost(accessToken, advertiserID, ksUpdateUnitStatusURI, body)
}

// UpdateUnitBudget 广告预算
func UpdateUnitBudget(accessToken string, advertiserID, unitID int, dayBudget float64) (bool, error) {
	body := map[string]interface{}{"unit_id": unitID, "day_budget": int(dayBudget * 100)}
	return ksPost(accessToken, advertiserID, ksUpdateUnitBudgetURI, body)
}

// UpdateUnitBid 广告出价
func UpdateUnitBid(accessToken string, advertiserID, unitID int, bid float64, deepBid bool) (bool, error) {
	body := map[string]interface{}{"unit_id": unitID}
	if deepBid {
		body["deep_conversion_bid"] = int(bid * 100)
	} else {
		body["bid"] = int(bid * 100)
	}
	return ksPost(accessToken, advertiserID, ksUpdateUnitBidURI, body)
}

// UpdateCreativeStatus 创意启停
func UpdateCreativeStatus(accessToken string, advertiserID int, status int, creativeIDs []int) (bool, error) {
	body := map[string]interface{}{"creative_id": creativeIDs[0], "put_status": status}
	return ksPost(accessToken, advertiserID, ksUpdateCreativeStatusURI, body)
}