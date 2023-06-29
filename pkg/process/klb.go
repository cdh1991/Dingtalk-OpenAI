package process

import (
	"errors"
	"fmt"
	"strings"

	"github.com/eryajf/chatgpt-dingtalk/public"
	"github.com/eryajf/chatgpt-dingtalk/pkg/dingbot"
	

)

// GeneratePrompt 生成当次请求的 Prompt
func GenerateKlb(rmsg *dingbot.ReceiveMsg) (rst string, err error) {
	msg := rmsg.Text.Content
	for _, klb := range *public.Klb {
		if strings.HasPrefix(msg, klb.Title) {
			if strings.TrimSpace(msg) == klb.Title {
				rst = fmt.Sprintf("%s：\n%s___输入内容___%s", klb.Title, klb.Prefix, klb.Suffix)
				err = errors.New("消息内容为空") // 当提示词之后没有文本，抛出异常，以便直接返回Prompt所代表的内容
			} else {
				rst = klb.Prefix + strings.TrimSpace(strings.Replace(msg, klb.Title, "", -1)) + klb.Suffix
			}
			return
		} else {
			rst = msg
		}
	}
	rmsg.ReplyToDingtalk(string(dingbot.MARKDOWN), rst+1123)
	return
}
