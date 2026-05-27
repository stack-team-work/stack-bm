package pay

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/pay"

	"gorm.io/gorm"
)

type PayPlatformRepository struct {
	db *gorm.DB
}

func NewPayPlatformRepository() *PayPlatformRepository {
	return &PayPlatformRepository{db: database.DBSdk}
}

func (r *PayPlatformRepository) Create(p *pay.PayPlatform) error { return r.db.Create(p).Error }

func (r *PayPlatformRepository) FindByID(id uint) (*pay.PayPlatform, error) {
	var p pay.PayPlatform
	err := r.db.First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PayPlatformRepository) FindPage(page, size int, keyword string) ([]pay.PayPlatform, int64, error) {
	var list []pay.PayPlatform
	var total int64
	query := r.db.Model(&pay.PayPlatform{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
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

func (r *PayPlatformRepository) FindAll() ([]pay.PayPlatform, error) {
	var list []pay.PayPlatform
	err := r.db.Find(&list).Error
	return list, err
}

func (r *PayPlatformRepository) Update(p *pay.PayPlatform) error { return r.db.Save(p).Error }

func (r *PayPlatformRepository) Delete(id uint) error {
	return r.db.Delete(&pay.PayPlatform{}, id).Error
}
