package sync

import (
	"fmt"
	"strconv"

	ksRepo "stack-bm/internal/repository/mkt/ks"
	"stack-bm/internal/service/mkt/ks/api"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/pkg/constants"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// RemoteSyncService 快手广告数据远程同步
type RemoteSyncService struct {
	auth       *oauth.ManagerAuth
	adDataRepo *ksRepo.AdDataRepository
}

func NewRemoteSyncService(auth *oauth.ManagerAuth) *RemoteSyncService {
	return &RemoteSyncService{auth: auth, adDataRepo: ksRepo.NewAdDataRepository()}
}

// SyncAccountBalance 拉取账户余额并回填 media_accounts
func (s *RemoteSyncService) SyncAccountBalance(accountID uint) error {
	uidStr, token, err := s.auth.GetAccountContext(accountID)
	if err != nil {
		return err
	}
	advID, _ := strconv.Atoi(uidStr)
	balance, err := api.GetAccountFund(token, advID)
	if err != nil {
		return err
	}
	return s.auth.AccountRepo.UpdateBalance(accountID, balance)
}

// SyncLevel 拉取层级列表并写入 mktAdData 集合
func (s *RemoteSyncService) SyncLevel(accountID uint, level string) error {
	uidStr, token, err := s.auth.GetAccountContext(accountID)
	if err != nil {
		return err
	}
	advID, _ := strconv.Atoi(uidStr)

	var list []map[string]interface{}
	switch level {
	case constants.AdDataLevelCampaign:
		list, err = api.ListCampaigns(token, advID, 1, 100)
	case constants.AdDataLevelUnit:
		list, err = api.ListUnits(token, advID, 1, 100)
	case constants.AdDataLevelCreative:
		list, err = api.ListCreatives(token, advID, 1, 100)
	default:
		return fmt.Errorf("快手不支持同步层级: %s", level)
	}
	if err != nil {
		return err
	}

	idField := constants.AdDataLevelIDField[level]
	for _, row := range list {
		row["account_id"] = advID
		row["channel_id"] = "KS"
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