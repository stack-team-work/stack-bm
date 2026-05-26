package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"

	"gorm.io/gorm"
)

type GamePlatformRepository struct {
	db *gorm.DB
}

func NewGamePlatformRepository() *GamePlatformRepository {
	return &GamePlatformRepository{db: database.DBSdk}
}

func (r *GamePlatformRepository) Create(p *game.GamePlatform) error {
	return r.db.Create(p).Error
}

func (r *GamePlatformRepository) FindByID(id uint) (*game.GamePlatform, error) {
	var p game.GamePlatform
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *GamePlatformRepository) FindPage(page, size int, keyword string, status int) ([]game.GamePlatform, int64, error) {
	var platforms []game.GamePlatform
	var total int64

	query := r.db.Model(&game.GamePlatform{}).Where("is_deleted = 0")
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
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&platforms).Error; err != nil {
		return nil, 0, err
	}
	return platforms, total, nil
}

func (r *GamePlatformRepository) FindAll() ([]game.GamePlatform, error) {
	var platforms []game.GamePlatform
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&platforms).Error
	return platforms, err
}

func (r *GamePlatformRepository) Update(p *game.GamePlatform) error {
	return r.db.Save(p).Error
}

func (r *GamePlatformRepository) Delete(id uint) error {
	return r.db.Model(&game.GamePlatform{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
