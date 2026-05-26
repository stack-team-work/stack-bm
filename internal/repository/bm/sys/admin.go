package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"

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
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *SysAdminRepository) FindByUsername(username string) (*sys.SysAdmin, error) {
	var admin sys.SysAdmin
	err := r.db.Where("username = ? AND is_deleted = 0", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *SysAdminRepository) FindPage(page, size int, keyword string, groupID int) ([]sys.SysAdmin, int64, error) {
	var admins []sys.SysAdmin
	var total int64
	query := r.db.Model(&sys.SysAdmin{}).Where("is_deleted = 0")
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
	return r.db.Model(&sys.SysAdmin{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
