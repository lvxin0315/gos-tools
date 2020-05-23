package push

import (
	"github.com/lvxin0315/gos-tools/dingtalk/client"
	dingMessage "github.com/lvxin0315/gos-tools/dingtalk/message"
	"github.com/lvxin0315/gos-tools/etc"
)

func SendDingDingMessage(mess string) (err error) {
	dingTalk := client.DingTalk{
		AccessToken: etc.DingAccessToken,
		Secret:      etc.DingSecret,
	}
	msg := dingMessage.NewTextMessage().SetContent(mess)
	_, err = dingTalk.Send(msg)
	return err
}
