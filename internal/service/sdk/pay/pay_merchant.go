package pay

import (
	"errors"

	"stack-bm/internal/model/sdk/pay"
	payRepo "stack-bm/internal/repository/sdk/pay"
	"stack-bm/pkg/utils"
)

type PayMerchantService struct {
	repo *payRepo.PayMerchantRepository
}

func NewPayMerchantService() *PayMerchantService {
	return &PayMerchantService{repo: payRepo.NewPayMerchantRepository()}
}

func (s *PayMerchantService) Create(p *pay.PayMerchant) error {
	if p.Mark == "" { p.Mark = utils.ToPinYinMark(p.Name) }
	return s.repo.Create(p)
}

func (s *PayMerchantService) FindByID(id uint) (*pay.PayMerchant, error) { return s.repo.FindByID(id) }

func (s *PayMerchantService) FindPage(page, size int, keyword string, status int, payType int) ([]pay.PayMerchant, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, status, payType)
}

func (s *PayMerchantService) FindAll() ([]pay.PayMerchant, error) { return s.repo.FindAll() }

func (s *PayMerchantService) Update(id uint, p *pay.PayMerchant) error {
	existing, err := s.repo.FindByID(id)
	if err != nil { return errors.New("支付商户不存在") }
	if p.Name != "" { existing.Name = p.Name }
	if p.ShowName != "" { existing.ShowName = p.ShowName }
	if p.Type > 0 { existing.Type = p.Type }
	if p.PlatformMark > 0 { existing.PlatformMark = p.PlatformMark }
	if p.Mark != "" { existing.Mark = p.Mark }
	if p.URL != "" { existing.URL = p.URL }
	if p.Rate > 0 { existing.Rate = p.Rate }
	if p.Config != "" { existing.Config = p.Config }
	existing.Weight = p.Weight
	existing.Status = p.Status
	return s.repo.Update(existing)
}

func (s *PayMerchantService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil { return errors.New("支付商户不存在") }
	return s.repo.Delete(id)
}
