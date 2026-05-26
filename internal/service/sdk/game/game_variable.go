package game

import (
	"encoding/json"
	"errors"
	"strings"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/utils"
)

type GameVariableService struct {
	repo *gameRepo.GameVariableRepository
}

func NewGameVariableService() *GameVariableService {
	return &GameVariableService{repo: gameRepo.NewGameVariableRepository()}
}

func encodeValue(v string) string {
	s := strings.TrimSpace(v)
	if s == "" { return "" }
	if strings.Contains(s, "\n") {
		lines := strings.Split(s, "\n")
		arr := make([]string, 0, len(lines))
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" { arr = append(arr, trimmed) }
		}
		if len(arr) == 0 { return "" }
		data, _ := json.Marshal(arr)
		return string(data)
	}
	return s
}

func decodeValue(v string) string {
	s := strings.TrimSpace(v)
	if s == "" { return "" }
	if strings.HasPrefix(s, "[") {
		var arr []string
		if err := json.Unmarshal([]byte(s), &arr); err == nil { return strings.Join(arr, "\n") }
	}
	return s
}

func (s *GameVariableService) Create(v *game.GameVariable) error {
	if v.Mark == "" {
		v.Mark = utils.ToPinYinMark(v.Name)
	}
	v.Value = encodeValue(v.Value)
	return s.repo.Create(v)
}

func (s *GameVariableService) FindByID(id uint) (*game.GameVariable, error) {
	v, err := s.repo.FindByID(id)
	if err != nil { return nil, err }
	v.Value = decodeValue(v.Value)
	return v, nil
}

func (s *GameVariableService) FindPage(page, size int, keyword string, status int) ([]game.GameVariable, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	vars, total, err := s.repo.FindPage(page, size, keyword, status)
	if err != nil { return nil, 0, err }
	for i := range vars { vars[i].Value = decodeValue(vars[i].Value) }
	return vars, total, nil
}

func (s *GameVariableService) Update(id uint, v *game.GameVariable) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("变量不存在") }
	if v.Name != "" { existing.Name = v.Name }
	if v.Key != "" { existing.Key = v.Key }
	existing.Value = encodeValue(v.Value)
	if v.Mark != "" { existing.Mark = v.Mark }
	existing.Status = v.Status
	return s.repo.Update(existing)
}

func (s *GameVariableService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("变量不存在") }
	return s.repo.Delete(id)
}
