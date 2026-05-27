package pay

import (
	"errors"

	"stack-bm/internal/model/sdk/pay"
	payRepo "stack-bm/internal/repository/sdk/pay"
	"stack-bm/pkg/utils"
)

type PayPlatformService struct {
	repo *payRepo.PayPlatformRepository
}

func NewPayPlatformService() *PayPlatformService {
	return &PayPlatformService{repo: payRepo.NewPayPlatformRepository()}
}

func (s *PayPlatformService) Create(p *pay.PayPlatform) error {
	if p.Mark == "" { p.Mark = utils.ToPinYinMark(p.Name) }
	return s.repo.Create(p)
}

func (s *PayPlatformService) FindByID(id uint) (*pay.PayPlatform, error) { return s.repo.FindByID(id) }

func (s *PayPlatformService) FindPage(page, size int, keyword string) ([]pay.PayPlatform, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword)
}

func (s *PayPlatformService) FindAll() ([]pay.PayPlatform, error) { return s.repo.FindAll() }

func (s *PayPlatformService) Update(id uint, p *pay.PayPlatform) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("支付平台不存在") }
	if p.Name != "" { existing.Name = p.Name }
	if p.Mark != "" { existing.Mark = p.Mark }
	return s.repo.Update(existing)
}

func (s *PayPlatformService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("支付平台不存在") }
	return s.repo.Delete(id)
}
