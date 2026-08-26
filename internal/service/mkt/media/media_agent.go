package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/utils"
)

type MediaAgentService struct {
	repo *mediaRepo.MediaAgentRepository
}

func NewMediaAgentService() *MediaAgentService {
	return &MediaAgentService{repo: mediaRepo.NewMediaAgentRepository()}
}

func (s *MediaAgentService) Create(m *media.MediaAgent) error {
	if m.Mark == "" { m.Mark = utils.ToPinYinMark(m.Name) }
	return s.repo.Create(m)
}

func (s *MediaAgentService) FindByID(id uint) (*media.MediaAgent, error) { return s.repo.FindByID(id) }

func (s *MediaAgentService) FindPage(page, size int, keyword string, status int, subjectID int) ([]media.MediaAgent, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status, subjectID)
}

func (s *MediaAgentService) FindAll() ([]media.MediaAgent, error) { return s.repo.FindAll() }

func (s *MediaAgentService) Update(id uint, m *media.MediaAgent) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("渠道代理不存在") }
	if m.Name != "" { existing.Name = m.Name }
	if m.Mark != "" { existing.Mark = m.Mark }
	if m.SubjectID > 0 { existing.SubjectID = m.SubjectID }
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *MediaAgentService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("渠道代理不存在") }
	return s.repo.Delete(id)
}

func (s *MediaAgentService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
