package media

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/mkt/media"
	"stack-bm/pkg/dict"

	"gorm.io/gorm"
)

type MediaAccountRepository struct {
	db *gorm.DB
}

func NewMediaAccountRepository() *MediaAccountRepository {
	return &MediaAccountRepository{db: database.DBMkt}
}

func (r *MediaAccountRepository) Create(m *media.MediaAccount) error { return r.db.Create(m).Error }

func (r *MediaAccountRepository) FindByID(id uint) (*media.MediaAccount, error) {
	var m media.MediaAccount
	if err := r.db.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MediaAccountRepository) FindByUID(uid string) (*media.MediaAccount, error) {
	var m media.MediaAccount
	if err := r.db.Where("uid = ?", uid).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MediaAccountRepository) FindPage(page, size int, keyword string, status int, mediaSubID int, subjectID int, adminID int) ([]media.MediaAccount, int64, error) {
	var list []media.MediaAccount
	var total int64
	query := r.db.Model(&media.MediaAccount{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR username LIKE ? OR uid LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if mediaSubID > 0 {
		query = query.Where("media_sub_id = ?", mediaSubID)
	}
	if subjectID > 0 {
		query = query.Where("subject_id = ?", subjectID)
	}
	if adminID > 0 {
		query = query.Where("admin_id = ?", adminID)
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

func (r *MediaAccountRepository) FindAll() ([]media.MediaAccount, error) {
	var list []media.MediaAccount
	err := r.db.Order("id DESC").Find(&list).Error
	return list, err
}

func (r *MediaAccountRepository) FindAllByChannel(mediaID int) ([]media.MediaAccount, error) {
	var list []media.MediaAccount
	// 渠道账户通过 media_sub 关联媒体渠道
	query := r.db.Model(&media.MediaAccount{}).
		Joins("JOIN media_sub ON media_sub.id = media_accounts.media_sub_id").
		Where("media_sub.media_id = ? AND media_accounts.status = 1", mediaID)
	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *MediaAccountRepository) Update(m *media.MediaAccount) error { return r.db.Save(m).Error }

func (r *MediaAccountRepository) UpdateBalance(id uint, balance float64) error {
	return r.db.Model(&media.MediaAccount{}).Where("id = ?", id).Update("balance", balance).Error
}

func (r *MediaAccountRepository) Delete(id uint) error {
	return r.db.Delete(&media.MediaAccount{}, id).Error
}

func (r *MediaAccountRepository) FindOptions() ([]dict.Option, error) {
	var list []dict.Option
	err := r.db.Model(&media.MediaAccount{}).Select("name as label, id as value").Where("status = 1").Order("id ASC").Find(&list).Error
	return list, err
}