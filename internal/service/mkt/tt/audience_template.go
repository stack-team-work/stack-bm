package tt

import (
	"errors"

	ttModel "stack-bm/internal/model/mkt/tt"
	ttRepo "stack-bm/internal/repository/mkt/tt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AudienceTemplateService struct {
	repo *ttRepo.AudienceTemplateRepository
}

func NewAudienceTemplateService() *AudienceTemplateService {
	return &AudienceTemplateService{repo: ttRepo.NewAudienceTemplateRepository()}
}

func (s *AudienceTemplateService) Create(doc *ttModel.AudienceTemplate, adminID int) error {
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	doc.SysUserID = adminID
	doc.Display = ttModel.DisplayShow
	doc.CreatedAt = now()
	doc.UpdatedAt = now()
	return s.repo.Create(doc)
}

func (s *AudienceTemplateService) FindPage(page, size int, keyword string) ([]ttModel.AudienceTemplate, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword)
}

func (s *AudienceTemplateService) FindByID(id string) (*ttModel.AudienceTemplate, error) {
	return s.repo.FindByID(id)
}

func (s *AudienceTemplateService) Update(id string, doc *ttModel.AudienceTemplate) error {
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

func (s *AudienceTemplateService) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("定向包模板不存在")
	}
	return s.repo.SoftDelete(id)
}

func (s *AudienceTemplateService) Copy(id, newName string, adminID int) error {
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
	existing.Display = ttModel.DisplayShow
	existing.CreatedAt = now()
	existing.UpdatedAt = now()
	return s.repo.Create(existing)
}
