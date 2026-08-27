package ad

import (
	"errors"
	"strconv"

	biliadRepo "stack-bm/internal/repository/mkt/bili/ad"
	"stack-bm/internal/service/mkt/bili/v1/api"
	bitoken "stack-bm/internal/service/mkt/bili/v1/token"
)

// CreativeService B站创意（第三层级）业务：列表查询 + 平台操作
type CreativeService struct {
	repo   *biliadRepo.CreativeRepository
	tokens *bitoken.TokenService
}

func NewCreativeService() *CreativeService {
	return &CreativeService{
		repo:   biliadRepo.NewCreativeRepository(),
		tokens: bitoken.NewTokenService(),
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

func (s *CreativeService) Open(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, accID, api.BiliOpOpen, []int{id})
}

func (s *CreativeService) Pause(id int) (bool, error) {
	accID, token, err := s.resolve(id)
	if err != nil {
		return false, err
	}
	return api.UpdateCreativeStatus(token, accID, api.BiliOpStop, []int{id})
}

func (s *CreativeService) BatchUpdateStatus(ids []int, status int) (bool, error) {
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
	return api.UpdateCreativeStatus(token, accID, op, ids)
}

func (s *CreativeService) Preview(id int) (string, error) {
	// 创意预览URL待媒体接口确认后实现
	return "", nil
}
