package game

import (
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/internal/model/sdk/game"
)

type GameVoucherUseService struct{ repo *gameRepo.GameVoucherUseRepository }

func NewGameVoucherUseService() *GameVoucherUseService {
	return &GameVoucherUseService{repo: gameRepo.NewGameVoucherUseRepository()}
}

func (s *GameVoucherUseService) FindPage(page, size int, keyword string, voucherID int) ([]game.GameVoucherUse, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, voucherID)
}
