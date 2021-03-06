package page

import (
	"log"
	"net/http"

	"text/template"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/gorilla/mux"
)

type PageData struct {
	Title   string
	Content string

	CSRF     string
	PageData interface{}
}

func RenderPageTemplate(w http.ResponseWriter, page string, base string, pd PageData) {
	paths := core.Paths{DirPath: core.BinPath, StaticUrl: core.StaticUrl}
	pagePath := core.PagePath(paths.DirPath, page)
	tmpl, err := template.ParseFiles(core.PagePath(paths.DirPath, base), pagePath)
	if err != nil {
		log.Println("Template error")
	}

	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println("Template exe error")
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	page := PageData{
		Title: "Home page", Content: "Home page content",
	}
	RenderPageTemplate(w, "home.html", "base.html", page)
}

func StaticPage(w http.ResponseWriter, r *http.Request) {

}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/", HomePage)
}
