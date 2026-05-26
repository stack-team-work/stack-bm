package sys

import (
	sdkSysRepo "stack-bm/internal/repository/sdk/sys"
	"stack-bm/internal/model/sdk/sys"
)

type SysLogService struct {
	repo *sdkSysRepo.SysLogRepository
}

func NewSysLogService() *SysLogService {
	return &SysLogService{repo: sdkSysRepo.NewSysLogRepository()}
}

func (s *SysLogService) FindPage(page, size int, keyword string, level int, type_ int) ([]sys.SysLog, int64, error) {
	if page < 1 { page = 1 }
	if size < 1 { size = 10 }
	return s.repo.FindPage(page, size, keyword, level, type_)
}
