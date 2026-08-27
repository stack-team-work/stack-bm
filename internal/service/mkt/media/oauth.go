package media

import (
	"errors"
	"fmt"

	mediaModel "stack-bm/internal/model/mkt/media"
	"stack-bm/internal/service/mkt/bili"
	biliSync "stack-bm/internal/service/mkt/bili/v1/sync"
	"stack-bm/internal/service/mkt/ks"
	ksSync "stack-bm/internal/service/mkt/ks/v1/sync"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/internal/service/mkt/tc"
	tcSync "stack-bm/internal/service/mkt/tc/v1/sync"
	"stack-bm/internal/service/mkt/tt"
	ttSync "stack-bm/internal/service/mkt/tt/v1/sync"
	"stack-bm/pkg/constants"
	"stack-bm/pkg/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// oauthService 各媒体管家授权服务统一接口
type oauthService interface {
	BuildOauthURL(appID, state, redirectURI string) string
	FinishOauth(params map[string]string) error
}

// advertiserSyncer 各媒体广告主同步服务统一接口
type advertiserSyncer interface {
	Sync(manager *mediaModel.MediaManager, authInfo bson.M) error
}

// ChannelAuthService 管家授权调度器：生成授权 URL、分发回调与广告主同步
type ChannelAuthService struct {
	auth         *oauth.ManagerAuth
	oauthSvcs    map[string]oauthService
	advertisers  map[string]advertiserSyncer
}

func NewChannelAuthService() *ChannelAuthService {
	auth := oauth.NewManagerAuth()
	return &ChannelAuthService{
		auth: auth,
		oauthSvcs: map[string]oauthService{
			constants.ChannelBili: bili.NewOauthService(auth),
			constants.ChannelKs:   ks.NewOauthService(auth),
			constants.ChannelTt:   tt.NewOauthService(auth),
			constants.ChannelTc:   tc.NewOauthService(auth),
		},
		advertisers: map[string]advertiserSyncer{
			constants.ChannelBili: biliSync.NewAdvertiserSyncService(auth),
			constants.ChannelKs:   ksSync.NewAdvertiserSyncService(auth),
			constants.ChannelTt:   ttSync.NewAdvertiserSyncService(auth),
			constants.ChannelTc:   tcSync.NewAdvertiserSyncService(auth),
		},
	}
}

// ResolveChannel 通过 manager 关联的 media.mark 解析渠道标识
func (s *ChannelAuthService) ResolveChannel(manager *mediaModel.MediaManager) (string, error) {
	if manager.MediaID == 0 {
		return "", errors.New("管家未关联媒体渠道")
	}
	m, err := s.auth.MediaRepo.FindByID(uint(manager.MediaID))
	if err != nil {
		return "", errors.New("媒体渠道不存在")
	}
	channel, ok := constants.MediaMarkChannel[m.Mark]
	if !ok {
		return "", fmt.Errorf("媒体渠道【%s】暂不支持管家授权", m.Mark)
	}
	return channel, nil
}

// GetOauthUrl 生成授权跳转 URL
func (s *ChannelAuthService) GetOauthUrl(managerID uint) (string, error) {
	manager, err := s.auth.ManagerRepo.FindByID(managerID)
	if err != nil {
		return "", errors.New("管家不存在")
	}
	channel, err := s.ResolveChannel(manager)
	if err != nil {
		return "", err
	}
	svc, ok := s.oauthSvcs[channel]
	if !ok {
		return "", fmt.Errorf("渠道【%s】暂未开发管家授权", channel)
	}
	app, err := s.auth.GetApp(manager)
	if err != nil {
		return "", err
	}

	state := utils.XorEncrypt(fmt.Sprintf(`{"mkt_account_manager_id":%d}`, managerID), utils.XorKey)
	redirectURI := s.auth.RedirectURI(channel)
	return svc.BuildOauthURL(app.AppID, state, redirectURI), nil
}

// FinishOauth 回调分发到对应渠道
func (s *ChannelAuthService) FinishOauth(channel string, params map[string]string) error {
	svc, ok := s.oauthSvcs[channel]
	if !ok {
		return fmt.Errorf("渠道【%s】不支持授权", channel)
	}
	return svc.FinishOauth(params)
}

// SyncAdvertiser 同步管家广告主列表
func (s *ChannelAuthService) SyncAdvertiser(managerID uint) error {
	manager, err := s.auth.ManagerRepo.FindByID(managerID)
	if err != nil {
		return errors.New("管家不存在")
	}
	if manager.AuthStatus != constants.ManagerAuthStatusComplete {
		return errors.New("管家账号暂未授权成功，不支持刷新广告主列表")
	}
	channel, err := s.ResolveChannel(manager)
	if err != nil {
		return err
	}
	syncSvc, ok := s.advertisers[channel]
	if !ok {
		return fmt.Errorf("渠道【%s】暂不支持广告主同步", channel)
	}
	authInfo, err := s.auth.TokenRepo.FindByManagerID(int(managerID))
	if err != nil {
		return err
	}
	if authInfo == nil {
		return errors.New("管家授权信息不存在")
	}
	return syncSvc.Sync(manager, authInfo)
}