package post

import (
	"errors"
	"net/http"

	"github.com/devasiajoseph/wemebox/api"
	"github.com/gorilla/mux"
)

func (ol *List) Init(r *http.Request) {
	ol.Data = make([]Object, 0)
	ol.Page = api.Page(r)
	ol.Limit = api.Limit(r)
	ol.DomainID = api.QueryInt(r, "domain_id")
}

func ListObjects(w http.ResponseWriter, r *http.Request) {
	var l List
	l.Init(r)
	if l.DomainID == 0 {
		api.Error(w, errors.New("invalid domain id"))
		return
	}
	err := l.Fetch()
	if err != nil {
		api.ServerError(w)
		return
	}

	api.ObjectResponse(w, l)
}

func GetObject(w http.ResponseWriter, r *http.Request) {
	var obj Object
	obj.PostID = api.ObjID(r, "post_id")
	if obj.PostID == 0 {
		api.Error(w, errors.New("invalid domain id"))
		return
	}
	err := obj.Data()
	if err != nil {
		api.ServerError(w)
		return
	}
	if obj.PostID == 0 {
		api.ObjectNotFound(w)
		return
	}
	api.ObjectResponse(w, obj)
}

func DeleteObject(w http.ResponseWriter, r *http.Request) {

}

func SaveObject(w http.ResponseWriter, r *http.Request) {

}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/post", ListObjects).Methods("GET")
	r.HandleFunc("/api/post/{post_id}", GetObject).Methods("GET")

}

//Start initializes uauth based functions
func Start(r *mux.Router) {
	AddRoutes(r)
}
