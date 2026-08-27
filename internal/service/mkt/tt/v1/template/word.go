package template

import (
	ttModel "stack-bm/internal/model/mkt/tt"
	ttRepo "stack-bm/internal/repository/mkt/tt"
)

// WordService 头条词包
type WordService struct {
	repo *ttRepo.WordListRepository
}

func NewWordService() *WordService {
	return &WordService{repo: ttRepo.NewWordListRepository()}
}

func (s *WordService) ListAll() ([]ttModel.WordList, error) {
	return s.repo.ListAll()
}
