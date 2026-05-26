package media

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/mkt/media"

	"gorm.io/gorm"
)

type MediaManagerRepository struct {
	db *gorm.DB
}

func NewMediaManagerRepository() *MediaManagerRepository {
	return &MediaManagerRepository{db: database.DBMkt}
}

func (r *MediaManagerRepository) Create(m *media.MediaManager) error { return r.db.Create(m).Error }

func (r *MediaManagerRepository) FindByID(id uint) (*media.MediaManager, error) {
	var m media.MediaManager
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&m).Error
	if err != nil { return nil, err }
	return &m, nil
}

func (r *MediaManagerRepository) FindPage(page, size int, keyword string, status int, mediaID int, applicationID int) ([]media.MediaManager, int64, error) {
	var list []media.MediaManager
	var total int64
	query := r.db.Model(&media.MediaManager{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ? OR account LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if status >= 0 { query = query.Where("status = ?", status) }
	if mediaID > 0 { query = query.Where("media_id = ?", mediaID) }
	if applicationID > 0 { query = query.Where("application_id = ?", applicationID) }
	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&list).Error; err != nil { return nil, 0, err }
	return list, total, nil
}

func (r *MediaManagerRepository) FindAll() ([]media.MediaManager, error) {
	var list []media.MediaManager
	err := r.db.Where("is_deleted = 0").Order("id DESC").Find(&list).Error
	return list, err
}

func (r *MediaManagerRepository) Update(m *media.MediaManager) error { return r.db.Save(m).Error }

func (r *MediaManagerRepository) Delete(id uint) error {
	return r.db.Model(&media.MediaManager{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
