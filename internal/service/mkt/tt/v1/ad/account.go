package ad

import (
	ttadRepo "stack-bm/internal/repository/mkt/tt/ad"
)

// AccountService 头条账户业务：列表查询（V3暂无余额接口，余额同步为占位）
type AccountService struct {
	repo *ttadRepo.AccountRepository
}

func NewAccountService() *AccountService {
	return &AccountService{repo: ttadRepo.NewAccountRepository()}
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
