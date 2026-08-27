package sync

import (
	"fmt"
	"strconv"

	ksadRepo "stack-bm/internal/repository/mkt/ks/ad"
	kuaishouAPI "stack-bm/internal/service/mkt/ks/v1/api"
	kstoken "stack-bm/internal/service/mkt/ks/v1/token"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// SyncUnit 拉取快手广告列表并全量分页写入 mktAdData 集合
func SyncUnit(accountID uint) error {
	tokens := kstoken.NewTokenService()
	uidStr, token, err := tokens.Context(accountID)
	if err != nil {
		return err
	}
	advID, err := strconv.Atoi(uidStr)
	if err != nil {
		return fmt.Errorf("平台账户UID非法: %s", uidStr)
	}

	repo := ksadRepo.NewUnitRepository()
	for page := 1; ; page++ {
		rows, err := kuaishouAPI.ListUnits(token, advID, page, pageSize)
		if err != nil {
			return err
		}
		for _, row := range rows {
			doc := bson.M{}
			for k, v := range row {
				doc[k] = v
			}
			if _, ok := doc["aid"]; !ok {
				continue
			}
			doc["account_id"] = advID
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
