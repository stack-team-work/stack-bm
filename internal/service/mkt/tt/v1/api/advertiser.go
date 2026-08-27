package api

import (
	"fmt"
	"net/url"

	"stack-bm/pkg/utils"
)

const AdvertiserURI = "https://ad.oceanengine.com/open_api/oauth2/advertiser/get/"

// GetAdvertiserList 获取头条广告主列表
func GetAdvertiserList(accessToken, appID, appSecret string) ([]string, error) {
	query := url.Values{}
	query.Set("access_token", accessToken)
	query.Set("app_id", appID)
	query.Set("secret", appSecret)

	data, err := utils.HTTPGet(AdvertiserURI, query, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			List []struct {
				AdvertiserID string `json:"advertiser_id"`
			} `json:"list"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("头条广告主列表获取失败_%d_%s", res.Code, res.Message)
	}
	ids := make([]string, 0, len(res.Data.List))
	for _, item := range res.Data.List {
		ids = append(ids, item.AdvertiserID)
	}
	return ids, nil
}
