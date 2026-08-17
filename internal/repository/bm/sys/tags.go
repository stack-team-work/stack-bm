package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"
	"stack-bm/pkg/dict"

	"gorm.io/gorm"
)

type SysTagRepository struct {
	db *gorm.DB
}

func NewSysTagRepository() *SysTagRepository {
	return &SysTagRepository{db: database.DBBM}
}

func (r *SysTagRepository) Create(t *sys.SysTag) error {
	return r.db.Create(t).Error
}

func (r *SysTagRepository) FindByID(id uint) (*sys.SysTag, error) {
	var t sys.SysTag
	err := r.db.First(&t, id).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *SysTagRepository) FindPage(page, size int, keyword string, tagType int, status int) ([]sys.SysTag, int64, error) {
	var list []sys.SysTag
	var total int64

	query := r.db.Model(&sys.SysTag{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR remark LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if tagType > 0 {
		query = query.Where("type = ?", tagType)
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
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

func (r *SysTagRepository) FindAllByType(tagType int) ([]sys.SysTag, error) {
	var list []sys.SysTag
	query := r.db.Where("status = 1")
	if tagType > 0 {
		query = query.Where("type = ?", tagType)
	}
	err := query.Order("id DESC").Find(&list).Error
	return list, err
}

func (r *SysTagRepository) FindOptionsByType(tagType int) ([]dict.Option, error) {
	var list []dict.Option
	query := r.db.Model(&sys.SysTag{}).Select("name as label, id as value").Where("status = 1")
	if tagType > 0 {
		query = query.Where("type = ?", tagType)
	}
	err := query.Order("id ASC").Find(&list).Error
	return list, err
}

func (r *SysTagRepository) Update(t *sys.SysTag) error {
	return r.db.Save(t).Error
}
