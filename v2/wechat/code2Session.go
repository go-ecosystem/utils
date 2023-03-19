package wechat

import (
	"github.com/go-resty/resty/v2"
)

// Code2SessionResult code2Session result
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Code2Session code2Session api
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
func Code2Session(appid, secret, code string) (*Code2SessionResult, error) {
	client := resty.New()
	resp, err := client.SetRetryCount(retryCount).R().
		SetQueryParams(map[string]string{
			"appid":      appid,
			"secret":     secret,
			"js_code":    code,
			"grant_type": "authorization_code",
		}).
		SetHeader("Accept", "application/json").
		SetResult(&Code2SessionResult{}).
		// 不设置这个，无法解析到结构体
		ForceContentType("application/json").
		Get(baseURL + "/sns/jscode2session")

	if err != nil {
		return nil, err
	}

	result := resp.Result().(*Code2SessionResult)
	return result, nil
}
