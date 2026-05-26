package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/utils"
)

type GameService struct {
	repo *gameRepo.GameRepository
}

func NewGameService() *GameService {
	return &GameService{repo: gameRepo.NewGameRepository()}
}

func (s *GameService) Create(g *game.Game) error {
	if g.Mark == "" {
		g.Mark = utils.ToPinYinMark(g.Name)
	}
	return s.repo.Create(g)
}

func (s *GameService) FindByID(id uint) (*game.Game, error) { return s.repo.FindByID(id) }

func (s *GameService) FindPage(page, size int, keyword string, status int) ([]game.Game, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status)
}

func (s *GameService) FindAll() ([]game.Game, error) { return s.repo.FindAll() }

func (s *GameService) Update(id uint, g *game.Game) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("游戏不存在") }
	if g.Name != "" { existing.Name = g.Name }
	if g.Mark != "" { existing.Mark = g.Mark }
	if g.WebName != "" { existing.WebName = g.WebName }
	if g.Icon != "" { existing.Icon = g.Icon }
	existing.CpID = g.CpID
	existing.Status = g.Status
	return s.repo.Update(existing)
}

func (s *GameService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("游戏不存在") }
	return s.repo.Delete(id)
}
