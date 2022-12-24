/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, August 2019
 */

package location

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/website"
	"github.com/gorilla/mux"
)

func getCountries(w http.ResponseWriter, r *http.Request) {

	objList, err := FetchCountries()
	w.Header().Set("Content-Type", "application/json")
	responseJson, err := json.Marshal(objList)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return

}

func getStates(w http.ResponseWriter, r *http.Request) {
	country_id := r.URL.Query().Get("country_id")

	objList, err := FetchStates(country_id)
	w.Header().Set("Content-Type", "application/json")
	responseJson, err := json.Marshal(objList)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return

}

func getDistricts(w http.ResponseWriter, r *http.Request) {
	state_id, err := strconv.Atoi(r.URL.Query().Get("state_id"))
	if err != nil {
		log.Println("Invalid State Id")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	objList, err := FetchDistricts(state_id)
	w.Header().Set("Content-Type", "application/json")
	responseJson, err := json.Marshal(objList)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return

}

func getCities(w http.ResponseWriter, r *http.Request) {
	districtID, err := strconv.Atoi(r.URL.Query().Get("district_id"))
	if err != nil {
		log.Println("Invalid District ID")
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	objList, err := FetchCities(districtID)
	w.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(objList)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJSON)
	return

}

func getLocalBodyTypes(w http.ResponseWriter, r *http.Request) {
	country_id := r.URL.Query().Get("country_id")

	objList, err := FetchLocalBodyTypes(country_id)
	w.Header().Set("Content-Type", "application/json")
	responseJson, err := json.Marshal(objList)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return

}

func getLocalBodies(w http.ResponseWriter, r *http.Request) {
	district_id, err := strconv.Atoi(r.URL.Query().Get("district_id"))

	if err != nil {
		log.Println("Invalid district Id")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	local_body_type_id, err := strconv.Atoi(r.URL.Query().Get("local_body_type_id"))

	if err != nil {
		log.Println("Invalid local body type Id")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	objList, err := FetchLocalBodies(district_id, local_body_type_id)
	w.Header().Set("Content-Type", "application/json")
	responseJson, err := json.Marshal(objList)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return

}

type CityOptions struct {
	Cities []City
}

func cityOptions(w http.ResponseWriter, r *http.Request) {
	districtID, err := strconv.Atoi(r.URL.Query().Get("district_id"))
	if err != nil {
		log.Println("Invalid District ID")
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	objList, err := FetchCities(districtID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	co := CityOptions{Cities: objList}
	paths := website.Paths{DirPath: core.BinPath, StaticUrl: core.StaticUrl}
	tmpl, err := template.ParseFiles(website.PagePath(paths.DirPath, "cities.html"))
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, co)
}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/location/countries", getCountries).Methods("GET")
	r.HandleFunc("/api/location/states", getStates).Methods("GET")
	r.HandleFunc("/api/location/districts", getDistricts).Methods("GET")
	r.HandleFunc("/api/location/cities", getCities).Methods("GET")
	r.HandleFunc("/api/location/localbodytypes", getLocalBodyTypes).Methods("GET")
	r.HandleFunc("/api/location/localbodies", getLocalBodies).Methods("GET")
	r.HandleFunc("/html/cities", cityOptions).Methods("GET")

}

func Start(r *mux.Router) {
	AddRoutes(r)
}
