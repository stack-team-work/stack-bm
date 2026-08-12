package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"
	"stack-bm/pkg/dict"

	"gorm.io/gorm"
)

type GameAppTemplateRepository struct {
	db *gorm.DB
}

func NewGameAppTemplateRepository() *GameAppTemplateRepository {
	return &GameAppTemplateRepository{db: database.DBSdk}
}

func (r *GameAppTemplateRepository) Create(t *game.GameAppTemplate) error {
	return r.db.Create(t).Error
}

func (r *GameAppTemplateRepository) FindByID(id uint) (*game.GameAppTemplate, error) {
	var t game.GameAppTemplate
	if err := r.db.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *GameAppTemplateRepository) FindPage(page, size int, keyword string, status int) ([]game.GameAppTemplate, int64, error) {
	var list []game.GameAppTemplate
	var total int64

	query := r.db.Model(&game.GameAppTemplate{})
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
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

func (r *GameAppTemplateRepository) FindAll() ([]game.GameAppTemplate, error) {
	var list []game.GameAppTemplate
	err := r.db.Where("status = 1").Order("id DESC").Find(&list).Error
	return list, err
}

func (r *GameAppTemplateRepository) Update(t *game.GameAppTemplate) error {
	return r.db.Save(t).Error
}

func (r *GameAppTemplateRepository) Delete(id uint) error {
	return r.db.Delete(&game.GameAppTemplate{}, id).Error
}

func (r *GameAppTemplateRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&game.GameAppTemplate{}).
		Select("name as label, id as value").
		Where("status = 1").Order("id ASC").Scan(&list).Error
	return list, err
}
