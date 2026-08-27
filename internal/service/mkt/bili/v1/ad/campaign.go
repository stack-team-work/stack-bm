package ad

import (
	"errors"
	"strconv"

	biliadRepo "stack-bm/internal/repository/mkt/bili/ad"
	"stack-bm/internal/service/mkt/bili/v1/api"
	bitoken "stack-bm/internal/service/mkt/bili/v1/token"
)

var errEmptyIDs = errors.New("ids不能为空")

// CampaignService B站计划（第一层级）业务：列表查询 + 平台操作
type CampaignService struct {
	repo   *biliadRepo.CampaignRepository
	tokens *bitoken.TokenService
}

func NewCampaignService() *CampaignService {
	return &CampaignService{
		repo:   biliadRepo.NewCampaignRepository(),
		tokens: bitoken.NewTokenService(),
	}
}

// List 分页查询计划数据
func (s *CampaignService) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(page, size, columns, filters)
}

// resolve 计划ID(cpid) -> 平台账户ID + access_token
func (s *CampaignService) resolve(id int) (int, string, error) {
	if id <= 0 {
		return 0, "", errors.New("id不能为空")
	}
	accID, err := s.repo.ResolveAccount(id)
	if err != nil {
		return 0, "", err
	}
	token, err := s.tokens.AccessTokenByUID(strconv.Itoa(accID))
	if err != nil {
		return 0, "", err
	}
	return accID, token, nil
}

func (s *CampaignService) Open(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpOpen, []int{id})
}

func (s *CampaignService) Pause(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpStop, []int{id})
}

func (s *CampaignService) Delete(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpDelete, []int{id})
}

func (s *CampaignService) SetBudget(id int, budget float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignBudget(token, accID, id, budget)
}

func (s *CampaignService) SetBid(id int, bid float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignBid(token, accID, id, bid, false)
}

func (s *CampaignService) SetDeepBid(id int, bid float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignBid(token, accID, id, bid, true)
}

// BatchUpdateStatus 批量启停：status==1 开启，其余暂停（同一批应属同一账户，取首个解析）
func (s *CampaignService) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	op := api.BiliOpOpen
	if status != 1 {
		op = api.BiliOpStop
	}
	accID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, op, ids)
}

func (s *CampaignService) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	accID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateCampaignStatus(token, accID, api.BiliOpDelete, ids)
}
