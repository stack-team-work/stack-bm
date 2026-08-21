package media

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"stack-bm/internal/config"
	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/constants"
	"stack-bm/pkg/utils"
)

// ChannelAuthService 管家授权：生成授权 URL、分发回调
type ChannelAuthService struct {
	managerRepo *mediaRepo.MediaManagerRepository
	appRepo     *mediaRepo.MediaApplicationRepository
	mediaRepo   *mediaRepo.MediaRepository
	tokenRepo   *mediaRepo.MediaManagerTokenRepository
}

func NewChannelAuthService() *ChannelAuthService {
	return &ChannelAuthService{
		managerRepo: mediaRepo.NewMediaManagerRepository(),
		appRepo:     mediaRepo.NewMediaApplicationRepository(),
		mediaRepo:   mediaRepo.NewMediaRepository(),
		tokenRepo:   mediaRepo.NewMediaManagerTokenRepository(),
	}
}

// ResolveChannel 通过 manager 关联的 media.mark 解析渠道标识
func (s *ChannelAuthService) ResolveChannel(manager *media.MediaManager) (string, error) {
	if manager.MediaID == 0 {
		return "", errors.New("管家未关联媒体渠道")
	}
	m, err := s.mediaRepo.FindByID(uint(manager.MediaID))
	if err != nil {
		return "", errors.New("媒体渠道不存在")
	}
	channel, ok := constants.MediaMarkChannel[m.Mark]
	if !ok {
		return "", fmt.Errorf("媒体渠道【%s】暂不支持管家授权", m.Mark)
	}
	return channel, nil
}

// getApp 获取管家关联的应用配置
func (s *ChannelAuthService) getApp(manager *media.MediaManager) (*media.MediaApplication, error) {
	if manager.ApplicationID == 0 {
		return nil, errors.New("管家未绑定授权应用")
	}
	app, err := s.appRepo.FindByID(uint(manager.ApplicationID))
	if err != nil {
		return nil, errors.New("授权应用不存在")
	}
	if app.AppID == "" {
		return nil, errors.New("授权应用 app_id 未配置")
	}
	return app, nil
}

// redirectURI 拼接各渠道回调地址
func (s *ChannelAuthService) redirectURI(channel string) string {
	path := config.AppConfig.OAuth.RedirectPath
	if path == "" {
		path = "http://127.0.0.1:8080/oauth/callback/%s"
	}
	return fmt.Sprintf(path, channel)
}

// GetOauthUrl 生成授权跳转 URL
func (s *ChannelAuthService) GetOauthUrl(managerID uint) (string, error) {
	manager, err := s.managerRepo.FindByID(managerID)
	if err != nil {
		return "", errors.New("管家不存在")
	}
	channel, err := s.ResolveChannel(manager)
	if err != nil {
		return "", err
	}
	app, err := s.getApp(manager)
	if err != nil {
		return "", err
	}

	state := utils.XorEncrypt(fmt.Sprintf(`{"mkt_account_manager_id":%d}`, managerID), utils.XorKey)
	redirectURI := s.redirectURI(channel)

	switch channel {
	case constants.ChannelTt:
		return fmt.Sprintf("http://ad.toutiao.com/openapi/audit/oauth.html?app_id=%s&state=%s&redirect_uri=%s",
			url.QueryEscape(app.AppID), url.QueryEscape(state), url.QueryEscape(redirectURI)), nil
	case constants.ChannelTc:
		return fmt.Sprintf("https://developers.e.qq.com/oauth/authorize?client_id=%s&state=%s&redirect_uri=%s",
			url.QueryEscape(app.AppID), url.QueryEscape(state), url.QueryEscape(redirectURI)), nil
	case constants.ChannelBili:
		return fmt.Sprintf("https://ad.bilibili.com/developer/developer-authorization?client_id=%s&redirect_uri=%s&state=%s",
			url.QueryEscape(app.AppID), url.QueryEscape(redirectURI), url.QueryEscape(state)), nil
	case constants.ChannelKs:
		query := url.Values{}
		query.Set("app_id", app.AppID)
		query.Set("scope", `["report_service","account_service","ad_query","ad_manage"]`)
		query.Set("redirect_uri", redirectURI)
		query.Set("oauth_type", "advertiser")
		query.Set("state", state)
		return "https://developers.e.kuaishou.com/tools/authorize?" + query.Encode(), nil
	default:
		return "", fmt.Errorf("渠道【%s】暂未开发管家授权", channel)
	}
}

// FinishOauth 回调分发到对应渠道
func (s *ChannelAuthService) FinishOauth(channel string, params map[string]string) error {
	switch channel {
	case constants.ChannelTt:
		return s.ttFinishOauth(params)
	case constants.ChannelTc:
		return s.tcFinishOauth(params)
	case constants.ChannelBili:
		return s.biliFinishOauth(params)
	case constants.ChannelKs:
		return s.ksFinishOauth(params)
	default:
		return fmt.Errorf("渠道【%s】不支持授权", channel)
	}
}

// updateAuthStatus 更新管家授权状态，auth_message 存入 extra 扩展信息
func (s *ChannelAuthService) updateAuthStatus(managerID uint, status int8, authMessage string) error {
	manager, err := s.managerRepo.FindByID(managerID)
	if err != nil {
		return err
	}
	manager.AuthStatus = status
	if authMessage != "" {
		extra := map[string]interface{}{}
		if manager.Extra != "" {
			_ = json.Unmarshal([]byte(manager.Extra), &extra)
		}
		extra["auth_message"] = authMessage
		raw, _ := json.Marshal(extra)
		manager.Extra = string(raw)
	}
	return s.managerRepo.Update(manager)
}
