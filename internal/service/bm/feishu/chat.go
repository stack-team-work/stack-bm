package feishu

import (
	"errors"
	"stack-bm/internal/model/bm/feishu"
	feishuRepo "stack-bm/internal/repository/bm/feishu"
)

type FeishuChatService struct{ repo *feishuRepo.FeishuChatRepository }

func NewFeishuChatService() *FeishuChatService {
	return &FeishuChatService{repo: feishuRepo.NewFeishuChatRepository()}
}

func (s *FeishuChatService) Create(m *feishu.FeishuChat) error { return s.repo.Create(m) }
func (s *FeishuChatService) FindByID(id uint) (*feishu.FeishuChat, error) {
	return s.repo.FindByID(id)
}
func (s *FeishuChatService) FindPage(page, size int, keyword string, status int, feishuAppID int) ([]feishu.FeishuChat, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status, feishuAppID)
}
func (s *FeishuChatService) FindAll() ([]feishu.FeishuChat, error) { return s.repo.FindAll() }

func (s *FeishuChatService) Update(id uint, m *feishu.FeishuChat) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书聊天不存在")
	}
	if m.ChatID != "" {
		existing.ChatID = m.ChatID
	}
	if m.DefaultAtList != "" {
		existing.DefaultAtList = m.DefaultAtList
	}
	if m.AtList != "" {
		existing.AtList = m.AtList
	}
	if m.CallAction != "" {
		existing.CallAction = m.CallAction
	}
	if m.ActionTitle != "" {
		existing.ActionTitle = m.ActionTitle
	}
	existing.Type = m.Type
	existing.AtType = m.AtType
	if m.FeishuAppID > 0 {
		existing.FeishuAppID = m.FeishuAppID
	}
	if m.AdminID > 0 {
		existing.AdminID = m.AdminID
	}
	existing.Status = m.Status
	return s.repo.Update(existing)
}

func (s *FeishuChatService) UpdateStatus(id uint, status int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书聊天不存在")
	}
	existing.Status = status
	return s.repo.Update(existing)
}

func (s *FeishuChatService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("飞书聊天不存在")
	}
	return s.repo.Delete(id)
}
