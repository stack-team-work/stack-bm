package pay

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/pay"

	"gorm.io/gorm"
)

type PayMerchantRepository struct {
	db *gorm.DB
}

func NewPayMerchantRepository() *PayMerchantRepository {
	return &PayMerchantRepository{db: database.DBSdk}
}

func (r *PayMerchantRepository) Create(p *pay.PayMerchant) error { return r.db.Create(p).Error }

func (r *PayMerchantRepository) FindByID(id uint) (*pay.PayMerchant, error) {
	var p pay.PayMerchant
	err := r.db.First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PayMerchantRepository) FindPage(page, size int, keyword string, status int, payType int) ([]pay.PayMerchant, int64, error) {
	var list []pay.PayMerchant
	var total int64
	query := r.db.Model(&pay.PayMerchant{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR show_name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if payType > 0 {
		query = query.Where("type = ?", payType)
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

func (r *PayMerchantRepository) FindAll() ([]pay.PayMerchant, error) {
	var list []pay.PayMerchant
	err := r.db.Where("status = 1").Find(&list).Error
	return list, err
}

func (r *PayMerchantRepository) Update(p *pay.PayMerchant) error { return r.db.Save(p).Error }

func (r *PayMerchantRepository) Delete(id uint) error {
	return r.db.Delete(&pay.PayMerchant{}, id).Error
}
