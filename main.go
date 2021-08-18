package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/NYTimes/gziphandler"
	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/website"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

var SKey = "JPcT1k6SwyqA2JX-oyGjGfOfHzKsN2BQdI4Cr56KG9M="
var StaticUrl = "/static"
var staticPath = AbsolutePath("static")
var certsPath = AbsolutePath("certs")

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
	Title     string
	Content   string
	PageData  interface{}
	Csrf      string
	StaticUrl string
}

func RenderPageTemplate(w http.ResponseWriter, page string, pd PageData) {
	paths := Paths{DirPath: BinPath, StaticUrl: StaticUrl}
	pagePath := PagePath(paths.DirPath, page)
	fmt.Println(PagePath(paths.DirPath, "base.html"))
	fmt.Println(pagePath)
	tmpl, err := template.ParseFiles(PagePath(paths.DirPath, "base.html"), pagePath)
	if err != nil {
		log.Println("Template error")
	}
	pd.StaticUrl = "/static"
	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println(err)
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

func AddMultiRoutes(r *mux.Router) {
	fs := gziphandler.GzipHandler(http.FileServer(http.Dir(staticPath)))
	s := http.StripPrefix("/static/", fs)
	r.PathPrefix("/static/").Handler(s)
}

func StartMultiHttps(r *mux.Router) {
	log.Println("Starting Secure webserver at port 80")
	AddMultiRoutes(r)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(certsPath),
		HostPolicy: autocert.HostWhitelist("wemebox.com", "www.wemebox.com"),
	}

	protectionMiddleware := func(handler http.Handler) http.Handler {
		protectionFn := csrf.Protect(
			[]byte(core.SKey),
			csrf.Secure(core.Secure),
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

	server := &http.Server{
		Addr:    ":443",
		Handler: protectionMiddleware(r),
		TLSConfig: &tls.Config{
			//Certificates: nil, // <-- s.ListenAndServeTLS will populate this field
			GetCertificate: certManager.GetCertificate,
			NextProtos:     []string{acme.ALPNProto},
		},
	}

	//go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Println(err)
	}
}

func StartHttp(r *mux.Router) {
	log.Println("Starting  webserver at port 8080")

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))
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

func main() {
	//website.Init()

	//StartHttp()
	r := mux.NewRouter()

	r.HandleFunc("/", website.HomePage)
	StartHttp(r)

}
