package sync

import (
	"errors"
	"fmt"
	"strconv"

	mediaRepo "stack-bm/internal/repository/mkt/media"
	biliSync "stack-bm/internal/service/mkt/bili/sync"
	ksSync "stack-bm/internal/service/mkt/ks/sync"
	"stack-bm/internal/service/mkt/oauth"
	tcSync "stack-bm/internal/service/mkt/tc/sync"
	ttSync "stack-bm/internal/service/mkt/tt/sync"
	"stack-bm/pkg/constants"
)

// channelSyncer 各渠道远程同步接口
type channelSyncer interface {
	SyncAccountBalance(accountID uint) error
	SyncLevel(accountID uint, level string) error
}

// Runner 广告数据定时同步调度器（供 Linux cron 调用）
type Runner struct {
	auth        *oauth.ManagerAuth
	accountRepo *mediaRepo.MediaAccountRepository
	mediaRepo   *mediaRepo.MediaRepository
	syncers     map[string]channelSyncer
}

func NewRunner() *Runner {
	auth := oauth.NewManagerAuth()
	return &Runner{
		auth:        auth,
		accountRepo: mediaRepo.NewMediaAccountRepository(),
		mediaRepo:   mediaRepo.NewMediaRepository(),
		syncers: map[string]channelSyncer{
			constants.ChannelBili: biliSync.NewRemoteSyncService(auth),
			constants.ChannelKs:   ksSync.NewRemoteSyncService(auth),
			constants.ChannelTt:   ttSync.NewRemoteSyncService(auth),
			constants.ChannelTc:   tcSync.NewRemoteSyncService(auth),
		},
	}
}

// Run 同步指定渠道的指定层级；channelMark 为 media.mark（tt/tc/bili/ks），account=0 同步全部
func (r *Runner) Run(channelMark, level, accountID string) error {
	channel, ok := constants.MediaMarkChannel[channelMark]
	if !ok {
		return errors.New("未知渠道标识: " + channelMark)
	}
	syncer, ok := r.syncers[channel]
	if !ok {
		return errors.New("渠道暂不支持同步: " + channel)
	}

	media, err := r.mediaRepo.FindByMark(channelMark)
	if err != nil {
		return errors.New("媒体渠道不存在: " + channelMark)
	}

	accounts, err := r.accountRepo.FindAllByChannel(int(media.ID))
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		return errors.New("该渠道暂无渠道账户")
	}

	for _, acc := range accounts {
		if accountID != "" && accountID != "0" {
			id, _ := strconv.Atoi(accountID)
			if int(acc.ID) != id {
				continue
			}
		}
		fmt.Printf("同步渠道[%s]账户[%d:%s] 层级[%s]\n", channelMark, acc.ID, acc.Name, level)
		if level == constants.AdDataLevelAccount {
			if err := syncer.SyncAccountBalance(acc.ID); err != nil {
				fmt.Printf("  账户[%d]余额同步失败: %v\n", acc.ID, err)
				continue
			}
		} else {
			if err := syncer.SyncLevel(acc.ID, level); err != nil {
				fmt.Printf("  账户[%d]层级[%s]同步失败: %v\n", acc.ID, level, err)
				continue
			}
		}
	}
	return nil
}