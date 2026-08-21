package media

import (
	"errors"
	"fmt"
	"time"

	"stack-bm/pkg/constants"
	"stack-bm/pkg/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const ttTokenURI = "https://ad.oceanengine.com/open_api/oauth2/access_token/"

func (s *ChannelAuthService) ttFinishOauth(params map[string]string) error {
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
		"app_id":     app.AppID,
		"secret":     app.AppSecret,
		"grant_type": "auth_code",
		"auth_code":  authCode,
	}
	data, err := utils.HTTPPostJSON(ttTokenURI, body, nil)
	if err != nil {
		return err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Request string `json:"request_id"`
		Data    struct {
			AccessToken          string   `json:"access_token"`
			ExpiresIn            int64    `json:"expires_in"`
			RefreshToken         string   `json:"refresh_token"`
			RefreshTokenExpIn    int64    `json:"refresh_token_expires_in"`
			AdvertiserIDs        []string `json:"advertiser_ids"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return err
	}
	if res.Code != 0 || res.Data.AccessToken == "" {
		msg := fmt.Sprintf("头条账户管家授权失败_%d_%s_%s", res.Code, res.Message, res.Request)
		_ = s.updateAuthStatus(managerID, constants.ManagerAuthStatusAuthFail, msg)
		return errors.New(msg)
	}

	now := time.Now().Unix()
	doc := bson.M{
		"mkt_account_manager_id":     managerID,
		"channel":                    constants.ChannelTt,
		"access_token":               res.Data.AccessToken,
		"access_token_expires_in":    res.Data.ExpiresIn,
		"access_token_expires_time":  now + res.Data.ExpiresIn - 300,
		"refresh_token":              res.Data.RefreshToken,
		"refresh_token_expires_in":   res.Data.RefreshTokenExpIn,
		"refresh_token_expires_time": now + res.Data.RefreshTokenExpIn - 300,
		"advertiser_ids":             res.Data.AdvertiserIDs,
		"updated_time":               time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.tokenRepo.Upsert(int(managerID), doc); err != nil {
		return err
	}

	_ = s.updateAuthStatus(managerID, constants.ManagerAuthStatusComplete, "")
	s.pushSyncQueue(managerID)
	return nil
}

// parseManagerFromState 解密 state 获取 managerID
func (s *ChannelAuthService) parseManagerFromState(state string) (uint, error) {
	if state == "" {
		return 0, errors.New("授权state不存在")
	}
	m, err := utils.XorDecryptJSON(state, utils.XorKey)
	if err != nil {
		return 0, errors.New("state解析失败")
	}
	id, ok := m["mkt_account_manager_id"].(float64)
	if !ok || id == 0 {
		return 0, errors.New("state中管家ID无效")
	}
	return uint(id), nil
}
