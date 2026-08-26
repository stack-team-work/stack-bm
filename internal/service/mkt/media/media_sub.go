package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/utils"
)

type MediaSubService struct {
	repo *mediaRepo.MediaSubRepository
}

func NewMediaSubService() *MediaSubService {
	return &MediaSubService{repo: mediaRepo.NewMediaSubRepository()}
}

func (s *MediaSubService) Create(sub *media.MediaSub) error {
	if sub.Mark == "" {
		sub.Mark = utils.ToPinYinMark(sub.Name)
	}
	return s.repo.Create(sub)
}
func (s *MediaSubService) FindByID(id uint) (*media.MediaSub, error) { return s.repo.FindByID(id) }
func (s *MediaSubService) FindPage(page, size int, keyword string, mediaID int, status int) ([]media.MediaSub, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, mediaID, status)
}
func (s *MediaSubService) FindAll() ([]media.MediaSub, error) { return s.repo.FindAll() }

func (s *MediaSubService) Update(id uint, sub *media.MediaSub) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("媒体子渠道不存在") }
	if sub.Name != "" { existing.Name = sub.Name }
	if sub.Mark != "" { existing.Mark = sub.Mark }
	if sub.MediaID > 0 { existing.MediaID = sub.MediaID }
	existing.Status = sub.Status
	return s.repo.Update(existing)
}

func (s *MediaSubService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("媒体子渠道不存在") }
	return s.repo.Delete(id)
}

func (s *MediaSubService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
