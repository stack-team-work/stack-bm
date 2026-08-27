package api

import (
	"fmt"
	"net/url"

	"stack-bm/pkg/utils"
)

const TokenURI = "https://cm.bilibili.com/takumi/api/open_api/oauth2/token"

// BuildOauthURL 生成B站授权URL
func BuildOauthURL(appID, state, redirectURI string) string {
	return fmt.Sprintf("https://ad.bilibili.com/developer/developer-authorization?client_id=%s&redirect_uri=%s&state=%s",
		url.QueryEscape(appID), url.QueryEscape(redirectURI), url.QueryEscape(state))
}

// TokenResult B站 token 结果
type TokenResult struct {
	AccessToken  string
	ExpiresIn    int64
	RefreshToken string
	TokenType    string
}

// ExchangeToken 授权码换取 token
func ExchangeToken(appID, appSecret, code, redirectURI string) (*TokenResult, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("client_id", appID)
	form.Set("client_secret", appSecret)
	form.Set("redirect_uri", redirectURI)

	data, err := utils.HTTPPostForm(TokenURI, form, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		Message string `json:"message"`
		Data    struct {
			AccessToken  string `json:"access_token"`
			ExpiresIn    int64  `json:"expires_in"`
			RefreshToken string `json:"refresh_token"`
			TokenType    string `json:"token_type"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 || res.Data.AccessToken == "" {
		return nil, fmt.Errorf("B站token获取失败_%d_%s_%s", res.Code, res.Msg, res.Message)
	}
	return &TokenResult{
		AccessToken:  res.Data.AccessToken,
		ExpiresIn:    res.Data.ExpiresIn,
		RefreshToken: res.Data.RefreshToken,
		TokenType:    res.Data.TokenType,
	}, nil
}
