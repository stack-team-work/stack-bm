package game

import (
	"encoding/json"
	"errors"
	"strings"

	"stack-bm/internal/model/game"
	gameRepo "stack-bm/internal/repository/game"
	"stack-bm/pkg/utils"
)

type GameService struct {
	repo *gameRepo.GameRepository
}

func NewGameService() *GameService {
	return &GameService{
		repo: gameRepo.NewGameRepository(),
	}
}

func (s *GameService) Create(g *game.Game) error {
	return s.repo.Create(g)
}

func (s *GameService) FindByID(id uint) (*game.Game, error) {
	return s.repo.FindByID(id)
}

func (s *GameService) FindPage(page, size int, keyword string, status int) ([]game.Game, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status)
}

func (s *GameService) FindAll() ([]game.Game, error) {
	return s.repo.FindAll()
}

func (s *GameService) Update(id uint, g *game.Game) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("游戏不存在")
	}

	if g.Name != "" {
		existing.Name = g.Name
	}
	if g.Mark != "" {
		existing.Mark = g.Mark
	}
	if g.WebName != "" {
		existing.WebName = g.WebName
	}
	if g.Icon != "" {
		existing.Icon = g.Icon
	}
	existing.CpID = g.CpID
	existing.Status = g.Status

	return s.repo.Update(existing)
}

func (s *GameService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("游戏不存在")
	}
	return s.repo.Delete(id)
}

type GameAppService struct {
	repo *gameRepo.GameAppRepository
}

func NewGameAppService() *GameAppService {
	return &GameAppService{
		repo: gameRepo.NewGameAppRepository(),
	}
}

func (s *GameAppService) Create(app *game.GameApp) error {
	app.AppKey = utils.RandomString(16)
	app.AppSecret = utils.RandomString(32)
	return s.repo.Create(app)
}

func (s *GameAppService) FindByID(id uint) (*game.GameApp, error) {
	return s.repo.FindByID(id)
}

func (s *GameAppService) FindPage(page, size int, keyword string, gameID int, status int) ([]game.GameApp, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, gameID, status)
}

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

type GameCpService struct {
	repo *gameRepo.GameCpRepository
}

func NewGameCpService() *GameCpService {
	return &GameCpService{
		repo: gameRepo.NewGameCpRepository(),
	}
}

func (s *GameCpService) Create(cp *game.GameCp) error {
	return s.repo.Create(cp)
}

func (s *GameCpService) FindByID(id uint) (*game.GameCp, error) {
	return s.repo.FindByID(id)
}

func (s *GameCpService) FindPage(page, size int, keyword string, status int) ([]game.GameCp, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status)
}

func (s *GameCpService) FindAll() ([]game.GameCp, error) {
	return s.repo.FindAll()
}

func (s *GameCpService) Update(id uint, cp *game.GameCp) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("CP不存在")
	}

	if cp.Name != "" {
		existing.Name = cp.Name
	}
	if cp.Mark != "" {
		existing.Mark = cp.Mark
	}
	if cp.Phone != "" {
		existing.Phone = cp.Phone
	}
	if cp.Addr != "" {
		existing.Addr = cp.Addr
	}
	existing.Status = cp.Status

	return s.repo.Update(existing)
}

func (s *GameCpService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("CP不存在")
	}
	return s.repo.Delete(id)
}

type GameTagService struct {
	repo *gameRepo.GameTagRepository
}

func NewGameTagService() *GameTagService {
	return &GameTagService{
		repo: gameRepo.NewGameTagRepository(),
	}
}

func (s *GameTagService) Create(tag *game.GameTag) error {
	return s.repo.Create(tag)
}

func (s *GameTagService) FindByID(id uint) (*game.GameTag, error) {
	return s.repo.FindByID(id)
}

func (s *GameTagService) FindPage(page, size int, keyword string, tagType int, status int) ([]game.GameTag, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, tagType, status)
}

func (s *GameTagService) FindAllByType(tagType int) ([]game.GameTag, error) {
	return s.repo.FindAllByType(tagType)
}

func (s *GameTagService) Update(id uint, tag *game.GameTag) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("标签不存在")
	}

	if tag.Name != "" {
		existing.Name = tag.Name
	}
	if tag.Mark != "" {
		existing.Mark = tag.Mark
	}
	existing.Status = tag.Status

	return s.repo.Update(existing)
}

func (s *GameTagService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("标签不存在")
	}
	return s.repo.Delete(id)
}

type GameVariableService struct {
	repo *gameRepo.GameVariableRepository
}

func NewGameVariableService() *GameVariableService {
	return &GameVariableService{
		repo: gameRepo.NewGameVariableRepository(),
	}
}

func encodeValue(v string) string {
	s := strings.TrimSpace(v)
	if s == "" {
		return ""
	}
	if strings.Contains(s, "\n") {
		lines := strings.Split(s, "\n")
		arr := make([]string, 0, len(lines))
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" {
				arr = append(arr, trimmed)
			}
		}
		if len(arr) == 0 {
			return ""
		}
		data, _ := json.Marshal(arr)
		return string(data)
	}
	return s
}

func decodeValue(v string) string {
	s := strings.TrimSpace(v)
	if s == "" {
		return ""
	}
	if strings.HasPrefix(s, "[") {
		var arr []string
		if err := json.Unmarshal([]byte(s), &arr); err == nil {
			return strings.Join(arr, "\n")
		}
	}
	return s
}

func (s *GameVariableService) Create(v *game.GameVariable) error {
	v.Value = encodeValue(v.Value)
	return s.repo.Create(v)
}

func (s *GameVariableService) FindByID(id uint) (*game.GameVariable, error) {
	v, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	v.Value = decodeValue(v.Value)
	return v, nil
}

func (s *GameVariableService) FindPage(page, size int, keyword string, status int) ([]game.GameVariable, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	vars, total, err := s.repo.FindPage(page, size, keyword, status)
	if err != nil {
		return nil, 0, err
	}
	for i := range vars {
		vars[i].Value = decodeValue(vars[i].Value)
	}
	return vars, total, nil
}

func (s *GameVariableService) Update(id uint, v *game.GameVariable) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("变量不存在")
	}

	if v.Name != "" {
		existing.Name = v.Name
	}
	if v.Key != "" {
		existing.Key = v.Key
	}
	existing.Value = encodeValue(v.Value)
	if v.Mark != "" {
		existing.Mark = v.Mark
	}
	existing.Status = v.Status

	return s.repo.Update(existing)
}

func (s *GameVariableService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("变量不存在")
	}
	return s.repo.Delete(id)
}
