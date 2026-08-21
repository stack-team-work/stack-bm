package media

import (
	"fmt"
	"net/url"

	"stack-bm/internal/model/mkt/media"
	"stack-bm/pkg/constants"
	"stack-bm/pkg/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const ttAdvertiserURI = "https://ad.oceanengine.com/open_api/oauth2/advertiser/get/"

// ttSyncAdvertiser 头条广告主列表同步
func (s *ChannelAuthService) ttSyncAdvertiser(manager *media.MediaManager, authInfo bson.M) error {
	accessToken, _ := authInfo["access_token"].(string)
	app, err := s.getApp(manager)
	if err != nil {
		return err
	}
	query := url.Values{}
	query.Set("access_token", accessToken)
	query.Set("app_id", app.AppID)
	query.Set("secret", app.AppSecret)

	data, err := utils.HTTPGet(ttAdvertiserURI, query, nil)
	if err != nil {
		return err
	}
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			List []struct {
				AdvertiserID string `json:"advertiser_id"`
			} `json:"list"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return err
	}
	if res.Code != 0 {
		return fmt.Errorf("头条广告主列表获取失败_%d_%s", res.Code, res.Message)
	}
	_ = s.updateManagerAccountNum(manager, len(res.Data.List))
	return s.syncChannelAccounts(constants.ChannelTt, manager.ID)
}

// biliSyncAdvertiser B站主体（广告主）列表同步
func (s *ChannelAuthService) biliSyncAdvertiser(manager *media.MediaManager, authInfo bson.M) error {
	accessToken, _ := authInfo["access_token"].(string)
	subjects, err := s.biliGetSubjectList(accessToken)
	if err != nil {
		return err
	}
	_ = s.updateManagerAccountNum(manager, len(subjects))
	return s.syncChannelAccounts(constants.ChannelBili, manager.ID)
}

// ksSyncAdvertiser 快手广告主列表同步
func (s *ChannelAuthService) ksSyncAdvertiser(manager *media.MediaManager, authInfo bson.M) error {
	// 快手广告主列表接口调用（预留实现）
	_ = s.updateManagerAccountNum(manager, 0)
	return s.syncChannelAccounts(constants.ChannelKs, manager.ID)
}

// tcSyncAdvertiser 腾讯广告主列表同步
func (s *ChannelAuthService) tcSyncAdvertiser(manager *media.MediaManager, authInfo bson.M) error {
	// 腾讯广告主列表接口调用（预留实现）
	_ = s.updateManagerAccountNum(manager, 0)
	return s.syncChannelAccounts(constants.ChannelTc, manager.ID)
}

// updateManagerAccountNum 更新管家绑定账户数
func (s *ChannelAuthService) updateManagerAccountNum(manager *media.MediaManager, num int) error {
	manager.AccountNum = num
	return s.managerRepo.Update(manager)
}
