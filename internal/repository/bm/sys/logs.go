package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"

	"gorm.io/gorm"
)

type SysLogRepository struct {
	db *gorm.DB
}

func NewSysLogRepository() *SysLogRepository {
	return &SysLogRepository{db: database.DBBM}
}

func (r *SysLogRepository) Create(log *sys.SysLog) error {
	return r.db.Create(log).Error
}

func (r *SysLogRepository) FindPage(page, size int, keyword string, level int) ([]sys.SysLog, int64, error) {
	var logs []sys.SysLog
	var total int64
	query := r.db.Model(&sys.SysLog{})
	if keyword != "" {
		query = query.Where("path LIKE ? OR username LIKE ? OR `desc` LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if level > 0 {
		query = query.Where("level = ?", level)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&logs).Error; err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

func (r *SysLogRepository) ClearAll() error {
	return r.db.Where("1 = 1").Delete(&sys.SysLog{}).Error
}
