package api

import (
	"fmt"

	"stack-bm/pkg/utils"
)

const (
	SubjectsURI = "https://cm.bilibili.com/takumi/api/open_api/oauth2/subjects"
	cashURI     = "https://cm.bilibili.com/takumi/api/open_api/report/v2/cash"
)

// GetSubjectList 获取B站授权主体（广告主）
func GetSubjectList(accessToken string) ([]interface{}, error) {
	data, err := utils.HTTPGet(SubjectsURI, nil, bearerHeaders(accessToken))
	if err != nil {
		return nil, err
	}
	var res struct {
		Code    int           `json:"code"`
		Message string        `json:"message"`
		Data    []interface{} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("B站获取主体失败_%d_%s", res.Code, res.Message)
	}
	return res.Data, nil
}

// GetAccountCash 获取B站账户余额
func GetAccountCash(accessToken string) (float64, error) {
	data, err := utils.HTTPGet(cashURI, nil, bearerHeaders(accessToken))
	if err != nil {
		return 0, err
	}
	var res struct {
		Status string `json:"status"`
		Data   struct {
			Cash float64 `json:"cash"`
		} `json:"data"`
	}
	if err := utils.ParseJSON(data, &res); err != nil {
		return 0, err
	}
	if err := checkStatus(res.Status, "获取余额"); err != nil {
		return 0, err
	}
	return res.Data.Cash, nil
}
