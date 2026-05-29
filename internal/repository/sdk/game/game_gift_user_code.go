package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"
	"gorm.io/gorm"
)

type GameGiftUserCodeRepository struct{ db *gorm.DB }

func NewGameGiftUserCodeRepository() *GameGiftUserCodeRepository {
	return &GameGiftUserCodeRepository{db: database.DBSdk}
}

func (r *GameGiftUserCodeRepository) FindPage(page, size int, keyword string, giftID int) ([]game.GameGiftUserCode, int64, error) {
	var list []game.GameGiftUserCode
	var total int64
	query := r.db.Model(&game.GameGiftUserCode{})
	if keyword != "" {
		query = query.Where("code LIKE ? OR user_id LIKE ? OR role_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if giftID > 0 {
		query = query.Where("gift_id = ?", giftID)
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
