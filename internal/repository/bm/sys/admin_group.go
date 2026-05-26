package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"

	"gorm.io/gorm"
)

type SysAdminGroupRepository struct {
	db *gorm.DB
}

func NewSysAdminGroupRepository() *SysAdminGroupRepository {
	return &SysAdminGroupRepository{db: database.DBBM}
}

func (r *SysAdminGroupRepository) Create(group *sys.SysAdminGroup) error {
	return r.db.Create(group).Error
}

func (r *SysAdminGroupRepository) FindByID(id uint) (*sys.SysAdminGroup, error) {
	var group sys.SysAdminGroup
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *SysAdminGroupRepository) FindPage(page, size int, keyword string) ([]sys.SysAdminGroup, int64, error) {
	var groups []sys.SysAdminGroup
	var total int64
	query := r.db.Model(&sys.SysAdminGroup{}).Where("is_deleted = 0")
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&groups).Error; err != nil {
		return nil, 0, err
	}
	return groups, total, nil
}

func (r *SysAdminGroupRepository) FindAll() ([]sys.SysAdminGroup, error) {
	var groups []sys.SysAdminGroup
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&groups).Error
	return groups, err
}

func (r *SysAdminGroupRepository) Update(group *sys.SysAdminGroup) error {
	return r.db.Save(group).Error
}

func (r *SysAdminGroupRepository) Delete(id uint) error {
	return r.db.Model(&sys.SysAdminGroup{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
