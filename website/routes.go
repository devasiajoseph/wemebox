package website

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddMultiRoutes(r *mux.Router) {
	r.HandleFunc("/", HomePage)
	r.HandleFunc("/{slug}", SlugPage)
	//fs := gziphandler.GzipHandler(http.FileServer(http.Dir(StaticDir)))
	fs := http.FileServer(http.Dir(StaticDir))
	s := http.StripPrefix("/static/", fs)
	r.PathPrefix("/static/").Handler(s)
}
