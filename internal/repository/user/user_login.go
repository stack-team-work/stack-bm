package user

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/user"

	"gorm.io/gorm"
)

type UserLoginRepository struct {
	db *gorm.DB
}

func NewUserLoginRepository() *UserLoginRepository {
	return &UserLoginRepository{db: database.DBSdk}
}

func (r *UserLoginRepository) FindPage(page, size int, appID int, userID string, startAt, endAt int64) ([]user.UserLogin, int64, error) {
	var list []user.UserLogin
	var total int64
	query := r.db.Model(&user.UserLogin{})
	if appID > 0 {
		query = query.Where("app_id = ?", appID)
	}
	if userID != "" {
		query = query.Where("user_id LIKE ?", "%"+userID+"%")
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