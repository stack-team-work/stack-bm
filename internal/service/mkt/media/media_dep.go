package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/utils"
)

type MediaDepService struct {
	repo *mediaRepo.MediaDepRepository
}

func NewMediaDepService() *MediaDepService {
	return &MediaDepService{repo: mediaRepo.NewMediaDepRepository()}
}

func (s *MediaDepService) Create(m *media.MediaDep) error {
	if m.Mark == "" {
		m.Mark = utils.ToPinYinMark(m.Name)
	}
	return s.repo.Create(m)
}

func (s *MediaDepService) FindByID(id uint) (*media.MediaDep, error) { return s.repo.FindByID(id) }

func (s *MediaDepService) FindPage(page, size int, keyword string, status int) ([]media.MediaDep, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status)
}

func (s *MediaDepService) FindAll() ([]media.MediaDep, error) { return s.repo.FindAll() }

func (s *MediaDepService) Update(id uint, m *media.MediaDep) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("账户部门不存在")
	}
	if m.Name != "" {
		existing.Name = m.Name
	}
	if m.Mark != "" {
		existing.Mark = m.Mark
	}
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *MediaDepService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("账户部门不存在")
	}
	return s.repo.Delete(id)
}

func (s *MediaDepService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
