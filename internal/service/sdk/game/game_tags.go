package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/utils"
)

type GameTagService struct {
	repo *gameRepo.GameTagRepository
}

func NewGameTagService() *GameTagService {
	return &GameTagService{repo: gameRepo.NewGameTagRepository()}
}

func (s *GameTagService) Create(tag *game.GameTag) error {
	if tag.Mark == "" {
		tag.Mark = utils.ToPinYinMark(tag.Name)
	}
	return s.repo.Create(tag)
}

func (s *GameTagService) FindByID(id uint) (*game.GameTag, error) { return s.repo.FindByID(id) }

func (s *GameTagService) FindPage(page, size int, keyword string, tagType int, status int) ([]game.GameTag, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, tagType, status)
}

func (s *GameTagService) FindAllByType(tagType int) ([]game.GameTag, error) {
	return s.repo.FindAllByType(tagType)
}

func (s *GameTagService) Update(id uint, tag *game.GameTag) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("标签不存在") }
	if tag.Name != "" { existing.Name = tag.Name }
	if tag.Mark != "" { existing.Mark = tag.Mark }
	existing.Status = tag.Status
	return s.repo.Update(existing)
}

func (s *GameTagService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("标签不存在") }
	return s.repo.Delete(id)
}
