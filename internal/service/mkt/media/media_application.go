package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/dict"
)

type MediaApplicationService struct {
	repo *mediaRepo.MediaApplicationRepository
}

func NewMediaApplicationService() *MediaApplicationService {
	return &MediaApplicationService{repo: mediaRepo.NewMediaApplicationRepository()}
}

func (s *MediaApplicationService) Create(m *media.MediaApplication) error { return s.repo.Create(m) }
func (s *MediaApplicationService) FindByID(id uint) (*media.MediaApplication, error) { return s.repo.FindByID(id) }
func (s *MediaApplicationService) FindPage(page, size int, keyword string, status int, mediaID int) ([]media.MediaApplication, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status, mediaID)
}
func (s *MediaApplicationService) FindAll() ([]media.MediaApplication, error) { return s.repo.FindAll() }

func (s *MediaApplicationService) Update(id uint, m *media.MediaApplication) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("mkt应用不存在") }
	if m.Name != "" { existing.Name = m.Name }
	if m.MediaID > 0 { existing.MediaID = m.MediaID }
	if m.AppID > 0 { existing.AppID = m.AppID }
	if m.AppSecret > 0 { existing.AppSecret = m.AppSecret }
	if m.Remark != "" { existing.Remark = m.Remark }
	if m.Extra != "" { existing.Extra = m.Extra }
	if m.AdminID > 0 { existing.AdminID = m.AdminID }
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *MediaApplicationService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("mkt应用不存在") }
	return s.repo.Delete(id)
}

func (s *MediaApplicationService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
