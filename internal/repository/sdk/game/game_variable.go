package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"

	"gorm.io/gorm"
)

type GameVariableRepository struct {
	db *gorm.DB
}

func NewGameVariableRepository() *GameVariableRepository {
	return &GameVariableRepository{db: database.DBSdk}
}

func (r *GameVariableRepository) Create(v *game.GameVariable) error {
	return r.db.Create(v).Error
}

func (r *GameVariableRepository) FindByID(id uint) (*game.GameVariable, error) {
	var v game.GameVariable
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *GameVariableRepository) FindPage(page, size int, keyword string, status int) ([]game.GameVariable, int64, error) {
	var vars []game.GameVariable
	var total int64

	query := r.db.Model(&game.GameVariable{}).Where("is_deleted = 0")
	if keyword != "" {
		query = query.Where("name LIKE ? OR `key` LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&vars).Error; err != nil {
		return nil, 0, err
	}
	return vars, total, nil
}

func (r *GameVariableRepository) Update(v *game.GameVariable) error {
	return r.db.Save(v).Error
}

func (r *GameVariableRepository) Delete(id uint) error {
	return r.db.Model(&game.GameVariable{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
