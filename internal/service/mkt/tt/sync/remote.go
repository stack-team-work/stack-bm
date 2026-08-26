package sync

import (
	"fmt"

	ttRepo "stack-bm/internal/repository/mkt/tt"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tt/api"
	"stack-bm/pkg/constants"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// RemoteSyncService 头条V3广告数据远程同步
type RemoteSyncService struct {
	auth       *oauth.ManagerAuth
	adDataRepo *ttRepo.AdDataRepository
}

func NewRemoteSyncService(auth *oauth.ManagerAuth) *RemoteSyncService {
	return &RemoteSyncService{auth: auth, adDataRepo: ttRepo.NewAdDataRepository()}
}

// SyncAccountBalance 头条账户余额（暂无可用的钱包接口，占位）
func (s *RemoteSyncService) SyncAccountBalance(accountID uint) error {
	return nil
}

// SyncLevel 拉取层级列表并写入 mktAdData 集合（头条V3仅项目/第一层级）
func (s *RemoteSyncService) SyncLevel(accountID uint, level string) error {
	advID, token, err := s.auth.GetAccountContext(accountID)
	if err != nil {
		return err
	}
	if level != constants.AdDataLevelCampaign {
		return fmt.Errorf("头条V3不支持同步层级: %s", level)
	}

	list, err := api.ListProjects(token, advID, 1, 100)
	if err != nil {
		return err
	}
	idField := constants.AdDataLevelIDField[level]
	for _, row := range list {
		row["account_id"] = advID
		row["channel_id"] = "TT"
		doc := bson.M{}
		for k, v := range row {
			doc[k] = v
		}
		if _, ok := doc[idField]; !ok {
			continue
		}
		if err := s.adDataRepo.Upsert(level, idField, doc); err != nil {
			return err
		}
	}
	return nil
}