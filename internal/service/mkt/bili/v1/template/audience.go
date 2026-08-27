package template

import (
	"errors"

	biliModel "stack-bm/internal/model/mkt/bili"
	biliRepo "stack-bm/internal/repository/mkt/bili"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AudienceService B站定向包模板
type AudienceService struct {
	repo *biliRepo.AudienceTemplateRepository
}

func NewAudienceService() *AudienceService {
	return &AudienceService{repo: biliRepo.NewAudienceTemplateRepository()}
}

func (s *AudienceService) Create(doc *biliModel.AudienceTemplate, adminID int) error {
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	doc.SysUserID = adminID
	doc.Display = biliModel.DisplayShow
	doc.CreatedAt = now()
	doc.UpdatedAt = now()
	return s.repo.Create(doc)
}

func (s *AudienceService) FindPage(page, size int, keyword string) ([]biliModel.AudienceTemplate, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword)
}

func (s *AudienceService) FindByID(id string) (*biliModel.AudienceTemplate, error) {
	return s.repo.FindByID(id)
}

func (s *AudienceService) Update(id string, doc *biliModel.AudienceTemplate) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("定向包模板不存在")
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

func (s *AudienceService) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("定向包模板不存在")
	}
	return s.repo.SoftDelete(id)
}

func (s *AudienceService) Copy(id, newName string, adminID int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("定向包模板不存在")
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
