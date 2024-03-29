package main

import (
	"github.com/devasiajoseph/wemebox/apps/post"
	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/dashboard"
	"github.com/devasiajoseph/wemebox/db/postgres"
	"github.com/devasiajoseph/wemebox/uauth"
	"github.com/devasiajoseph/wemebox/website"
	"github.com/gorilla/mux"
)

func main() {
	core.Start()
	postgres.InitDB()
	r := mux.NewRouter()
	dashboard.Start(r)
	website.Start(r)
	uauth.Start(r)
	post.Start(r)
	if core.Config.Ssl {
		website.StartMultiHttps(r)
	} else {
		website.StartHttp(r)
	}
}
