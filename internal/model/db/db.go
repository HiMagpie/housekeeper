package db

import (
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	"housekeeper/internal/com/cfg"
	"housekeeper/internal/com/logger"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *dbx.DB
)

func init() {
	var err error
	db, err = dbx.Open("mysql", fmt.Sprintf("%s@/%s", cfg.C.Db.Username, cfg.C.Db.Database))
	if err != nil {
		panic(err.Error())
	}

	db.LogFunc = func(format string, a ...interface{}) {
		logger.Debug("sql", format, a)
	}
	db.DB().SetMaxOpenConns(1000)
	db.DB().SetMaxIdleConns(10)
}

func GetDb() *dbx.DB {
	return db
}