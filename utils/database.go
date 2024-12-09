package utils

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/punkestu/dodollaptoponline-go/config"
)

var dbSingleton *sql.DB = nil

func DB() *sql.DB {
	if dbSingleton != nil {
		return dbSingleton
	}

	driver, url := config.GetDBConfig()

	db, err := sql.Open(driver, url)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dbSingleton = db

	return db
}
