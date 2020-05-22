package common

import (
	"fmt"
	"github.com/lvxin0315/gos-tools/etc"
	"github.com/lvxin0315/gos-tools/tools/push"
	"time"
)

var (
	DomainName string
	LogPath    string
)

var NowDay = time.Now().Day()

/**
发送error
*/
func SendErr(errMsg string) {
	if etc.Debug {
		fmt.Println(errMsg)
		return
	}
	_ = push.SendDingDingMessage(errMsg)
}
