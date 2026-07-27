package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"
	"gorm.io/gorm"
)

type FeishuUserRepository struct{ db *gorm.DB }

func NewFeishuUserRepository() *FeishuUserRepository { return &FeishuUserRepository{db: database.DBBM} }

func (r *FeishuUserRepository) Create(m *sys.FeishuUser) error { return r.db.Create(m).Error }

func (r *FeishuUserRepository) FindByID(id uint) (*sys.FeishuUser, error) {
	var m sys.FeishuUser
	err := r.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *FeishuUserRepository) FindPage(page, size int, keyword string, status int, adminID int) ([]sys.FeishuUser, int64, error) {
	var list []sys.FeishuUser
	var total int64
	query := r.db.Model(&sys.FeishuUser{})
	if keyword != "" {
		query = query.Where("feishu_user_id LIKE ?", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if adminID > 0 {
		query = query.Where("admin_id = ?", adminID)
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

func (r *FeishuUserRepository) FindAll() ([]sys.FeishuUser, error) {
	var list []sys.FeishuUser
	err := r.db.Where("status = 1").Find(&list).Error
	return list, err
}

func (r *FeishuUserRepository) Update(m *sys.FeishuUser) error { return r.db.Save(m).Error }
func (r *FeishuUserRepository) Delete(id uint) error          { return r.db.Delete(&sys.FeishuUser{}, id).Error }
