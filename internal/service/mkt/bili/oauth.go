package bili

import (
	"errors"
	"time"

	"stack-bm/internal/service/mkt/bili/api"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/pkg/constants"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// OauthService B站管家授权
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
	authCode := params["code"]
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
	redirectURI := s.auth.RedirectURI(constants.ChannelBili)

	token, err := api.ExchangeToken(app.AppID, app.AppSecret, authCode, redirectURI)
	if err != nil {
		_ = s.auth.UpdateAuthStatus(managerID, constants.ManagerAuthStatusAuthFail, err.Error())
		return err
	}

	now := time.Now().Unix()
	refreshExpIn := int64(20 * 86400)
	doc := bson.M{
		"mkt_account_manager_id":     managerID,
		"channel":                    constants.ChannelBili,
		"token_type":                 token.TokenType,
		"access_token":               token.AccessToken,
		"access_token_expires_in":    token.ExpiresIn,
		"access_token_expires_time":  now + token.ExpiresIn - 300,
		"refresh_token":              token.RefreshToken,
		"refresh_token_expires_in":   refreshExpIn,
		"refresh_token_expires_time": now + refreshExpIn - 300,
		"updated_time":               time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.auth.TokenRepo.Upsert(int(managerID), doc); err != nil {
		return err
	}

	_ = s.auth.UpdateAuthStatus(managerID, constants.ManagerAuthStatusComplete, "")
	s.auth.PushSyncQueue(managerID)
	return nil
}