package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
)

type GamePlatformService struct {
	repo *gameRepo.GamePlatformRepository
}

func NewGamePlatformService() *GamePlatformService {
	return &GamePlatformService{repo: gameRepo.NewGamePlatformRepository()}
}

func (s *GamePlatformService) Create(p *game.GamePlatform) error { return s.repo.Create(p) }

func (s *GamePlatformService) FindByID(id uint) (*game.GamePlatform, error) { return s.repo.FindByID(id) }

func (s *GamePlatformService) FindPage(page, size int, keyword string, status int) ([]game.GamePlatform, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status)
}

func (s *GamePlatformService) FindAll() ([]game.GamePlatform, error) { return s.repo.FindAll() }

func (s *GamePlatformService) Update(id uint, p *game.GamePlatform) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("平台不存在") }
	if p.Name != "" { existing.Name = p.Name }
	if p.Mark != "" { existing.Mark = p.Mark }
	existing.Status = p.Status
	return s.repo.Update(existing)
}

func (s *GamePlatformService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("平台不存在") }
	return s.repo.Delete(id)
}
