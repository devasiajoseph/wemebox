package main

import (
	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/db/postgres"
	"github.com/devasiajoseph/wemebox/website"
	"github.com/gorilla/mux"
)

func main() {
	core.Start()
	postgres.InitDB()
	r := mux.NewRouter()
	website.AddMultiRoutes(r)
	website.StartHttp(r)

}
