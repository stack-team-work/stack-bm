package sys

import (
	"errors"

	"stack-bm/internal/model/bm/sys"
	bmSysRepo "stack-bm/internal/repository/bm/sys"
	"stack-bm/pkg/dict"
)

type SysTagService struct {
	repo *bmSysRepo.SysTagRepository
}

func NewSysTagService() *SysTagService {
	return &SysTagService{repo: bmSysRepo.NewSysTagRepository()}
}

func (s *SysTagService) Create(t *sys.SysTag) error {
	if t.Type <= 0 {
		return errors.New("标签类型不能为空")
	}
	if t.Name == "" {
		return errors.New("标签名称不能为空")
	}
	return s.repo.Create(t)
}

func (s *SysTagService) FindByID(id uint) (*sys.SysTag, error) {
	return s.repo.FindByID(id)
}

func (s *SysTagService) FindPage(page, size int, keyword string, tagType int, status int) ([]sys.SysTag, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, tagType, status)
}

func (s *SysTagService) FindAllByType(tagType int) ([]sys.SysTag, error) {
	return s.repo.FindAllByType(tagType)
}

func (s *SysTagService) FindOptions(tagType int) ([]dict.Option, error) {
	return s.repo.FindOptionsByType(tagType)
}

func (s *SysTagService) Update(id uint, t *sys.SysTag) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("标签不存在")
	}
	if t.Type > 0 {
		existing.Type = t.Type
	}
	if t.Name != "" {
		existing.Name = t.Name
	}
	existing.Remark = t.Remark
	existing.Status = t.Status
	return s.repo.Update(existing)
}
