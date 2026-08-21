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

const tcTokenURI = "https://api.e.qq.com/oauth/token"

func (s *ChannelAuthService) tcFinishOauth(params map[string]string) error {
	authCode := params["authorization_code"]
	if authCode == "" {
		return errors.New("授权authorization_code不存在")
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

	query := url.Values{}
	query.Set("client_id", app.AppID)
	query.Set("client_secret", app.AppSecret)
	query.Set("grant_type", "authorization_code")
	query.Set("authorization_code", authCode)
	query.Set("redirect_uri", s.redirectURI(constants.ChannelTc))

	data, err := utils.HTTPGet(tcTokenURI, query, nil)
	if err != nil {
		return err
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
		return err
	}
	if res.Code != 0 || res.Data.AccessToken == "" {
		msg := fmt.Sprintf("腾讯账户管家授权失败_%d_%s_%s", res.Code, res.Message, res.MsgCn)
		_ = s.updateAuthStatus(managerID, constants.ManagerAuthStatusAuthFail, msg)
		return errors.New(msg)
	}

	now := time.Now().Unix()
	doc := bson.M{
		"mkt_account_manager_id":     managerID,
		"channel":                    constants.ChannelTc,
		"access_token":               res.Data.AccessToken,
		"access_token_expires_in":    res.Data.AccessTokenExpIn,
		"access_token_expires_time":  now + res.Data.AccessTokenExpIn - 300,
		"refresh_token":              res.Data.RefreshToken,
		"refresh_token_expires_in":   res.Data.RefreshTokenExpIn,
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
