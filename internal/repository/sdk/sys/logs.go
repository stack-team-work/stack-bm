package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/sys"

	"gorm.io/gorm"
)

type SysLogRepository struct {
	db *gorm.DB
}

func NewSysLogRepository() *SysLogRepository {
	return &SysLogRepository{db: database.DBSdk}
}

func (r *SysLogRepository) FindPage(page, size int, keyword string, level int, type_ int) ([]sys.SysLog, int64, error) {
	var logs []sys.SysLog
	var total int64
	query := r.db.Model(&sys.SysLog{})
	if keyword != "" {
		query = query.Where("ip LIKE ? OR `desc` LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if level > 0 {
		query = query.Where("level = ?", level)
	}
	if type_ > 0 {
		query = query.Where("type = ?", type_)
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
