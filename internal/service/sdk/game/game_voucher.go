package game

import (
	"errors"

	"stack-bm/internal/model/sdk/game"
	gameRepo "stack-bm/internal/repository/sdk/game"
	"stack-bm/pkg/dict"
)

type GameVoucherService struct{ repo *gameRepo.GameVoucherRepository }

func NewGameVoucherService() *GameVoucherService {
	return &GameVoucherService{repo: gameRepo.NewGameVoucherRepository()}
}

func (s *GameVoucherService) Create(v *game.GameVoucher) error { return s.repo.Create(v) }
func (s *GameVoucherService) FindByID(id uint) (*game.GameVoucher, error) { return s.repo.FindByID(id) }
func (s *GameVoucherService) FindPage(page, size int, keyword string, status int) ([]game.GameVoucher, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status)
}
func (s *GameVoucherService) FindAll() ([]game.GameVoucher, error) { return s.repo.FindAll() }

func (s *GameVoucherService) Update(id uint, v *game.GameVoucher) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("代金券不存在")
	}
	if v.Name != "" {
		existing.Name = v.Name
	}
	if v.Desc != "" {
		existing.Desc = v.Desc
	}
	existing.UseType = v.UseType
	existing.UseLimit = v.UseLimit
	existing.Total = v.Total
	existing.TotalFee = v.TotalFee
	if v.Stime > 0 {
		existing.Stime = v.Stime
	}
	if v.Etime > 0 {
		existing.Etime = v.Etime
	}
	existing.Status = v.Status
	return s.repo.Update(existing)
}

func (s *GameVoucherService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("代金券不存在")
	}
	return s.repo.Delete(id)
}

func (s *GameVoucherService) FindOptions() ([]dict.Option, error) { return s.repo.FindOptions() }
