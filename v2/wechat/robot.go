package wechat

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type NotifyResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// https://developer.work.weixin.qq.com/document/path/91770
// https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=1029c86a-6f89-403c-809f-105cf160e73d
func Notify(key, content string, mentionedList []string) error {
	client := resty.New()
	resp, err := client.SetRetryCount(retryCount).R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		SetQueryParam("key", key).
		SetBody(map[string]any{
			"msgtype": "text",
			"text": map[string]any{
				"content": content,
			},
			"mentioned_list": mentionedList,
		}).
		SetResult(&NotifyResult{}).
		// 不设置这个，无法解析到结构体
		ForceContentType("application/json").
		Post(robotBaseURL)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("error code: %d", resp.StatusCode())
	}

	r := resp.Result().(*NotifyResult)

	if r.ErrCode != 0 {
		return fmt.Errorf("error code: %d, error message: %v", r.ErrCode, r.ErrMsg)
	}

	return nil
}
