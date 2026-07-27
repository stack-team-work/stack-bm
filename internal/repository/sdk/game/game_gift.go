package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"
	"stack-bm/pkg/dict"
	"gorm.io/gorm"
)

type GameGiftRepository struct{ db *gorm.DB }

func NewGameGiftRepository() *GameGiftRepository { return &GameGiftRepository{db: database.DBSdk} }

func (r *GameGiftRepository) Create(g *game.GameGift) error { return r.db.Create(g).Error }

func (r *GameGiftRepository) FindByID(id uint) (*game.GameGift, error) {
	var g game.GameGift
	err := r.db.First(&g, id).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *GameGiftRepository) FindPage(page, size int, keyword string, status int) ([]game.GameGift, int64, error) {
	var list []game.GameGift
	var total int64
	query := r.db.Model(&game.GameGift{})
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

func (r *GameGiftRepository) FindAll() ([]game.GameGift, error) {
	var list []game.GameGift
	err := r.db.Where("status = 1").Find(&list).Error
	return list, err
}

func (r *GameGiftRepository) Update(g *game.GameGift) error { return r.db.Save(g).Error }

func (r *GameGiftRepository) Delete(id uint) error { return r.db.Delete(&game.GameGift{}, id).Error }

func (r *GameGiftRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&game.GameGift{}).Select("name as label, id as value").Where("status = 1").Order("id ASC").Find(&list).Error
	return list, err
}
