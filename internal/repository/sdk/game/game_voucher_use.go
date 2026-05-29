package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"
	"gorm.io/gorm"
)

type GameVoucherUseRepository struct{ db *gorm.DB }

func NewGameVoucherUseRepository() *GameVoucherUseRepository {
	return &GameVoucherUseRepository{db: database.DBSdk}
}

func (r *GameVoucherUseRepository) FindPage(page, size int, keyword string, voucherID int) ([]game.GameVoucherUse, int64, error) {
	var list []game.GameVoucherUse
	var total int64
	query := r.db.Model(&game.GameVoucherUse{})
	if keyword != "" {
		query = query.Where("user_id LIKE ? OR role_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if voucherID > 0 {
		query = query.Where("voucher_id = ?", voucherID)
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
