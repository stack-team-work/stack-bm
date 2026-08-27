package sync

import (
	"errors"
	"fmt"
	"strconv"

	mediaRepo "stack-bm/internal/repository/mkt/media"
	biliSync "stack-bm/internal/service/mkt/bili/v1/sync"
	ksSync "stack-bm/internal/service/mkt/ks/v1/sync"
	tcSync "stack-bm/internal/service/mkt/tc/v1/sync"
	ttSync "stack-bm/internal/service/mkt/tt/v1/sync"
	"stack-bm/pkg/constants"
)

// Runner 广告数据定时同步调度器（供 Linux cron 调用）
// 仅负责渠道识别与渠道账户枚举；各渠道同步逻辑自治于各自 v1/sync，
// 无共享同步接口，渠道互不影响
type Runner struct {
	accountRepo *mediaRepo.MediaAccountRepository
	mediaRepo   *mediaRepo.MediaRepository
}

func NewRunner() *Runner {
	return &Runner{
		accountRepo: mediaRepo.NewMediaAccountRepository(),
		mediaRepo:   mediaRepo.NewMediaRepository(),
	}
}

// Run 同步指定渠道的指定层级；channelMark 为 media.mark（tt/tc/bili/ks），account=0 同步全部
func (r *Runner) Run(channelMark, level, accountID string) error {
	channel, ok := constants.MediaMarkChannel[channelMark]
	if !ok {
		return errors.New("未知渠道标识: " + channelMark)
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
		var err error
		switch channel {
		case constants.ChannelBili:
			err = biliSync.Run(level, acc.ID)
		case constants.ChannelKs:
			err = ksSync.Run(level, acc.ID)
		case constants.ChannelTt:
			err = ttSync.Run(level, acc.ID)
		case constants.ChannelTc:
			err = tcSync.Run(level, acc.ID)
		default:
			err = errors.New("渠道暂不支持同步: " + channel)
		}
		if err != nil {
			fmt.Printf("  账户[%d]层级[%s]同步失败: %v\n", acc.ID, level, err)
			continue
		}
	}
	return nil
}
