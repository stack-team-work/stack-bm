package sync

import (
	mediaRepo "stack-bm/internal/repository/mkt/media"
	tcAPI "stack-bm/internal/service/mkt/tc/v1/api"
	tctoken "stack-bm/internal/service/mkt/tc/v1/token"
)

// SyncBalance 拉取腾讯账户钱包余额并回填 media_accounts.balance
func SyncBalance(accountID uint) error {
	tokens := tctoken.NewTokenService()
	uidStr, token, err := tokens.Context(accountID)
	if err != nil {
		return err
	}
	balance, err := tcAPI.GetAccountWallet(token, uidStr)
	if err != nil {
		return err
	}
	return mediaRepo.NewMediaAccountRepository().UpdateBalance(accountID, balance)
}
