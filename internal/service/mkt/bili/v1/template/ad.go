package template

import (
	"errors"
	"time"

	biliModel "stack-bm/internal/model/mkt/bili"
	biliRepo "stack-bm/internal/repository/mkt/bili"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AdService B站广告模板
type AdService struct {
	repo *biliRepo.AdTemplateRepository
}

func NewAdService() *AdService {
	return &AdService{repo: biliRepo.NewAdTemplateRepository()}
}

func now() string { return time.Now().Format("2006-01-02 15:04:05") }

func (s *AdService) Create(doc *biliModel.AdTemplate, adminID int) error {
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	doc.SysUserID = adminID
	doc.Display = biliModel.DisplayShow
	doc.CreatedAt = now()
	doc.UpdatedAt = now()
	return s.repo.Create(doc)
}

func (s *AdService) FindPage(page, size int, keyword string) ([]biliModel.AdTemplate, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword)
}

func (s *AdService) FindByID(id string) (*biliModel.AdTemplate, error) {
	return s.repo.FindByID(id)
}

func (s *AdService) Update(id string, doc *biliModel.AdTemplate) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("广告模板不存在")
	}
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	doc.ID = existing.ID
	doc.SysUserID = existing.SysUserID
	doc.Display = existing.Display
	doc.CreatedAt = existing.CreatedAt
	doc.UpdatedAt = now()
	return s.repo.Update(doc)
}

func (s *AdService) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("广告模板不存在")
	}
	return s.repo.SoftDelete(id)
}

func (s *AdService) Copy(id, newName string, adminID int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("广告模板不存在")
	}
	if newName == "" {
		return errors.New("模板名称不能为空")
	}
	existing.ID = bson.ObjectID{}
	existing.TemplateName = newName
	existing.SysUserID = adminID
	existing.Display = biliModel.DisplayShow
	existing.CreatedAt = now()
	existing.UpdatedAt = now()
	return s.repo.Create(existing)
}
