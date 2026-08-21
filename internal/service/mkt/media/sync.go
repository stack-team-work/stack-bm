package media

import (
	"context"
	"encoding/json"
	"fmt"

	"stack-bm/internal/database"
	"stack-bm/pkg/constants"
)

// pushSyncQueue 将管家ID推入广告主同步队列（Redis）
func (s *ChannelAuthService) pushSyncQueue(managerID uint) {
	if database.Redis == nil {
		return
	}
	payload, _ := json.Marshal(map[string]uint{"mkt_account_manager_id": managerID})
	_ = database.Redis.RPush(context.Background(), constants.RedisKeyMktManagerSyncAdvertiser, payload).Err()
}

// SyncAdvertiser 同步管家广告主列表
// 先实现各媒体广告主/主体列表接口调用，预留 syncChannelAccounts 方法用于同步渠道账户（该功能暂无）
func (s *ChannelAuthService) SyncAdvertiser(managerID uint) error {
	manager, err := s.managerRepo.FindByID(managerID)
	if err != nil {
		return fmt.Errorf("管家不存在")
	}
	if manager.AuthStatus != constants.ManagerAuthStatusComplete {
		return fmt.Errorf("管家账号暂未授权成功，不支持刷新广告主列表")
	}
	channel, err := s.ResolveChannel(manager)
	if err != nil {
		return err
	}
	authInfo, err := s.tokenRepo.FindByManagerID(int(managerID))
	if err != nil {
		return err
	}
	if authInfo == nil {
		return fmt.Errorf("管家授权信息不存在")
	}

	switch channel {
	case constants.ChannelTt:
		err = s.ttSyncAdvertiser(manager, authInfo)
	case constants.ChannelBili:
		err = s.biliSyncAdvertiser(manager, authInfo)
	case constants.ChannelKs:
		err = s.ksSyncAdvertiser(manager, authInfo)
	case constants.ChannelTc:
		err = s.tcSyncAdvertiser(manager, authInfo)
	default:
		return fmt.Errorf("渠道【%s】暂不支持广告主同步", channel)
	}
	if err != nil {
		return err
	}
	return nil
}

// syncChannelAccounts 预留：同步更新渠道账户
// 当前渠道账户功能尚未开发，先空实现，待渠道账户模块就绪后实现
func (s *ChannelAuthService) syncChannelAccounts(channel string, managerID uint) error {
	return nil
}
