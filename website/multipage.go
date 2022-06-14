/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, January 2019
 */

package website

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

func GetDomain(r *http.Request) string {
	hostName := r.Host
	domain := strings.Replace(hostName, "www.", "", -1)
	domain = strings.Replace(domain, ":8080", "", -1)
	return domain

}

func DomainDir(r *http.Request) string {
	hostName := r.Host
	domain := strings.Replace(hostName, "www.", "", -1)
	domain = strings.Replace(domain, ":8080", "", -1)
	return domain
}

func MultiDir(r *http.Request) string {
	return BinPath + "/html/" + DomainDir(r)
}

func MultiStaticUrl(r *http.Request) string {
	return "/static/" + DomainDir(r)
}

func GetMultiPaths(r *http.Request) Paths {
	return Paths{DirPath: MultiDir(r), StaticUrl: MultiStaticUrl(r)}
}

func MultiPagePath(dirPath string, spage string) string {
	return FixPathSlash(dirPath + "/" + spage)
}

func RenderMultiPageTemplate(w http.ResponseWriter, r *http.Request, pd PageData) {
	//pageData := pd
	paths := Paths{DirPath: MultiDir(r), StaticUrl: MultiStaticUrl(r)}
	pagePath := MultiPagePath(paths.DirPath, pd.PageFile)
	basePagePath := MultiPagePath(paths.DirPath, pd.BasePageFile)
	pd.StaticUrl = MultiStaticUrl(r)
	tmpl, err := template.ParseFiles(basePagePath, pagePath)
	if err != nil {
		log.Println("Template error")
	}

	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println(err)
		log.Println("Template exe error")
	}
}
