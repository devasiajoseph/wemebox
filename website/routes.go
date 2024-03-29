package website

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/devasiajoseph/wemebox/core"
	"github.com/gorilla/mux"
)

func AddMultiRoutes(r *mux.Router) {
	r.HandleFunc("/", HomePage)
	r.HandleFunc("/page/{slug}", SlugPage)

	var fs http.Handler
	var gs http.Handler
	var dfs http.Handler
	var cljsfs http.Handler
	if core.Config.Compress {
		fs = gziphandler.GzipHandler(http.FileServer(http.Dir(StaticDir)))
		gs = gziphandler.GzipHandler(http.FileServer(http.Dir(GlobalStaticDir)))
		dfs = gziphandler.GzipHandler(http.FileServer(http.Dir(DashboardStaticDir)))
		cljsfs = gziphandler.GzipHandler(http.FileServer(http.Dir(CljsDir)))
	} else {
		fs = http.FileServer(http.Dir(StaticDir))
		gs = http.FileServer(http.Dir(GlobalStaticDir))
		dfs = http.FileServer(http.Dir(DashboardStaticDir))
		cljsfs = http.FileServer(http.Dir(CljsDir))
	}

	s := http.StripPrefix("/static/", fs)
	r.PathPrefix("/static/").Handler(s)

	gls := http.StripPrefix("/gstatic/", gs)
	r.PathPrefix("/gstatic/").Handler(gls)

	ds := http.StripPrefix("/dstatic/", dfs)
	r.PathPrefix("/dstatic/").Handler(ds)

	cljs := http.StripPrefix("/cljs-out/", cljsfs)
	r.PathPrefix("/cljs-out/").Handler(cljs)
}

func Start(r *mux.Router) {
	AddMultiRoutes(r)
	AddApiRoutes(r)
}
