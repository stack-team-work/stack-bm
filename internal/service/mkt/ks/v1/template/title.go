package template

import (
	"errors"

	ksModel "stack-bm/internal/model/mkt/ks"
	ksRepo "stack-bm/internal/repository/mkt/ks"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// TitleService B站标题包模板
type TitleService struct {
	repo *ksRepo.TitleTemplateRepository
}

func NewTitleService() *TitleService {
	return &TitleService{repo: ksRepo.NewTitleTemplateRepository()}
}

func (s *TitleService) Create(doc *ksModel.TitleTemplate, adminID int) error {
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	if len(doc.TitleMaterials) > 0 {
		doc.TitleNum = len(doc.TitleMaterials)
	}
	doc.SysUserID = adminID
	doc.Display = ksModel.DisplayShow
	doc.CreatedAt = now()
	doc.UpdatedAt = now()
	return s.repo.Create(doc)
}

func (s *TitleService) FindPage(page, size int, keyword string) ([]ksModel.TitleTemplate, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword)
}

func (s *TitleService) FindByID(id string) (*ksModel.TitleTemplate, error) {
	return s.repo.FindByID(id)
}

func (s *TitleService) Update(id string, doc *ksModel.TitleTemplate) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("标题包模板不存在")
	}
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	doc.ID = existing.ID
	doc.SysUserID = existing.SysUserID
	doc.Display = existing.Display
	doc.CreatedAt = existing.CreatedAt
	doc.UpdatedAt = now()
	if len(doc.TitleMaterials) > 0 {
		doc.TitleNum = len(doc.TitleMaterials)
	}
	return s.repo.Update(doc)
}

func (s *TitleService) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("标题包模板不存在")
	}
	return s.repo.SoftDelete(id)
}

func (s *TitleService) Copy(id, newName string, adminID int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("标题包模板不存在")
	}
	if newName == "" {
		return errors.New("模板名称不能为空")
	}
	existing.ID = bson.ObjectID{}
	existing.TemplateName = newName
	existing.SysUserID = adminID
	existing.Display = ksModel.DisplayShow
	existing.CreatedAt = now()
	existing.UpdatedAt = now()
	return s.repo.Create(existing)
}
