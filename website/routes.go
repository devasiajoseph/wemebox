package website

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/devasiajoseph/wemebox/core"
	"github.com/gorilla/mux"
)

func AddMultiRoutes(r *mux.Router) {
	r.HandleFunc("/", HomePage)
	r.HandleFunc("/p/{slug}", SlugPage)

	var fs http.Handler
	var dfs http.Handler
	var cljsfs http.Handler
	if core.Config.Compress {
		fs = gziphandler.GzipHandler(http.FileServer(http.Dir(StaticDir)))
		dfs = gziphandler.GzipHandler(http.FileServer(http.Dir(DashboardStaticDir)))
		cljsfs = gziphandler.GzipHandler(http.FileServer(http.Dir(CljsDir)))
	} else {
		fs = http.FileServer(http.Dir(StaticDir))
		dfs = http.FileServer(http.Dir(DashboardStaticDir))
		cljsfs = http.FileServer(http.Dir(CljsDir))
	}

	s := http.StripPrefix("/static/", fs)
	r.PathPrefix("/static/").Handler(s)

	ds := http.StripPrefix("/dstatic/", dfs)
	r.PathPrefix("/dstatic/").Handler(ds)

	cljs := http.StripPrefix("/cljs-out/", cljsfs)
	r.PathPrefix("/cljs-out/").Handler(cljs)
}
