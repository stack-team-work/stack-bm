package game

import (
	"errors"
	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
)

type GameGiftCodeService struct{ repo *gameRepo.GameGiftCodeRepository }

func NewGameGiftCodeService() *GameGiftCodeService {
	return &GameGiftCodeService{repo: gameRepo.NewGameGiftCodeRepository()}
}

func (s *GameGiftCodeService) Create(c *game.GameGiftCode) error { return s.repo.Create(c) }
func (s *GameGiftCodeService) FindByID(id uint) (*game.GameGiftCode, error) { return s.repo.FindByID(id) }
func (s *GameGiftCodeService) FindPage(page, size int, keyword string, status int, giftID int) ([]game.GameGiftCode, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status, giftID)
}

func (s *GameGiftCodeService) Update(id uint, c *game.GameGiftCode) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("礼包码不存在")
	}
	if c.Code != "" {
		existing.Code = c.Code
	}
	if c.GiftID > 0 {
		existing.GiftID = c.GiftID
	}
	existing.Status = c.Status
	return s.repo.Update(existing)
}

func (s *GameGiftCodeService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("礼包码不存在")
	}
	return s.repo.Delete(id)
}
