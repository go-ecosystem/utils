package wechat

import "github.com/go-resty/resty/v2"

type AccessTokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken GetAccessToken api
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
func GetAccessToken(appid, secret string) (*AccessTokenResult, error) {
	client := resty.New()
	resp, err := client.SetRetryCount(retryCount).R().
		SetQueryParams(map[string]string{
			"appid":      appid,
			"secret":     secret,
			"grant_type": "client_credential",
		}).
		SetHeader("Accept", "application/json").
		SetResult(&AccessTokenResult{}).
		// 不设置这个，无法解析到结构体
		ForceContentType("application/json").
		Get(baseURL + "/cgi-bin/token")

	if err != nil {
		return nil, err
	}

	result := resp.Result().(*AccessTokenResult)
	return result, nil
}
