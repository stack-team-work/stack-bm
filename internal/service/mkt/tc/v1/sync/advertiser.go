package sync

import (
	mediaModel "stack-bm/internal/model/mkt/media"
	"stack-bm/internal/service/mkt/oauth"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AdvertiserSyncService 腾讯广告主同步（预留实现）
type AdvertiserSyncService struct {
	auth *oauth.ManagerAuth
}

func NewAdvertiserSyncService(auth *oauth.ManagerAuth) *AdvertiserSyncService {
	return &AdvertiserSyncService{auth: auth}
}

// Sync 同步腾讯广告主列表
func (s *AdvertiserSyncService) Sync(manager *mediaModel.MediaManager, authInfo bson.M) error {
	_ = s.auth.UpdateAccountNum(manager, 0)
	return s.syncChannelAccounts(manager.ID)
}

// syncChannelAccounts 预留：同步更新渠道账户（当前渠道账户功能尚未开发）
func (s *AdvertiserSyncService) syncChannelAccounts(managerID uint) error {
	return nil
}
