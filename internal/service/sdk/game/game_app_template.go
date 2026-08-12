package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/dict"
)

type GameAppTemplateService struct {
	repo *gameRepo.GameAppTemplateRepository
}

func NewGameAppTemplateService() *GameAppTemplateService {
	return &GameAppTemplateService{repo: gameRepo.NewGameAppTemplateRepository()}
}

func (s *GameAppTemplateService) Create(t *game.GameAppTemplate) error {
	return s.repo.Create(t)
}

func (s *GameAppTemplateService) FindByID(id uint) (*game.GameAppTemplate, error) {
	return s.repo.FindByID(id)
}

func (s *GameAppTemplateService) FindPage(page, size int, keyword string, status int) ([]game.GameAppTemplate, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status)
}

func (s *GameAppTemplateService) FindAll() ([]game.GameAppTemplate, error) {
	return s.repo.FindAll()
}

func (s *GameAppTemplateService) Update(id uint, t *game.GameAppTemplate) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("SDK模板不存在")
	}
	if t.Name != "" {
		existing.Name = t.Name
	}
	if t.AllowAge > 0 {
		existing.AllowAge = t.AllowAge
	}
	if t.PrivacyURL != "" {
		existing.PrivacyURL = t.PrivacyURL
	}
	if t.AgreementURL != "" {
		existing.AgreementURL = t.AgreementURL
	}
	existing.IsOpenRealname = t.IsOpenRealname
	existing.IsOpenRegister = t.IsOpenRegister
	existing.IsOpenCharge = t.IsOpenCharge
	existing.IsAlertEmail = t.IsAlertEmail
	existing.IsAlertPhone = t.IsAlertPhone
	existing.IsAlertAuth = t.IsAlertAuth
	existing.IsOpenFloat = t.IsOpenFloat
	existing.Status = t.Status
	return s.repo.Update(existing)
}

func (s *GameAppTemplateService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("SDK模板不存在")
	}
	return s.repo.Delete(id)
}

func (s *GameAppTemplateService) FindOptions() ([]dict.Option, error) {
	return s.repo.FindOptions()
}
