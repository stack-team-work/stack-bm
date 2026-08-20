package user

import (
	"stack-bm/internal/model/user"
	userRepo "stack-bm/internal/repository/user"
)

type UserActiveService struct {
	repo *userRepo.UserActiveRepository
}

func NewUserActiveService() *UserActiveService {
	return &UserActiveService{repo: userRepo.NewUserActiveRepository()}
}

func (s *UserActiveService) FindPage(page, size int, appID int, clientID string, startAt, endAt int64) ([]user.UserActive, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, appID, clientID, startAt, endAt)
}