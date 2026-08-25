package sync

import (
	"errors"

	ttModel "stack-bm/internal/model/mkt/tt"
	ttRepo "stack-bm/internal/repository/mkt/tt"
)

// AdDataService 广告数据
type AdDataService struct {
	repo *ttRepo.AdDataRepository
}

func NewAdDataService() *AdDataService {
	return &AdDataService{repo: ttRepo.NewAdDataRepository()}
}

// List 广告数据列表（表未建，返回空）
func (s *AdDataService) List(level string, page, size int, params map[string]interface{}) ([]map[string]interface{}, int64, error) {
	collection, ok := ttModel.AdDataLevelCollections[level]
	if !ok {
		return nil, 0, errors.New("未知的广告数据层级: " + level)
	}
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.List(collection, page, size, params)
}