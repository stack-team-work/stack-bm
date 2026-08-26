package sync

import (
	"errors"

	tcModel "stack-bm/internal/model/mkt/tc"
	tcRepo "stack-bm/internal/repository/mkt/tc"
)

// AdDataService 广告数据
type AdDataService struct {
	repo *tcRepo.AdDataRepository
}

func NewAdDataService() *AdDataService {
	return &AdDataService{repo: tcRepo.NewAdDataRepository()}
}

// List 广告数据列表（表未建，返回空；columns 用于字段投影，filters 用于筛选条件）
func (s *AdDataService) List(level string, page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	collection, ok := tcModel.AdDataLevelCollections[level]
	if !ok {
		return nil, 0, errors.New("未知的广告数据层级: " + level)
	}
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(collection, page, size, columns, filters)
}