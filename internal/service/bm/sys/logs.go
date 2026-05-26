package sys

import (
	bmSysRepo "stack-bm/internal/repository/bm/sys"
	"stack-bm/internal/model/bm/sys"
)

type SysLogService struct {
	repo *bmSysRepo.SysLogRepository
}

func NewSysLogService() *SysLogService {
	return &SysLogService{repo: bmSysRepo.NewSysLogRepository()}
}

func (s *SysLogService) Create(log *sys.SysLog) error { return s.repo.Create(log) }

func (s *SysLogService) FindPage(page, size int, keyword string, level string) ([]sys.SysLog, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, level)
}

func (s *SysLogService) ClearAll() error { return s.repo.ClearAll() }
