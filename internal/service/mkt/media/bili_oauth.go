package media

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"stack-bm/pkg/constants"
	"stack-bm/pkg/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	biliTokenURI    = "https://cm.bilibili.com/takumi/api/open_api/oauth2/token"
	biliSubjectsURI = "https://cm.bilibili.com/takumi/api/open_api/oauth2/subjects"
)

func (s *ChannelAuthService) biliFinishOauth(params map[string]string) error {
	authCode := params["code"]
	if authCode == "" {
		return errors.New("授权code不存在")
	}
	managerID, err := s.parseManagerFromState(params["state"])
	if err != nil {
		return err
	}
	manager, err := s.managerRepo.FindByID(managerID)
	if err != nil {
		return errors.New("管家不存在")
	}
	app, err := s.getApp(manager)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", authCode)
	form.Set("client_id", app.AppID)
	form.Set("client_secret", app.AppSecret)
	form.Set("redirect_uri", s.redirectURI(constants.ChannelBili))

	data, err := utils.HTTPPostForm(biliTokenURI, form, nil)
	if err != nil {
		return err
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
		return err
	}
	if res.Code != 0 || res.Data.AccessToken == "" {
		msg := fmt.Sprintf("B站账户管家授权失败_%d_%s_%s", res.Code, res.Msg, res.Message)
		_ = s.updateAuthStatus(managerID, constants.ManagerAuthStatusAuthFail, msg)
		return errors.New(msg)
	}

	now := time.Now().Unix()
	refreshExpIn := int64(20 * 86400)
	doc := bson.M{
		"mkt_account_manager_id":     managerID,
		"channel":                    constants.ChannelBili,
		"token_type":                 res.Data.TokenType,
		"access_token":               res.Data.AccessToken,
		"access_token_expires_in":    res.Data.ExpiresIn,
		"access_token_expires_time":  now + res.Data.ExpiresIn - 300,
		"refresh_token":              res.Data.RefreshToken,
		"refresh_token_expires_in":   refreshExpIn,
		"refresh_token_expires_time": now + refreshExpIn - 300,
		"updated_time":               time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.tokenRepo.Upsert(int(managerID), doc); err != nil {
		return err
	}

	_ = s.updateAuthStatus(managerID, constants.ManagerAuthStatusComplete, "")
	s.pushSyncQueue(managerID)
	return nil
}

// biliGetSubjectList 获取B站授权主体（广告主）
func (s *ChannelAuthService) biliGetSubjectList(accessToken string) ([]interface{}, error) {
	headers := map[string]string{"Authorization": "Bearer " + accessToken}
	data, err := utils.HTTPGet(biliSubjectsURI, nil, headers)
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("B站获取主体失败_%d_%s", res.Code, res.Message)
	}
	return res.Data, nil
}
