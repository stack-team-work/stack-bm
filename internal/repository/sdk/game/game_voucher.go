package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"
	"stack-bm/pkg/dict"
	"gorm.io/gorm"
)

type GameVoucherRepository struct{ db *gorm.DB }

func NewGameVoucherRepository() *GameVoucherRepository {
	return &GameVoucherRepository{db: database.DBSdk}
}

func (r *GameVoucherRepository) Create(v *game.GameVoucher) error { return r.db.Create(v).Error }

func (r *GameVoucherRepository) FindByID(id uint) (*game.GameVoucher, error) {
	var v game.GameVoucher
	err := r.db.First(&v, id).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *GameVoucherRepository) FindPage(page, size int, keyword string, status int) ([]game.GameVoucher, int64, error) {
	var list []game.GameVoucher
	var total int64
	query := r.db.Model(&game.GameVoucher{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR `desc` LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
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

func (r *GameVoucherRepository) FindAll() ([]game.GameVoucher, error) {
	var list []game.GameVoucher
	err := r.db.Where("status = 1").Find(&list).Error
	return list, err
}

func (r *GameVoucherRepository) Update(v *game.GameVoucher) error { return r.db.Save(v).Error }

func (r *GameVoucherRepository) Delete(id uint) error { return r.db.Delete(&game.GameVoucher{}, id).Error }

func (r *GameVoucherRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&game.GameVoucher{}).Select("name as label, id as value").Where("status = 1").Order("id ASC").Find(&list).Error
	return list, err
}
