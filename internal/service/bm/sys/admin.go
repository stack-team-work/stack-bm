package sys

import (
	"errors"

	"stack-bm/internal/model/bm/sys"
	bmSysRepo "stack-bm/internal/repository/bm/sys"
	"stack-bm/pkg/utils"
)

type SysAdminService struct {
	repo *bmSysRepo.SysAdminRepository
}

func NewSysAdminService() *SysAdminService {
	return &SysAdminService{repo: bmSysRepo.NewSysAdminRepository()}
}

func (s *SysAdminService) Create(admin *sys.SysAdmin) error {
	existing, _ := s.repo.FindByUsername(admin.Username)
	if existing != nil { return errors.New("用户名已存在") }
	salt := utils.GenerateSalt()
	admin.Password = utils.MD5WithSalt(admin.Password, salt)
	admin.Salt = salt
	return s.repo.Create(admin)
}

func (s *SysAdminService) FindByID(id uint) (*sys.SysAdmin, error) { return s.repo.FindByID(id) }

func (s *SysAdminService) FindPage(page, size int, keyword string, groupID int) ([]sys.SysAdmin, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, groupID)
}

func (s *SysAdminService) Update(id uint, admin *sys.SysAdmin) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("管理员不存在") }
	if admin.Username != "" && admin.Username != existing.Username {
		dup, _ := s.repo.FindByUsername(admin.Username)
		if dup != nil { return errors.New("用户名已存在") }
		existing.Username = admin.Username
	}
	if admin.Password != "" {
		salt := utils.GenerateSalt()
		existing.Password = utils.MD5WithSalt(admin.Password, salt)
		existing.Salt = salt
	}
	if admin.Name != "" { existing.Name = admin.Name }
	if admin.Phone != "" { existing.Phone = admin.Phone }
	if admin.GroupID > 0 { existing.GroupID = admin.GroupID }
	existing.Status = admin.Status
	return s.repo.Update(existing)
}

func (s *SysAdminService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("管理员不存在") }
	return s.repo.Delete(id)
}
