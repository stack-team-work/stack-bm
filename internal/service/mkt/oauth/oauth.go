package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"stack-bm/internal/config"
	"stack-bm/internal/database"
	mediaModel "stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/constants"
	"stack-bm/pkg/utils"
)

// ManagerAuth 管家授权公共操作，各媒体 service 共享
type ManagerAuth struct {
	ManagerRepo *mediaRepo.MediaManagerRepository
	AppRepo     *mediaRepo.MediaApplicationRepository
	MediaRepo   *mediaRepo.MediaRepository
	TokenRepo   *mediaRepo.MediaManagerTokenRepository
	AccountRepo *mediaRepo.MediaAccountRepository
}

func NewManagerAuth() *ManagerAuth {
	return &ManagerAuth{
		ManagerRepo: mediaRepo.NewMediaManagerRepository(),
		AppRepo:     mediaRepo.NewMediaApplicationRepository(),
		MediaRepo:   mediaRepo.NewMediaRepository(),
		TokenRepo:   mediaRepo.NewMediaManagerTokenRepository(),
		AccountRepo: mediaRepo.NewMediaAccountRepository(),
	}
}

// GetAccessToken 通过管家ID获取授权 access_token（读 mongo mkt_account_manager_token）
func (m *ManagerAuth) GetAccessToken(managerID uint) (string, error) {
	if managerID == 0 {
		return "", errors.New("管家ID无效")
	}
	authInfo, err := m.TokenRepo.FindByManagerID(int(managerID))
	if err != nil {
		return "", err
	}
	if authInfo == nil {
		return "", errors.New("管家授权信息不存在")
	}
	token, _ := authInfo["access_token"].(string)
	if token == "" {
		return "", errors.New("管家未授权或token为空")
	}
	return token, nil
}

// ResolveManagerByAccount 渠道账户 -> 管家ID（1:1，通过 media_manager_manager_id）
func (m *ManagerAuth) ResolveManagerByAccount(accountID uint) (uint, error) {
	if accountID == 0 {
		return 0, errors.New("渠道账户ID无效")
	}
	account, err := m.AccountRepo.FindByID(accountID)
	if err != nil {
		return 0, errors.New("渠道账户不存在")
	}
	if account.MediaManagerManagerID == 0 {
		return 0, errors.New("渠道账户未绑定管家")
	}
	return uint(account.MediaManagerManagerID), nil
}

// GetAccountContext 渠道账户 -> {平台UID, 授权token}
func (m *ManagerAuth) GetAccountContext(accountID uint) (string, string, error) {
	if accountID == 0 {
		return "", "", errors.New("渠道账户ID无效")
	}
	account, err := m.AccountRepo.FindByID(accountID)
	if err != nil {
		return "", "", errors.New("渠道账户不存在")
	}
	token, err := m.GetAccessToken(uint(account.MediaManagerManagerID))
	if err != nil {
		return "", "", err
	}
	return account.UID, token, nil
}

// ParseManagerFromState 解密 state 获取管家ID
func (m *ManagerAuth) ParseManagerFromState(state string) (uint, error) {
	if state == "" {
		return 0, errors.New("授权state不存在")
	}
	data, err := utils.XorDecryptJSON(state, utils.XorKey)
	if err != nil {
		return 0, errors.New("state解析失败")
	}
	id, ok := data["mkt_account_manager_id"].(float64)
	if !ok || id == 0 {
		return 0, errors.New("state中管家ID无效")
	}
	return uint(id), nil
}

// GetManager 获取管家
func (m *ManagerAuth) GetManager(id uint) (*mediaModel.MediaManager, error) {
	manager, err := m.ManagerRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("管家不存在")
	}
	return manager, nil
}

// GetApp 获取管家关联授权应用
func (m *ManagerAuth) GetApp(manager *mediaModel.MediaManager) (*mediaModel.MediaApplication, error) {
	if manager.ApplicationID == 0 {
		return nil, errors.New("管家未绑定授权应用")
	}
	app, err := m.AppRepo.FindByID(uint(manager.ApplicationID))
	if err != nil {
		return nil, errors.New("授权应用不存在")
	}
	if app.AppID == "" {
		return nil, errors.New("授权应用 app_id 未配置")
	}
	return app, nil
}

// RedirectURI 拼接渠道回调地址
func (m *ManagerAuth) RedirectURI(channel string) string {
	path := config.AppConfig.OAuth.RedirectPath
	if path == "" {
		path = "http://127.0.0.1:8080/oauth/callback/%s"
	}
	return fmt.Sprintf(path, channel)
}

// UpdateAuthStatus 更新管家授权状态，auth_message 存入 extra 扩展信息
func (m *ManagerAuth) UpdateAuthStatus(managerID uint, status int8, message string) error {
	manager, err := m.ManagerRepo.FindByID(managerID)
	if err != nil {
		return err
	}
	manager.AuthStatus = status
	if message != "" {
		extra := map[string]interface{}{}
		if manager.Extra != "" {
			_ = json.Unmarshal([]byte(manager.Extra), &extra)
		}
		extra["auth_message"] = message
		raw, _ := json.Marshal(extra)
		manager.Extra = string(raw)
	}
	return m.ManagerRepo.Update(manager)
}

// UpdateAccountNum 更新管家绑定账户数
func (m *ManagerAuth) UpdateAccountNum(manager *mediaModel.MediaManager, num int) error {
	manager.AccountNum = num
	return m.ManagerRepo.Update(manager)
}

// PushSyncQueue 推送管家ID到广告主同步队列（Redis）
func (m *ManagerAuth) PushSyncQueue(managerID uint) {
	if database.Redis == nil {
		return
	}
	payload, _ := json.Marshal(map[string]uint{"mkt_account_manager_id": managerID})
	_ = database.Redis.RPush(context.Background(), constants.RedisKeyMktManagerSyncAdvertiser, payload).Err()
}