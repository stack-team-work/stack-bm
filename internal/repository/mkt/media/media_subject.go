package media

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/mkt/media"
	"stack-bm/pkg/dict"

	"gorm.io/gorm"
)

type MediaSubjectRepository struct {
	db *gorm.DB
}

func NewMediaSubjectRepository() *MediaSubjectRepository {
	return &MediaSubjectRepository{db: database.DBMkt}
}

func (r *MediaSubjectRepository) Create(m *media.MediaSubject) error { return r.db.Create(m).Error }

func (r *MediaSubjectRepository) FindByID(id uint) (*media.MediaSubject, error) {
	var m media.MediaSubject
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&m).Error
	if err != nil { return nil, err }
	return &m, nil
}

func (r *MediaSubjectRepository) FindPage(page, size int, keyword string, status int) ([]media.MediaSubject, int64, error) {
	var list []media.MediaSubject
	var total int64
	query := r.db.Model(&media.MediaSubject{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if status >= 0 { query = query.Where("status = ?", status) }
	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&list).Error; err != nil { return nil, 0, err }
	return list, total, nil
}

func (r *MediaSubjectRepository) FindAll() ([]media.MediaSubject, error) {
	var list []media.MediaSubject
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&list).Error
	return list, err
}

func (r *MediaSubjectRepository) Update(m *media.MediaSubject) error { return r.db.Save(m).Error }

func (r *MediaSubjectRepository) Delete(id uint) error {
	return r.db.Model(&media.MediaSubject{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

func (r *MediaSubjectRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&media.MediaSubject{}).Select("name as label, id as value").Where("is_deleted = 0 AND status = 1").Order("id ASC").Find(&list).Error
	return list, err
}
