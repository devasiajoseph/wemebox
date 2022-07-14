package website

import (
	"github.com/devasiajoseph/wemebox/db/sqlite"
)

func Init() {
	sqlite.InitDB()
}
