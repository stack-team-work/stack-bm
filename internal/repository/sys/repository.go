package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sys"

	"gorm.io/gorm"
)

type SysAdminRepository struct {
	db *gorm.DB
}

func NewSysAdminRepository() *SysAdminRepository { return &SysAdminRepository{db: database.DBBM} }

func (r *SysAdminRepository) Create(admin *sys.SysAdmin) error { return r.db.Create(admin).Error }

func (r *SysAdminRepository) FindByID(id uint) (*sys.SysAdmin, error) {
	var admin sys.SysAdmin
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&admin).Error
	if err != nil { return nil, err }
	return &admin, nil
}

func (r *SysAdminRepository) FindByUsername(username string) (*sys.SysAdmin, error) {
	var admin sys.SysAdmin
	err := r.db.Where("username = ? AND is_deleted = 0", username).First(&admin).Error
	if err != nil { return nil, err }
	return &admin, nil
}

func (r *SysAdminRepository) FindPage(page, size int, keyword string, groupID int) ([]sys.SysAdmin, int64, error) {
	var admins []sys.SysAdmin
	var total int64
	query := r.db.Model(&sys.SysAdmin{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("username LIKE ? OR name LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if groupID > 0 { query = query.Where("group_id = ?", groupID) }
	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&admins).Error; err != nil { return nil, 0, err }
	return admins, total, nil
}

func (r *SysAdminRepository) Update(admin *sys.SysAdmin) error { return r.db.Save(admin).Error }

func (r *SysAdminRepository) Delete(id uint) error {
	return r.db.Model(&sys.SysAdmin{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

type SysAdminGroupRepository struct {
	db *gorm.DB
}

func NewSysAdminGroupRepository() *SysAdminGroupRepository { return &SysAdminGroupRepository{db: database.DBBM} }

func (r *SysAdminGroupRepository) Create(group *sys.SysAdminGroup) error { return r.db.Create(group).Error }

func (r *SysAdminGroupRepository) FindByID(id uint) (*sys.SysAdminGroup, error) {
	var group sys.SysAdminGroup
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&group).Error
	if err != nil { return nil, err }
	return &group, nil
}

func (r *SysAdminGroupRepository) FindPage(page, size int, keyword string) ([]sys.SysAdminGroup, int64, error) {
	var groups []sys.SysAdminGroup
	var total int64
	query := r.db.Model(&sys.SysAdminGroup{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ?", "%"+keyword+"%") }
	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&groups).Error; err != nil { return nil, 0, err }
	return groups, total, nil
}

func (r *SysAdminGroupRepository) FindAll() ([]sys.SysAdminGroup, error) {
	var groups []sys.SysAdminGroup
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&groups).Error
	return groups, err
}

func (r *SysAdminGroupRepository) Update(group *sys.SysAdminGroup) error { return r.db.Save(group).Error }

func (r *SysAdminGroupRepository) Delete(id uint) error {
	return r.db.Model(&sys.SysAdminGroup{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

type SysLogRepository struct {
	db *gorm.DB
}

func NewSysLogRepository() *SysLogRepository { return &SysLogRepository{db: database.DBBM} }

func (r *SysLogRepository) Create(log *sys.SysLog) error { return r.db.Create(log).Error }

func (r *SysLogRepository) FindPage(page, size int, keyword string, level string) ([]sys.SysLog, int64, error) {
	var logs []sys.SysLog
	var total int64
	query := r.db.Model(&sys.SysLog{})
	if keyword != "" { query = query.Where("path LIKE ? OR username LIKE ? OR `desc` LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%") }
	if level != "" { query = query.Where("level = ?", level) }
	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&logs).Error; err != nil { return nil, 0, err }
	return logs, total, nil
}

func (r *SysLogRepository) ClearAll() error {
	return r.db.Where("1 = 1").Delete(&sys.SysLog{}).Error
}

type SysMenuRepository struct {
	db *gorm.DB
}

func NewSysMenuRepository() *SysMenuRepository { return &SysMenuRepository{db: database.DBBM} }

func (r *SysMenuRepository) Create(m *sys.SysMenu) error { return r.db.Create(m).Error }

func (r *SysMenuRepository) FindByID(id uint) (*sys.SysMenu, error) {
	var m sys.SysMenu
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&m).Error
	if err != nil { return nil, err }
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

func (r *SysMenuRepository) Update(m *sys.SysMenu) error { return r.db.Save(m).Error }

func (r *SysMenuRepository) Delete(id uint) error {
	return r.db.Model(&sys.SysMenu{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
