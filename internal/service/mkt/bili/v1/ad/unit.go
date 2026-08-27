package ad

import (
	"errors"
	"strconv"

	biliadRepo "stack-bm/internal/repository/mkt/bili/ad"
	"stack-bm/internal/service/mkt/bili/v1/api"
	bitoken "stack-bm/internal/service/mkt/bili/v1/token"
)

// UnitService B站单元（第二层级）业务：列表查询 + 平台操作
type UnitService struct {
	repo   *biliadRepo.UnitRepository
	tokens *bitoken.TokenService
}

func NewUnitService() *UnitService {
	return &UnitService{
		repo:   biliadRepo.NewUnitRepository(),
		tokens: bitoken.NewTokenService(),
	}
}

// List 分页查询单元数据
func (s *UnitService) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(page, size, columns, filters)
}

// resolve 单元ID(aid) -> 平台账户ID + access_token
func (s *UnitService) resolve(id int) (int, string, error) {
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

func (s *UnitService) Open(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpOpen, []int{id})
}

func (s *UnitService) Pause(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpStop, []int{id})
}

func (s *UnitService) Delete(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpDelete, []int{id})
}

func (s *UnitService) SetBudget(id int, budget float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBudget(token, accID, id, budget)
}

func (s *UnitService) SetBid(id int, bid float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, accID, id, bid, false)
}

func (s *UnitService) SetDeepBid(id int, bid float64) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateUnitBid(token, accID, id, bid, true)
}

func (s *UnitService) SetBeginDate(id int, beginDate string) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	// 投放时间操作：起投日期参数待B站接口字段确认，当前仅按操作类型提交
	return api.UpdateUnitStatus(token, accID, api.BiliOpDate, []int{id})
}

func (s *UnitService) BatchUpdateStatus(ids []int, status int) (bool, error) {
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
	return api.UpdateUnitStatus(token, accID, op, ids)
}

func (s *UnitService) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	accID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return api.UpdateUnitStatus(token, accID, api.BiliOpDelete, ids)
}

func (s *UnitService) SetRaise(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.OpenAccelerate(token, accID, id, 0, 1, 6)
}

func (s *UnitService) StopRaise(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.CloseAccelerate(token, accID, id)
}

func (s *UnitService) BatchSetRaise(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	for _, id := range ids {
		if _, err := s.SetRaise(id); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (s *UnitService) BatchStopRaise(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	for _, id := range ids {
		if _, err := s.StopRaise(id); err != nil {
			return false, err
		}
	}
	return true, nil
}

func (s *UnitService) Collect(id int) (bool, error) {
	// 收藏为本地功能，待收藏表就绪后实现
	return true, nil
}

func (s *UnitService) CancelCollect(id int) (bool, error) {
	// 收藏为本地功能，待收藏表就绪后实现
	return true, nil
}

func (s *UnitService) RaiseInfo(id int) (map[string]interface{}, error) {
	return map[string]interface{}{"status": "无起量信息"}, nil
}
