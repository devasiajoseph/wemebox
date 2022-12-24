/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, August 2019
 */

package location

import (
	//"errors"
	//"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/devasiajoseph/wemebox/db/postgres"
	"github.com/devasiajoseph/wemebox/settings"

	//"github.com/devasiajoseph/golibs/validator"
	"log"

	_ "github.com/lib/pq"
)

var CountrySettings = "country"

//Country data
type Country struct {
	CountryID        string `db:"country_id" json:"country_id"`
	CountryName      string `db:"country_name" json:"country_name"`
	CountryPhoneCode string `db:"country_phone_code" json:"country_phone_code"`
}

//State data
type State struct {
	StateID   int    `db:"state_id" json:"state_id"`
	StateName string `db:"state_name" json:"state_name"`
	CountryID string `db:"country_id" json:"country_id"`
}

//District data
type District struct {
	DistrictID   string `db:"district_id" json:"district_id"`
	DistrictName string `db:"district_name" json:"district_name"`
	StateID      int    `db:"state_id" json:"state_id"`
}

//City town village data
type City struct {
	CityID     int    `db:"city_id" json:"city_id"`
	CityName   string `db:"city_name" json:"city_name"`
	CityType   string `db:"city_type" json:"city_type"`
	DistrictID int    `db:"district_id" json:"district_id"`
	Latitude   string `db:"latitude" json:"latitude"`
	Longitude  string `db:"longitude" json:"longitude"`
}

type LocalBodyType struct {
	LocalBodyTypeId   int    `db:"local_body_type_id" json:"LocalBodyTypeId"`
	LocalBodyTypeName string `db:"local_body_type_name" json:"LocalBodyTypeName"`
	CountryID         string `db:"country_id" json:"country_id"`
}

type LocalBody struct {
	LocalBodyId     int    `db:"local_body_id" json:"LocalBodyId"`
	LocalBodyName   string `db:"local_body_name" json:"LocalBodyName"`
	LocalBodyTypeId int    `db:"local_body_type_id" json:"LocalBodyTypeId"`
	DistrictId      int    `db:"district_id" json:"DistrictId"`
}

//FetchStates gets list of states
func FetchStates(countryID string) ([]State, error) {
	db := postgres.Db
	objs := []State{}
	err := db.Select(&objs, sqlSelectState, countryID)
	if err != nil {
		log.Println("Error fetching states")
		return objs, err
	}

	return objs, err
}

//FetchCountries gets list of countries
func FetchCountries() ([]Country, error) {
	db := postgres.Db
	objs := []Country{}
	err := db.Select(&objs, sqlSelectCountry)
	if err != nil {
		log.Println("Error fetching countries")
		return objs, err
	}

	return objs, err
}

func FetchLocalBodyTypes(countryID string) ([]LocalBodyType, error) {
	db := postgres.Db
	objs := []LocalBodyType{}
	err := db.Select(&objs, sqlSelectLocalBodyType, countryID)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching local body types")
		return objs, err
	}

	return objs, err
}

func FetchDistricts(stateID int) ([]District, error) {
	db := postgres.Db
	objs := []District{}
	err := db.Select(&objs, sqlSelectDistrict, stateID)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching dsitricts")
		return objs, err
	}

	return objs, err
}

func FetchCities(districtID int) ([]City, error) {
	db := postgres.Db
	objs := []City{}
	err := db.Select(&objs, sqlSelectCity, districtID)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching cities")
		return objs, err
	}

	return objs, err
}

func FetchLocalBodies(district_id int, local_body_type_id int) ([]LocalBody, error) {
	db := postgres.Db
	objs := []LocalBody{}
	err := db.Select(&objs, sqlSelectLocalBody, district_id, local_body_type_id)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching local bodies")
		return objs, err
	}

	return objs, err
}

func GetCity(cityID int) (City, error) {
	db := postgres.Db
	city := City{}
	err := db.Get(&city, "Select city_id,city_name,latitude,longitude from city where city_id=$1", cityID)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching city")
	}

	return city, err
}

func GetCityDistrictID(cityID int) (int, error) {
	db := postgres.Db
	dID := 0
	err := db.Get(&dID, "select district_id from city where city_id=$1;", cityID)

	if err != nil {
		log.Println("Invalid city_id error")
		log.Println(err)
	}

	return dID, err
}

func GetCountryCode(countryName string) (string, error) {
	db := postgres.Db
	var cc string
	err := db.Get(&cc, "select country_code from country where country_name=$1", countryName)
	if err != nil {
		log.Println("Error getting country code from db")
		log.Println(err)
	}

	return cc, err
}

func GetAppCountryCode() (string, error) {
	appCountry, err := settings.GetSettings(CountrySettings)

	if err != nil {
		log.Println("Error getting app country code")
		log.Println(err)
	}
	return GetCountryCode(appCountry.SettingsValue)
}

func GetDistrictID(name string) (int, error) {
	db := postgres.Db
	var id int
	err := db.Get(&id, "select district_id from district where district_name=$1", strings.TrimSpace(name))
	if err != nil {
		log.Println("Error fetching district id")
		log.Println(err)
	}

	return id, err
}

func GetCityID(name string) (int, error) {
	db := postgres.Db
	var id int
	cleaned := strings.TrimSpace(name)
	err := db.Get(&id, "select city_id from city where city_name=$1;", cleaned)
	if err != nil {
		log.Println("Error fetching city id from : " + cleaned)
		log.Println()
		log.Println(err)
	}

	return id, err
}

type CityDistrictState struct {
	CityID       int    `db:"city_id" json:"city_id"`
	CityName     string `db:"city_name" json:"city_name"`
	DistrictID   int    `db:"district_id" json:"district_id"`
	DistrictName string `db:"district_name" json:"district_name"`
	StateID      int    `db:"state_id" json:"state_id"`
	StateName    string `db:"state_name" json:"state_name"`
}

func GetCityDistrict(city string, district string) (CityDistrictState, error) {
	db := postgres.Db
	var cds CityDistrictState
	cds.CityName = strings.Title(strings.ToLower(strings.TrimSpace(city)))
	cds.DistrictName = strings.Title(strings.ToLower(strings.TrimSpace(district)))
	sqlCityDistrict := "select city.city_id,city.city_name,district.district_id,district.district_name " +
		" from city left join district  on city.district_id=district.district_id " +
		" where city.city_name=$1 and district.district_name=$2;"
	err := db.Get(&cds, sqlCityDistrict, cds.CityName, cds.DistrictName)
	if err != nil {
		log.Println(err)
		err = errors.New("Error fetching city id from : " + cds.CityName + " | " + cds.DistrictName)
	}

	return cds, err
}

func GetDistrictCentre(dID int) (City, error) {
	db := postgres.Db
	var c City
	err := db.Get(&c, "select city_id,city_name,latitude,longitude from city where district_id=$1 and district_centre=true", dID)

	if err != nil {
		fmt.Println("Error getting district centre")
		fmt.Println(err)
	}

	return c, err
}
