package main

import (
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/lvxin0315/gos-tools/etc"
	"github.com/lvxin0315/gos-tools/tools/push"
	"strings"
	"time"
)

//扫描tp5日志error

var NowDay = time.Now().Day()

var logPath string

func init() {
	flag.StringVar(&logPath, "l", "/data/log", "日志目录")
}

func main() {
	flag.Parse()
	for {
		filePath := logPath + fmt.Sprintf("/%s/%d.log", time.Now().Format("200601"), time.Now().Day())
		tailF(filePath)
		if etc.Debug {
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
	if etc.Debug {
		fmt.Println(errMsg)
		return
	}
	_ = push.SendDingDingMessage(errMsg)
}
