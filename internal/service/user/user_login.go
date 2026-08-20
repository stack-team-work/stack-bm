package user

import (
	"stack-bm/internal/model/user"
	userRepo "stack-bm/internal/repository/user"
)

type UserLoginService struct {
	repo *userRepo.UserLoginRepository
}

func NewUserLoginService() *UserLoginService {
	return &UserLoginService{repo: userRepo.NewUserLoginRepository()}
}

func (s *UserLoginService) FindPage(page, size int, appID int, userID string, startAt, endAt int64) ([]user.UserLogin, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, appID, userID, startAt, endAt)
}