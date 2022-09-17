package website

import (
	"net/http"

	"github.com/devasiajoseph/wemebox/api"
	"github.com/devasiajoseph/wemebox/uauth"
	"github.com/gorilla/mux"
)

func saveWebsite(w http.ResponseWriter, r *http.Request) {

}

func getWebsites(w http.ResponseWriter, r *http.Request) {
	au, err := uauth.GetAuthenticatedUser(r)
	if err != nil {
		api.AuthError(w)
		return
	}
	wb, err := GetUserWebsites(au.UserAccountID)

	if err != nil {
		api.ServerError(w)
		return
	}

	api.ObjectResponse(w, wb)
}

func AddApiRoutes(r *mux.Router) {
	r.HandleFunc("/api/website/save", saveWebsite)
	r.HandleFunc("/api/websites", getWebsites)
}
