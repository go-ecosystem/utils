package wechat

import "github.com/go-resty/resty/v2"

type SecCheckResponse struct {
	ErrCode int              `json:"errcode"`
	ErrMsg  string           `json:"errmsg"`
	Detail  []SecCheckDetail `json:"detail"`
	TraceID string           `json:"trace_id"`
	Result  SecCheckResult   `json:"result"`
}

type SecCheckDetail struct {
	Strategy string `json:"strategy"`
	ErrCode  int    `json:"errcode"`
	Suggest  string `json:"suggest"`
	Label    int    `json:"label"`
	Keyword  string `json:"keyword"`
	Prob     int    `json:"prob"`
}

type SecCheckResult struct {
	Suggest string `json:"suggest"`
	Label   int    `json:"label"`
}

const version = 2

// MsgSecCheck MsgSecCheck api
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html
func MsgSecCheck(openid, content, accessToken string, scene int) (*SecCheckResponse, error) {
	client := resty.New()
	resp, err := client.SetRetryCount(retryCount).R().
		SetQueryParams(map[string]string{
			"access_token": accessToken,
		}).
		SetBody(map[string]any{
			"content": content,
			"version": version,
			"scene":   scene,
			"openid":  openid,
		}).
		SetHeader("Accept", "application/json").
		SetResult(&SecCheckResponse{}).
		// 不设置这个，无法解析到结构体
		ForceContentType("application/json").
		Post(baseURL + "/wxa/msg_sec_check")

	if err != nil {
		return nil, err
	}

	result := resp.Result().(*SecCheckResponse)
	return result, nil
}
