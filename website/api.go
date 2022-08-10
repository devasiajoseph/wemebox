package website

import (
	"net/http"

	"github.com/gorilla/mux"
)

func saveWebsite(w http.ResponseWriter, r *http.Request) {

}

func AddApiRoutes(r *mux.Router) {
	r.HandleFunc("/api/website/save", saveWebsite)
}
