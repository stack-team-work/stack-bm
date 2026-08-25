package sync

import (
	mediaModel "stack-bm/internal/model/mkt/media"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tt/api"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AdvertiserSyncService 头条广告主同步
type AdvertiserSyncService struct {
	auth *oauth.ManagerAuth
}

func NewAdvertiserSyncService(auth *oauth.ManagerAuth) *AdvertiserSyncService {
	return &AdvertiserSyncService{auth: auth}
}

// Sync 同步头条广告主列表
func (s *AdvertiserSyncService) Sync(manager *mediaModel.MediaManager, authInfo bson.M) error {
	accessToken, _ := authInfo["access_token"].(string)
	app, err := s.auth.GetApp(manager)
	if err != nil {
		return err
	}
	ids, err := api.GetAdvertiserList(accessToken, app.AppID, app.AppSecret)
	if err != nil {
		return err
	}
	_ = s.auth.UpdateAccountNum(manager, len(ids))
	return s.syncChannelAccounts(manager.ID)
}

// syncChannelAccounts 预留：同步更新渠道账户（当前渠道账户功能尚未开发）
func (s *AdvertiserSyncService) syncChannelAccounts(managerID uint) error {
	return nil
}