package user

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/user"

	"gorm.io/gorm"
)

type UserOrderRepository struct {
	db *gorm.DB
}

func NewUserOrderRepository() *UserOrderRepository {
	return &UserOrderRepository{db: database.DBSdk}
}

func (r *UserOrderRepository) FindPage(page, size int, appID int, userID string, payStatus, status int, startAt, endAt int64) ([]user.UserOrder, int64, error) {
	var list []user.UserOrder
	var total int64
	query := r.db.Model(&user.UserOrder{})
	if appID > 0 {
		query = query.Where("app_id = ?", appID)
	}
	if userID != "" {
		query = query.Where("user_id LIKE ?", "%"+userID+"%")
	}
	if payStatus > 0 {
		query = query.Where("pay_status = ?", payStatus)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
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