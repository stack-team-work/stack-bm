package media

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/mkt/media"

	"gorm.io/gorm"
)

type MediaSubRepository struct {
	db *gorm.DB
}

func NewMediaSubRepository() *MediaSubRepository {
	return &MediaSubRepository{db: database.DBMkt}
}

func (r *MediaSubRepository) Create(s *media.MediaSub) error {
	return r.db.Create(s).Error
}

func (r *MediaSubRepository) FindByID(id uint) (*media.MediaSub, error) {
	var s media.MediaSub
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&s).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *MediaSubRepository) FindPage(page, size int, keyword string, mediaID int, status int) ([]media.MediaSub, int64, error) {
	var list []media.MediaSub
	var total int64

	query := r.db.Model(&media.MediaSub{}).Where("is_deleted = 0")
	if keyword != "" {
		query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if mediaID > 0 {
		query = query.Where("media_id = ?", mediaID)
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

func (r *MediaSubRepository) FindAll() ([]media.MediaSub, error) {
	var list []media.MediaSub
	err := r.db.Where("is_deleted = 0").Order("id DESC").Find(&list).Error
	return list, err
}

func (r *MediaSubRepository) Update(s *media.MediaSub) error {
	return r.db.Save(s).Error
}

func (r *MediaSubRepository) Delete(id uint) error {
	return r.db.Model(&media.MediaSub{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
