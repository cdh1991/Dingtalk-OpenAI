package process

import (
	"errors"
	"fmt"
	"strings"

	"public"
	"/pkg/dingbot"
	

)

// GeneratePrompt 生成当次请求的 Prompt
func GenerateKlb(msg string) (rst string, err error) {
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
	msg.ReplyToDingtalk(string(dingbot.MARKDOWN), rst)
	return
}
