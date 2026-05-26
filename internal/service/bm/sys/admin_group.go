package sys

import (
	"errors"

	"stack-bm/internal/model/bm/sys"
	bmSysRepo "stack-bm/internal/repository/bm/sys"
)

type SysAdminGroupService struct {
	repo *bmSysRepo.SysAdminGroupRepository
}

func NewSysAdminGroupService() *SysAdminGroupService {
	return &SysAdminGroupService{repo: bmSysRepo.NewSysAdminGroupRepository()}
}

func (s *SysAdminGroupService) Create(group *sys.SysAdminGroup) error { return s.repo.Create(group) }

func (s *SysAdminGroupService) FindByID(id uint) (*sys.SysAdminGroup, error) { return s.repo.FindByID(id) }

func (s *SysAdminGroupService) FindPage(page, size int, keyword string) ([]sys.SysAdminGroup, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword)
}

func (s *SysAdminGroupService) FindAll() ([]sys.SysAdminGroup, error) { return s.repo.FindAll() }

func (s *SysAdminGroupService) Update(id uint, group *sys.SysAdminGroup) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("分组不存在") }
	if group.Name != "" { existing.Name = group.Name }
	if group.Mark != "" { existing.Mark = group.Mark }
	if group.Description != "" { existing.Description = group.Description }
	if group.MenuPermit != "" { existing.MenuPermit = group.MenuPermit }
	if group.ColumnPermit != "" { existing.ColumnPermit = group.ColumnPermit }
	existing.Status = group.Status
	return s.repo.Update(existing)
}

func (s *SysAdminGroupService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("分组不存在") }
	return s.repo.Delete(id)
}
