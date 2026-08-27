package sync

import (
	mediaModel "stack-bm/internal/model/mkt/media"
	"stack-bm/internal/service/mkt/bili/v1/api"
	"stack-bm/internal/service/mkt/oauth"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// AdvertiserSyncService B站广告主（授权主体）同步
type AdvertiserSyncService struct {
	auth *oauth.ManagerAuth
}

func NewAdvertiserSyncService(auth *oauth.ManagerAuth) *AdvertiserSyncService {
	return &AdvertiserSyncService{auth: auth}
}

// Sync 同步B站主体（广告主）列表并更新管家账户数
func (s *AdvertiserSyncService) Sync(manager *mediaModel.MediaManager, authInfo bson.M) error {
	accessToken, _ := authInfo["access_token"].(string)
	subjects, err := api.GetSubjectList(accessToken)
	if err != nil {
		return err
	}
	_ = s.auth.UpdateAccountNum(manager, len(subjects))
	return s.syncChannelAccounts(manager.ID)
}

// syncChannelAccounts 预留：同步更新渠道账户（当前渠道账户功能尚未开发）
func (s *AdvertiserSyncService) syncChannelAccounts(managerID uint) error {
	return nil
}
