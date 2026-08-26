package media

import (
	"errors"

	"stack-bm/internal/model/mkt/media"
	mediaRepo "stack-bm/internal/repository/mkt/media"
	"stack-bm/pkg/dict"
)

type MediaAccountService struct {
	repo *mediaRepo.MediaAccountRepository
}

func NewMediaAccountService() *MediaAccountService {
	return &MediaAccountService{repo: mediaRepo.NewMediaAccountRepository()}
}

func (s *MediaAccountService) Create(m *media.MediaAccount) error { return s.repo.Create(m) }

func (s *MediaAccountService) FindByID(id uint) (*media.MediaAccount, error) {
	return s.repo.FindByID(id)
}

func (s *MediaAccountService) FindPage(page, size int, keyword string, status int, mediaSubID int, subjectID int, adminID int) ([]media.MediaAccount, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status, mediaSubID, subjectID, adminID)
}

func (s *MediaAccountService) FindAll() ([]media.MediaAccount, error) {
	return s.repo.FindAll()
}

// ResolveManagerByAccount 渠道账户 -> 管家ID（1:1，通过 media_manager_manager_id）
func (s *MediaAccountService) ResolveManagerByAccount(accountID uint) (uint, error) {
	account, err := s.repo.FindByID(accountID)
	if err != nil {
		return 0, errors.New("渠道账户不存在")
	}
	if account.MediaManagerManagerID == 0 {
		return 0, errors.New("渠道账户未绑定管家")
	}
	return uint(account.MediaManagerManagerID), nil
}

// FindByUID 通过平台UID查找渠道账户
func (s *MediaAccountService) FindByUID(uid string) (*media.MediaAccount, error) {
	return s.repo.FindByUID(uid)
}

// UpdateBalance 回填账户余额
func (s *MediaAccountService) UpdateBalance(id uint, balance float64) error {
	return s.repo.UpdateBalance(id, balance)
}

func (s *MediaAccountService) Update(id uint, m *media.MediaAccount) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("渠道账户不存在")
	}
	if m.Name != "" {
		existing.Name = m.Name
	}
	if m.Username != "" {
		existing.Username = m.Username
	}
	if m.UID != "" {
		existing.UID = m.UID
	}
	if m.AgentID > 0 {
		existing.AgentID = m.AgentID
	}
	if m.MediaSubID > 0 {
		existing.MediaSubID = m.MediaSubID
	}
	if m.SubjectID > 0 {
		existing.SubjectID = m.SubjectID
	}
	if m.AdminID > 0 {
		existing.AdminID = m.AdminID
	}
	existing.Rebate = m.Rebate
	existing.Balance = m.Balance
	if m.UseType > 0 {
		existing.UseType = m.UseType
	}
	existing.MediaManagerManagerID = m.MediaManagerManagerID
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *MediaAccountService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("渠道账户不存在")
	}
	return s.repo.Delete(id)
}

func (s *MediaAccountService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }