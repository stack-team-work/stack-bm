package sync

import (
	"fmt"

	tcRepo "stack-bm/internal/repository/mkt/tc"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tc/api"
	"stack-bm/pkg/constants"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// RemoteSyncService 腾讯V3广告数据远程同步
type RemoteSyncService struct {
	auth       *oauth.ManagerAuth
	adDataRepo *tcRepo.AdDataRepository
}

func NewRemoteSyncService(auth *oauth.ManagerAuth) *RemoteSyncService {
	return &RemoteSyncService{auth: auth, adDataRepo: tcRepo.NewAdDataRepository()}
}

// SyncAccountBalance 拉取账户钱包余额并回填 media_accounts
func (s *RemoteSyncService) SyncAccountBalance(accountID uint) error {
	accIDStr, token, err := s.auth.GetAccountContext(accountID)
	if err != nil {
		return err
	}
	balance, err := api.GetAccountWallet(token, accIDStr)
	if err != nil {
		return err
	}
	return s.auth.AccountRepo.UpdateBalance(accountID, balance)
}

// SyncLevel 拉取层级列表并写入 mktAdData 集合（腾讯V3仅广告组/第一层级）
func (s *RemoteSyncService) SyncLevel(accountID uint, level string) error {
	accIDStr, token, err := s.auth.GetAccountContext(accountID)
	if err != nil {
		return err
	}
	if level != constants.AdDataLevelCampaign {
		return fmt.Errorf("腾讯V3不支持同步层级: %s", level)
	}

	list, err := api.ListAdgroups(token, accIDStr, 1, 100)
	if err != nil {
		return err
	}
	idField := constants.AdDataLevelIDField[level]
	for _, row := range list {
		row["account_id"] = accIDStr
		row["channel_id"] = "TC"
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