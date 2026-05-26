package game

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/sdk/game"

	"gorm.io/gorm"
)

type GameTagRepository struct {
	db *gorm.DB
}

func NewGameTagRepository() *GameTagRepository {
	return &GameTagRepository{db: database.DBSdk}
}

func (r *GameTagRepository) Create(tag *game.GameTag) error {
	return r.db.Create(tag).Error
}

func (r *GameTagRepository) FindByID(id uint) (*game.GameTag, error) {
	var tag game.GameTag
	err := r.db.Where("id = ? AND is_deleted = 0", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *GameTagRepository) FindPage(page, size int, keyword string, tagType int, status int) ([]game.GameTag, int64, error) {
	var tags []game.GameTag
	var total int64

	query := r.db.Model(&game.GameTag{}).Where("is_deleted = 0")
	if keyword != "" {
		query = query.Where("name LIKE ? OR mark LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if tagType > 0 {
		query = query.Where("type = ?", tagType)
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("id DESC").Find(&tags).Error; err != nil {
		return nil, 0, err
	}
	return tags, total, nil
}

func (r *GameTagRepository) FindAllByType(tagType int) ([]game.GameTag, error) {
	var tags []game.GameTag
	err := r.db.Where("type = ? AND status = 1 AND is_deleted = 0", tagType).Order("id DESC").Find(&tags).Error
	return tags, err
}

func (r *GameTagRepository) Update(tag *game.GameTag) error {
	return r.db.Save(tag).Error
}

func (r *GameTagRepository) Delete(id uint) error {
	return r.db.Model(&game.GameTag{}).Where("id = ?", id).Update("is_deleted", 1).Error
}
