package tt

import (
	"errors"
	"time"

	mediaModel "stack-bm/internal/model/mkt/media"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tt/api"
	"stack-bm/pkg/constants"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// OauthService 头条管家授权
type OauthService struct {
	auth *oauth.ManagerAuth
}

func NewOauthService(auth *oauth.ManagerAuth) *OauthService {
	return &OauthService{auth: auth}
}

// BuildOauthURL 生成授权跳转URL
func (s *OauthService) BuildOauthURL(appID, state, redirectURI string) string {
	return api.BuildOauthURL(appID, state, redirectURI)
}

// FinishOauth 回调完成授权
func (s *OauthService) FinishOauth(params map[string]string) error {
	authCode := params["auth_code"]
	if authCode == "" {
		return errors.New("授权code不存在")
	}
	managerID, err := s.auth.ParseManagerFromState(params["state"])
	if err != nil {
		return err
	}
	manager, err := s.auth.GetManager(managerID)
	if err != nil {
		return err
	}
	app, err := s.auth.GetApp(manager)
	if err != nil {
		return err
	}

	token, err := api.ExchangeToken(app.AppID, app.AppSecret, authCode)
	if err != nil {
		_ = s.auth.UpdateAuthStatus(managerID, constants.ManagerAuthStatusAuthFail, err.Error())
		return err
	}

	now := time.Now().Unix()
	doc := bson.M{
		"mkt_account_manager_id":     managerID,
		"channel":                    constants.ChannelTt,
		"access_token":               token.AccessToken,
		"access_token_expires_in":    token.ExpiresIn,
		"access_token_expires_time":  now + token.ExpiresIn - 300,
		"refresh_token":              token.RefreshToken,
		"refresh_token_expires_in":   token.RefreshTokenExpIn,
		"refresh_token_expires_time": now + token.RefreshTokenExpIn - 300,
		"advertiser_ids":             token.AdvertiserIDs,
		"updated_time":               time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.auth.TokenRepo.Upsert(int(managerID), doc); err != nil {
		return err
	}

	_ = s.auth.UpdateAuthStatus(managerID, constants.ManagerAuthStatusComplete, "")
	s.auth.PushSyncQueue(managerID)
	return nil
}

// SyncAdvertiser 同步头条广告主列表
func (s *OauthService) SyncAdvertiser(manager *mediaModel.MediaManager, authInfo bson.M) error {
	accessToken, _ := authInfo["access_token"].(string)
	app, err := s.auth.GetApp(manager)
	if err != nil {
		return err
	}
	ids, err := api.GetAdvertiserList(accessToken, app.AppID, app.AppSecret)
	if err != nil {
		return err
	}
	_ = s.auth.UpdateAccountNum(manager, len(ids))
	return s.syncChannelAccounts(manager.ID)
}

// syncChannelAccounts 预留：同步更新渠道账户（当前渠道账户功能尚未开发）
func (s *OauthService) syncChannelAccounts(managerID uint) error {
	return nil
}