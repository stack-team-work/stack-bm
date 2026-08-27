package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

// 快手开放API 公共封装：Access-Token 头、advertiser_id 注入、code==0 校验

const (
	listCampaignsURI = "https://ad.e.kuaishou.com/rest/openapi/gw/dsp/campaign/list"
	listUnitsURI     = "https://ad.e.kuaishou.com/rest/openapi/gw/dsp/unit/list"
	listCreativesURI = "https://ad.e.kuaishou.com/rest/openapi/gw/dsp/creative/list"
)

// 操作状态（对齐源 KuaiShouConstants）
const (
	KsStatusOpen   = 1
	KsStatusPause  = 2
	KsStatusDelete = 3
)

func headers(accessToken string) map[string]string {
	return map[string]string{"Access-Token": accessToken}
}

// post 快手 POST：注入 advertiser_id，成功 code==0
func post(accessToken string, advertiserID int, uri string, body map[string]interface{}) (bool, error) {
	body["advertiser_id"] = advertiserID
	data, err := utils.HTTPPostJSON(uri, body, headers(accessToken))
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

// list 分页拉取层级列表（快手 data 为裸数组）
func list(accessToken string, advertiserID int, uri string, page, pageSize int) ([]map[string]interface{}, error) {
	body := map[string]interface{}{"advertiser_id": advertiserID, "page": page, "page_size": pageSize}
	data, err := utils.HTTPPostJSON(uri, body, headers(accessToken))
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int                      `json:"code"`
		Message string                   `json:"message"`
		Data    []map[string]interface{} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("快手列表获取失败: code=%d %s", res.Code, res.Message)
	}
	return res.Data, nil
}
