package main

import (
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/lvxin0315/gos-tools/common"
	"github.com/lvxin0315/gos-tools/etc"
	"strings"
	"time"
)

//扫描nginx日志error
func init() {
	flag.StringVar(&common.LogPath, "l", "/data/log", "日志目录")
	flag.StringVar(&common.DomainName, "d", "www.baidu.com", "域名")
}

func main() {
	flag.Parse()
	for {
		filePath := fmt.Sprintf("%s/%s-error.log", common.LogPath, common.DomainName)
		tailNginxLog(filePath)
		if etc.Debug {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(5 * time.Minute)
		}
	}
}

func tailNginxLog(filePath string) {
	t, err := tail.TailFile(filePath, tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
		return
	}
	for line := range t.Lines {
		if strings.Index(line.Text, "[error]") > 0 {
			common.SendErr(line.Text)
		}
		//跨日重新tail
		if common.NowDay != time.Now().Day() {
			break
		}
	}
}
