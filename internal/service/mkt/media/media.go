package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
)

type MediaService struct {
	repo *mediaRepo.MediaRepository
}

func NewMediaService() *MediaService {
	return &MediaService{repo: mediaRepo.NewMediaRepository()}
}

func (s *MediaService) Create(m *media.Media) error { return s.repo.Create(m) }
func (s *MediaService) FindByID(id uint) (*media.Media, error) { return s.repo.FindByID(id) }
func (s *MediaService) FindPage(page, size int, keyword string, status int) ([]media.Media, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status)
}
func (s *MediaService) FindAll() ([]media.Media, error) { return s.repo.FindAll() }

func (s *MediaService) Update(id uint, m *media.Media) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("媒体渠道不存在") }
	if m.Name != "" { existing.Name = m.Name }
	if m.Mark != "" { existing.Mark = m.Mark }
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *MediaService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("媒体渠道不存在") }
	return s.repo.Delete(id)
}
