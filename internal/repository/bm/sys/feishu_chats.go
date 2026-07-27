package sys

import (
	"stack-bm/internal/database"
	"stack-bm/internal/model/bm/sys"
	"gorm.io/gorm"
)

type FeishuChatRepository struct{ db *gorm.DB }

func NewFeishuChatRepository() *FeishuChatRepository { return &FeishuChatRepository{db: database.DBBM} }

func (r *FeishuChatRepository) Create(m *sys.FeishuChat) error { return r.db.Create(m).Error }

func (r *FeishuChatRepository) FindByID(id uint) (*sys.FeishuChat, error) {
	var m sys.FeishuChat
	err := r.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *FeishuChatRepository) FindPage(page, size int, keyword string, status int, feishuAppID int) ([]sys.FeishuChat, int64, error) {
	var list []sys.FeishuChat
	var total int64
	query := r.db.Model(&sys.FeishuChat{})
	if keyword != "" {
		query = query.Where("chat_id LIKE ? OR call_action LIKE ? OR action_title LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if status >= 0 {
		query = query.Where("status = ?", status)
	}
	if feishuAppID > 0 {
		query = query.Where("feishu_app_id = ?", feishuAppID)
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

func (r *FeishuChatRepository) FindAll() ([]sys.FeishuChat, error) {
	var list []sys.FeishuChat
	err := r.db.Where("status = 1").Find(&list).Error
	return list, err
}

func (r *FeishuChatRepository) Update(m *sys.FeishuChat) error { return r.db.Save(m).Error }
func (r *FeishuChatRepository) Delete(id uint) error          { return r.db.Delete(&sys.FeishuChat{}, id).Error }
