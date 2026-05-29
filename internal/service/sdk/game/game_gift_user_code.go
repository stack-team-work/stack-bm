package game

import (
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/internal/model/sdk/game"
)

type GameGiftUserCodeService struct{ repo *gameRepo.GameGiftUserCodeRepository }

func NewGameGiftUserCodeService() *GameGiftUserCodeService {
	return &GameGiftUserCodeService{repo: gameRepo.NewGameGiftUserCodeRepository()}
}

func (s *GameGiftUserCodeService) FindPage(page, size int, keyword string, giftID int) ([]game.GameGiftUserCode, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, giftID)
}
