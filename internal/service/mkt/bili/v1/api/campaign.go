package api

const campaignListURI = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/campaign/list_campaigns"

// ListCampaigns 分页拉取B站计划列表
func ListCampaigns(accessToken string, accountID, page, pageSize int) ([]map[string]interface{}, error) {
	return listAll(accessToken, campaignListURI, map[string]interface{}{
		"account_id": accountID,
		"page":       page,
		"page_size":  pageSize,
		"status":     "ALL",
	})
}

// UpdateCampaignStatus 计划启停/删除（批量）
func UpdateCampaignStatus(accessToken string, accountID int, operationType int, campaignIDs []int) (bool, error) {
	body := map[string]interface{}{
		"operation_type": operationType,
		"operations":     []map[string]interface{}{{"campaign_ids": campaignIDs}},
	}
	return updateGeneric(accessToken, accountID, updateCampaignURI, body)
}

// UpdateCampaignBudget 修改计划预算
func UpdateCampaignBudget(accessToken string, accountID, campaignID int, budget float64) (bool, error) {
	body := map[string]interface{}{
		"operation_type": BiliOpBudget,
		"operations":     []map[string]interface{}{{"campaign_ids": []int{campaignID}, "budget": budgetOp(budget)}},
	}
	return updateGeneric(accessToken, accountID, updateCampaignURI, body)
}

// UpdateCampaignBid 修改计划出价
func UpdateCampaignBid(accessToken string, accountID, campaignID int, bid float64, deepBid bool) (bool, error) {
	body := map[string]interface{}{
		"operation_type": BiliOpBid,
		"operations":     []map[string]interface{}{{"campaign_ids": []int{campaignID}, "bid": bidOp(bid, deepBid)}},
	}
	return updateGeneric(accessToken, accountID, updateCampaignURI, body)
}
