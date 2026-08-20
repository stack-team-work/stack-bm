package user

import (
	"stack-bm/internal/model/user"
	userRepo "stack-bm/internal/repository/user"
)

type UserInfoService struct {
	repo *userRepo.UserInfoRepository
}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{repo: userRepo.NewUserInfoRepository()}
}

func (s *UserInfoService) FindPage(page, size int, appID int, userID string, startAt, endAt int64) ([]user.UserInfo, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, appID, userID, startAt, endAt)
}