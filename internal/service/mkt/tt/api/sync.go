package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	ttProjectListURI = "https://ad.oceanengine.com/open_api/v3.0/project/get/"
)

// ListProjects 拉取头条V3项目列表
func ListProjects(accessToken string, advertiserID string, page, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{
		"advertiser_id": advertiserID,
		"page":          page,
		"page_size":     pageSize,
		"filtering":     map[string]interface{}{"status": "ALL"},
	}
	data, err := utils.HTTPPostJSON(ttProjectListURI, body, map[string]string{"Access-Token": accessToken})
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
		return nil, fmt.Errorf("头条项目列表获取失败: code=%d", res.Code)
	}
	return res.Data.List, nil
}