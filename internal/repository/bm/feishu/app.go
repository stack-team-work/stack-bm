package feishu

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/feishu"
	"stack-bm/pkg/dict"
	"gorm.io/gorm"
)

type FeishuAppRepository struct{ db *gorm.DB }

func NewFeishuAppRepository() *FeishuAppRepository { return &FeishuAppRepository{db: database.DBBM} }

func (r *FeishuAppRepository) Create(m *feishu.FeishuApp) error { return r.db.Create(m).Error }

func (r *FeishuAppRepository) FindByID(id uint) (*feishu.FeishuApp, error) {
	var m feishu.FeishuApp
	err := r.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *FeishuAppRepository) FindPage(page, size int, keyword string, status int) ([]feishu.FeishuApp, int64, error) {
	var list []feishu.FeishuApp
	var total int64
	query := r.db.Model(&feishu.FeishuApp{})
	if keyword != "" {
		query = query.Where("app_name LIKE ? OR app_id LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
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

func (r *FeishuAppRepository) FindAll() ([]feishu.FeishuApp, error) {
	var list []feishu.FeishuApp
	err := r.db.Where("status = 1").Find(&list).Error
	return list, err
}

func (r *FeishuAppRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&feishu.FeishuApp{}).Select("app_name as label, id as value").Where("status = 1").Order("id ASC").Find(&list).Error
	return list, err
}

func (r *FeishuAppRepository) Update(m *feishu.FeishuApp) error { return r.db.Save(m).Error }
func (r *FeishuAppRepository) Delete(id uint) error          { return r.db.Delete(&feishu.FeishuApp{}, id).Error }
