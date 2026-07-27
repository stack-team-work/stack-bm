package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/utils"
)

type MediaSubjectService struct {
	repo *mediaRepo.MediaSubjectRepository
}

func NewMediaSubjectService() *MediaSubjectService {
	return &MediaSubjectService{repo: mediaRepo.NewMediaSubjectRepository()}
}

func (s *MediaSubjectService) Create(m *media.MediaSubject) error {
	if m.Mark == "" { m.Mark = utils.ToPinYinMark(m.Name) }
	return s.repo.Create(m)
}

func (s *MediaSubjectService) FindByID(id uint) (*media.MediaSubject, error) { return s.repo.FindByID(id) }
func (s *MediaSubjectService) FindPage(page, size int, keyword string, status int) ([]media.MediaSubject, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status)
}
func (s *MediaSubjectService) FindAll() ([]media.MediaSubject, error) { return s.repo.FindAll() }

func (s *MediaSubjectService) Update(id uint, m *media.MediaSubject) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("主体不存在") }
	if m.Name != "" { existing.Name = m.Name }
	if m.Mark != "" { existing.Mark = m.Mark }
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *MediaSubjectService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("主体不存在") }
	return s.repo.Delete(id)
}

func (s *MediaSubjectService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
