package sys

import (
	"errors"

	"stack-bm/internal/model/sys"
	sysRepo "stack-bm/internal/repository/sys"
	"stack-bm/pkg/utils"
)

type SysAdminService struct {
	repo *sysRepo.SysAdminRepository
}

func NewSysAdminService() *SysAdminService {
	return &SysAdminService{
		repo: sysRepo.NewSysAdminRepository(),
	}
}

func (s *SysAdminService) Create(admin *sys.SysAdmin) error {
	existing, _ := s.repo.FindByUsername(admin.Username)
	if existing != nil {
		return errors.New("用户名已存在")
	}

	salt := utils.GenerateSalt()
	admin.Password = utils.MD5WithSalt(admin.Password, salt)
	admin.Salt = salt

	return s.repo.Create(admin)
}

func (s *SysAdminService) FindByID(id uint) (*sys.SysAdmin, error) {
	return s.repo.FindByID(id)
}

func (s *SysAdminService) FindPage(page, size int, keyword string, groupID int) ([]sys.SysAdmin, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, groupID)
}

func (s *SysAdminService) Update(id uint, admin *sys.SysAdmin) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("管理员不存在")
	}

	if admin.Username != "" && admin.Username != existing.Username {
		dup, _ := s.repo.FindByUsername(admin.Username)
		if dup != nil {
			return errors.New("用户名已存在")
		}
		existing.Username = admin.Username
	}

	if admin.Password != "" {
		salt := utils.GenerateSalt()
		existing.Password = utils.MD5WithSalt(admin.Password, salt)
		existing.Salt = salt
	}

	if admin.Name != "" {
		existing.Name = admin.Name
	}
	if admin.Phone != "" {
		existing.Phone = admin.Phone
	}
	if admin.GroupID > 0 {
		existing.GroupID = admin.GroupID
	}
	existing.Status = admin.Status

	return s.repo.Update(existing)
}

func (s *SysAdminService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("管理员不存在")
	}
	return s.repo.Delete(id)
}

type SysAdminGroupService struct {
	repo *sysRepo.SysAdminGroupRepository
}

func NewSysAdminGroupService() *SysAdminGroupService {
	return &SysAdminGroupService{
		repo: sysRepo.NewSysAdminGroupRepository(),
	}
}

func (s *SysAdminGroupService) Create(group *sys.SysAdminGroup) error {
	return s.repo.Create(group)
}

func (s *SysAdminGroupService) FindByID(id uint) (*sys.SysAdminGroup, error) {
	return s.repo.FindByID(id)
}

func (s *SysAdminGroupService) FindPage(page, size int, keyword string) ([]sys.SysAdminGroup, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword)
}

func (s *SysAdminGroupService) FindAll() ([]sys.SysAdminGroup, error) {
	return s.repo.FindAll()
}

func (s *SysAdminGroupService) Update(id uint, group *sys.SysAdminGroup) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("分组不存在")
	}

	if group.Name != "" {
		existing.Name = group.Name
	}
	if group.Mark != "" {
		existing.Mark = group.Mark
	}
	if group.Description != "" {
		existing.Description = group.Description
	}
	if group.MenuPermit != "" {
		existing.MenuPermit = group.MenuPermit
	}
	if group.ColumnPermit != "" {
		existing.ColumnPermit = group.ColumnPermit
	}
	existing.Status = group.Status

	return s.repo.Update(existing)
}

func (s *SysAdminGroupService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("分组不存在")
	}
	return s.repo.Delete(id)
}

type SysLogService struct {
	repo *sysRepo.SysLogRepository
}

func NewSysLogService() *SysLogService {
	return &SysLogService{repo: sysRepo.NewSysLogRepository()}
}

func (s *SysLogService) Create(log *sys.SysLog) error {
	return s.repo.Create(log)
}

func (s *SysLogService) FindPage(page, size int, keyword string, level string) ([]sys.SysLog, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, level)
}

func (s *SysLogService) ClearAll() error {
	return s.repo.ClearAll()
}

type SysMenuService struct {
	repo *sysRepo.SysMenuRepository
}

func NewSysMenuService() *SysMenuService {
	return &SysMenuService{repo: sysRepo.NewSysMenuRepository()}
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
