package main

import (
	"flag"
	"fmt"
	"github.com/CatchZeng/dingtalk/client"
	dingMessage "github.com/CatchZeng/dingtalk/message"
	"github.com/hpcloud/tail"
	"strings"
	"time"
)

//扫描tp5日志error

const Debug = false

var NowDay = time.Now().Day()

const DingAccessToken = "a4603a75fcf43a592cb43f939ccf01e74c9732e3591c92fa0505cd64818cd538"
const DingSecret = "SEC4ab56819b4d6ae20f203064320d44612a3940875ba2042b47c8a4d77ffd36239"

var logPath string

func init() {
	flag.StringVar(&logPath, "l", "/data/log", "日志目录")
}

func main() {
	flag.Parse()
	for {
		filePath := logPath + fmt.Sprintf("/%s/%d.log", time.Now().Format("200601"), time.Now().Day())
		tailF(filePath)
		if Debug {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(5 * time.Minute)
		}
	}
}

func tailF(filePath string) {
	t, err := tail.TailFile(filePath, tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
		return
	}
	for line := range t.Lines {
		if strings.Index(line.Text, "err") > 0 {
			sendErr(line.Text)
		}
		//跨日重新tail
		if NowDay != time.Now().Day() {
			break
		}
	}
}

func sendErr(errMsg string) {
	if Debug {
		fmt.Println(errMsg)
		return
	}
	dingTalk := client.DingTalk{
		AccessToken: DingAccessToken,
		Secret:      DingSecret,
	}
	msg := dingMessage.NewTextMessage().SetContent(errMsg)
	dingTalk.Send(msg)
}
