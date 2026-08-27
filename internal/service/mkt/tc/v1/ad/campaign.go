package ad

import (
	"errors"
	"strconv"

	tcadRepo "stack-bm/internal/repository/mkt/tc/ad"
	tcAPI "stack-bm/internal/service/mkt/tc/v1/api"
	tctoken "stack-bm/internal/service/mkt/tc/v1/token"
)

var errEmptyIDs = errors.New("ids不能为空")

// CampaignService 腾讯广告组（第一层级）业务：列表查询 + 平台操作
type CampaignService struct {
	repo   *tcadRepo.CampaignRepository
	tokens *tctoken.TokenService
}

func NewCampaignService() *CampaignService {
	return &CampaignService{
		repo:   tcadRepo.NewCampaignRepository(),
		tokens: tctoken.NewTokenService(),
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

// resolve 广告组ID(cpid) -> 平台账户ID(字符串, 广点通接口参数) + access_token
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
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return tcAPI.UpdateAdgroupStatus(token, accID, id, tcAPI.TcAdStatusNormal)
}

func (s *CampaignService) Pause(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return tcAPI.UpdateAdgroupStatus(token, accID, id, tcAPI.TcAdStatusPause)
}

func (s *CampaignService) Delete(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return tcAPI.DeleteAdgroup(token, accID, id)
}

func (s *CampaignService) SetBudget(id int, budget float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return tcAPI.UpdateAdgroupBudget(token, accID, id, budget)
}

func (s *CampaignService) SetBid(id int, bid float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return tcAPI.UpdateAdgroupBid(token, accID, id, bid)
}

func (s *CampaignService) SetDeepBid(id int, bid float64) (bool, error) {
	// 广点通无深度出价接口，沿用普通出价端点（行为与旧实现一致）
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return tcAPI.UpdateAdgroupBid(token, accID, id, bid)
}

// BatchUpdateStatus 批量启停：status==1 启用，其余暂停；逐条提交
func (s *CampaignService) BatchUpdateStatus(ids []int, status int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	cs := tcAPI.TcAdStatusNormal
	if status != 1 {
		cs = tcAPI.TcAdStatusPause
	}
	for _, id := range ids {
		accID, token, err := s.resolve(id)
		if err != nil {
			return false, err
		}
		if _, err := tcAPI.UpdateAdgroupStatus(token, accID, id, cs); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (s *CampaignService) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	for _, id := range ids {
		accID, token, err := s.resolve(id)
		if err != nil {
			return false, err
		}
		if _, err := tcAPI.DeleteAdgroup(token, accID, id); err != nil {
			return false, err
		}
	}
	return true, nil
}
