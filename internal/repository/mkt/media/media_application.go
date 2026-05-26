package media

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/mkt/media"

	"gorm.io/gorm"
)

type MediaApplicationRepository struct {
	db *gorm.DB
}

func NewMediaApplicationRepository() *MediaApplicationRepository {
	return &MediaApplicationRepository{db: database.DBMkt}
}

func (r *MediaApplicationRepository) Create(m *media.MediaApplication) error { return r.db.Create(m).Error }

func (r *MediaApplicationRepository) FindByID(id uint) (*media.MediaApplication, error) {
	var m media.MediaApplication
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&m).Error
	if err != nil { return nil, err }
	return &m, nil
}

func (r *MediaApplicationRepository) FindPage(page, size int, keyword string, status int, mediaID int) ([]media.MediaApplication, int64, error) {
	var list []media.MediaApplication
	var total int64
	query := r.db.Model(&media.MediaApplication{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ? OR remark LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if status >= 0 { query = query.Where("status = ?", status) }
	if mediaID > 0 { query = query.Where("media_id = ?", mediaID) }
	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&list).Error; err != nil { return nil, 0, err }
	return list, total, nil
}

func (r *MediaApplicationRepository) FindAll() ([]media.MediaApplication, error) {
	var list []media.MediaApplication
	err := r.db.Where("is_deleted = 0").Order("id DESC").Find(&list).Error
	return list, err
}

func (r *MediaApplicationRepository) Update(m *media.MediaApplication) error { return r.db.Save(m).Error }

func (r *MediaApplicationRepository) Delete(id uint) error {
	return r.db.Model(&media.MediaApplication{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
