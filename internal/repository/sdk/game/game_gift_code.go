package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"
	"gorm.io/gorm"
)

type GameGiftCodeRepository struct{ db *gorm.DB }

func NewGameGiftCodeRepository() *GameGiftCodeRepository {
	return &GameGiftCodeRepository{db: database.DBSdk}
}

func (r *GameGiftCodeRepository) Create(c *game.GameGiftCode) error { return r.db.Create(c).Error }

func (r *GameGiftCodeRepository) FindByID(id uint) (*game.GameGiftCode, error) {
	var c game.GameGiftCode
	err := r.db.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *GameGiftCodeRepository) FindPage(page, size int, keyword string, status int, giftID int) ([]game.GameGiftCode, int64, error) {
	var list []game.GameGiftCode
	var total int64
	query := r.db.Model(&game.GameGiftCode{})
	if keyword != "" {
		query = query.Where("code LIKE ?", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
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

func (r *GameGiftCodeRepository) Update(c *game.GameGiftCode) error { return r.db.Save(c).Error }

func (r *GameGiftCodeRepository) Delete(id uint) error { return r.db.Delete(&game.GameGiftCode{}, id).Error }
