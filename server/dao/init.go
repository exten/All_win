/*
 *  flower9
 *  17/12/2017
 *
 */
package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"All_win/server/common"
	"strconv"
)

//mySql 数据库连接
var MY_SQL_DBC *sql.DB
//MS SQL 数据库
var MS_SQL_DBC *sql.DB
//oracle 数据库
var OR_SQL_DBC *sql.DB
//postgraduate
var PG_SQL_DBC *sql.DB

func init() {
	log.Println("common init ...")
	initMysqlConnection()
}

//初始化mysql数据库连接
func initMysqlConnection() {
	//"pk10:123456@tcp(103.56.17.126:3306)/pk_10?charset=utf8"
	//"pk10:123456@tcp(103.56.17.126:8081)/pk_10?charset=utf8"
	dataSourceName := common.S_config.DbUser + ":" + common.S_config.DbPswd + "@tcp(" + common.S_config.DbHost + ":" + strconv.Itoa(common.S_config.DbPoint) + ")/" + common.S_config.DbSchema + "?charset=utf8"
	log.Println("dataSourceName => ", dataSourceName)
	_db, err := sql.Open("mysql", dataSourceName)
	if err == nil {
		MY_SQL_DBC = _db
		log.Println("init connection db", _db)
	} else {
		log.Println("Connection DB error", err)
	}

	err = _db.Ping()
	if err != nil {
		log.Println("_db error ...", err)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
