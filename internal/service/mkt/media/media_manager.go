package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
)

type MediaManagerService struct {
	repo *mediaRepo.MediaManagerRepository
}

func NewMediaManagerService() *MediaManagerService {
	return &MediaManagerService{repo: mediaRepo.NewMediaManagerRepository()}
}

func (s *MediaManagerService) Create(m *media.MediaManager) error { return s.repo.Create(m) }
func (s *MediaManagerService) FindByID(id uint) (*media.MediaManager, error) { return s.repo.FindByID(id) }
func (s *MediaManagerService) FindPage(page, size int, keyword string, status int, mediaID int, applicationID int) ([]media.MediaManager, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status, mediaID, applicationID)
}
func (s *MediaManagerService) FindAll() ([]media.MediaManager, error) { return s.repo.FindAll() }

func (s *MediaManagerService) Update(id uint, m *media.MediaManager) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("mkt管家不存在") }
	if m.Name != "" { existing.Name = m.Name }
	if m.MediaID > 0 { existing.MediaID = m.MediaID }
	if m.ApplicationID > 0 { existing.ApplicationID = m.ApplicationID }
	if m.Account != "" { existing.Account = m.Account }
	if m.AccountID != "" { existing.AccountID = m.AccountID }
	existing.AccountNum = m.AccountNum
	if m.AuthStatus > 0 { existing.AuthStatus = m.AuthStatus }
	if m.Remark != "" { existing.Remark = m.Remark }
	if m.Extra != "" { existing.Extra = m.Extra }
	if m.AdminID > 0 { existing.AdminID = m.AdminID }
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *MediaManagerService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("mkt管家不存在") }
	return s.repo.Delete(id)
}
