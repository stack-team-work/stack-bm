package media

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/mkt/media"
	"stack-bm/pkg/dict"

	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository() *MediaRepository {
	return &MediaRepository{db: database.DBMkt}
}

func (r *MediaRepository) Create(m *media.Media) error {
	return r.db.Create(m).Error
}

func (r *MediaRepository) FindByID(id uint) (*media.Media, error) {
	var m media.Media
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MediaRepository) FindByMark(mark string) (*media.Media, error) {
	var m media.Media
	err := r.db.Where("mark = ? AND is_deleted = 0", mark).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MediaRepository) FindPage(page, size int, keyword string, status int) ([]media.Media, int64, error) {
	var list []media.Media
	var total int64

	query := r.db.Model(&media.Media{}).Where("is_deleted = 0")
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
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *MediaRepository) FindAll() ([]media.Media, error) {
	var list []media.Media
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&list).Error
	return list, err
}

func (r *MediaRepository) Update(m *media.Media) error {
	return r.db.Save(m).Error
}

func (r *MediaRepository) Delete(id uint) error {
	return r.db.Model(&media.Media{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

func (r *MediaRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&media.Media{}).Select("name as label, id as value").Where("is_deleted = 0 AND status = 1").Order("id ASC").Find(&list).Error
	return list, err
}
