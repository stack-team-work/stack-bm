package ad

import (
	"errors"
	"strconv"

	ksadRepo "stack-bm/internal/repository/mkt/ks/ad"
	kuaishouAPI "stack-bm/internal/service/mkt/ks/v1/api"
	kstoken "stack-bm/internal/service/mkt/ks/v1/token"
)

// CreativeService 快手创意（第三层级）业务：列表查询 + 平台操作
type CreativeService struct {
	repo   *ksadRepo.CreativeRepository
	tokens *kstoken.TokenService
}

func NewCreativeService() *CreativeService {
	return &CreativeService{
		repo:   ksadRepo.NewCreativeRepository(),
		tokens: kstoken.NewTokenService(),
	}
}

// List 分页查询创意数据
func (s *CreativeService) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(page, size, columns, filters)
}

// resolve 创意ID(cid) -> 平台账户ID + access_token
func (s *CreativeService) resolve(id int) (int, string, error) {
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

func (s *CreativeService) Open(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateCreativeStatus(token, advID, kuaishouAPI.KsStatusOpen, []int{id})
}

func (s *CreativeService) Pause(id int) (bool, error) {
	advID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return kuaishouAPI.UpdateCreativeStatus(token, advID, kuaishouAPI.KsStatusPause, []int{id})
}

// BatchUpdateStatus 批量启停：接口单条语义，仅提交首个ID（行为与旧实现一致）
func (s *CreativeService) BatchUpdateStatus(ids []int, status int) (bool, error) {
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
	return kuaishouAPI.UpdateCreativeStatus(token, advID, op, ids)
}

func (s *CreativeService) Preview(id int) (string, error) {
	// 创意预览URL待媒体接口确认后实现
	return "", nil
}
