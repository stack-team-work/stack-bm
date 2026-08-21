package media

import (
	"errors"
	"fmt"
	"time"

	"stack-bm/pkg/constants"
	"stack-bm/pkg/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const ksTokenURI = "https://ad.e.kuaishou.com/rest/openapi/oauth2/authorize/access_token"

func (s *ChannelAuthService) ksFinishOauth(params map[string]string) error {
	authCode := params["auth_code"]
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

	body := map[string]interface{}{
		"app_id":   app.AppID,
		"secret":   app.AppSecret,
		"auth_code": authCode,
	}
	data, err := utils.HTTPPostJSON(ksTokenURI, body, nil)
	if err != nil {
		return err
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
		return err
	}
	if res.Code != 0 || res.Data.AccessToken == "" {
		msg := fmt.Sprintf("快手账户管家授权失败_%d_%s", res.Code, res.Message)
		_ = s.updateAuthStatus(managerID, constants.ManagerAuthStatusAuthFail, msg)
		return errors.New(msg)
	}

	now := time.Now().Unix()
	doc := bson.M{
		"mkt_account_manager_id":     managerID,
		"channel":                    constants.ChannelKs,
		"access_token":               res.Data.AccessToken,
		"access_token_expires_in":    now + res.Data.AccessTokenExpIn - 300,
		"access_token_expires_time":  now + res.Data.AccessTokenExpIn - 300,
		"refresh_token":              res.Data.RefreshToken,
		"refresh_token_expires_in":   now + res.Data.RefreshTokenExpIn - 300,
		"refresh_token_expires_time": now + res.Data.RefreshTokenExpIn - 300,
		"updated_time":               time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.tokenRepo.Upsert(int(managerID), doc); err != nil {
		return err
	}

	_ = s.updateAuthStatus(managerID, constants.ManagerAuthStatusComplete, "")
	s.pushSyncQueue(managerID)
	return nil
}
