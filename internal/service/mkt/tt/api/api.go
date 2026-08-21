package api

import (
	"fmt"
	"net/url"

	"stack-bm/pkg/utils"
)

const (
	TokenURI      = "https://ad.oceanengine.com/open_api/oauth2/access_token/"
	AdvertiserURI = "https://ad.oceanengine.com/open_api/oauth2/advertiser/get/"
)

// BuildOauthURL 生成头条授权URL
func BuildOauthURL(appID, state, redirectURI string) string {
	return fmt.Sprintf("http://ad.toutiao.com/openapi/audit/oauth.html?app_id=%s&state=%s&redirect_uri=%s",
		url.QueryEscape(appID), url.QueryEscape(state), url.QueryEscape(redirectURI))
}

// TokenResult 头条 token 结果
type TokenResult struct {
	AccessToken       string
	ExpiresIn         int64
	RefreshToken      string
	RefreshTokenExpIn int64
	AdvertiserIDs     []string
}

// ExchangeToken 授权码换取 token
func ExchangeToken(appID, appSecret, authCode string) (*TokenResult, error) {
	body := map[string]interface{}{
		"app_id":     appID,
		"secret":     appSecret,
		"grant_type": "auth_code",
		"auth_code":  authCode,
	}
	data, err := utils.HTTPPostJSON(TokenURI, body, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Request string `json:"request_id"`
		Data    struct {
			AccessToken       string   `json:"access_token"`
			ExpiresIn         int64    `json:"expires_in"`
			RefreshToken      string   `json:"refresh_token"`
			RefreshTokenExpIn int64    `json:"refresh_token_expires_in"`
			AdvertiserIDs     []string `json:"advertiser_ids"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 || res.Data.AccessToken == "" {
		return nil, fmt.Errorf("头条token获取失败_%d_%s_%s", res.Code, res.Message, res.Request)
	}
	return &TokenResult{
		AccessToken:       res.Data.AccessToken,
		ExpiresIn:         res.Data.ExpiresIn,
		RefreshToken:      res.Data.RefreshToken,
		RefreshTokenExpIn: res.Data.RefreshTokenExpIn,
		AdvertiserIDs:     res.Data.AdvertiserIDs,
	}, nil
}

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