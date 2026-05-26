package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"

	"gorm.io/gorm"
)

type GameCpRepository struct {
	db *gorm.DB
}

func NewGameCpRepository() *GameCpRepository {
	return &GameCpRepository{db: database.DBSdk}
}

func (r *GameCpRepository) Create(cp *game.GameCp) error {
	return r.db.Create(cp).Error
}

func (r *GameCpRepository) FindByID(id uint) (*game.GameCp, error) {
	var cp game.GameCp
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&cp).Error
	if err != nil {
		return nil, err
	}
	return &cp, nil
}

func (r *GameCpRepository) FindPage(page, size int, keyword string, status int) ([]game.GameCp, int64, error) {
	var cps []game.GameCp
	var total int64

	query := r.db.Model(&game.GameCp{}).Where("is_deleted = 0")
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
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&cps).Error; err != nil {
		return nil, 0, err
	}
	return cps, total, nil
}

func (r *GameCpRepository) FindAll() ([]game.GameCp, error) {
	var cps []game.GameCp
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&cps).Error
	return cps, err
}

func (r *GameCpRepository) Update(cp *game.GameCp) error {
	return r.db.Save(cp).Error
}

func (r *GameCpRepository) Delete(id uint) error {
	return r.db.Model(&game.GameCp{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
