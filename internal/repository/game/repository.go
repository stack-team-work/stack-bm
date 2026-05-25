package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/game"

	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository() *GameRepository {
	return &GameRepository{db: database.DBApi}
}

func (r *GameRepository) Create(g *game.Game) error {
	return r.db.Create(g).Error
}

func (r *GameRepository) FindByID(id uint) (*game.Game, error) {
	var g game.Game
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&g).Error
	if err != nil { return nil, err }
	return &g, nil
}

func (r *GameRepository) FindPage(page, size int, keyword string, status int) ([]game.Game, int64, error) {
	var games []game.Game
	var total int64

	query := r.db.Model(&game.Game{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if status >= 0 { query = query.Where("status = ?", status) }

	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&games).Error; err != nil { return nil, 0, err }
	return games, total, nil
}

func (r *GameRepository) FindAll() ([]game.Game, error) {
	var games []game.Game
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&games).Error
	return games, err
}

func (r *GameRepository) Update(g *game.Game) error { return r.db.Save(g).Error }

func (r *GameRepository) Delete(id uint) error {
	return r.db.Model(&game.Game{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

type GameAppRepository struct {
	db *gorm.DB
}

func NewGameAppRepository() *GameAppRepository {
	return &GameAppRepository{db: database.DBApi}
}

func (r *GameAppRepository) Create(app *game.GameApp) error { return r.db.Create(app).Error }

func (r *GameAppRepository) FindByID(id uint) (*game.GameApp, error) {
	var app game.GameApp
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&app).Error
	if err != nil { return nil, err }
	return &app, nil
}

func (r *GameAppRepository) FindPage(page, size int, keyword string, gameID int, status int) ([]game.GameApp, int64, error) {
	var apps []game.GameApp
	var total int64

	query := r.db.Model(&game.GameApp{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ? OR package_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if gameID > 0 { query = query.Where("pid = ?", gameID) }
	if status >= 0 { query = query.Where("status = ?", status) }

	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&apps).Error; err != nil { return nil, 0, err }
	return apps, total, nil
}

func (r *GameAppRepository) Update(app *game.GameApp) error { return r.db.Save(app).Error }

func (r *GameAppRepository) Delete(id uint) error {
	return r.db.Model(&game.GameApp{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

type GameCpRepository struct {
	db *gorm.DB
}

func NewGameCpRepository() *GameCpRepository {
	return &GameCpRepository{db: database.DBApi}
}

func (r *GameCpRepository) Create(cp *game.GameCp) error { return r.db.Create(cp).Error }

func (r *GameCpRepository) FindByID(id uint) (*game.GameCp, error) {
	var cp game.GameCp
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&cp).Error
	if err != nil { return nil, err }
	return &cp, nil
}

func (r *GameCpRepository) FindPage(page, size int, keyword string, status int) ([]game.GameCp, int64, error) {
	var cps []game.GameCp
	var total int64

	query := r.db.Model(&game.GameCp{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if status >= 0 { query = query.Where("status = ?", status) }

	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&cps).Error; err != nil { return nil, 0, err }
	return cps, total, nil
}

func (r *GameCpRepository) FindAll() ([]game.GameCp, error) {
	var cps []game.GameCp
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&cps).Error
	return cps, err
}

func (r *GameCpRepository) Update(cp *game.GameCp) error { return r.db.Save(cp).Error }

func (r *GameCpRepository) Delete(id uint) error {
	return r.db.Model(&game.GameCp{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
