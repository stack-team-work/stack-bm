package sync

import (
	"fmt"
	"strconv"

	biliAPI "stack-bm/internal/service/mkt/bili/v1/api"
	bitoken "stack-bm/internal/service/mkt/bili/v1/token"
	mediaRepo "stack-bm/internal/repository/mkt/media"
)

// SyncBalance 拉取B站账户余额并回填 media_accounts.balance
func SyncBalance(accountID uint) error {
	tokens := bitoken.NewTokenService()
	uidStr, token, err := tokens.Context(accountID)
	if err != nil {
		return err
	}
	if _, err := strconv.Atoi(uidStr); err != nil {
		return fmt.Errorf("平台账户UID非法: %s", uidStr)
	}
	balance, err := biliAPI.GetAccountCash(token)
	if err != nil {
		return err
	}
	return mediaRepo.NewMediaAccountRepository().UpdateBalance(accountID, balance)
}
