package website

import (
	"github.com/devasiajoseph/wemebox/db/sqlite"
)

func initWebsiteDb() {
	db := sqlite.DB
	db.MustExec("create tables")
}

func Init() {
	sqlite.InitDB()
}
