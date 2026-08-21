package api

import (
	"fmt"
	"net/url"

	"stack-bm/pkg/utils"
)

const TokenURI = "https://ad.e.kuaishou.com/rest/openapi/oauth2/authorize/access_token"

// BuildOauthURL 生成快手授权URL
func BuildOauthURL(appID, state, redirectURI string) string {
	query := url.Values{}
	query.Set("app_id", appID)
	query.Set("scope", `["report_service","account_service","ad_query","ad_manage"]`)
	query.Set("redirect_uri", redirectURI)
	query.Set("oauth_type", "advertiser")
	query.Set("state", state)
	return "https://developers.e.kuaishou.com/tools/authorize?" + query.Encode()
}

// TokenResult 快手 token 结果
type TokenResult struct {
	AccessToken       string
	RefreshToken      string
	AccessTokenExpIn  int64
	RefreshTokenExpIn int64
}

// ExchangeToken 授权码换取 token
func ExchangeToken(appID, appSecret, authCode string) (*TokenResult, error) {
	body := map[string]interface{}{
		"app_id":   appID,
		"secret":   appSecret,
		"auth_code": authCode,
	}
	data, err := utils.HTTPPostJSON(TokenURI, body, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			AccessToken       string `json:"access_token"`
			RefreshToken      string `json:"refresh_token"`
			AccessTokenExpIn  int64  `json:"access_token_expires_in"`
			RefreshTokenExpIn int64  `json:"refresh_token_expires_in"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 || res.Data.AccessToken == "" {
		return nil, fmt.Errorf("快手token获取失败_%d_%s", res.Code, res.Message)
	}
	return &TokenResult{
		AccessToken:       res.Data.AccessToken,
		RefreshToken:      res.Data.RefreshToken,
		AccessTokenExpIn:  res.Data.AccessTokenExpIn,
		RefreshTokenExpIn: res.Data.RefreshTokenExpIn,
	}, nil
}