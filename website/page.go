package website

import (
	"fmt"
	"log"
	"net/http"

	"text/template"

	"github.com/devasiajoseph/wemebox/api"
)

type Object struct {
	Title    string
	Content  string
	CSRF     string
	PageData interface{}
}

func RenderPageTemplate(w http.ResponseWriter, r *http.Request, pd PageData) {
	//pageData := pd
	paths := Paths{DirPath: BinPath, StaticUrl: StaticUrl}
	pagePath := PagePath(paths.DirPath, pd.PageFile)
	pd.StaticUrl = "/static"
	tmpl, err := template.ParseFiles(PagePath(paths.DirPath, pd.BasePageFile), pagePath)
	if err != nil {
		log.Println("Template error")
	}

	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println(err)
		log.Println("Template exe error")
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	var pd PageData
	pd.Slug = "home"
	err := pd.DomainPage(r)
	if err != nil {
		log.Println(err)
		fmt.Printf("%+v\n", pd)
		fmt.Fprintf(w, "Error loading page")
		return
	}
	RenderMultiPageTemplate(w, r, pd)
}

func SlugPage(w http.ResponseWriter, r *http.Request) {
	slug := api.Vars(r, "slug")
	//fmt.Println(slug)
	pd := PageData{Slug: slug}
	err := pd.DomainPage(r)
	if err != nil {
		log.Println(err)
		fmt.Printf("%+v\n", pd)
		fmt.Fprintf(w, "Error loading page")
		return
	}
	RenderMultiPageTemplate(w, r, pd)
}

func StaticPage(w http.ResponseWriter, r *http.Request) {

}
