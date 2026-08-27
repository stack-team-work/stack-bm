package api

const updateCampaignStatusURI = "https://ad.e.kuaishou.com/rest/openapi/v1/campaign/update/status"

// ListCampaigns 分页拉取快手广告组列表
func ListCampaigns(accessToken string, advertiserID int, page, pageSize int) ([]map[string]interface{}, error) {
	return list(accessToken, advertiserID, listCampaignsURI, page, pageSize)
}

// UpdateCampaignStatus 广告组启停/删除（接口单条，取首个ID提交，批量由上层循环）
func UpdateCampaignStatus(accessToken string, advertiserID int, status int, campaignIDs []int) (bool, error) {
	body := map[string]interface{}{"campaign_id": campaignIDs[0], "put_status": status}
	return post(accessToken, advertiserID, updateCampaignStatusURI, body)
}
