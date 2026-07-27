package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"
	"stack-bm/pkg/dict"

	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository() *GameRepository {
	return &GameRepository{db: database.DBSdk}
}

func (r *GameRepository) Create(g *game.Game) error {
	return r.db.Create(g).Error
}

func (r *GameRepository) FindByID(id uint) (*game.Game, error) {
	var g game.Game
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&g).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *GameRepository) FindPage(page, size int, keyword string, status int) ([]game.Game, int64, error) {
	var games []game.Game
	var total int64

	query := r.db.Model(&game.Game{}).Where("is_deleted = 0")
	if keyword != "" {
		query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&games).Error; err != nil {
		return nil, 0, err
	}
	return games, total, nil
}

func (r *GameRepository) FindAll() ([]game.Game, error) {
	var games []game.Game
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&games).Error
	return games, err
}

func (r *GameRepository) Update(g *game.Game) error {
	return r.db.Save(g).Error
}

func (r *GameRepository) Delete(id uint) error {
	return r.db.Model(&game.Game{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

func (r *GameRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&game.Game{}).Select("name as label, id as value").Where("is_deleted = 0 AND status = 1").Order("id ASC").Find(&list).Error
	return list, err
}
