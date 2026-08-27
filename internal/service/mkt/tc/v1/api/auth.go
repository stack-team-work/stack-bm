package api

import (
	"fmt"
	"net/url"

	"stack-bm/pkg/utils"
)

const TokenURI = "https://api.e.qq.com/oauth/token"

// BuildOauthURL 生成腾讯授权URL
func BuildOauthURL(appID, state, redirectURI string) string {
	return fmt.Sprintf("https://developers.e.qq.com/oauth/authorize?client_id=%s&state=%s&redirect_uri=%s",
		url.QueryEscape(appID), url.QueryEscape(state), url.QueryEscape(redirectURI))
}

// TokenResult 腾讯 token 结果
type TokenResult struct {
	AccessToken       string
	RefreshToken      string
	AccessTokenExpIn  int64
	RefreshTokenExpIn int64
}

// ExchangeToken 授权码换取 token
func ExchangeToken(appID, appSecret, code, redirectURI string) (*TokenResult, error) {
	query := url.Values{}
	query.Set("client_id", appID)
	query.Set("client_secret", appSecret)
	query.Set("grant_type", "authorization_code")
	query.Set("authorization_code", code)
	query.Set("redirect_uri", redirectURI)

	data, err := utils.HTTPGet(TokenURI, query, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		MsgCn   string `json:"message_cn"`
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
		return nil, fmt.Errorf("腾讯token获取失败_%d_%s_%s", res.Code, res.Message, res.MsgCn)
	}
	return &TokenResult{
		AccessToken:       res.Data.AccessToken,
		RefreshToken:      res.Data.RefreshToken,
		AccessTokenExpIn:  res.Data.AccessTokenExpIn,
		RefreshTokenExpIn: res.Data.RefreshTokenExpIn,
	}, nil
}
