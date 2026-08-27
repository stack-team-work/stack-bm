package token

import (
	"errors"

	mediaRepo "stack-bm/internal/repository/mkt/media"
)

// TokenService 腾讯渠道 token 服务：解析本渠道 access_token，
// 不再依赖共享 oauth.ManagerAuth
type TokenService struct {
	accountRepo *mediaRepo.MediaAccountRepository
	tokenRepo   *mediaRepo.MediaManagerTokenRepository
}

func NewTokenService() *TokenService {
	return &TokenService{
		accountRepo: mediaRepo.NewMediaAccountRepository(),
		tokenRepo:   mediaRepo.NewMediaManagerTokenRepository(),
	}
}

// AccessToken 渠道账户记录ID(media_accounts.id) -> access_token（余额/层级同步用）
func (s *TokenService) AccessToken(accountID uint) (string, error) {
	if accountID == 0 {
		return "", errors.New("渠道账户ID无效")
	}
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return "", errors.New("渠道账户不存在")
	}
	return s.accessTokenOfManager(account.MediaManagerManagerID)
}

// AccessTokenByUID 平台账户UID -> 定位渠道账户行 -> access_token
// （工具操作经 cpid 反查 account_id 后使用）
func (s *TokenService) AccessTokenByUID(uid string) (string, error) {
	if uid == "" {
		return "", errors.New("平台账户UID为空")
	}
	account, err := s.accountRepo.FindByUID(uid)
	if err != nil {
		return "", errors.New("渠道账户不存在")
	}
	return s.accessTokenOfManager(account.MediaManagerManagerID)
}

// Context 渠道账户记录ID -> 平台UID字符串 + access_token（同步链路用）
func (s *TokenService) Context(accountID uint) (string, string, error) {
	if accountID == 0 {
		return "", "", errors.New("渠道账户ID无效")
	}
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return "", "", errors.New("渠道账户不存在")
	}
	token, err := s.accessTokenOfManager(account.MediaManagerManagerID)
	if err != nil {
		return "", "", err
	}
	return account.UID, token, nil
}

func (s *TokenService) accessTokenOfManager(managerID int) (string, error) {
	if managerID == 0 {
		return "", errors.New("渠道账户未绑定管家")
	}
	authInfo, err := s.tokenRepo.FindByManagerID(int(managerID))
	if err != nil {
		return "", err
	}
	if authInfo == nil {
		return "", errors.New("管家授权信息不存在")
	}
	accessToken, _ := authInfo["access_token"].(string)
	if accessToken == "" {
		return "", errors.New("管家未授权或token为空")
	}
	return accessToken, nil
}
