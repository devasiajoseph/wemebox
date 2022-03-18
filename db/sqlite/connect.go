package sqlite

import (
	"log"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func connectSQLiteDB() *sqlx.DB {
	log.Println(core.AbsolutePath("dbf/wemebox.db"))
	db, err := sqlx.Open("sqlite3", core.AbsolutePath("dbf/wemebox.db"))
	if err != nil {
		panic(err)
	}
	return db
}

func InitDB() {
	DB = connectSQLiteDB()
}
