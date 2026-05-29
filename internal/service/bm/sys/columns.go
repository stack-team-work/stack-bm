package sys

import (
	"errors"

	"stack-bm/internal/model/bm/sys"
	bmSysRepo "stack-bm/internal/repository/bm/sys"
	"stack-bm/pkg/utils"
)

type SysColumnService struct {
	repo *bmSysRepo.SysColumnRepository
}

func NewSysColumnService() *SysColumnService {
	return &SysColumnService{repo: bmSysRepo.NewSysColumnRepository()}
}

func (s *SysColumnService) Create(c *sys.SysColumn) error {
	if c.Mark == "" {
		c.Mark = utils.ToPinYinMark(c.Name)
	}
	return s.repo.Create(c)
}

func (s *SysColumnService) FindByID(id uint) (*sys.SysColumn, error) {
	return s.repo.FindByID(id)
}

func (s *SysColumnService) FindPage(page, size int, keyword string, status int, reportType int, indicatorType int) ([]sys.SysColumn, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return s.repo.FindPage(page, size, keyword, status, reportType, indicatorType)
}

func (s *SysColumnService) FindAll() ([]sys.SysColumn, error) {
	return s.repo.FindAll()
}

func (s *SysColumnService) Update(id uint, c *sys.SysColumn) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("报表指标不存在")
	}
	if c.Name != "" {
		existing.Name = c.Name
	}
	if c.Field != "" {
		existing.Field = c.Field
	}
	if c.Mark != "" {
		existing.Mark = c.Mark
	}
	if c.ReportType > 0 {
		existing.ReportType = c.ReportType
	}
	if c.IndicatorType > 0 {
		existing.IndicatorType = c.IndicatorType
	}
	existing.Default = c.Default
	existing.Status = c.Status
	if c.AdminID > 0 {
		existing.AdminID = c.AdminID
	}
	return s.repo.Update(existing)
}

func (s *SysColumnService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("报表指标不存在")
	}
	return s.repo.Delete(id)
}
