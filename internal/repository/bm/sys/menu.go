package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"

	"gorm.io/gorm"
)

type SysMenuRepository struct {
	db *gorm.DB
}

func NewSysMenuRepository() *SysMenuRepository {
	return &SysMenuRepository{db: database.DBBM}
}

func (r *SysMenuRepository) Create(m *sys.SysMenu) error {
	return r.db.Create(m).Error
}

func (r *SysMenuRepository) FindByID(id uint) (*sys.SysMenu, error) {
	var m sys.SysMenu
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *SysMenuRepository) FindPage(page, size int) ([]sys.SysMenu, int64, error) {
	var menus []sys.SysMenu
	var total int64
	query := r.db.Model(&sys.SysMenu{}).Where("is_deleted = 0")
	query.Count(&total)
	offset := (page - 1) * size
	err := query.Offset(offset).Limit(size).Order("sort ASC, id ASC").Find(&menus).Error
	return menus, total, err
}

func (r *SysMenuRepository) FindAll() ([]sys.SysMenu, error) {
	var menus []sys.SysMenu
	err := r.db.Where("is_deleted = 0").Order("sort ASC, id ASC").Find(&menus).Error
	return menus, err
}

func (r *SysMenuRepository) Update(m *sys.SysMenu) error {
	return r.db.Save(m).Error
}

func (r *SysMenuRepository) Delete(id uint) error {
	return r.db.Model(&sys.SysMenu{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
