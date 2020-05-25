package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//DB 数据库连接池
var DB *xorm.Engine

// SQL2JAVAMap sql => java
var SQL2JAVAMap = map[string]string{
	"VARCHAR":   "String",
	"CHAR":      "String",
	"BLOB":      "byte[]",
	"TEXT":      "String",
	"INTEGER":   "Long",
	"TINYINT":   "Integer",
	"SMALLINT":  "Integer",
	"BIT":       "Boolean",
	"BIGINT":    "Long",
	"INT":       "Integer",
	"FLOAT":     "Float",
	"DOUBLE":    "Double",
	"DECIMAL":   "BigDecimal",
	"BOOLEAN":   "Boolean",
	"DATE":      "Date",
	"TIME":      "Date",
	"DATETIME":  "Date",
	"TIMESTAMP": "Date",
}

func init() {

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.IP, dbConfig.Port, dbConfig.Dbname)

	DB, _ = xorm.NewEngine("mysql", url)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	// DB.SetTableMapper(tbMapper)
	// DB.SetColumnMapper(core.SnakeMapper{})

	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")

}
