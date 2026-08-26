package media

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/mkt/media"
	"stack-bm/pkg/dict"

	"gorm.io/gorm"
)

type MediaAgentRepository struct {
	db *gorm.DB
}

func NewMediaAgentRepository() *MediaAgentRepository {
	return &MediaAgentRepository{db: database.DBMkt}
}

func (r *MediaAgentRepository) Create(m *media.MediaAgent) error { return r.db.Create(m).Error }

func (r *MediaAgentRepository) FindByID(id uint) (*media.MediaAgent, error) {
	var m media.MediaAgent
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&m).Error
	if err != nil { return nil, err }
	return &m, nil
}

func (r *MediaAgentRepository) FindPage(page, size int, keyword string, status int, subjectID int) ([]media.MediaAgent, int64, error) {
	var list []media.MediaAgent
	var total int64
	query := r.db.Model(&media.MediaAgent{}).Where("is_deleted = 0")
	if keyword != "" { query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%") }
	if status >= 0 { query = query.Where("status = ?", status) }
	if subjectID > 0 { query = query.Where("subject_id = ?", subjectID) }
	if err := query.Count(&total).Error; err != nil { return nil, 0, err }
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&list).Error; err != nil { return nil, 0, err }
	return list, total, nil
}

func (r *MediaAgentRepository) FindAll() ([]media.MediaAgent, error) {
	var list []media.MediaAgent
	err := r.db.Where("status = 1 AND is_deleted = 0").Find(&list).Error
	return list, err
}

func (r *MediaAgentRepository) Update(m *media.MediaAgent) error { return r.db.Save(m).Error }

func (r *MediaAgentRepository) Delete(id uint) error {
	return r.db.Model(&media.MediaAgent{}).Where("id = ?", id).Update("is_deleted", 1).Error
}

func (r *MediaAgentRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&media.MediaAgent{}).Select("name as label, id as value").Where("is_deleted = 0 AND status = 1").Order("id ASC").Find(&list).Error
	return list, err
}
