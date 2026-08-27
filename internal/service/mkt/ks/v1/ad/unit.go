package ad

import (
	"errors"
	"strconv"

	ksadRepo "stack-bm/internal/repository/mkt/ks/ad"
	kuaishouAPI "stack-bm/internal/service/mkt/ks/v1/api"
	kstoken "stack-bm/internal/service/mkt/ks/v1/token"
)

// UnitService 快手广告（第二层级）业务：列表查询 + 平台操作
type UnitService struct {
	repo   *ksadRepo.UnitRepository
	tokens *kstoken.TokenService
}

func NewUnitService() *UnitService {
	return &UnitService{
		repo:   ksadRepo.NewUnitRepository(),
		tokens: kstoken.NewTokenService(),
	}
}

// List 分页查询广告数据
func (s *UnitService) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(page, size, columns, filters)
}

// resolve 广告ID(aid) -> 平台账户ID + access_token
func (s *UnitService) resolve(id int) (int, string, error) {
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

func (s *UnitService) Open(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitStatus(token, advID, kuaishouAPI.KsStatusOpen, []int{id})
}

func (s *UnitService) Pause(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitStatus(token, advID, kuaishouAPI.KsStatusPause, []int{id})
}

func (s *UnitService) Delete(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitStatus(token, advID, kuaishouAPI.KsStatusDelete, []int{id})
}

func (s *UnitService) SetBudget(id int, budget float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitBudget(token, advID, id, budget)
}

func (s *UnitService) SetBid(id int, bid float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitBid(token, advID, id, bid, false)
}

func (s *UnitService) SetDeepBid(id int, bid float64) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitBid(token, advID, id, bid, true)
}

func (s *UnitService) SetBeginDate(id int, beginDate string) (bool, error) {
	// 快手广告投放日期更新接口暂未实现
	return true, nil
}

// BatchUpdateStatus 批量启停：接口单条语义，仅提交首个ID（行为与旧实现一致）
func (s *UnitService) BatchUpdateStatus(ids []int, status int) (bool, error) {
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
	return kuaishouAPI.UpdateUnitStatus(token, advID, op, ids)
}

func (s *UnitService) BatchDelete(ids []int) (bool, error) {
	if len(ids) == 0 {
		return false, errEmptyIDs
	}
	advID, token, err := s.resolve(ids[0])
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateUnitStatus(token, advID, kuaishouAPI.KsStatusDelete, ids)
}

func (s *UnitService) SetRaise(id int) (bool, error) { return true, nil }
func (s *UnitService) StopRaise(id int) (bool, error) {
	return true, nil
}
func (s *UnitService) BatchSetRaise(ids []int) (bool, error)  { return true, nil }
func (s *UnitService) BatchStopRaise(ids []int) (bool, error) { return true, nil }
func (s *UnitService) Collect(id int) (bool, error)           { return true, nil }
func (s *UnitService) CancelCollect(id int) (bool, error)     { return true, nil }

func (s *UnitService) RaiseInfo(id int) (map[string]interface{}, error) {
	return map[string]interface{}{"status": "无起量信息"}, nil
}
