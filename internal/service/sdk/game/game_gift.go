package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/dict"
)

type GameGiftService struct{ repo *gameRepo.GameGiftRepository }

func NewGameGiftService() *GameGiftService { return &GameGiftService{repo: gameRepo.NewGameGiftRepository()} }

func (s *GameGiftService) Create(g *game.GameGift) error { return s.repo.Create(g) }
func (s *GameGiftService) FindByID(id uint) (*game.GameGift, error) { return s.repo.FindByID(id) }
func (s *GameGiftService) FindPage(page, size int, keyword string, status int) ([]game.GameGift, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status)
}
func (s *GameGiftService) FindAll() ([]game.GameGift, error) { return s.repo.FindAll() }

func (s *GameGiftService) Update(id uint, g *game.GameGift) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("礼包不存在")
	}
	if g.Name != "" {
		existing.Name = g.Name
	}
	existing.GetType = g.GetType
	existing.IsCode = g.IsCode
	existing.Type = g.Type
	if g.Cond != "" {
		existing.Cond = g.Cond
	}
	if g.Desc != "" {
		existing.Desc = g.Desc
	}
	if g.Stime > 0 {
		existing.Stime = g.Stime
	}
	if g.Etime > 0 {
		existing.Etime = g.Etime
	}
	existing.Status = g.Status
	return s.repo.Update(existing)
}

func (s *GameGiftService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("礼包不存在")
	}
	return s.repo.Delete(id)
}

func (s *GameGiftService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
