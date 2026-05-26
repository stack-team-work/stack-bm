package service

import (
	"errors"
	"fmt"

	"stack-bm/internal/config"
	"stack-bm/internal/model/bm/sys"
	bmSysRepo "stack-bm/internal/repository/bm/sys"
	"stack-bm/pkg/utils"
)

type AuthService struct {
	adminRepo *bmSysRepo.SysAdminRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		adminRepo: bmSysRepo.NewSysAdminRepository(),
	}
}

func (s *AuthService) Login(username, password string) (*sys.SysAdmin, error) {
	admin, err := s.adminRepo.FindByUsername(username)
	if err != nil {
		if config.AppConfig.Server.Mode == "dev" &&
			username == config.AppConfig.Dev.DefaultUsername &&
			password == config.AppConfig.Dev.DefaultPassword {
			return s.createDevDefaultAdmin()
		}
		return nil, errors.New("用户名或密码错误")
	}

	if admin.Status != 1 {
		return nil, errors.New("账号已被禁用")
	}

	if admin.Password != utils.MD5WithSalt(password, admin.Salt) {
		return nil, errors.New("用户名或密码错误")
	}

	return admin, nil
}

func (s *AuthService) createDevDefaultAdmin() (*sys.SysAdmin, error) {
	admin := &sys.SysAdmin{
		Username: config.AppConfig.Dev.DefaultUsername,
		Password: utils.MD5(config.AppConfig.Dev.DefaultPassword),
		Salt:     "",
		Name:     "管理员",
		Status:   1,
	}

	existing, _ := s.adminRepo.FindByUsername(admin.Username)
	if existing != nil {
		return existing, nil
	}

	if err := s.adminRepo.Create(admin); err != nil {
		return nil, fmt.Errorf("创建默认管理员失败: %w", err)
	}

	return admin, nil
}
