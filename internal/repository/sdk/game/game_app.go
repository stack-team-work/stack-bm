package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"

	"gorm.io/gorm"
)

type GameAppRepository struct {
	db *gorm.DB
}

func NewGameAppRepository() *GameAppRepository {
	return &GameAppRepository{db: database.DBSdk}
}

func (r *GameAppRepository) Create(app *game.GameApp) error {
	return r.db.Create(app).Error
}

func (r *GameAppRepository) FindByID(id uint) (*game.GameApp, error) {
	var app game.GameApp
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&app).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *GameAppRepository) FindPage(page, size int, keyword string, gameID int, status int) ([]game.GameApp, int64, error) {
	var apps []game.GameApp
	var total int64

	query := r.db.Model(&game.GameApp{}).Where("is_deleted = 0")
	if keyword != "" {
		query = query.Where("name LIKE ? OR package_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if gameID > 0 {
		query = query.Where("pid = ?", gameID)
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&apps).Error; err != nil {
		return nil, 0, err
	}
	return apps, total, nil
}

func (r *GameAppRepository) FindAll() ([]game.GameApp, error) {
	var apps []game.GameApp
	err := r.db.Where("is_deleted = 0").Order("id DESC").Find(&apps).Error
	return apps, err
}

func (r *GameAppRepository) Update(app *game.GameApp) error {
	return r.db.Save(app).Error
}

func (r *GameAppRepository) Delete(id uint) error {
	return r.db.Model(&game.GameApp{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
