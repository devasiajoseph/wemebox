package website

import "github.com/devasiajoseph/wemebox/db"

func initWebsiteDb() {
	db := db.DB
	db.MustExec("create tables")
}

func Init() {
	db.InitSQLiteDB()
}
