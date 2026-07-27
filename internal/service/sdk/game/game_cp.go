package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/utils"
)

type GameCpService struct {
	repo *gameRepo.GameCpRepository
}

func NewGameCpService() *GameCpService {
	return &GameCpService{repo: gameRepo.NewGameCpRepository()}
}

func (s *GameCpService) Create(cp *game.GameCp) error {
	if cp.Mark == "" {
		cp.Mark = utils.ToPinYinMark(cp.Name)
	}
	return s.repo.Create(cp)
}

func (s *GameCpService) FindByID(id uint) (*game.GameCp, error) { return s.repo.FindByID(id) }

func (s *GameCpService) FindPage(page, size int, keyword string, status int) ([]game.GameCp, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status)
}

func (s *GameCpService) FindAll() ([]game.GameCp, error) { return s.repo.FindAll() }

func (s *GameCpService) Update(id uint, cp *game.GameCp) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("CP不存在") }
	if cp.Name != "" { existing.Name = cp.Name }
	if cp.Mark != "" { existing.Mark = cp.Mark }
	if cp.Phone != "" { existing.Phone = cp.Phone }
	if cp.Addr != "" { existing.Addr = cp.Addr }
	existing.Status = cp.Status
	return s.repo.Update(existing)
}

func (s *GameCpService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("CP不存在") }
	return s.repo.Delete(id)
}

func (s *GameCpService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
