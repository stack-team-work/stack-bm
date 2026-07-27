package sys

import (
	"errors"
	"stack-bm/internal/model/bm/sys"
	bmSysRepo "stack-bm/internal/repository/bm/sys"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/utils"
)

type FeishuAppService struct{ repo *bmSysRepo.FeishuAppRepository }

func NewFeishuAppService() *FeishuAppService {
	return &FeishuAppService{repo: bmSysRepo.NewFeishuAppRepository()}
}

func (s *FeishuAppService) Create(m *sys.FeishuApp) error {
	if m.Mark == "" {
		m.Mark = utils.ToPinYinMark(m.AppName)
	}
	return s.repo.Create(m)
}
func (s *FeishuAppService) FindByID(id uint) (*sys.FeishuApp, error) { return s.repo.FindByID(id) }
func (s *FeishuAppService) FindOptions() ([]dict.Option, error)      { return s.repo.FindOptions() }
func (s *FeishuAppService) FindPage(page, size int, keyword string, status int) ([]sys.FeishuApp, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status)
}
func (s *FeishuAppService) FindAll() ([]sys.FeishuApp, error) { return s.repo.FindAll() }

func (s *FeishuAppService) Update(id uint, m *sys.FeishuApp) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书应用不存在")
	}
	if m.AppName != "" {
		existing.AppName = m.AppName
	}
	if m.AppID != "" {
		existing.AppID = m.AppID
	}
	if m.AppSecret != "" {
		existing.AppSecret = m.AppSecret
	}
	if m.Mark != "" {
		existing.Mark = m.Mark
	}
	if m.AdminID > 0 {
		existing.AdminID = m.AdminID
	}
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *FeishuAppService) UpdateStatus(id uint, status int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书应用不存在")
	}
	existing.Status = status
	return s.repo.Update(existing)
}

func (s *FeishuAppService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书应用不存在")
	}
	return s.repo.Delete(id)
}
