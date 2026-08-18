package feishu

import (
	"errors"
	"stack-bm/internal/model/bm/feishu"
	feishuRepo "stack-bm/internal/repository/bm/feishu"
)

type FeishuUserService struct{ repo *feishuRepo.FeishuUserRepository }

func NewFeishuUserService() *FeishuUserService {
	return &FeishuUserService{repo: feishuRepo.NewFeishuUserRepository()}
}

func (s *FeishuUserService) Create(m *feishu.FeishuUser) error { return s.repo.Create(m) }
func (s *FeishuUserService) FindByID(id uint) (*feishu.FeishuUser, error) {
	return s.repo.FindByID(id)
}
func (s *FeishuUserService) FindPage(page, size int, keyword string, status int, adminID int) ([]feishu.FeishuUser, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status, adminID)
}
func (s *FeishuUserService) FindAll() ([]feishu.FeishuUser, error) { return s.repo.FindAll() }

func (s *FeishuUserService) Update(id uint, m *feishu.FeishuUser) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书用户不存在")
	}
	if m.FeishuUserID != "" {
		existing.FeishuUserID = m.FeishuUserID
	}
	if m.AdminID > 0 {
		existing.AdminID = m.AdminID
	}
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *FeishuUserService) UpdateStatus(id uint, status int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书用户不存在")
	}
	existing.Status = status
	return s.repo.Update(existing)
}

func (s *FeishuUserService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书用户不存在")
	}
	return s.repo.Delete(id)
}
