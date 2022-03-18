package website

import (
	"log"
	"net/http"

	"text/template"
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
	//fmt.Println(PagePath(paths.DirPath, pd.BasePageFile))
	//fmt.Println(pagePath)
	/*
		authUser, err := uauth.GetAuthenticatedUser(r)
		if err == nil {
			pd.UAuthLoggedIn = true
			pd.LoggedInUser = authUser.FullName
		}*/

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
	page := PageData{
		Title:        "Home page",
		Content:      "Home page content",
		PageFile:     "home.html",
		BasePageFile: "base.html",
	}
	RenderMultiPageTemplate(w, r, page)
}

func StaticPage(w http.ResponseWriter, r *http.Request) {

}
