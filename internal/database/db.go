package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDB() *sqlx.DB {
	dsn := "root:rootoor@tcp(127.0.0.1:3306)/gochi"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db
}
