package tt

import (
	ttModel "stack-bm/internal/model/mkt/tt"
	ttRepo "stack-bm/internal/repository/mkt/tt"
)

type WordListService struct {
	repo *ttRepo.WordListRepository
}

func NewWordListService() *WordListService {
	return &WordListService{repo: ttRepo.NewWordListRepository()}
}

func (s *WordListService) ListAll() ([]ttModel.WordList, error) {
	return s.repo.ListAll()
}
