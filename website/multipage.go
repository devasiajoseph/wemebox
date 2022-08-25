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

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/file"
)

func GetDomain(r *http.Request) string {
	hostName := r.Host
	domain := strings.Split(hostName, ":")[0]
	domain = strings.Replace(domain, "www.", "", -1)
	return domain
}

func DomainDir(r *http.Request) string {
	domain := GetDomain(r)
	if domain == "localhost" {
		return core.Config.LocalHost
	}
	return GetDomain(r)
}

func MultiDir(r *http.Request) string {
	return BinPath + "/sites/html/" + DomainDir(r)
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

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 - Page not found"))
}

func RenderMultiPageTemplate(w http.ResponseWriter, r *http.Request, pd PageData) {
	paths := Paths{DirPath: MultiDir(r), StaticUrl: MultiStaticUrl(r)}
	pagePath := MultiPagePath(paths.DirPath, pd.PageFile)
	basePagePath := MultiPagePath(paths.DirPath, pd.BasePageFile)
	pd.StaticUrl = MultiStaticUrl(r)
	if !file.FileExist(basePagePath) {
		PageNotFound(w, r)
		return
	}
	if !file.FileExist(pagePath) {
		PageNotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(basePagePath, pagePath)
	if err != nil {
		log.Println("Template error")
		log.Println(err)
	}

	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println(err)
		log.Println("Template exe error")
	}
}
