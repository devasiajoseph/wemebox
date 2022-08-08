package dashboard

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/website"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

var dashTemplate = core.AbsolutePath("html/mdb.html")

func Dashboard(w http.ResponseWriter, r *http.Request) {
	p := website.PageData{
		Csrf:      csrf.Token(r),
		StaticUrl: "/dstatic",
	}

	tmpl, err := template.ParseFiles(dashTemplate)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Dashboard Error")
	}
	err = tmpl.Execute(w, p)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Dashboard Error")
	}
}

func Start(r *mux.Router) {
	r.HandleFunc("/wb/dashboard", Dashboard)
}
