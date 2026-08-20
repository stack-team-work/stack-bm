package user

import (
	"stack-bm/internal/model/user"
	userRepo "stack-bm/internal/repository/user"
)

type UserOrderService struct {
	repo *userRepo.UserOrderRepository
}

func NewUserOrderService() *UserOrderService {
	return &UserOrderService{repo: userRepo.NewUserOrderRepository()}
}

func (s *UserOrderService) FindPage(page, size int, appID int, userID string, payStatus, status int, startAt, endAt int64) ([]user.UserOrder, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, appID, userID, payStatus, status, startAt, endAt)
}