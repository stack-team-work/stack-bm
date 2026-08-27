package sync

import (
	"fmt"
	"strconv"

	mediaRepo "stack-bm/internal/repository/mkt/media"
	kuaishouAPI "stack-bm/internal/service/mkt/ks/v1/api"
	kstoken "stack-bm/internal/service/mkt/ks/v1/token"
)

// SyncBalance 拉取快手账户余额并回填 media_accounts.balance
func SyncBalance(accountID uint) error {
	tokens := kstoken.NewTokenService()
	uidStr, token, err := tokens.Context(accountID)
	if err != nil {
		return err
	}
	advID, err := strconv.Atoi(uidStr)
	if err != nil {
		return fmt.Errorf("平台账户UID非法: %s", uidStr)
	}
	balance, err := kuaishouAPI.GetAccountFund(token, advID)
	if err != nil {
		return err
	}
	return mediaRepo.NewMediaAccountRepository().UpdateBalance(accountID, balance)
}
