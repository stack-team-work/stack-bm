package ks

import (
	"errors"
	"time"

	ksModel "stack-bm/internal/model/mkt/ks"
	ksRepo "stack-bm/internal/repository/mkt/ks"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func now() string { return time.Now().Format("2006-01-02 15:04:05") }

type AdTemplateService struct {
	repo *ksRepo.AdTemplateRepository
}

func NewAdTemplateService() *AdTemplateService {
	return &AdTemplateService{repo: ksRepo.NewAdTemplateRepository()}
}

func (s *AdTemplateService) Create(doc *ksModel.AdTemplate, adminID int) error {
	if doc.TemplateName == "" {
		return errors.New("模板名称不能为空")
	}
	doc.SysUserID = adminID
	doc.Display = ksModel.DisplayShow
	doc.CreatedAt = now()
	doc.UpdatedAt = now()
	return s.repo.Create(doc)
}

func (s *AdTemplateService) FindPage(page, size int, keyword string) ([]ksModel.AdTemplate, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword)
}

func (s *AdTemplateService) FindByID(id string) (*ksModel.AdTemplate, error) {
	return s.repo.FindByID(id)
}

func (s *AdTemplateService) Update(id string, doc *ksModel.AdTemplate) error {
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

func (s *AdTemplateService) Delete(id string) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return errors.New("广告模板不存在")
	}
	return s.repo.SoftDelete(id)
}

func (s *AdTemplateService) Copy(id, newName string, adminID int) error {
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
	existing.Display = ksModel.DisplayShow
	existing.CreatedAt = now()
	existing.UpdatedAt = now()
	return s.repo.Create(existing)
}
