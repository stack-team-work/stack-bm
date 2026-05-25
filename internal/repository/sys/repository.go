package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sys"

	"gorm.io/gorm"
)

type SysAdminRepository struct {
	db *gorm.DB
}

func NewSysAdminRepository() *SysAdminRepository {
	return &SysAdminRepository{db: database.DBBM}
}

func (r *SysAdminRepository) Create(admin *sys.SysAdmin) error {
	return r.db.Create(admin).Error
}

func (r *SysAdminRepository) FindByID(id uint) (*sys.SysAdmin, error) {
	var admin sys.SysAdmin
	err := r.db.First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *SysAdminRepository) FindByUsername(username string) (*sys.SysAdmin, error) {
	var admin sys.SysAdmin
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *SysAdminRepository) FindPage(page, size int, keyword string, groupID int) ([]sys.SysAdmin, int64, error) {
	var admins []sys.SysAdmin
	var total int64

	query := r.db.Model(&sys.SysAdmin{})
	if keyword != "" {
		query = query.Where("username LIKE ? OR name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if groupID > 0 {
		query = query.Where("group_id = ?", groupID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&admins).Error; err != nil {
		return nil, 0, err
	}

	return admins, total, nil
}

func (r *SysAdminRepository) Update(admin *sys.SysAdmin) error {
	return r.db.Save(admin).Error
}

func (r *SysAdminRepository) Delete(id uint) error {
	return r.db.Delete(&sys.SysAdmin{}, id).Error
}

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
	err := r.db.First(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *SysAdminGroupRepository) FindPage(page, size int, keyword string) ([]sys.SysAdminGroup, int64, error) {
	var groups []sys.SysAdminGroup
	var total int64

	query := r.db.Model(&sys.SysAdminGroup{})
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
	err := r.db.Where("status = ?", 1).Find(&groups).Error
	return groups, err
}

func (r *SysAdminGroupRepository) Update(group *sys.SysAdminGroup) error {
	return r.db.Save(group).Error
}

func (r *SysAdminGroupRepository) Delete(id uint) error {
	return r.db.Delete(&sys.SysAdminGroup{}, id).Error
}
