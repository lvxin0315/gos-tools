package base

import (
	"fmt"
	"github.com/lvxin0315/gos-tools/etc"
	"github.com/lvxin0315/gos-tools/tools/push"
	"github.com/sirupsen/logrus"
	"strings"
)

type ListenerCallback interface {
	Call(ist *informationSchemaTable, tld *tableListenerDiscrepantData) (err error)
}

//调试callback
type TestListenerCallback struct {
}

func (callback *TestListenerCallback) Call(ist *informationSchemaTable, tld *tableListenerDiscrepantData) (err error) {
	logrus.Info(ist)
	logrus.Info(tld)
	return
}

type DingDingListenerCallback struct {
}

func (callback *DingDingListenerCallback) Call(ist *informationSchemaTable, tld *tableListenerDiscrepantData) (err error) {
	msg := fmt.Sprintf("钉钉callback mysql(%s:%s)/%s的table:%s发生了变化，event：%s;自增主键：%d",
		etc.MysqlHost,
		etc.MysqlPort,
		etc.MysqlDatabase,
		ist.TableName,
		strings.Join(tld.Events, ","),
		ist.AutoIncrement)
	push.SendDingDingMessage(msg)
	return
}
