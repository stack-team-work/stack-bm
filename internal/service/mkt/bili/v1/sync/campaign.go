package sync

import (
	"fmt"
	"strconv"

	biliadRepo "stack-bm/internal/repository/mkt/bili/ad"
	"stack-bm/internal/service/mkt/bili/v1/api"
	bitoken "stack-bm/internal/service/mkt/bili/v1/token"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// SyncCampaign 拉取B站计划列表并全量分页写入 mktAdData 集合
func SyncCampaign(accountID uint) error {
	tokens := bitoken.NewTokenService()
	uidStr, token, err := tokens.Context(accountID)
	if err != nil {
		return err
	}
	accID, err := strconv.Atoi(uidStr)
	if err != nil {
		return fmt.Errorf("平台账户UID非法: %s", uidStr)
	}

	repo := biliadRepo.NewCampaignRepository()
	for page := 1; ; page++ {
		rows, err := api.ListCampaigns(token, accID, page, pageSize)
		if err != nil {
			return err
		}
		for _, row := range rows {
			doc := bson.M{}
			for k, v := range row {
				doc[k] = v
			}
			if _, ok := doc["cpid"]; !ok {
				continue
			}
			doc["account_id"] = accID
			doc["channel_id"] = ChannelID
			if err := repo.Upsert(doc); err != nil {
				return err
			}
		}
		if len(rows) < pageSize {
			break
		}
	}
	return nil
}
