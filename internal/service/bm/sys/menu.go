package sys

import (
	"errors"

	"stack-bm/internal/model/bm/sys"
	bmSysRepo "stack-bm/internal/repository/bm/sys"
)

type SysMenuService struct {
	repo *bmSysRepo.SysMenuRepository
}

func NewSysMenuService() *SysMenuService {
	return &SysMenuService{repo: bmSysRepo.NewSysMenuRepository()}
}

func (s *SysMenuService) Create(m *sys.SysMenu) error { return s.repo.Create(m) }
func (s *SysMenuService) FindByID(id uint) (*sys.SysMenu, error) { return s.repo.FindByID(id) }
func (s *SysMenuService) FindPage(page, size int) ([]sys.SysMenu, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size)
}
func (s *SysMenuService) FindAll() ([]sys.SysMenu, error) { return s.repo.FindAll() }

func (s *SysMenuService) Update(id uint, m *sys.SysMenu) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("菜单不存在") }
	if m.Name != "" { existing.Name = m.Name }
	if m.Path != "" { existing.Path = m.Path }
	if m.Icon != "" { existing.Icon = m.Icon }
	if m.Author != "" { existing.Author = m.Author }
	existing.Parent = m.Parent
	existing.Sort = m.Sort
	existing.Type = m.Type
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *SysMenuService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("菜单不存在") }
	return s.repo.Delete(id)
}
