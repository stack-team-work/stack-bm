package sync

import (
	"fmt"
	"strconv"

	ttadRepo "stack-bm/internal/repository/mkt/tt/ad"
	ttAPI "stack-bm/internal/service/mkt/tt/v1/api"
	tttoken "stack-bm/internal/service/mkt/tt/v1/token"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// SyncCampaign 拉取头条V3项目列表并全量分页写入 mktAdData 集合
func SyncCampaign(accountID uint) error {
	tokens := tttoken.NewTokenService()
	uidStr, token, err := tokens.Context(accountID)
	if err != nil {
		return err
	}
	if uidStr == "" {
		return fmt.Errorf("平台账户UID为空")
	}

	repo := ttadRepo.NewCampaignRepository()
	for page := 1; ; page++ {
		rows, err := ttAPI.ListProjects(token, uidStr, page, pageSize)
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
			// account_id 统一写 int（存量历史数据为字符串，读取侧已做兼容）
			accID, err := strconv.Atoi(uidStr)
			if err != nil {
				return fmt.Errorf("平台账户UID非法: %s", uidStr)
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
