package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/utils"
)

type GameAppService struct {
	repo *gameRepo.GameAppRepository
}

func NewGameAppService() *GameAppService {
	return &GameAppService{repo: gameRepo.NewGameAppRepository()}
}

func (s *GameAppService) Create(app *game.GameApp) error {
	app.AppKey = utils.RandomString(16)
	app.AppSecret = utils.RandomString(32)
	return s.repo.Create(app)
}

func (s *GameAppService) FindByID(id uint) (*game.GameApp, error) { return s.repo.FindByID(id) }

func (s *GameAppService) FindPage(page, size int, keyword string, gameID int, status int) ([]game.GameApp, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, gameID, status)
}

func (s *GameAppService) FindAll() ([]game.GameApp, error) { return s.repo.FindAll() }

func (s *GameAppService) Update(id uint, app *game.GameApp) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("应用不存在")
	}
	if app.Name != "" {
		existing.Name = app.Name
	}
	if app.Pid > 0 {
		existing.Pid = app.Pid
	}
	if app.PackageName != "" {
		existing.PackageName = app.PackageName
	}
	if app.AppName != "" {
		existing.AppName = app.AppName
	}
	if app.AppTemplateID > 0 {
		existing.AppTemplateID = app.AppTemplateID
	}
	if app.Os > 0 {
		existing.Os = app.Os
	}
	if app.SdkVer != "" {
		existing.SdkVer = app.SdkVer
	}
	if app.AppVer != "" {
		existing.AppVer = app.AppVer
	}
	if app.CallbackURL != "" {
		existing.CallbackURL = app.CallbackURL
	}
	if app.ApiDomain != "" {
		existing.ApiDomain = app.ApiDomain
	}
	if app.PayDomain != "" {
		existing.PayDomain = app.PayDomain
	}
	if app.CsParams != "" {
		existing.CsParams = app.CsParams
	}
	if app.PayParams != "" {
		existing.PayParams = app.PayParams
	}
	if app.H5Params != "" {
		existing.H5Params = app.H5Params
	}
	existing.Status = app.Status
	return s.repo.Update(existing)
}

func (s *GameAppService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("应用不存在")
	}
	return s.repo.Delete(id)
}
