package website

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

var certsPath = core.AbsolutePath("certs")

func HttpsRedirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}

func StartHttp(r *mux.Router) {
	log.Println("Starting  webserver at port " + core.Config.Port)

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
	err := http.ListenAndServe(":"+core.Config.Port,
		protectionMiddleware(r),
	)

	if err != nil {
		log.Println(err)
	}
}

func StartMultiHttps(r *mux.Router) {
	log.Println("Starting Secure webserver at port 443")
	go http.ListenAndServe(":80", http.HandlerFunc(HttpsRedirect))
	AddMultiRoutes(r)
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(certsPath),
		//HostPolicy: autocert.HostWhitelist("wemebox.com", "www.wemebox.com"),
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
