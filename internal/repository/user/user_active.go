package user

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/user"

	"gorm.io/gorm"
)

type UserActiveRepository struct {
	db *gorm.DB
}

func NewUserActiveRepository() *UserActiveRepository {
	return &UserActiveRepository{db: database.DBSdk}
}

func (r *UserActiveRepository) FindPage(page, size int, appID int, clientID string, startAt, endAt int64) ([]user.UserActive, int64, error) {
	var list []user.UserActive
	var total int64
	query := r.db.Model(&user.UserActive{})
	if appID > 0 {
		query = query.Where("app_id = ?", appID)
	}
	if clientID != "" {
		query = query.Where("client_id LIKE ?", "%"+clientID+"%")
	}
	if startAt > 0 {
		query = query.Where("created_at >= ?", startAt)
	}
	if endAt > 0 {
		query = query.Where("created_at <= ?", endAt)
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