package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gos-tools/etc"
	"github.com/lvxin0315/gos-tools/mysql/base"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func main() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		etc.MysqlUser,
		etc.MysqlPass,
		etc.MysqlHost,
		etc.MysqlPort,
		etc.MysqlDatabase)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	db.DB().SetMaxOpenConns(5)
	db.DB().SetMaxIdleConns(1)

	//加载需要监听的table
	initTable(db)

	select {}
}

//加载监听的table
func initTable(db *gorm.DB) {
	tableList := strings.Split(etc.MysqlListenTables, ",")
	for _, t := range tableList {
		var callback base.ListenerCallback
		go func(table string) {
			if etc.Debug {
				callback = new(base.TestListenerCallback)
			} else {
				callback = new(base.DingDingListenerCallback)
			}
			tl := base.NewTableListener(etc.MysqlDatabase, table, callback)
			err := tl.Init(db)
			if err != nil {
				logrus.Error("NewTableListener error: ", err)
				return
			}
			for {
				time.Sleep(etc.MysqlListenTime * time.Millisecond)
				err = tl.Update(db)
				if err != nil {
					logrus.Error("tableListener update error: ", err)
					continue
				}
			}
		}(t)
	}
}
