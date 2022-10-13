package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"time"

	"github.com/devasiajoseph/wemebox/validator"
	"github.com/gorilla/mux"
)

//ResponseObj has generic response structure
type ResponseObj struct {
	Message string
	Code    int
	Valid   bool
}

//ObjID Fetches obj id from the request
func ObjID(r *http.Request, param string) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[param])
	if err != nil {
		return 0
	}
	return id
}

//VarsInt fetches mux var value as integer from the request
func VarsInt(r *http.Request, param string) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[param])
	if err != nil {
		return 0
	}
	return id
}

//Vars fetches mux vars value from request
func Vars(r *http.Request, param string) string {
	vars := mux.Vars(r)
	return vars[param]
}

//QueryInt fetches query parameter value as integer from the request
func QueryInt(r *http.Request, param string) int {
	id, err := strconv.Atoi(r.URL.Query().Get(param))
	if err != nil {
		return 0
	}
	return id

}

//QueryInt fetches query parameter value as integer from the request
func Page(r *http.Request) int {
	val, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		return 1
	}
	return val
}

func Limit(r *http.Request) int {
	val, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		return 50
	}
	return val
}

//PostInt fetches post parameter value as integer from the request
func PostInt(r *http.Request, param string) int {
	id, err := strconv.Atoi(r.FormValue(param))
	if err != nil {
		return 0
	}
	return id

}

//QueryParam fetches query parameter value from the request
func QueryParam(r *http.Request, param string) string {
	return r.URL.Query().Get(param)
}

//ObjQueryID Fetches obj id from the request
func ObjQueryID(r *http.Request, param string) int {
	id, err := strconv.Atoi(r.URL.Query().Get(param))
	if err != nil {
		return 0
	}
	return id
}

//Error Generic error response
func Error(w http.ResponseWriter, err error) {
	aE := ResponseObj{
		Message: err.Error(),
		Code:    503,
	}
	rj, err := json.Marshal(aE)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write(rj)
}

//ResponseError generic response
func ResponseError(w http.ResponseWriter, resObj ResponseObj) {
	rj, err := json.Marshal(resObj)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resObj.Code)
	w.Write(rj)
}

//GenericResponse returns a generic response json
func GenericResponse(w http.ResponseWriter, resObj ResponseObj) {
	rj, err := json.Marshal(resObj)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resObj.Code)
	w.Write(rj)
}

//GenericError returns a generic Error response json
func GenericError(w http.ResponseWriter) {
	resObj := ResponseObj{Code: 400, Message: "Generic Error"}
	rj, err := json.Marshal(resObj)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resObj.Code)
	w.Write(rj)
}

func ObjectNotFound(w http.ResponseWriter) {
	resObj := ResponseObj{Code: 404, Message: "Object Not Found"}
	rj, err := json.Marshal(resObj)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resObj.Code)
	w.Write(rj)
}

//ServerError 500
func ServerError(w http.ResponseWriter) {
	http.Error(w, "Server Error", http.StatusInternalServerError)
	return
}

//ObjectResponse generic response json object
func ObjectResponse(w http.ResponseWriter, resObj interface{}) {
	rj, err := json.Marshal(resObj)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(200)
	w.Write(rj)
}

//ValidationError returns validation error response
func ValidationError(w http.ResponseWriter, vRes validator.ValidatorResponse) {
	w.Header().Set("Content-Type", "application/json")
	responseJSON, _ := json.Marshal(vRes)
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(responseJSON)
	return
}

//RequestComplete acknowledges a request
func RequestComplete(w http.ResponseWriter, message string) {
	aE := ResponseObj{
		Message: message,
		Code:    200,
	}
	rj, err := json.Marshal(aE)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rj)
}

//AuthError authentication response error
func AuthError(w http.ResponseWriter) {
	aE := ResponseObj{
		Message: "Access denied",
		Code:    401,
	}
	rj, err := json.Marshal(aE)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	//http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	w.Write(rj)
}

//ObjIDError response if a bad object id is passed
func ObjIDError(w http.ResponseWriter) {
	aE := ResponseObj{
		Message: "Invalid ID",
		Code:    400,
	}
	rj, err := json.Marshal(aE)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write(rj)
}

//Response standard api response
func Response(rj []byte, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.Write(rj)
		return
	}
	Error(w, err)
}

var defaultLimit = 50

//PageLimit gets page and limit from the request
func PageLimit(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = defaultLimit
	}
	return page, limit
}

var myClient = &http.Client{Timeout: 5 * time.Second}

//GetJson gets and parse json
func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
