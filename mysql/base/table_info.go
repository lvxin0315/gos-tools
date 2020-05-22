package base

import (
	"github.com/jinzhu/gorm"
	"time"
)

const tableInfoSql = `SELECT
	*
FROM
	information_schema.TABLES
WHERE
	table_schema = ?
	AND table_name = ? `

type informationSchemaTable struct {
	TableSchema   string    `gorm:"column:TABLE_SCHEMA"`
	TableName     string    `gorm:"column:TABLE_NAME"`
	TableRows     int64     `gorm:"column:TABLE_ROWS"`
	AutoIncrement uint      `gorm:"column:AUTO_INCREMENT"`
	UpdateTime    time.Time `gorm:"column:UPDATE_TIME"`
}

func GetTableInfo(db *gorm.DB, database, table string) (ist *informationSchemaTable, err error) {
	ist = new(informationSchemaTable)
	err = db.Raw(tableInfoSql, database, table).Scan(ist).Error
	return
}
