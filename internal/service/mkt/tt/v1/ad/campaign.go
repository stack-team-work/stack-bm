package ad

import (
	"errors"
	"strconv"

	ttadRepo "stack-bm/internal/repository/mkt/tt/ad"
	ttAPI "stack-bm/internal/service/mkt/tt/v1/api"
	tttoken "stack-bm/internal/service/mkt/tt/v1/token"
)

var errEmptyIDs = errors.New("ids不能为空")

// CampaignService 头条V3项目（第一层级）业务：列表查询 + 平台操作
type CampaignService struct {
	repo   *ttadRepo.CampaignRepository
	tokens *tttoken.TokenService
}

func NewCampaignService() *CampaignService {
	return &CampaignService{
		repo:   ttadRepo.NewCampaignRepository(),
		tokens: tttoken.NewTokenService(),
	}
}

// List 分页查询项目数据
func (s *CampaignService) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(page, size, columns, filters)
}

// resolve 项目ID(cpid) -> 平台账户ID(字符串, 巨量接口参数) + access_token
func (s *CampaignService) resolve(id int) (string, string, error) {
	if id <= 0 {
		return "", "", errors.New("id不能为空")
	}
	accID, err := s.repo.ResolveAccount(id)
	if err != nil {
		return "", "", err
	}
	token, err := s.tokens.AccessTokenByUID(strconv.Itoa(accID))
	if err != nil {
		return "", "", err
	}
	return strconv.Itoa(accID), token, nil
}

func (s *CampaignService) Open(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return ttAPI.UpdateProjectStatus(token, advID, ttAPI.TtOptStatusEnable, []int{id})
}

func (s *CampaignService) Pause(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return ttAPI.UpdateProjectStatus(token, advID, ttAPI.TtOptStatusDisable, []int{id})
}

func (s *CampaignService) Delete(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return ttAPI.DeleteProject(token, advID, []int{id})
}

func (s *CampaignService) SetBudget(id int, budget float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return ttAPI.UpdateProjectBudget(token, advID, id, budget)
}

func (s *CampaignService) SetBid(id int, bid float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return ttAPI.UpdateProjectBid(token, advID, id, bid, false)
}

func (s *CampaignService) SetDeepBid(id int, bid float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return ttAPI.UpdateProjectBid(token, advID, id, bid, true)
}

// BatchUpdateStatus 批量启停：status==1 开启，其余暂停
func (s *CampaignService) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	opt := ttAPI.TtOptStatusEnable
	if status != 1 {
		opt = ttAPI.TtOptStatusDisable
	}
	advID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return ttAPI.UpdateProjectStatus(token, advID, opt, ids)
}

func (s *CampaignService) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	advID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return ttAPI.DeleteProject(token, advID, ids)
}
