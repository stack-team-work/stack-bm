package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"

	"gorm.io/gorm"
)

type SysColumnRepository struct {
	db *gorm.DB
}

func NewSysColumnRepository() *SysColumnRepository {
	return &SysColumnRepository{db: database.DBBM}
}

func (r *SysColumnRepository) Create(c *sys.SysColumn) error {
	return r.db.Create(c).Error
}

func (r *SysColumnRepository) FindByID(id uint) (*sys.SysColumn, error) {
	var c sys.SysColumn
	err := r.db.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SysColumnRepository) FindPage(page, size int, keyword string, status int, reportType int, indicatorType int) ([]sys.SysColumn, int64, error) {
	var list []sys.SysColumn
	var total int64
	query := r.db.Model(&sys.SysColumn{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR field LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if reportType > 0 {
		query = query.Where("report_type = ?", reportType)
	}
	if indicatorType > 0 {
		query = query.Where("indicator_type = ?", indicatorType)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *SysColumnRepository) FindAll() ([]sys.SysColumn, error) {
	var list []sys.SysColumn
	err := r.db.Where("status = 1").Find(&list).Error
	return list, err
}

func (r *SysColumnRepository) Update(c *sys.SysColumn) error {
	return r.db.Save(c).Error
}

func (r *SysColumnRepository) Delete(id uint) error {
	return r.db.Delete(&sys.SysColumn{}, id).Error
}
