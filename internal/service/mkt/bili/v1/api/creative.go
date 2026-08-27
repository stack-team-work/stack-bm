package api

const creativeListURI = "https://cm.bilibili.com/takumi/api/open_api/cpc/v3/creative/list_creatives"

// ListCreatives 分页拉取B站创意列表
func ListCreatives(accessToken string, accountID, page, pageSize int) ([]map[string]interface{}, error) {
	return listAll(accessToken, creativeListURI, map[string]interface{}{
		"account_id": accountID,
		"page":       page,
		"page_size":  pageSize,
		"status":     "ALL",
	})
}

// UpdateCreativeStatus 创意启停/删除（批量）
func UpdateCreativeStatus(accessToken string, accountID int, operationType int, creativeIDs []int) (bool, error) {
	body := map[string]interface{}{
		"operation_type": operationType,
		"operations":     []map[string]interface{}{{"creative_ids": creativeIDs}},
	}
	return updateGeneric(accessToken, accountID, updateCreativeURI, body)
}
