package api

const updateCreativeStatusURI = "https://ad.e.kuaishou.com/rest/openapi/v1/creative/update/status"

// ListCreatives 分页拉取快手创意列表
func ListCreatives(accessToken string, advertiserID int, page, pageSize int) ([]map[string]interface{}, error) {
	return list(accessToken, advertiserID, listCreativesURI, page, pageSize)
}

// UpdateCreativeStatus 创意启停（接口单条，取首个ID提交，批量由上层循环）
func UpdateCreativeStatus(accessToken string, advertiserID int, status int, creativeIDs []int) (bool, error) {
	body := map[string]interface{}{"creative_id": creativeIDs[0], "put_status": status}
	return post(accessToken, advertiserID, updateCreativeStatusURI, body)
}
