package tc

import (
	"errors"
	"time"

	mediaModel "stack-bm/internal/model/mkt/media"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tc/api"
	"stack-bm/pkg/constants"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// OauthService 腾讯管家授权
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
	authCode := params["authorization_code"]
	if authCode == "" {
		return errors.New("授权authorization_code不存在")
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
	redirectURI := s.auth.RedirectURI(constants.ChannelTc)

	token, err := api.ExchangeToken(app.AppID, app.AppSecret, authCode, redirectURI)
	if err != nil {
		_ = s.auth.UpdateAuthStatus(managerID, constants.ManagerAuthStatusAuthFail, err.Error())
		return err
	}

	now := time.Now().Unix()
	doc := bson.M{
		"mkt_account_manager_id":     managerID,
		"channel":                    constants.ChannelTc,
		"access_token":               token.AccessToken,
		"access_token_expires_in":    token.AccessTokenExpIn,
		"access_token_expires_time":  now + token.AccessTokenExpIn - 300,
		"refresh_token":              token.RefreshToken,
		"refresh_token_expires_in":   token.RefreshTokenExpIn,
		"refresh_token_expires_time": now + token.RefreshTokenExpIn - 300,
		"updated_time":               time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.auth.TokenRepo.Upsert(int(managerID), doc); err != nil {
		return err
	}

	_ = s.auth.UpdateAuthStatus(managerID, constants.ManagerAuthStatusComplete, "")
	s.auth.PushSyncQueue(managerID)
	return nil
}

// SyncAdvertiser 同步腾讯广告主列表（预留实现）
func (s *OauthService) SyncAdvertiser(manager *mediaModel.MediaManager, authInfo bson.M) error {
	_ = s.auth.UpdateAccountNum(manager, 0)
	return s.syncChannelAccounts(manager.ID)
}

// syncChannelAccounts 预留：同步更新渠道账户（当前渠道账户功能尚未开发）
func (s *OauthService) syncChannelAccounts(managerID uint) error {
	return nil
}