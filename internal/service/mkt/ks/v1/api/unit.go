package api

const (
	updateUnitStatusURI = "https://ad.e.kuaishou.com/rest/openapi/v1/ad_unit/update/status"
	updateUnitBudgetURI = "https://ad.e.kuaishou.com/rest/openapi/v1/ad_unit/update/day_budget"
	updateUnitBidURI    = "https://ad.e.kuaishou.com/rest/openapi/v1/ad_unit/update/bid"
)

// ListUnits 分页拉取快手广告列表
func ListUnits(accessToken string, advertiserID int, page, pageSize int) ([]map[string]interface{}, error) {
	return list(accessToken, advertiserID, listUnitsURI, page, pageSize)
}

// UpdateUnitStatus 广告启停/删除（接口单条，取首个ID提交，批量由上层循环）
func UpdateUnitStatus(accessToken string, advertiserID int, status int, unitIDs []int) (bool, error) {
	body := map[string]interface{}{"unit_id": unitIDs[0], "put_status": status}
	return post(accessToken, advertiserID, updateUnitStatusURI, body)
}

// UpdateUnitBudget 广告日预算（元 -> 分）
func UpdateUnitBudget(accessToken string, advertiserID, unitID int, dayBudget float64) (bool, error) {
	body := map[string]interface{}{"unit_id": unitID, "day_budget": int(dayBudget * 100)}
	return post(accessToken, advertiserID, updateUnitBudgetURI, body)
}

// UpdateUnitBid 广告出价/深度出价（元 -> 分）
func UpdateUnitBid(accessToken string, advertiserID, unitID int, bid float64, deepBid bool) (bool, error) {
	body := map[string]interface{}{"unit_id": unitID}
	if deepBid {
		body["deep_conversion_bid"] = int(bid * 100)
	} else {
		body["bid"] = int(bid * 100)
	}
	return post(accessToken, advertiserID, updateUnitBidURI, body)
}
