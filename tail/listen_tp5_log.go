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

//扫描tp5日志error

func init() {
	flag.StringVar(&common.LogPath, "l", "/data/log", "日志目录")
}

func main() {
	flag.Parse()
	for {
		filePath := common.LogPath + fmt.Sprintf("/%s/%d.log", time.Now().Format("200601"), time.Now().Day())
		tailTp5Log(filePath)
		if etc.Debug {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(5 * time.Minute)
		}
	}
}

func tailTp5Log(filePath string) {
	t, err := tail.TailFile(filePath, tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
		return
	}
	for line := range t.Lines {
		if strings.Index(line.Text, "err") > 0 {
			common.SendErr(line.Text)
		}
		//跨日重新tail
		if common.NowDay != time.Now().Day() {
			break
		}
	}
}
