/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, January 2019
 */

package website

import (
	"net/http"
	"strings"
)

func DomainDir(r *http.Request) string {
	hostName := r.Host
	domain := strings.Replace(hostName, "www.", "", -1)
	domain = strings.Replace(domain, ":8080", "", -1)
	return domain
}

func MultiDir(r *http.Request) string {
	return BinPath + "/sites/" + DomainDir(r)
}

func MultiStaticUrl(r *http.Request) string {
	return "/sites/" + DomainDir(r) + "/static"
}

func GetMultiPaths(r *http.Request) Paths {
	return Paths{DirPath: MultiDir(r), StaticUrl: MultiStaticUrl(r)}
}
