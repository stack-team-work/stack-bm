package ad

import (
	tcadRepo "stack-bm/internal/repository/mkt/tc/ad"
)

// AccountService 腾讯账户业务：列表查询（余额经 sync.SyncBalance 回填展示）
type AccountService struct {
	repo *tcadRepo.AccountRepository
}

func NewAccountService() *AccountService {
	return &AccountService{repo: tcadRepo.NewAccountRepository()}
}

// List 分页查询账户数据
func (s *AccountService) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(page, size, columns, filters)
}
