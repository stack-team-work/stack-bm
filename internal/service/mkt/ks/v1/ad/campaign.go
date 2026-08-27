package ad

import (
	"errors"
	"strconv"

	ksadRepo "stack-bm/internal/repository/mkt/ks/ad"
	kuaishouAPI "stack-bm/internal/service/mkt/ks/v1/api"
	kstoken "stack-bm/internal/service/mkt/ks/v1/token"
)

var errEmptyIDs = errors.New("ids不能为空")

// CampaignService 快手广告组（第一层级）业务：列表查询 + 平台操作
type CampaignService struct {
	repo   *ksadRepo.CampaignRepository
	tokens *kstoken.TokenService
}

func NewCampaignService() *CampaignService {
	return &CampaignService{
		repo:   ksadRepo.NewCampaignRepository(),
		tokens: kstoken.NewTokenService(),
	}
}

// List 分页查询广告组数据
func (s *CampaignService) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(page, size, columns, filters)
}

// resolve 广告组ID(cpid) -> 平台账户ID + access_token
func (s *CampaignService) resolve(id int) (int, string, error) {
	if id <= 0 {
		return 0, "", errors.New("id不能为空")
	}
	advID, err := s.repo.ResolveAccount(id)
	if err != nil {
		return 0, "", err
	}
	token, err := s.tokens.AccessTokenByUID(strconv.Itoa(advID))
	if err != nil {
		return 0, "", err
	}
	return advID, token, nil
}

func (s *CampaignService) Open(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateCampaignStatus(token, advID, kuaishouAPI.KsStatusOpen, []int{id})
}

func (s *CampaignService) Pause(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateCampaignStatus(token, advID, kuaishouAPI.KsStatusPause, []int{id})
}

func (s *CampaignService) Delete(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateCampaignStatus(token, advID, kuaishouAPI.KsStatusDelete, []int{id})
}

// SetBudget 广告组预算：快手无广告组预算接口，沿用旧实现的单元预算端点（行为保持）
func (s *CampaignService) SetBudget(id int, budget float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitBudget(token, advID, id, budget)
}

// SetBid 广告组出价：快手无广告组出价接口，沿用旧实现的单元出价端点（行为保持）
func (s *CampaignService) SetBid(id int, bid float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitBid(token, advID, id, bid, false)
}

func (s *CampaignService) SetDeepBid(id int, bid float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitBid(token, advID, id, bid, true)
}

// BatchUpdateStatus 批量启停：status==1 开启，其余暂停；接口单条语义，仅提交首个ID（行为与旧实现一致）
func (s *CampaignService) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	op := kuaishouAPI.KsStatusOpen
	if status != 1 {
		op = kuaishouAPI.KsStatusPause
	}
	advID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateCampaignStatus(token, advID, op, ids)
}

func (s *CampaignService) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	advID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateCampaignStatus(token, advID, kuaishouAPI.KsStatusDelete, ids)
}
