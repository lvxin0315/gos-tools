package base

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	AutoIncrementEvent = "AutoIncrement"
	TotalRowsEvent     = "TotalRows"
	UpdateTimeEvent    = "UpdateTime"
)

type tableListenerDiscrepantData struct {
	Events    []string
	TableRows int64
}

type tableListener struct {
	Database                     string
	Table                        string
	DiscrepantData               *tableListenerDiscrepantData
	LatestInformationSchemaTable *informationSchemaTable
	LatestRunTime                time.Time
	Callback                     ListenerCallback
}

//初始化
func NewTableListener(database, table string, callback ListenerCallback) *tableListener {
	return &tableListener{
		Database:                     database,
		Table:                        table,
		DiscrepantData:               new(tableListenerDiscrepantData),
		LatestInformationSchemaTable: new(informationSchemaTable),
		LatestRunTime:                time.Now(),
		Callback:                     callback,
	}
}

//第一获取table信息
func (listener *tableListener) Init(db *gorm.DB) (err error) {
	ist, err := GetTableInfo(db.New(), listener.Database, listener.Table)
	if err != nil {
		return
	}
	listener.LatestInformationSchemaTable = ist
	listener.DiscrepantData.TableRows = ist.TableRows
	listener.LatestRunTime = time.Now()
	return
}

//更新信息
func (listener *tableListener) Update(db *gorm.DB) (err error) {
	ist, err := GetTableInfo(db.New(), listener.Database, listener.Table)
	if err != nil {
		return
	}
	//判断变化内容
	tld := new(tableListenerDiscrepantData)
	//自增id变化
	if ist.AutoIncrement != listener.LatestInformationSchemaTable.AutoIncrement {
		tld.Events = append(tld.Events, AutoIncrementEvent)
	}
	//总行数变化
	if ist.TableRows != listener.LatestInformationSchemaTable.TableRows {
		tld.Events = append(tld.Events, TotalRowsEvent)
	}
	//更新时间变化
	if !ist.UpdateTime.IsZero() && ist.UpdateTime != listener.LatestInformationSchemaTable.UpdateTime {
		tld.Events = append(tld.Events, UpdateTimeEvent)
	}
	//无变化，终止
	if len(tld.Events) == 0 {
		return
	}
	//有变化，先记录行数
	tld.TableRows = ist.TableRows
	listener.DiscrepantData = tld
	//记录最后信息
	listener.LatestInformationSchemaTable = ist
	listener.LatestRunTime = time.Now()
	//是否有回调
	if listener.Callback != nil {
		err = listener.Callback.Call(listener.LatestInformationSchemaTable, listener.DiscrepantData)
	}
	return
}
