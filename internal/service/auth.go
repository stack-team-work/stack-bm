package service

import (
	"errors"

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
