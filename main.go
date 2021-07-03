package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

var SKey = "JPcT1k6SwyqA2JX-oyGjGfOfHzKsN2BQdI4Cr56KG9M="
var StaticUrl = "/static"

type Paths struct {
	DirPath   string
	StaticUrl string
}

func CleanPath(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	return path
}

func AbsolutePath(path string) string {
	if flag.Lookup("test.v") == nil {
		return CleanPath(filepath.FromSlash(BinPath + "/" + path))
	} else {
		return CleanPath(filepath.FromSlash(path))
	}
}

func BinaryPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Error getting file path")
	}
	return dir
}
func FixPathSlash(path string) string {
	return filepath.FromSlash(path)
}
func PagePath(dirPath string, spage string) string {
	return FixPathSlash(dirPath + "/html/" + spage)
}

var BinPath = BinaryPath()
var StaticDir = AbsolutePath("static")

type Page struct {
	Title   string
	Content string
}

type PageData struct {
	Title    string
	Content  string
	PageData interface{}
}

func RenderPageTemplate(w http.ResponseWriter, page string, pd PageData) {
	paths := Paths{DirPath: BinPath, StaticUrl: StaticUrl}
	pagePath := PagePath(paths.DirPath, page)
	tmpl, err := template.ParseFiles(PagePath(paths.DirPath, "base.html"), pagePath)
	if err != nil {
		log.Println("Template error")
	}

	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println("Template exe error")
	}
}

type HomePage struct {
	HomeContent string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page := PageData{
		Title: "Home page", Content: "Home page content",
		PageData: HomePage{HomeContent: "Hello home content"}}
	RenderPageTemplate(w, "home.html", page)
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	page := PageData{
		Title: "News page", Content: "Home page content",
		PageData: HomePage{HomeContent: "Hello home content"}}
	RenderPageTemplate(w, "home.html", page)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/news/{slug}", NewsHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))
	protectionMiddleware := func(handler http.Handler) http.Handler {
		protectionFn := csrf.Protect(
			[]byte(SKey),
			csrf.Secure(false),
			csrf.Path("/"),
		)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use some kind of condition here to see if the router should use
			// the CSRF protection. For the sake of this example, we'll check
			// the path prefix.
			if !strings.HasPrefix(r.URL.Path, "/webhooks") {
				protectionFn(handler).ServeHTTP(w, r)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
	err := http.ListenAndServe(":8080",
		protectionMiddleware(r),
	)

	if err != nil {
		log.Println(err)
	}
}
